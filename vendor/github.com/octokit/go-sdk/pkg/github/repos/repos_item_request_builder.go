package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ReposItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}
type ReposItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByOwnerId gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item collection
// returns a *ItemOwnerItemRequestBuilder when successful
func (m *ReposItemRequestBuilder) ByOwnerId(ownerId string)(*ItemOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ownerId != "" {
        urlTplParams["Owner%2Did"] = ownerId
    }
    return NewItemOwnerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewReposItemRequestBuilderInternal instantiates a new ReposItemRequestBuilder and sets the default values.
func NewReposItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReposItemRequestBuilder) {
    m := &ReposItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}", pathParameters),
    }
    return m
}
// NewReposItemRequestBuilder instantiates a new ReposItemRequestBuilder and sets the default values.
func NewReposItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReposItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReposItemRequestBuilderInternal(urlParams, requestAdapter)
}
