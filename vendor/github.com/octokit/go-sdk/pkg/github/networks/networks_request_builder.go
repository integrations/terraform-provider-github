package networks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// NetworksRequestBuilder builds and executes requests for operations under \networks
type NetworksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByOwner gets an item from the github.com/octokit/go-sdk/pkg/github.networks.item collection
// returns a *WithOwnerItemRequestBuilder when successful
func (m *NetworksRequestBuilder) ByOwner(owner string)(*WithOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if owner != "" {
        urlTplParams["owner"] = owner
    }
    return NewWithOwnerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewNetworksRequestBuilderInternal instantiates a new NetworksRequestBuilder and sets the default values.
func NewNetworksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NetworksRequestBuilder) {
    m := &NetworksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networks", pathParameters),
    }
    return m
}
// NewNetworksRequestBuilder instantiates a new NetworksRequestBuilder and sets the default values.
func NewNetworksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NetworksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNetworksRequestBuilderInternal(urlParams, requestAdapter)
}
