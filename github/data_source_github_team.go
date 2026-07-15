package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v89/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubTeamRead,

		Description: "Data source to lookup a team.",

		Schema: map[string]*schema.Schema{
			"team_id": {
				Description:      "ID of the team. One of `team_id` or `slug` must be specified.",
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"team_id", "slug"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
			},
			"slug": {
				Description:  "Slug of the team name. One of `team_id` or `slug` must be specified.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"team_id", "slug"},
			},
			"summary_only": {
				Description: "If true, non-default team details such as `members` & `repositories` will be omitted.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"membership_type": {
				Description:      "If `summary_only` is `false` this controls which members are returned; this can be set to either `all` or `immediate`.",
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "all",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "immediate"}, false)),
			},
			"node_id": {
				Description: "Node ID of the team.",
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
				Description: "List of members of the team.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"repositories": {
				Description: "List of repositories the team has access to.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Deprecated:  "The `repositories` attribute is deprecated and will be removed in a future version of the provider. Use `repositories_detailed` instead.",
			},
			"repositories_detailed": {
				Description: "List of repositories the team has access to.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo_id": {
							Description: "ID of the repository.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"repo_name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"role_name": {
							Description: "Role the team has for the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"results_per_page": {
				Description:      "This is unused and will be removed in a future version of the provider.",
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          100,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
				Deprecated:       "The `results_per_page` argument is deprecated and will be removed in a future version of the provider.",
			},
		},
	}
}

func dataSourceGithubTeamRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	summaryOnly, _ := d.Get("summary_only").(bool)

	var team *github.Team
	if v, ok := d.GetOk("team_id"); ok {
		teamIDInt, _ := v.(int)
		teamID := int64(teamIDInt)
		t, _, err := client.Teams.GetTeamByID(ctx, meta.id, teamID)
		if err != nil {
			return diag.FromErr(err)
		}
		team = t
	} else {
		slug, _ := d.Get("slug").(string)
		t, _, err := client.Teams.GetTeamBySlug(ctx, meta.name, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		team = t
	}

	t := map[string]any{
		"team_id":              int(team.GetID()),
		"slug":                 team.GetSlug(),
		"node_id":              team.GetNodeID(),
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
	}

	var members, repositories []string
	var repositoriesDetailed []map[string]any
	if !summaryOnly {
		membershipType, _ := d.Get("membership_type").(string)

		for member, err := range client.Teams.ListTeamMembersBySlugIter(ctx, meta.name, team.GetSlug(), &github.TeamListTeamMembersOptions{ListOptions: github.ListOptions{PerPage: maxPerPage}}) {
			if err != nil {
				return diag.FromErr(err)
			}

			if membershipType == "immediate" && member.GetInherited() {
				continue
			}

			members = append(members, member.GetLogin())
		}

		for repo, err := range client.Teams.ListTeamReposBySlugIter(ctx, meta.name, team.GetSlug(), &github.ListOptions{PerPage: maxPerPage}) {
			if err != nil {
				return diag.FromErr(err)
			}

			repositories = append(repositories, repo.GetName())

			repositoriesDetailed = append(repositoriesDetailed, map[string]any{
				"repo_id":   repo.GetID(),
				"repo_name": repo.GetName(),
				"role_name": repo.GetRoleName(),
			})
		}
	}

	t["members"] = members
	t["repositories"] = repositories
	t["repositories_detailed"] = repositoriesDetailed

	d.SetId(strconv.FormatInt(team.GetID(), 10))

	for k, v := range t {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
