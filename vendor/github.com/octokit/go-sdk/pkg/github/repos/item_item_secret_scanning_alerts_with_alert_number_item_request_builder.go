package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\secret-scanning\alerts\{alert_number}
type ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal instantiates a new ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    m := &ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/secret-scanning/alerts/{alert_number}", pathParameters),
    }
    return m
}
// NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder instantiates a new ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a single secret scanning alert detected in an eligible repository.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a SecretScanningAlertable when successful
// returns a SecretScanningAlert503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/secret-scanning/secret-scanning#get-a-secret-scanning-alert
func (m *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSecretScanningAlert503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSecretScanningAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningAlertable), nil
}
// Locations the locations property
// returns a *ItemItemSecretScanningAlertsItemLocationsRequestBuilder when successful
func (m *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) Locations()(*ItemItemSecretScanningAlertsItemLocationsRequestBuilder) {
    return NewItemItemSecretScanningAlertsItemLocationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch updates the status of a secret scanning alert in an eligible repository.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a SecretScanningAlertable when successful
// returns a SecretScanningAlert503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/secret-scanning/secret-scanning#update-a-secret-scanning-alert
func (m *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) Patch(ctx context.Context, body ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningAlertable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSecretScanningAlert503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSecretScanningAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningAlertable), nil
}
// ToGetRequestInformation gets a single secret scanning alert detected in an eligible repository.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a *RequestInformation when successful
func (m *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates the status of a secret scanning alert in an eligible repository.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a *RequestInformation when successful
func (m *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder when successful
func (m *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    return NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
