package github

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsEnvironmentSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsEnvironmentSecretCreate,
		ReadContext:   resourceGithubActionsEnvironmentSecretRead,
		UpdateContext: resourceGithubActionsEnvironmentSecretUpdate,
		DeleteContext: resourceGithubActionsEnvironmentSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsEnvironmentSecretImport,
		},

		CustomizeDiff: customdiff.All(
			diffRepository,
			diffSecret,
		),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsEnvironmentSecretV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsEnvironmentSecretStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to manage a GitHub Actions secrets for a repository environment.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the environment.",
			},
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
				Description:   "ID of the public key used to encrypt the secret. This is required when setting `value_encrypted`.",
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
				Description:      "Value encrypted with the GitHub public key, defined by `key_id`, in Base64 format.",
			},
			"encrypted_value": {
				Type:             schema.TypeString,
				Optional:         true,
				Sensitive:        true,
				ExactlyOneOf:     []string{"value", "value_encrypted", "encrypted_value", "plaintext_value"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
				Description:      "Encrypted value of the secret using the GitHub public key in Base64 format.",
				Deprecated:       "Use `value_encrypted` and `key_id`.",
			},
			"plaintext_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"value", "value_encrypted", "encrypted_value", "plaintext_value"},
				Description:  "Plaintext value of the secret to be encrypted.",
				Deprecated:   "Use `value`.",
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
		},
	}
}

func resourceGithubActionsEnvironmentSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	secretName, _ := d.Get("secret_name").(string)
	keyID, _ := d.Get("key_id").(string)
	encryptedValue, _ := resourceKeysGetOk[string](d, "value_encrypted", "encrypted_value")

	escapedEnvName := url.PathEscape(envName)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		ki, pk, err := getEnvironmentPublicKeyDetails(ctx, meta, owner, repoName, escapedEnvName)
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

	secretReq := github.SecretRequest{
		EncryptedValue: encryptedValue,
		KeyID:          keyID,
	}

	if _, err := client.Actions.CreateOrUpdateEnvSecret(ctx, owner, repoName, escapedEnvName, secretName, secretReq); err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), secretName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("repository_id", repoID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on create so we have to lookup the secret to get timestamps.
	if secret, err := retryUntilResourceFound(ctx, func() (*github.Secret, error) {
		val, _, err := client.Actions.GetEnvSecret(ctx, owner, repoName, escapedEnvName, secretName)
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

func resourceGithubActionsEnvironmentSecretRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	secretName, _ := d.Get("secret_name").(string)

	secret, _, err := client.Actions.GetEnvSecret(ctx, owner, repoName, url.PathEscape(envName), secretName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing environment secret from state because it no longer exists in GitHub.", map[string]any{"secret_name": secretName, "environment": envName, "repository": repoName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), secretName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

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

	return nil
}

func resourceGithubActionsEnvironmentSecretUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	secretName, _ := d.Get("secret_name").(string)
	keyID, _ := d.Get("key_id").(string)
	encryptedValue, _ := resourceKeysGetOk[string](d, "value_encrypted", "encrypted_value")

	escapedEnvName := url.PathEscape(envName)

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		ki, pk, err := getEnvironmentPublicKeyDetails(ctx, meta, owner, repoName, escapedEnvName)
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

	secretReq := github.SecretRequest{
		EncryptedValue: encryptedValue,
		KeyID:          keyID,
	}

	if _, err := client.Actions.CreateOrUpdateEnvSecret(ctx, owner, repoName, escapedEnvName, secretName, secretReq); err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), secretName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on update so we have to lookup the secret to get timestamps.
	if secret, err := retryUntilResourceFound(ctx, func() (*github.Secret, error) {
		val, _, err := client.Actions.GetEnvSecret(ctx, owner, repoName, escapedEnvName, secretName)
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

func resourceGithubActionsEnvironmentSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	secretName, _ := d.Get("secret_name").(string)

	tflog.Info(ctx, "Deleting actions environment secret.", map[string]any{"secret_name": secretName, "environment": envName, "repository": repoName})
	_, err := client.Actions.DeleteEnvSecret(ctx, owner, repoName, url.PathEscape(envName), secretName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsEnvironmentSecretImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, secretName, err := parseID3(d.Id())
	if err != nil {
		return nil, err
	}

	envName := unescapeIDPart(envNamePart)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	repoID := int(repo.GetID())

	secret, _, err := client.Actions.GetEnvSecret(ctx, owner, repoName, url.PathEscape(envName), secretName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}
	if err := d.Set("environment", envName); err != nil {
		return nil, err
	}
	if err := d.Set("secret_name", secretName); err != nil {
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

func getEnvironmentPublicKeyDetails(ctx context.Context, meta *Owner, owner, repoName, envNameEscaped string) (string, string, error) {
	client := meta.v3client

	publicKey, _, err := client.Actions.GetEnvPublicKey(ctx, owner, repoName, envNameEscaped)
	if err != nil {
		return "", "", err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), nil
}
