package github

import (
	"context"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/shurcooL/githubv4"
)

const (
	GITHUB_IP_RANGE_GIT      = "git"
	GITHUB_IP_RANGE_HOOKS    = "hooks"
	GITHUB_IP_RANGE_IMPORTER = "importer"
	GITHUB_IP_RANGE_PAGES    = "pages"
)

func dataSourceGithubIpRanges() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			GITHUB_IP_RANGE_GIT: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			GITHUB_IP_RANGE_HOOKS: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			GITHUB_IP_RANGE_IMPORTER: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			GITHUB_IP_RANGE_PAGES: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},

		Read: resourceGithubAppInitIpRangesRead,
	}
}

func resourceGithubAppInitIpRangesRead(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Meta struct {
			GitIpAddresses      []githubv4.String
			HookIpAddresses     []githubv4.String
			ImporterIpAddresses []githubv4.String
			PagesIpAddresses    []githubv4.String
		}
	}
	variables := map[string]interface{}{}

	ctx := context.Background()
	client := meta.(*Organization).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return err
	}

	err = d.Set(GITHUB_IP_RANGE_GIT, query.Meta.GitIpAddresses)
	if err != nil {
		return err
	}

	err = d.Set(GITHUB_IP_RANGE_HOOKS, query.Meta.HookIpAddresses)
	if err != nil {
		return err
	}

	err = d.Set(GITHUB_IP_RANGE_IMPORTER, query.Meta.ImporterIpAddresses)
	if err != nil {
		return err
	}

	err = d.Set(GITHUB_IP_RANGE_PAGES, query.Meta.PagesIpAddresses)
	if err != nil {
		return err
	}

	d.SetId("github/ip_ranges")

	return nil
}
