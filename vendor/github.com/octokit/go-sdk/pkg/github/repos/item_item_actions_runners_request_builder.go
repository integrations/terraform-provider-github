package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runners
type ItemItemActionsRunnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunnersRequestBuilderGetQueryParameters lists all self-hosted runners configured in a repository.Authenticated users must have admin access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
type ItemItemActionsRunnersRequestBuilderGetQueryParameters struct {
    // The name of a self-hosted runner.
    Name *string `uriparametername:"name"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByRunner_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.actions.runners.item collection
// returns a *ItemItemActionsRunnersWithRunner_ItemRequestBuilder when successful
func (m *ItemItemActionsRunnersRequestBuilder) ByRunner_id(runner_id int32)(*ItemItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["runner_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(runner_id), 10)
    return NewItemItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunnersRequestBuilderInternal instantiates a new ItemItemActionsRunnersRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRequestBuilder) {
    m := &ItemItemActionsRunnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runners{?name*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersRequestBuilder instantiates a new ItemItemActionsRunnersRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Downloads the downloads property
// returns a *ItemItemActionsRunnersDownloadsRequestBuilder when successful
func (m *ItemItemActionsRunnersRequestBuilder) Downloads()(*ItemItemActionsRunnersDownloadsRequestBuilder) {
    return NewItemItemActionsRunnersDownloadsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateJitconfig the generateJitconfig property
// returns a *ItemItemActionsRunnersGenerateJitconfigRequestBuilder when successful
func (m *ItemItemActionsRunnersRequestBuilder) GenerateJitconfig()(*ItemItemActionsRunnersGenerateJitconfigRequestBuilder) {
    return NewItemItemActionsRunnersGenerateJitconfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get lists all self-hosted runners configured in a repository.Authenticated users must have admin access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ItemItemActionsRunnersGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#list-self-hosted-runners-for-a-repository
func (m *ItemItemActionsRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunnersRequestBuilderGetQueryParameters])(ItemItemActionsRunnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersGetResponseable), nil
}
// RegistrationToken the registrationToken property
// returns a *ItemItemActionsRunnersRegistrationTokenRequestBuilder when successful
func (m *ItemItemActionsRunnersRequestBuilder) RegistrationToken()(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveToken the removeToken property
// returns a *ItemItemActionsRunnersRemoveTokenRequestBuilder when successful
func (m *ItemItemActionsRunnersRequestBuilder) RemoveToken()(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all self-hosted runners configured in a repository.Authenticated users must have admin access to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunnersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunnersRequestBuilder when successful
func (m *ItemItemActionsRunnersRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersRequestBuilder) {
    return NewItemItemActionsRunnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
