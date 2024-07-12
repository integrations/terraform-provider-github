package projects

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ColumnsRequestBuilder builds and executes requests for operations under \projects\columns
type ColumnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByColumn_id gets an item from the github.com/octokit/go-sdk/pkg/github.projects.columns.item collection
// returns a *ColumnsWithColumn_ItemRequestBuilder when successful
func (m *ColumnsRequestBuilder) ByColumn_id(column_id int32)(*ColumnsWithColumn_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["column_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(column_id), 10)
    return NewColumnsWithColumn_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Cards the cards property
// returns a *ColumnsCardsRequestBuilder when successful
func (m *ColumnsRequestBuilder) Cards()(*ColumnsCardsRequestBuilder) {
    return NewColumnsCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewColumnsRequestBuilderInternal instantiates a new ColumnsRequestBuilder and sets the default values.
func NewColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsRequestBuilder) {
    m := &ColumnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/columns", pathParameters),
    }
    return m
}
// NewColumnsRequestBuilder instantiates a new ColumnsRequestBuilder and sets the default values.
func NewColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewColumnsRequestBuilderInternal(urlParams, requestAdapter)
}
