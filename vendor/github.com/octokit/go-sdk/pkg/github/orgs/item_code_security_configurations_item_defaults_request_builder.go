package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder builds and executes requests for operations under \orgs\{org}\code-security\configurations\{configuration_id}\defaults
type ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCodeSecurityConfigurationsItemDefaultsRequestBuilderInternal instantiates a new ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsItemDefaultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder) {
    m := &ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/code-security/configurations/{configuration_id}/defaults", pathParameters),
    }
    return m
}
// NewItemCodeSecurityConfigurationsItemDefaultsRequestBuilder instantiates a new ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsItemDefaultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeSecurityConfigurationsItemDefaultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Put sets a code security configuration as a default to be applied to new repositories in your organization.This configuration will be applied to the matching repository type (all, none, public, private and internal) by default when they are created.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
// returns a ItemCodeSecurityConfigurationsItemDefaultsPutResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-security/configurations#set-a-code-security-configuration-as-a-default-for-an-organization
func (m *ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder) Put(ctx context.Context, body ItemCodeSecurityConfigurationsItemDefaultsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemCodeSecurityConfigurationsItemDefaultsPutResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCodeSecurityConfigurationsItemDefaultsPutResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCodeSecurityConfigurationsItemDefaultsPutResponseable), nil
}
// ToPutRequestInformation sets a code security configuration as a default to be applied to new repositories in your organization.This configuration will be applied to the matching repository type (all, none, public, private and internal) by default when they are created.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemCodeSecurityConfigurationsItemDefaultsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder when successful
func (m *ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder) WithUrl(rawUrl string)(*ItemCodeSecurityConfigurationsItemDefaultsRequestBuilder) {
    return NewItemCodeSecurityConfigurationsItemDefaultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
