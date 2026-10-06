package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubAgentsOrganizationSecretRepositories() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubAgentsOrganizationSecretRepositoriesCreateOrUpdate,
		ReadContext:   resourceGithubAgentsOrganizationSecretRepositoriesRead,
		UpdateContext: resourceGithubAgentsOrganizationSecretRepositoriesCreateOrUpdate,
		DeleteContext: resourceGithubAgentsOrganizationSecretRepositoriesDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubAgentsOrganizationSecretImport},
		Description:   "Manages the complete set of repositories with access to an organization agent secret. Requires selected visibility.",
		Schema: map[string]*schema.Schema{
			"secret_name":             {Type: schema.TypeString, Required: true, ForceNew: true, ValidateDiagFunc: validateSecretNameFunc, Description: "Name of the existing organization agent secret."},
			"selected_repository_ids": {Type: schema.TypeSet, Required: true, Set: schema.HashInt, Elem: &schema.Schema{Type: schema.TypeInt, ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1))}, Description: "Complete set of repository IDs allowed to access the secret. An empty set removes all access."},
		},
	}
}

func resourceGithubAgentsOrganizationSecretRepositoriesCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	ids := []int64{}
	selectedIDs, _ := d.Get("selected_repository_ids").(*schema.Set)
	for _, id := range selectedIDs.List() {
		repoID, _ := id.(int)
		ids = append(ids, int64(repoID))
	}
	name, _ := d.Get("secret_name").(string)
	if _, err := meta.v3client.Agents.SetSelectedReposForOrgSecret(ctx, meta.name, name, ids); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(name)
	return nil
}

func resourceGithubAgentsOrganizationSecretRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	ids := []int64{}
	for repo, err := range meta.v3client.Agents.ListSelectedReposForOrgSecretIter(ctx, meta.name, name, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
		ids = append(ids, repo.GetID())
	}
	return diag.FromErr(d.Set("selected_repository_ids", ids))
}

func resourceGithubAgentsOrganizationSecretRepositoriesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	_, err := meta.v3client.Agents.SetSelectedReposForOrgSecret(ctx, meta.name, name, []int64{})
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
		return nil
	}
	return diag.FromErr(err)
}
