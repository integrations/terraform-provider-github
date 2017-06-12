package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGithubTeam() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTeamRead,

		Schema: map[string]*schema.Schema{
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"privacy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	slug := d.Get("slug").(string)
	log.Printf("[INFO] Refreshing Gitub Team: %s", slug)

	client := meta.(*Organization).client
	ctx := context.Background()

	team, err := getGithubTeamBySlug(client, meta.(*Organization).name, slug)
	if err != nil {
		return err
	}

	member, _, err := client.Organizations.ListTeamMembers(ctx, team.GetID(), nil)
	if err != nil {
		return err
	}

	members := []string{}
	for _, v := range member {
		members = append(members, v.GetLogin())
	}

	d.SetId(strconv.Itoa(team.GetID()))
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
		teams, resp, err := client.Organizations.ListTeams(context.TODO(), org, opt)
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
