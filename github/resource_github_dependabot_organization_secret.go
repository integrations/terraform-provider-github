package github

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubDependabotOrganizationSecret() *schema.Resource {
	return &schema.Resource{
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
				Description:      "Configures the access that repositories have to the organization secret. Must be one of 'all', 'private' or 'selected'. 'selected_repository_ids' is required if set to 'selected'.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Set:  schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "An array of repository ids that can access the organization secret.",
				Deprecated:  "This field is deprecated and will be removed in a future release. Please use the `github_dependabot_organization_secret_repositories` or `github_dependabot_organization_secret_repository` resources to manage repository access to organization secrets.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'dependabot_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'dependabot_secret' update.",
			},
			"remote_updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of secret update at the remote.",
			},
		},

		CustomizeDiff: customdiff.All(
			diffSecret,
			diffSecretVariableVisibility,
		),

		CreateContext: resourceGithubDependabotOrganizationSecretCreate,
		ReadContext:   resourceGithubDependabotOrganizationSecretRead,
		UpdateContext: resourceGithubDependabotOrganizationSecretUpdate,
		DeleteContext: resourceGithubDependabotOrganizationSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubDependabotOrganizationSecretImport,
		},
	}
}

func resourceGithubDependabotOrganizationSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)
	keyID := d.Get("key_id").(string)
	encryptedValue, _ := resourceKeysGetOk[string](d, "value_encrypted", "encrypted_value")
	visibility := d.Get("visibility").(string)
	repoIDs := github.DependabotSecretsSelectedRepoIDs{}

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		ki, pk, err := getDependabotOrganizationPublicKeyDetails(ctx, meta)
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

	secret := github.DependabotEncryptedSecret{
		Name:                  secretName,
		KeyID:                 keyID,
		EncryptedValue:        encryptedValue,
		Visibility:            visibility,
		SelectedRepositoryIDs: repoIDs,
	}

	_, err := client.Dependabot.CreateOrUpdateOrgSecret(ctx, owner, &secret)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(secretName)

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on create so we have to lookup the secret to get timestamps
	if secret, _, err := client.Dependabot.GetOrgSecret(ctx, owner, secretName); err == nil {
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

func resourceGithubDependabotOrganizationSecretRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)

	secret, _, err := client.Dependabot.GetOrgSecret(ctx, owner, secretName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing Dependabot organization secret %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
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
			repoIDs := []int64{}
			opt := &github.ListOptions{
				PerPage: maxPerPage,
			}
			for {
				results, resp, err := client.Dependabot.ListSelectedReposForOrgSecret(ctx, owner, secretName, opt)
				if err != nil {
					return diag.FromErr(err)
				}

				for _, repo := range results.Repositories {
					repoIDs = append(repoIDs, repo.GetID())
				}

				if resp.NextPage == 0 {
					break
				}
				opt.Page = resp.NextPage
			}

			if err := d.Set("selected_repository_ids", repoIDs); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubDependabotOrganizationSecretUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)
	keyID := d.Get("key_id").(string)
	encryptedValue, _ := resourceKeysGetOk[string](d, "value_encrypted", "encrypted_value")
	visibility := d.Get("visibility").(string)
	repoIDs := github.DependabotSecretsSelectedRepoIDs{}

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		ki, pk, err := getDependabotOrganizationPublicKeyDetails(ctx, meta)
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

	secret := github.DependabotEncryptedSecret{
		Name:                  secretName,
		KeyID:                 keyID,
		EncryptedValue:        encryptedValue,
		Visibility:            visibility,
		SelectedRepositoryIDs: repoIDs,
	}

	_, err := client.Dependabot.CreateOrUpdateOrgSecret(ctx, owner, &secret)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on update so we have to lookup the secret to get timestamps
	if secret, _, err := client.Dependabot.GetOrgSecret(ctx, owner, secretName); err == nil {
		if err := d.Set("created_at", secret.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("remote_updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("updated_at", nil); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("remote_updated_at", nil); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubDependabotOrganizationSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)

	log.Printf("[INFO] Deleting Dependabot organization secret: %s", d.Id())
	_, err := client.Dependabot.DeleteOrgSecret(ctx, owner, secretName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubDependabotOrganizationSecretImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Id()

	secret, _, err := client.Dependabot.GetOrgSecret(ctx, owner, secretName)
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

	selectedRepositoryIDs := []int64{}
	if secret.Visibility == "selected" {
		opt := &github.ListOptions{
			PerPage: maxPerPage,
		}
		for {
			results, resp, err := client.Dependabot.ListSelectedReposForOrgSecret(ctx, owner, secretName, opt)
			if err != nil {
				return nil, err
			}

			for _, repo := range results.Repositories {
				selectedRepositoryIDs = append(selectedRepositoryIDs, repo.GetID())
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}

	if err := d.Set("selected_repository_ids", selectedRepositoryIDs); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func getDependabotOrganizationPublicKeyDetails(ctx context.Context, meta *Owner) (string, string, error) {
	client := meta.v3client
	owner := meta.name

	publicKey, _, err := client.Dependabot.GetOrgPublicKey(ctx, owner)
	if err != nil {
		return "", "", err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), err
}
