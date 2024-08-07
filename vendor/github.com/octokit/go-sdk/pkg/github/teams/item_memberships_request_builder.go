package teams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembershipsRequestBuilder builds and executes requests for operations under \teams\{team_id}\memberships
type ItemMembershipsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUsername gets an item from the github.com/octokit/go-sdk/pkg/github.teams.item.memberships.item collection
// Deprecated: 
// returns a *ItemMembershipsWithUsernameItemRequestBuilder when successful
func (m *ItemMembershipsRequestBuilder) ByUsername(username string)(*ItemMembershipsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemMembershipsWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMembershipsRequestBuilderInternal instantiates a new ItemMembershipsRequestBuilder and sets the default values.
func NewItemMembershipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembershipsRequestBuilder) {
    m := &ItemMembershipsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/memberships", pathParameters),
    }
    return m
}
// NewItemMembershipsRequestBuilder instantiates a new ItemMembershipsRequestBuilder and sets the default values.
func NewItemMembershipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembershipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembershipsRequestBuilderInternal(urlParams, requestAdapter)
}
