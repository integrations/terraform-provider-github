package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningCodeqlDatabasesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql\databases
type ItemItemCodeScanningCodeqlDatabasesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByLanguage gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.codeScanning.codeql.databases.item collection
// returns a *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) ByLanguage(language string)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if language != "" {
        urlTplParams["language"] = language
    }
    return NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlDatabasesRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlDatabasesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql/databases", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlDatabasesRequestBuilder instantiates a new ItemItemCodeScanningCodeqlDatabasesRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the CodeQL databases that are available in a repository.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a []CodeScanningCodeqlDatabaseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a Databases503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#list-codeql-databases-for-a-repository
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningCodeqlDatabaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDatabases503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningCodeqlDatabaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningCodeqlDatabaseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningCodeqlDatabaseable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the CodeQL databases that are available in a repository.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningCodeqlDatabasesRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlDatabasesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    return NewItemItemCodeScanningCodeqlDatabasesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
