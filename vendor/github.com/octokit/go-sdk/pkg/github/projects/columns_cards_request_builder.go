package projects

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ColumnsCardsRequestBuilder builds and executes requests for operations under \projects\columns\cards
type ColumnsCardsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCard_id gets an item from the github.com/octokit/go-sdk/pkg/github.projects.columns.cards.item collection
// returns a *ColumnsCardsWithCard_ItemRequestBuilder when successful
func (m *ColumnsCardsRequestBuilder) ByCard_id(card_id int32)(*ColumnsCardsWithCard_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["card_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(card_id), 10)
    return NewColumnsCardsWithCard_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewColumnsCardsRequestBuilderInternal instantiates a new ColumnsCardsRequestBuilder and sets the default values.
func NewColumnsCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsCardsRequestBuilder) {
    m := &ColumnsCardsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/columns/cards", pathParameters),
    }
    return m
}
// NewColumnsCardsRequestBuilder instantiates a new ColumnsCardsRequestBuilder and sets the default values.
func NewColumnsCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsCardsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewColumnsCardsRequestBuilderInternal(urlParams, requestAdapter)
}
