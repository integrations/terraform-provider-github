package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryWebhooks() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the webhooks for a repository.",
		Read:        dataSourceGithubRepositoryWebhooksRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
			"webhooks": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An array of GitHub webhooks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the webhook.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the webhook.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the webhook.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the webhook.",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the webhook is active.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryWebhooksRead(d *schema.ResourceData, meta any) error {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client
	ctx := context.Background()

	options := &github.ListOptions{
		PerPage: 100,
	}

	results := make([]map[string]any, 0)
	for {
		hooks, resp, err := client.Repositories.ListHooks(ctx, owner, repository, options)
		if err != nil {
			return err
		}

		results = append(results, flattenGitHubWebhooks(hooks)...)

		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(fmt.Sprintf("%s/%s", owner, repository))
	if err := d.Set("repository", repository); err != nil {
		return err
	}
	if err := d.Set("webhooks", results); err != nil {
		return err
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
