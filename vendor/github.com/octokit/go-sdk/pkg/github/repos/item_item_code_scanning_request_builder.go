package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning
type ItemItemCodeScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
// returns a *ItemItemCodeScanningAlertsRequestBuilder when successful
func (m *ItemItemCodeScanningRequestBuilder) Alerts()(*ItemItemCodeScanningAlertsRequestBuilder) {
    return NewItemItemCodeScanningAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analyses the analyses property
// returns a *ItemItemCodeScanningAnalysesRequestBuilder when successful
func (m *ItemItemCodeScanningRequestBuilder) Analyses()(*ItemItemCodeScanningAnalysesRequestBuilder) {
    return NewItemItemCodeScanningAnalysesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codeql the codeql property
// returns a *ItemItemCodeScanningCodeqlRequestBuilder when successful
func (m *ItemItemCodeScanningRequestBuilder) Codeql()(*ItemItemCodeScanningCodeqlRequestBuilder) {
    return NewItemItemCodeScanningCodeqlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningRequestBuilderInternal instantiates a new ItemItemCodeScanningRequestBuilder and sets the default values.
func NewItemItemCodeScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningRequestBuilder) {
    m := &ItemItemCodeScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningRequestBuilder instantiates a new ItemItemCodeScanningRequestBuilder and sets the default values.
func NewItemItemCodeScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultSetup the defaultSetup property
// returns a *ItemItemCodeScanningDefaultSetupRequestBuilder when successful
func (m *ItemItemCodeScanningRequestBuilder) DefaultSetup()(*ItemItemCodeScanningDefaultSetupRequestBuilder) {
    return NewItemItemCodeScanningDefaultSetupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sarifs the sarifs property
// returns a *ItemItemCodeScanningSarifsRequestBuilder when successful
func (m *ItemItemCodeScanningRequestBuilder) Sarifs()(*ItemItemCodeScanningSarifsRequestBuilder) {
    return NewItemItemCodeScanningSarifsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
