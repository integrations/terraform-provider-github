package repositories

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithRepository_ItemRequestBuilder builds and executes requests for operations under \repositories\{repository_id}
type WithRepository_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithRepository_ItemRequestBuilderInternal instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewWithRepository_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithRepository_ItemRequestBuilder) {
    m := &WithRepository_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}", pathParameters),
    }
    return m
}
// NewWithRepository_ItemRequestBuilder instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewWithRepository_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithRepository_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithRepository_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Environments the environments property
// returns a *ItemEnvironmentsRequestBuilder when successful
func (m *WithRepository_ItemRequestBuilder) Environments()(*ItemEnvironmentsRequestBuilder) {
    return NewItemEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
