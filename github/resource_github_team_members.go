package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func resourceGithubTeamMembers() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubTeamMembersCreate,
		ReadContext:   resourceGithubTeamMembersRead,
		UpdateContext: resourceGithubTeamMembersUpdate,
		DeleteContext: resourceGithubTeamMembersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubTeamMembersImport,
		},

		CustomizeDiff: customdiff.Sequence(diffLegacyTeamID, diffLegacyTeam),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubTeamMembersV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubTeamMembersStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to authoritatively manage GitHub team members.",

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				Description:      "Slug of the GitHub team to manage membership for.",
			},
			"team_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
				Description:      "ID or slug of the GitHub team to manage membership for.",
				Deprecated:       "Use `team_slug` instead; this field will be made computed only in a future version of the provider.",
			},
			"members": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of users that should be members of the team.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
							Description:      "User to add to the team.",
							// This seems to be the only way to ensure that the username is in lowercase.
							// Without this the tests fail because the value is compared in a case-sensitive manner.
							StateFunc: func(v any) string {
								val, _ := v.(string)
								return strings.ToLower(val)
							},
						},
						"role": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "member",
							Description:      "Role to grant the user within the team; must be one of `member` or `maintainer`.",
							ValidateDiagFunc: validateValueFunc([]string{"member", "maintainer"}),
						},
					},
				},
			},
		},
	}
}

func resourceGithubTeamMembersCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)
	teamIDString, _ := d.Get("team_id").(string)
	members, _ := d.Get("members").(*schema.Set)

	var teamID int64
	if slug == "" {
		team := newLegacyTeamIdentity(teamIDString)
		if s, ok := team.getSlugOK(); ok {
			slug = s
		} else {
			teamID = team.getID()

			s, err := lookupTeamSlug(ctx, client, meta.id, teamID)
			if err != nil {
				return diag.FromErr(err)
			}
			slug = s
		}
	}

	if teamID == 0 {
		id, err := lookupTeamID(ctx, client, orgName, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		teamID = id
	}

	tflog.Debug(ctx, "Creating team members.", map[string]any{"team_slug": slug, "team_id": teamID})

	teamMembers, err := newUserMembers(members.List())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse team members: %w", err))
	}

	if err := updateTeamMembers(ctx, meta, slug, teamMembers); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(teamID, 10))

	if err := d.Set("team_slug", slug); err != nil {
		return diag.FromErr(err)
	}

	if teamIDString == "" {
		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Debug(ctx, "Created team members.", map[string]any{"team_slug": slug, "team_id": teamID})

	return nil
}

func resourceGithubTeamMembersRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)

	teamIDStr, _ := d.Get("team_id").(string)
	if teamIDStr == "" {
		tflog.Debug(ctx, "Looking up team ID from slug.", map[string]any{"team_slug": slug})

		teamID, err := lookupTeamID(ctx, client, orgName, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Debug(ctx, "Reading team members.", map[string]any{"team_slug": slug})

	teamMembers, err := getTeamMembers(ctx, meta, slug)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Team not found during read, removing from state.", map[string]any{"team_slug": slug})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("members", teamMembers.flatten()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubTeamMembersUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	slug, _ := d.Get("team_slug").(string)
	members, _ := d.Get("members").(*schema.Set)

	if idStr, _ := d.Get("team_id").(string); idStr == "" {
		tflog.Debug(ctx, "Looking up team ID from slug.", map[string]any{"team_slug": slug})

		teamID, err := lookupTeamID(ctx, client, orgName, slug)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return diag.FromErr(err)
		}
	}

	teamMembers, err := newUserMembers(members.List())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse team members: %w", err))
	}

	if err := updateTeamMembers(ctx, meta, slug, teamMembers); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubTeamMembersDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	slug, _ := d.Get("team_slug").(string)

	tflog.Debug(ctx, "Removing all members from team.", map[string]any{"team_slug": slug})

	if err := updateTeamMembers(ctx, meta, slug, nil); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Team not found during delete, assuming already removed.", map[string]any{"team_slug": slug})
			return nil
		}

		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Removed all members from team.", map[string]any{"team_slug": slug})

	return nil
}

func resourceGithubTeamMembersImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)

	idString := d.Id()

	tflog.Debug(ctx, "Importing team members.", map[string]any{"id": idString})

	team := newLegacyTeamIdentity(idString)
	slug, slugOK := team.getSlugOK()
	teamID := team.getID()
	if slugOK {
		id, err := lookupTeamID(ctx, meta.v3client, meta.name, slug)
		if err != nil {
			return nil, err
		}
		teamID = id

		if err := d.Set("team_id", strconv.FormatInt(teamID, 10)); err != nil {
			return nil, err
		}
	} else {
		s, err := lookupTeamSlug(ctx, meta.v3client, meta.id, teamID)
		if err != nil {
			return nil, err
		}
		slug = s

		if err := d.Set("team_id", idString); err != nil {
			return nil, err
		}
	}

	teamMembers, err := getTeamMembers(ctx, meta, slug)
	if err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(teamID, 10))

	if err := d.Set("team_slug", slug); err != nil {
		return nil, err
	}

	if err := d.Set("members", teamMembers.flatten()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func getTeamMembers(ctx context.Context, meta *Owner, slug string) (userMembers, error) {
	tflog.Debug(ctx, "Getting team members.", map[string]any{"team_slug": slug})

	// Use the GraphQL API with membership:IMMEDIATE so that only direct members
	// of the team are returned. The REST list-members endpoint always includes
	// members inherited from child teams (it has no immediate-only option), which
	// would otherwise surface as drift for organizations that use nested teams.
	// See https://github.com/integrations/terraform-provider-github/issues/3497
	var query struct {
		Organization struct {
			Team *struct {
				Members struct {
					Edges []struct {
						Node struct {
							Login string
						}
						Role string
					}
					PageInfo struct {
						EndCursor   githubv4.String
						HasNextPage bool
					}
				} `graphql:"members(membership: IMMEDIATE, first: 100, after: $after)"`
			} `graphql:"team(slug: $slug)"`
		} `graphql:"organization(login: $login)"`
	}

	variables := map[string]any{
		"login": githubv4.String(meta.name),
		"slug":  githubv4.String(slug),
		"after": (*githubv4.String)(nil),
	}

	var teamMembers userMembers
	for {
		if err := meta.v4client.Query(ctx, &query, variables); err != nil {
			return nil, err
		}

		// A null team means it no longer exists (e.g. deleted out-of-band). The
		// GraphQL API returns a null node rather than an error in this case, so
		// surface a 404 to match the REST endpoints and let callers remove the
		// resource from state.
		if query.Organization.Team == nil {
			return nil, &github.ErrorResponse{
				Response: &http.Response{StatusCode: http.StatusNotFound},
				Message:  fmt.Sprintf("team %q not found in organization %q", slug, meta.name),
			}
		}

		for _, edge := range query.Organization.Team.Members.Edges {
			teamMembers = append(teamMembers, userMember{
				userIdentity: userIdentity{
					login: strings.ToLower(edge.Node.Login),
				},
				role: strings.ToLower(edge.Role),
			})
		}

		if !query.Organization.Team.Members.PageInfo.HasNextPage {
			break
		}

		variables["after"] = query.Organization.Team.Members.PageInfo.EndCursor
	}

	tflog.Debug(ctx, "Got team members.", map[string]any{"team_slug": slug, "team_members": teamMembers})

	return teamMembers, nil
}

func updateTeamMembers(ctx context.Context, meta *Owner, slug string, wantMembers userMembers) error {
	client := meta.v3client
	orgName := meta.name

	tflog.Debug(ctx, "Updating team members.", map[string]any{"team_slug": slug, "members": wantMembers.flatten()})

	roleLookup := map[string]int{
		"maintainer": 1,
		"member":     2,
	}

	slices.SortFunc(wantMembers, func(a, b userMember) int {
		return roleLookup[a.role] - roleLookup[b.role]
	})

	currentMembers, err := getTeamMembers(ctx, meta, slug)
	if err != nil {
		return err
	}

	lookup := make(map[string]userMember, len(currentMembers))
	want := make(map[string]struct{}, len(wantMembers))

	for _, member := range currentMembers {
		lookup[member.login] = member
	}

	for _, member := range wantMembers {
		login := strings.ToLower(member.login)
		if current, ok := lookup[login]; !ok || current.role != member.role {
			tflog.Debug(ctx, "Adding/updating team member.", map[string]any{"team_slug": slug, "username": login, "role": member.role})

			if _, _, err := client.Teams.AddTeamMembershipBySlug(ctx, orgName, slug, login, &github.TeamAddTeamMembershipOptions{Role: member.role}); err != nil {
				return fmt.Errorf("could not add team member %q: %w", login, err)
			}
		}

		want[login] = struct{}{}
	}

	for _, member := range currentMembers {
		if _, ok := want[member.login]; !ok {
			tflog.Debug(ctx, "Removing team member.", map[string]any{"team_slug": slug, "username": member.login})

			if _, err := client.Teams.RemoveTeamMembershipBySlug(ctx, orgName, slug, member.login); err != nil {
				return fmt.Errorf("could not remove existing team member %q: %w", member.login, err)
			}
		}
	}

	tflog.Debug(ctx, "Updated team members.", map[string]any{"team_slug": slug})
	return nil
}
