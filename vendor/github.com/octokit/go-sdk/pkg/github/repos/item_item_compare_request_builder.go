package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCompareRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\compare
type ItemItemCompareRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByBasehead gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.compare.item collection
// returns a *ItemItemCompareWithBaseheadItemRequestBuilder when successful
func (m *ItemItemCompareRequestBuilder) ByBasehead(basehead string)(*ItemItemCompareWithBaseheadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if basehead != "" {
        urlTplParams["basehead"] = basehead
    }
    return NewItemItemCompareWithBaseheadItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCompareRequestBuilderInternal instantiates a new ItemItemCompareRequestBuilder and sets the default values.
func NewItemItemCompareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCompareRequestBuilder) {
    m := &ItemItemCompareRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/compare", pathParameters),
    }
    return m
}
// NewItemItemCompareRequestBuilder instantiates a new ItemItemCompareRequestBuilder and sets the default values.
func NewItemItemCompareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCompareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCompareRequestBuilderInternal(urlParams, requestAdapter)
}
