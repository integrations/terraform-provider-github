package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\environments\{environment_name}\secrets\public-key
type ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal instantiates a new ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    m := &ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/environments/{environment_name}/secrets/public-key", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder instantiates a new ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the public key for an environment, which you need to encrypt environmentsecrets. You need to encrypt a secret before you can create or update secrets.Anyone with read access to the repository can use this endpoint.If the repository is private, OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ActionsPublicKeyable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#get-an-environment-public-key
func (m *ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ActionsPublicKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateActionsPublicKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ActionsPublicKeyable), nil
}
// ToGetRequestInformation get the public key for an environment, which you need to encrypt environmentsecrets. You need to encrypt a secret before you can create or update secrets.Anyone with read access to the repository can use this endpoint.If the repository is private, OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder when successful
func (m *ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    return NewItemItemEnvironmentsItemSecretsPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
