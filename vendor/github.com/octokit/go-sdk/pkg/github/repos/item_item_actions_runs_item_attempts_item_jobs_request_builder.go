package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\attempts\{attempt_number}\jobs
type ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsItemAttemptsItemJobsRequestBuilderGetQueryParameters lists jobs for a specific workflow run attempt. You can use parameters to narrow the list of results. For more informationabout using parameters, see [Parameters](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint  with a private repository.
type ItemItemActionsRunsItemAttemptsItemJobsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilderInternal instantiates a new ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/attempts/{attempt_number}/jobs{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilder instantiates a new ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists jobs for a specific workflow run attempt. You can use parameters to narrow the list of results. For more informationabout using parameters, see [Parameters](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint  with a private repository.
// returns a ItemItemActionsRunsItemAttemptsItemJobsGetResponseable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-jobs#list-jobs-for-a-workflow-run-attempt
func (m *ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunsItemAttemptsItemJobsRequestBuilderGetQueryParameters])(ItemItemActionsRunsItemAttemptsItemJobsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunsItemAttemptsItemJobsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunsItemAttemptsItemJobsGetResponseable), nil
}
// ToGetRequestInformation lists jobs for a specific workflow run attempt. You can use parameters to narrow the list of results. For more informationabout using parameters, see [Parameters](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint  with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunsItemAttemptsItemJobsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder when successful
func (m *ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
