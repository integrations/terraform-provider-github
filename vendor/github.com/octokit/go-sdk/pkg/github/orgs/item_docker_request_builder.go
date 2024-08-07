package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDockerRequestBuilder builds and executes requests for operations under \orgs\{org}\docker
type ItemDockerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Conflicts the conflicts property
// returns a *ItemDockerConflictsRequestBuilder when successful
func (m *ItemDockerRequestBuilder) Conflicts()(*ItemDockerConflictsRequestBuilder) {
    return NewItemDockerConflictsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDockerRequestBuilderInternal instantiates a new ItemDockerRequestBuilder and sets the default values.
func NewItemDockerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDockerRequestBuilder) {
    m := &ItemDockerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/docker", pathParameters),
    }
    return m
}
// NewItemDockerRequestBuilder instantiates a new ItemDockerRequestBuilder and sets the default values.
func NewItemDockerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDockerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDockerRequestBuilderInternal(urlParams, requestAdapter)
}
