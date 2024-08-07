package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSecurityManagersTeamsRequestBuilder builds and executes requests for operations under \orgs\{org}\security-managers\teams
type ItemSecurityManagersTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTeam_slug gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.securityManagers.teams.item collection
// returns a *ItemSecurityManagersTeamsWithTeam_slugItemRequestBuilder when successful
func (m *ItemSecurityManagersTeamsRequestBuilder) ByTeam_slug(team_slug string)(*ItemSecurityManagersTeamsWithTeam_slugItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if team_slug != "" {
        urlTplParams["team_slug"] = team_slug
    }
    return NewItemSecurityManagersTeamsWithTeam_slugItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSecurityManagersTeamsRequestBuilderInternal instantiates a new ItemSecurityManagersTeamsRequestBuilder and sets the default values.
func NewItemSecurityManagersTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersTeamsRequestBuilder) {
    m := &ItemSecurityManagersTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/security-managers/teams", pathParameters),
    }
    return m
}
// NewItemSecurityManagersTeamsRequestBuilder instantiates a new ItemSecurityManagersTeamsRequestBuilder and sets the default values.
func NewItemSecurityManagersTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityManagersTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
