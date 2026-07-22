package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubOrganizationTeams() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationTeamsRead,

		Description: "Data source to list all organization teams.",

		Schema: map[string]*schema.Schema{
			"root_teams_only": {
				Description: "If true, only root teams (teams without a parent) will be returned.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"summary_only": {
				Description: "If true, non-default team details such as `members` & `repositories` will be omitted.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"results_per_page": {
				Description:      "This is unused and will be removed in a future version of the provider.",
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          100,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
				Deprecated:       "The `results_per_page` argument is deprecated and will be removed in a future version of the provider.",
			},
			"teams": {
				Description: "Organization teams.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the team.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"node_id": {
							Description: "Node ID of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"slug": {
							Description: "Slug of the team name.",
							Type:        schema.TypeString,
							Required:    true,
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
						"notification_setting": {
							Description: "Notification setting for the team; one of `notifications_enabled`, or `notifications_disabled`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"permission": {
							Description: "Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.",
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
						"members": {
							Description: "List of members in the team.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"repositories": {
							Description: "List of repositories the team has access to.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"parent": {
							Description: "Map of parent team attributes; only set if this team is not a root team.",
							Type:        schema.TypeMap,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Deprecated:  "The `parent` attribute is deprecated and will be removed in a future version of the provider. Use `parent_team` instead.",
						},
						"parent_team_id": {
							Description: "ID of the parent team; only set if this team is not a root team.",
							Type:        schema.TypeString,
							Computed:    true,
							Deprecated:  "The `parent_team_id` attribute is deprecated and will be removed in a future version of the provider. Use `parent_team` instead.",
						},
						"parent_team_slug": {
							Description: "Slug of the parent team; only set if this team is not a root team.",
							Type:        schema.TypeString,
							Computed:    true,
							Deprecated:  "The `parent_team_slug` attribute is deprecated and will be removed in a future version of the provider. Use `parent_team` instead.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationTeamsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	rootTeamsOnly, _ := d.Get("root_teams_only").(bool)
	summaryOnly, _ := d.Get("summary_only").(bool)

	teams := make([]map[string]any, 0)
	for team, err := range meta.v3client.Teams.ListTeamsIter(ctx, meta.name, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}

		if rootTeamsOnly && team.Parent != nil {
			continue
		}

		t := map[string]any{
			"id":                   int(team.GetID()),
			"node_id":              team.GetNodeID(),
			"slug":                 team.GetSlug(),
			"name":                 team.GetName(),
			"description":          team.GetDescription(),
			"type":                 team.GetType(),
			"privacy":              team.GetPrivacy(),
			"notification_setting": team.GetNotificationSetting(),
			"permission":           team.GetPermission(),
		}

		if team.Parent != nil {
			t["parent_team"] = []map[string]any{
				{
					"id":   int(team.Parent.GetID()),
					"slug": team.Parent.GetSlug(),
				},
			}

			t["parent"] = map[string]any{
				"id":   team.Parent.GetNodeID(),
				"slug": team.Parent.GetSlug(),
				"name": team.Parent.GetName(),
			}

			t["parent_team_id"] = strconv.FormatInt(team.Parent.GetID(), 10)
			t["parent_team_slug"] = team.Parent.GetSlug()
		} else {
			t["parent_team"] = nil

			t["parent"] = map[string]any{
				"id":   "",
				"slug": "",
				"name": "",
			}

			t["parent_team_id"] = ""
			t["parent_team_slug"] = ""
		}

		if !summaryOnly {
			var members, repositories []string

			for member, err := range meta.v3client.Teams.ListTeamMembersBySlugIter(ctx, meta.name, team.GetSlug(), &github.TeamListTeamMembersOptions{ListOptions: github.ListOptions{PerPage: meta.maxPerPage}}) {
				if err != nil {
					if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
						tflog.Warn(ctx, "Team members are not accessible, this is likely because the team has been deleted.", map[string]any{"team": team.GetSlug()})
						continue
					}
					return diag.FromErr(err)
				}

				if member.GetInherited() {
					continue
				}

				members = append(members, member.GetLogin())
			}

			for repo, err := range meta.v3client.Teams.ListTeamReposBySlugIter(ctx, meta.name, team.GetSlug(), &github.ListOptions{PerPage: meta.maxPerPage}) {
				if err != nil {
					if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
						tflog.Warn(ctx, "Team repositories are not accessible, this is likely because the team has been deleted.", map[string]any{"team": team.GetSlug()})
						continue
					}
					return diag.FromErr(err)
				}

				repositories = append(repositories, repo.GetName())
			}

			t["members"] = members
			t["repositories"] = repositories
		}

		teams = append(teams, t)
	}

	d.SetId(meta.name)

	if err := d.Set("teams", teams); err != nil {
		return diag.Errorf("error setting teams: %v", err)
	}

	return nil
}
