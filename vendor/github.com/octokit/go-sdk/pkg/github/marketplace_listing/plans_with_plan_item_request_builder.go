package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PlansWithPlan_ItemRequestBuilder builds and executes requests for operations under \marketplace_listing\plans\{plan_id}
type PlansWithPlan_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *PlansItemAccountsRequestBuilder when successful
func (m *PlansWithPlan_ItemRequestBuilder) Accounts()(*PlansItemAccountsRequestBuilder) {
    return NewPlansItemAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPlansWithPlan_ItemRequestBuilderInternal instantiates a new PlansWithPlan_ItemRequestBuilder and sets the default values.
func NewPlansWithPlan_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlansWithPlan_ItemRequestBuilder) {
    m := &PlansWithPlan_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing/plans/{plan_id}", pathParameters),
    }
    return m
}
// NewPlansWithPlan_ItemRequestBuilder instantiates a new PlansWithPlan_ItemRequestBuilder and sets the default values.
func NewPlansWithPlan_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlansWithPlan_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlansWithPlan_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
