package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemSsh_signing_keysRequestBuilder builds and executes requests for operations under \users\{username}\ssh_signing_keys
type ItemSsh_signing_keysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSsh_signing_keysRequestBuilderGetQueryParameters lists the SSH signing keys for a user. This operation is accessible by anyone.
type ItemSsh_signing_keysRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemSsh_signing_keysRequestBuilderInternal instantiates a new ItemSsh_signing_keysRequestBuilder and sets the default values.
func NewItemSsh_signing_keysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSsh_signing_keysRequestBuilder) {
    m := &ItemSsh_signing_keysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/ssh_signing_keys{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemSsh_signing_keysRequestBuilder instantiates a new ItemSsh_signing_keysRequestBuilder and sets the default values.
func NewItemSsh_signing_keysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSsh_signing_keysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSsh_signing_keysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the SSH signing keys for a user. This operation is accessible by anyone.
// returns a []SshSigningKeyable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/ssh-signing-keys#list-ssh-signing-keys-for-a-user
func (m *ItemSsh_signing_keysRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSsh_signing_keysRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SshSigningKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSshSigningKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SshSigningKeyable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SshSigningKeyable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the SSH signing keys for a user. This operation is accessible by anyone.
// returns a *RequestInformation when successful
func (m *ItemSsh_signing_keysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSsh_signing_keysRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSsh_signing_keysRequestBuilder when successful
func (m *ItemSsh_signing_keysRequestBuilder) WithUrl(rawUrl string)(*ItemSsh_signing_keysRequestBuilder) {
    return NewItemSsh_signing_keysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
