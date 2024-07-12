package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MigrationsItemReposRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}\repos
type MigrationsItemReposRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRepo_name gets an item from the github.com/octokit/go-sdk/pkg/github.user.migrations.item.repos.item collection
// returns a *MigrationsItemReposWithRepo_nameItemRequestBuilder when successful
func (m *MigrationsItemReposRequestBuilder) ByRepo_name(repo_name string)(*MigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if repo_name != "" {
        urlTplParams["repo_name"] = repo_name
    }
    return NewMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMigrationsItemReposRequestBuilderInternal instantiates a new MigrationsItemReposRequestBuilder and sets the default values.
func NewMigrationsItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemReposRequestBuilder) {
    m := &MigrationsItemReposRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/migrations/{migration_id}/repos", pathParameters),
    }
    return m
}
// NewMigrationsItemReposRequestBuilder instantiates a new MigrationsItemReposRequestBuilder and sets the default values.
func NewMigrationsItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
