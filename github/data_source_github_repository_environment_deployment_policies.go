package github

import (
	"context"
	"net/url"

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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An array of deployment branch policies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of deployment policy (branch or tag).",
						},
						"pattern": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name pattern that branches or tags must match in order to deploy to the environment.",
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

	policies, _, err := client.Repositories.ListDeploymentBranchPolicies(ctx, owner, repoName, url.PathEscape(envName))
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]map[string]any, 0)

	for _, policy := range policies.BranchPolicies {
		policyMap := make(map[string]any)
		policyMap["type"] = policy.GetType()
		policyMap["pattern"] = policy.GetName()
		results = append(results, policyMap)
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
