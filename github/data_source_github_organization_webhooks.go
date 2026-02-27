package github

import (
	"context"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationWebhooks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationWebhooksRead,

		Schema: map[string]*schema.Schema{
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

func dataSourceGithubOrganizationWebhooksRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client

	options := &github.ListOptions{
		PerPage: 100,
	}

	results := make([]map[string]any, 0)
	for {
		hooks, resp, err := client.Organizations.ListHooks(ctx, owner, options)
		if err != nil {
			return diag.FromErr(err)
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
		return diag.FromErr(err)
	}

	return nil
}
