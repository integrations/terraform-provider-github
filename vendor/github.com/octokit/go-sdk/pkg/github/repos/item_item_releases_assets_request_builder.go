package repos

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemReleasesAssetsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\releases\assets
type ItemItemReleasesAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByAsset_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.releases.assets.item collection
// returns a *ItemItemReleasesAssetsWithAsset_ItemRequestBuilder when successful
func (m *ItemItemReleasesAssetsRequestBuilder) ByAsset_id(asset_id int32)(*ItemItemReleasesAssetsWithAsset_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["asset_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(asset_id), 10)
    return NewItemItemReleasesAssetsWithAsset_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemReleasesAssetsRequestBuilderInternal instantiates a new ItemItemReleasesAssetsRequestBuilder and sets the default values.
func NewItemItemReleasesAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesAssetsRequestBuilder) {
    m := &ItemItemReleasesAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/releases/assets", pathParameters),
    }
    return m
}
// NewItemItemReleasesAssetsRequestBuilder instantiates a new ItemItemReleasesAssetsRequestBuilder and sets the default values.
func NewItemItemReleasesAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReleasesAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
