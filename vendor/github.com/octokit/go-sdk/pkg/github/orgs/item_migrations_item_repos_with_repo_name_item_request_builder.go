package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMigrationsItemReposWithRepo_nameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}\repos\{repo_name}
type ItemMigrationsItemReposWithRepo_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal instantiates a new ItemMigrationsItemReposWithRepo_nameItemRequestBuilder and sets the default values.
func NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) {
    m := &ItemMigrationsItemReposWithRepo_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/migrations/{migration_id}/repos/{repo_name}", pathParameters),
    }
    return m
}
// NewItemMigrationsItemReposWithRepo_nameItemRequestBuilder instantiates a new ItemMigrationsItemReposWithRepo_nameItemRequestBuilder and sets the default values.
func NewItemMigrationsItemReposWithRepo_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Lock the lock property
// returns a *ItemMigrationsItemReposItemLockRequestBuilder when successful
func (m *ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) Lock()(*ItemMigrationsItemReposItemLockRequestBuilder) {
    return NewItemMigrationsItemReposItemLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
