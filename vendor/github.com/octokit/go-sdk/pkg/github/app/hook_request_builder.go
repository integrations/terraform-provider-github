package app

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HookRequestBuilder builds and executes requests for operations under \app\hook
type HookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Config the config property
// returns a *HookConfigRequestBuilder when successful
func (m *HookRequestBuilder) Config()(*HookConfigRequestBuilder) {
    return NewHookConfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewHookRequestBuilderInternal instantiates a new HookRequestBuilder and sets the default values.
func NewHookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookRequestBuilder) {
    m := &HookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/hook", pathParameters),
    }
    return m
}
// NewHookRequestBuilder instantiates a new HookRequestBuilder and sets the default values.
func NewHookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHookRequestBuilderInternal(urlParams, requestAdapter)
}
// Deliveries the deliveries property
// returns a *HookDeliveriesRequestBuilder when successful
func (m *HookRequestBuilder) Deliveries()(*HookDeliveriesRequestBuilder) {
    return NewHookDeliveriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
