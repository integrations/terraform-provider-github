package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamWithTeam_slugItemRequestBuilder builds and executes requests for operations under \orgs\{org}\team\{team_slug}
type ItemTeamWithTeam_slugItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTeamWithTeam_slugItemRequestBuilderInternal instantiates a new ItemTeamWithTeam_slugItemRequestBuilder and sets the default values.
func NewItemTeamWithTeam_slugItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamWithTeam_slugItemRequestBuilder) {
    m := &ItemTeamWithTeam_slugItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/team/{team_slug}", pathParameters),
    }
    return m
}
// NewItemTeamWithTeam_slugItemRequestBuilder instantiates a new ItemTeamWithTeam_slugItemRequestBuilder and sets the default values.
func NewItemTeamWithTeam_slugItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamWithTeam_slugItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamWithTeam_slugItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copilot the copilot property
// returns a *ItemTeamItemCopilotRequestBuilder when successful
func (m *ItemTeamWithTeam_slugItemRequestBuilder) Copilot()(*ItemTeamItemCopilotRequestBuilder) {
    return NewItemTeamItemCopilotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
