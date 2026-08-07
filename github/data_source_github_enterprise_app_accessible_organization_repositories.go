package github

import (
	"context"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseAppAccessibleOrganizationRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseAppAccessibleOrganizationRepositoriesRead,
		Description: "Use this data source to retrieve the repositories of an enterprise-owned organization that GitHub Apps can be granted access to. " +
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
			"repositories": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of repositories of the organization that GitHub Apps can be granted access to.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the repository.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the repository.",
						},
						"full_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The full name of the repository, in the format `<organization>/<name>`.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseAppAccessibleOrganizationRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)

	opts := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	results := make([]map[string]any, 0)
	for {
		repositories, resp, err := client.Enterprise.ListAppAccessibleOrganizationRepositories(ctx, enterpriseSlug, org, opts)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, repository := range repositories {
			results = append(results, map[string]any{
				"id":        repository.ID,
				"name":      repository.Name,
				"full_name": repository.FullName,
			})
		}
		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	d.SetId(buildTwoPartID(enterpriseSlug, org))
	if err := d.Set("repositories", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
