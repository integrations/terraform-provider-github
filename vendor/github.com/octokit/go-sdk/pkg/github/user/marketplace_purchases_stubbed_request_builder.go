package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// Marketplace_purchasesStubbedRequestBuilder builds and executes requests for operations under \user\marketplace_purchases\stubbed
type Marketplace_purchasesStubbedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Marketplace_purchasesStubbedRequestBuilderGetQueryParameters lists the active subscriptions for the authenticated user.
type Marketplace_purchasesStubbedRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewMarketplace_purchasesStubbedRequestBuilderInternal instantiates a new Marketplace_purchasesStubbedRequestBuilder and sets the default values.
func NewMarketplace_purchasesStubbedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_purchasesStubbedRequestBuilder) {
    m := &Marketplace_purchasesStubbedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/marketplace_purchases/stubbed{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewMarketplace_purchasesStubbedRequestBuilder instantiates a new Marketplace_purchasesStubbedRequestBuilder and sets the default values.
func NewMarketplace_purchasesStubbedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_purchasesStubbedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMarketplace_purchasesStubbedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the active subscriptions for the authenticated user.
// returns a []UserMarketplacePurchaseable when successful
// returns a BasicError error when the service returns a 401 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/marketplace#list-subscriptions-for-the-authenticated-user-stubbed
func (m *Marketplace_purchasesStubbedRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[Marketplace_purchasesStubbedRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.UserMarketplacePurchaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateUserMarketplacePurchaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.UserMarketplacePurchaseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.UserMarketplacePurchaseable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the active subscriptions for the authenticated user.
// returns a *RequestInformation when successful
func (m *Marketplace_purchasesStubbedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[Marketplace_purchasesStubbedRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Marketplace_purchasesStubbedRequestBuilder when successful
func (m *Marketplace_purchasesStubbedRequestBuilder) WithUrl(rawUrl string)(*Marketplace_purchasesStubbedRequestBuilder) {
    return NewMarketplace_purchasesStubbedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
