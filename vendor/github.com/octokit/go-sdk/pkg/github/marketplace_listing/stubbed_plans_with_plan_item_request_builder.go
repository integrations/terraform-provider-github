package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StubbedPlansWithPlan_ItemRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed\plans\{plan_id}
type StubbedPlansWithPlan_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *StubbedPlansItemAccountsRequestBuilder when successful
func (m *StubbedPlansWithPlan_ItemRequestBuilder) Accounts()(*StubbedPlansItemAccountsRequestBuilder) {
    return NewStubbedPlansItemAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStubbedPlansWithPlan_ItemRequestBuilderInternal instantiates a new StubbedPlansWithPlan_ItemRequestBuilder and sets the default values.
func NewStubbedPlansWithPlan_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedPlansWithPlan_ItemRequestBuilder) {
    m := &StubbedPlansWithPlan_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing/stubbed/plans/{plan_id}", pathParameters),
    }
    return m
}
// NewStubbedPlansWithPlan_ItemRequestBuilder instantiates a new StubbedPlansWithPlan_ItemRequestBuilder and sets the default values.
func NewStubbedPlansWithPlan_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedPlansWithPlan_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedPlansWithPlan_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
