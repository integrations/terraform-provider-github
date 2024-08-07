package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependabotRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\dependabot
type ItemItemDependabotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
// returns a *ItemItemDependabotAlertsRequestBuilder when successful
func (m *ItemItemDependabotRequestBuilder) Alerts()(*ItemItemDependabotAlertsRequestBuilder) {
    return NewItemItemDependabotAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemDependabotRequestBuilderInternal instantiates a new ItemItemDependabotRequestBuilder and sets the default values.
func NewItemItemDependabotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotRequestBuilder) {
    m := &ItemItemDependabotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/dependabot", pathParameters),
    }
    return m
}
// NewItemItemDependabotRequestBuilder instantiates a new ItemItemDependabotRequestBuilder and sets the default values.
func NewItemItemDependabotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependabotRequestBuilderInternal(urlParams, requestAdapter)
}
// Secrets the secrets property
// returns a *ItemItemDependabotSecretsRequestBuilder when successful
func (m *ItemItemDependabotRequestBuilder) Secrets()(*ItemItemDependabotSecretsRequestBuilder) {
    return NewItemItemDependabotSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
