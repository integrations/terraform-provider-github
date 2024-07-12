package gitignore

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GitignoreRequestBuilder builds and executes requests for operations under \gitignore
type GitignoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewGitignoreRequestBuilderInternal instantiates a new GitignoreRequestBuilder and sets the default values.
func NewGitignoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GitignoreRequestBuilder) {
    m := &GitignoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/gitignore", pathParameters),
    }
    return m
}
// NewGitignoreRequestBuilder instantiates a new GitignoreRequestBuilder and sets the default values.
func NewGitignoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GitignoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGitignoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Templates the templates property
// returns a *TemplatesRequestBuilder when successful
func (m *GitignoreRequestBuilder) Templates()(*TemplatesRequestBuilder) {
    return NewTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
