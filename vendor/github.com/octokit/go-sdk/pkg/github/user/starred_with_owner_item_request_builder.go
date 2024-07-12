package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StarredWithOwnerItemRequestBuilder builds and executes requests for operations under \user\starred\{owner}
type StarredWithOwnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo gets an item from the github.com/octokit/go-sdk/pkg/github.user.starred.item.item collection
// returns a *StarredItemWithRepoItemRequestBuilder when successful
func (m *StarredWithOwnerItemRequestBuilder) ByRepo(repo string)(*StarredItemWithRepoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo != "" {
        urlTplParams["repo"] = repo
    }
    return NewStarredItemWithRepoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewStarredWithOwnerItemRequestBuilderInternal instantiates a new StarredWithOwnerItemRequestBuilder and sets the default values.
func NewStarredWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StarredWithOwnerItemRequestBuilder) {
    m := &StarredWithOwnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/starred/{owner}", pathParameters),
    }
    return m
}
// NewStarredWithOwnerItemRequestBuilder instantiates a new StarredWithOwnerItemRequestBuilder and sets the default values.
func NewStarredWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StarredWithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStarredWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
