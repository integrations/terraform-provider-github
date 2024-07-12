package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInstallationsRequestBuilder builds and executes requests for operations under \orgs\{org}\installations
type ItemInstallationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstallationsRequestBuilderGetQueryParameters lists all GitHub Apps in an organization. The installation count includesall GitHub Apps installed on repositories in the organization.The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:read` scope to use this endpoint.
type ItemInstallationsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemInstallationsRequestBuilderInternal instantiates a new ItemInstallationsRequestBuilder and sets the default values.
func NewItemInstallationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstallationsRequestBuilder) {
    m := &ItemInstallationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/installations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemInstallationsRequestBuilder instantiates a new ItemInstallationsRequestBuilder and sets the default values.
func NewItemInstallationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstallationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstallationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all GitHub Apps in an organization. The installation count includesall GitHub Apps installed on repositories in the organization.The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:read` scope to use this endpoint.
// returns a ItemInstallationsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/orgs#list-app-installations-for-an-organization
func (m *ItemInstallationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInstallationsRequestBuilderGetQueryParameters])(ItemInstallationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInstallationsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInstallationsGetResponseable), nil
}
// ToGetRequestInformation lists all GitHub Apps in an organization. The installation count includesall GitHub Apps installed on repositories in the organization.The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:read` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemInstallationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInstallationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInstallationsRequestBuilder when successful
func (m *ItemInstallationsRequestBuilder) WithUrl(rawUrl string)(*ItemInstallationsRequestBuilder) {
    return NewItemInstallationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
