package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSecretScanningRequestBuilder builds and executes requests for operations under \orgs\{org}\secret-scanning
type ItemSecretScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
// returns a *ItemSecretScanningAlertsRequestBuilder when successful
func (m *ItemSecretScanningRequestBuilder) Alerts()(*ItemSecretScanningAlertsRequestBuilder) {
    return NewItemSecretScanningAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSecretScanningRequestBuilderInternal instantiates a new ItemSecretScanningRequestBuilder and sets the default values.
func NewItemSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecretScanningRequestBuilder) {
    m := &ItemSecretScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/secret-scanning", pathParameters),
    }
    return m
}
// NewItemSecretScanningRequestBuilder instantiates a new ItemSecretScanningRequestBuilder and sets the default values.
func NewItemSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
