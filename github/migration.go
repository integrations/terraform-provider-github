package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// migrateRepositoryWithID is a helper function to migrate a raw state where the repository is set to make usre that the repository_id is added to the state.
func migrateRepositoryWithID(ctx context.Context, client *github.Client, owner string, rawState map[string]any) (map[string]any, error) {
	tflog.Debug(ctx, "Migrating state to add repository_id.")

	repoNameVal, ok := rawState["repository"]
	if !ok {
		return nil, fmt.Errorf("repository name not found in state")
	}

	repoName, ok := repoNameVal.(string)
	if !ok {
		return nil, fmt.Errorf("repository name is not a string")
	}

	if repoID, ok := rawState["repository_id"]; ok {
		if _, ok := repoID.(int); ok {
			tflog.Debug(ctx, "Found repository_id in state, skipping migration.", map[string]any{"repository": repoName, "repository_id": repoID})
			return rawState, nil
		}
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	repoID := int(repo.GetID())
	rawState["repository_id"] = repoID

	tflog.Debug(ctx, "State migrated to add repository_id.", map[string]any{"repository": repoName, "repository_id": repoID})

	return rawState, nil
}
