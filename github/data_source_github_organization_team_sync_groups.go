package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationTeamSyncGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the identity provider groups for team synchronization in an organization.",
		Read:        dataSourceGithubOrganizationTeamSyncGroupsRead,

		Schema: map[string]*schema.Schema{
			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An array of GitHub Identity Provider Groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the IdP group.",
						},
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the IdP group.",
						},
						"group_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the IdP group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationTeamSyncGroupsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	orgName := meta.(*Owner).name
	options := &github.ListIDPGroupsOptions{
		ListCursorOptions: github.ListCursorOptions{
			PerPage: maxPerPage,
		},
	}

	groups := make([]any, 0)
	for {
		idpGroupList, resp, err := client.Teams.ListIDPGroupsInOrganization(ctx, orgName, options)
		if err != nil {
			return err
		}

		result := flattenGithubIDPGroupList(idpGroupList)

		groups = append(groups, result...)

		if resp.NextPageToken == "" {
			break
		}
		options.Page = resp.NextPageToken
	}

	d.SetId(fmt.Sprintf("%s/github-org-team-sync-groups", orgName))
	if err := d.Set("groups", groups); err != nil {
		return fmt.Errorf("error setting groups: %w", err)
	}

	return nil
}
