package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// importEnterpriseTeamSelector preserves the selector used in the import ID.
func importEnterpriseTeamSelector(d *schema.ResourceData, enterpriseSlug, selector string) error {
	if strings.TrimSpace(enterpriseSlug) == "" || strings.TrimSpace(selector) == "" {
		return fmt.Errorf("enterprise and team selector must not be empty")
	}
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}
	id, err := strconv.Atoi(selector)
	if err == nil {
		if id <= 0 {
			return fmt.Errorf("team ID must be positive: %q", selector)
		}
		return d.Set("team_id", id)
	}
	if strings.Trim(selector, "0123456789") == "" {
		return fmt.Errorf("invalid team ID %q: %w", selector, err)
	}
	return d.Set("team_slug", selector)
}

// buildEnterpriseTeamMembershipID creates an ID for enterprise team membership resources.
// Uses "/" as separator because team slugs contain ":" (e.g., "ent:team-name").
// Note: GitHub slugs only allow alphanumeric characters, hyphens, and colons - never "/".
func buildEnterpriseTeamMembershipID(enterpriseSlug, teamSlug, username string) string {
	return fmt.Sprintf("%s/%s/%s", enterpriseSlug, teamSlug, username)
}

// parseEnterpriseTeamMembershipID parses the ID for enterprise team membership resources.
func parseEnterpriseTeamMembershipID(id string) (enterpriseSlug, teamSlug, username string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("unexpected ID format (%q); expected enterprise_slug/team_slug/username", id)
	}
	for _, part := range parts {
		if strings.TrimSpace(part) == "" {
			return "", "", "", fmt.Errorf("enterprise, team selector and username must not be empty: %q", id)
		}
	}
	return parts[0], parts[1], parts[2], nil
}

// buildEnterpriseTeamOrganizationsID creates an ID for enterprise team organizations resources.
// Uses "/" as separator because team slugs contain ":" (e.g., "ent:team-name").
// Note: GitHub slugs only allow alphanumeric characters, hyphens, and colons - never "/".
func buildEnterpriseTeamOrganizationsID(enterpriseSlug, teamSlug string) string {
	return fmt.Sprintf("%s/%s", enterpriseSlug, teamSlug)
}

// parseEnterpriseTeamOrganizationsID parses the ID for enterprise team organizations resources.
func parseEnterpriseTeamOrganizationsID(id string) (enterpriseSlug, teamSlug string, err error) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected ID format (%q); expected enterprise_slug/team_slug", id)
	}
	return parts[0], parts[1], nil
}

// findEnterpriseTeamByID stops paging when it finds the requested team's metadata.
// Unlike membership and assignment endpoints, GetTeam only documents slug lookups.
func findEnterpriseTeamByID(meta *Owner, ctx context.Context, enterpriseSlug string, id int64) (*github.EnterpriseTeam, error) {
	opt := &github.ListOptions{PerPage: meta.maxPerPage}
	for {
		teams, resp, err := meta.v3client.Enterprise.ListTeams(ctx, enterpriseSlug, opt)
		if err != nil {
			return nil, err
		}
		for _, team := range teams {
			if team.ID == id {
				return team, nil
			}
		}
		if resp.NextPage == 0 {
			return nil, nil
		}
		opt.Page = resp.NextPage
	}
}

// storedEnterpriseTeamID returns the numeric team identity recorded in state,
// preferring resolved_team_id and falling back to a configured team_id.
func storedEnterpriseTeamID(d *schema.ResourceData) int64 {
	if v, ok := d.GetOk("resolved_team_id"); ok {
		if id, _ := v.(int); id > 0 {
			return int64(id)
		}
	}
	if v, ok := d.GetOk("team_id"); ok {
		if id, _ := v.(int); id > 0 {
			return int64(id)
		}
	}
	return 0
}

// enterpriseTeamIDForOperations returns the numeric identity to address the
// team by in membership and organization-assignment API calls, which accept
// either a slug or a numeric ID. State with a recorded identity needs no team
// lookup at all; legacy state and slug-based imports bootstrap the ID with a
// single GetTeam. Returns 0 when the team no longer exists.
func enterpriseTeamIDForOperations(meta *Owner, ctx context.Context, enterpriseSlug, slug string, d *schema.ResourceData) (int64, error) {
	if id := storedEnterpriseTeamID(d); id > 0 {
		return id, nil
	}
	team, _, err := meta.v3client.Enterprise.GetTeam(ctx, enterpriseSlug, slug)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			return 0, nil
		}
		return 0, err
	}
	if team == nil {
		return 0, nil
	}
	return team.ID, nil
}

// resolveEnterpriseTeamForCreate resolves the configured selector to the
// team's numeric ID, looking the team up only when team_slug is configured.
// The returned slug is empty when the caller configured team_id.
func resolveEnterpriseTeamForCreate(meta *Owner, ctx context.Context, enterpriseSlug string, d *schema.ResourceData) (int64, string, error) {
	if v, ok := d.GetOk("team_slug"); ok {
		team, _, err := meta.v3client.Enterprise.GetTeam(ctx, enterpriseSlug, v.(string))
		if err != nil {
			return 0, "", err
		}
		if team == nil {
			return 0, "", fmt.Errorf("enterprise team not found")
		}
		return team.ID, team.Slug, nil
	}
	return int64(d.Get("team_id").(int)), "", nil
}

// findEnterpriseTeamByIdentity treats the slug as a hint, never as proof of identity.
func findEnterpriseTeamByIdentity(meta *Owner, ctx context.Context, enterpriseSlug, slug string, id int64) (*github.EnterpriseTeam, error) {
	if slug != "" {
		team, _, err := meta.v3client.Enterprise.GetTeam(ctx, enterpriseSlug, slug)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); !ok || ghErr.Response == nil || ghErr.Response.StatusCode != http.StatusNotFound {
				return nil, err
			}
		} else if team != nil && team.ID == id {
			return team, nil
		}
	}
	return findEnterpriseTeamByID(meta, ctx, enterpriseSlug, id)
}

// organizationSlugs extracts the non-empty logins of the given organizations.
func organizationSlugs(orgs []*github.Organization) []string {
	slugs := make([]string, 0, len(orgs))
	for _, org := range orgs {
		if login := org.GetLogin(); login != "" {
			slugs = append(slugs, login)
		}
	}
	return slugs
}

// listAllEnterpriseTeamOrganizations returns all organizations assigned to an enterprise team with pagination handled.
func listAllEnterpriseTeamOrganizations(meta *Owner, ctx context.Context, enterpriseSlug, enterpriseTeam string) ([]*github.Organization, error) {
	var all []*github.Organization
	opt := &github.ListOptions{PerPage: meta.maxPerPage}

	for {
		orgs, resp, err := meta.v3client.Enterprise.ListAssignments(ctx, enterpriseSlug, enterpriseTeam, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, orgs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return all, nil
}

// listAllEnterpriseTeams returns all enterprise teams with pagination handled.
func listAllEnterpriseTeams(meta *Owner, ctx context.Context, enterpriseSlug string) ([]*github.EnterpriseTeam, error) {
	var all []*github.EnterpriseTeam
	opt := &github.ListOptions{PerPage: meta.maxPerPage}

	for {
		teams, resp, err := meta.v3client.Enterprise.ListTeams(ctx, enterpriseSlug, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, teams...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return all, nil
}
