package teams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemReposWithOwnerItemRequestBuilder builds and executes requests for operations under \teams\{team_id}\repos\{owner}
type ItemReposWithOwnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo gets an item from the github.com/octokit/go-sdk/pkg/github.teams.item.repos.item.item collection
// Deprecated: 
// returns a *ItemReposItemWithRepoItemRequestBuilder when successful
func (m *ItemReposWithOwnerItemRequestBuilder) ByRepo(repo string)(*ItemReposItemWithRepoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo != "" {
        urlTplParams["repo"] = repo
    }
    return NewItemReposItemWithRepoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemReposWithOwnerItemRequestBuilderInternal instantiates a new ItemReposWithOwnerItemRequestBuilder and sets the default values.
func NewItemReposWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposWithOwnerItemRequestBuilder) {
    m := &ItemReposWithOwnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/repos/{owner}", pathParameters),
    }
    return m
}
// NewItemReposWithOwnerItemRequestBuilder instantiates a new ItemReposWithOwnerItemRequestBuilder and sets the default values.
func NewItemReposWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposWithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReposWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
