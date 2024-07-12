package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\environments
type ItemItemEnvironmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsRequestBuilderGetQueryParameters lists the environments for a repository.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
type ItemItemEnvironmentsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByEnvironment_name gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.environments.item collection
// returns a *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder when successful
func (m *ItemItemEnvironmentsRequestBuilder) ByEnvironment_name(environment_name string)(*ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if environment_name != "" {
        urlTplParams["environment_name"] = environment_name
    }
    return NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemEnvironmentsRequestBuilderInternal instantiates a new ItemItemEnvironmentsRequestBuilder and sets the default values.
func NewItemItemEnvironmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsRequestBuilder) {
    m := &ItemItemEnvironmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/environments{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsRequestBuilder instantiates a new ItemItemEnvironmentsRequestBuilder and sets the default values.
func NewItemItemEnvironmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the environments for a repository.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a ItemItemEnvironmentsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#list-environments
func (m *ItemItemEnvironmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemEnvironmentsRequestBuilderGetQueryParameters])(ItemItemEnvironmentsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsGetResponseable), nil
}
// ToGetRequestInformation lists the environments for a repository.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemEnvironmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemEnvironmentsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemEnvironmentsRequestBuilder when successful
func (m *ItemItemEnvironmentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsRequestBuilder) {
    return NewItemItemEnvironmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
