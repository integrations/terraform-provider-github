package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryDeploymentBranchPolicies() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This data source is deprecated in favour of the github_repository_environment_deployment_policies data source.",

		ReadContext: dataSourceGithubRepositoryDeploymentBranchPoliciesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository name.",
			},
			"environment_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target environment name.",
			},
			"deployment_branch_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryDeploymentBranchPoliciesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment_name").(string)

	results := make([]map[string]any, 0)
	listOptions := &github.ListOptions{PerPage: maxPerPage}
	for {
		policies, resp, err := client.Repositories.ListDeploymentBranchPolicies(ctx, owner, repoName, environmentName, listOptions)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, policy := range policies.BranchPolicies {
			policyMap := make(map[string]any)
			policyMap["id"] = strconv.FormatInt(*policy.ID, 10)
			policyMap["name"] = policy.Name
			results = append(results, policyMap)
		}

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	id, err := buildID(repoName, environmentName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("deployment_branch_policies", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
