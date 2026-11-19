package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubAgentsSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubAgentsSecretCreate,
		ReadContext:   resourceGithubAgentsSecretRead,
		UpdateContext: resourceGithubAgentsSecretUpdate,
		DeleteContext: resourceGithubAgentsSecretDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubAgentsSecretImport},
		CustomizeDiff: customdiff.All(diffRepository, diffSecret),
		Description:   "Manages an agent secret for a GitHub repository.",
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
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSecretNameFunc,
				Description:      "Name of the secret.",
			},
			"key_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				RequiredWith:  []string{"value_encrypted"},
				ConflictsWith: []string{"value"},
				Description:   "ID of the agent public key used to encrypt the secret.",
			},
			"value": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"value", "value_encrypted"},
				Description:  "Plaintext value to be encrypted. This value is stored in Terraform state.",
			},
			"value_encrypted": {
				Type:             schema.TypeString,
				Optional:         true,
				Sensitive:        true,
				ExactlyOneOf:     []string{"value", "value_encrypted"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
				Description:      "Base64-encoded value encrypted with the agent public key identified by key_id.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the secret was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the secret was last updated by the provider.",
			},
			"remote_updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the secret was last updated on GitHub.",
			},
		},
	}
}

func resourceGithubAgentsSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	repository, _ := d.Get("repository").(string)
	repo, _, err := meta.v3client.Repositories.Get(ctx, meta.name, repository)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}
	return resourceGithubAgentsSecretUpdate(ctx, d, m)
}

func resourceGithubAgentsSecretUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	repository, _ := d.Get("repository").(string)
	name, _ := d.Get("secret_name").(string)
	request, err := agentSecretValue(ctx, d, meta, repository)
	if err != nil {
		return diag.FromErr(err)
	}
	if _, err := meta.v3client.Agents.CreateOrUpdateRepoSecret(ctx, meta.name, repository, name, request); err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(repository, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := d.Set("key_id", request.KeyID); err != nil {
		return diag.FromErr(err)
	}
	// PUT returns no metadata, so fetch the timestamp baseline for drift detection.
	secret, err := retryUntilResourceFound(ctx, func() (*github.Secret, error) {
		secret, _, err := meta.v3client.Agents.GetRepoSecret(ctx, meta.name, repository, name)
		return secret, err
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(setAgentSecretTimestamps(d, secret, true))
}

func resourceGithubAgentsSecretRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	repository, _ := d.Get("repository").(string)
	name, _ := d.Get("secret_name").(string)
	secret, _, err := meta.v3client.Agents.GetRepoSecret(ctx, meta.name, repository, name)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing agent secret from state because it no longer exists", map[string]any{"repository": repository, "secret_name": name})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	id, err := buildID(repository, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return diag.FromErr(setAgentSecretTimestamps(d, secret, false))
}

func resourceGithubAgentsSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	repository, _ := d.Get("repository").(string)
	name, _ := d.Get("secret_name").(string)
	_, err := meta.v3client.Agents.DeleteRepoSecret(ctx, meta.name, repository, name)
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
		return nil
	}
	return diag.FromErr(err)
}

func resourceGithubAgentsSecretImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	repository, name, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}
	meta, _ := m.(*Owner)
	repo, _, err := meta.v3client.Repositories.Get(ctx, meta.name, repository)
	if err != nil {
		return nil, err
	}
	if err := d.Set("repository", repository); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}
	if err := d.Set("secret_name", name); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
