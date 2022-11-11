package github

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func resourceGithubTeamSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamSettingsCreate,
		Read:   resourceGithubTeamSettingsRead,
		Update: resourceGithubTeamSettingsUpdate,
		Delete: resourceGithubTeamSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID or slug of team",
			},
			"team_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The slug of the Team within the Organization",
			},
			"team_uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of the Team on GitHub. Corresponds to the ID of the github_team_settings resource",
			},
			"review_request_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Algorithm to use when determining team members to be assigned to a pull request. Allowed values are ROUND_ROBIN and LOAD_BALANCE",
				Default:     "ROUND_ROBIN",
				ValidateFunc: func(v interface{}, key string) (we []string, errs []error) {
					algorithm, ok := v.(string)
					if !ok {
						return nil, []error{fmt.Errorf("expected type of %s to be string", key)}
					}

					if !(algorithm == "ROUND_ROBIN" || algorithm == "LOAD_BALANCE") {
						errs = append(errs, errors.New("review request delegation algorithm must be one of [\"ROUND_ROBIN\", \"LOAD_BALANCE\"]"))
					}

					return we, errs
				},
			},
			"review_request_delegation": {
				Type:         schema.TypeBool,
				Default:      false,
				Optional:     true,
				Description:  "Enable review request delegation for the Team",
				RequiredWith: []string{"review_request_algorithm", "review_request_count", "review_request_notify"},
			},
			"review_request_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"review_request_delegation"},
				Description:  "The number of reviewers to be assigned to a pull request from this team",
				ValidateFunc: func(v interface{}, key string) (we []string, errs []error) {
					count, ok := v.(int)
					if !ok {
						return nil, []error{fmt.Errorf("expected type of %s to be an integer", key)}
					}
					if count <= 0 {
						errs = append(errs, errors.New("review request delegation reviewer count must be a positive number"))
					}
					return we, errs
				},
			},
			"review_request_notify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Notify the entire team when a pull request is assigned to a member of the team",
			},
		},
	}
}

func resourceGithubTeamSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	// Given a string that is either a team id or team slug, return the
	// get the basic details of the team including node_id and slug
	ctx := context.Background()
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	orgId := meta.(*Owner).id

	teamIDString, _ := d.Get("team_id").(string)

	teamId, parseIntErr := strconv.ParseInt(teamIDString, 10, 64)
	if parseIntErr != nil {
		// The given id not an integer, assume it is a team slug
		team, _, slugErr := client.Teams.GetTeamBySlug(ctx, orgName, teamIDString)
		if slugErr != nil {
			return errors.New(parseIntErr.Error() + slugErr.Error())
		}
		d.SetId(team.GetNodeID())
		d.Set("team_slug", team.GetSlug())
	} else {
		// The given id is an integer, assume it is a team id
		team, _, teamIdErr := client.Teams.GetTeamByID(ctx, orgId, teamId)
		if teamIdErr != nil {
			// There isn't a team with the given ID, assume it is a teamslug
			team, _, slugErr := client.Teams.GetTeamBySlug(ctx, orgName, teamIDString)
			if slugErr != nil {
				return errors.New(teamIdErr.Error() + slugErr.Error())
			}
			d.SetId(team.GetNodeID())
			d.Set("team_slug", team.GetSlug())
		}
		d.SetId(team.GetNodeID())
		d.Set("team_slug", team.GetSlug())
	}

	return resourceGithubTeamSettingsUpdate(d, meta)

}

func resourceGithubTeamSettingsRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	graphql := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)

	var query = queryTeamSettings{}
	variables := map[string]interface{}{
		"slug":  githubv4.String(teamSlug),
		"login": githubv4.String(orgName),
	}

	e := graphql.Query(meta.(*Owner).StopContext, &query, variables)
	if e != nil {
		return e
	}

	d.Set("review_request_algorithm", query.Organization.Team.ReviewRequestDelegationAlgorithm)
	d.Set("review_request_delegation", query.Organization.Team.ReviewRequestDelegation)
	d.Set("review_request_count", query.Organization.Team.ReviewRequestDelegationCount)
	d.Set("review_request_notify", query.Organization.Team.ReviewRequestDelegationNotifyAll)

	return nil

}

func resourceGithubTeamSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	if d.HasChanges(
		"review_request_algorithm",
		"review_request_delegation",
		"review_request_count",
		"review_request_notify") || d.IsNewResource() {
		ctx := context.WithValue(context.Background(), ctxId, d.Id())
		graphql := meta.(*Owner).v4client

		var mutation struct {
			UpdateTeamReviewAssignment struct {
				Team struct {
					ClientMutationId githubv4.ID
				} `graphql:"team"`
			} `graphql:"updateTeamReviewAssignment(input:$input)"`
		}

		e := graphql.Mutate(ctx, &mutation, UpdateTeamReviewAssignmentInput{
			TeamID:                           d.Id(),
			ReviewRequestDelegation:          d.Get("review_request_delegation").(bool),
			ReviewRequestDelegationAlgorithm: d.Get("review_request_algorithm").(string),
			ReviewRequestDelegationCount:     d.Get("review_request_count").(int),
			ReviewRequestDelegationNotifyAll: d.Get("review_request_notify").(bool),
		}, nil)
		if e != nil {
			return e
		}
	}

	return resourceGithubTeamSettingsRead(d, meta)

}

func resourceGithubTeamSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	graphql := meta.(*Owner).v4client

	var mutation struct {
		UpdateTeamReviewAssignment struct {
			Team struct {
				ClientMutationId githubv4.ID
			} `graphql:"team"`
		} `graphql:"updateTeamReviewAssignment(input:$input)"`
	}

	return graphql.Mutate(ctx, &mutation, UpdateTeamReviewAssignmentInput{
		TeamID:                           d.Id(),
		ReviewRequestDelegation:          false,
		ReviewRequestDelegationAlgorithm: "ROUND_ROBIN",
		ReviewRequestDelegationCount:     1,
		ReviewRequestDelegationNotifyAll: true,
	}, nil)

}

type UpdateTeamReviewAssignmentInput struct {
	ClientMutationID                 string `json:"clientMutationId,omitempty"`
	TeamID                           string `graphql:"id" json:"id"`
	ReviewRequestDelegation          bool   `graphql:"enabled" json:"enabled"`
	ReviewRequestDelegationAlgorithm string `graphql:"algorithm" json:"algorithm"`
	ReviewRequestDelegationCount     int    `graphql:"teamMemberCount" json:"teamMemberCount"`
	ReviewRequestDelegationNotifyAll bool   `graphql:"notifyTeam" json:"notifyTeam"`
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
