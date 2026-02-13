package github

import (
	"context"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationWebhooks() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the webhooks for an organization.",
		Read:        dataSourceGithubOrganizationWebhooksRead,

		Schema: map[string]*schema.Schema{
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

func dataSourceGithubOrganizationWebhooksRead(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client
	ctx := context.Background()

	options := &github.ListOptions{
		PerPage: 100,
	}

	results := make([]map[string]any, 0)
	for {
		hooks, resp, err := client.Organizations.ListHooks(ctx, owner, options)
		if err != nil {
			return err
		}

		results = append(results, flattenGitHubWebhooks(hooks)...)
		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(owner)
	err := d.Set("webhooks", results)
	if err != nil {
		return err
	}

	return nil
}
