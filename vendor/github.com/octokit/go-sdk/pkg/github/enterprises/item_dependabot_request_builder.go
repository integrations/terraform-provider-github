package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\dependabot
type ItemDependabotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
// returns a *ItemDependabotAlertsRequestBuilder when successful
func (m *ItemDependabotRequestBuilder) Alerts()(*ItemDependabotAlertsRequestBuilder) {
    return NewItemDependabotAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDependabotRequestBuilderInternal instantiates a new ItemDependabotRequestBuilder and sets the default values.
func NewItemDependabotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotRequestBuilder) {
    m := &ItemDependabotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/dependabot", pathParameters),
    }
    return m
}
// NewItemDependabotRequestBuilder instantiates a new ItemDependabotRequestBuilder and sets the default values.
func NewItemDependabotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotRequestBuilderInternal(urlParams, requestAdapter)
}
