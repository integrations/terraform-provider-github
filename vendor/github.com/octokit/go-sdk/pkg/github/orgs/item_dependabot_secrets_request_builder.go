package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotSecretsRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot\secrets
type ItemDependabotSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDependabotSecretsRequestBuilderGetQueryParameters lists all secrets available in an organization without revealing theirencrypted values.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemDependabotSecretsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// BySecret_name gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.dependabot.secrets.item collection
// returns a *ItemDependabotSecretsWithSecret_nameItemRequestBuilder when successful
func (m *ItemDependabotSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemDependabotSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemDependabotSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDependabotSecretsRequestBuilderInternal instantiates a new ItemDependabotSecretsRequestBuilder and sets the default values.
func NewItemDependabotSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsRequestBuilder) {
    m := &ItemDependabotSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/dependabot/secrets{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemDependabotSecretsRequestBuilder instantiates a new ItemDependabotSecretsRequestBuilder and sets the default values.
func NewItemDependabotSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in an organization without revealing theirencrypted values.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemDependabotSecretsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/secrets#list-organization-secrets
func (m *ItemDependabotSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDependabotSecretsRequestBuilderGetQueryParameters])(ItemDependabotSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemDependabotSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDependabotSecretsGetResponseable), nil
}
// PublicKey the publicKey property
// returns a *ItemDependabotSecretsPublicKeyRequestBuilder when successful
func (m *ItemDependabotSecretsRequestBuilder) PublicKey()(*ItemDependabotSecretsPublicKeyRequestBuilder) {
    return NewItemDependabotSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in an organization without revealing theirencrypted values.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemDependabotSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDependabotSecretsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDependabotSecretsRequestBuilder when successful
func (m *ItemDependabotSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemDependabotSecretsRequestBuilder) {
    return NewItemDependabotSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
