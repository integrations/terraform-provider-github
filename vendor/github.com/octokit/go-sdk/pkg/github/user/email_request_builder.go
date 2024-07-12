package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailRequestBuilder builds and executes requests for operations under \user\email
type EmailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewEmailRequestBuilderInternal instantiates a new EmailRequestBuilder and sets the default values.
func NewEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailRequestBuilder) {
    m := &EmailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/email", pathParameters),
    }
    return m
}
// NewEmailRequestBuilder instantiates a new EmailRequestBuilder and sets the default values.
func NewEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Visibility the visibility property
// returns a *EmailVisibilityRequestBuilder when successful
func (m *EmailRequestBuilder) Visibility()(*EmailVisibilityRequestBuilder) {
    return NewEmailVisibilityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
