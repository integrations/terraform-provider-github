package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamItemCopilotRequestBuilder builds and executes requests for operations under \orgs\{org}\team\{team_slug}\copilot
type ItemTeamItemCopilotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTeamItemCopilotRequestBuilderInternal instantiates a new ItemTeamItemCopilotRequestBuilder and sets the default values.
func NewItemTeamItemCopilotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamItemCopilotRequestBuilder) {
    m := &ItemTeamItemCopilotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/team/{team_slug}/copilot", pathParameters),
    }
    return m
}
// NewItemTeamItemCopilotRequestBuilder instantiates a new ItemTeamItemCopilotRequestBuilder and sets the default values.
func NewItemTeamItemCopilotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamItemCopilotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamItemCopilotRequestBuilderInternal(urlParams, requestAdapter)
}
// Usage the usage property
// returns a *ItemTeamItemCopilotUsageRequestBuilder when successful
func (m *ItemTeamItemCopilotRequestBuilder) Usage()(*ItemTeamItemCopilotUsageRequestBuilder) {
    return NewItemTeamItemCopilotUsageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
