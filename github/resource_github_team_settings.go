package github

import (
	"context"
	"errors"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func resourceGithubTeamSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamSettingsCreate,
		Read:   resourceGithubTeamSettingsRead,
		Update: resourceGithubTeamSettingsUpdate,
		Delete: resourceGithubTeamSettingsDelete,
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

func resourceGithubTeamSettingsCreate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	// Given a string that is either a team id or team slug, return the
	// get the basic details of the team including node_id and slug
	ctx := context.Background()

	teamIDString, _ := d.Get("team_id").(string)

	nodeId, slug, err := resolveTeamIDs(teamIDString, meta.(*Owner), ctx)
	if err != nil {
		return err
	}
	d.SetId(nodeId)
	if err = d.Set("team_slug", slug); err != nil {
		return err
	}
	if err = d.Set("team_uid", nodeId); err != nil {
		return err
	}
	return resourceGithubTeamSettingsUpdate(d, meta)
}

func resourceGithubTeamSettingsRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	graphql := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)

	query := queryTeamSettings{}
	variables := map[string]any{
		"slug":  githubv4.String(teamSlug),
		"login": githubv4.String(orgName),
	}

	e := graphql.Query(meta.(*Owner).StopContext, &query, variables)
	if e != nil {
		return e
	}

	if query.Organization.Team.ReviewRequestDelegation {
		reviewRequestDelegation := make(map[string]any)
		reviewRequestDelegation["algorithm"] = query.Organization.Team.ReviewRequestDelegationAlgorithm
		reviewRequestDelegation["member_count"] = query.Organization.Team.ReviewRequestDelegationCount
		reviewRequestDelegation["notify"] = query.Organization.Team.ReviewRequestDelegationNotifyAll
		if err = d.Set("review_request_delegation", []any{reviewRequestDelegation}); err != nil {
			return err
		}
	} else {
		if err = d.Set("review_request_delegation", []any{}); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubTeamSettingsUpdate(d *schema.ResourceData, meta any) error {
	if d.HasChange("review_request_delegation") || d.IsNewResource() {

		ctx := context.WithValue(context.Background(), ctxId, d.Id())
		graphql := meta.(*Owner).v4client
		if setting := d.Get("review_request_delegation").([]any); len(setting) == 0 {
			var mutation struct {
				UpdateTeamReviewAssignment struct {
					ClientMutationId githubv4.ID `graphql:"clientMutationId"`
				} `graphql:"updateTeamReviewAssignment(input:$input)"`
			}

			return graphql.Mutate(ctx, &mutation, defaultTeamReviewAssignmentSettings(d.Id()), nil)
		} else {
			settings := d.Get("review_request_delegation").([]any)[0].(map[string]any)

			var mutation struct {
				UpdateTeamReviewAssignment struct {
					ClientMutationId githubv4.ID `graphql:"clientMutationId"`
				} `graphql:"updateTeamReviewAssignment(input:$input)"`
			}

			teamReviewAlgorithm := githubv4.TeamReviewAssignmentAlgorithm(settings["algorithm"].(string))
			return graphql.Mutate(ctx, &mutation, githubv4.UpdateTeamReviewAssignmentInput{
				ID:              d.Id(),
				Enabled:         githubv4.Boolean(true),
				Algorithm:       &teamReviewAlgorithm,
				TeamMemberCount: githubv4.NewInt(githubv4.Int(settings["member_count"].(int))),
				NotifyTeam:      githubv4.NewBoolean(githubv4.Boolean(settings["notify"].(bool))),
			}, nil)
		}
	}

	return resourceGithubTeamSettingsRead(d, meta)
}

func resourceGithubTeamSettingsDelete(d *schema.ResourceData, meta any) error {
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	graphql := meta.(*Owner).v4client

	var mutation struct {
		UpdateTeamReviewAssignment struct {
			ClientMutationId githubv4.ID `graphql:"clientMutationId"`
		} `graphql:"updateTeamReviewAssignment(input:$input)"`
	}

	return graphql.Mutate(ctx, &mutation, defaultTeamReviewAssignmentSettings(d.Id()), nil)
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
