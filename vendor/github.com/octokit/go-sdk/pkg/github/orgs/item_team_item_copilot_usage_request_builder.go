package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemTeamItemCopilotUsageRequestBuilder builds and executes requests for operations under \orgs\{org}\team\{team_slug}\copilot\usage
type ItemTeamItemCopilotUsageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamItemCopilotUsageRequestBuilderGetQueryParameters **Note**: This endpoint is in beta and is subject to change.You can use this endpoint to see a daily breakdown of aggregated usage metrics for Copilot completions and Copilot Chat in the IDEfor users within a team, with a further breakdown of suggestions, acceptances, and number of active users by editor and language for each day.See the response schema tab for detailed metrics definitions.The response contains metrics for the prior 28 days. Usage metrics are processed once per day for the previous day,and the response will only include data up until yesterday. In order for an end user to be counted towards these metrics,they must have telemetry enabled in their IDE.**Note**: This endpoint will only return results for a given day if the team had five or more members on that day.Copilot Business or Copilot Enterprise organization owners for the organization that contains this team,and owners and billing managers of their parent enterprises, can view Copilot usage metrics for a team.OAuth app tokens and personal access tokens (classic) need the `copilot`, `manage_billing:copilot`, `admin:org`, `admin:enterprise`, or `manage_billing:enterprise` scope to use this endpoint.
type ItemTeamItemCopilotUsageRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of days of metrics to display per page (max 28). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Show usage metrics since this date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:MM:SSZ`). Maximum value is 28 days ago.
    Since *string `uriparametername:"since"`
    // Show usage metrics until this date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:MM:SSZ`) and should not preceed the `since` date if it is passed.
    Until *string `uriparametername:"until"`
}
// NewItemTeamItemCopilotUsageRequestBuilderInternal instantiates a new ItemTeamItemCopilotUsageRequestBuilder and sets the default values.
func NewItemTeamItemCopilotUsageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamItemCopilotUsageRequestBuilder) {
    m := &ItemTeamItemCopilotUsageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/team/{team_slug}/copilot/usage{?page*,per_page*,since*,until*}", pathParameters),
    }
    return m
}
// NewItemTeamItemCopilotUsageRequestBuilder instantiates a new ItemTeamItemCopilotUsageRequestBuilder and sets the default values.
func NewItemTeamItemCopilotUsageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamItemCopilotUsageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamItemCopilotUsageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note**: This endpoint is in beta and is subject to change.You can use this endpoint to see a daily breakdown of aggregated usage metrics for Copilot completions and Copilot Chat in the IDEfor users within a team, with a further breakdown of suggestions, acceptances, and number of active users by editor and language for each day.See the response schema tab for detailed metrics definitions.The response contains metrics for the prior 28 days. Usage metrics are processed once per day for the previous day,and the response will only include data up until yesterday. In order for an end user to be counted towards these metrics,they must have telemetry enabled in their IDE.**Note**: This endpoint will only return results for a given day if the team had five or more members on that day.Copilot Business or Copilot Enterprise organization owners for the organization that contains this team,and owners and billing managers of their parent enterprises, can view Copilot usage metrics for a team.OAuth app tokens and personal access tokens (classic) need the `copilot`, `manage_billing:copilot`, `admin:org`, `admin:enterprise`, or `manage_billing:enterprise` scope to use this endpoint.
// returns a []CopilotUsageMetricsable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-usage#get-a-summary-of-copilot-usage-for-a-team
func (m *ItemTeamItemCopilotUsageRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamItemCopilotUsageRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CopilotUsageMetricsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCopilotUsageMetricsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CopilotUsageMetricsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CopilotUsageMetricsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation **Note**: This endpoint is in beta and is subject to change.You can use this endpoint to see a daily breakdown of aggregated usage metrics for Copilot completions and Copilot Chat in the IDEfor users within a team, with a further breakdown of suggestions, acceptances, and number of active users by editor and language for each day.See the response schema tab for detailed metrics definitions.The response contains metrics for the prior 28 days. Usage metrics are processed once per day for the previous day,and the response will only include data up until yesterday. In order for an end user to be counted towards these metrics,they must have telemetry enabled in their IDE.**Note**: This endpoint will only return results for a given day if the team had five or more members on that day.Copilot Business or Copilot Enterprise organization owners for the organization that contains this team,and owners and billing managers of their parent enterprises, can view Copilot usage metrics for a team.OAuth app tokens and personal access tokens (classic) need the `copilot`, `manage_billing:copilot`, `admin:org`, `admin:enterprise`, or `manage_billing:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemTeamItemCopilotUsageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTeamItemCopilotUsageRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamItemCopilotUsageRequestBuilder when successful
func (m *ItemTeamItemCopilotUsageRequestBuilder) WithUrl(rawUrl string)(*ItemTeamItemCopilotUsageRequestBuilder) {
    return NewItemTeamItemCopilotUsageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
