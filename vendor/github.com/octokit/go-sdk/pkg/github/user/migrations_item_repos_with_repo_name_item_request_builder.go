package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MigrationsItemReposWithRepo_nameItemRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}\repos\{repo_name}
type MigrationsItemReposWithRepo_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMigrationsItemReposWithRepo_nameItemRequestBuilderInternal instantiates a new MigrationsItemReposWithRepo_nameItemRequestBuilder and sets the default values.
func NewMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemReposWithRepo_nameItemRequestBuilder) {
    m := &MigrationsItemReposWithRepo_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/migrations/{migration_id}/repos/{repo_name}", pathParameters),
    }
    return m
}
// NewMigrationsItemReposWithRepo_nameItemRequestBuilder instantiates a new MigrationsItemReposWithRepo_nameItemRequestBuilder and sets the default values.
func NewMigrationsItemReposWithRepo_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Lock the lock property
// returns a *MigrationsItemReposItemLockRequestBuilder when successful
func (m *MigrationsItemReposWithRepo_nameItemRequestBuilder) Lock()(*MigrationsItemReposItemLockRequestBuilder) {
    return NewMigrationsItemReposItemLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
