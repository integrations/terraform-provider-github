package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryDeployKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryDeployKeysRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"verified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryDeployKeysRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client

	options := &github.ListOptions{
		PerPage: 100,
	}

	results := make([]map[string]any, 0)
	for {
		keys, resp, err := client.Repositories.ListKeys(ctx, owner, repository, options)
		if err != nil {
			return diag.FromErr(err)
		}

		results = append(results, flattenGitHubDeployKeys(keys)...)

		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(fmt.Sprintf("%s/%s", owner, repository))
	err := d.Set("keys", results)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenGitHubDeployKeys(keys []*github.Key) []map[string]any {
	results := make([]map[string]any, 0)

	if keys == nil {
		return results
	}

	for _, c := range keys {
		result := make(map[string]any)

		result["id"] = c.ID
		result["key"] = c.Key
		result["title"] = c.Title
		result["verified"] = c.Verified

		results = append(results, result)
	}

	return results
}
