package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemReleasesTagsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\releases\tags
type ItemItemReleasesTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTag gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.releases.tags.item collection
// returns a *ItemItemReleasesTagsWithTagItemRequestBuilder when successful
func (m *ItemItemReleasesTagsRequestBuilder) ByTag(tag string)(*ItemItemReleasesTagsWithTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tag != "" {
        urlTplParams["tag"] = tag
    }
    return NewItemItemReleasesTagsWithTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemReleasesTagsRequestBuilderInternal instantiates a new ItemItemReleasesTagsRequestBuilder and sets the default values.
func NewItemItemReleasesTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesTagsRequestBuilder) {
    m := &ItemItemReleasesTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/releases/tags", pathParameters),
    }
    return m
}
// NewItemItemReleasesTagsRequestBuilder instantiates a new ItemItemReleasesTagsRequestBuilder and sets the default values.
func NewItemItemReleasesTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReleasesTagsRequestBuilderInternal(urlParams, requestAdapter)
}
