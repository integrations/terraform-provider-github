package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemAssigneesWithAssigneeItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\assignees\{assignee}
type ItemItemAssigneesWithAssigneeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal instantiates a new ItemItemAssigneesWithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    m := &ItemItemAssigneesWithAssigneeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/assignees/{assignee}", pathParameters),
    }
    return m
}
// NewItemItemAssigneesWithAssigneeItemRequestBuilder instantiates a new ItemItemAssigneesWithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemAssigneesWithAssigneeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks if a user has permission to be assigned to an issue in this repository.If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.Otherwise a `404` status code is returned.
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/assignees#check-if-a-user-can-be-assigned
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation checks if a user has permission to be assigned to an issue in this repository.If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.Otherwise a `404` status code is returned.
// returns a *RequestInformation when successful
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemAssigneesWithAssigneeItemRequestBuilder when successful
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    return NewItemItemAssigneesWithAssigneeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
