package github

import (
	"context"
	"errors"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func resourceGithubTeamSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubTeamSettingsCreate,
		ReadContext:   resourceGithubTeamSettingsRead,
		UpdateContext: resourceGithubTeamSettingsUpdate,
		DeleteContext: resourceGithubTeamSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubTeamSettingsImport,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub team id or the GitHub team slug.",
			},
			"team_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The slug of the Team within the Organization.",
			},
			"team_uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.",
			},
			"review_request_delegation": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The algorithm to use when assigning pull requests to team members. Supported values are " + string(githubv4.TeamReviewAssignmentAlgorithmRoundRobin) + " and " + string(githubv4.TeamReviewAssignmentAlgorithmLoadBalance) + ".",
							Default:          string(githubv4.TeamReviewAssignmentAlgorithmRoundRobin),
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{string(githubv4.TeamReviewAssignmentAlgorithmRoundRobin), string(githubv4.TeamReviewAssignmentAlgorithmLoadBalance)}, false)),
						},
						"member_count": {
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"review_request_delegation"},
							Description:  "The number of team members to assign to a pull request.",
							ValidateDiagFunc: validation.ToDiagFunc(validation.All(
								validation.IntAtLeast(1),
							)),
						},
						"notify": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "whether to notify the entire team when at least one member is also assigned to the pull request.",
						},
					},
				},
			},
		},
	}
}

func resourceGithubTeamSettingsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	graphql := meta.v4client

	teamIDString := d.Get("team_id").(string)

	tflog.Debug(ctx, "Resolving team_id to Team node_id and slug", map[string]any{
		"team_id": teamIDString,
	})
	// Given a string that is either a team id or team slug, return the
	// get the basic details of the team including node_id and slug
	nodeId, slug, err := resolveTeamIDs(teamIDString, meta, ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Trace(ctx, "Resolved team_id to Team node_id and slug", map[string]any{
		"node_id": nodeId,
		"slug":    slug,
	})
	d.SetId(nodeId)
	if err = d.Set("team_slug", slug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("team_uid", nodeId); err != nil {
		return diag.FromErr(err)
	}

	reviewRequestDelegation := d.Get("review_request_delegation").([]any)

	var mutation struct {
		UpdateTeamReviewAssignment struct {
			ClientMutationId githubv4.ID `graphql:"clientMutationId"`
		} `graphql:"updateTeamReviewAssignment(input:$input)"`
	}

	if len(reviewRequestDelegation) == 0 {
		tflog.Debug(ctx, "No review request delegation settings provided, disabling review request delegation", map[string]any{
			"team_id":   d.Id(),
			"team_slug": slug,
		})

		err := graphql.Mutate(ctx, &mutation, defaultTeamReviewAssignmentSettings(d.Id()), nil)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		settings := reviewRequestDelegation[0].(map[string]any)

		teamReviewAlgorithm := githubv4.TeamReviewAssignmentAlgorithm(settings["algorithm"].(string))
		updateTeamReviewAssignmentInput := githubv4.UpdateTeamReviewAssignmentInput{
			ID:              d.Id(),
			Enabled:         githubv4.Boolean(true),
			Algorithm:       &teamReviewAlgorithm,
			TeamMemberCount: githubv4.NewInt(githubv4.Int(settings["member_count"].(int))),
			NotifyTeam:      githubv4.NewBoolean(githubv4.Boolean(settings["notify"].(bool))),
		}
		err := graphql.Mutate(ctx, &mutation, updateTeamReviewAssignmentInput, nil)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceGithubTeamSettingsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	graphql := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)

	query := queryTeamSettings{}
	variables := map[string]any{
		"slug":  githubv4.String(teamSlug),
		"login": githubv4.String(orgName),
	}

	err := graphql.Query(ctx, &query, variables)
	if err != nil {
		return diag.FromErr(err)
	}

	if query.Organization.Team.ReviewRequestDelegation {
		reviewRequestDelegation := make(map[string]any)
		reviewRequestDelegation["algorithm"] = query.Organization.Team.ReviewRequestDelegationAlgorithm
		reviewRequestDelegation["member_count"] = query.Organization.Team.ReviewRequestDelegationCount
		reviewRequestDelegation["notify"] = query.Organization.Team.ReviewRequestDelegationNotifyAll
		if err = d.Set("review_request_delegation", []any{reviewRequestDelegation}); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err = d.Set("review_request_delegation", []any{}); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubTeamSettingsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if d.HasChange("review_request_delegation") {
		meta := m.(*Owner)
		graphql := meta.v4client
		reviewRequestDelegation := d.Get("review_request_delegation").([]any)

		var mutation struct {
			UpdateTeamReviewAssignment struct {
				ClientMutationId githubv4.ID `graphql:"clientMutationId"`
			} `graphql:"updateTeamReviewAssignment(input:$input)"`
		}

		if len(reviewRequestDelegation) == 0 {
			tflog.Debug(ctx, "No review request delegation settings provided, disabling review request delegation", map[string]any{
				"team_id":   d.Id(),
				"team_slug": d.Get("team_slug").(string),
			})

			err := graphql.Mutate(ctx, &mutation, defaultTeamReviewAssignmentSettings(d.Id()), nil)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			settings := reviewRequestDelegation[0].(map[string]any)

			teamReviewAlgorithm := githubv4.TeamReviewAssignmentAlgorithm(settings["algorithm"].(string))
			updateTeamReviewAssignmentInput := githubv4.UpdateTeamReviewAssignmentInput{
				ID:              d.Id(),
				Enabled:         githubv4.Boolean(true),
				Algorithm:       &teamReviewAlgorithm,
				TeamMemberCount: githubv4.NewInt(githubv4.Int(settings["member_count"].(int))),
				NotifyTeam:      githubv4.NewBoolean(githubv4.Boolean(settings["notify"].(bool))),
			}
			err := graphql.Mutate(ctx, &mutation, updateTeamReviewAssignmentInput, nil)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubTeamSettingsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	graphql := meta.(*Owner).v4client

	var mutation struct {
		UpdateTeamReviewAssignment struct {
			ClientMutationId githubv4.ID `graphql:"clientMutationId"`
		} `graphql:"updateTeamReviewAssignment(input:$input)"`
	}

	err := graphql.Mutate(ctx, &mutation, defaultTeamReviewAssignmentSettings(d.Id()), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubTeamSettingsImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	nodeId, slug, err := resolveTeamIDs(d.Id(), meta.(*Owner), ctx)
	if err != nil {
		return nil, err
	}
	if err = d.Set("team_id", d.Id()); err != nil {
		return nil, err
	}
	d.SetId(nodeId)
	if err = d.Set("team_slug", slug); err != nil {
		return nil, err
	}
	if err = d.Set("team_uid", nodeId); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resolveTeamIDs(idOrSlug string, meta *Owner, ctx context.Context) (nodeId, slug string, err error) {
	client := meta.v3client
	orgName := meta.name
	orgId := meta.id

	teamId, parseIntErr := strconv.ParseInt(idOrSlug, 10, 64)
	if parseIntErr != nil {
		// The given id not an integer, assume it is a team slug
		team, _, slugErr := client.Teams.GetTeamBySlug(ctx, orgName, idOrSlug)
		if slugErr != nil {
			return "", "", errors.New(parseIntErr.Error() + slugErr.Error())
		}
		return team.GetNodeID(), team.GetSlug(), nil
	} else {
		// The given id is an integer, assume it is a team id
		team, _, teamIdErr := client.Teams.GetTeamByID(ctx, orgId, teamId)
		if teamIdErr != nil {
			// There isn't a team with the given ID, assume it is a teamslug
			team, _, slugErr := client.Teams.GetTeamBySlug(ctx, orgName, idOrSlug)
			if slugErr != nil {
				return "", "", errors.New(teamIdErr.Error() + slugErr.Error())
			}

			return team.GetNodeID(), team.GetSlug(), nil
		}

		return team.GetNodeID(), team.GetSlug(), nil
	}
}

func defaultTeamReviewAssignmentSettings(id string) githubv4.UpdateTeamReviewAssignmentInput {
	roundRobinAlgo := githubv4.TeamReviewAssignmentAlgorithmRoundRobin
	return githubv4.UpdateTeamReviewAssignmentInput{
		ID:              id,
		Enabled:         githubv4.Boolean(false),
		Algorithm:       &roundRobinAlgo,
		TeamMemberCount: githubv4.NewInt(githubv4.Int(1)),
		NotifyTeam:      githubv4.NewBoolean(true),
	}
}

type queryTeamSettings struct {
	Organization struct {
		Team struct {
			Name                             string `graphql:"name"`
			Slug                             string `graphql:"slug"`
			ID                               string `graphql:"id"`
			ReviewRequestDelegation          bool   `graphql:"reviewRequestDelegationEnabled"`
			ReviewRequestDelegationAlgorithm string `graphql:"reviewRequestDelegationAlgorithm"`
			ReviewRequestDelegationCount     int    `graphql:"reviewRequestDelegationMemberCount"`
			ReviewRequestDelegationNotifyAll bool   `graphql:"reviewRequestDelegationNotifyTeam"`
		} `graphql:"team(slug:$slug)"`
	} `graphql:"organization(login:$login)"`
}
