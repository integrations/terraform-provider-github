package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsRunsWithRun_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}
type ItemItemActionsRunsWithRun_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters gets a specific workflow run.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
type ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters struct {
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool `uriparametername:"exclude_pull_requests"`
}
// Approvals the approvals property
// returns a *ItemItemActionsRunsItemApprovalsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Approvals()(*ItemItemActionsRunsItemApprovalsRequestBuilder) {
    return NewItemItemActionsRunsItemApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Approve the approve property
// returns a *ItemItemActionsRunsItemApproveRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Approve()(*ItemItemActionsRunsItemApproveRequestBuilder) {
    return NewItemItemActionsRunsItemApproveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Artifacts the artifacts property
// returns a *ItemItemActionsRunsItemArtifactsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Artifacts()(*ItemItemActionsRunsItemArtifactsRequestBuilder) {
    return NewItemItemActionsRunsItemArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attempts the attempts property
// returns a *ItemItemActionsRunsItemAttemptsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Attempts()(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel the cancel property
// returns a *ItemItemActionsRunsItemCancelRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Cancel()(*ItemItemActionsRunsItemCancelRequestBuilder) {
    return NewItemItemActionsRunsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal instantiates a new ItemItemActionsRunsWithRun_ItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    m := &ItemItemActionsRunsWithRun_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}{?exclude_pull_requests*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsWithRun_ItemRequestBuilder instantiates a new ItemItemActionsRunsWithRun_ItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsWithRun_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a specific workflow run.Anyone with write access to the repository can use this endpoint.If the repository is private, OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#delete-a-workflow-run
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Deployment_protection_rule the deployment_protection_rule property
// returns a *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Deployment_protection_rule()(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ForceCancel the forceCancel property
// returns a *ItemItemActionsRunsItemForceCancelRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ForceCancel()(*ItemItemActionsRunsItemForceCancelRequestBuilder) {
    return NewItemItemActionsRunsItemForceCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a specific workflow run.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a WorkflowRunable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#get-a-workflow-run
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateWorkflowRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable), nil
}
// Jobs the jobs property
// returns a *ItemItemActionsRunsItemJobsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Jobs()(*ItemItemActionsRunsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logs the logs property
// returns a *ItemItemActionsRunsItemLogsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Logs()(*ItemItemActionsRunsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pending_deployments the pending_deployments property
// returns a *ItemItemActionsRunsItemPending_deploymentsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Pending_deployments()(*ItemItemActionsRunsItemPending_deploymentsRequestBuilder) {
    return NewItemItemActionsRunsItemPending_deploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rerun the rerun property
// returns a *ItemItemActionsRunsItemRerunRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Rerun()(*ItemItemActionsRunsItemRerunRequestBuilder) {
    return NewItemItemActionsRunsItemRerunRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RerunFailedJobs the rerunFailedJobs property
// returns a *ItemItemActionsRunsItemRerunFailedJobsRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) RerunFailedJobs()(*ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) {
    return NewItemItemActionsRunsItemRerunFailedJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Timing the timing property
// returns a *ItemItemActionsRunsItemTimingRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Timing()(*ItemItemActionsRunsItemTimingRequestBuilder) {
    return NewItemItemActionsRunsItemTimingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a specific workflow run.Anyone with write access to the repository can use this endpoint.If the repository is private, OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation gets a specific workflow run.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsWithRun_ItemRequestBuilder when successful
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    return NewItemItemActionsRunsWithRun_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
