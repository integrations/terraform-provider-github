package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubAgentsOrganizationSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubAgentsOrganizationSecretCreate,
		ReadContext:   resourceGithubAgentsOrganizationSecretRead,
		UpdateContext: resourceGithubAgentsOrganizationSecretUpdate,
		DeleteContext: resourceGithubAgentsOrganizationSecretDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubAgentsOrganizationSecretImport},
		CustomizeDiff: diffSecret,
		Description:   "Manages an agent secret for a GitHub organization. Repository access is managed separately.",
		Schema: map[string]*schema.Schema{
			"secret_name":             {Type: schema.TypeString, Required: true, ForceNew: true, ValidateDiagFunc: validateSecretNameFunc, Description: "Name of the secret."},
			"key_id":                  {Type: schema.TypeString, Optional: true, Computed: true, RequiredWith: []string{"value_encrypted"}, ConflictsWith: []string{"value"}, Description: "ID of the agent public key used to encrypt the secret."},
			"value":                   {Type: schema.TypeString, Optional: true, Sensitive: true, ExactlyOneOf: []string{"value", "value_encrypted"}, Description: "Plaintext value to be encrypted. This value is stored in Terraform state."},
			"value_encrypted":         {Type: schema.TypeString, Optional: true, Sensitive: true, ExactlyOneOf: []string{"value", "value_encrypted"}, ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64), Description: "Base64-encoded value encrypted with the agent public key identified by key_id."},
			"visibility":              {Type: schema.TypeString, Required: true, ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "private", "selected"}, false)), Description: "Repository access to the secret: all, private, or selected."},
			"selected_repository_ids": {Type: schema.TypeSet, Computed: true, Set: schema.HashInt, Elem: &schema.Schema{Type: schema.TypeInt}, Description: "Observed repository access, used to distinguish access changes from value drift. Manage access with the separate repository association resources."},
			"created_at":              {Type: schema.TypeString, Computed: true, Description: "Timestamp of when the secret was created."},
			"updated_at":              {Type: schema.TypeString, Computed: true, Description: "Timestamp baseline for detecting value drift. Updated after provider writes or observed repository access changes."},
			"remote_updated_at":       {Type: schema.TypeString, Computed: true, Description: "Timestamp of when the secret was last updated on GitHub."},
		},
	}
}

func resourceGithubAgentsOrganizationSecretCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return resourceGithubAgentsOrganizationSecretUpdate(ctx, d, m)
}

func resourceGithubAgentsOrganizationSecretUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	value, err := agentSecretValue(ctx, d, meta, "")
	if err != nil {
		return diag.FromErr(err)
	}
	visibility, _ := d.Get("visibility").(string)
	request := github.SecretOrgRequest{KeyID: value.KeyID, EncryptedValue: value.EncryptedValue, Visibility: visibility}
	if d.Id() != "" && visibility == "selected" {
		secret, _, err := meta.v3client.Agents.GetOrgSecret(ctx, meta.name, name)
		if err != nil {
			return diag.FromErr(err)
		}
		if secret.Visibility == "selected" {
			// GitHub clears access when this field is omitted. Preserve the live
			// list, including associations managed outside this Terraform state.
			ids, err := agentOrganizationSecretRepositoryIDs(ctx, meta, name)
			if err != nil {
				return diag.FromErr(err)
			}
			request.SelectedRepositoryIDs = ids
		}
	}
	if _, err := meta.v3client.Agents.CreateOrUpdateOrgSecret(ctx, meta.name, name, request); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(name)
	if err := d.Set("key_id", value.KeyID); err != nil {
		return diag.FromErr(err)
	}
	secret, err := retryUntilResourceFound(ctx, func() (*github.Secret, error) {
		secret, _, err := meta.v3client.Agents.GetOrgSecret(ctx, meta.name, name)
		return secret, err
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(setAgentOrganizationSecretMetadata(ctx, d, meta, secret, true))
}

func resourceGithubAgentsOrganizationSecretRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	secret, _, err := meta.v3client.Agents.GetOrgSecret(ctx, meta.name, name)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing organization agent secret from state because it no longer exists", map[string]any{"secret_name": d.Id()})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	return diag.FromErr(setAgentOrganizationSecretMetadata(ctx, d, meta, secret, false))
}

func setAgentOrganizationSecretMetadata(ctx context.Context, d *schema.ResourceData, meta *Owner, secret *github.Secret, written bool) error {
	ids := []any{}
	if secret.Visibility == "selected" {
		name, _ := d.Get("secret_name").(string)
		repositoryIDs, err := agentOrganizationSecretRepositoryIDs(ctx, meta, name)
		if err != nil {
			return err
		}
		for _, id := range repositoryIDs {
			ids = append(ids, int(id))
		}
	}
	previousIDs, _ := d.Get("selected_repository_ids").(*schema.Set)
	currentIDs := schema.NewSet(schema.HashInt, ids)
	previousVisibility, _ := d.Get("visibility").(string)
	accessChanged := !previousIDs.Equal(currentIDs) || previousVisibility != secret.Visibility
	if err := d.Set("selected_repository_ids", currentIDs); err != nil {
		return err
	}
	if err := d.Set("visibility", secret.Visibility); err != nil {
		return err
	}
	// GitHub also advances updated_at on access edits. With no value-specific
	// version, simultaneous access and value changes cannot be distinguished.
	return setAgentSecretTimestamps(d, secret, written || accessChanged)
}

func agentOrganizationSecretRepositoryIDs(ctx context.Context, meta *Owner, name string) ([]int64, error) {
	ids := []int64{}
	for repo, err := range meta.v3client.Agents.ListSelectedReposForOrgSecretIter(ctx, meta.name, name, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			return nil, err
		}
		ids = append(ids, repo.GetID())
	}
	return ids, nil
}

func resourceGithubAgentsOrganizationSecretDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	_, err := meta.v3client.Agents.DeleteOrgSecret(ctx, meta.name, name)
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
		return nil
	}
	return diag.FromErr(err)
}

func resourceGithubAgentsOrganizationSecretImport(_ context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	if err := checkOrganization(m); err != nil {
		return nil, err
	}
	if err := d.Set("secret_name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
