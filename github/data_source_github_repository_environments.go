package github

import (
	"context"

	"github.com/google/go-github/v74/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryEnvironments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryEnvironmentsRead,

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

func dataSourceGithubRepositoryEnvironmentsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	results := make([]map[string]any, 0)

	var listOptions *github.EnvironmentListOptions
	for {
		environments, resp, err := client.Repositories.ListEnvironments(context.Background(), orgName, repoName, listOptions)
		if err != nil {
			return err
		}

		results = append(results, flattenEnvironments(environments)...)

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	d.SetId(repoName)
	err := d.Set("environments", results)
	if err != nil {
		return err
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
