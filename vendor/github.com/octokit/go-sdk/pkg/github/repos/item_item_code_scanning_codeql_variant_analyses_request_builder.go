package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql\variant-analyses
type ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCodeql_variant_analysis_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.codeScanning.codeql.variantAnalyses.item collection
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) ByCodeql_variant_analysis_id(codeql_variant_analysis_id int32)(*ItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["codeql_variant_analysis_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(codeql_variant_analysis_id), 10)
    return NewItemItemCodeScanningCodeqlVariantAnalysesWithCodeql_variant_analysis_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql/variant-analyses", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a new CodeQL variant analysis, which will run a CodeQL query against one or more repositories.Get started by learning more about [running CodeQL queries at scale with Multi-Repository Variant Analysis](https://docs.github.com/code-security/codeql-for-vs-code/getting-started-with-codeql-for-vs-code/running-codeql-queries-at-scale-with-multi-repository-variant-analysis).Use the `owner` and `repo` parameters in the URL to specify the controller repository thatwill be used for running GitHub Actions workflows and storing the results of the CodeQL variant analysis.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a CodeScanningVariantAnalysisable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 422 status code
// returns a CodeScanningVariantAnalysis503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#create-a-codeql-variant-analysis
func (m *ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) Post(ctx context.Context, body ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
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
// ToPostRequestInformation creates a new CodeQL variant analysis, which will run a CodeQL query against one or more repositories.Get started by learning more about [running CodeQL queries at scale with Multi-Repository Variant Analysis](https://docs.github.com/code-security/codeql-for-vs-code/getting-started-with-codeql-for-vs-code/running-codeql-queries-at-scale-with-multi-repository-variant-analysis).Use the `owner` and `repo` parameters in the URL to specify the controller repository thatwill be used for running GitHub Actions workflows and storing the results of the CodeQL variant analysis.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) {
    return NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
