package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTarballWithRefItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\tarball\{ref}
type ItemItemTarballWithRefItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemTarballWithRefItemRequestBuilderInternal instantiates a new ItemItemTarballWithRefItemRequestBuilder and sets the default values.
func NewItemItemTarballWithRefItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballWithRefItemRequestBuilder) {
    m := &ItemItemTarballWithRefItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/tarball/{ref}", pathParameters),
    }
    return m
}
// NewItemItemTarballWithRefItemRequestBuilder instantiates a new ItemItemTarballWithRefItemRequestBuilder and sets the default values.
func NewItemItemTarballWithRefItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballWithRefItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTarballWithRefItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download a tar archive for a repository. If you omit `:ref`, the repository’s default branch (usually`main`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to usethe `Location` header to make a second `GET` request.**Note**: For private repositories, these links are temporary and expire after five minutes.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#download-a-repository-archive-tar
func (m *ItemItemTarballWithRefItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation gets a redirect URL to download a tar archive for a repository. If you omit `:ref`, the repository’s default branch (usually`main`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to usethe `Location` header to make a second `GET` request.**Note**: For private repositories, these links are temporary and expire after five minutes.
// returns a *RequestInformation when successful
func (m *ItemItemTarballWithRefItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemTarballWithRefItemRequestBuilder when successful
func (m *ItemItemTarballWithRefItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemTarballWithRefItemRequestBuilder) {
    return NewItemItemTarballWithRefItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
