package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v66/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRoleTeams() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationRoleTeamsRead,

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The unique identifier of the organization role.",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"teams": {
				Description: "List of teams assigned to the organization role.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Unique identifier of the team.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"slug": {
							Description: "Slug of the team name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"permission": {
							Description: "Permission that the team will have for its repositories.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						// TODO: Add these fields when go-github adds the functionality to get a custom org
						// See https://github.com/google/go-github/issues/3364
						// "assignment": {
						// 	Description: "Determines if the team has a direct, indirect, or mixed relationship to a role.",
						// 	Type:        schema.TypeString,
						// 	Computed:    true,
						// },
						// "parent_team_id": {
						// 	Description: "The ID of the parent team if this is an indirect assignment.",
						// 	Type:        schema.TypeString,
						// 	Computed:    true,
						// },
						// "parent_team_slug": {
						// 	Description: "The slug of the parent team if this is an indirect assignment.",
						// 	Type:        schema.TypeString,
						// 	Computed:    true,
						// },
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRoleTeamsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))

	allTeams := make([]any, 0)

	opts := &github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		teams, resp, err := client.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, roleId, opts)
		if err != nil {
			return err
		}

		for _, team := range teams {
			t := map[string]any{
				"id":         team.GetID(),
				"slug":       team.GetSlug(),
				"name":       team.GetName(),
				"permission": team.GetPermission(),
			}
			allTeams = append(allTeams, t)
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	d.SetId(fmt.Sprintf("%d", roleId))
	if err := d.Set("teams", allTeams); err != nil {
		return fmt.Errorf("error setting teams: %s", err)
	}

	return nil
}
