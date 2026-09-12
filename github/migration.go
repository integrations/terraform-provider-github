package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// migrateResourceID is a helper function to migrate a raw state where the resource ID contains a legacy separator.
func migrateResourceID(ctx context.Context, currentIDSeparator string, rawState map[string]any) (map[string]any, error) {
	tflog.Debug(ctx, "migrating state to replace legacy separator in resource ID", map[string]any{"currentIDSeparator": currentIDSeparator})

	if currentIDSeparator == idSeparator {
		tflog.Debug(ctx, "provided currentIDSeparator is the same as the new separator.", map[string]any{"currentIDSeparator": currentIDSeparator})
		return nil, fmt.Errorf("currentIDSeparator cannot be the same as the new separator")
	}
	if len(currentIDSeparator) == 0 {
		tflog.Debug(ctx, "provided currentIDSeparator is empty.", map[string]any{"currentIDSeparator": currentIDSeparator})
		return nil, fmt.Errorf("currentIDSeparator cannot be empty")
	}

	currentID, ok := rawState["id"]
	if !ok {
		return nil, fmt.Errorf("resource ID not found in state")
	}

	currentIDStr, ok := currentID.(string)
	if !ok {
		return nil, fmt.Errorf("resource ID is not a string")
	}

	if strings.Contains(currentIDStr, idSeparator) {
		tflog.Debug(ctx, "resource ID already contains the new separator, replacing with escaped ID separator", map[string]any{"currentID": currentIDStr})
		currentIDStr = strings.ReplaceAll(currentIDStr, idSeparator, idSeparatorEscaped)
	}

	newID := strings.ReplaceAll(currentIDStr, currentIDSeparator, idSeparator)

	rawState["id"] = newID
	return rawState, nil
}

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
