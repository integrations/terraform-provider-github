package repositories

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsItemSecretsRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}\secrets
type ItemEnvironmentsItemSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters lists all secrets available in an environment without revealing theirencrypted values.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
type ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// BySecret_name gets an item from the github.com/octokit/go-sdk/pkg/github/.repositories.item.environments.item.secrets.item collection
// returns a *ItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilder when successful
func (m *ItemEnvironmentsItemSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEnvironmentsItemSecretsRequestBuilderInternal instantiates a new ItemEnvironmentsItemSecretsRequestBuilder and sets the default values.
func NewItemEnvironmentsItemSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemSecretsRequestBuilder) {
    m := &ItemEnvironmentsItemSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}/environments/{environment_name}/secrets{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemEnvironmentsItemSecretsRequestBuilder instantiates a new ItemEnvironmentsItemSecretsRequestBuilder and sets the default values.
func NewItemEnvironmentsItemSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsItemSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in an environment without revealing theirencrypted values.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ItemEnvironmentsItemSecretsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-environment-secrets
func (m *ItemEnvironmentsItemSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters])(ItemEnvironmentsItemSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemEnvironmentsItemSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEnvironmentsItemSecretsGetResponseable), nil
}
// PublicKey the publicKey property
// returns a *ItemEnvironmentsItemSecretsPublicKeyRequestBuilder when successful
func (m *ItemEnvironmentsItemSecretsRequestBuilder) PublicKey()(*ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    return NewItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in an environment without revealing theirencrypted values.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemEnvironmentsItemSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemEnvironmentsItemSecretsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemEnvironmentsItemSecretsRequestBuilder when successful
func (m *ItemEnvironmentsItemSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemEnvironmentsItemSecretsRequestBuilder) {
    return NewItemEnvironmentsItemSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
