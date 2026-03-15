package github

import (
	"context"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryEnvironments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryEnvironmentsRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryEnvironmentsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	results := make([]map[string]any, 0)

	for {
		listOptions := &github.EnvironmentListOptions{}
		environments, resp, err := client.Repositories.ListEnvironments(ctx, orgName, repoName, listOptions)
		if err != nil {
			return diag.FromErr(err)
		}

		results = append(results, flattenEnvironments(environments)...)

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	d.SetId(repoName)
	if err := d.Set("environments", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenEnvironments(environments *github.EnvResponse) []map[string]any {
	results := make([]map[string]any, 0)
	if environments == nil {
		return results
	}

	for _, environment := range environments.Environments {
		environmentMap := make(map[string]any)
		environmentMap["name"] = environment.GetName()
		environmentMap["node_id"] = environment.GetNodeID()
		results = append(results, environmentMap)
	}

	return results
}
