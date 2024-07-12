package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemWithRepoItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}
type ItemWithRepoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWithRepoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithRepoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemWithRepoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithRepoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemWithRepoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithRepoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Actions the actions property
func (m *ItemWithRepoItemRequestBuilder) Actions()(*ItemItemActionsRequestBuilder) {
    return NewItemItemActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Activity the activity property
func (m *ItemWithRepoItemRequestBuilder) Activity()(*ItemItemActivityRequestBuilder) {
    return NewItemItemActivityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignees the assignees property
func (m *ItemWithRepoItemRequestBuilder) Assignees()(*ItemItemAssigneesRequestBuilder) {
    return NewItemItemAssigneesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Autolinks the autolinks property
func (m *ItemWithRepoItemRequestBuilder) Autolinks()(*ItemItemAutolinksRequestBuilder) {
    return NewItemItemAutolinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AutomatedSecurityFixes the automatedSecurityFixes property
func (m *ItemWithRepoItemRequestBuilder) AutomatedSecurityFixes()(*ItemItemAutomatedSecurityFixesRequestBuilder) {
    return NewItemItemAutomatedSecurityFixesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Branches the branches property
func (m *ItemWithRepoItemRequestBuilder) Branches()(*ItemItemBranchesRequestBuilder) {
    return NewItemItemBranchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckRuns the checkRuns property
func (m *ItemWithRepoItemRequestBuilder) CheckRuns()(*ItemItemCheckRunsRequestBuilder) {
    return NewItemItemCheckRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckSuites the checkSuites property
func (m *ItemWithRepoItemRequestBuilder) CheckSuites()(*ItemItemCheckSuitesRequestBuilder) {
    return NewItemItemCheckSuitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codeowners the codeowners property
func (m *ItemWithRepoItemRequestBuilder) Codeowners()(*ItemItemCodeownersRequestBuilder) {
    return NewItemItemCodeownersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CodeScanning the codeScanning property
func (m *ItemWithRepoItemRequestBuilder) CodeScanning()(*ItemItemCodeScanningRequestBuilder) {
    return NewItemItemCodeScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codespaces the codespaces property
func (m *ItemWithRepoItemRequestBuilder) Codespaces()(*ItemItemCodespacesRequestBuilder) {
    return NewItemItemCodespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Collaborators the collaborators property
func (m *ItemWithRepoItemRequestBuilder) Collaborators()(*ItemItemCollaboratorsRequestBuilder) {
    return NewItemItemCollaboratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
func (m *ItemWithRepoItemRequestBuilder) Comments()(*ItemItemCommentsRequestBuilder) {
    return NewItemItemCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commits the commits property
func (m *ItemWithRepoItemRequestBuilder) Commits()(*ItemItemCommitsRequestBuilder) {
    return NewItemItemCommitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Community the community property
func (m *ItemWithRepoItemRequestBuilder) Community()(*ItemItemCommunityRequestBuilder) {
    return NewItemItemCommunityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Compare the compare property
func (m *ItemWithRepoItemRequestBuilder) Compare()(*ItemItemCompareRequestBuilder) {
    return NewItemItemCompareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemWithRepoItemRequestBuilderInternal instantiates a new WithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithRepoItemRequestBuilder) {
    m := &ItemWithRepoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}", pathParameters),
    }
    return m
}
// NewItemWithRepoItemRequestBuilder instantiates a new WithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithRepoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithRepoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Contents the contents property
func (m *ItemWithRepoItemRequestBuilder) Contents()(*ItemItemContentsRequestBuilder) {
    return NewItemItemContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contributors the contributors property
func (m *ItemWithRepoItemRequestBuilder) Contributors()(*ItemItemContributorsRequestBuilder) {
    return NewItemItemContributorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deleting a repository requires admin access. If OAuth is used, the `delete_repo` scope is required.If an organization owner has configured the organization to prevent members from deleting organization-ownedrepositories, you will get a `403 Forbidden` response.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#delete-a-repository
func (m *ItemWithRepoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": CreateItemItemWithRepo403ErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Dependabot the dependabot property
func (m *ItemWithRepoItemRequestBuilder) Dependabot()(*ItemItemDependabotRequestBuilder) {
    return NewItemItemDependabotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DependencyGraph the dependencyGraph property
func (m *ItemWithRepoItemRequestBuilder) DependencyGraph()(*ItemItemDependencyGraphRequestBuilder) {
    return NewItemItemDependencyGraphRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deployments the deployments property
func (m *ItemWithRepoItemRequestBuilder) Deployments()(*ItemItemDeploymentsRequestBuilder) {
    return NewItemItemDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dispatches the dispatches property
func (m *ItemWithRepoItemRequestBuilder) Dispatches()(*ItemItemDispatchesRequestBuilder) {
    return NewItemItemDispatchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Environments the environments property
func (m *ItemWithRepoItemRequestBuilder) Environments()(*ItemItemEnvironmentsRequestBuilder) {
    return NewItemItemEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
func (m *ItemWithRepoItemRequestBuilder) Events()(*ItemItemEventsRequestBuilder) {
    return NewItemItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forks the forks property
func (m *ItemWithRepoItemRequestBuilder) Forks()(*ItemItemForksRequestBuilder) {
    return NewItemItemForksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Generate the generate property
func (m *ItemWithRepoItemRequestBuilder) Generate()(*ItemItemGenerateRequestBuilder) {
    return NewItemItemGenerateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#get-a-repository
func (m *ItemWithRepoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FullRepositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateFullRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FullRepositoryable), nil
}
// Git the git property
func (m *ItemWithRepoItemRequestBuilder) Git()(*ItemItemGitRequestBuilder) {
    return NewItemItemGitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hooks the hooks property
func (m *ItemWithRepoItemRequestBuilder) Hooks()(*ItemItemHooksRequestBuilder) {
    return NewItemItemHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportEscaped the import property
func (m *ItemWithRepoItemRequestBuilder) ImportEscaped()(*ItemItemImportRequestBuilder) {
    return NewItemItemImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
func (m *ItemWithRepoItemRequestBuilder) Installation()(*ItemItemInstallationRequestBuilder) {
    return NewItemItemInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InteractionLimits the interactionLimits property
func (m *ItemWithRepoItemRequestBuilder) InteractionLimits()(*ItemItemInteractionLimitsRequestBuilder) {
    return NewItemItemInteractionLimitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Invitations the invitations property
func (m *ItemWithRepoItemRequestBuilder) Invitations()(*ItemItemInvitationsRequestBuilder) {
    return NewItemItemInvitationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
func (m *ItemWithRepoItemRequestBuilder) Issues()(*ItemItemIssuesRequestBuilder) {
    return NewItemItemIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keys the keys property
func (m *ItemWithRepoItemRequestBuilder) Keys()(*ItemItemKeysRequestBuilder) {
    return NewItemItemKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Labels the labels property
func (m *ItemWithRepoItemRequestBuilder) Labels()(*ItemItemLabelsRequestBuilder) {
    return NewItemItemLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Languages the languages property
func (m *ItemWithRepoItemRequestBuilder) Languages()(*ItemItemLanguagesRequestBuilder) {
    return NewItemItemLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// License the license property
func (m *ItemWithRepoItemRequestBuilder) License()(*ItemItemLicenseRequestBuilder) {
    return NewItemItemLicenseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Merges the merges property
func (m *ItemWithRepoItemRequestBuilder) Merges()(*ItemItemMergesRequestBuilder) {
    return NewItemItemMergesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MergeUpstream the mergeUpstream property
func (m *ItemWithRepoItemRequestBuilder) MergeUpstream()(*ItemItemMergeUpstreamRequestBuilder) {
    return NewItemItemMergeUpstreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Milestones the milestones property
func (m *ItemWithRepoItemRequestBuilder) Milestones()(*ItemItemMilestonesRequestBuilder) {
    return NewItemItemMilestonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
func (m *ItemWithRepoItemRequestBuilder) Notifications()(*ItemItemNotificationsRequestBuilder) {
    return NewItemItemNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages the pages property
func (m *ItemWithRepoItemRequestBuilder) Pages()(*ItemItemPagesRequestBuilder) {
    return NewItemItemPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch **Note**: To edit a repository's topics, use the [Replace all repository topics](https://docs.github.com/rest/repos/repos#replace-all-repository-topics) endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#update-a-repository
func (m *ItemWithRepoItemRequestBuilder) Patch(ctx context.Context, body ItemItemWithRepoPatchRequestBodyable, requestConfiguration *ItemWithRepoItemRequestBuilderPatchRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FullRepositoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateFullRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FullRepositoryable), nil
}
// PrivateVulnerabilityReporting the privateVulnerabilityReporting property
func (m *ItemWithRepoItemRequestBuilder) PrivateVulnerabilityReporting()(*ItemItemPrivateVulnerabilityReportingRequestBuilder) {
    return NewItemItemPrivateVulnerabilityReportingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
func (m *ItemWithRepoItemRequestBuilder) Projects()(*ItemItemProjectsRequestBuilder) {
    return NewItemItemProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Properties the properties property
func (m *ItemWithRepoItemRequestBuilder) Properties()(*ItemItemPropertiesRequestBuilder) {
    return NewItemItemPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pulls the pulls property
func (m *ItemWithRepoItemRequestBuilder) Pulls()(*ItemItemPullsRequestBuilder) {
    return NewItemItemPullsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Readme the readme property
func (m *ItemWithRepoItemRequestBuilder) Readme()(*ItemItemReadmeRequestBuilder) {
    return NewItemItemReadmeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Releases the releases property
func (m *ItemWithRepoItemRequestBuilder) Releases()(*ItemItemReleasesRequestBuilder) {
    return NewItemItemReleasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rules the rules property
func (m *ItemWithRepoItemRequestBuilder) Rules()(*ItemItemRulesRequestBuilder) {
    return NewItemItemRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rulesets the rulesets property
func (m *ItemWithRepoItemRequestBuilder) Rulesets()(*ItemItemRulesetsRequestBuilder) {
    return NewItemItemRulesetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretScanning the secretScanning property
func (m *ItemWithRepoItemRequestBuilder) SecretScanning()(*ItemItemSecretScanningRequestBuilder) {
    return NewItemItemSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityAdvisories the securityAdvisories property
func (m *ItemWithRepoItemRequestBuilder) SecurityAdvisories()(*ItemItemSecurityAdvisoriesRequestBuilder) {
    return NewItemItemSecurityAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stargazers the stargazers property
func (m *ItemWithRepoItemRequestBuilder) Stargazers()(*ItemItemStargazersRequestBuilder) {
    return NewItemItemStargazersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stats the stats property
func (m *ItemWithRepoItemRequestBuilder) Stats()(*ItemItemStatsRequestBuilder) {
    return NewItemItemStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Statuses the statuses property
func (m *ItemWithRepoItemRequestBuilder) Statuses()(*ItemItemStatusesRequestBuilder) {
    return NewItemItemStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscribers the subscribers property
func (m *ItemWithRepoItemRequestBuilder) Subscribers()(*ItemItemSubscribersRequestBuilder) {
    return NewItemItemSubscribersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscription the subscription property
func (m *ItemWithRepoItemRequestBuilder) Subscription()(*ItemItemSubscriptionRequestBuilder) {
    return NewItemItemSubscriptionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags the tags property
func (m *ItemWithRepoItemRequestBuilder) Tags()(*ItemItemTagsRequestBuilder) {
    return NewItemItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tarball the tarball property
func (m *ItemWithRepoItemRequestBuilder) Tarball()(*ItemItemTarballRequestBuilder) {
    return NewItemItemTarballRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
func (m *ItemWithRepoItemRequestBuilder) Teams()(*ItemItemTeamsRequestBuilder) {
    return NewItemItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deleting a repository requires admin access. If OAuth is used, the `delete_repo` scope is required.If an organization owner has configured the organization to prevent members from deleting organization-ownedrepositories, you will get a `403 Forbidden` response.
func (m *ItemWithRepoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
func (m *ItemWithRepoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation **Note**: To edit a repository's topics, use the [Replace all repository topics](https://docs.github.com/rest/repos/repos#replace-all-repository-topics) endpoint.
func (m *ItemWithRepoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemWithRepoPatchRequestBodyable, requestConfiguration *ItemWithRepoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Topics the topics property
func (m *ItemWithRepoItemRequestBuilder) Topics()(*ItemItemTopicsRequestBuilder) {
    return NewItemItemTopicsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Traffic the traffic property
func (m *ItemWithRepoItemRequestBuilder) Traffic()(*ItemItemTrafficRequestBuilder) {
    return NewItemItemTrafficRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transfer the transfer property
func (m *ItemWithRepoItemRequestBuilder) Transfer()(*ItemItemTransferRequestBuilder) {
    return NewItemItemTransferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VulnerabilityAlerts the vulnerabilityAlerts property
func (m *ItemWithRepoItemRequestBuilder) VulnerabilityAlerts()(*ItemItemVulnerabilityAlertsRequestBuilder) {
    return NewItemItemVulnerabilityAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemWithRepoItemRequestBuilder) WithUrl(rawUrl string)(*ItemWithRepoItemRequestBuilder) {
    return NewItemWithRepoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Zipball the zipball property
func (m *ItemWithRepoItemRequestBuilder) Zipball()(*ItemItemZipballRequestBuilder) {
    return NewItemItemZipballRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
