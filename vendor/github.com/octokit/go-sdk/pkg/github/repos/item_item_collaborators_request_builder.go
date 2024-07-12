package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    ief3bd7d40cba4706c2c0a30c6789d312e44d4a4ae592e9c24c47ec933ab6cd48 "github.com/octokit/go-sdk/pkg/github/repos/item/item/collaborators"
)

// ItemItemCollaboratorsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\collaborators
type ItemItemCollaboratorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCollaboratorsRequestBuilderGetQueryParameters for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.The authenticated user must have push access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` and `repo` scopes to use this endpoint.
type ItemItemCollaboratorsRequestBuilderGetQueryParameters struct {
    // Filter collaborators returned by their affiliation. `outside` means all outside collaborators of an organization-owned repository. `direct` means all collaborators with permissions to an organization-owned repository, regardless of organization membership status. `all` means all collaborators the authenticated user can see.
    Affiliation *ief3bd7d40cba4706c2c0a30c6789d312e44d4a4ae592e9c24c47ec933ab6cd48.GetAffiliationQueryParameterType `uriparametername:"affiliation"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Filter collaborators by the permissions they have on the repository. If not specified, all collaborators will be returned.
    Permission *ief3bd7d40cba4706c2c0a30c6789d312e44d4a4ae592e9c24c47ec933ab6cd48.GetPermissionQueryParameterType `uriparametername:"permission"`
}
// ByUsername gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.collaborators.item collection
// returns a *ItemItemCollaboratorsWithUsernameItemRequestBuilder when successful
func (m *ItemItemCollaboratorsRequestBuilder) ByUsername(username string)(*ItemItemCollaboratorsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemItemCollaboratorsWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCollaboratorsRequestBuilderInternal instantiates a new ItemItemCollaboratorsRequestBuilder and sets the default values.
func NewItemItemCollaboratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsRequestBuilder) {
    m := &ItemItemCollaboratorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/collaborators{?affiliation*,page*,per_page*,permission*}", pathParameters),
    }
    return m
}
// NewItemItemCollaboratorsRequestBuilder instantiates a new ItemItemCollaboratorsRequestBuilder and sets the default values.
func NewItemItemCollaboratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCollaboratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.The authenticated user must have push access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` and `repo` scopes to use this endpoint.
// returns a []Collaboratorable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/collaborators#list-repository-collaborators
func (m *ItemItemCollaboratorsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCollaboratorsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Collaboratorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCollaboratorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Collaboratorable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Collaboratorable)
        }
    }
    return val, nil
}
// ToGetRequestInformation for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.The authenticated user must have push access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` and `repo` scopes to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemCollaboratorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCollaboratorsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCollaboratorsRequestBuilder when successful
func (m *ItemItemCollaboratorsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCollaboratorsRequestBuilder) {
    return NewItemItemCollaboratorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
