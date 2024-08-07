package networks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithRepoItemRequestBuilder builds and executes requests for operations under \networks\{owner}\{repo}
type ItemWithRepoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemWithRepoItemRequestBuilderInternal instantiates a new ItemWithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithRepoItemRequestBuilder) {
    m := &ItemWithRepoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networks/{owner}/{repo}", pathParameters),
    }
    return m
}
// NewItemWithRepoItemRequestBuilder instantiates a new ItemWithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithRepoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithRepoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
// returns a *ItemItemEventsRequestBuilder when successful
func (m *ItemWithRepoItemRequestBuilder) Events()(*ItemItemEventsRequestBuilder) {
    return NewItemItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
