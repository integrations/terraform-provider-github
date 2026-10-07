package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubCodespacesSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubCodespacesSecretCreateOrUpdate,
		ReadContext:   resourceGithubCodespacesSecretRead,
		DeleteContext: resourceGithubCodespacesSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubCodespacesSecretImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository.",
			},
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"encrypted_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"plaintext_value"},
				Description:   "Encrypted value of the secret using the GitHub public key in Base64 format.",
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"encrypted_value"},
				Description:   "Plaintext value of the secret to be encrypted.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'codespaces_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'codespaces_secret' update.",
			},
		},
	}
}

func resourceGithubCodespacesSecretCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo := d.Get("repository").(string)
	secretName := d.Get("secret_name").(string)
	plaintextValue := d.Get("plaintext_value").(string)
	var encryptedValue string

	keyId, publicKey, err := getCodespacesPublicKeyDetails(ctx, owner, repo, meta)
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

	_, err = client.Codespaces.CreateOrUpdateRepoSecret(ctx, owner, repo, eSecret)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(repo, secretName))
	return resourceGithubCodespacesSecretRead(ctx, d, meta)
}

func resourceGithubCodespacesSecretRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, secretName, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	secret, _, err := client.Codespaces.GetRepoSecret(ctx, owner, repoName, secretName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Warn(ctx, "Removing Codespaces secret from state because it no longer exists in GitHub", map[string]any{"secret_id": d.Id(), "owner": owner, "repository": repoName, "secret_name": secretName})
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
	// timestamp we've persisted in the state. In that case, we can no longer
	// trust that the value (which we don't see) is equal to what we've declared
	// previously.
	//
	// The only solution to enforce consistency between is to mark the resource
	// as deleted (unset the ID) in order to fix potential drift by recreating
	// the resource.
	if updatedAt, ok := d.GetOk("updated_at"); ok && updatedAt != secret.UpdatedAt.String() {
		tflog.Warn(ctx, "Secret has been externally updated in GitHub", map[string]any{"secret_id": d.Id(), "owner": owner, "repository": repoName, "secret_name": secretName})
		d.SetId("")
	} else if !ok {
		if err = d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubCodespacesSecretDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx = context.WithValue(ctx, ctxId, d.Id())

	repoName, secretName, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, "Deleting secret", map[string]any{"secret_id": d.Id(), "owner": orgName, "repository": repoName, "secret_name": secretName})
	_, err = client.Codespaces.DeleteRepoSecret(ctx, orgName, repoName, secretName)

	return diag.FromErr(err)
}

func resourceGithubCodespacesSecretImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <repository>/<secret_name>")
	}

	d.SetId(buildTwoPartID(parts[0], parts[1]))

	repoName, secretName, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}

	secret, _, err := client.Codespaces.GetRepoSecret(ctx, owner, repoName, secretName)
	if err != nil {
		return nil, err
	}

	if err = d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err = d.Set("secret_name", secretName); err != nil {
		return nil, err
	}

	// encrypted_value or plaintext_value can not be imported

	if err = d.Set("created_at", secret.CreatedAt.String()); err != nil {
		return nil, err
	}
	if err = d.Set("updated_at", secret.UpdatedAt.String()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func getCodespacesPublicKeyDetails(ctx context.Context, owner, repository string, meta any) (keyId, pkValue string, err error) {
	client := meta.(*Owner).v3client

	publicKey, _, err := client.Codespaces.GetRepoPublicKey(ctx, owner, repository)
	if err != nil {
		return keyId, pkValue, err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), err
}
