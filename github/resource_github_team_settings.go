package github

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

// getBatchUserNodeIds retrieves the GraphQL node IDs for multiple usernames in a single request.
func getBatchUserNodeIds(ctx context.Context, meta any, usernames []string) (map[string]string, error) {
	if len(usernames) == 0 {
		return make(map[string]string), nil
	}

	client := meta.(*Owner).v4client

	// Create GraphQL variables and query struct using reflection (similar to data_source_github_users.go)
	type UserFragment struct {
		ID string `graphql:"id"`
	}

	var fields []reflect.StructField
	variables := make(map[string]any)

	for idx, username := range usernames {
		label := fmt.Sprintf("User%d", idx)
		variables[label] = githubv4.String(username)
		fields = append(fields, reflect.StructField{
			Name: label,
			Type: reflect.TypeFor[UserFragment](),
			Tag:  reflect.StructTag(fmt.Sprintf("graphql:\"%[1]s: user(login: $%[1]s)\"", label)),
		})
	}

	query := reflect.New(reflect.StructOf(fields)).Elem()

	err := client.Query(ctx, query.Addr().Interface(), variables)
	if err != nil && !strings.Contains(err.Error(), "Could not resolve to a User with the login of") {
		return nil, fmt.Errorf("failed to query users in batch: %w", err)
	}

	result := make(map[string]string)
	for idx, username := range usernames {
		label := fmt.Sprintf("User%d", idx)
		user := query.FieldByName(label).Interface().(UserFragment)
		if user.ID != "" {
			result[username] = user.ID
		}
	}

	return result, nil
}

func resourceGithubTeamSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamSettingsCreate,
		Read:   resourceGithubTeamSettingsRead,
		Update: resourceGithubTeamSettingsUpdate,
		Delete: resourceGithubTeamSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubTeamSettingsImport,
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
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The algorithm to use when assigning pull requests to team members. Supported values are 'ROUND_ROBIN' and 'LOAD_BALANCE'.",
							Default:     "ROUND_ROBIN",
							ValidateDiagFunc: toDiagFunc(func(v any, key string) (we []string, errs []error) {
								algorithm, ok := v.(string)
								if !ok {
									return nil, []error{fmt.Errorf("expected type of %s to be string", key)}
								}

								if algorithm != "ROUND_ROBIN" && algorithm != "LOAD_BALANCE" {
									errs = append(errs, errors.New("review request delegation algorithm must be one of [\"ROUND_ROBIN\", \"LOAD_BALANCE\"]"))
								}

								return we, errs
							}, "algorithm"),
						},
						"member_count": {
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"review_request_delegation"},
							Description:  "The number of team members to assign to a pull request.",
							ValidateDiagFunc: toDiagFunc(func(v any, key string) (we []string, errs []error) {
								count, ok := v.(int)
								if !ok {
									return nil, []error{fmt.Errorf("expected type of %s to be an integer", key)}
								}
								if count <= 0 {
									errs = append(errs, errors.New("review request delegation reviewer count must be a positive number"))
								}
								return we, errs
							}, "member_count"),
						},
						"notify": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "whether to notify the entire team when at least one member is also assigned to the pull request.",
						},
						"excluded_members": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of team member usernames to exclude from the PR review process.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

		// NOTE: The exclusion list is not available via the GraphQL read query yet.
		// The excluded_team_member_node_ids field can be set but cannot be read back from the GitHub API.
		// This is because the GraphQL API for team review assignments is currently in preview.
		// As a workaround, we preserve the excluded_members from the current state.
		if currentDelegation := d.Get("review_request_delegation").([]any); len(currentDelegation) > 0 {
			if currentSettings, ok := currentDelegation[0].(map[string]any); ok {
				if excludedMembers, exists := currentSettings["excluded_members"]; exists {
					reviewRequestDelegation["excluded_members"] = excludedMembers
				}
			}
		}

		if err = d.Set("review_request_delegation", []any{reviewRequestDelegation}); err != nil {
			return err
		}
	} else {
		if err = d.Set("review_request_delegation", []any{}); err != nil {
			return err
		}
	}
	// NOTE: The excluded members are preserved from the current state in the read logic above
	// since the GitHub API doesn't currently support reading them back.

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

			exclusionList := make([]githubv4.ID, 0)
			if excludedMembers, ok := settings["excluded_members"]; ok && excludedMembers != nil {
				// Collect all usernames first
				usernames := make([]string, 0)
				for _, v := range excludedMembers.(*schema.Set).List() {
					if v != nil {
						username := v.(string)
						usernames = append(usernames, username)
					}
				}

				// Get all node IDs in a single batch request
				if len(usernames) > 0 {
					nodeIds, err := getBatchUserNodeIds(ctx, meta, usernames)
					if err != nil {
						return fmt.Errorf("failed to get node IDs for excluded members: %w", err)
					}

					// Convert to the exclusion list
					for _, username := range usernames {
						if nodeId, exists := nodeIds[username]; exists {
							exclusionList = append(exclusionList, githubv4.ID(nodeId))
						} else {
							return fmt.Errorf("failed to get node ID for user %s: user not found", username)
						}
					}
				}
			}

			return graphql.Mutate(ctx, &mutation, UpdateTeamReviewAssignmentInput{
				TeamID:                           d.Id(),
				ReviewRequestDelegation:          true,
				ReviewRequestDelegationAlgorithm: settings["algorithm"].(string),
				ReviewRequestDelegationCount:     settings["member_count"].(int),
				ReviewRequestDelegationNotifyAll: settings["notify"].(bool),
				ExcludedTeamMemberIds:            exclusionList,
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

func resourceGithubTeamSettingsImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	nodeId, slug, err := resolveTeamIDs(d.Id(), meta.(*Owner), context.Background())
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
	return []*schema.ResourceData{d}, resourceGithubTeamSettingsRead(d, meta)
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

type UpdateTeamReviewAssignmentInput struct {
	ClientMutationID                 string        `json:"clientMutationId,omitempty"`
	TeamID                           string        `graphql:"id" json:"id"`
	ReviewRequestDelegation          bool          `graphql:"enabled" json:"enabled"`
	ReviewRequestDelegationAlgorithm string        `graphql:"algorithm" json:"algorithm"`
	ReviewRequestDelegationCount     int           `graphql:"teamMemberCount" json:"teamMemberCount"`
	ReviewRequestDelegationNotifyAll bool          `graphql:"notifyTeam" json:"notifyTeam"`
	ExcludedTeamMemberIds            []githubv4.ID `graphql:"excludedTeamMemberIds" json:"excludedTeamMemberIds"`
}

func defaultTeamReviewAssignmentSettings(id string) UpdateTeamReviewAssignmentInput {
	return UpdateTeamReviewAssignmentInput{
		TeamID:                           id,
		ReviewRequestDelegation:          false,
		ReviewRequestDelegationAlgorithm: "ROUND_ROBIN",
		ReviewRequestDelegationCount:     1,
		ReviewRequestDelegationNotifyAll: true,
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
