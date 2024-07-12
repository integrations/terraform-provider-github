package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPropertiesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\properties
type ItemItemPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemPropertiesRequestBuilderInternal instantiates a new ItemItemPropertiesRequestBuilder and sets the default values.
func NewItemItemPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPropertiesRequestBuilder) {
    m := &ItemItemPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/properties", pathParameters),
    }
    return m
}
// NewItemItemPropertiesRequestBuilder instantiates a new ItemItemPropertiesRequestBuilder and sets the default values.
func NewItemItemPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Values the values property
// returns a *ItemItemPropertiesValuesRequestBuilder when successful
func (m *ItemItemPropertiesRequestBuilder) Values()(*ItemItemPropertiesValuesRequestBuilder) {
    return NewItemItemPropertiesValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
