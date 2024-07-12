package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsCacheRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\cache
type ItemActionsCacheRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsCacheRequestBuilderInternal instantiates a new ItemActionsCacheRequestBuilder and sets the default values.
func NewItemActionsCacheRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheRequestBuilder) {
    m := &ItemActionsCacheRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/cache", pathParameters),
    }
    return m
}
// NewItemActionsCacheRequestBuilder instantiates a new ItemActionsCacheRequestBuilder and sets the default values.
func NewItemActionsCacheRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsCacheRequestBuilderInternal(urlParams, requestAdapter)
}
// Usage the usage property
// returns a *ItemActionsCacheUsageRequestBuilder when successful
func (m *ItemActionsCacheRequestBuilder) Usage()(*ItemActionsCacheUsageRequestBuilder) {
    return NewItemActionsCacheUsageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsageByRepository the usageByRepository property
// returns a *ItemActionsCacheUsageByRepositoryRequestBuilder when successful
func (m *ItemActionsCacheRequestBuilder) UsageByRepository()(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    return NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
