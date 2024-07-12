package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i6248a706a395d319f3e66611fac0b2654fe48b61e0e7152b109cc83e1d984972 "github.com/octokit/go-sdk/pkg/github/user/memberships/orgs"
)

// MembershipsOrgsRequestBuilder builds and executes requests for operations under \user\memberships\orgs
type MembershipsOrgsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MembershipsOrgsRequestBuilderGetQueryParameters lists all of the authenticated user's organization memberships.
type MembershipsOrgsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Indicates the state of the memberships to return. If not specified, the API returns both active and pending memberships.
    State *i6248a706a395d319f3e66611fac0b2654fe48b61e0e7152b109cc83e1d984972.GetStateQueryParameterType `uriparametername:"state"`
}
// ByOrg gets an item from the github.com/octokit/go-sdk/pkg/github.user.memberships.orgs.item collection
// returns a *MembershipsOrgsWithOrgItemRequestBuilder when successful
func (m *MembershipsOrgsRequestBuilder) ByOrg(org string)(*MembershipsOrgsWithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if org != "" {
        urlTplParams["org"] = org
    }
    return NewMembershipsOrgsWithOrgItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMembershipsOrgsRequestBuilderInternal instantiates a new MembershipsOrgsRequestBuilder and sets the default values.
func NewMembershipsOrgsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsOrgsRequestBuilder) {
    m := &MembershipsOrgsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/memberships/orgs{?page*,per_page*,state*}", pathParameters),
    }
    return m
}
// NewMembershipsOrgsRequestBuilder instantiates a new MembershipsOrgsRequestBuilder and sets the default values.
func NewMembershipsOrgsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsOrgsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembershipsOrgsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all of the authenticated user's organization memberships.
// returns a []OrgMembershipable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#list-organization-memberships-for-the-authenticated-user
func (m *MembershipsOrgsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[MembershipsOrgsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgMembershipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrgMembershipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgMembershipable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgMembershipable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all of the authenticated user's organization memberships.
// returns a *RequestInformation when successful
func (m *MembershipsOrgsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[MembershipsOrgsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MembershipsOrgsRequestBuilder when successful
func (m *MembershipsOrgsRequestBuilder) WithUrl(rawUrl string)(*MembershipsOrgsRequestBuilder) {
    return NewMembershipsOrgsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
