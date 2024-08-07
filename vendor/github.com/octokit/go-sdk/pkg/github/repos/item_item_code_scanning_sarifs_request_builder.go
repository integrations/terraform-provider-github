package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningSarifsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\sarifs
type ItemItemCodeScanningSarifsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BySarif_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.codeScanning.sarifs.item collection
// returns a *ItemItemCodeScanningSarifsWithSarif_ItemRequestBuilder when successful
func (m *ItemItemCodeScanningSarifsRequestBuilder) BySarif_id(sarif_id string)(*ItemItemCodeScanningSarifsWithSarif_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sarif_id != "" {
        urlTplParams["sarif_id"] = sarif_id
    }
    return NewItemItemCodeScanningSarifsWithSarif_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningSarifsRequestBuilderInternal instantiates a new ItemItemCodeScanningSarifsRequestBuilder and sets the default values.
func NewItemItemCodeScanningSarifsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningSarifsRequestBuilder) {
    m := &ItemItemCodeScanningSarifsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/sarifs", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningSarifsRequestBuilder instantiates a new ItemItemCodeScanningSarifsRequestBuilder and sets the default values.
func NewItemItemCodeScanningSarifsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningSarifsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningSarifsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post uploads SARIF data containing the results of a code scanning analysis to make the results available in a repository. For troubleshooting information, see "[Troubleshooting SARIF uploads](https://docs.github.com/code-security/code-scanning/troubleshooting-sarif)."There are two places where you can upload code scanning results. - If you upload to a pull request, for example `--ref refs/pull/42/merge` or `--ref refs/pull/42/head`, then the results appear as alerts in a pull request check. For more information, see "[Triaging code scanning alerts in pull requests](/code-security/secure-coding/triaging-code-scanning-alerts-in-pull-requests)." - If you upload to a branch, for example `--ref refs/heads/my-branch`, then the results appear in the **Security** tab for your repository. For more information, see "[Managing code scanning alerts for your repository](/code-security/secure-coding/managing-code-scanning-alerts-for-your-repository#viewing-the-alerts-for-a-repository)."You must compress the SARIF-formatted analysis data that you want to upload, using `gzip`, and then encode it as a Base64 format string. For example:```gzip -c analysis-data.sarif | base64 -w0```SARIF upload supports a maximum number of entries per the following data objects, and an analysis will be rejected if any of these objects is above its maximum value. For some objects, there are additional values over which the entries will be ignored while keeping the most important entries whenever applicable.To get the most out of your analysis when it includes data above the supported limits, try to optimize the analysis configuration. For example, for the CodeQL tool, identify and remove the most noisy queries. For more information, see "[SARIF results exceed one or more limits](https://docs.github.com/code-security/code-scanning/troubleshooting-sarif/results-exceed-limit)."| **SARIF data**                   | **Maximum values** | **Additional limits**                                                            ||----------------------------------|:------------------:|----------------------------------------------------------------------------------|| Runs per file                    |         20         |                                                                                  || Results per run                  |       25,000       | Only the top 5,000 results will be included, prioritized by severity.            || Rules per run                    |       25,000       |                                                                                  || Tool extensions per run          |        100         |                                                                                  || Thread Flow Locations per result |       10,000       | Only the top 1,000 Thread Flow Locations will be included, using prioritization. || Location per result             |       1,000        | Only 100 locations will be included.                                             || Tags per rule                   |         20         | Only 10 tags will be included.                                                   |The `202 Accepted` response includes an `id` value.You can use this ID to check the status of the upload by using it in the `/sarifs/{sarif_id}` endpoint.For more information, see "[Get information about a SARIF upload](/rest/code-scanning/code-scanning#get-information-about-a-sarif-upload)."OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.This endpoint is limited to 1,000 requests per hour for each user or app installation calling it.
// returns a CodeScanningSarifsReceiptable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningSarifsReceipt503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#upload-an-analysis-as-sarif-data
func (m *ItemItemCodeScanningSarifsRequestBuilder) Post(ctx context.Context, body ItemItemCodeScanningSarifsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningSarifsReceiptable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningSarifsReceipt503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningSarifsReceiptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningSarifsReceiptable), nil
}
// ToPostRequestInformation uploads SARIF data containing the results of a code scanning analysis to make the results available in a repository. For troubleshooting information, see "[Troubleshooting SARIF uploads](https://docs.github.com/code-security/code-scanning/troubleshooting-sarif)."There are two places where you can upload code scanning results. - If you upload to a pull request, for example `--ref refs/pull/42/merge` or `--ref refs/pull/42/head`, then the results appear as alerts in a pull request check. For more information, see "[Triaging code scanning alerts in pull requests](/code-security/secure-coding/triaging-code-scanning-alerts-in-pull-requests)." - If you upload to a branch, for example `--ref refs/heads/my-branch`, then the results appear in the **Security** tab for your repository. For more information, see "[Managing code scanning alerts for your repository](/code-security/secure-coding/managing-code-scanning-alerts-for-your-repository#viewing-the-alerts-for-a-repository)."You must compress the SARIF-formatted analysis data that you want to upload, using `gzip`, and then encode it as a Base64 format string. For example:```gzip -c analysis-data.sarif | base64 -w0```SARIF upload supports a maximum number of entries per the following data objects, and an analysis will be rejected if any of these objects is above its maximum value. For some objects, there are additional values over which the entries will be ignored while keeping the most important entries whenever applicable.To get the most out of your analysis when it includes data above the supported limits, try to optimize the analysis configuration. For example, for the CodeQL tool, identify and remove the most noisy queries. For more information, see "[SARIF results exceed one or more limits](https://docs.github.com/code-security/code-scanning/troubleshooting-sarif/results-exceed-limit)."| **SARIF data**                   | **Maximum values** | **Additional limits**                                                            ||----------------------------------|:------------------:|----------------------------------------------------------------------------------|| Runs per file                    |         20         |                                                                                  || Results per run                  |       25,000       | Only the top 5,000 results will be included, prioritized by severity.            || Rules per run                    |       25,000       |                                                                                  || Tool extensions per run          |        100         |                                                                                  || Thread Flow Locations per result |       10,000       | Only the top 1,000 Thread Flow Locations will be included, using prioritization. || Location per result             |       1,000        | Only 100 locations will be included.                                             || Tags per rule                   |         20         | Only 10 tags will be included.                                                   |The `202 Accepted` response includes an `id` value.You can use this ID to check the status of the upload by using it in the `/sarifs/{sarif_id}` endpoint.For more information, see "[Get information about a SARIF upload](/rest/code-scanning/code-scanning#get-information-about-a-sarif-upload)."OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.This endpoint is limited to 1,000 requests per hour for each user or app installation calling it.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningSarifsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemCodeScanningSarifsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemCodeScanningSarifsRequestBuilder when successful
func (m *ItemItemCodeScanningSarifsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningSarifsRequestBuilder) {
    return NewItemItemCodeScanningSarifsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
