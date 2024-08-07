package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\stats
type ItemItemStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Code_frequency the code_frequency property
// returns a *ItemItemStatsCode_frequencyRequestBuilder when successful
func (m *ItemItemStatsRequestBuilder) Code_frequency()(*ItemItemStatsCode_frequencyRequestBuilder) {
    return NewItemItemStatsCode_frequencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commit_activity the commit_activity property
// returns a *ItemItemStatsCommit_activityRequestBuilder when successful
func (m *ItemItemStatsRequestBuilder) Commit_activity()(*ItemItemStatsCommit_activityRequestBuilder) {
    return NewItemItemStatsCommit_activityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemStatsRequestBuilderInternal instantiates a new ItemItemStatsRequestBuilder and sets the default values.
func NewItemItemStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsRequestBuilder) {
    m := &ItemItemStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/stats", pathParameters),
    }
    return m
}
// NewItemItemStatsRequestBuilder instantiates a new ItemItemStatsRequestBuilder and sets the default values.
func NewItemItemStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Contributors the contributors property
// returns a *ItemItemStatsContributorsRequestBuilder when successful
func (m *ItemItemStatsRequestBuilder) Contributors()(*ItemItemStatsContributorsRequestBuilder) {
    return NewItemItemStatsContributorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Participation the participation property
// returns a *ItemItemStatsParticipationRequestBuilder when successful
func (m *ItemItemStatsRequestBuilder) Participation()(*ItemItemStatsParticipationRequestBuilder) {
    return NewItemItemStatsParticipationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Punch_card the punch_card property
// returns a *ItemItemStatsPunch_cardRequestBuilder when successful
func (m *ItemItemStatsRequestBuilder) Punch_card()(*ItemItemStatsPunch_cardRequestBuilder) {
    return NewItemItemStatsPunch_cardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
