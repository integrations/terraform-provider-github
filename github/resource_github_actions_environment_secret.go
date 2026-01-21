package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsEnvironmentSecret() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsEnvironmentSecretV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsEnvironmentSecretStateUpgradeV0,
				Version: 0,
			},
		},

		CustomizeDiff: resourceGithubActionsEnvironmentSecretDiff,
		CreateContext: resourceGithubActionsEnvironmentSecretCreate,
		ReadContext:   resourceGithubActionsEnvironmentSecretRead,
		UpdateContext: resourceGithubActionsEnvironmentSecretUpdate,
		DeleteContext: resourceGithubActionsEnvironmentSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsEnvironmentSecretImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
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
				Description:   "ID of the public key used to encrypt the secret.",
				ConflictsWith: []string{"plaintext_value"},
			},
			"encrypted_value": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Encrypted value of the secret using the GitHub public key in Base64 format.",
				ExactlyOneOf:     []string{"encrypted_value", "plaintext_value"},
				ValidateDiagFunc: toDiagFunc(validation.StringIsBase64, "encrypted_value"),
			},
			"plaintext_value": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				Description:  "Plaintext value of the secret to be encrypted.",
				ExactlyOneOf: []string{"encrypted_value", "plaintext_value"},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_environment_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_environment_secret' update.",
			},
			"remote_updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of remote 'actions_environment_secret' update.",
			},
		},
	}
}

func resourceGithubActionsEnvironmentSecretDiff(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	if len(diff.Id()) == 0 {
		return nil
	}

	if diff.HasChange("repository") {
		repoIDString, _, _, _, err := parseID4(diff.Id())
		if err != nil {
			return err
		}

		repoID, err := strconv.Atoi(repoIDString)
		if err != nil {
			return fmt.Errorf("failed to convert repository ID %s to integer: %w", repoIDString, err)
		}

		repoName := diff.Get("repository").(string)

		repo, _, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) {
				if ghErr.Response.StatusCode != http.StatusNotFound {
					return err
				}

				log.Printf("[INFO] Repository %s not found when checking repository change for actions environment secret %s", repoName, diff.Id())
			} else {
				return err
			}
		} else {
			log.Printf("[INFO] Repository %s found when checking repository change for actions environment secret %s", repoName, diff.Id())

			if repoID != int(repo.GetID()) {
				return diff.ForceNew("repository")
			}
		}
	}

	if diff.HasChange("remote_updated_at") {
		remoteUpdatedAt := diff.Get("remote_updated_at").(string)
		if len(remoteUpdatedAt) == 0 {
			return nil
		}

		updatedAt := diff.Get("updated_at").(string)
		if updatedAt != remoteUpdatedAt {
			if len(updatedAt) == 0 {
				return diff.SetNew("updated_at", remoteUpdatedAt)
			}

			return diff.SetNewComputed("updated_at")
		}
	}

	return nil
}

func resourceGithubActionsEnvironmentSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	secretName := d.Get("secret_name").(string)
	keyID := d.Get("key_id").(string)
	encryptedValue := d.Get("encrypted_value").(string)

	escapedEnvName := url.PathEscape(envName)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		keyID, publicKey, err = getEnvironmentPublicKeyDetails(ctx, meta, repoID, escapedEnvName)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if len(encryptedValue) == 0 {
		plaintextValue := d.Get("plaintext_value").(string)

		encryptedBytes, err := encryptPlaintext(plaintextValue, publicKey)
		if err != nil {
			return diag.FromErr(err)
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	_, err = client.Actions.CreateOrUpdateEnvSecret(ctx, repoID, escapedEnvName, &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          keyID,
		EncryptedValue: encryptedValue,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if id, err := buildID(strconv.Itoa(repoID), repoName, escapeIDPart(envName), secretName); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on create so we have to lookup the secret to get timestamps
	if secret, _, err := client.Actions.GetEnvSecret(ctx, repoID, escapedEnvName, secretName); err == nil {
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
	meta := m.(*Owner)
	client := meta.v3client

	repoIDString, _, _, _, err := parseID4(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	repoID, err := strconv.Atoi(repoIDString)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to convert repository ID %s to integer: %w", repoIDString, err))
	}

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	secretName := d.Get("secret_name").(string)

	secret, _, err := client.Actions.GetEnvSecret(ctx, repoID, url.PathEscape(envName), secretName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing environment secret %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if id, err := buildID(strconv.Itoa(repoID), repoName, escapeIDPart(envName), secretName); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

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
	meta := m.(*Owner)
	client := meta.v3client

	repoIDString, _, _, _, err := parseID4(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	repoID, err := strconv.Atoi(repoIDString)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to convert repository ID %s to integer: %w", repoIDString, err))
	}

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	secretName := d.Get("secret_name").(string)
	keyID := d.Get("key_id").(string)
	encryptedValue := d.Get("encrypted_value").(string)

	escapedEnvName := url.PathEscape(envName)

	var publicKey string
	if len(keyID) == 0 || len(encryptedValue) == 0 {
		keyID, publicKey, err = getEnvironmentPublicKeyDetails(ctx, meta, repoID, escapedEnvName)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if len(encryptedValue) == 0 {
		plaintextValue := d.Get("plaintext_value").(string)

		encryptedBytes, err := encryptPlaintext(plaintextValue, publicKey)
		if err != nil {
			return diag.FromErr(err)
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	_, err = client.Actions.CreateOrUpdateEnvSecret(ctx, repoID, escapedEnvName, &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          keyID,
		EncryptedValue: encryptedValue,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if id, err := buildID(strconv.Itoa(repoID), repoName, escapeIDPart(envName), secretName); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	if err := d.Set("key_id", keyID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on update so we have to lookup the secret to get timestamps
	if secret, _, err := client.Actions.GetEnvSecret(ctx, repoID, escapedEnvName, secretName); err == nil {
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

func resourceGithubActionsEnvironmentSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	repoIDString, _, envNamePart, secretName, err := parseID4(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	repoID, err := strconv.Atoi(repoIDString)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to convert repository ID %s to integer: %w", repoIDString, err))
	}

	log.Printf("[INFO] Deleting environment secret: %s", d.Id())
	_, err = client.Actions.DeleteEnvSecret(ctx, repoID, unescapeIDPart(envNamePart), secretName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsEnvironmentSecretImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
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

	secret, _, err := client.Actions.GetEnvSecret(ctx, repoID, url.PathEscape(envName), secretName)
	if err != nil {
		return nil, err
	}

	if id, err := buildID(strconv.Itoa(repoID), repoName, envNamePart, secretName); err != nil {
		return nil, err
	} else {
		d.SetId(id)
	}

	if err := d.Set("repository", repoName); err != nil {
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

func getEnvironmentPublicKeyDetails(ctx context.Context, meta *Owner, repoID int, envNameEscaped string) (string, string, error) {
	client := meta.v3client

	publicKey, _, err := client.Actions.GetEnvPublicKey(ctx, repoID, envNameEscaped)
	if err != nil {
		return "", "", err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), nil
}
