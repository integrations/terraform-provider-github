package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v83/github"
)

// teamIdentity represents a GitHub team.
type teamIdentity struct {
	id     *int64
	slug   *string
	teamID *string
}

// newLegacyTeamIdentity creates a new teamIdentity from the given teamID string representing either the team ID or slug.
func newLegacyTeamIdentity(teamID string) teamIdentity {
	t := teamIdentity{teamID: &teamID}

	if id, ok := parseTeamID(teamID); ok {
		t.id = &id
	} else {
		t.slug = &teamID
	}

	return t
}

// getID returns the team ID as an int64 if it's available, or 0 if it's not.
func (t teamIdentity) getID() int64 {
	id, _ := t.getIDOK()
	return id
}

// getIDOK returns the team ID as an int64 if it's available.
func (t teamIdentity) getIDOK() (int64, bool) {
	if t.id != nil {
		return *t.id, true
	}
	return 0, false
}

// getSlug returns the slug of the team if it's available, or an empty string if it's not.
func (t teamIdentity) getSlug() string {
	slug, _ := t.getSlugOK()
	return slug
}

// getSlugOK returns the slug of the team if it's available.
func (t teamIdentity) getSlugOK() (string, bool) {
	if t.slug != nil {
		return *t.slug, true
	}
	return "", false
}

// getTeamID returns the legacy team ID of the team if it's available, or an empty string if it's not.
func (t teamIdentity) getTeamID() string {
	teamID, _ := t.getTeamIDOK()
	return teamID
}

// getTeamIDOK returns the legacy team ID of the team if it's available.
func (t teamIdentity) getTeamIDOK() (string, bool) {
	if t.teamID != nil {
		return *t.teamID, true
	}
	return "", false
}

// teamCollaborator represents a GitHub team collaborator with its identity and permission level.
type teamCollaborator struct {
	teamIdentity
	permission string
}

// flatten converts the teamCollaborator into a format suitable for Terraform schema.
func (t teamCollaborator) flatten() any {
	if teamID, ok := t.getTeamIDOK(); ok {
		return map[string]any{
			"team_id":    teamID,
			"permission": t.permission,
		}
	}

	return map[string]any{
		"team_id":    t.getSlug(),
		"permission": t.permission,
	}
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
	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		return id, true
	}
	return 0, false
}

// getTeamSlug returns the slug of the team identified by the given string, which may be either a team ID or a team slug.
func getTeamSlug(ctx context.Context, meta *Owner, s string) (string, error) {
	if id, ok := parseTeamID(s); ok {
		// The given id is an integer, assume it is the team ID and look up the corresponding team slug.
		return lookupTeamSlug(ctx, meta.v3client, meta.id, id)
	}
	// The given id not an integer, assume it is a team slug.
	return s, nil
}

// getTeamID returns the id of the team identified by the given string, which may be either a team ID or a team slug.
func getTeamID(ctx context.Context, meta *Owner, s string) (int64, error) {
	if id, ok := parseTeamID(s); ok {
		// The given id is an integer, assume it is the team ID.
		return id, nil
	}
	// The given id not an integer, assume it is a team slug and look up the corresponding team ID.
	return lookupTeamID(ctx, meta.v3client, meta.name, s)
}

// lookupTeamSlug looks up the slug of a team by its ID.
func lookupTeamSlug(ctx context.Context, client *github.Client, orgID, id int64) (string, error) {
	team, _, err := client.Teams.GetTeamByID(ctx, orgID, id) //nolint:staticcheck
	if err != nil {
		return "", err
	}
	return team.GetSlug(), nil
}

// lookupTeamID looks up the ID of a team by its slug.
func lookupTeamID(ctx context.Context, client *github.Client, orgName, slug string) (int64, error) {
	team, _, err := client.Teams.GetTeamBySlug(ctx, orgName, slug)
	if err != nil {
		return 0, err
	}
	return team.GetID(), nil
}
