package github

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseAppAccessibleOrganizationRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseAppAccessibleOrganizationRepositoriesRead,
		Description: "Use this data source to retrieve repositories of an enterprise-owned organization that a GitHub App can be granted access to.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The login of the enterprise-owned organization.",
			},
			"repositories": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Repositories of the organization a GitHub App can access.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"full_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseAppAccessibleOrganizationRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)

	opts := &github.ListOptions{PerPage: maxPerPage}
	results := make([]map[string]any, 0)
	for {
		repos, resp, err := client.Enterprise.ListAppAccessibleOrganizationRepositories(ctx, enterprise, org, opts)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, r := range repos {
			results = append(results, map[string]any{
				"id":        r.ID,
				"name":      r.Name,
				"full_name": r.FullName,
			})
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	d.SetId(buildTwoPartID(enterprise, org))
	if err := d.Set("repositories", results); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
