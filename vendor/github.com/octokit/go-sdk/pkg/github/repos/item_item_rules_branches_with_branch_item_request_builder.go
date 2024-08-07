package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemRulesBranchesWithBranchItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\rules\branches\{branch}
type ItemItemRulesBranchesWithBranchItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters returns all active rules that apply to the specified branch. The branch does not need to exist; rules that would applyto a branch with that name will be returned. All active rules that apply will be returned, regardless of the levelat which they are configured (e.g. repository or organization). Rules in rulesets with "evaluate" or "disabled"enforcement statuses are not returned.
type ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal instantiates a new ItemItemRulesBranchesWithBranchItemRequestBuilder and sets the default values.
func NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    m := &ItemItemRulesBranchesWithBranchItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/rules/branches/{branch}{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemRulesBranchesWithBranchItemRequestBuilder instantiates a new ItemItemRulesBranchesWithBranchItemRequestBuilder and sets the default values.
func NewItemItemRulesBranchesWithBranchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all active rules that apply to the specified branch. The branch does not need to exist; rules that would applyto a branch with that name will be returned. All active rules that apply will be returned, regardless of the levelat which they are configured (e.g. repository or organization). Rules in rulesets with "evaluate" or "disabled"enforcement statuses are not returned.
// returns a []RepositoryRuleDetailedable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/rules#get-rules-for-a-branch
func (m *ItemItemRulesBranchesWithBranchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryRuleDetailedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRepositoryRuleDetailedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryRuleDetailedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryRuleDetailedable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns all active rules that apply to the specified branch. The branch does not need to exist; rules that would applyto a branch with that name will be returned. All active rules that apply will be returned, regardless of the levelat which they are configured (e.g. repository or organization). Rules in rulesets with "evaluate" or "disabled"enforcement statuses are not returned.
// returns a *RequestInformation when successful
func (m *ItemItemRulesBranchesWithBranchItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemRulesBranchesWithBranchItemRequestBuilder when successful
func (m *ItemItemRulesBranchesWithBranchItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    return NewItemItemRulesBranchesWithBranchItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
