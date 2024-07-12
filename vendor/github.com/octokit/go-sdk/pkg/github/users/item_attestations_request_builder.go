package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAttestationsRequestBuilder builds and executes requests for operations under \users\{username}\attestations
type ItemAttestationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BySubject_digest gets an item from the github.com/octokit/go-sdk/pkg/github.users.item.attestations.item collection
// returns a *ItemAttestationsWithSubject_digestItemRequestBuilder when successful
func (m *ItemAttestationsRequestBuilder) BySubject_digest(subject_digest string)(*ItemAttestationsWithSubject_digestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if subject_digest != "" {
        urlTplParams["subject_digest"] = subject_digest
    }
    return NewItemAttestationsWithSubject_digestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAttestationsRequestBuilderInternal instantiates a new ItemAttestationsRequestBuilder and sets the default values.
func NewItemAttestationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttestationsRequestBuilder) {
    m := &ItemAttestationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/attestations", pathParameters),
    }
    return m
}
// NewItemAttestationsRequestBuilder instantiates a new ItemAttestationsRequestBuilder and sets the default values.
func NewItemAttestationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAttestationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAttestationsRequestBuilderInternal(urlParams, requestAdapter)
}
