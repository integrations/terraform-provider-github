package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OwnerItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}
type OwnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepoId gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item collection
// returns a *ItemRepoItemRequestBuilder when successful
func (m *OwnerItemRequestBuilder) ByRepoId(repoId string)(*ItemRepoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repoId != "" {
        urlTplParams["repo%2Did"] = repoId
    }
    return NewItemRepoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOwnerItemRequestBuilderInternal instantiates a new OwnerItemRequestBuilder and sets the default values.
func NewOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnerItemRequestBuilder) {
    m := &OwnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}", pathParameters),
    }
    return m
}
// NewOwnerItemRequestBuilder instantiates a new OwnerItemRequestBuilder and sets the default values.
func NewOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
