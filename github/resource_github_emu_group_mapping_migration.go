package github

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEMUGroupMappingV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Slug of the GitHub team.",
			},
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Integer corresponding to the external group ID to be linked.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubEMUGroupMappingStateUpgradeV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	tflog.Trace(ctx, "GitHub EMU Group Mapping State before migration", map[string]any{"state": rawState, "owner": orgName})

	teamSlug := rawState["team_slug"].(string)
	// We need to bypass the etag because we need to get the latest group
	ctx = context.WithValue(ctx, ctxEtag, nil)
	groupsList, resp, err := client.Teams.ListExternalGroupsForTeamBySlug(ctx, orgName, teamSlug)
	if err != nil {
		if resp != nil && (resp.StatusCode == http.StatusNotFound) {
			// If the Group is not found, remove it from state
			tflog.Info(ctx, "Removing EMU group mapping from state because team no longer exists in GitHub", map[string]any{
				"resource_id": rawState["id"],
			})
			return nil, err
		}
		return nil, err
	}

	group := groupsList.Groups[0]
	teamID, err := lookupTeamID(ctx, client, orgName, teamSlug)
	if err != nil {
		return nil, err
	}
	rawState["team_id"] = int(teamID)
	resourceID, err := buildID(strconv.FormatInt(teamID, 10), teamSlug, strconv.FormatInt(group.GetGroupID(), 10))
	if err != nil {
		return nil, err
	}
	rawState["id"] = resourceID

	tflog.Trace(ctx, "GitHub EMU Group Mapping State after migration", map[string]any{"state": rawState})
	return rawState, nil
}
