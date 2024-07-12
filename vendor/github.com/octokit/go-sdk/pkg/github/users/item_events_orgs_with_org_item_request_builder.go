package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemEventsOrgsWithOrgItemRequestBuilder builds and executes requests for operations under \users\{username}\events\orgs\{org}
type ItemEventsOrgsWithOrgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters this is the user's organization dashboard. You must be authenticated as the user to view this.
type ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemEventsOrgsWithOrgItemRequestBuilderInternal instantiates a new ItemEventsOrgsWithOrgItemRequestBuilder and sets the default values.
func NewItemEventsOrgsWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    m := &ItemEventsOrgsWithOrgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/events/orgs/{org}{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemEventsOrgsWithOrgItemRequestBuilder instantiates a new ItemEventsOrgsWithOrgItemRequestBuilder and sets the default values.
func NewItemEventsOrgsWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsOrgsWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this is the user's organization dashboard. You must be authenticated as the user to view this.
// returns a []Eventable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/events#list-organization-events-for-the-authenticated-user
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Eventable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Eventable)
        }
    }
    return val, nil
}
// ToGetRequestInformation this is the user's organization dashboard. You must be authenticated as the user to view this.
// returns a *RequestInformation when successful
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemEventsOrgsWithOrgItemRequestBuilder when successful
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) WithUrl(rawUrl string)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    return NewItemEventsOrgsWithOrgItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
