package github

import (
	"context"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationVariableRepositories() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"variable_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSecretNameFunc,
				Description:      "Name of the existing variable.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Set:  schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Required:    true,
				Description: "An array of repository ids that can access the organization variable.",
			},
		},

		CreateContext: resourceGithubActionsOrganizationVariableRepositoriesCreateOrUpdate,
		ReadContext:   resourceGithubActionsOrganizationVariableRepositoriesRead,
		UpdateContext: resourceGithubActionsOrganizationVariableRepositoriesCreateOrUpdate,
		DeleteContext: resourceGithubActionsOrganizationVariableRepositoriesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationVariableRepositoriesImport,
		},
	}
}

func resourceGithubActionsOrganizationVariableRepositoriesCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	variableName := d.Get("variable_name").(string)
	repoIDs := []int64{}

	ids := d.Get("selected_repository_ids").(*schema.Set).List()
	for _, id := range ids {
		repoIDs = append(repoIDs, int64(id.(int)))
	}

	_, err := client.Actions.SetSelectedReposForOrgVariable(ctx, owner, variableName, repoIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(variableName)

	return nil
}

func resourceGithubActionsOrganizationVariableRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	variableName := d.Get("variable_name").(string)

	repoIDs := []int64{}
	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}
	for {
		results, resp, err := client.Actions.ListSelectedReposForOrgVariable(ctx, owner, variableName, opt)
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

	return nil
}

func resourceGithubActionsOrganizationVariableRepositoriesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	_, err := client.Actions.SetSelectedReposForOrgVariable(ctx, owner, d.Id(), []int64{})
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationVariableRepositoriesImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	variableName := d.Id()

	if err := d.Set("variable_name", variableName); err != nil {
		return nil, err
	}

	repoIDs := []int64{}
	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}
	for {
		results, resp, err := client.Actions.ListSelectedReposForOrgVariable(ctx, owner, variableName, opt)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
