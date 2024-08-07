package marketplace_listing

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StubbedAccountsRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed\accounts
type StubbedAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByAccount_id gets an item from the github.com/octokit/go-sdk/pkg/github.marketplace_listing.stubbed.accounts.item collection
// returns a *StubbedAccountsWithAccount_ItemRequestBuilder when successful
func (m *StubbedAccountsRequestBuilder) ByAccount_id(account_id int32)(*StubbedAccountsWithAccount_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["account_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(account_id), 10)
    return NewStubbedAccountsWithAccount_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewStubbedAccountsRequestBuilderInternal instantiates a new StubbedAccountsRequestBuilder and sets the default values.
func NewStubbedAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedAccountsRequestBuilder) {
    m := &StubbedAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing/stubbed/accounts", pathParameters),
    }
    return m
}
// NewStubbedAccountsRequestBuilder instantiates a new StubbedAccountsRequestBuilder and sets the default values.
func NewStubbedAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
