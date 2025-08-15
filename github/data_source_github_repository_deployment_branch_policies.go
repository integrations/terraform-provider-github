package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryDeploymentBranchPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryDeploymentBranchPoliciesRead,

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

func dataSourceGithubRepositoryDeploymentBranchPoliciesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment_name").(string)

	policies, _, err := client.Repositories.ListDeploymentBranchPolicies(context.Background(), owner, repoName, environmentName)
	if err != nil {
		return nil
	}

	results := make([]map[string]any, 0)

	for _, policy := range policies.BranchPolicies {
		policyMap := make(map[string]any)
		policyMap["id"] = strconv.FormatInt(*policy.ID, 10)
		policyMap["name"] = policy.Name
		results = append(results, policyMap)
	}

	d.SetId(repoName + ":" + environmentName)
	err = d.Set("deployment_branch_policies", results)
	if err != nil {
		return err
	}

	return nil
}
