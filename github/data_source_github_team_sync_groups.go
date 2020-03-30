package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubTeamSyncGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTeamSyncGroupsRead,

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"org_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"team_id": {
				Type:     schema.TypeInt,
				Computed: true,
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
						"description": {
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
	slug := d.Get("team_slug").(string)
	log.Printf("[INFO] Refreshing GitHub Team Sync Groups: %s", slug)

	client := meta.(*Organization).client
	ctx := context.Background()

	team, err := getGithubTeamBySlug(ctx, client, meta.(*Organization).name, slug)
	if err != nil {
		return err
	}

	idpGroups, _, err := client.Teams.ListIDPGroupsForTeam(ctx, string(team.GetID()))
	if err != nil {
		return err
	}

	groups := []map[string]string{}
	for _, g := range idpGroups.Groups {
		group := map[string]string{
			"group_id":    g.GetGroupID(),
			"group_name":  g.GetGroupName(),
			"description": g.GetGroupDescription(),
		}
		groups = append(groups, group)
	}

	d.SetId("github-team-sync-groups")
	d.Set("org_id", meta.(*Organization).id)
	d.Set("org_name", meta.(*Organization).name)
	d.Set("team_id", team.GetID())
	d.Set("groups", groups)

	return nil
}
