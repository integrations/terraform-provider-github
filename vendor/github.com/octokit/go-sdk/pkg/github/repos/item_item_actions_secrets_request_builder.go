package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsSecretsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\secrets
type ItemItemActionsSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsSecretsRequestBuilderGetQueryParameters lists all secrets available in a repository without revealing their encryptedvalues.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
type ItemItemActionsSecretsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// BySecret_name gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.actions.secrets.item collection
// returns a *ItemItemActionsSecretsWithSecret_nameItemRequestBuilder when successful
func (m *ItemItemActionsSecretsRequestBuilder) BySecret_name(secret_name string)(*ItemItemActionsSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewItemItemActionsSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsSecretsRequestBuilderInternal instantiates a new ItemItemActionsSecretsRequestBuilder and sets the default values.
func NewItemItemActionsSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsSecretsRequestBuilder) {
    m := &ItemItemActionsSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/secrets{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsSecretsRequestBuilder instantiates a new ItemItemActionsSecretsRequestBuilder and sets the default values.
func NewItemItemActionsSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available in a repository without revealing their encryptedvalues.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ItemItemActionsSecretsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#list-repository-secrets
func (m *ItemItemActionsSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsSecretsRequestBuilderGetQueryParameters])(ItemItemActionsSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsSecretsGetResponseable), nil
}
// PublicKey the publicKey property
// returns a *ItemItemActionsSecretsPublicKeyRequestBuilder when successful
func (m *ItemItemActionsSecretsRequestBuilder) PublicKey()(*ItemItemActionsSecretsPublicKeyRequestBuilder) {
    return NewItemItemActionsSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all secrets available in a repository without revealing their encryptedvalues.Authenticated users must have collaborator access to a repository to create, update, or read secrets.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsSecretsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsSecretsRequestBuilder when successful
func (m *ItemItemActionsSecretsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsSecretsRequestBuilder) {
    return NewItemItemActionsSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
