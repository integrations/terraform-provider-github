package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemGitRefRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\git\ref
type ItemItemGitRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRef gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.git.ref.item collection
// returns a *ItemItemGitRefWithRefItemRequestBuilder when successful
func (m *ItemItemGitRefRequestBuilder) ByRef(ref string)(*ItemItemGitRefWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ref != "" {
        urlTplParams["ref"] = ref
    }
    return NewItemItemGitRefWithRefItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemGitRefRequestBuilderInternal instantiates a new ItemItemGitRefRequestBuilder and sets the default values.
func NewItemItemGitRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitRefRequestBuilder) {
    m := &ItemItemGitRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/git/ref", pathParameters),
    }
    return m
}
// NewItemItemGitRefRequestBuilder instantiates a new ItemItemGitRefRequestBuilder and sets the default values.
func NewItemItemGitRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitRefRequestBuilderInternal(urlParams, requestAdapter)
}
