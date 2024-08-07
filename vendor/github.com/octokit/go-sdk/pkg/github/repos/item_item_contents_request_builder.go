package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemContentsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\contents
type ItemItemContentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPath gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.contents.item collection
// returns a *ItemItemContentsWithPathItemRequestBuilder when successful
func (m *ItemItemContentsRequestBuilder) ByPath(path string)(*ItemItemContentsWithPathItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if path != "" {
        urlTplParams["path"] = path
    }
    return NewItemItemContentsWithPathItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemContentsRequestBuilderInternal instantiates a new ItemItemContentsRequestBuilder and sets the default values.
func NewItemItemContentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsRequestBuilder) {
    m := &ItemItemContentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/contents", pathParameters),
    }
    return m
}
// NewItemItemContentsRequestBuilder instantiates a new ItemItemContentsRequestBuilder and sets the default values.
func NewItemItemContentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemContentsRequestBuilderInternal(urlParams, requestAdapter)
}
