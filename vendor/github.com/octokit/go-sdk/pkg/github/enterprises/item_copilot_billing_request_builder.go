package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCopilotBillingRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\copilot\billing
type ItemCopilotBillingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCopilotBillingRequestBuilderInternal instantiates a new ItemCopilotBillingRequestBuilder and sets the default values.
func NewItemCopilotBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingRequestBuilder) {
    m := &ItemCopilotBillingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/copilot/billing", pathParameters),
    }
    return m
}
// NewItemCopilotBillingRequestBuilder instantiates a new ItemCopilotBillingRequestBuilder and sets the default values.
func NewItemCopilotBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotBillingRequestBuilderInternal(urlParams, requestAdapter)
}
// Seats the seats property
// returns a *ItemCopilotBillingSeatsRequestBuilder when successful
func (m *ItemCopilotBillingRequestBuilder) Seats()(*ItemCopilotBillingSeatsRequestBuilder) {
    return NewItemCopilotBillingSeatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
