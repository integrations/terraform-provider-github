package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemRulesBranchesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\rules\branches
type ItemItemRulesBranchesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByBranch gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.rules.branches.item collection
// returns a *ItemItemRulesBranchesWithBranchItemRequestBuilder when successful
func (m *ItemItemRulesBranchesRequestBuilder) ByBranch(branch string)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if branch != "" {
        urlTplParams["branch"] = branch
    }
    return NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemRulesBranchesRequestBuilderInternal instantiates a new ItemItemRulesBranchesRequestBuilder and sets the default values.
func NewItemItemRulesBranchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesBranchesRequestBuilder) {
    m := &ItemItemRulesBranchesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/rules/branches", pathParameters),
    }
    return m
}
// NewItemItemRulesBranchesRequestBuilder instantiates a new ItemItemRulesBranchesRequestBuilder and sets the default values.
func NewItemItemRulesBranchesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesBranchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesBranchesRequestBuilderInternal(urlParams, requestAdapter)
}
