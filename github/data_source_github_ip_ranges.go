package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubIpRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubIpRangesRead,

		Schema: map[string]*schema.Schema{
			"hooks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"web": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"api": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"git": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"importer": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubIpRangesRead(d *schema.ResourceData, meta interface{}) error {
	org := meta.(*Organization)

	api, _, err := org.v3client.APIMeta(org.StopContext)
	if err != nil {
		return err
	}

	if len(api.Hooks)+len(api.Web)+len(api.Api)+len(api.Git)+len(api.Pages)+len(api.Importer) > 0 {
		d.SetId("github-ip-ranges")
	}
	if len(api.Hooks) > 0 {
		d.Set("hooks", api.Hooks)
	}
	if len(api.Web) > 0 {
		d.Set("web", api.Web)
	}
	if len(api.Api) > 0 {
		d.Set("api", api.Api)
	}
	if len(api.Git) > 0 {
		d.Set("git", api.Git)
	}
	if len(api.Pages) > 0 {
		d.Set("pages", api.Pages)
	}
	if len(api.Importer) > 0 {
		d.Set("importer", api.Importer)
	}

	return nil
}
