package codes_of_conduct

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// Codes_of_conductRequestBuilder builds and executes requests for operations under \codes_of_conduct
type Codes_of_conductRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByKey gets an item from the github.com/octokit/go-sdk/pkg/github.codes_of_conduct.item collection
// returns a *WithKeyItemRequestBuilder when successful
func (m *Codes_of_conductRequestBuilder) ByKey(key string)(*WithKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if key != "" {
        urlTplParams["key"] = key
    }
    return NewWithKeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCodes_of_conductRequestBuilderInternal instantiates a new Codes_of_conductRequestBuilder and sets the default values.
func NewCodes_of_conductRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Codes_of_conductRequestBuilder) {
    m := &Codes_of_conductRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/codes_of_conduct", pathParameters),
    }
    return m
}
// NewCodes_of_conductRequestBuilder instantiates a new Codes_of_conductRequestBuilder and sets the default values.
func NewCodes_of_conductRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Codes_of_conductRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodes_of_conductRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns array of all GitHub's codes of conduct.
// returns a []CodeOfConductable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codes-of-conduct/codes-of-conduct#get-all-codes-of-conduct
func (m *Codes_of_conductRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeOfConductable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeOfConductFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeOfConductable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeOfConductable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns array of all GitHub's codes of conduct.
// returns a *RequestInformation when successful
func (m *Codes_of_conductRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Codes_of_conductRequestBuilder when successful
func (m *Codes_of_conductRequestBuilder) WithUrl(rawUrl string)(*Codes_of_conductRequestBuilder) {
    return NewCodes_of_conductRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
