package field

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

type Gateway struct{ client *githubv4.Client }

func NewGateway(client *githubv4.Client) *Gateway { return &Gateway{client: client} }

func (gateway *Gateway) Create(ctx context.Context, input application.CreateInput) (string, error) {
	options, iteration := configurationInput(input.Configuration)
	variables := githubv4.CreateProjectV2FieldInput{ProjectID: githubv4.ID(input.ProjectID), Name: githubv4.String(input.Name), DataType: githubv4.ProjectV2CustomFieldType(input.DataType), SingleSelectOptions: options, IterationConfiguration: iteration}
	var mutation struct {
		CreateProjectV2Field struct {
			Field node `graphql:"projectV2Field"`
		} `graphql:"createProjectV2Field(input: $input)"`
	}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return "", projectgraphql.Error("creating Projects V2 field", err)
	}
	return nodeID(mutation.CreateProjectV2Field.Field), nil
}

func (gateway *Gateway) Get(ctx context.Context, id string) (application.Result, error) {
	var query struct {
		Node node `graphql:"node(id: $id)"`
	}
	if err := gateway.client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)}); err != nil {
		return application.Result{}, projectgraphql.Error(fmt.Sprintf("querying Projects V2 field %q", id), err)
	}
	return resultFromNode(query.Node)
}

func (gateway *Gateway) Update(ctx context.Context, input application.UpdateInput) error {
	name := githubv4.String(input.Name)
	variables := githubv4.UpdateProjectV2FieldInput{FieldID: githubv4.ID(input.ID), Name: &name}
	if input.Configuration != nil {
		variables.SingleSelectOptions, variables.IterationConfiguration = configurationInput(*input.Configuration)
	}
	var mutation struct {
		UpdateProjectV2Field struct {
			Field node `graphql:"projectV2Field"`
		} `graphql:"updateProjectV2Field(input: $input)"`
	}
	if err := gateway.client.Mutate(ctx, &mutation, variables, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("updating Projects V2 field %q", input.ID), err)
	}
	return nil
}

func (gateway *Gateway) Delete(ctx context.Context, id string) error {
	var mutation struct {
		DeleteProjectV2Field struct{ ClientMutationID githubv4.String } `graphql:"deleteProjectV2Field(input: $input)"`
	}
	if err := gateway.client.Mutate(ctx, &mutation, githubv4.DeleteProjectV2FieldInput{FieldID: githubv4.ID(id)}, nil); err != nil {
		return projectgraphql.Error(fmt.Sprintf("deleting Projects V2 field %q", id), err)
	}
	return nil
}

var _ application.Store = (*Gateway)(nil)
