package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryEnvironmentDeploymentPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryEnvironmentDeploymentPoliciesRead,

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

func dataSourceGithubRepositoryEnvironmentDeploymentPoliciesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment").(string)

	policies, _, err := client.Repositories.ListDeploymentBranchPolicies(context.Background(), owner, repoName, environmentName)
	if err != nil {
		return err
	}

	results := make([]map[string]any, 0)

	for _, policy := range policies.BranchPolicies {
		policyMap := make(map[string]any)
		policyMap["type"] = policy.GetType()
		policyMap["pattern"] = policy.GetName()
		results = append(results, policyMap)
	}

	d.SetId(fmt.Sprintf("%s:%s", repoName, environmentName))
	return d.Set("policies", results)
}
