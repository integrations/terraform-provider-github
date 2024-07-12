package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\deployment_protection_rule
type ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Deployment_protection_rulePostRequestBody composed type wrapper for classes i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
type Deployment_protection_rulePostRequestBody struct {
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
    deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
    deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
    reviewCustomGatesCommentRequired i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
    reviewCustomGatesStateRequired i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
}
// NewDeployment_protection_rulePostRequestBody instantiates a new Deployment_protection_rulePostRequestBody and sets the default values.
func NewDeployment_protection_rulePostRequestBody()(*Deployment_protection_rulePostRequestBody) {
    m := &Deployment_protection_rulePostRequestBody{
    }
    return m
}
// CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeployment_protection_rulePostRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable); ok {
                result.SetReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable); ok {
                result.SetReviewCustomGatesStateRequired(cast)
            }
        }
    }
    return result, nil
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
// returns a ReviewCustomGatesCommentRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
// returns a ReviewCustomGatesStateRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Deployment_protection_rulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Deployment_protection_rulePostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
// returns a ReviewCustomGatesCommentRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesCommentRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable) {
    return m.reviewCustomGatesCommentRequired
}
// GetReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
// returns a ReviewCustomGatesStateRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesStateRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable) {
    return m.reviewCustomGatesStateRequired
}
// Serialize serializes information the current object
func (m *Deployment_protection_rulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired = value
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired = value
}
// SetReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesCommentRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable)() {
    m.reviewCustomGatesCommentRequired = value
}
// SetReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesStateRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable)() {
    m.reviewCustomGatesStateRequired = value
}
type Deployment_protection_rulePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable)
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable)
    GetReviewCustomGatesCommentRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable)
    GetReviewCustomGatesStateRequired()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable)
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable)()
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable)()
    SetReviewCustomGatesCommentRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesCommentRequiredable)()
    SetReviewCustomGatesStateRequired(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ReviewCustomGatesStateRequiredable)()
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal instantiates a new ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    m := &ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/deployment_protection_rule", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder instantiates a new ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."**Note:** GitHub Apps can only review their own custom deployment protection rules.To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#review-custom-deployment-protection-rules-for-a-workflow-run
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) Post(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."**Note:** GitHub Apps can only review their own custom deployment protection rules.To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) ToPostRequestInformation(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder when successful
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
