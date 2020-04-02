package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func dataSourceGithubOrganizationTeamSyncGroupsRead(d *schema.ResourceData, meta interface{}) error {
	log.Print("[INFO] Refreshing GitHub Organization Team-Sync Groups")

	client := meta.(*Organization).client
	ctx := context.Background()

	orgName := meta.(*Organization).name
	options := &github.ListCursorOptions{PerPage: maxPerPage}

	groups := make([]interface{}, 0)
	for {
		idpGroupList, resp, err := client.Teams.ListIDPGroupsInOrganization(ctx, orgName, options)
		if err != nil {
			return err
		}

		result, err := flattenGithubIDPGroupList(idpGroupList)
		if err != nil {
			return fmt.Errorf("unable to flatten IdP Groups in Github Organization(Org: %q) : %+v", orgName, err)
		}

		groups = append(groups, result...)

		if resp.NextPage == 0 {
			break
		}
		options.Page = string(resp.NextPage)
	}

	d.SetId("github-org-team-sync-groups")
	d.Set("groups", groups)

	return nil
}

func flattenGithubIDPGroupList(idpGroupList *github.IDPGroupList) ([]interface{}, error) {
	if idpGroupList == nil {
		return make([]interface{}, 0), nil
	}
	results := make([]interface{}, 0)
	for _, group := range idpGroupList.Groups {
		result := make(map[string]interface{})
		result["group_id"] = group.GetGroupID()
		result["group_name"] = group.GetGroupName()
		result["group_description"] = group.GetGroupDescription()
		results = append(results, result)
	}

	return results, nil
}
