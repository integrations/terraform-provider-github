package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsBillingRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\billing
type ItemSettingsBillingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Actions the actions property
// returns a *ItemSettingsBillingActionsRequestBuilder when successful
func (m *ItemSettingsBillingRequestBuilder) Actions()(*ItemSettingsBillingActionsRequestBuilder) {
    return NewItemSettingsBillingActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSettingsBillingRequestBuilderInternal instantiates a new ItemSettingsBillingRequestBuilder and sets the default values.
func NewItemSettingsBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingRequestBuilder) {
    m := &ItemSettingsBillingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/billing", pathParameters),
    }
    return m
}
// NewItemSettingsBillingRequestBuilder instantiates a new ItemSettingsBillingRequestBuilder and sets the default values.
func NewItemSettingsBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingRequestBuilderInternal(urlParams, requestAdapter)
}
// Packages the packages property
// returns a *ItemSettingsBillingPackagesRequestBuilder when successful
func (m *ItemSettingsBillingRequestBuilder) Packages()(*ItemSettingsBillingPackagesRequestBuilder) {
    return NewItemSettingsBillingPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedStorage the sharedStorage property
// returns a *ItemSettingsBillingSharedStorageRequestBuilder when successful
func (m *ItemSettingsBillingRequestBuilder) SharedStorage()(*ItemSettingsBillingSharedStorageRequestBuilder) {
    return NewItemSettingsBillingSharedStorageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
