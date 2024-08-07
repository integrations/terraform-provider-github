package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsCode_frequencyRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\stats\code_frequency
type ItemItemStatsCode_frequencyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemStatsCode_frequencyRequestBuilderInternal instantiates a new ItemItemStatsCode_frequencyRequestBuilder and sets the default values.
func NewItemItemStatsCode_frequencyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCode_frequencyRequestBuilder) {
    m := &ItemItemStatsCode_frequencyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/stats/code_frequency", pathParameters),
    }
    return m
}
// NewItemItemStatsCode_frequencyRequestBuilder instantiates a new ItemItemStatsCode_frequencyRequestBuilder and sets the default values.
func NewItemItemStatsCode_frequencyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCode_frequencyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsCode_frequencyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a weekly aggregate of the number of additions and deletions pushed to a repository.**Note:** This endpoint can only be used for repositories with fewer than 10,000 commits. If the repository contains10,000 or more commits, a 422 status code will be returned.
// returns a []int32 when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/statistics#get-the-weekly-commit-activity
func (m *ItemItemStatsCode_frequencyRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    val := make([]int32, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*int32))
        }
    }
    return val, nil
}
// ToGetRequestInformation returns a weekly aggregate of the number of additions and deletions pushed to a repository.**Note:** This endpoint can only be used for repositories with fewer than 10,000 commits. If the repository contains10,000 or more commits, a 422 status code will be returned.
// returns a *RequestInformation when successful
func (m *ItemItemStatsCode_frequencyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemStatsCode_frequencyRequestBuilder when successful
func (m *ItemItemStatsCode_frequencyRequestBuilder) WithUrl(rawUrl string)(*ItemItemStatsCode_frequencyRequestBuilder) {
    return NewItemItemStatsCode_frequencyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
