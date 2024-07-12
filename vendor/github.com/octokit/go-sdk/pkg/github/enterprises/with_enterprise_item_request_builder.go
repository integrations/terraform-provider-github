package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithEnterpriseItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}
type WithEnterpriseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithEnterpriseItemRequestBuilderInternal instantiates a new WithEnterpriseItemRequestBuilder and sets the default values.
func NewWithEnterpriseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterpriseItemRequestBuilder) {
    m := &WithEnterpriseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}", pathParameters),
    }
    return m
}
// NewWithEnterpriseItemRequestBuilder instantiates a new WithEnterpriseItemRequestBuilder and sets the default values.
func NewWithEnterpriseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterpriseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithEnterpriseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copilot the copilot property
// returns a *ItemCopilotRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Copilot()(*ItemCopilotRequestBuilder) {
    return NewItemCopilotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dependabot the dependabot property
// returns a *ItemDependabotRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) Dependabot()(*ItemDependabotRequestBuilder) {
    return NewItemDependabotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretScanning the secretScanning property
// returns a *ItemSecretScanningRequestBuilder when successful
func (m *WithEnterpriseItemRequestBuilder) SecretScanning()(*ItemSecretScanningRequestBuilder) {
    return NewItemSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
