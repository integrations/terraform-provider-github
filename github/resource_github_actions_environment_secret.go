package github

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsEnvironmentSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsEnvironmentSecretCreateOrUpdate,
		ReadContext:   resourceGithubActionsEnvironmentSecretRead,
		DeleteContext: resourceGithubActionsEnvironmentSecretDelete,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
			"encrypted_value": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				Description:      "Encrypted value of the secret using the GitHub public key in Base64 format.",
				ConflictsWith:    []string{"plaintext_value"},
				ValidateDiagFunc: toDiagFunc(validation.StringIsBase64, "encrypted_value"),
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Sensitive:     true,
				Description:   "Plaintext value of the secret to be encrypted.",
				ConflictsWith: []string{"encrypted_value"},
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
		},
	}
}

func resourceGithubActionsEnvironmentSecretCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	secretName := d.Get("secret_name").(string)
	plaintextValue := d.Get("plaintext_value").(string)
	var encryptedValue string

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	keyId, publicKey, err := getEnvironmentPublicKeyDetails(ctx, repo.GetID(), url.PathEscape(envName), meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if encryptedText, ok := d.GetOk("encrypted_value"); ok {
		encryptedValue = encryptedText.(string)
	} else {
		encryptedBytes, err := encryptPlaintext(plaintextValue, publicKey)
		if err != nil {
			return diag.FromErr(err)
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	// Create an EncryptedSecret and encrypt the plaintext value into it
	eSecret := &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          keyId,
		EncryptedValue: encryptedValue,
	}

	_, err = client.Actions.CreateOrUpdateEnvSecret(ctx, int(repo.GetID()), url.PathEscape(envName), eSecret)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), secretName)
	if err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	return resourceGithubActionsEnvironmentSecretRead(ctx, d, meta)
}

func resourceGithubActionsEnvironmentSecretRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, envNamePart, secretName, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	envName := unescapeIDPart(envNamePart)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing environment secret %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	secret, _, err := client.Actions.GetEnvSecret(ctx, int(repo.GetID()), url.PathEscape(envName), secretName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing environment secret %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("encrypted_value", d.Get("encrypted_value")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("plaintext_value", d.Get("plaintext_value")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("created_at", secret.CreatedAt.String()); err != nil {
		return diag.FromErr(err)
	}

	// This is a drift detection mechanism based on timestamps.
	//
	// If we do not currently store the "updated_at" field, it means we've only
	// just created the resource and the value is most likely what we want it to
	// be.
	//
	// If the resource is changed externally in the meantime then reading back
	// the last update timestamp will return a result different than the
	// timestamp we've persisted in the state. In this case, we can no longer
	// trust that the value matches what is in the state file.
	//
	// To solve this, we must unset the values and allow Terraform to decide whether or
	// not this resource should be modified or left as-is (ignore_changes).
	if updatedAt, ok := d.GetOk("updated_at"); ok && updatedAt != secret.UpdatedAt.String() {
		log.Printf("[INFO] The environment secret %s has been externally updated in GitHub", d.Id())
		_ = d.Set("encrypted_value", "")
		_ = d.Set("plaintext_value", "")
	} else if !ok {
		if err = d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsEnvironmentSecretDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, envNamePart, secretName, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	envName := unescapeIDPart(envNamePart)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] Deleting environment secret: %s", d.Id())
	_, err = client.Actions.DeleteEnvSecret(ctx, int(repo.GetID()), url.PathEscape(envName), secretName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func getEnvironmentPublicKeyDetails(ctx context.Context, repoID int64, envNameEscaped string, meta any) (string, string, error) {
	client := meta.(*Owner).v3client

	publicKey, _, err := client.Actions.GetEnvPublicKey(ctx, int(repoID), envNameEscaped)
	if err != nil {
		return "", "", err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), nil
}
