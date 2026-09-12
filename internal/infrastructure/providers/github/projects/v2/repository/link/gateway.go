package link

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

type Gateway struct{ client *githubv4.Client }

func NewGateway(client *githubv4.Client) *Gateway { return &Gateway{client: client} }

func (gateway *Gateway) Resolve(ctx context.Context, owner, name string) (application.Result, error) {
	var query struct {
		Repository repositoryNode `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]any{"owner": githubv4.String(owner), "name": githubv4.String(name)}
	if err := gateway.client.Query(ctx, &query, variables); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying repository %q/%q", owner, name), err)
	}
	return resultFromNode("", query.Repository), nil
}

func (gateway *Gateway) Attach(ctx context.Context, projectID, repositoryID string) (application.Result, error) {
	var mutation struct {
		LinkProjectV2ToRepository struct{ Repository repositoryNode } `graphql:"linkProjectV2ToRepository(input: $input)"`
	}
	variables := githubv4.LinkProjectV2ToRepositoryInput{ProjectID: githubv4.ID(projectID), RepositoryID: githubv4.ID(repositoryID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("linking repository %q to Projects V2 project %q", repositoryID, projectID), err)
	}
	if mutation.LinkProjectV2ToRepository.Repository.ID == "" {
		return application.Result{}, fmt.Errorf("GitHub returned a linked repository without an ID")
	}
	return resultFromNode(projectID, mutation.LinkProjectV2ToRepository.Repository), nil
}

func (gateway *Gateway) Get(ctx context.Context, projectID, repositoryID string) (application.Result, error) {
	return gateway.find(ctx, projectID, repositoryID)
}

func (gateway *Gateway) Detach(ctx context.Context, projectID, repositoryID string) error {
	var mutation struct {
		UnlinkProjectV2FromRepository struct{ Repository repositoryNode } `graphql:"unlinkProjectV2FromRepository(input: $input)"`
	}
	variables := githubv4.UnlinkProjectV2FromRepositoryInput{ProjectID: githubv4.ID(projectID), RepositoryID: githubv4.ID(repositoryID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("unlinking repository %q from Projects V2 project %q", repositoryID, projectID), err)
	}
	return nil
}

func (gateway *Gateway) find(ctx context.Context, projectID, repositoryID string) (application.Result, error) {
	var after *githubv4.String
	for {
		var query struct {
			Node struct {
				Project struct {
					Repositories struct {
						Nodes    []repositoryNode
						PageInfo pageInfo
					} `graphql:"repositories(first: 100, after: $after)"`
				} `graphql:"... on ProjectV2"`
			} `graphql:"node(id: $id)"`
		}
		if err := gateway.client.Query(ctx, &query, map[string]any{"id": githubv4.ID(projectID), "after": after}); err != nil {
			return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying repositories linked to Projects V2 project %q", projectID), err)
		}
		for _, repository := range query.Node.Project.Repositories.Nodes {
			if string(repository.ID) == repositoryID {
				return resultFromNode(projectID, repository), nil
			}
		}
		if !bool(query.Node.Project.Repositories.PageInfo.HasNextPage) {
			return application.Result{}, fmt.Errorf("repository %q link to project %q: %w", repositoryID, projectID, projects.ErrNotFound)
		}
		after = new(query.Node.Project.Repositories.PageInfo.EndCursor)
	}
}

var _ application.Store = (*Gateway)(nil)
