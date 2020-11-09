package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func getRepositoryID(name string, meta interface{}) (githubv4.ID, error) {

	// Interperet `name` as a node ID and return
	exists, err := repositoryNodeIDExists(name, meta)
	if exists {
		return githubv4.ID(name), nil
	}
	if err != nil {
		return nil, err
	}

	// Resolve `name` to a node ID and return
	var query struct {
		Repository struct {
			ID githubv4.ID
		} `graphql:"repository(owner:$owner, name:$name)"`
	}
	variables := map[string]interface{}{
		"owner": githubv4.String(meta.(*Owner).name),
		"name":  githubv4.String(name),
	}
	ctx := context.Background()
	client := meta.(*Owner).v4client
	err = client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return query.Repository.ID, nil
}

func repositoryNodeIDExists(name string, meta interface{}) (bool, error) {
	// Quick check for node ID length
	if len(name) != 32 {
		return false, nil
	}

	// API check if node ID exists
	var query struct {
		Node struct {
			ID githubv4.ID
		} `graphql:"node(id:$id)"`
	}
	variables := map[string]interface{}{
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
