package github

import (
	"context"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationSecretRepositories() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSecretNameFunc,
				Description:      "Name of the existing secret.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Set:  schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Required:    true,
				Description: "An array of repository ids that can access the organization secret.",
			},
		},

		CreateContext: resourceGithubActionsOrganizationSecretRepositoriesCreateOrUpdate,
		ReadContext:   resourceGithubActionsOrganizationSecretRepositoriesRead,
		UpdateContext: resourceGithubActionsOrganizationSecretRepositoriesCreateOrUpdate,
		DeleteContext: resourceGithubActionsOrganizationSecretRepositoriesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationSecretRepositoriesImport,
		},
	}
}

func resourceGithubActionsOrganizationSecretRepositoriesCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)
	repoIDs := []int64{}

	ids := d.Get("selected_repository_ids").(*schema.Set).List()
	for _, id := range ids {
		repoIDs = append(repoIDs, int64(id.(int)))
	}

	_, err := client.Actions.SetSelectedReposForOrgSecret(ctx, owner, secretName, repoIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(secretName)

	return nil
}

func resourceGithubActionsOrganizationSecretRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Get("secret_name").(string)

	repoIDs := []int64{}
	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}
	for {
		results, resp, err := client.Actions.ListSelectedReposForOrgSecret(ctx, owner, secretName, opt)
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

func resourceGithubActionsOrganizationSecretRepositoriesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	_, err := client.Actions.SetSelectedReposForOrgSecret(ctx, owner, d.Id(), []int64{})
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationSecretRepositoriesImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName := d.Id()

	if err := d.Set("secret_name", secretName); err != nil {
		return nil, err
	}

	repoIDs := []int64{}
	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}
	for {
		results, resp, err := client.Actions.ListSelectedReposForOrgSecret(ctx, owner, secretName, opt)
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
