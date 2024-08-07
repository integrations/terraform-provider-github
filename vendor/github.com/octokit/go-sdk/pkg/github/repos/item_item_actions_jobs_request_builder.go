package repos

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsJobsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\jobs
type ItemItemActionsJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByJob_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.actions.jobs.item collection
// returns a *ItemItemActionsJobsWithJob_ItemRequestBuilder when successful
func (m *ItemItemActionsJobsRequestBuilder) ByJob_id(job_id int32)(*ItemItemActionsJobsWithJob_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["job_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(job_id), 10)
    return NewItemItemActionsJobsWithJob_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsJobsRequestBuilderInternal instantiates a new ItemItemActionsJobsRequestBuilder and sets the default values.
func NewItemItemActionsJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsRequestBuilder) {
    m := &ItemItemActionsJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/jobs", pathParameters),
    }
    return m
}
// NewItemItemActionsJobsRequestBuilder instantiates a new ItemItemActionsJobsRequestBuilder and sets the default values.
func NewItemItemActionsJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsJobsRequestBuilderInternal(urlParams, requestAdapter)
}
