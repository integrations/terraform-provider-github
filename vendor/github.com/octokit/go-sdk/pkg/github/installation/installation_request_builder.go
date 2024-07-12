package installation

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstallationRequestBuilder builds and executes requests for operations under \installation
type InstallationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewInstallationRequestBuilderInternal instantiates a new InstallationRequestBuilder and sets the default values.
func NewInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationRequestBuilder) {
    m := &InstallationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/installation", pathParameters),
    }
    return m
}
// NewInstallationRequestBuilder instantiates a new InstallationRequestBuilder and sets the default values.
func NewInstallationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// Repositories the repositories property
// returns a *RepositoriesRequestBuilder when successful
func (m *InstallationRequestBuilder) Repositories()(*RepositoriesRequestBuilder) {
    return NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Token the token property
// returns a *TokenRequestBuilder when successful
func (m *InstallationRequestBuilder) Token()(*TokenRequestBuilder) {
    return NewTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
