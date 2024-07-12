package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemAttestationsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\attestations
type ItemItemAttestationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BySubject_digest gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.attestations.item collection
// returns a *ItemItemAttestationsWithSubject_digestItemRequestBuilder when successful
func (m *ItemItemAttestationsRequestBuilder) BySubject_digest(subject_digest string)(*ItemItemAttestationsWithSubject_digestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if subject_digest != "" {
        urlTplParams["subject_digest"] = subject_digest
    }
    return NewItemItemAttestationsWithSubject_digestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemAttestationsRequestBuilderInternal instantiates a new ItemItemAttestationsRequestBuilder and sets the default values.
func NewItemItemAttestationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAttestationsRequestBuilder) {
    m := &ItemItemAttestationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/attestations", pathParameters),
    }
    return m
}
// NewItemItemAttestationsRequestBuilder instantiates a new ItemItemAttestationsRequestBuilder and sets the default values.
func NewItemItemAttestationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAttestationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemAttestationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post store an artifact attestation and associate it with a repository.The authenticated user must have write permission to the repository and, if using a fine-grained access token the `attestations:write` permission is required.Artifact attestations are meant to be created using the [attest action](https://github.com/actions/attest). For amore information, see our guide on [using artifact attestations to establish a build's provenance](https://docs.github.com/actions/security-guides/using-artifact-attestations-to-establish-provenance-for-builds).
// returns a ItemItemAttestationsPostResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#create-an-attestation
func (m *ItemItemAttestationsRequestBuilder) Post(ctx context.Context, body ItemItemAttestationsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemItemAttestationsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemAttestationsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemAttestationsPostResponseable), nil
}
// ToPostRequestInformation store an artifact attestation and associate it with a repository.The authenticated user must have write permission to the repository and, if using a fine-grained access token the `attestations:write` permission is required.Artifact attestations are meant to be created using the [attest action](https://github.com/actions/attest). For amore information, see our guide on [using artifact attestations to establish a build's provenance](https://docs.github.com/actions/security-guides/using-artifact-attestations-to-establish-provenance-for-builds).
// returns a *RequestInformation when successful
func (m *ItemItemAttestationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemAttestationsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemAttestationsRequestBuilder when successful
func (m *ItemItemAttestationsRequestBuilder) WithUrl(rawUrl string)(*ItemItemAttestationsRequestBuilder) {
    return NewItemItemAttestationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
