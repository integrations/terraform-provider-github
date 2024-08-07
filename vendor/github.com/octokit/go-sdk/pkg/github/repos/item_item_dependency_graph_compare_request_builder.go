package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependencyGraphCompareRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\dependency-graph\compare
type ItemItemDependencyGraphCompareRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByBasehead gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.dependencyGraph.compare.item collection
// returns a *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder when successful
func (m *ItemItemDependencyGraphCompareRequestBuilder) ByBasehead(basehead string)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if basehead != "" {
        urlTplParams["basehead"] = basehead
    }
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemDependencyGraphCompareRequestBuilderInternal instantiates a new ItemItemDependencyGraphCompareRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareRequestBuilder) {
    m := &ItemItemDependencyGraphCompareRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/dependency-graph/compare", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphCompareRequestBuilder instantiates a new ItemItemDependencyGraphCompareRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphCompareRequestBuilderInternal(urlParams, requestAdapter)
}
