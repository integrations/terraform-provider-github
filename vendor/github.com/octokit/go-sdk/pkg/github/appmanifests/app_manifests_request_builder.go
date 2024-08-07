package appmanifests

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppManifestsRequestBuilder builds and executes requests for operations under \app-manifests
type AppManifestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCode gets an item from the github.com/octokit/go-sdk/pkg/github.appManifests.item collection
// returns a *WithCodeItemRequestBuilder when successful
func (m *AppManifestsRequestBuilder) ByCode(code string)(*WithCodeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if code != "" {
        urlTplParams["code"] = code
    }
    return NewWithCodeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppManifestsRequestBuilderInternal instantiates a new AppManifestsRequestBuilder and sets the default values.
func NewAppManifestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManifestsRequestBuilder) {
    m := &AppManifestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app-manifests", pathParameters),
    }
    return m
}
// NewAppManifestsRequestBuilder instantiates a new AppManifestsRequestBuilder and sets the default values.
func NewAppManifestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManifestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppManifestsRequestBuilderInternal(urlParams, requestAdapter)
}
