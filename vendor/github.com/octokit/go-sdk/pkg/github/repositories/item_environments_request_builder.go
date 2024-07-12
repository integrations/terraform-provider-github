package repositories

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments
type ItemEnvironmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByEnvironment_name gets an item from the github.com/octokit/go-sdk/pkg/github/.repositories.item.environments.item collection
// returns a *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder when successful
func (m *ItemEnvironmentsRequestBuilder) ByEnvironment_name(environment_name string)(*ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if environment_name != "" {
        urlTplParams["environment_name"] = environment_name
    }
    return NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEnvironmentsRequestBuilderInternal instantiates a new ItemEnvironmentsRequestBuilder and sets the default values.
func NewItemEnvironmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsRequestBuilder) {
    m := &ItemEnvironmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}/environments", pathParameters),
    }
    return m
}
// NewItemEnvironmentsRequestBuilder instantiates a new ItemEnvironmentsRequestBuilder and sets the default values.
func NewItemEnvironmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsRequestBuilderInternal(urlParams, requestAdapter)
}
