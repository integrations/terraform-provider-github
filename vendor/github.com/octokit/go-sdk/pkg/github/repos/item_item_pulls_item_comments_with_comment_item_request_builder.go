package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPullsItemCommentsWithComment_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\pulls\{pull_number}\comments\{comment_id}
type ItemItemPullsItemCommentsWithComment_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal instantiates a new ItemItemPullsItemCommentsWithComment_ItemRequestBuilder and sets the default values.
func NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) {
    m := &ItemItemPullsItemCommentsWithComment_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/pulls/{pull_number}/comments/{comment_id}", pathParameters),
    }
    return m
}
// NewItemItemPullsItemCommentsWithComment_ItemRequestBuilder instantiates a new ItemItemPullsItemCommentsWithComment_ItemRequestBuilder and sets the default values.
func NewItemItemPullsItemCommentsWithComment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Replies the replies property
// returns a *ItemItemPullsItemCommentsItemRepliesRequestBuilder when successful
func (m *ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) Replies()(*ItemItemPullsItemCommentsItemRepliesRequestBuilder) {
    return NewItemItemPullsItemCommentsItemRepliesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
