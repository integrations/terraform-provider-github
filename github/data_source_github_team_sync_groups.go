package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceGithubTeamSyncGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTeamSyncGroupsRead,

		Schema: map[string]*schema.Schema{
			"retrieve_by": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"id",
					"slug",
				}, false),
			},
			"org_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"team_slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubTeamSyncGroupsRead(d *schema.ResourceData, meta interface{}) error {
	retrieveBy := d.Get("retrieve_by").(string)
	log.Print("[INFO] Refreshing GitHub Team Sync Groups")

	client := meta.(*Organization).client
	ctx := context.Background()

	var team *github.Team
	var err error

	if retrieveBy == "id" {
		orgID := meta.(*Organization).id
		teamID, err := strconv.ParseInt(d.Get("team_id").(string), 10, 64)
		if err != nil {
			return err
		}
		team, _, err = client.Teams.GetTeamByID(ctx, orgID, teamID)
	} else {
		orgName := meta.(*Organization).name
		slug := d.Get("team_slug").(string)
		team, err = getGithubTeamBySlug(ctx, client, orgName, slug)
	}

	if err != nil {
		return err
	}

	idpGroups, _, err := client.Teams.ListIDPGroupsForTeam(ctx, string(team.GetID()))
	if err != nil {
		return fmt.Errorf("Could not find team with slug: %s", d.Get("team_slug").(string))
	}

	groups := []map[string]string{}
	for _, g := range idpGroups.Groups {
		group := map[string]string{
			"group_id":          g.GetGroupID(),
			"group_name":        g.GetGroupName(),
			"group_description": g.GetGroupDescription(),
		}
		groups = append(groups, group)
	}

	d.SetId("github-team-sync-groups")
	d.Set("org_name", meta.(*Organization).name)
	d.Set("org_id", meta.(*Organization).id)
	if d.Get("team_id") == nil {
		d.Set("team_id", string(team.GetID()))
	}
	d.Set("groups", groups)

	return nil
}
