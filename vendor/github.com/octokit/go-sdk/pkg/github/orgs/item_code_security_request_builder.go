package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCodeSecurityRequestBuilder builds and executes requests for operations under \orgs\{org}\code-security
type ItemCodeSecurityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Configurations the configurations property
// returns a *ItemCodeSecurityConfigurationsRequestBuilder when successful
func (m *ItemCodeSecurityRequestBuilder) Configurations()(*ItemCodeSecurityConfigurationsRequestBuilder) {
    return NewItemCodeSecurityConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCodeSecurityRequestBuilderInternal instantiates a new ItemCodeSecurityRequestBuilder and sets the default values.
func NewItemCodeSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityRequestBuilder) {
    m := &ItemCodeSecurityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/code-security", pathParameters),
    }
    return m
}
// NewItemCodeSecurityRequestBuilder instantiates a new ItemCodeSecurityRequestBuilder and sets the default values.
func NewItemCodeSecurityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
