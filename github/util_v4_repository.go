package github

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/shurcooL/githubv4"
)

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

// Maintain compatibility with deprecated Global ID format
// https://github.blog/2021-02-10-new-global-id-format-coming-to-graphql/
func repositoryLegacyNodeIDExists(name string, meta any) (bool, error) {
	// Check if the name is a base 64 encoded node ID
	_, err := base64.StdEncoding.DecodeString(name)
	if err != nil {
		return false, nil
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
	err = client.Query(ctx, &query, variables)
	if err != nil {
		return false, err
	}

	return query.Node.ID.(string) == name, nil
}
