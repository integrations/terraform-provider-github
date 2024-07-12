package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstallationsWithInstallation_ItemRequestBuilder builds and executes requests for operations under \user\installations\{installation_id}
type InstallationsWithInstallation_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewInstallationsWithInstallation_ItemRequestBuilderInternal instantiates a new InstallationsWithInstallation_ItemRequestBuilder and sets the default values.
func NewInstallationsWithInstallation_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationsWithInstallation_ItemRequestBuilder) {
    m := &InstallationsWithInstallation_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/installations/{installation_id}", pathParameters),
    }
    return m
}
// NewInstallationsWithInstallation_ItemRequestBuilder instantiates a new InstallationsWithInstallation_ItemRequestBuilder and sets the default values.
func NewInstallationsWithInstallation_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationsWithInstallation_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstallationsWithInstallation_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Repositories the repositories property
// returns a *InstallationsItemRepositoriesRequestBuilder when successful
func (m *InstallationsWithInstallation_ItemRequestBuilder) Repositories()(*InstallationsItemRepositoriesRequestBuilder) {
    return NewInstallationsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
