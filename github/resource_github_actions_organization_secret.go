package github

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsOrganizationSecretCreate,
		ReadContext:   resourceGithubActionsOrganizationSecretRead,
		UpdateContext: resourceGithubActionsOrganizationSecretUpdate,
		DeleteContext: resourceGithubActionsOrganizationSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationSecretImport,
		},

		CustomizeDiff: customdiff.All(
			diffSecret,
			diffSecretVariableVisibility,
		),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsOrganizationSecretV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsOrganizationSecretStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to manage a GitHub Actions secret for an organization.",

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"key_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				RequiredWith:  []string{"value_encrypted"},
				ConflictsWith: []string{"value", "plaintext_value"},
				Description:   "ID of the public key used to encrypt the secret.",
			},
			"value": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"value", "value_encrypted", "encrypted_value", "plaintext_value"},
				Description:  "Plaintext value to be encrypted.",
			},
			"value_encrypted": {
				Type:             schema.TypeString,
				Optional:         true,
				Sensitive:        true,
				ExactlyOneOf:     []string{"value", "value_encrypted", "encrypted_value", "plaintext_value"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
				Description:      "Value encrypted with the GitHub public key, defined by key_id, in Base64 format.",
			},
			"encrypted_value": {
				Type:             schema.TypeString,
				Optional:         true,
				Sensitive:        true,
				ExactlyOneOf:     []string{"value", "value_encrypted", "encrypted_value", "plaintext_value"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
				Description:      "Encrypted value of the secret using the GitHub public key in Base64 format.",
				Deprecated:       "Use value_encrypted and key_id.",
			},
			"plaintext_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"value", "value_encrypted", "encrypted_value", "plaintext_value"},
				Description:  "Plaintext value of the secret to be encrypted.",
				Deprecated:   "Use value.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "private", "selected"}, false)),
				Description:      "Configures the access that repositories have to the organization secret. Must be one of 'all', 'private', or 'selected'.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Set:  schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "An array of repository IDs that can access the organization secret.",
				Deprecated:  "This field is deprecated and will be removed in a future release. Please use the `github_actions_organization_secret_repositories` or `github_actions_organization_secret_repository` resources to manage repository access to organization secrets.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp for when the secret was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp for when the secret was last updated by the provider.",
			},
			"remote_updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp for when the secret was last updated.",
			},
			"destroy_on_drift": {
				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "This is no longer required and will be removed in a future release. Drift detection is now always performed, and external changes will result in the secret being updated to match the Terraform configuration. If you want to ignore external changes, you can use the `lifecycle` block with `ignore_changes` on the `updated_at` field.",
			},
		},
	}
}

func resourceGithubActionsOrganizationSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)
	keyID := d.Get("key_id").(string)
	encryptedValue, _ := resourceKeysGetOk[string](d, "value_encrypted", "encrypted_value")
	visibility := d.Get("visibility").(string)

	var repoIDs []int64

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		ki, pk, err := getOrganizationPublicKeyDetails(ctx, meta)
		if err != nil {
			return diag.FromErr(err)
		}

		keyID = ki
		publicKey = pk
	}

	if len(encryptedValue) == 0 {
		plaintextValue, _ := resourceKeysGetOk[string](d, "value", "plaintext_value")

		encryptedBytes, err := encryptPlaintext(plaintextValue, publicKey)
		if err != nil {
			return diag.FromErr(err)
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	secretReq := github.OrgSecretRequest{
		KeyID:                 keyID,
		EncryptedValue:        encryptedValue,
		Visibility:            visibility,
		SelectedRepositoryIDs: repoIDs,
	}

	if _, err := client.Actions.CreateOrUpdateOrgSecret(ctx, owner, secretName, secretReq); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(secretName)

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on create so we have to lookup the secret to get timestamps.
	if secret, err := retryUntilResourceFound(ctx, func() (*github.Secret, error) {
		val, _, err := client.Actions.GetOrgSecret(ctx, owner, secretName)
		return val, err
	}, nil); err == nil {
		if err := d.Set("created_at", secret.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("remote_updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsOrganizationSecretRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName, _ := d.Get("secret_name").(string)

	secret, _, err := client.Actions.GetOrgSecret(ctx, owner, secretName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			log.Printf("[INFO] Removing actions organization secret %s from state because it no longer exists in GitHub", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	// Due to the eventually consistent behavior of this API we may not get created_at/updated_at
	// values on the first read after creation, so we only set them here if they are not already set.
	if len(d.Get("created_at").(string)) == 0 {
		if err = d.Set("created_at", secret.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(d.Get("updated_at").(string)) == 0 {
		if err = d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}
	if err = d.Set("remote_updated_at", secret.UpdatedAt.String()); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("created_at", secret.CreatedAt.String()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("visibility", secret.Visibility); err != nil {
		return diag.FromErr(err)
	}

	if secret.Visibility == "selected" {
		if _, ok := d.GetOk("selected_repository_ids"); ok {
			var repoIDs []int64
			for repo, err := range client.Actions.ListSelectedReposForOrgSecretIter(ctx, owner, secretName, &github.ListOptions{PerPage: maxPerPage}) {
				if err != nil {
					return diag.FromErr(err)
				}

				repoIDs = append(repoIDs, repo.GetID())
			}

			if err := d.Set("selected_repository_ids", repoIDs); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubActionsOrganizationSecretUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName, _ := d.Get("secret_name").(string)
	keyID, _ := d.Get("key_id").(string)
	encryptedValue, _ := resourceKeysGetOk[string](d, "value_encrypted", "encrypted_value")
	visibility, _ := d.Get("visibility").(string)

	var repoIDs []int64

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		ki, pk, err := getOrganizationPublicKeyDetails(ctx, meta)
		if err != nil {
			return diag.FromErr(err)
		}

		keyID = ki
		publicKey = pk
	}

	if len(encryptedValue) == 0 {
		plaintextValue, _ := resourceKeysGetOk[string](d, "value", "plaintext_value")

		encryptedBytes, err := encryptPlaintext(plaintextValue, publicKey)
		if err != nil {
			return diag.FromErr(err)
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	secretReq := github.OrgSecretRequest{
		KeyID:                 keyID,
		EncryptedValue:        encryptedValue,
		Visibility:            visibility,
		SelectedRepositoryIDs: repoIDs,
	}

	if _, err := client.Actions.CreateOrUpdateOrgSecret(ctx, owner, secretName, secretReq); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on update so we have to lookup the secret to get timestamps.
	if secret, err := retryUntilResourceFound(ctx, func() (*github.Secret, error) {
		val, _, err := client.Actions.GetOrgSecret(ctx, owner, secretName)
		return val, err
	}, nil); err == nil {
		if err := d.Set("created_at", secret.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("remote_updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsOrganizationSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName, _ := d.Get("secret_name").(string)

	tflog.Info(ctx, "Deleting actions organization secret", map[string]any{"secret_name": secretName})

	if _, err := client.Actions.DeleteOrgSecret(ctx, owner, secretName); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationSecretImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Id()

	secret, _, err := client.Actions.GetOrgSecret(ctx, owner, secretName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("secret_name", secretName); err != nil {
		return nil, err
	}
	if err := d.Set("visibility", secret.Visibility); err != nil {
		return nil, err
	}
	if err := d.Set("created_at", secret.CreatedAt.String()); err != nil {
		return nil, err
	}
	if err := d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
		return nil, err
	}
	if err := d.Set("remote_updated_at", secret.UpdatedAt.String()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func getOrganizationPublicKeyDetails(ctx context.Context, meta *Owner) (string, string, error) {
	client := meta.v3client
	owner := meta.name

	publicKey, _, err := client.Actions.GetOrgPublicKey(ctx, owner)
	if err != nil {
		return "", "", err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), err
}
