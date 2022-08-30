package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v47/github"

	"github.com/shurcooL/githubv4"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"membership_type": {
				Type:         schema.TypeString,
				Default:      "all",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"all", "immediate"}, false),
			},
		},
	}
}

func dataSourceGithubTeamRead(d *schema.ResourceData, meta interface{}) error {
	slug := d.Get("slug").(string)

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id
	ctx := context.Background()

	team, _, err := client.Teams.GetTeamBySlug(ctx, meta.(*Owner).name, slug)
	if err != nil {
		return err
	}

	options := github.TeamListTeamMembersOptions{
		ListOptions: github.ListOptions{
			PerPage: maxPerPage,
		},
	}

	var members []string
	if d.Get("membership_type").(string) == "all" {
		for {
			member, resp, err := client.Teams.ListTeamMembersByID(ctx, orgId, team.GetID(), &options)
			if err != nil {
				return err
			}

			for _, v := range member {
				members = append(members, v.GetLogin())
			}

			if resp.NextPage == 0 {
				break
			}
			options.Page = resp.NextPage
		}
	} else {
		type member struct {
			Login string
		}
		var query struct {
			Organization struct {
				Team struct {
					Members struct {
						Nodes    []member
						PageInfo struct {
							EndCursor   githubv4.String
							HasNextPage bool
						}
					} `graphql:"members(first:100,after:$memberCursor,membership:IMMEDIATE)"`
				} `graphql:"team(slug:$slug)"`
			} `graphql:"organization(login:$owner)"`
		}
		variables := map[string]interface{}{
			"owner":        githubv4.String(meta.(*Owner).name),
			"slug":         githubv4.String(slug),
			"memberCursor": (*githubv4.String)(nil),
		}
		client := meta.(*Owner).v4client
		for {
			nameErr := client.Query(ctx, &query, variables)
			if nameErr != nil {
				return nameErr
			}
			for _, v := range query.Organization.Team.Members.Nodes {
				members = append(members, v.Login)
			}
			if query.Organization.Team.Members.PageInfo.HasNextPage {
				variables["memberCursor"] = query.Organization.Team.Members.PageInfo.EndCursor
			} else {
				break
			}
		}
	}
	var repositories []string
	for {
		repository, resp, err := client.Teams.ListTeamReposByID(ctx, orgId, team.GetID(), &options.ListOptions)
		if err != nil {
			return err
		}

		for _, v := range repository {
			repositories = append(repositories, v.GetName())
		}

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(strconv.FormatInt(team.GetID(), 10))
	d.Set("name", team.GetName())
	d.Set("members", members)
	d.Set("repositories", repositories)
	d.Set("description", team.GetDescription())
	d.Set("privacy", team.GetPrivacy())
	d.Set("permission", team.GetPermission())
	d.Set("node_id", team.GetNodeID())

	return nil
}
