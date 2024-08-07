package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\deployments\{deployment_id}\statuses\{status_id}
type ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal instantiates a new ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    m := &ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/deployments/{deployment_id}/statuses/{status_id}", pathParameters),
    }
    return m
}
// NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder instantiates a new ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get users with pull access can view a deployment status for a deployment:
// returns a DeploymentStatusable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/statuses#get-a-deployment-status
func (m *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentStatusable), nil
}
// ToGetRequestInformation users with pull access can view a deployment status for a deployment:
// returns a *RequestInformation when successful
func (m *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder when successful
func (m *ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    return NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
