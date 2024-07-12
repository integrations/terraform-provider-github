package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\environments\{environment_name}\deployment_protection_rules\{protection_rule_id}
type ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/environments/{environment_name}/deployment_protection_rules/{protection_rule_id}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete disables a custom deployment protection rule for an environment.The authenticated user must have admin or owner permissions to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#disable-a-custom-protection-rule-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get gets an enabled custom deployment protection rule for an environment. Anyone with read access to the repository can use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see [`GET /apps/{app_slug}`](https://docs.github.com/rest/apps/apps#get-an-app).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a DeploymentProtectionRuleable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#get-a-custom-deployment-protection-rule
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDeploymentProtectionRuleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable), nil
}
// ToDeleteRequestInformation disables a custom deployment protection rule for an environment.The authenticated user must have admin or owner permissions to the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation gets an enabled custom deployment protection rule for an environment. Anyone with read access to the repository can use this endpoint. For more information about environments, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see [`GET /apps/{app_slug}`](https://docs.github.com/rest/apps/apps#get-an-app).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesWithProtection_rule_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
