package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\code-security\configurations\{configuration_id}\repositories
type ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderGetQueryParameters lists the repositories associated with a code security configuration in an organization.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
type ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Before *string `uriparametername:"before"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // A comma-separated list of statuses. If specified, only repositories with these attachment statuses will be returned.Can be: `all`, `attached`, `attaching`, `detached`, `enforced`, `failed`, `updating`
    Status *string `uriparametername:"status"`
}
// NewItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderInternal instantiates a new ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder) {
    m := &ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/code-security/configurations/{configuration_id}/repositories{?after*,before*,per_page*,status*}", pathParameters),
    }
    return m
}
// NewItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder instantiates a new ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder and sets the default values.
func NewItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the repositories associated with a code security configuration in an organization.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
// returns a []CodeSecurityConfigurationRepositoriesable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-security/configurations#get-repositories-associated-with-a-code-security-configuration
func (m *ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeSecurityConfigurationRepositoriesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeSecurityConfigurationRepositoriesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeSecurityConfigurationRepositoriesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeSecurityConfigurationRepositoriesable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the repositories associated with a code security configuration in an organization.The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `write:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder when successful
func (m *ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*ItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder) {
    return NewItemCodeSecurityConfigurationsItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
