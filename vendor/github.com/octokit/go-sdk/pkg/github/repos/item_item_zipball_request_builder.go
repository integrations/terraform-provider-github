package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemZipballRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\zipball
type ItemItemZipballRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRef gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.zipball.item collection
// returns a *ItemItemZipballWithRefItemRequestBuilder when successful
func (m *ItemItemZipballRequestBuilder) ByRef(ref string)(*ItemItemZipballWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ref != "" {
        urlTplParams["ref"] = ref
    }
    return NewItemItemZipballWithRefItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemZipballRequestBuilderInternal instantiates a new ItemItemZipballRequestBuilder and sets the default values.
func NewItemItemZipballRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemZipballRequestBuilder) {
    m := &ItemItemZipballRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/zipball", pathParameters),
    }
    return m
}
// NewItemItemZipballRequestBuilder instantiates a new ItemItemZipballRequestBuilder and sets the default values.
func NewItemItemZipballRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemZipballRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemZipballRequestBuilderInternal(urlParams, requestAdapter)
}
