package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMigrationsItemReposRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}\repos
type ItemMigrationsItemReposRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo_name gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.migrations.item.repos.item collection
// returns a *ItemMigrationsItemReposWithRepo_nameItemRequestBuilder when successful
func (m *ItemMigrationsItemReposRequestBuilder) ByRepo_name(repo_name string)(*ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo_name != "" {
        urlTplParams["repo_name"] = repo_name
    }
    return NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMigrationsItemReposRequestBuilderInternal instantiates a new ItemMigrationsItemReposRequestBuilder and sets the default values.
func NewItemMigrationsItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposRequestBuilder) {
    m := &ItemMigrationsItemReposRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/migrations/{migration_id}/repos", pathParameters),
    }
    return m
}
// NewItemMigrationsItemReposRequestBuilder instantiates a new ItemMigrationsItemReposRequestBuilder and sets the default values.
func NewItemMigrationsItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
