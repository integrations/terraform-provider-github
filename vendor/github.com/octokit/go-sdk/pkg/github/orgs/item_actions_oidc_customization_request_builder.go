package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsOidcCustomizationRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\oidc\customization
type ItemActionsOidcCustomizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsOidcCustomizationRequestBuilderInternal instantiates a new ItemActionsOidcCustomizationRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationRequestBuilder) {
    m := &ItemActionsOidcCustomizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/oidc/customization", pathParameters),
    }
    return m
}
// NewItemActionsOidcCustomizationRequestBuilder instantiates a new ItemActionsOidcCustomizationRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsOidcCustomizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Sub the sub property
// returns a *ItemActionsOidcCustomizationSubRequestBuilder when successful
func (m *ItemActionsOidcCustomizationRequestBuilder) Sub()(*ItemActionsOidcCustomizationSubRequestBuilder) {
    return NewItemActionsOidcCustomizationSubRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
