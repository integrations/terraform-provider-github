package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommitsItemCheckSuitesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\commits\{commit_sha-id}\check-suites
type ItemItemCommitsItemCheckSuitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint on a private repository.
type ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters struct {
    // Filters check suites by GitHub App `id`.
    App_id *int32 `uriparametername:"app_id"`
    // Returns check runs with the specified `name`.
    Check_name *string `uriparametername:"check_name"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemCommitsItemCheckSuitesRequestBuilderInternal instantiates a new ItemItemCommitsItemCheckSuitesRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    m := &ItemItemCommitsItemCheckSuitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/commits/{commit_sha%2Did}/check-suites{?app_id*,check_name*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemCommitsItemCheckSuitesRequestBuilder instantiates a new ItemItemCommitsItemCheckSuitesRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsItemCheckSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint on a private repository.
// returns a ItemItemCommitsItemCheckSuitesGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/suites#list-check-suites-for-a-git-reference
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters])(ItemItemCommitsItemCheckSuitesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCommitsItemCheckSuitesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCommitsItemCheckSuitesGetResponseable), nil
}
// ToGetRequestInformation lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint on a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCommitsItemCheckSuitesRequestBuilder when successful
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    return NewItemItemCommitsItemCheckSuitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
