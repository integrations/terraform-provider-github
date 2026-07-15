package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubOrganizationRoleTeams() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRoleTeamsRead,

		Description: "Data source to list all teams assigned to a custom organization role.",

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description:      "ID of the organization role.",
				Type:             schema.TypeInt,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
			},
			"teams": {
				Description: "Teams assigned to the organization role.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the team.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"team_id": {
							Description: "ID of the team.",
							Type:        schema.TypeInt,
							Computed:    true,
							Deprecated:  "The `team_id` attribute is deprecated and will be removed in a future version of the provider. Use `id` instead.",
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
						"description": {
							Description: "Description of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "Ownership type of the team; one of `enterprise` or `organization`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"privacy": {
							Description: "Privacy level of the team; one of `secret` or `closed`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"permission": {
							Description: "Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"assignment": {
							Description: "Relationship a team has with a role; one of `direct`, `indirect`, or `mixed`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"parent_team": {
							Description: "Parent team; only set if this team is not a root team.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Description: "ID of the parent team.",
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"slug": {
										Description: "Slug of the parent team name.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRoleTeamsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	roleIdInt, _ := d.Get("role_id").(int)
	roleID := int64(roleIdInt)

	teams := make([]any, 0)
	for team, err := range meta.v3client.Organizations.ListTeamsAssignedToOrgRoleIter(ctx, meta.name, roleID, &github.ListOptions{PerPage: maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}

		t := map[string]any{
			"id":          int(team.GetID()),
			"team_id":     int(team.GetID()),
			"slug":        team.GetSlug(),
			"name":        team.GetName(),
			"description": team.GetDescription(),
			"type":        team.GetType(),
			"privacy":     team.GetPrivacy(),
			"permission":  team.GetPermission(),
			"assignment":  team.GetAssignment(),
		}

		if team.Parent != nil {
			t["parent_team"] = []map[string]any{
				{
					"id":   int(team.Parent.GetID()),
					"slug": team.Parent.GetSlug(),
				},
			}
		}

		teams = append(teams, t)
	}

	d.SetId(strconv.FormatInt(roleID, 10))

	if err := d.Set("teams", teams); err != nil {
		return diag.Errorf("error setting teams: %v", err)
	}

	return nil
}
