package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	tflog.Debug(ctx, "GitHub Organization Team Sync Groups State before migration", map[string]any{"state": rawState})

	if rawState == nil {
		return nil, fmt.Errorf("cannot migrate nil state")
	}

	orgName := meta.(*Owner).name

	newID, err := buildID(orgName, "team-sync-groups", "")
	if err != nil {
		return nil, fmt.Errorf("error building migrated ID: %w", err)
	}

	rawState["id"] = newID

	tflog.Debug(ctx, "GitHub Organization Team Sync Groups State after migration", map[string]any{"state": rawState})

	return rawState, nil
}
