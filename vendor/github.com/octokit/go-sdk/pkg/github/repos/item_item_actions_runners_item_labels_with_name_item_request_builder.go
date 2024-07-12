package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runners\{runner_id}\labels\{name}
type ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal instantiates a new ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder and sets the default values.
func NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    m := &ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runners/{runner_id}/labels/{name}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder instantiates a new ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder and sets the default values.
func NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a custom label from a self-hosted runner configuredin a repository. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.Authenticated users must have admin access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationErrorSimple error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#remove-a-custom-label-from-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseable), nil
}
// ToDeleteRequestInformation remove a custom label from a self-hosted runner configuredin a repository. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.Authenticated users must have admin access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder when successful
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    return NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
