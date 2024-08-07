package appmanifests

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithCodeItemRequestBuilder builds and executes requests for operations under \app-manifests\{code}
type WithCodeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithCodeItemRequestBuilderInternal instantiates a new WithCodeItemRequestBuilder and sets the default values.
func NewWithCodeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCodeItemRequestBuilder) {
    m := &WithCodeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app-manifests/{code}", pathParameters),
    }
    return m
}
// NewWithCodeItemRequestBuilder instantiates a new WithCodeItemRequestBuilder and sets the default values.
func NewWithCodeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCodeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithCodeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Conversions the conversions property
// returns a *ItemConversionsRequestBuilder when successful
func (m *WithCodeItemRequestBuilder) Conversions()(*ItemConversionsRequestBuilder) {
    return NewItemConversionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
