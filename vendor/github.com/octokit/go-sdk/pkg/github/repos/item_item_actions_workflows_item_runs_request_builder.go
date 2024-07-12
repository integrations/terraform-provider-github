package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia77747759b3f91e25296248277ed17f00dd8b68db982a03262ffc61b4720274c "github.com/octokit/go-sdk/pkg/github/repos/item/item/actions/workflows/item/runs"
)

// ItemItemActionsWorkflowsItemRunsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\workflows\{workflow_id}\runs
type ItemItemActionsWorkflowsItemRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#parameters).Anyone with read access to the repository can use this endpointOAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
type ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters struct {
    // Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run.
    Actor *string `uriparametername:"actor"`
    // Returns workflow runs associated with a branch. Use the name of the branch of the `push`.
    Branch *string `uriparametername:"branch"`
    // Returns workflow runs with the `check_suite_id` that you specify.
    Check_suite_id *int32 `uriparametername:"check_suite_id"`
    // Returns workflow runs created within the given date-time range. For more information on the syntax, see "[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates)."
    Created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"created"`
    // Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see "[Events that trigger workflows](https://docs.github.com/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows)."
    Event *string `uriparametername:"event"`
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool `uriparametername:"exclude_pull_requests"`
    // Only returns workflow runs that are associated with the specified `head_sha`.
    Head_sha *string `uriparametername:"head_sha"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub Actions can set a status of `waiting`, `pending`, or `requested`.
    Status *ia77747759b3f91e25296248277ed17f00dd8b68db982a03262ffc61b4720274c.GetStatusQueryParameterType `uriparametername:"status"`
}
// NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal instantiates a new ItemItemActionsWorkflowsItemRunsRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    m := &ItemItemActionsWorkflowsItemRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/workflows/{workflow_id}/runs{?actor*,branch*,check_suite_id*,created*,event*,exclude_pull_requests*,head_sha*,page*,per_page*,status*}", pathParameters),
    }
    return m
}
// NewItemItemActionsWorkflowsItemRunsRequestBuilder instantiates a new ItemItemActionsWorkflowsItemRunsRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#parameters).Anyone with read access to the repository can use this endpointOAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a ItemItemActionsWorkflowsItemRunsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#list-workflow-runs-for-a-workflow
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters])(ItemItemActionsWorkflowsItemRunsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsWorkflowsItemRunsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsWorkflowsItemRunsGetResponseable), nil
}
// ToGetRequestInformation list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#parameters).Anyone with read access to the repository can use this endpointOAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsWorkflowsItemRunsRequestBuilder when successful
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    return NewItemItemActionsWorkflowsItemRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
