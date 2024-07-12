package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StubbedRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed
type StubbedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *StubbedAccountsRequestBuilder when successful
func (m *StubbedRequestBuilder) Accounts()(*StubbedAccountsRequestBuilder) {
    return NewStubbedAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStubbedRequestBuilderInternal instantiates a new StubbedRequestBuilder and sets the default values.
func NewStubbedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedRequestBuilder) {
    m := &StubbedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing/stubbed", pathParameters),
    }
    return m
}
// NewStubbedRequestBuilder instantiates a new StubbedRequestBuilder and sets the default values.
func NewStubbedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedRequestBuilderInternal(urlParams, requestAdapter)
}
// Plans the plans property
// returns a *StubbedPlansRequestBuilder when successful
func (m *StubbedRequestBuilder) Plans()(*StubbedPlansRequestBuilder) {
    return NewStubbedPlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
