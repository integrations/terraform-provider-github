package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnterprisesRequestBuilder builds and executes requests for operations under \enterprises
type EnterprisesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByEnterprise gets an item from the github.com/octokit/go-sdk/pkg/github.enterprises.item collection
// returns a *WithEnterpriseItemRequestBuilder when successful
func (m *EnterprisesRequestBuilder) ByEnterprise(enterprise string)(*WithEnterpriseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if enterprise != "" {
        urlTplParams["enterprise"] = enterprise
    }
    return NewWithEnterpriseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterprisesRequestBuilderInternal instantiates a new EnterprisesRequestBuilder and sets the default values.
func NewEnterprisesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterprisesRequestBuilder) {
    m := &EnterprisesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises", pathParameters),
    }
    return m
}
// NewEnterprisesRequestBuilder instantiates a new EnterprisesRequestBuilder and sets the default values.
func NewEnterprisesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterprisesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterprisesRequestBuilderInternal(urlParams, requestAdapter)
}
