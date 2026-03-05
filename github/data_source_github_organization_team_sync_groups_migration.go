package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationTeamSyncGroupsV0() *schema.Resource {
	return &schema.Resource{
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

func dataSourceGithubOrganizationTeamSyncGroupsStateUpgradeV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Organization Team Sync Groups State before migration: %#v", rawState)

	if rawState == nil {
		return nil, fmt.Errorf("cannot migrate nil state")
	}

	orgName := meta.(*Owner).name

	newID, err := buildID(orgName, "team-sync-groups")
	if err != nil {
		return nil, fmt.Errorf("error building migrated ID: %w", err)
	}

	rawState["id"] = newID

	log.Printf("[DEBUG] GitHub Organization Team Sync Groups State after migration: %#v", rawState)

	return rawState, nil
}
