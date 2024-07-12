package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCollaboratorsItemPermissionRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\collaborators\{username}\permission
type ItemItemCollaboratorsItemPermissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCollaboratorsItemPermissionRequestBuilderInternal instantiates a new ItemItemCollaboratorsItemPermissionRequestBuilder and sets the default values.
func NewItemItemCollaboratorsItemPermissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    m := &ItemItemCollaboratorsItemPermissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/collaborators/{username}/permission", pathParameters),
    }
    return m
}
// NewItemItemCollaboratorsItemPermissionRequestBuilder instantiates a new ItemItemCollaboratorsItemPermissionRequestBuilder and sets the default values.
func NewItemItemCollaboratorsItemPermissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCollaboratorsItemPermissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks the repository permission of a collaborator. The possible repositorypermissions are `admin`, `write`, `read`, and `none`.*Note*: The `permission` attribute provides the legacy base roles of `admin`, `write`, `read`, and `none`, where the`maintain` role is mapped to `write` and the `triage` role is mapped to `read`. To determine the role assigned to thecollaborator, see the `role_name` attribute, which will provide the full role name, including custom roles. The`permissions` hash can also be used to determine which base level of access the collaborator has to the repository.
// returns a RepositoryCollaboratorPermissionable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/collaborators#get-repository-permissions-for-a-user
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryCollaboratorPermissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRepositoryCollaboratorPermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryCollaboratorPermissionable), nil
}
// ToGetRequestInformation checks the repository permission of a collaborator. The possible repositorypermissions are `admin`, `write`, `read`, and `none`.*Note*: The `permission` attribute provides the legacy base roles of `admin`, `write`, `read`, and `none`, where the`maintain` role is mapped to `write` and the `triage` role is mapped to `read`. To determine the role assigned to thecollaborator, see the `role_name` attribute, which will provide the full role name, including custom roles. The`permissions` hash can also be used to determine which base level of access the collaborator has to the repository.
// returns a *RequestInformation when successful
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCollaboratorsItemPermissionRequestBuilder when successful
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) WithUrl(rawUrl string)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    return NewItemItemCollaboratorsItemPermissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
