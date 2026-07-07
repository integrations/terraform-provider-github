package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// diffRepository checks if the repository has changed and forces a new resource if the repository ID does not match.
// The resource must have both "repository" and "repository_id" attributes.
func diffRepository(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	if len(diff.Id()) == 0 {
		return nil
	}

	ctx = tflog.SetField(ctx, "id", diff.Id())

	if diff.HasChange("repository") {
		meta := m.(*Owner)
		client := meta.v3client
		owner := meta.name

		var repoName string
		old, n := diff.GetChange("repository")
		if v, ok := n.(string); ok {
			repoName = v
		} else {
			return fmt.Errorf("repository is not a string")
		}

		var repoID int
		if o, ok := diff.GetOk("repository_id"); ok {
			if v, ok := o.(int); ok {
				repoID = v
			} else {
				return fmt.Errorf("repository_id is not an int")
			}
		} else {
			tflog.Info(ctx, "No repository_id in schema, cannot verify if repository change is a rename or a new repository. Forcing new resource.", map[string]any{
				"old_repository": old,
				"new_repository": repoName,
			})
			return diff.ForceNew("repository")
		}

		repo, _, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) {
				if ghErr.Response.StatusCode != http.StatusNotFound {
					return ghErr
				}

				tflog.Info(ctx, "Repository not found, assuming it was deleted and will be recreated. Forcing new resource.", map[string]any{"repository": repoName})
			} else {
				return err
			}
		} else {
			tflog.Debug(ctx, "Repository found when checking repository change.", map[string]any{"repository": repoName})

			if repoID != int(repo.GetID()) {
				tflog.Info(ctx, "Repository ID changed, forcing new resource.", map[string]any{
					"old_repository":    old,
					"old_repository_id": repoID,
					"new_repository":    repoName,
					"new_repository_id": repo.GetID(),
				})
				return diff.ForceNew("repository")
			}
		}
	}

	return nil
}

// diffSecret compares the remote_updated_at and updated_at fields to determine if the secret has changed remotely.
func diffSecret(ctx context.Context, diff *schema.ResourceDiff, _ any) error {
	if len(diff.Id()) == 0 {
		return nil
	}

	if diff.HasChange("remote_updated_at") {
		remoteUpdatedAt := diff.Get("remote_updated_at").(string)
		if len(remoteUpdatedAt) == 0 {
			return nil
		}

		updatedAt := diff.Get("updated_at").(string)
		if updatedAt != remoteUpdatedAt {
			if len(updatedAt) == 0 {
				return diff.SetNew("updated_at", remoteUpdatedAt)
			}

			return diff.SetNewComputed("updated_at")
		}
	}

	return nil
}

// diffSecretVariableVisibility ensures that selected_repository_ids is only set when visibility is set to selected.
func diffSecretVariableVisibility(ctx context.Context, d *schema.ResourceDiff, _ any) error {
	if len(d.Id()) == 0 {
		return nil
	}

	visibility := d.Get("visibility").(string)
	if visibility != "selected" {
		if _, ok := d.GetOk("selected_repository_ids"); ok {
			return fmt.Errorf("cannot use selected_repository_ids without visibility being set to selected")
		}
	}

	return nil
}

// diffLegacyTeamID checks for an unset team_id previously storing a team slug and updates this to carry the numeric ID.
func diffLegacyTeamID(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	// Skip for new resources - no existing team_id to compare against
	if len(diff.Id()) == 0 || diff.GetRawConfig().IsNull() {
		return nil
	}

	ctx = tflog.SetField(ctx, "id", diff.Id())

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	if diff.GetRawConfig().AsValueMap()["team_id"].IsNull() {
		teamIDStr, _ := diff.Get("team_id").(string)
		if _, err := strconv.ParseInt(teamIDStr, 10, 64); err != nil {
			tflog.Debug(ctx, "Computing team_id from old slug.", map[string]any{"team_slug": teamIDStr})

			team, _, err := client.Teams.GetTeamBySlug(ctx, owner, teamIDStr)
			if err != nil {
				if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusNotFound {
					tflog.Debug(ctx, "Team not found so set ID as computed.", map[string]any{"team_slug": teamIDStr})

					if err := diff.SetNewComputed("team_id"); err != nil {
						return fmt.Errorf("failed to set new computed for team_id: %w", err)
					}

					return nil
				}

				return fmt.Errorf("failed to lookup team id from slug (%s): %w", teamIDStr, err)
			}

			if err := diff.SetNew("team_id", strconv.FormatInt(team.GetID(), 10)); err != nil {
				return fmt.Errorf("failed to set new for team_id: %w", err)
			}
		}
	}

	if diff.HasChange("team_id") {
		oldTeamIDV, newTeamIDV := diff.GetChange("team_id")
		oldTeamIDStr, _ := oldTeamIDV.(string)
		newTeamIDStr, _ := newTeamIDV.(string)

		oldTeam, err := getTeam(ctx, meta, oldTeamIDStr)
		if err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusNotFound {
				tflog.Debug(ctx, "Old team not found for legacy team_id so stop calculating diff.", map[string]any{"team_id": oldTeamIDStr})

				return nil
			}

			return fmt.Errorf("failed to lookup old team (%s): %w", oldTeamIDStr, err)
		}

		newTeam, err := getTeam(ctx, meta, newTeamIDStr)
		if err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusNotFound {
				tflog.Debug(ctx, "New team not found for legacy team_id so stop calculating diff.", map[string]any{"team_id": newTeamIDStr})

				return nil
			}

			return fmt.Errorf("failed to lookup new team (%s): %w", newTeamIDStr, err)
		}

		if oldTeam.GetID() != newTeam.GetID() {
			tflog.Debug(ctx, "Team ID changed, forcing new resource", map[string]any{"old_team_id": oldTeam.GetID(), "new_team_id": newTeam.GetID()})
			if err := diff.ForceNew("team_id"); err != nil {
				return fmt.Errorf("failed to force new for team_id: %w", err)
			}
		}
	}

	return nil
}

// diffLegacyTeam compares the legacy team_id string and team_slug fields to determine if the team has changed.
func diffLegacyTeam(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	// Skip for new resources - no existing team_id to compare against
	if len(diff.Id()) == 0 {
		return nil
	}

	if !diff.HasChanges("team_slug") {
		return nil
	}

	ctx = tflog.SetField(ctx, "id", diff.Id())

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	teamIDV := diff.Get("team_id")
	teamIDStr, _ := teamIDV.(string)
	teamID, _ := strconv.ParseInt(teamIDStr, 10, 64)

	if teamID == 0 {
		tflog.Debug(ctx, "No team ID set so skipping team diff.", map[string]any{"team_slug": diff.Get("team_slug"), "team_id": teamIDStr})
		return nil
	}

	slug, _ := diff.Get("team_slug").(string)

	changed, err := diffTeamCheck(ctx, client, owner, teamID, slug)
	if err != nil {
		return err
	}

	if changed {
		if err := diff.ForceNew("team_slug"); err != nil {
			return fmt.Errorf("failed to force new for team_slug: %w", err)
		}
	}

	return nil
}

// diffTeam compares the team_id and team_slug fields to determine if the team has changed.
func diffTeam(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	// Skip for new resources - no existing team_id to compare against
	if len(diff.Id()) == 0 {
		return nil
	}

	if !diff.HasChanges("team_slug") {
		return nil
	}

	ctx = tflog.SetField(ctx, "id", diff.Id())

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	teamID := toInt64(diff.Get("team_id"))
	slug, _ := diff.Get("team_slug").(string)

	changed, err := diffTeamCheck(ctx, client, owner, teamID, slug)
	if err != nil {
		return err
	}

	if changed {
		if err := diff.ForceNew("team_slug"); err != nil {
			return fmt.Errorf("failed to force new for team_slug: %w", err)
		}
	}

	return nil
}

func diffTeamCheck(ctx context.Context, client *github.Client, owner string, teamID int64, slug string) (bool, error) {
	team, _, err := client.Teams.GetTeamBySlug(ctx, owner, slug)
	if err != nil {
		if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusNotFound {
			tflog.Debug(ctx, "Team not found when checking team change, skipping diff.", map[string]any{"team_slug": slug})

			return false, nil
		}

		return false, fmt.Errorf("failed to lookup team from slug: %w", err)
	}

	if team.GetID() != teamID {
		tflog.Debug(ctx, "Team ID changed, forcing new resource.", map[string]any{"old_team_id": teamID, "new_team_slug": slug, "new_team_id": team.GetID()})

		return true, nil
	}

	return false, nil
}
