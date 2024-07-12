package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\teams\{team_slug}
type ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByRole_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.organizationRoles.teams.item.item collection
// returns a *ItemOrganizationRolesTeamsItemWithRole_ItemRequestBuilder when successful
func (m *ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) ByRole_id(role_id int32)(*ItemOrganizationRolesTeamsItemWithRole_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["role_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(role_id), 10)
    return NewItemOrganizationRolesTeamsItemWithRole_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilderInternal instantiates a new ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder and sets the default values.
func NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) {
    m := &ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/teams/{team_slug}", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder instantiates a new ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder and sets the default values.
func NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes all assigned organization roles from a team. For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."The authenticated user must be an administrator for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#remove-all-organization-roles-for-a-team
func (m *ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation removes all assigned organization roles from a team. For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."The authenticated user must be an administrator for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder when successful
func (m *ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder) {
    return NewItemOrganizationRolesTeamsWithTeam_slugItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
