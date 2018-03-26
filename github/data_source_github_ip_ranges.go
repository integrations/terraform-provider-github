package github

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGithubIpRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubIpRangesRead,

		Schema: map[string]*schema.Schema{
			"hooks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"git": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pages": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubIpRangesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	ctx := context.Background()

	api, _, err := client.APIMeta(ctx)
	if err != nil {
		return err
	}

	d.SetId("github-ip-ranges")
	d.Set("hooks", api.Hooks)
	d.Set("git", api.Git)
	d.Set("pages", api.Pages)

	return nil
}
