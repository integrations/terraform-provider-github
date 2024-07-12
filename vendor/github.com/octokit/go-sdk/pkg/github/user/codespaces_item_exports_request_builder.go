package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// CodespacesItemExportsRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\exports
type CodespacesItemExportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByExport_id gets an item from the github.com/octokit/go-sdk/pkg/github.user.codespaces.item.exports.item collection
// returns a *CodespacesItemExportsWithExport_ItemRequestBuilder when successful
func (m *CodespacesItemExportsRequestBuilder) ByExport_id(export_id string)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if export_id != "" {
        urlTplParams["export_id"] = export_id
    }
    return NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCodespacesItemExportsRequestBuilderInternal instantiates a new CodespacesItemExportsRequestBuilder and sets the default values.
func NewCodespacesItemExportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsRequestBuilder) {
    m := &CodespacesItemExportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}/exports", pathParameters),
    }
    return m
}
// NewCodespacesItemExportsRequestBuilder instantiates a new CodespacesItemExportsRequestBuilder and sets the default values.
func NewCodespacesItemExportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemExportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post triggers an export of the specified codespace and returns a URL and ID where the status of the export can be monitored.If changes cannot be pushed to the codespace's repository, they will be pushed to a new or previously-existing fork instead.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
// returns a CodespaceExportDetailsable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#export-a-codespace-for-the-authenticated-user
func (m *CodespacesItemExportsRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodespaceExportDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespaceExportDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodespaceExportDetailsable), nil
}
// ToPostRequestInformation triggers an export of the specified codespace and returns a URL and ID where the status of the export can be monitored.If changes cannot be pushed to the codespace's repository, they will be pushed to a new or previously-existing fork instead.OAuth app tokens and personal access tokens (classic) need the `codespace` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *CodespacesItemExportsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CodespacesItemExportsRequestBuilder when successful
func (m *CodespacesItemExportsRequestBuilder) WithUrl(rawUrl string)(*CodespacesItemExportsRequestBuilder) {
    return NewCodespacesItemExportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
