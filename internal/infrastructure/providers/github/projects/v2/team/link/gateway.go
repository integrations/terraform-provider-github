package link

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

type Gateway struct{ client *githubv4.Client }

func NewGateway(client *githubv4.Client) *Gateway { return &Gateway{client: client} }

func (gateway *Gateway) Resolve(ctx context.Context, organization, slug string) (application.Result, error) {
	var query struct {
		Organization struct {
			Team teamNode `graphql:"team(slug: $slug)"`
		} `graphql:"organization(login: $organization)"`
	}
	variables := map[string]any{"organization": githubv4.String(organization), "slug": githubv4.String(slug)}
	if err := gateway.client.Query(ctx, &query, variables); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying team %q/%q", organization, slug), err)
	}
	return resultFromNode("", query.Organization.Team), nil
}

func (gateway *Gateway) Attach(ctx context.Context, projectID, teamID string) (application.Result, error) {
	var mutation struct {
		LinkProjectV2ToTeam struct{ Team teamNode } `graphql:"linkProjectV2ToTeam(input: $input)"`
	}
	variables := githubv4.LinkProjectV2ToTeamInput{ProjectID: githubv4.ID(projectID), TeamID: githubv4.ID(teamID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("linking team %q to Projects V2 project %q", teamID, projectID), err)
	}
	if mutation.LinkProjectV2ToTeam.Team.ID == "" {
		return application.Result{}, fmt.Errorf("GitHub returned a linked team without an ID")
	}
	return resultFromNode(projectID, mutation.LinkProjectV2ToTeam.Team), nil
}

func (gateway *Gateway) Get(ctx context.Context, projectID, teamID string) (application.Result, error) {
	return gateway.find(ctx, projectID, teamID)
}

func (gateway *Gateway) Detach(ctx context.Context, projectID, teamID string) error {
	var mutation struct {
		UnlinkProjectV2FromTeam struct{ Team teamNode } `graphql:"unlinkProjectV2FromTeam(input: $input)"`
	}
	variables := githubv4.UnlinkProjectV2FromTeamInput{ProjectID: githubv4.ID(projectID), TeamID: githubv4.ID(teamID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("unlinking team %q from Projects V2 project %q", teamID, projectID), err)
	}
	return nil
}

func (gateway *Gateway) find(ctx context.Context, projectID, teamID string) (application.Result, error) {
	var after *githubv4.String
	for {
		var query struct {
			Node struct {
				Project struct {
					Teams struct {
						Nodes    []teamNode
						PageInfo pageInfo
					} `graphql:"teams(first: 100, after: $after)"`
				} `graphql:"... on ProjectV2"`
			} `graphql:"node(id: $id)"`
		}
		if err := gateway.client.Query(ctx, &query, map[string]any{"id": githubv4.ID(projectID), "after": after}); err != nil {
			return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying teams linked to Projects V2 project %q", projectID), err)
		}
		for _, team := range query.Node.Project.Teams.Nodes {
			if string(team.ID) == teamID {
				return resultFromNode(projectID, team), nil
			}
		}
		if !bool(query.Node.Project.Teams.PageInfo.HasNextPage) {
			return application.Result{}, fmt.Errorf("team %q link to project %q: %w", teamID, projectID, projects.ErrNotFound)
		}
		after = new(query.Node.Project.Teams.PageInfo.EndCursor)
	}
}

var _ application.Store = (*Gateway)(nil)
