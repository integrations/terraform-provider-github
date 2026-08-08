package item

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

type Gateway struct{ client *githubv4.Client }

func NewGateway(client *githubv4.Client) *Gateway { return &Gateway{client: client} }

func (gateway *Gateway) Add(ctx context.Context, projectID, contentID string) (application.Result, error) {
	var mutation struct {
		AddProjectV2ItemByID struct{ Item node } `graphql:"addProjectV2ItemById(input: $input)"`
	}
	variables := githubv4.AddProjectV2ItemByIdInput{ProjectID: githubv4.ID(projectID), ContentID: githubv4.ID(contentID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("adding content to Projects V2 project %q", projectID), err)
	}
	return resultFromNode(mutation.AddProjectV2ItemByID.Item)
}

func (gateway *Gateway) Get(ctx context.Context, id string) (application.Result, error) {
	operation := fmt.Sprintf("querying Projects V2 item %q", id)
	var query struct {
		Node struct {
			Item node `graphql:"... on ProjectV2Item"`
		} `graphql:"node(id: $id)"`
	}
	if err := gateway.client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)}); err != nil {
		return application.Result{}, projectgraphql.Error(operation, err)
	}
	if query.Node.Item.ID == "" {
		return application.Result{}, projectgraphql.NotFound(operation)
	}
	return resultFromNode(query.Node.Item)
}

func (gateway *Gateway) SetArchived(ctx context.Context, projectID, itemID string, archived bool) (application.Result, error) {
	if archived {
		var mutation struct {
			ArchiveProjectV2Item struct{ Item node } `graphql:"archiveProjectV2Item(input: $input)"`
		}
		variables := githubv4.ArchiveProjectV2ItemInput{ProjectID: githubv4.ID(projectID), ItemID: githubv4.ID(itemID)}
		if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
			return application.Result{}, projectgraphql.Error(fmt.Sprintf("archiving Projects V2 item %q", itemID), err)
		}
		return resultFromNode(mutation.ArchiveProjectV2Item.Item)
	}
	var mutation struct {
		UnarchiveProjectV2Item struct{ Item node } `graphql:"unarchiveProjectV2Item(input: $input)"`
	}
	variables := githubv4.UnarchiveProjectV2ItemInput{ProjectID: githubv4.ID(projectID), ItemID: githubv4.ID(itemID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("unarchiving Projects V2 item %q", itemID), err)
	}
	return resultFromNode(mutation.UnarchiveProjectV2Item.Item)
}

func (gateway *Gateway) Remove(ctx context.Context, projectID, itemID string) error {
	var mutation struct {
		DeleteProjectV2Item struct{ DeletedItemID githubv4.String } `graphql:"deleteProjectV2Item(input: $input)"`
	}
	variables := githubv4.DeleteProjectV2ItemInput{ProjectID: githubv4.ID(projectID), ItemID: githubv4.ID(itemID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("deleting Projects V2 item %q", itemID), err)
	}
	return nil
}

var _ application.Store = (*Gateway)(nil)
