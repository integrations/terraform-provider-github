package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPropertiesRequestBuilder builds and executes requests for operations under \orgs\{org}\properties
type ItemPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPropertiesRequestBuilderInternal instantiates a new ItemPropertiesRequestBuilder and sets the default values.
func NewItemPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesRequestBuilder) {
    m := &ItemPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/properties", pathParameters),
    }
    return m
}
// NewItemPropertiesRequestBuilder instantiates a new ItemPropertiesRequestBuilder and sets the default values.
func NewItemPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Schema the schema property
// returns a *ItemPropertiesSchemaRequestBuilder when successful
func (m *ItemPropertiesRequestBuilder) Schema()(*ItemPropertiesSchemaRequestBuilder) {
    return NewItemPropertiesSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Values the values property
// returns a *ItemPropertiesValuesRequestBuilder when successful
func (m *ItemPropertiesRequestBuilder) Values()(*ItemPropertiesValuesRequestBuilder) {
    return NewItemPropertiesValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
