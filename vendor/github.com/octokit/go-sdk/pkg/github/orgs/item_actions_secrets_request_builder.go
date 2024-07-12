package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsSecretsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\secrets
type ItemActionsSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsSecretsRequestBuilderGetQueryParameters lists all secrets available in an organization without revealing theirencrypted values.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
type ItemActionsSecretsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// BySecret_name gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.actions.secrets.item collection
// returns a *ItemActionsSecretsWithSecret_nameItemRequestBuilder when successful
func (m *ItemActionsSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemActionsSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemActionsSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsSecretsRequestBuilderInternal instantiates a new ItemActionsSecretsRequestBuilder and sets the default values.
func NewItemActionsSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsRequestBuilder) {
    m := &ItemActionsSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/secrets{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsSecretsRequestBuilder instantiates a new ItemActionsSecretsRequestBuilder and sets the default values.
func NewItemActionsSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in an organization without revealing theirencrypted values.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
// returns a ItemActionsSecretsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-organization-secrets
func (m *ItemActionsSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsSecretsRequestBuilderGetQueryParameters])(ItemActionsSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsSecretsGetResponseable), nil
}
// PublicKey the publicKey property
// returns a *ItemActionsSecretsPublicKeyRequestBuilder when successful
func (m *ItemActionsSecretsRequestBuilder) PublicKey()(*ItemActionsSecretsPublicKeyRequestBuilder) {
    return NewItemActionsSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in an organization without revealing theirencrypted values.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
// returns a *RequestInformation when successful
func (m *ItemActionsSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsSecretsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsSecretsRequestBuilder when successful
func (m *ItemActionsSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsSecretsRequestBuilder) {
    return NewItemActionsSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
