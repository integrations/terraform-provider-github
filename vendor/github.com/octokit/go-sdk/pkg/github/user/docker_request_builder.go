package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DockerRequestBuilder builds and executes requests for operations under \user\docker
type DockerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Conflicts the conflicts property
// returns a *DockerConflictsRequestBuilder when successful
func (m *DockerRequestBuilder) Conflicts()(*DockerConflictsRequestBuilder) {
    return NewDockerConflictsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDockerRequestBuilderInternal instantiates a new DockerRequestBuilder and sets the default values.
func NewDockerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DockerRequestBuilder) {
    m := &DockerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/docker", pathParameters),
    }
    return m
}
// NewDockerRequestBuilder instantiates a new DockerRequestBuilder and sets the default values.
func NewDockerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DockerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDockerRequestBuilderInternal(urlParams, requestAdapter)
}
