package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatusesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\statuses
type ItemItemStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BySha gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.statuses.item collection
// returns a *ItemItemStatusesWithShaItemRequestBuilder when successful
func (m *ItemItemStatusesRequestBuilder) BySha(sha string)(*ItemItemStatusesWithShaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sha != "" {
        urlTplParams["sha"] = sha
    }
    return NewItemItemStatusesWithShaItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemStatusesRequestBuilderInternal instantiates a new ItemItemStatusesRequestBuilder and sets the default values.
func NewItemItemStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatusesRequestBuilder) {
    m := &ItemItemStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/statuses", pathParameters),
    }
    return m
}
// NewItemItemStatusesRequestBuilder instantiates a new ItemItemStatusesRequestBuilder and sets the default values.
func NewItemItemStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
