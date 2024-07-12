package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsRequestBuilder builds and executes requests for operations under \users\{username}\settings
type ItemSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Billing the billing property
// returns a *ItemSettingsBillingRequestBuilder when successful
func (m *ItemSettingsRequestBuilder) Billing()(*ItemSettingsBillingRequestBuilder) {
    return NewItemSettingsBillingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSettingsRequestBuilderInternal instantiates a new ItemSettingsRequestBuilder and sets the default values.
func NewItemSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsRequestBuilder) {
    m := &ItemSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/settings", pathParameters),
    }
    return m
}
// NewItemSettingsRequestBuilder instantiates a new ItemSettingsRequestBuilder and sets the default values.
func NewItemSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
