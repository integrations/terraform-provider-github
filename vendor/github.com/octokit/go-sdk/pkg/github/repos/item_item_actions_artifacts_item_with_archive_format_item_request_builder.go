package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\artifacts\{artifact_id}\{archive_format}
type ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal instantiates a new ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    m := &ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/artifacts/{artifact_id}/{archive_format}", pathParameters),
    }
    return m
}
// NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder instantiates a new ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` inthe response header to find the URL for the download. The `:archive_format` must be `zip`.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a BasicError error when the service returns a 410 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/artifacts#download-an-artifact
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "410": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` inthe response header to find the URL for the download. The `:archive_format` must be `zip`.OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder when successful
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    return NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
