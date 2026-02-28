package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryEnvironmentDeploymentPolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryEnvironmentDeploymentPoliciesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the GitHub repository.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the environment.",
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pattern": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryEnvironmentDeploymentPoliciesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)

	results := make([]map[string]any, 0)
	listOptions := &github.ListOptions{PerPage: maxPerPage}
	for {
		policies, resp, err := client.Repositories.ListDeploymentBranchPolicies(ctx, owner, repoName, url.PathEscape(envName), listOptions)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, policy := range policies.BranchPolicies {
			policyMap := make(map[string]any)
			policyMap["type"] = policy.GetType()
			policyMap["pattern"] = policy.GetName()
			results = append(results, policyMap)
		}

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	id, err := buildID(repoName, escapeIDPart(envName))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err = d.Set("policies", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
