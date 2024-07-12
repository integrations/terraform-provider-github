package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemSecretScanningRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\secret-scanning
type ItemItemSecretScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
// returns a *ItemItemSecretScanningAlertsRequestBuilder when successful
func (m *ItemItemSecretScanningRequestBuilder) Alerts()(*ItemItemSecretScanningAlertsRequestBuilder) {
    return NewItemItemSecretScanningAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemSecretScanningRequestBuilderInternal instantiates a new ItemItemSecretScanningRequestBuilder and sets the default values.
func NewItemItemSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningRequestBuilder) {
    m := &ItemItemSecretScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/secret-scanning", pathParameters),
    }
    return m
}
// NewItemItemSecretScanningRequestBuilder instantiates a new ItemItemSecretScanningRequestBuilder and sets the default values.
func NewItemItemSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
