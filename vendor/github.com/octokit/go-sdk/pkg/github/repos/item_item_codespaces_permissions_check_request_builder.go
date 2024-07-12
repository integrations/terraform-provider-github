package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodespacesPermissions_checkRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\codespaces\permissions_check
type ItemItemCodespacesPermissions_checkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters checks whether the permissions defined by a given devcontainer configuration have been accepted by the authenticated user.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
type ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters struct {
    // Path to the devcontainer.json configuration to use for the permission check.
    Devcontainer_path *string `uriparametername:"devcontainer_path"`
    // The git reference that points to the location of the devcontainer configuration to use for the permission check. The value of `ref` will typically be a branch name (`heads/BRANCH_NAME`). For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
    Ref *string `uriparametername:"ref"`
}
// NewItemItemCodespacesPermissions_checkRequestBuilderInternal instantiates a new ItemItemCodespacesPermissions_checkRequestBuilder and sets the default values.
func NewItemItemCodespacesPermissions_checkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    m := &ItemItemCodespacesPermissions_checkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/codespaces/permissions_check?devcontainer_path={devcontainer_path}&ref={ref}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesPermissions_checkRequestBuilder instantiates a new ItemItemCodespacesPermissions_checkRequestBuilder and sets the default values.
func NewItemItemCodespacesPermissions_checkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesPermissions_checkRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks whether the permissions defined by a given devcontainer configuration have been accepted by the authenticated user.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
// returns a CodespacesPermissionsCheckForDevcontainerable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a CodespacesPermissionsCheckForDevcontainer503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#check-if-permissions-defined-by-a-devcontainer-have-been-accepted-by-the-authenticated-user
func (m *ItemItemCodespacesPermissions_checkRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodespacesPermissionsCheckForDevcontainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespacesPermissionsCheckForDevcontainer503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespacesPermissionsCheckForDevcontainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodespacesPermissionsCheckForDevcontainerable), nil
}
// ToGetRequestInformation checks whether the permissions defined by a given devcontainer configuration have been accepted by the authenticated user.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemCodespacesPermissions_checkRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodespacesPermissions_checkRequestBuilder when successful
func (m *ItemItemCodespacesPermissions_checkRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    return NewItemItemCodespacesPermissions_checkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
