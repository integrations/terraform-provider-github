package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemOwnerItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}
type ItemOwnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Actions the actions property
// returns a *ItemItemActionsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Actions()(*ItemItemActionsRequestBuilder) {
    return NewItemItemActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Activity the activity property
// returns a *ItemItemActivityRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Activity()(*ItemItemActivityRequestBuilder) {
    return NewItemItemActivityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignees the assignees property
// returns a *ItemItemAssigneesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Assignees()(*ItemItemAssigneesRequestBuilder) {
    return NewItemItemAssigneesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Autolinks the autolinks property
// returns a *ItemItemAutolinksRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Autolinks()(*ItemItemAutolinksRequestBuilder) {
    return NewItemItemAutolinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AutomatedSecurityFixes the automatedSecurityFixes property
// returns a *ItemItemAutomatedSecurityFixesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) AutomatedSecurityFixes()(*ItemItemAutomatedSecurityFixesRequestBuilder) {
    return NewItemItemAutomatedSecurityFixesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Branches the branches property
// returns a *ItemItemBranchesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Branches()(*ItemItemBranchesRequestBuilder) {
    return NewItemItemBranchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckRuns the checkRuns property
// returns a *ItemItemCheckRunsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) CheckRuns()(*ItemItemCheckRunsRequestBuilder) {
    return NewItemItemCheckRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckSuites the checkSuites property
// returns a *ItemItemCheckSuitesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) CheckSuites()(*ItemItemCheckSuitesRequestBuilder) {
    return NewItemItemCheckSuitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codeowners the codeowners property
// returns a *ItemItemCodeownersRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Codeowners()(*ItemItemCodeownersRequestBuilder) {
    return NewItemItemCodeownersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CodeScanning the codeScanning property
// returns a *ItemItemCodeScanningRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) CodeScanning()(*ItemItemCodeScanningRequestBuilder) {
    return NewItemItemCodeScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codespaces the codespaces property
// returns a *ItemItemCodespacesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Codespaces()(*ItemItemCodespacesRequestBuilder) {
    return NewItemItemCodespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Collaborators the collaborators property
// returns a *ItemItemCollaboratorsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Collaborators()(*ItemItemCollaboratorsRequestBuilder) {
    return NewItemItemCollaboratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
// returns a *ItemItemCommentsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Comments()(*ItemItemCommentsRequestBuilder) {
    return NewItemItemCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commits the commits property
// returns a *ItemItemCommitsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Commits()(*ItemItemCommitsRequestBuilder) {
    return NewItemItemCommitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Community the community property
// returns a *ItemItemCommunityRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Community()(*ItemItemCommunityRequestBuilder) {
    return NewItemItemCommunityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Compare the compare property
// returns a *ItemItemCompareRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Compare()(*ItemItemCompareRequestBuilder) {
    return NewItemItemCompareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOwnerItemRequestBuilderInternal instantiates a new ItemOwnerItemRequestBuilder and sets the default values.
func NewItemOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnerItemRequestBuilder) {
    m := &ItemOwnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}", pathParameters),
    }
    return m
}
// NewItemOwnerItemRequestBuilder instantiates a new ItemOwnerItemRequestBuilder and sets the default values.
func NewItemOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Contents the contents property
// returns a *ItemItemContentsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Contents()(*ItemItemContentsRequestBuilder) {
    return NewItemItemContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contributors the contributors property
// returns a *ItemItemContributorsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Contributors()(*ItemItemContributorsRequestBuilder) {
    return NewItemItemContributorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deleting a repository requires admin access.If an organization owner has configured the organization to prevent members from deleting organization-ownedrepositories, you will get a `403 Forbidden` response.OAuth app tokens and personal access tokens (classic) need the `delete_repo` scope to use this endpoint.
// returns a ItemItemOwner403Error error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#delete-a-repository
func (m *ItemOwnerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": CreateItemItemOwner403ErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Dependabot the dependabot property
// returns a *ItemItemDependabotRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Dependabot()(*ItemItemDependabotRequestBuilder) {
    return NewItemItemDependabotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DependencyGraph the dependencyGraph property
// returns a *ItemItemDependencyGraphRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) DependencyGraph()(*ItemItemDependencyGraphRequestBuilder) {
    return NewItemItemDependencyGraphRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deployments the deployments property
// returns a *ItemItemDeploymentsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Deployments()(*ItemItemDeploymentsRequestBuilder) {
    return NewItemItemDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dispatches the dispatches property
// returns a *ItemItemDispatchesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Dispatches()(*ItemItemDispatchesRequestBuilder) {
    return NewItemItemDispatchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Environments the environments property
// returns a *ItemItemEnvironmentsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Environments()(*ItemItemEnvironmentsRequestBuilder) {
    return NewItemItemEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
// returns a *ItemItemEventsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Events()(*ItemItemEventsRequestBuilder) {
    return NewItemItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forks the forks property
// returns a *ItemItemForksRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Forks()(*ItemItemForksRequestBuilder) {
    return NewItemItemForksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Generate the generate property
// returns a *ItemItemGenerateRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Generate()(*ItemItemGenerateRequestBuilder) {
    return NewItemItemGenerateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// returns a FullRepositoryable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#get-a-repository
func (m *ItemOwnerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FullRepositoryable, error) {
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
// returns a *ItemItemGitRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Git()(*ItemItemGitRequestBuilder) {
    return NewItemItemGitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hooks the hooks property
// returns a *ItemItemHooksRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Hooks()(*ItemItemHooksRequestBuilder) {
    return NewItemItemHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportEscaped the import property
// returns a *ItemItemImportRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) ImportEscaped()(*ItemItemImportRequestBuilder) {
    return NewItemItemImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
// returns a *ItemItemInstallationRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Installation()(*ItemItemInstallationRequestBuilder) {
    return NewItemItemInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InteractionLimits the interactionLimits property
// returns a *ItemItemInteractionLimitsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) InteractionLimits()(*ItemItemInteractionLimitsRequestBuilder) {
    return NewItemItemInteractionLimitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Invitations the invitations property
// returns a *ItemItemInvitationsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Invitations()(*ItemItemInvitationsRequestBuilder) {
    return NewItemItemInvitationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *ItemItemIssuesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Issues()(*ItemItemIssuesRequestBuilder) {
    return NewItemItemIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keys the keys property
// returns a *ItemItemKeysRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Keys()(*ItemItemKeysRequestBuilder) {
    return NewItemItemKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Labels the labels property
// returns a *ItemItemLabelsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Labels()(*ItemItemLabelsRequestBuilder) {
    return NewItemItemLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Languages the languages property
// returns a *ItemItemLanguagesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Languages()(*ItemItemLanguagesRequestBuilder) {
    return NewItemItemLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// License the license property
// returns a *ItemItemLicenseRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) License()(*ItemItemLicenseRequestBuilder) {
    return NewItemItemLicenseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Merges the merges property
// returns a *ItemItemMergesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Merges()(*ItemItemMergesRequestBuilder) {
    return NewItemItemMergesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MergeUpstream the mergeUpstream property
// returns a *ItemItemMergeUpstreamRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) MergeUpstream()(*ItemItemMergeUpstreamRequestBuilder) {
    return NewItemItemMergeUpstreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Milestones the milestones property
// returns a *ItemItemMilestonesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Milestones()(*ItemItemMilestonesRequestBuilder) {
    return NewItemItemMilestonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
// returns a *ItemItemNotificationsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Notifications()(*ItemItemNotificationsRequestBuilder) {
    return NewItemItemNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages the pages property
// returns a *ItemItemPagesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Pages()(*ItemItemPagesRequestBuilder) {
    return NewItemItemPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch **Note**: To edit a repository's topics, use the [Replace all repository topics](https://docs.github.com/rest/repos/repos#replace-all-repository-topics) endpoint.
// returns a FullRepositoryable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#update-a-repository
func (m *ItemOwnerItemRequestBuilder) Patch(ctx context.Context, body ItemItemOwnerPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FullRepositoryable, error) {
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
// returns a *ItemItemPrivateVulnerabilityReportingRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) PrivateVulnerabilityReporting()(*ItemItemPrivateVulnerabilityReportingRequestBuilder) {
    return NewItemItemPrivateVulnerabilityReportingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ItemItemProjectsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Projects()(*ItemItemProjectsRequestBuilder) {
    return NewItemItemProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Properties the properties property
// returns a *ItemItemPropertiesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Properties()(*ItemItemPropertiesRequestBuilder) {
    return NewItemItemPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pulls the pulls property
// returns a *ItemItemPullsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Pulls()(*ItemItemPullsRequestBuilder) {
    return NewItemItemPullsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Readme the readme property
// returns a *ItemItemReadmeRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Readme()(*ItemItemReadmeRequestBuilder) {
    return NewItemItemReadmeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Releases the releases property
// returns a *ItemItemReleasesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Releases()(*ItemItemReleasesRequestBuilder) {
    return NewItemItemReleasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rules the rules property
// returns a *ItemItemRulesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Rules()(*ItemItemRulesRequestBuilder) {
    return NewItemItemRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rulesets the rulesets property
// returns a *ItemItemRulesetsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Rulesets()(*ItemItemRulesetsRequestBuilder) {
    return NewItemItemRulesetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretScanning the secretScanning property
// returns a *ItemItemSecretScanningRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) SecretScanning()(*ItemItemSecretScanningRequestBuilder) {
    return NewItemItemSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityAdvisories the securityAdvisories property
// returns a *ItemItemSecurityAdvisoriesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) SecurityAdvisories()(*ItemItemSecurityAdvisoriesRequestBuilder) {
    return NewItemItemSecurityAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stargazers the stargazers property
// returns a *ItemItemStargazersRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Stargazers()(*ItemItemStargazersRequestBuilder) {
    return NewItemItemStargazersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stats the stats property
// returns a *ItemItemStatsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Stats()(*ItemItemStatsRequestBuilder) {
    return NewItemItemStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Statuses the statuses property
// returns a *ItemItemStatusesRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Statuses()(*ItemItemStatusesRequestBuilder) {
    return NewItemItemStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscribers the subscribers property
// returns a *ItemItemSubscribersRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Subscribers()(*ItemItemSubscribersRequestBuilder) {
    return NewItemItemSubscribersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscription the subscription property
// returns a *ItemItemSubscriptionRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Subscription()(*ItemItemSubscriptionRequestBuilder) {
    return NewItemItemSubscriptionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags the tags property
// returns a *ItemItemTagsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Tags()(*ItemItemTagsRequestBuilder) {
    return NewItemItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tarball the tarball property
// returns a *ItemItemTarballRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Tarball()(*ItemItemTarballRequestBuilder) {
    return NewItemItemTarballRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
// returns a *ItemItemTeamsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Teams()(*ItemItemTeamsRequestBuilder) {
    return NewItemItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deleting a repository requires admin access.If an organization owner has configured the organization to prevent members from deleting organization-ownedrepositories, you will get a `403 Forbidden` response.OAuth app tokens and personal access tokens (classic) need the `delete_repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemOwnerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// returns a *RequestInformation when successful
func (m *ItemOwnerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation **Note**: To edit a repository's topics, use the [Replace all repository topics](https://docs.github.com/rest/repos/repos#replace-all-repository-topics) endpoint.
// returns a *RequestInformation when successful
func (m *ItemOwnerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemOwnerPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Topics the topics property
// returns a *ItemItemTopicsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Topics()(*ItemItemTopicsRequestBuilder) {
    return NewItemItemTopicsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Traffic the traffic property
// returns a *ItemItemTrafficRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Traffic()(*ItemItemTrafficRequestBuilder) {
    return NewItemItemTrafficRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transfer the transfer property
// returns a *ItemItemTransferRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Transfer()(*ItemItemTransferRequestBuilder) {
    return NewItemItemTransferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VulnerabilityAlerts the vulnerabilityAlerts property
// returns a *ItemItemVulnerabilityAlertsRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) VulnerabilityAlerts()(*ItemItemVulnerabilityAlertsRequestBuilder) {
    return NewItemItemVulnerabilityAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOwnerItemRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) WithUrl(rawUrl string)(*ItemOwnerItemRequestBuilder) {
    return NewItemOwnerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Zipball the zipball property
// returns a *ItemItemZipballRequestBuilder when successful
func (m *ItemOwnerItemRequestBuilder) Zipball()(*ItemItemZipballRequestBuilder) {
    return NewItemItemZipballRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
