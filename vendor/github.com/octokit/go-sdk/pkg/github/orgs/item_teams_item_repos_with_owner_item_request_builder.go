package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamsItemReposWithOwnerItemRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\repos\{owner}
type ItemTeamsItemReposWithOwnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.teams.item.repos.item.item collection
// returns a *ItemTeamsItemReposItemWithRepoItemRequestBuilder when successful
func (m *ItemTeamsItemReposWithOwnerItemRequestBuilder) ByRepo(repo string)(*ItemTeamsItemReposItemWithRepoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo != "" {
        urlTplParams["repo"] = repo
    }
    return NewItemTeamsItemReposItemWithRepoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal instantiates a new ItemTeamsItemReposWithOwnerItemRequestBuilder and sets the default values.
func NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemReposWithOwnerItemRequestBuilder) {
    m := &ItemTeamsItemReposWithOwnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/repos/{owner}", pathParameters),
    }
    return m
}
// NewItemTeamsItemReposWithOwnerItemRequestBuilder instantiates a new ItemTeamsItemReposWithOwnerItemRequestBuilder and sets the default values.
func NewItemTeamsItemReposWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemReposWithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
