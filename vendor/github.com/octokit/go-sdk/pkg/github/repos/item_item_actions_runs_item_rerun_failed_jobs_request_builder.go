package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsRunsItemRerunFailedJobsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\rerun-failed-jobs
type ItemItemActionsRunsItemRerunFailedJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsRunsItemRerunFailedJobsRequestBuilderInternal instantiates a new ItemItemActionsRunsItemRerunFailedJobsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemRerunFailedJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) {
    m := &ItemItemActionsRunsItemRerunFailedJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/rerun-failed-jobs", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemRerunFailedJobsRequestBuilder instantiates a new ItemItemActionsRunsItemRerunFailedJobsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemRerunFailedJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemRerunFailedJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post re-run all of the failed jobs and their dependent jobs in a workflow run using the `id` of the workflow run.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a EmptyObjectable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#re-run-failed-jobs-from-a-workflow-run
func (m *ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) Post(ctx context.Context, body ItemItemActionsRunsItemRerunFailedJobsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EmptyObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateEmptyObjectFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EmptyObjectable), nil
}
// ToPostRequestInformation re-run all of the failed jobs and their dependent jobs in a workflow run using the `id` of the workflow run.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemActionsRunsItemRerunFailedJobsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemRerunFailedJobsRequestBuilder when successful
func (m *ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) {
    return NewItemItemActionsRunsItemRerunFailedJobsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
