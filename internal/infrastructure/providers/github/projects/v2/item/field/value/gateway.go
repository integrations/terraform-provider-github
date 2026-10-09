package value

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

type Gateway struct{ client *githubv4.Client }

func NewGateway(client *githubv4.Client) *Gateway { return &Gateway{client: client} }

func (gateway *Gateway) Set(ctx context.Context, input application.SetInput) (application.Result, error) {
	var mutation struct {
		UpdateProjectV2ItemFieldValue struct {
			Item struct{ ID githubv4.String } `graphql:"projectV2Item"`
		} `graphql:"updateProjectV2ItemFieldValue(input: $input)"`
	}
	variables := githubv4.UpdateProjectV2ItemFieldValueInput{ProjectID: githubv4.ID(input.ProjectID), ItemID: githubv4.ID(input.ItemID), FieldID: githubv4.ID(input.FieldID), Value: fieldValueInput(input.Value)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("setting Projects V2 item %q field %q", input.ItemID, input.FieldID), err)
	}
	return input.Value, nil
}

func (gateway *Gateway) Get(ctx context.Context, itemID, fieldID string) (application.Result, error) {
	var after *githubv4.String
	for {
		var query struct {
			Node struct {
				Item struct {
					FieldValues struct {
						Nodes    []node
						PageInfo pageInfo
					} `graphql:"fieldValues(first: 100, after: $after)"`
				} `graphql:"... on ProjectV2Item"`
			} `graphql:"node(id: $itemID)"`
		}
		variables := map[string]any{"itemID": githubv4.ID(itemID), "after": after}
		if err := gateway.client.Query(ctx, &query, variables); err != nil {
			return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying Projects V2 item %q field %q", itemID, fieldID), err)
		}
		for _, fieldValue := range query.Node.Item.FieldValues.Nodes {
			if nodeFieldID(fieldValue) == fieldID {
				return resultFromNode(fieldValue)
			}
		}
		if !bool(query.Node.Item.FieldValues.PageInfo.HasNextPage) {
			return application.Result{}, nil
		}
		after = new(query.Node.Item.FieldValues.PageInfo.EndCursor)
	}
}

func (gateway *Gateway) Clear(ctx context.Context, projectID, itemID, fieldID string) error {
	var mutation struct {
		ClearProjectV2ItemFieldValue struct {
			Item struct{ ID githubv4.String } `graphql:"projectV2Item"`
		} `graphql:"clearProjectV2ItemFieldValue(input: $input)"`
	}
	variables := githubv4.ClearProjectV2ItemFieldValueInput{ProjectID: githubv4.ID(projectID), ItemID: githubv4.ID(itemID), FieldID: githubv4.ID(fieldID)}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("clearing Projects V2 item %q field %q", itemID, fieldID), err)
	}
	return nil
}

var _ application.Store = (*Gateway)(nil)
