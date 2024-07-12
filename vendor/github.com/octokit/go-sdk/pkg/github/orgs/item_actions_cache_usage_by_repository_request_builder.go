package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsCacheUsageByRepositoryRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\cache\usage-by-repository
type ItemActionsCacheUsageByRepositoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.OAuth tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
type ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemActionsCacheUsageByRepositoryRequestBuilderInternal instantiates a new ItemActionsCacheUsageByRepositoryRequestBuilder and sets the default values.
func NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    m := &ItemActionsCacheUsageByRepositoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/cache/usage-by-repository{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsCacheUsageByRepositoryRequestBuilder instantiates a new ItemActionsCacheUsageByRepositoryRequestBuilder and sets the default values.
func NewItemActionsCacheUsageByRepositoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.OAuth tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
// returns a ItemActionsCacheUsageByRepositoryGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/cache#list-repositories-with-github-actions-cache-usage-for-an-organization
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters])(ItemActionsCacheUsageByRepositoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsCacheUsageByRepositoryGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsCacheUsageByRepositoryGetResponseable), nil
}
// ToGetRequestInformation lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.OAuth tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsCacheUsageByRepositoryRequestBuilder when successful
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) WithUrl(rawUrl string)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    return NewItemActionsCacheUsageByRepositoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
