package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql\variant-analyses\{codeql_variant_analysis_id}
type ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql/variant-analyses/{codeql_variant_analysis_id}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the summary of a CodeQL variant analysis.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningVariantAnalysisable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningVariantAnalysis503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#get-the-summary-of-a-codeql-variant-analysis
func (m *ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningVariantAnalysis503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningVariantAnalysisFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisable), nil
}
// Repos the repos property
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) Repos()(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder) {
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets the summary of a CodeQL variant analysis.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) {
    return NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
