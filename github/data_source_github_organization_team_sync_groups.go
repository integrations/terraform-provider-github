package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationTeamSyncGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationTeamSyncGroupsRead,

		Schema: map[string]*schema.Schema{
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_description": {
							Type:     schema.TypeString,
							Computed: true,
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
