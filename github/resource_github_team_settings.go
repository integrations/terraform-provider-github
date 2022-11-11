package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v48/github"
	"github.com/google/uuid"
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
	return resourceGithubTeamSettingsRead(d, meta)

}

func resourceGithubTeamSettingsRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	graphql := meta.(*Owner).v4client
	rest := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	id, ok := d.Get("team_id").(string)
	if !ok {
		return errors.New("team_id must be provided as a string")
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	teamSlug := ""
	teamId, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		// If err != nil then team_id was provided with the numerical ID, this must be converted to the team slug
		team, _, err := rest.Teams.GetTeamByID(ctx, orgId, teamId)
		if err != nil {
			if ghErr, ok := err.(*github.ErrorResponse); ok {
				if ghErr.Response.StatusCode == http.StatusNotModified {
					return nil
				}
				if ghErr.Response.StatusCode == http.StatusNotFound {
					log.Printf("[INFO] Removing team %s from state because it no longer exists in GitHub",
						d.Id())
					d.SetId("")
					return nil
				}
			}
			return err
		}
		teamSlug = *team.Slug
	} else {
		teamSlug = id
	}

	orgName := meta.(*Owner).name

	var query struct {
		Organization struct {
			Team struct {
				Name                             string `graphql:"name"`
				ReviewRequestDelegation          bool   `graphql:"reviewRequestDelegationEnabled"`
				ReviewRequestDelegationAlgorithm string `graphql:"reviewRequestDelegationAlgorithm"`
				ReviewRequestDelegationCount     int    `graphql:"reviewRequestDelegationMemberCount"`
				ReviewRequestDelegationNotifyAll bool   `graphql:"reviewRequestDelegationNotifyTeam"`
			} `graphql:"team(slug:$slug)"`
		} `graphql:"organization(login:$login)"`
	}
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
		"review_request_notify") {
		ctx := context.WithValue(context.Background(), ctxId, d.Id())
		graphql := meta.(*Owner).v4client

		var mutation struct {
			UpdateTeamReviewAssignment struct {
				Team struct {
					Name                             string `graphql:"name"`
					ID                               string `graphql:"id"`
					ReviewRequestDelegation          bool   `graphql:"reviewRequestDelegationEnabled"`
					ReviewRequestDelegationAlgorithm string `graphql:"reviewRequestDelegationAlgorithm"`
					ReviewRequestDelegationCount     int    `graphql:"reviewRequestDelegationMemberCount"`
					ReviewRequestDelegationNotifyAll bool   `graphql:"reviewRequestDelegationNotifyTeam"`
				} `graphql:"team"`
			} `graphql:"updateTeamReviewAssignment(input:$input)"`
		}

		e := graphql.Mutate(ctx, mutation, updateTeamReviewAssignment{
			MutationID:                       uuid.NewString(),
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
	return nil
}

type updateTeamReviewAssignment struct {
	MutationID                       string `graphql:"clientMutationId"`
	TeamID                           string `graphql:"id"`
	ReviewRequestDelegation          bool   `graphql:"enabled"`
	ReviewRequestDelegationAlgorithm string `graphql:"algorithm"`
	ReviewRequestDelegationCount     int    `graphql:"teamMemberCount"`
	ReviewRequestDelegationNotifyAll bool   `graphql:"notifyTeam"`
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
