package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryWebhooks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryWebhooksRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"webhooks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryWebhooksRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	owner := meta.name
	client := meta.v3client

	repository, _ := d.Get("repository").(string)

	options := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	results := make([]map[string]any, 0)
	for {
		hooks, resp, err := client.Repositories.ListHooks(ctx, owner, repository, options)
		if err != nil {
			return diag.FromErr(err)
		}

		results = append(results, flattenGitHubWebhooks(hooks)...)

		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(fmt.Sprintf("%s/%s", owner, repository))
	if err := d.Set("repository", repository); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("webhooks", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenGitHubWebhooks(hooks []*github.Hook) []map[string]any {
	results := make([]map[string]any, 0)

	if hooks == nil {
		return results
	}

	for _, hook := range hooks {
		result := make(map[string]any)

		result["id"] = hook.ID
		result["type"] = hook.Type
		result["name"] = hook.Name
		result["url"] = hook.URL
		result["active"] = hook.Active

		results = append(results, result)
	}

	return results
}
