package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

func getRepositoryID(name string, meta interface{}) (githubv4.ID, error) {
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
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return query.Repository.ID, nil
}

func getRepositoryName(id string, meta interface{}) (githubv4.String, error) {

	// {
	// 	node(id: "MDEwOlJlcG9zaXRvcnkzMDkxNjc3NjY=") {
	// 		... on Repository {
	// 			name
	// 		}
	// 	}
	// }

	var query struct {
		Node struct {
			Repository struct {
				Name githubv4.String
			} `graphql:"... on Repository"`
		} `graphql:"node(id:$id)"`
	}

	variables := map[string]interface{}{
		"id": githubv4.ID(id),
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return "", err
	}

	return query.Node.Repository.Name, nil
}

func getBranchProtectionID(name string, pattern string, meta interface{}) (githubv4.ID, error) {
	var query struct {
		Node struct {
			Repository struct {
				BranchProtectionRules struct {
					Nodes []struct {
						ID      string
						Pattern string
					}
					PageInfo PageInfo
				} `graphql:"branchProtectionRules(first: $first, after: $cursor)"`
				ID string
			} `graphql:"... on Repository"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"owner":  githubv4.String(meta.(*Owner).name),
		"name":   githubv4.String(name),
		"first":  githubv4.Int(100),
		"cursor": (*githubv4.String)(nil),
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client

	var allRules []struct {
		ID      string
		Pattern string
	}
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}

		allRules = append(allRules, query.Node.Repository.BranchProtectionRules.Nodes...)

		if !query.Node.Repository.BranchProtectionRules.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Node.Repository.BranchProtectionRules.PageInfo.EndCursor)
	}

	var id string
	for i := range allRules {
		if allRules[i].Pattern == pattern {
			id = allRules[i].ID
			break
		}
	}

	return id, nil
}
