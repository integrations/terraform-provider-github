package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql\variant-analyses\{codeql_variant_analysis_id}\repos\{repo_owner}
type ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo_name gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.codeScanning.codeql.variantAnalyses.item.repos.item.item collection
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder) ByRepo_name(repo_name string)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo_name != "" {
        urlTplParams["repo_name"] = repo_name
    }
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposItemWithRepo_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql/variant-analyses/{codeql_variant_analysis_id}/repos/{repo_owner}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlVariantAnalysesItemReposWithRepo_ownerItemRequestBuilderInternal(urlParams, requestAdapter)
}
