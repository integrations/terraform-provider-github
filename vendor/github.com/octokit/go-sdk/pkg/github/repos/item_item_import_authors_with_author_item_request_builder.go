package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemImportAuthorsWithAuthor_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\import\authors\{author_id}
type ItemItemImportAuthorsWithAuthor_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemImportAuthorsWithAuthor_ItemRequestBuilderInternal instantiates a new ItemItemImportAuthorsWithAuthor_ItemRequestBuilder and sets the default values.
func NewItemItemImportAuthorsWithAuthor_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) {
    m := &ItemItemImportAuthorsWithAuthor_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/import/authors/{author_id}", pathParameters),
    }
    return m
}
// NewItemItemImportAuthorsWithAuthor_ItemRequestBuilder instantiates a new ItemItemImportAuthorsWithAuthor_ItemRequestBuilder and sets the default values.
func NewItemItemImportAuthorsWithAuthor_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemImportAuthorsWithAuthor_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Patch update an author's identity for the import. Your application can continue updating authors any time before you pushnew commits to the repository.**Warning:** Due to very low levels of usage and available alternatives, this endpoint is deprecated and will no longer be available from 00:00 UTC on April 12, 2024. For more details and alternatives, see the [changelog](https://gh.io/source-imports-api-deprecation).
// Deprecated: 
// returns a PorterAuthorable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/source-imports#map-a-commit-author
func (m *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) Patch(ctx context.Context, body ItemItemImportAuthorsItemWithAuthor_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PorterAuthorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePorterAuthorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PorterAuthorable), nil
}
// ToPatchRequestInformation update an author's identity for the import. Your application can continue updating authors any time before you pushnew commits to the repository.**Warning:** Due to very low levels of usage and available alternatives, this endpoint is deprecated and will no longer be available from 00:00 UTC on April 12, 2024. For more details and alternatives, see the [changelog](https://gh.io/source-imports-api-deprecation).
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemImportAuthorsItemWithAuthor_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder when successful
func (m *ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemImportAuthorsWithAuthor_ItemRequestBuilder) {
    return NewItemItemImportAuthorsWithAuthor_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
