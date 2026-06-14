package github

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseAppInstallations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseAppInstallationsRead,
		Description: "Use this data source to retrieve the GitHub App installations on an enterprise-owned organization.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise that owns the organization.",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The login of the enterprise-owned organization.",
			},
			"installations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of GitHub App installations on the organization.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"app_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"app_slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"target_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository_selection": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"permissions": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"events": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"suspended": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"single_file_paths": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseAppInstallationsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)

	opts := &github.ListOptions{PerPage: maxPerPage}
	results := make([]map[string]any, 0)
	for {
		installations, resp, err := client.Enterprise.ListAppInstallations(ctx, enterprise, org, opts)
		if err != nil {
			return diag.FromErr(err)
		}
		results = append(results, flattenGitHubAppInstallations(installations)...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	d.SetId(buildTwoPartID(enterprise, org))
	if err := d.Set("installations", results); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
