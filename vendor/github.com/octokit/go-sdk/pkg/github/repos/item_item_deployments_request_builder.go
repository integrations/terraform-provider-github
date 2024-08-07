package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemDeploymentsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\deployments
type ItemItemDeploymentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDeploymentsRequestBuilderGetQueryParameters simple filtering of deployments is available via query parameters:
type ItemItemDeploymentsRequestBuilderGetQueryParameters struct {
    // The name of the environment that was deployed to (e.g., `staging` or `production`).
    Environment *string `uriparametername:"environment"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The name of the ref. This can be a branch, tag, or SHA.
    Ref *string `uriparametername:"ref"`
    // The SHA recorded at creation time.
    Sha *string `uriparametername:"sha"`
    // The name of the task for the deployment (e.g., `deploy` or `deploy:migrations`).
    Task *string `uriparametername:"task"`
}
// ByDeployment_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.deployments.item collection
// returns a *ItemItemDeploymentsWithDeployment_ItemRequestBuilder when successful
func (m *ItemItemDeploymentsRequestBuilder) ByDeployment_id(deployment_id int32)(*ItemItemDeploymentsWithDeployment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["deployment_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(deployment_id), 10)
    return NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemDeploymentsRequestBuilderInternal instantiates a new ItemItemDeploymentsRequestBuilder and sets the default values.
func NewItemItemDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsRequestBuilder) {
    m := &ItemItemDeploymentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/deployments{?environment*,page*,per_page*,ref*,sha*,task*}", pathParameters),
    }
    return m
}
// NewItemItemDeploymentsRequestBuilder instantiates a new ItemItemDeploymentsRequestBuilder and sets the default values.
func NewItemItemDeploymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get simple filtering of deployments is available via query parameters:
// returns a []Deploymentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/deployments#list-deployments
func (m *ItemItemDeploymentsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemDeploymentsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Deploymentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDeploymentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Deploymentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Deploymentable)
        }
    }
    return val, nil
}
// Post deployments offer a few configurable parameters with certain defaults.The `ref` parameter can be any named branch, tag, or SHA. At GitHub we often deploy branches and verify thembefore we merge a pull request.The `environment` parameter allows deployments to be issued to different runtime environments. Teams often havemultiple environments for verifying their applications, such as `production`, `staging`, and `qa`. This parametermakes it easier to track which environments have requested deployments. The default environment is `production`.The `auto_merge` parameter is used to ensure that the requested ref is not behind the repository's default branch. Ifthe ref _is_ behind the default branch for the repository, we will attempt to merge it for you. If the merge succeeds,the API will return a successful merge commit. If merge conflicts prevent the merge from succeeding, the API willreturn a failure response.By default, [commit statuses](https://docs.github.com/rest/commits/statuses) for every submitted context must be in a `success`state. The `required_contexts` parameter allows you to specify a subset of contexts that must be `success`, or tospecify contexts that have not yet been submitted. You are not required to use commit statuses to deploy. If you donot require any contexts or create any commit statuses, the deployment will always succeed.The `payload` parameter is available for any extra information that a deployment system might need. It is a JSON textfield that will be passed on when a deployment event is dispatched.The `task` parameter is used by the deployment system to allow different execution paths. In the web world this mightbe `deploy:migrations` to run schema changes on the system. In the compiled world this could be a flag to compile anapplication with debugging enabled.Merged branch response:You will see this response when GitHub automatically merges the base branch into the topic branch instead of creatinga deployment. This auto-merge happens when:*   Auto-merge option is enabled in the repository*   Topic branch does not include the latest changes on the base branch, which is `master` in the response example*   There are no merge conflictsIf there are no new commits in the base branch, a new request to create a deployment should give a successfulresponse.Merge conflict response:This error happens when the `auto_merge` option is enabled and when the default branch (in this case `master`), can'tbe merged into the branch that's being deployed (in this case `topic-branch`), due to merge conflicts.Failed commit status checks:This error happens when the `required_contexts` parameter indicates that one or more contexts need to have a `success`status for the commit to be deployed, but one or more of the required contexts do not have a state of `success`.OAuth app tokens and personal access tokens (classic) need the `repo` or `repo_deployment` scope to use this endpoint.
// returns a Deploymentable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/deployments#create-a-deployment
func (m *ItemItemDeploymentsRequestBuilder) Post(ctx context.Context, body ItemItemDeploymentsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Deploymentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Deploymentable), nil
}
// ToGetRequestInformation simple filtering of deployments is available via query parameters:
// returns a *RequestInformation when successful
func (m *ItemItemDeploymentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemDeploymentsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation deployments offer a few configurable parameters with certain defaults.The `ref` parameter can be any named branch, tag, or SHA. At GitHub we often deploy branches and verify thembefore we merge a pull request.The `environment` parameter allows deployments to be issued to different runtime environments. Teams often havemultiple environments for verifying their applications, such as `production`, `staging`, and `qa`. This parametermakes it easier to track which environments have requested deployments. The default environment is `production`.The `auto_merge` parameter is used to ensure that the requested ref is not behind the repository's default branch. Ifthe ref _is_ behind the default branch for the repository, we will attempt to merge it for you. If the merge succeeds,the API will return a successful merge commit. If merge conflicts prevent the merge from succeeding, the API willreturn a failure response.By default, [commit statuses](https://docs.github.com/rest/commits/statuses) for every submitted context must be in a `success`state. The `required_contexts` parameter allows you to specify a subset of contexts that must be `success`, or tospecify contexts that have not yet been submitted. You are not required to use commit statuses to deploy. If you donot require any contexts or create any commit statuses, the deployment will always succeed.The `payload` parameter is available for any extra information that a deployment system might need. It is a JSON textfield that will be passed on when a deployment event is dispatched.The `task` parameter is used by the deployment system to allow different execution paths. In the web world this mightbe `deploy:migrations` to run schema changes on the system. In the compiled world this could be a flag to compile anapplication with debugging enabled.Merged branch response:You will see this response when GitHub automatically merges the base branch into the topic branch instead of creatinga deployment. This auto-merge happens when:*   Auto-merge option is enabled in the repository*   Topic branch does not include the latest changes on the base branch, which is `master` in the response example*   There are no merge conflictsIf there are no new commits in the base branch, a new request to create a deployment should give a successfulresponse.Merge conflict response:This error happens when the `auto_merge` option is enabled and when the default branch (in this case `master`), can'tbe merged into the branch that's being deployed (in this case `topic-branch`), due to merge conflicts.Failed commit status checks:This error happens when the `required_contexts` parameter indicates that one or more contexts need to have a `success`status for the commit to be deployed, but one or more of the required contexts do not have a state of `success`.OAuth app tokens and personal access tokens (classic) need the `repo` or `repo_deployment` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemDeploymentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemDeploymentsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemDeploymentsRequestBuilder when successful
func (m *ItemItemDeploymentsRequestBuilder) WithUrl(rawUrl string)(*ItemItemDeploymentsRequestBuilder) {
    return NewItemItemDeploymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
