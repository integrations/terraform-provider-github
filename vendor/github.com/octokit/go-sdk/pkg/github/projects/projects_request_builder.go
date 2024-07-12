package projects

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsRequestBuilder builds and executes requests for operations under \projects
type ProjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByProject_id gets an item from the github.com/octokit/go-sdk/pkg/github.projects.item collection
// returns a *WithProject_ItemRequestBuilder when successful
func (m *ProjectsRequestBuilder) ByProject_id(project_id int32)(*WithProject_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["project_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(project_id), 10)
    return NewWithProject_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Columns the columns property
// returns a *ColumnsRequestBuilder when successful
func (m *ProjectsRequestBuilder) Columns()(*ColumnsRequestBuilder) {
    return NewColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewProjectsRequestBuilderInternal instantiates a new ProjectsRequestBuilder and sets the default values.
func NewProjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsRequestBuilder) {
    m := &ProjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects", pathParameters),
    }
    return m
}
// NewProjectsRequestBuilder instantiates a new ProjectsRequestBuilder and sets the default values.
func NewProjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsRequestBuilderInternal(urlParams, requestAdapter)
}
