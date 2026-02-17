package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v82/github"

	"github.com/shurcooL/githubv4"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubTeamRead,

		Schema: map[string]*schema.Schema{
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The team slug.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The team's full name.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The team's description.",
			},
			"privacy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The team's privacy type (closed or secret).",
			},
			"notification_setting": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The team's notification setting (notifications_enabled or notifications_disabled).",
			},
			"permission": {
				Type:        schema.TypeString,
				Computed:    true,
				Deprecated:  "Closing down notice.",
				Description: "The permission that new repositories will be added to the team with.",
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of team members (GitHub usernames). Not returned if summary_only is true.",
			},
			"repositories": {
				Deprecated:  "Use repositories_detailed instead.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of team repositories (repo names). Not returned if summary_only is true.",
			},
			"repositories_detailed": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of team repositories with detailed information. Not returned if summary_only is true.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the repository.",
						},
						"repo_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the repository.",
						},
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The team's role name on the repository (pull, triage, push, maintain, admin).",
						},
					},
				},
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Node ID of the team.",
			},
			"membership_type": {
				Type:             schema.TypeString,
				Default:          "all",
				Optional:         true,
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"all", "immediate"}, false), "membership_type"),
				Description:      "Type of membership to be requested to fill the list of members (all or immediate).",
			},
			"summary_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Exclude the members and repositories of the team from the returned result.",
			},
			"results_per_page": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          100,
				ValidateDiagFunc: toDiagFunc(validation.IntBetween(0, 100), "results_per_page"),
				Deprecated:       "This is deprecated and will be removed in a future release.",
				Description:      "Set the number of results per REST API query (0-100).",
			},
		},
	}
}

func dataSourceGithubTeamRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	slug := d.Get("slug").(string)

	team, _, err := client.Teams.GetTeamBySlug(ctx, meta.(*Owner).name, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(team.GetID(), 10))
	if err = d.Set("name", team.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", team.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("privacy", team.GetPrivacy()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("notification_setting", team.GetNotificationSetting()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("permission", team.GetPermission()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("node_id", team.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}

	var members []string
	var repositories []string
	var repositories_detailed []any

	summaryOnly := d.Get("summary_only").(bool)
	if !summaryOnly {
		resultsPerPage := d.Get("results_per_page").(int)
		options := github.TeamListTeamMembersOptions{
			ListOptions: github.ListOptions{
				PerPage: resultsPerPage,
			},
		}

		if d.Get("membership_type").(string) == "all" {
			for {
				member, resp, err := client.Teams.ListTeamMembersBySlug(ctx, owner, team.GetSlug(), &options)
				if err != nil {
					return diag.FromErr(err)
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
			variables := map[string]any{
				"owner":        githubv4.String(meta.(*Owner).name),
				"slug":         githubv4.String(slug),
				"memberCursor": (*githubv4.String)(nil),
			}
			client := meta.(*Owner).v4client
			for {
				nameErr := client.Query(ctx, &query, variables)
				if nameErr != nil {
					return diag.FromErr(nameErr)
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

		repositories_detailed = make([]any, 0, resultsPerPage)
		for {
			repository, resp, err := client.Teams.ListTeamReposBySlug(ctx, owner, team.GetSlug(), &options.ListOptions)
			if err != nil {
				return diag.FromErr(err)
			}

			for _, v := range repository {
				repositories = append(repositories, v.GetName())
				repositories_detailed = append(repositories_detailed, map[string]any{
					"repo_id":   v.GetID(),
					"repo_name": v.GetName(),
					"role_name": v.GetRoleName(),
				})
			}

			if resp.NextPage == 0 {
				break
			}
			options.Page = resp.NextPage
		}
	}

	if err = d.Set("members", members); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repositories", repositories); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repositories_detailed", repositories_detailed); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
