package github

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseAppInstallableOrganizations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseAppInstallableOrganizationsRead,
		Description: "Use this data source to retrieve the organizations in an enterprise that a GitHub App can be installed on.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"organizations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Organizations in the enterprise that can have a GitHub App installed on them.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"login": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"accessible_repositories_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseAppInstallableOrganizationsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)

	opts := &github.ListOptions{PerPage: maxPerPage}
	results := make([]map[string]any, 0)
	for {
		orgs, resp, err := client.Enterprise.ListAppInstallableOrganizations(ctx, enterprise, opts)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, o := range orgs {
			entry := map[string]any{
				"id":    o.ID,
				"login": o.Login,
			}
			if o.AccessibleRepositoriesURL != nil {
				entry["accessible_repositories_url"] = *o.AccessibleRepositoriesURL
			}
			results = append(results, entry)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	d.SetId(enterprise)
	if err := d.Set("organizations", results); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
