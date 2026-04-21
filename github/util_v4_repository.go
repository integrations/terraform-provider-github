package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/shurcooL/githubv4"
)

// PullRequestCreationPolicy mirrors the GitHub GraphQL enum type of the same
// name so we can query and mutate the field even when the vendored client
// model lags behind the live schema.
type PullRequestCreationPolicy string

const (
	PullRequestCreationPolicyAll               PullRequestCreationPolicy = "ALL"
	PullRequestCreationPolicyCollaboratorsOnly PullRequestCreationPolicy = "COLLABORATORS_ONLY"
)

// UpdateRepositoryInput intentionally mirrors the GitHub GraphQL input type
// name so the graphql client emits the correct variable type in mutations.
// We only model the fields needed for pullRequestCreationPolicy updates.
type UpdateRepositoryInput struct {
	RepositoryID              githubv4.ID                `json:"repositoryId"`
	PullRequestCreationPolicy *PullRequestCreationPolicy `json:"pullRequestCreationPolicy,omitempty"`
	ClientMutationID          *githubv4.String           `json:"clientMutationId,omitempty"`
}

func getRepositoryID(name string, meta any) (githubv4.ID, error) {
	// Interpret `name` as a node ID
	exists, nodeIDerr := repositoryNodeIDExists(name, meta)
	if exists {
		return githubv4.ID(name), nil
	}

	// Interpret `name` as a legacy node ID
	exists, _ = repositoryLegacyNodeIDExists(name, meta)
	if exists {
		return githubv4.ID(name), nil
	}

	// Could not find repo by node ID, interpret `name` as repo name
	var query struct {
		Repository struct {
			ID githubv4.ID
		} `graphql:"repository(owner:$owner, name:$name)"`
	}
	variables := map[string]any{
		"owner": githubv4.String(meta.(*Owner).name),
		"name":  githubv4.String(name),
	}
	ctx := context.Background()
	client := meta.(*Owner).v4client
	nameErr := client.Query(ctx, &query, variables)
	if nameErr != nil {
		if nodeIDerr != nil {
			// Could not find repo by node ID or repo name, return both errors
			return nil, errors.New(nodeIDerr.Error() + nameErr.Error())
		}
		return nil, nameErr
	}

	return query.Repository.ID, nil
}

func repositoryNodeIDExists(name string, meta any) (bool, error) {
	// API check if node ID exists
	var query struct {
		Node struct {
			ID githubv4.ID
		} `graphql:"node(id:$id)"`
	}
	variables := map[string]any{
		"id": githubv4.ID(name),
	}
	ctx := context.Background()
	client := meta.(*Owner).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return false, err
	}

	return query.Node.ID.(string) == name, nil
}

func flattenPullRequestCreationPolicy(policy PullRequestCreationPolicy) (string, error) {
	switch policy {
	case PullRequestCreationPolicyAll:
		return "all", nil
	case PullRequestCreationPolicyCollaboratorsOnly:
		return "collaborators_only", nil
	case "":
		return "", nil
	default:
		return "", fmt.Errorf("unsupported GraphQL pull request creation policy %q", policy)
	}
}

func expandPullRequestCreationPolicy(policy string) (PullRequestCreationPolicy, error) {
	switch policy {
	case "all":
		return PullRequestCreationPolicyAll, nil
	case "collaborators_only":
		return PullRequestCreationPolicyCollaboratorsOnly, nil
	default:
		return "", fmt.Errorf("unsupported Terraform pull request creation policy %q", policy)
	}
}

func getRepositoryPullRequestCreationPolicy(ctx context.Context, owner, name string, meta any) (string, error) {
	var query struct {
		Repository struct {
			PullRequestCreationPolicy PullRequestCreationPolicy
		} `graphql:"repository(owner:$owner, name:$name)"`
	}

	variables := map[string]any{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	client := meta.(*Owner).v4client
	if err := client.Query(ctx, &query, variables); err != nil {
		return "", err
	}

	return flattenPullRequestCreationPolicy(query.Repository.PullRequestCreationPolicy)
}

func updateRepositoryPullRequestCreationPolicy(ctx context.Context, repositoryID githubv4.ID, policy string, meta any) error {
	expandedPolicy, err := expandPullRequestCreationPolicy(policy)
	if err != nil {
		return err
	}

	input := UpdateRepositoryInput{
		RepositoryID:              repositoryID,
		PullRequestCreationPolicy: &expandedPolicy,
	}

	var mutation struct {
		UpdateRepository struct {
			Repository struct {
				ID githubv4.ID
			}
		} `graphql:"updateRepository(input:$input)"`
	}

	client := meta.(*Owner).v4client
	return client.Mutate(ctx, &mutation, input, nil)
}

// Maintain compatibility with deprecated Global ID format
// https://github.blog/2021-02-10-new-global-id-format-coming-to-graphql/
func repositoryLegacyNodeIDExists(name string, meta any) (bool, error) {
	// Check if the name is a base 64 encoded node ID
	if _, err := base64.StdEncoding.DecodeString(name); err != nil {
		var corrErr base64.CorruptInputError
		ok := errors.As(err, &corrErr)
		if ok {
			return false, nil
		}

		return false, err
	}

	// API check if node ID exists
	var query struct {
		Node struct {
			ID githubv4.ID
		} `graphql:"node(id:$id)"`
	}
	variables := map[string]any{
		"id": githubv4.ID(name),
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client
	if err := client.Query(ctx, &query, variables); err != nil {
		return false, err
	}

	return query.Node.ID.(string) == name, nil
}
