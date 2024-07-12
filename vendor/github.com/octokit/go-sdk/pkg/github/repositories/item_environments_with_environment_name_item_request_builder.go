package repositories

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsWithEnvironment_nameItemRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}
type ItemEnvironmentsWithEnvironment_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal instantiates a new ItemEnvironmentsWithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    m := &ItemEnvironmentsWithEnvironment_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}/environments/{environment_name}", pathParameters),
    }
    return m
}
// NewItemEnvironmentsWithEnvironment_nameItemRequestBuilder instantiates a new ItemEnvironmentsWithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemEnvironmentsWithEnvironment_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Secrets the secrets property
// returns a *ItemEnvironmentsItemSecretsRequestBuilder when successful
func (m *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Secrets()(*ItemEnvironmentsItemSecretsRequestBuilder) {
    return NewItemEnvironmentsItemSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Variables the variables property
// returns a *ItemEnvironmentsItemVariablesRequestBuilder when successful
func (m *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Variables()(*ItemEnvironmentsItemVariablesRequestBuilder) {
    return NewItemEnvironmentsItemVariablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
