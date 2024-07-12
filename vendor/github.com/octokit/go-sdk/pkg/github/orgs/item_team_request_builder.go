package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamRequestBuilder builds and executes requests for operations under \orgs\{org}\team
type ItemTeamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTeam_slug gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.team.item collection
// returns a *ItemTeamWithTeam_slugItemRequestBuilder when successful
func (m *ItemTeamRequestBuilder) ByTeam_slug(team_slug string)(*ItemTeamWithTeam_slugItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if team_slug != "" {
        urlTplParams["team_slug"] = team_slug
    }
    return NewItemTeamWithTeam_slugItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamRequestBuilderInternal instantiates a new ItemTeamRequestBuilder and sets the default values.
func NewItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamRequestBuilder) {
    m := &ItemTeamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/team", pathParameters),
    }
    return m
}
// NewItemTeamRequestBuilder instantiates a new ItemTeamRequestBuilder and sets the default values.
func NewItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}
