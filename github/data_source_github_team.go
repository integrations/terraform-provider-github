package github

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubTeam() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTeamRead,

		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privacy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	slug := d.Get("slug").(string)
	log.Printf("[INFO] Refreshing GitHub Team: %s", slug)

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id
	ctx := context.Background()

	team, _, err := client.Teams.GetTeamBySlug(ctx, meta.(*Owner).name, slug)
	if err != nil {
		return err
	}

	member, _, err := client.Teams.ListTeamMembersByID(ctx, orgId, team.GetID(), nil)
	if err != nil {
		return err
	}

	members := []string{}
	for _, v := range member {
		members = append(members, v.GetLogin())
	}

	d.SetId(strconv.FormatInt(team.GetID(), 10))
	d.Set("name", team.GetName())
	d.Set("members", members)
	d.Set("description", team.GetDescription())
	d.Set("privacy", team.GetPrivacy())
	d.Set("permission", team.GetPermission())
	d.Set("node_id", team.GetNodeID())

	return nil
}
