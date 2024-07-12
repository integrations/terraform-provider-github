package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodespacesMachinesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\codespaces\machines
type ItemItemCodespacesMachinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesMachinesRequestBuilderGetQueryParameters list the machine types available for a given repository based on its configuration.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
type ItemItemCodespacesMachinesRequestBuilderGetQueryParameters struct {
    // IP for location auto-detection when proxying a request
    Client_ip *string `uriparametername:"client_ip"`
    // The location to check for available machines. Assigned by IP if not provided.
    Location *string `uriparametername:"location"`
    // The branch or commit to check for prebuild availability and devcontainer restrictions.
    Ref *string `uriparametername:"ref"`
}
// NewItemItemCodespacesMachinesRequestBuilderInternal instantiates a new ItemItemCodespacesMachinesRequestBuilder and sets the default values.
func NewItemItemCodespacesMachinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesMachinesRequestBuilder) {
    m := &ItemItemCodespacesMachinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/codespaces/machines{?client_ip*,location*,ref*}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesMachinesRequestBuilder instantiates a new ItemItemCodespacesMachinesRequestBuilder and sets the default values.
func NewItemItemCodespacesMachinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesMachinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesMachinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the machine types available for a given repository based on its configuration.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
// returns a ItemItemCodespacesMachinesGetResponseable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/machines#list-available-machine-types-for-a-repository
func (m *ItemItemCodespacesMachinesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodespacesMachinesRequestBuilderGetQueryParameters])(ItemItemCodespacesMachinesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesMachinesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesMachinesGetResponseable), nil
}
// ToGetRequestInformation list the machine types available for a given repository based on its configuration.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemCodespacesMachinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodespacesMachinesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodespacesMachinesRequestBuilder when successful
func (m *ItemItemCodespacesMachinesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesMachinesRequestBuilder) {
    return NewItemItemCodespacesMachinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
