package github

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/shurcooL/githubv4"
)

func getRepositoryID(name string, m any) (githubv4.ID, error) {
	meta, _ := m.(*Owner)

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
		"owner": githubv4.String(meta.name),
		"name":  githubv4.String(name),
	}
	ctx := context.Background()
	client := meta.v4client
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

func repositoryNodeIDExists(name string, m any) (bool, error) {
	meta, _ := m.(*Owner)

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
	client := meta.v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return false, err
	}

	id, ok := query.Node.ID.(string)
	return ok && id == name, nil
}

// Maintain compatibility with deprecated Global ID format
// https://github.blog/2021-02-10-new-global-id-format-coming-to-graphql/
func repositoryLegacyNodeIDExists(name string, m any) (bool, error) {
	meta, _ := m.(*Owner)

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
	client := meta.v4client
	if err := client.Query(ctx, &query, variables); err != nil {
		return false, err
	}

	id, ok := query.Node.ID.(string)
	return ok && id == name, nil
}
