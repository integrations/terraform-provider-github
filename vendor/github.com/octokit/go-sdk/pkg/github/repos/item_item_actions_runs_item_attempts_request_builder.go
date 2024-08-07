package repos

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunsItemAttemptsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\attempts
type ItemItemActionsRunsItemAttemptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByAttempt_number gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.actions.runs.item.attempts.item collection
// returns a *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder when successful
func (m *ItemItemActionsRunsItemAttemptsRequestBuilder) ByAttempt_number(attempt_number int32)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["attempt_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(attempt_number), 10)
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunsItemAttemptsRequestBuilderInternal instantiates a new ItemItemActionsRunsItemAttemptsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/attempts", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemAttemptsRequestBuilder instantiates a new ItemItemActionsRunsItemAttemptsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
