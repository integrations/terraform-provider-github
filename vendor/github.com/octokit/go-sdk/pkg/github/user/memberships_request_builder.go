package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MembershipsRequestBuilder builds and executes requests for operations under \user\memberships
type MembershipsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMembershipsRequestBuilderInternal instantiates a new MembershipsRequestBuilder and sets the default values.
func NewMembershipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsRequestBuilder) {
    m := &MembershipsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/memberships", pathParameters),
    }
    return m
}
// NewMembershipsRequestBuilder instantiates a new MembershipsRequestBuilder and sets the default values.
func NewMembershipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembershipsRequestBuilderInternal(urlParams, requestAdapter)
}
// Orgs the orgs property
// returns a *MembershipsOrgsRequestBuilder when successful
func (m *MembershipsRequestBuilder) Orgs()(*MembershipsOrgsRequestBuilder) {
    return NewMembershipsOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
