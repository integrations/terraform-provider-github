package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesTeamsRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\teams
type ItemOrganizationRolesTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTeam_slug gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.organizationRoles.teams.item collection
// returns a *ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder when successful
func (m *ItemOrganizationRolesTeamsRequestBuilder) ByTeam_slug(team_slug string)(*ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if team_slug != "" {
        urlTplParams["team_slug"] = team_slug
    }
    return NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOrganizationRolesTeamsRequestBuilderInternal instantiates a new ItemOrganizationRolesTeamsRequestBuilder and sets the default values.
func NewItemOrganizationRolesTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesTeamsRequestBuilder) {
    m := &ItemOrganizationRolesTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/teams", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesTeamsRequestBuilder instantiates a new ItemOrganizationRolesTeamsRequestBuilder and sets the default values.
func NewItemOrganizationRolesTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
