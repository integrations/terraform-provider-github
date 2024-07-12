package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependencyGraphRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\dependency-graph
type ItemItemDependencyGraphRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Compare the compare property
// returns a *ItemItemDependencyGraphCompareRequestBuilder when successful
func (m *ItemItemDependencyGraphRequestBuilder) Compare()(*ItemItemDependencyGraphCompareRequestBuilder) {
    return NewItemItemDependencyGraphCompareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemDependencyGraphRequestBuilderInternal instantiates a new ItemItemDependencyGraphRequestBuilder and sets the default values.
func NewItemItemDependencyGraphRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphRequestBuilder) {
    m := &ItemItemDependencyGraphRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/dependency-graph", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphRequestBuilder instantiates a new ItemItemDependencyGraphRequestBuilder and sets the default values.
func NewItemItemDependencyGraphRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphRequestBuilderInternal(urlParams, requestAdapter)
}
// Sbom the sbom property
// returns a *ItemItemDependencyGraphSbomRequestBuilder when successful
func (m *ItemItemDependencyGraphRequestBuilder) Sbom()(*ItemItemDependencyGraphSbomRequestBuilder) {
    return NewItemItemDependencyGraphSbomRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Snapshots the snapshots property
// returns a *ItemItemDependencyGraphSnapshotsRequestBuilder when successful
func (m *ItemItemDependencyGraphRequestBuilder) Snapshots()(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    return NewItemItemDependencyGraphSnapshotsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
