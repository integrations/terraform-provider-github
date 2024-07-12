package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemGitTagsWithTag_shaItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\git\tags\{tag_sha}
type ItemItemGitTagsWithTag_shaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemGitTagsWithTag_shaItemRequestBuilderInternal instantiates a new ItemItemGitTagsWithTag_shaItemRequestBuilder and sets the default values.
func NewItemItemGitTagsWithTag_shaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitTagsWithTag_shaItemRequestBuilder) {
    m := &ItemItemGitTagsWithTag_shaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/git/tags/{tag_sha}", pathParameters),
    }
    return m
}
// NewItemItemGitTagsWithTag_shaItemRequestBuilder instantiates a new ItemItemGitTagsWithTag_shaItemRequestBuilder and sets the default values.
func NewItemItemGitTagsWithTag_shaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitTagsWithTag_shaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitTagsWithTag_shaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on their account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
// returns a GitTagable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 409 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/git/tags#get-a-tag
func (m *ItemItemGitTagsWithTag_shaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GitTagable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "409": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateGitTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GitTagable), nil
}
// ToGetRequestInformation **Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on their account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
// returns a *RequestInformation when successful
func (m *ItemItemGitTagsWithTag_shaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemGitTagsWithTag_shaItemRequestBuilder when successful
func (m *ItemItemGitTagsWithTag_shaItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemGitTagsWithTag_shaItemRequestBuilder) {
    return NewItemItemGitTagsWithTag_shaItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
