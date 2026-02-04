package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// diffRepository checks if the repository has changed and forces a new resource if the repository ID does not match.
// The resource must have both "repository" and "repository_id" attributes.
func diffRepository(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	if len(diff.Id()) == 0 {
		return nil
	}

	if diff.HasChange("repository") {
		meta := m.(*Owner)
		client := meta.v3client
		owner := meta.name

		var repoID int
		if o, ok := diff.GetOk("repository_id"); ok {
			if v, ok := o.(int); ok {
				repoID = v
			} else {
				return fmt.Errorf("repository_id is not an int")
			}
		} else {
			return fmt.Errorf("repository_id is not set")
		}

		var repoName string
		if o, ok := diff.GetOk("repository"); ok {
			if v, ok := o.(string); ok {
				repoName = v
			} else {
				return fmt.Errorf("repository is not a string")
			}
		} else {
			return fmt.Errorf("repository is not set")
		}

		repo, _, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) {
				if ghErr.Response.StatusCode != http.StatusNotFound {
					return ghErr
				}

				log.Printf("[INFO] Repository %s not found when checking repository change diff %s", repoName, diff.Id())
			} else {
				return err
			}
		} else {
			log.Printf("[INFO] Repository %s found when checking repository change diff %s", repoName, diff.Id())

			if repoID != int(repo.GetID()) {
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
