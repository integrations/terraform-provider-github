package marketplace_listing

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AccountsRequestBuilder builds and executes requests for operations under \marketplace_listing\accounts
type AccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByAccount_id gets an item from the github.com/octokit/go-sdk/pkg/github.marketplace_listing.accounts.item collection
// returns a *AccountsWithAccount_ItemRequestBuilder when successful
func (m *AccountsRequestBuilder) ByAccount_id(account_id int32)(*AccountsWithAccount_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["account_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(account_id), 10)
    return NewAccountsWithAccount_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAccountsRequestBuilderInternal instantiates a new AccountsRequestBuilder and sets the default values.
func NewAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountsRequestBuilder) {
    m := &AccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing/accounts", pathParameters),
    }
    return m
}
// NewAccountsRequestBuilder instantiates a new AccountsRequestBuilder and sets the default values.
func NewAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
