package github

import (
	"context"
	"fmt"
	"strconv"
)

// teamWithSlug is an interface representing a GitHub team that has a slug.
type teamWithSlug interface {
	getSlug() string
}

// teamIdentity represents a GitHub team by its slug.
// It also can optionally include the team ID as a string for legacy support.
type teamIdentity struct {
	slug   string
	teamID *string
}

// getSlug returns the slug of the team.
func (t teamIdentity) getSlug() string {
	return t.slug
}

// teamCollaborator represents a GitHub team collaborator with its identity and permission level.
type teamCollaborator struct {
	teamIdentity
	permission string
}

// flatten converts the teamCollaborator into a format suitable for Terraform schema.
func (t teamCollaborator) flatten() any {
	m := map[string]any{
		"slug":       t.slug,
		"permission": t.permission,
	}
	if t.teamID != nil {
		m["team_id"] = *t.teamID
	}
	return m
}

// teamCollaborators is a slice of teamCollaborator.
type teamCollaborators []teamCollaborator

// flatten converts the teamCollaborators slice into a format suitable for Terraform schema.
func (tc teamCollaborators) flatten() any {
	items := make([]any, len(tc))

	for i, t := range tc {
		items[i] = t.flatten()
	}

	return items
}

// parseTeamID attempts to parse the given string as a team ID (int64).
// It returns the parsed ID and a boolean indicating whether the parsing was successful.
func parseTeamID(s string) (int64, bool) {
	id, err := strconv.ParseInt(s, 10, 64)
	return id, err == nil
}

// getTeamSlug returns the slug of the team identified by the given string, which may be either a team ID or a team slug.
func getTeamSlug(ctx context.Context, meta *Owner, s string) (string, error) {
	id, ok := parseTeamID(s)
	if !ok {
		// The given id not an integer, assume it is a team slug
		return s, nil
	}

	return lookupTeamSlug(ctx, meta, id)
}

// lookupTeamSlug looks up the slug of a team by its ID.
func lookupTeamSlug(ctx context.Context, meta *Owner, id int64) (string, error) {
	client := meta.v3client
	orgId := meta.id

	team, _, err := client.Teams.GetTeamByID(ctx, orgId, id) //nolint:staticcheck
	if err != nil {
		return "", err
	}
	return team.GetSlug(), nil
}

// lookupTeamID looks up the ID of a team by its slug.
func lookupTeamID(ctx context.Context, meta *Owner, slug string) (int64, error) {
	client := meta.v3client
	owner := meta.name

	team, _, err := client.Teams.GetTeamBySlug(ctx, owner, slug)
	if err != nil {
		return 0, err
	}
	return team.GetID(), nil
}

// getTeamIdentity returns a team identity represented by the input.
// The input may include a team slug or a team ID.
// The output will include the team slug.
func getTeamIdentity(ctx context.Context, meta *Owner, d any) (teamIdentity, error) {
	m, ok := d.(map[string]any)
	if !ok {
		return teamIdentity{}, fmt.Errorf("team input invalid")
	}

	if s, ok := m["slug"]; ok {
		if slug, ok := s.(string); ok && len(slug) > 0 {
			return teamIdentity{slug: slug}, nil
		}
	}

	if d, ok := m["team_id"]; ok {
		if teamID, ok := d.(string); ok && len(teamID) > 0 {
			slug, err := getTeamSlug(ctx, meta, teamID)
			if err != nil {
				return teamIdentity{}, err
			}

			return teamIdentity{slug: slug, teamID: &teamID}, nil
		}
	}

	return teamIdentity{}, fmt.Errorf("team input must include either 'slug' or 'team_id'")
}

// getTeamIdentity returns a team identity represented by the input.
// The input must include a team slug.
func getTeamIdentityStrict(d any) (teamIdentity, error) {
	m, ok := d.(map[string]any)
	if !ok {
		return teamIdentity{}, fmt.Errorf("team input invalid")
	}

	if s, ok := m["slug"]; ok {
		if slug, ok := s.(string); ok && len(slug) > 0 {
			return teamIdentity{slug: slug}, nil
		}
	}

	return teamIdentity{}, fmt.Errorf("input must include 'slug'")
}

// getTeamIdentities returns a list of team identities represented by the input.
// Each input may include a team slug or a team ID.
// Each team identity in the output will include the slug.
func getTeamIdentities(ctx context.Context, meta *Owner, col []any) ([]teamIdentity, error) {
	identities := make([]teamIdentity, len(col))

	for i, t := range col {
		id, err := getTeamIdentity(ctx, meta, t)
		if err != nil {
			return nil, err
		}
		identities[i] = id
	}

	return identities, nil
}

// getTeamCollaborators returns a list of team collaborators represented by the input.
// Each input may include a team slug or a team ID, along with a permission level.
// Each team collaborator in the output will include the slug.
func getTeamCollaborators(ctx context.Context, meta *Owner, col []any) (teamCollaborators, error) {
	collaborators := make([]teamCollaborator, len(col))

	for i, t := range col {
		m, ok := t.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("input invalid")
		}

		id, err := getTeamIdentity(ctx, meta, m)
		if err != nil {
			return nil, err
		}

		permission, ok := m["permission"].(string)
		if !ok || len(permission) == 0 {
			return nil, fmt.Errorf("team input must include 'permission'")
		}

		collaborators[i] = teamCollaborator{
			teamIdentity: id,
			permission:   permission,
		}
	}

	return collaborators, nil
}

// checkDuplicateTeams checks for duplicate team slugs in the given list of team identities.
func checkDuplicateTeams[T teamWithSlug](teams []T) bool {
	seen := make(map[string]any)

	for _, team := range teams {
		slug := team.getSlug()
		if _, ok := seen[slug]; ok {
			return true
		}
		seen[slug] = nil
	}

	return false
}
