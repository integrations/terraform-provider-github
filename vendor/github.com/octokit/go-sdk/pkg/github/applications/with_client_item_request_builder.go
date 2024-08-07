package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithClient_ItemRequestBuilder builds and executes requests for operations under \applications\{client_id}
type WithClient_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithClient_ItemRequestBuilderInternal instantiates a new WithClient_ItemRequestBuilder and sets the default values.
func NewWithClient_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithClient_ItemRequestBuilder) {
    m := &WithClient_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{client_id}", pathParameters),
    }
    return m
}
// NewWithClient_ItemRequestBuilder instantiates a new WithClient_ItemRequestBuilder and sets the default values.
func NewWithClient_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithClient_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithClient_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Grant the grant property
// returns a *ItemGrantRequestBuilder when successful
func (m *WithClient_ItemRequestBuilder) Grant()(*ItemGrantRequestBuilder) {
    return NewItemGrantRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Token the token property
// returns a *ItemTokenRequestBuilder when successful
func (m *WithClient_ItemRequestBuilder) Token()(*ItemTokenRequestBuilder) {
    return NewItemTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
