package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsVariablesItemRepositoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\variables\{name}\repositories
type ItemActionsVariablesItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsVariablesItemRepositoriesRequestBuilderGetQueryParameters lists all repositories that can access an organization variablethat is available to selected repositories.Authenticated users must have collaborator access to a repository to create, update, or read variables.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
type ItemActionsVariablesItemRepositoriesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByRepository_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.actions.variables.item.repositories.item collection
// returns a *ItemActionsVariablesItemRepositoriesWithRepository_ItemRequestBuilder when successful
func (m *ItemActionsVariablesItemRepositoriesRequestBuilder) ByRepository_id(repository_id int32)(*ItemActionsVariablesItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["repository_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(repository_id), 10)
    return NewItemActionsVariablesItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsVariablesItemRepositoriesRequestBuilderInternal instantiates a new ItemActionsVariablesItemRepositoriesRequestBuilder and sets the default values.
func NewItemActionsVariablesItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsVariablesItemRepositoriesRequestBuilder) {
    m := &ItemActionsVariablesItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/variables/{name}/repositories{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsVariablesItemRepositoriesRequestBuilder instantiates a new ItemActionsVariablesItemRepositoriesRequestBuilder and sets the default values.
func NewItemActionsVariablesItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsVariablesItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsVariablesItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all repositories that can access an organization variablethat is available to selected repositories.Authenticated users must have collaborator access to a repository to create, update, or read variables.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
// returns a ItemActionsVariablesItemRepositoriesGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-selected-repositories-for-an-organization-variable
func (m *ItemActionsVariablesItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsVariablesItemRepositoriesRequestBuilderGetQueryParameters])(ItemActionsVariablesItemRepositoriesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsVariablesItemRepositoriesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsVariablesItemRepositoriesGetResponseable), nil
}
// Put replaces all repositories for an organization variable that is availableto selected repositories. Organization variables that are available to selectedrepositories have their `visibility` field set to `selected`.Authenticated users must have collaborator access to a repository to create, update, or read variables.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#set-selected-repositories-for-an-organization-variable
func (m *ItemActionsVariablesItemRepositoriesRequestBuilder) Put(ctx context.Context, body ItemActionsVariablesItemRepositoriesPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation lists all repositories that can access an organization variablethat is available to selected repositories.Authenticated users must have collaborator access to a repository to create, update, or read variables.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
// returns a *RequestInformation when successful
func (m *ItemActionsVariablesItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsVariablesItemRepositoriesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation replaces all repositories for an organization variable that is availableto selected repositories. Organization variables that are available to selectedrepositories have their `visibility` field set to `selected`.Authenticated users must have collaborator access to a repository to create, update, or read variables.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint. If the repository is private, the `repo` scope is also required.
// returns a *RequestInformation when successful
func (m *ItemActionsVariablesItemRepositoriesRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemActionsVariablesItemRepositoriesPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsVariablesItemRepositoriesRequestBuilder when successful
func (m *ItemActionsVariablesItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*ItemActionsVariablesItemRepositoriesRequestBuilder) {
    return NewItemActionsVariablesItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
