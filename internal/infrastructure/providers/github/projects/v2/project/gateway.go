package project

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

type Gateway struct{ client *githubv4.Client }

func NewGateway(client *githubv4.Client) *Gateway { return &Gateway{client: client} }

func (gateway *Gateway) Create(ctx context.Context, input application.CreateInput) (application.Result, error) {
	ownerID, err := gateway.ownerID(ctx, input.OwnerKind, input.Owner)
	if err != nil {
		return application.Result{}, err
	}
	var mutation struct {
		CreateProjectV2 struct {
			Project node `graphql:"projectV2"`
		} `graphql:"createProjectV2(input: $input)"`
	}
	create := githubv4.CreateProjectV2Input{OwnerID: ownerID, Title: githubv4.String(input.Title)}
	if err := gateway.client.Mutate(ctx, &mutation, create, nil); err != nil {
		return application.Result{}, projectgraphql.Error("creating Projects V2 project", err)
	}
	return resultFromNode(mutation.CreateProjectV2.Project)
}

func (gateway *Gateway) Get(ctx context.Context, id string) (application.Result, error) {
	var query struct {
		Node struct {
			Project node `graphql:"... on ProjectV2"`
		} `graphql:"node(id: $id)"`
	}
	if err := gateway.client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)}); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying Projects V2 project %q", id), err)
	}
	return resultFromNode(query.Node.Project)
}

func (gateway *Gateway) Update(ctx context.Context, input application.UpdateInput) (application.Result, error) {
	title, description, readme := githubv4.String(input.Title), githubv4.String(input.ShortDescription), githubv4.String(input.Readme)
	public, closed := githubv4.Boolean(input.Public), githubv4.Boolean(input.Closed)
	variables := githubv4.UpdateProjectV2Input{ProjectID: githubv4.ID(input.ID), Title: &title, ShortDescription: &description, Readme: &readme, Public: &public, Closed: &closed}
	var mutation struct {
		UpdateProjectV2 struct {
			Project node `graphql:"projectV2"`
		} `graphql:"updateProjectV2(input: $input)"`
	}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("updating Projects V2 project %q", input.ID), err)
	}
	return resultFromNode(mutation.UpdateProjectV2.Project)
}

func (gateway *Gateway) Delete(ctx context.Context, id string) error {
	var mutation struct {
		DeleteProjectV2 struct{ ClientMutationID githubv4.String } `graphql:"deleteProjectV2(input: $input)"`
	}
	if err := gateway.client.Mutate(ctx, &mutation, githubv4.DeleteProjectV2Input{ProjectID: githubv4.ID(id)}, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("deleting Projects V2 project %q", id), err)
	}
	return nil
}

func (gateway *Gateway) ownerID(ctx context.Context, kind application.OwnerKind, login string) (githubv4.ID, error) {
	variables := map[string]any{"login": githubv4.String(login)}
	switch kind {
	case application.OwnerOrganization:
		var query struct {
			Organization struct{ ID githubv4.String } `graphql:"organization(login: $login)"`
		}
		if err := gateway.client.Query(ctx, &query, variables); err != nil {
			return nil, projectgraphql.Error(fmt.Sprintf("querying organization %q", login), err)
		}
		return githubv4.ID(query.Organization.ID), nil
	case application.OwnerUser:
		var query struct {
			User struct{ ID githubv4.String } `graphql:"user(login: $login)"`
		}
		if err := gateway.client.Query(ctx, &query, variables); err != nil {
			return nil, projectgraphql.Error(fmt.Sprintf("querying user %q", login), err)
		}
		return githubv4.ID(query.User.ID), nil
	default:
		return nil, fmt.Errorf("unsupported Projects V2 owner kind %q", kind)
	}
}

var _ application.Store = (*Gateway)(nil)
