package github

import (
	"context"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseAppInstallableOrganizations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseAppInstallableOrganizationsRead,
		Description: "Use this data source to retrieve the enterprise-owned organizations that GitHub Apps can be installed on. " +
			"This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"organizations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of organizations in the enterprise that GitHub Apps can be installed on.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the organization.",
						},
						"login": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The login of the organization.",
						},
						"accessible_repositories_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API URL listing the repositories that can be made accessible to a GitHub App installed on the organization.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseAppInstallableOrganizationsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)

	opts := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	results := make([]map[string]any, 0)
	for {
		organizations, resp, err := client.Enterprise.ListAppInstallableOrganizations(ctx, enterpriseSlug, opts)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, organization := range organizations {
			results = append(results, map[string]any{
				"id":                          organization.ID,
				"login":                       organization.Login,
				"accessible_repositories_url": organization.GetAccessibleRepositoriesURL(),
			})
		}
		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	d.SetId(enterpriseSlug)
	if err := d.Set("organizations", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
