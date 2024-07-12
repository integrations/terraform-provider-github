package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamsItemMembershipsRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\memberships
type ItemTeamsItemMembershipsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUsername gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.teams.item.memberships.item collection
// returns a *ItemTeamsItemMembershipsWithUsernameItemRequestBuilder when successful
func (m *ItemTeamsItemMembershipsRequestBuilder) ByUsername(username string)(*ItemTeamsItemMembershipsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemTeamsItemMembershipsWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamsItemMembershipsRequestBuilderInternal instantiates a new ItemTeamsItemMembershipsRequestBuilder and sets the default values.
func NewItemTeamsItemMembershipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemMembershipsRequestBuilder) {
    m := &ItemTeamsItemMembershipsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/memberships", pathParameters),
    }
    return m
}
// NewItemTeamsItemMembershipsRequestBuilder instantiates a new ItemTeamsItemMembershipsRequestBuilder and sets the default values.
func NewItemTeamsItemMembershipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemMembershipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemMembershipsRequestBuilderInternal(urlParams, requestAdapter)
}
