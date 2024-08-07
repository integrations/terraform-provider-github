package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEventsOrgsRequestBuilder builds and executes requests for operations under \users\{username}\events\orgs
type ItemEventsOrgsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByOrg gets an item from the github.com/octokit/go-sdk/pkg/github.users.item.events.orgs.item collection
// returns a *ItemEventsOrgsWithOrgItemRequestBuilder when successful
func (m *ItemEventsOrgsRequestBuilder) ByOrg(org string)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if org != "" {
        urlTplParams["org"] = org
    }
    return NewItemEventsOrgsWithOrgItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEventsOrgsRequestBuilderInternal instantiates a new ItemEventsOrgsRequestBuilder and sets the default values.
func NewItemEventsOrgsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsRequestBuilder) {
    m := &ItemEventsOrgsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/events/orgs", pathParameters),
    }
    return m
}
// NewItemEventsOrgsRequestBuilder instantiates a new ItemEventsOrgsRequestBuilder and sets the default values.
func NewItemEventsOrgsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsOrgsRequestBuilderInternal(urlParams, requestAdapter)
}
