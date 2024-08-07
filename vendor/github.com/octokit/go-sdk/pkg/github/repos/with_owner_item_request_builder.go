package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithOwnerItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}
type WithOwnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item collection
func (m *WithOwnerItemRequestBuilder) ByRepo(repo string)(*ItemWithRepoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo != "" {
        urlTplParams["repo"] = repo
    }
    return NewItemWithRepoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithOwnerItemRequestBuilderInternal instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOwnerItemRequestBuilder) {
    m := &WithOwnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}", pathParameters),
    }
    return m
}
// NewWithOwnerItemRequestBuilder instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
