package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCodespacesAccessRequestBuilder builds and executes requests for operations under \orgs\{org}\codespaces\access
type ItemCodespacesAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCodespacesAccessRequestBuilderInternal instantiates a new ItemCodespacesAccessRequestBuilder and sets the default values.
func NewItemCodespacesAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesAccessRequestBuilder) {
    m := &ItemCodespacesAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/codespaces/access", pathParameters),
    }
    return m
}
// NewItemCodespacesAccessRequestBuilder instantiates a new ItemCodespacesAccessRequestBuilder and sets the default values.
func NewItemCodespacesAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodespacesAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Put sets which users can access codespaces in an organization. This is synonymous with granting or revoking codespaces access permissions for users according to the visibility.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// Deprecated: 
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#manage-access-control-for-organization-codespaces
func (m *ItemCodespacesAccessRequestBuilder) Put(ctx context.Context, body ItemCodespacesAccessPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Selected_users the selected_users property
// returns a *ItemCodespacesAccessSelected_usersRequestBuilder when successful
func (m *ItemCodespacesAccessRequestBuilder) Selected_users()(*ItemCodespacesAccessSelected_usersRequestBuilder) {
    return NewItemCodespacesAccessSelected_usersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPutRequestInformation sets which users can access codespaces in an organization. This is synonymous with granting or revoking codespaces access permissions for users according to the visibility.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemCodespacesAccessRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemCodespacesAccessPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
// returns a *ItemCodespacesAccessRequestBuilder when successful
func (m *ItemCodespacesAccessRequestBuilder) WithUrl(rawUrl string)(*ItemCodespacesAccessRequestBuilder) {
    return NewItemCodespacesAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
