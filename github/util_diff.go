package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v84/github"
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

// diffTeam compares the team_id and team_slug fields to determine if the team has changed.
func diffTeam(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	// Skip for new resources - no existing team_id to compare against
	if len(diff.Id()) == 0 {
		return nil
	}

	ctx = tflog.SetField(ctx, "id", diff.Id())

	if diff.HasChange("team_slug") {
		if isNewTeamID(ctx, diff, m) {
			return diff.ForceNew("team_slug")
		}
	}

	return nil
}

// helper function to determine if the team has changed or was renamed.
func isNewTeamID(ctx context.Context, diff *schema.ResourceDiff, m any) bool {
	// Get old team_id from state
	oldTeamID := toInt64(diff.Get("team_id"))
	if oldTeamID == 0 {
		return false
	}
	meta := m.(*Owner)

	// Resolve new team_slug to team ID via API
	oldTeamSlug, newTeamSlug := diff.GetChange("team_slug")
	newTeamID, err := lookupTeamID(ctx, meta, newTeamSlug.(string))
	if err != nil {
		// If team doesn't exist or API fails, skip ForceNew check and let Read handle it
		tflog.Debug(ctx, "Unable to resolve new team_slug to team ID, skipping ForceNew check", map[string]any{
			"new_team_slug": newTeamSlug,
			"error":         err.Error(),
		})
		return false
	}

	if newTeamID != oldTeamID {
		tflog.Debug(ctx, "Team ID changed, forcing new resource", map[string]any{
			"old_team_id":   oldTeamID,
			"new_team_id":   newTeamID,
			"new_team_slug": newTeamSlug,
			"old_team_slug": oldTeamSlug,
		})
		return true
	}

	return false
}
