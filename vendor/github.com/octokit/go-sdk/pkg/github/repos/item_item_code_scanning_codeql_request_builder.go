package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningCodeqlRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\codeql
type ItemItemCodeScanningCodeqlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCodeScanningCodeqlRequestBuilderInternal instantiates a new ItemItemCodeScanningCodeqlRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/codeql", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlRequestBuilder instantiates a new ItemItemCodeScanningCodeqlRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlRequestBuilderInternal(urlParams, requestAdapter)
}
// Databases the databases property
// returns a *ItemItemCodeScanningCodeqlDatabasesRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlRequestBuilder) Databases()(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    return NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VariantAnalyses the variantAnalyses property
// returns a *ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder when successful
func (m *ItemItemCodeScanningCodeqlRequestBuilder) VariantAnalyses()(*ItemItemCodeScanningCodeqlVariantAnalysesRequestBuilder) {
    return NewItemItemCodeScanningCodeqlVariantAnalysesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
