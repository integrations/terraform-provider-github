package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsOidcRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\oidc
type ItemActionsOidcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsOidcRequestBuilderInternal instantiates a new ItemActionsOidcRequestBuilder and sets the default values.
func NewItemActionsOidcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcRequestBuilder) {
    m := &ItemActionsOidcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/oidc", pathParameters),
    }
    return m
}
// NewItemActionsOidcRequestBuilder instantiates a new ItemActionsOidcRequestBuilder and sets the default values.
func NewItemActionsOidcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsOidcRequestBuilderInternal(urlParams, requestAdapter)
}
// Customization the customization property
// returns a *ItemActionsOidcCustomizationRequestBuilder when successful
func (m *ItemActionsOidcRequestBuilder) Customization()(*ItemActionsOidcCustomizationRequestBuilder) {
    return NewItemActionsOidcCustomizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
