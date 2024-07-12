package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesUsersRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\users
type ItemOrganizationRolesUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUsername gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.organizationRoles.users.item collection
// returns a *ItemOrganizationRolesUsersWithUsernameItemRequestBuilder when successful
func (m *ItemOrganizationRolesUsersRequestBuilder) ByUsername(username string)(*ItemOrganizationRolesUsersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemOrganizationRolesUsersWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOrganizationRolesUsersRequestBuilderInternal instantiates a new ItemOrganizationRolesUsersRequestBuilder and sets the default values.
func NewItemOrganizationRolesUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesUsersRequestBuilder) {
    m := &ItemOrganizationRolesUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/users", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesUsersRequestBuilder instantiates a new ItemOrganizationRolesUsersRequestBuilder and sets the default values.
func NewItemOrganizationRolesUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesUsersRequestBuilderInternal(urlParams, requestAdapter)
}
