package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTrafficPopularRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\traffic\popular
type ItemItemTrafficPopularRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemTrafficPopularRequestBuilderInternal instantiates a new ItemItemTrafficPopularRequestBuilder and sets the default values.
func NewItemItemTrafficPopularRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularRequestBuilder) {
    m := &ItemItemTrafficPopularRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/traffic/popular", pathParameters),
    }
    return m
}
// NewItemItemTrafficPopularRequestBuilder instantiates a new ItemItemTrafficPopularRequestBuilder and sets the default values.
func NewItemItemTrafficPopularRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficPopularRequestBuilderInternal(urlParams, requestAdapter)
}
// Paths the paths property
// returns a *ItemItemTrafficPopularPathsRequestBuilder when successful
func (m *ItemItemTrafficPopularRequestBuilder) Paths()(*ItemItemTrafficPopularPathsRequestBuilder) {
    return NewItemItemTrafficPopularPathsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Referrers the referrers property
// returns a *ItemItemTrafficPopularReferrersRequestBuilder when successful
func (m *ItemItemTrafficPopularRequestBuilder) Referrers()(*ItemItemTrafficPopularReferrersRequestBuilder) {
    return NewItemItemTrafficPopularReferrersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
