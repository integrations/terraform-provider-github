package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions
type ItemItemActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Artifacts the artifacts property
// returns a *ItemItemActionsArtifactsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Artifacts()(*ItemItemActionsArtifactsRequestBuilder) {
    return NewItemItemActionsArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cache the cache property
// returns a *ItemItemActionsCacheRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Cache()(*ItemItemActionsCacheRequestBuilder) {
    return NewItemItemActionsCacheRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Caches the caches property
// returns a *ItemItemActionsCachesRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Caches()(*ItemItemActionsCachesRequestBuilder) {
    return NewItemItemActionsCachesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRequestBuilderInternal instantiates a new ItemItemActionsRequestBuilder and sets the default values.
func NewItemItemActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequestBuilder) {
    m := &ItemItemActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions", pathParameters),
    }
    return m
}
// NewItemItemActionsRequestBuilder instantiates a new ItemItemActionsRequestBuilder and sets the default values.
func NewItemItemActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Jobs the jobs property
// returns a *ItemItemActionsJobsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Jobs()(*ItemItemActionsJobsRequestBuilder) {
    return NewItemItemActionsJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oidc the oidc property
// returns a *ItemItemActionsOidcRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Oidc()(*ItemItemActionsOidcRequestBuilder) {
    return NewItemItemActionsOidcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OrganizationSecrets the organizationSecrets property
// returns a *ItemItemActionsOrganizationSecretsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) OrganizationSecrets()(*ItemItemActionsOrganizationSecretsRequestBuilder) {
    return NewItemItemActionsOrganizationSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OrganizationVariables the organizationVariables property
// returns a *ItemItemActionsOrganizationVariablesRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) OrganizationVariables()(*ItemItemActionsOrganizationVariablesRequestBuilder) {
    return NewItemItemActionsOrganizationVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions the permissions property
// returns a *ItemItemActionsPermissionsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Permissions()(*ItemItemActionsPermissionsRequestBuilder) {
    return NewItemItemActionsPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runners the runners property
// returns a *ItemItemActionsRunnersRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Runners()(*ItemItemActionsRunnersRequestBuilder) {
    return NewItemItemActionsRunnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runs the runs property
// returns a *ItemItemActionsRunsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Runs()(*ItemItemActionsRunsRequestBuilder) {
    return NewItemItemActionsRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Secrets the secrets property
// returns a *ItemItemActionsSecretsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Secrets()(*ItemItemActionsSecretsRequestBuilder) {
    return NewItemItemActionsSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Variables the variables property
// returns a *ItemItemActionsVariablesRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Variables()(*ItemItemActionsVariablesRequestBuilder) {
    return NewItemItemActionsVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Workflows the workflows property
// returns a *ItemItemActionsWorkflowsRequestBuilder when successful
func (m *ItemItemActionsRequestBuilder) Workflows()(*ItemItemActionsWorkflowsRequestBuilder) {
    return NewItemItemActionsWorkflowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
