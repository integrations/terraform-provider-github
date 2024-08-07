package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    ic8ccce7f7df3354ee09c704fd2c3c7a95354f442dcd3fefc8778b101a690d643 "github.com/octokit/go-sdk/pkg/github/repos/item/item/codescanning/analyses"
)

// ItemItemCodeScanningAnalysesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\analyses
type ItemItemCodeScanningAnalysesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters lists the details of all code scanning analyses for a repository,starting with the most recent.The response is paginated and you can use the `page` and `per_page` parametersto list the analyses you're interested in.By default 30 analyses are listed per page.The `rules_count` field in the response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.**Deprecation notice**:The `tool_name` field is deprecated and will, in future, not be included in the response for this endpoint. The example response reflects this change. The tool name can now be found inside the `tool` field.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
type ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *ic8ccce7f7df3354ee09c704fd2c3c7a95354f442dcd3fefc8778b101a690d643.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The Git reference for the analyses you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
    // Filter analyses belonging to the same SARIF upload.
    Sarif_id *string `uriparametername:"sarif_id"`
    // The property by which to sort the results.
    Sort *ic8ccce7f7df3354ee09c704fd2c3c7a95354f442dcd3fefc8778b101a690d643.GetSortQueryParameterType `uriparametername:"sort"`
    // The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both.
    Tool_guid *string `uriparametername:"tool_guid"`
    // The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both.
    Tool_name *string `uriparametername:"tool_name"`
}
// ByAnalysis_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.codeScanning.analyses.item collection
// returns a *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder when successful
func (m *ItemItemCodeScanningAnalysesRequestBuilder) ByAnalysis_id(analysis_id int32)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["analysis_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(analysis_id), 10)
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningAnalysesRequestBuilderInternal instantiates a new ItemItemCodeScanningAnalysesRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesRequestBuilder) {
    m := &ItemItemCodeScanningAnalysesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/analyses{?direction*,page*,per_page*,ref*,sarif_id*,sort*,tool_guid*,tool_name*}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAnalysesRequestBuilder instantiates a new ItemItemCodeScanningAnalysesRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAnalysesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the details of all code scanning analyses for a repository,starting with the most recent.The response is paginated and you can use the `page` and `per_page` parametersto list the analyses you're interested in.By default 30 analyses are listed per page.The `rules_count` field in the response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.**Deprecation notice**:The `tool_name` field is deprecated and will, in future, not be included in the response for this endpoint. The example response reflects this change. The tool name can now be found inside the `tool` field.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a []CodeScanningAnalysisable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a Analyses503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#list-code-scanning-analyses-for-a-repository
func (m *ItemItemCodeScanningAnalysesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateAnalyses503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAnalysisFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the details of all code scanning analyses for a repository,starting with the most recent.The response is paginated and you can use the `page` and `per_page` parametersto list the analyses you're interested in.By default 30 analyses are listed per page.The `rules_count` field in the response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.**Deprecation notice**:The `tool_name` field is deprecated and will, in future, not be included in the response for this endpoint. The example response reflects this change. The tool name can now be found inside the `tool` field.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAnalysesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeScanningAnalysesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningAnalysesRequestBuilder when successful
func (m *ItemItemCodeScanningAnalysesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAnalysesRequestBuilder) {
    return NewItemItemCodeScanningAnalysesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
