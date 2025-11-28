package github

import (
	"context"

	"github.com/google/go-github/v66/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationAppInstallations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationAppInstallationsRead,

		Schema: map[string]*schema.Schema{
			"installations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"app_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationAppInstallationsRead(d *schema.ResourceData, meta interface{}) error {
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client
	ctx := context.Background()

	options := &github.ListOptions{
		PerPage: 100,
	}

	results := make([]map[string]interface{}, 0)
	for {
		appInstallations, resp, err := client.Organizations.ListInstallations(ctx, owner, options)
		if err != nil {
			return err
		}

		results = append(results, flattenGitHubAppInstallations(appInstallations.Installations)...)
		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(owner)
	err := d.Set("installations", results)
	if err != nil {
		return err
	}

	return nil
}

func flattenGitHubAppInstallations(orgAppInstallations []*github.Installation) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	if orgAppInstallations == nil {
		return results
	}

	for _, appInstallation := range orgAppInstallations {
		result := make(map[string]interface{})

		result["id"] = appInstallation.ID
		result["slug"] = appInstallation.AppSlug
		result["app_id"] = appInstallation.AppID

		results = append(results, result)
	}

	return results
}
