package github

import (
	"context"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseAppInstallations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseAppInstallationsRead,
		Description: "Use this data source to retrieve the GitHub App installations on an enterprise-owned organization. " +
			"This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.",

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
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the GitHub App installation.",
						},
						"app_slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL-friendly name of the GitHub App.",
						},
						"app_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the GitHub App.",
						},
						"repository_selection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the installation has access to all repositories or only selected ones. Possible values are 'all' or 'selected'.",
						},
						"permissions": {
							Type:        schema.TypeMap,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The permissions granted to the GitHub App installation.",
						},
						"events": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The list of events the GitHub App installation subscribes to.",
						},
						"client_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The OAuth client ID of the GitHub App.",
						},
						"target_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the account the GitHub App is installed on.",
						},
						"target_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of account the GitHub App is installed on. Possible values are 'Organization' or 'User'.",
						},
						"suspended": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the GitHub App installation is currently suspended.",
						},
						"single_file_paths": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The list of single file paths the GitHub App installation has access to.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date the GitHub App installation was created.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date the GitHub App installation was last updated.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseAppInstallationsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)

	opts := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	results := make([]map[string]any, 0)
	for {
		installations, resp, err := client.Enterprise.ListAppInstallations(ctx, enterpriseSlug, org, opts)
		if err != nil {
			return diag.FromErr(err)
		}

		results = append(results, flattenGitHubAppInstallations(installations)...)
		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	d.SetId(buildTwoPartID(enterpriseSlug, org))
	if err := d.Set("installations", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
