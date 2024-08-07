package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql\variant-analyses\{codeql_variant_analysis_id}\repos
type ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo_owner gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.codeScanning.codeql.variantAnalyses.item.repos.item collection
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder) ByRepo_owner(repo_owner string)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo_owner != "" {
        urlTplParams["repo_owner"] = repo_owner
    }
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql/variant-analyses/{codeql_variant_analysis_id}/repos", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
