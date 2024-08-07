package octocat

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OctocatRequestBuilder builds and executes requests for operations under \octocat
type OctocatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OctocatRequestBuilderGetQueryParameters get the octocat as ASCII art
type OctocatRequestBuilderGetQueryParameters struct {
    // The words to show in Octocat's speech bubble
    S *string `uriparametername:"s"`
}
// NewOctocatRequestBuilderInternal instantiates a new OctocatRequestBuilder and sets the default values.
func NewOctocatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OctocatRequestBuilder) {
    m := &OctocatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/octocat{?s*}", pathParameters),
    }
    return m
}
// NewOctocatRequestBuilder instantiates a new OctocatRequestBuilder and sets the default values.
func NewOctocatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OctocatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOctocatRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the octocat as ASCII art
// returns a []byte when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#get-octocat
func (m *OctocatRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[OctocatRequestBuilderGetQueryParameters])([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the octocat as ASCII art
// returns a *RequestInformation when successful
func (m *OctocatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[OctocatRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/octocat-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OctocatRequestBuilder when successful
func (m *OctocatRequestBuilder) WithUrl(rawUrl string)(*OctocatRequestBuilder) {
    return NewOctocatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
