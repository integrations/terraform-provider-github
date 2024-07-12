package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\security-advisories\{ghsa_id}
type ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal instantiates a new ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    m := &ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/security-advisories/{ghsa_id}", pathParameters),
    }
    return m
}
// NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder instantiates a new ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Cve the cve property
// returns a *ItemItemSecurityAdvisoriesItemCveRequestBuilder when successful
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Cve()(*ItemItemSecurityAdvisoriesItemCveRequestBuilder) {
    return NewItemItemSecurityAdvisoriesItemCveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forks the forks property
// returns a *ItemItemSecurityAdvisoriesItemForksRequestBuilder when successful
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Forks()(*ItemItemSecurityAdvisoriesItemForksRequestBuilder) {
    return NewItemItemSecurityAdvisoriesItemForksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a repository security advisory using its GitHub Security Advisory (GHSA) identifier.Anyone can access any published security advisory on a public repository.The authenticated user can access an unpublished security advisory from a repository if they are a security manager or administrator of that repository, or if they are acollaborator on the security advisory.OAuth app tokens and personal access tokens (classic) need the `repo` or `repository_advisories:read` scope to to get a published security advisory in a private repository, or any unpublished security advisory that the authenticated user has access to.
// returns a RepositoryAdvisoryable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#get-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable), nil
}
// Patch update a repository security advisory using its GitHub Security Advisory (GHSA) identifier.In order to update any security advisory, the authenticated user must be a security manager or administrator of that repository,or a collaborator on the repository security advisory.OAuth app tokens and personal access tokens (classic) need the `repo` or `repository_advisories:write` scope to use this endpoint.
// returns a RepositoryAdvisoryable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#update-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) Patch(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryUpdateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable), nil
}
// ToGetRequestInformation get a repository security advisory using its GitHub Security Advisory (GHSA) identifier.Anyone can access any published security advisory on a public repository.The authenticated user can access an unpublished security advisory from a repository if they are a security manager or administrator of that repository, or if they are acollaborator on the security advisory.OAuth app tokens and personal access tokens (classic) need the `repo` or `repository_advisories:read` scope to to get a published security advisory in a private repository, or any unpublished security advisory that the authenticated user has access to.
// returns a *RequestInformation when successful
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update a repository security advisory using its GitHub Security Advisory (GHSA) identifier.In order to update any security advisory, the authenticated user must be a security manager or administrator of that repository,or a collaborator on the repository security advisory.OAuth app tokens and personal access tokens (classic) need the `repo` or `repository_advisories:write` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryUpdateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder when successful
func (m *ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    return NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
