package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\personal-access-token-requests\{pat_request_id}
type ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderInternal instantiates a new ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder and sets the default values.
func NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) {
    m := &ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/personal-access-token-requests/{pat_request_id}", pathParameters),
    }
    return m
}
// NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder instantiates a new ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder and sets the default values.
func NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approves or denies a pending request to access organization resources via a fine-grained personal access token.Only GitHub Apps can use this endpoint.
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#review-a-request-to-access-organization-resources-with-a-fine-grained-personal-access-token
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) Post(ctx context.Context, body ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Repositories the repositories property
// returns a *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder when successful
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) Repositories()(*ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPostRequestInformation approves or denies a pending request to access organization resources via a fine-grained personal access token.Only GitHub Apps can use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder when successful
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
