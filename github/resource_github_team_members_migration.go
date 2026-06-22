package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubTeamMembersV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub team id or slug",
			},
			"members": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of team members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
							Description:      "The user to add to the team.",
						},
						"role": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "member",
							Description:      "The role of the user within the team. Must be one of 'member' or 'maintainer'.",
							ValidateDiagFunc: validateValueFunc([]string{"member", "maintainer"}),
						},
					},
				},
			},
		},
	}
}

func resourceGithubTeamMembersStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	tflog.Debug(ctx, "Migrating GitHub Team Members from v0 to v1.", map[string]any{"raw_state": rawState})

	idv, ok := rawState["id"]
	if !ok {
		return nil, fmt.Errorf("missing id in raw state")
	}

	id, ok := idv.(string)
	if !ok {
		return nil, fmt.Errorf("id in raw state is not a string")
	}

	var teamID int64
	var teamSlug string
	team := newLegacyTeamIdentity(id)
	if s, ok := team.getSlugOK(); ok {
		teamSlug = s

		id, err := lookupTeamID(ctx, client, owner, teamSlug)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup team ID for slug %s: %w", teamSlug, err)
		}

		teamID = id
	} else {
		teamID = team.getID()

		s, err := lookupTeamSlug(ctx, client, meta.id, teamID)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup team slug for ID %d: %w", teamID, err)
		}

		teamSlug = s
	}

	rawState["id"] = strconv.FormatInt(teamID, 10)
	rawState["team_slug"] = teamSlug

	tflog.Debug(ctx, "GitHub Team Members migrated to v1.", map[string]any{"raw_state": rawState})

	return rawState, nil
}
