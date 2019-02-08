package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}
}

func dataSourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	slug := d.Get("slug").(string)
	log.Printf("[INFO] Refreshing GitHub Team: %s", slug)

	client := meta.(*Organization).client
	ctx := context.Background()

	team, err := getGithubTeamBySlug(client, meta.(*Organization).name, slug)
	if err != nil {
		return err
	}

	member, _, err := client.Teams.ListTeamMembers(ctx, team.GetID(), nil)
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

	return nil
}

func getGithubTeamBySlug(client *github.Client, org string, slug string) (team *github.Team, err error) {
	opt := &github.ListOptions{PerPage: 10}
	for {
		teams, resp, err := client.Teams.ListTeams(context.TODO(), org, opt)
		if err != nil {
			return team, err
		}

		for _, t := range teams {
			if *t.Slug == slug {
				return t, nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return team, fmt.Errorf("Could not find team with slug: %s", slug)
}
