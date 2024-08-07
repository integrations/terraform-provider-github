package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Marketplace_listingRequestBuilder builds and executes requests for operations under \marketplace_listing
type Marketplace_listingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *AccountsRequestBuilder when successful
func (m *Marketplace_listingRequestBuilder) Accounts()(*AccountsRequestBuilder) {
    return NewAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMarketplace_listingRequestBuilderInternal instantiates a new Marketplace_listingRequestBuilder and sets the default values.
func NewMarketplace_listingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_listingRequestBuilder) {
    m := &Marketplace_listingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing", pathParameters),
    }
    return m
}
// NewMarketplace_listingRequestBuilder instantiates a new Marketplace_listingRequestBuilder and sets the default values.
func NewMarketplace_listingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_listingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMarketplace_listingRequestBuilderInternal(urlParams, requestAdapter)
}
// Plans the plans property
// returns a *PlansRequestBuilder when successful
func (m *Marketplace_listingRequestBuilder) Plans()(*PlansRequestBuilder) {
    return NewPlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stubbed the stubbed property
// returns a *StubbedRequestBuilder when successful
func (m *Marketplace_listingRequestBuilder) Stubbed()(*StubbedRequestBuilder) {
    return NewStubbedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
