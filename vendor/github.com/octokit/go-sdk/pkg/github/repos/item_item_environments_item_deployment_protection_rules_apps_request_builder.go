package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\environments\{environment_name}\deployment_protection_rules\apps
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/environments/{environment_name}/deployment_protection_rules/apps{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/protection-rules#list-custom-deployment-rule-integrations-available-for-an-environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters])(ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable), nil
}
// ToGetRequestInformation gets all custom deployment protection rule integrations that are available for an environment. Anyone with read access to the repository can use this endpoint.For more information about environments, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."For more information about the app that is providing this custom deployment rule, see "[GET an app](https://docs.github.com/rest/apps/apps#get-an-app)".OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) WithUrl(rawUrl string)(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
