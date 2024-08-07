package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql\variant-analyses\{codeql_variant_analysis_id}\repos\{repo_owner}\{repo_name}
type ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql/variant-analyses/{codeql_variant_analysis_id}/repos/{repo_owner}/{repo_name}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the analysis status of a repository in a CodeQL variant analysis.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningVariantAnalysisRepoTaskable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningVariantAnalysisRepoTask503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#get-the-analysis-status-of-a-repository-in-a-codeql-variant-analysis
func (m *ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisRepoTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningVariantAnalysisRepoTask503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningVariantAnalysisRepoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisRepoTaskable), nil
}
// ToGetRequestInformation gets the analysis status of a repository in a CodeQL variant analysis.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) {
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
