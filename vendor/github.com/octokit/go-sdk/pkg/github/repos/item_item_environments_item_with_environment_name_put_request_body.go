package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody struct {
    // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
    deployment_branch_policy i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentBranchPolicySettingsable
    // Whether or not a user who created the job is prevented from approving their own job.
    prevent_self_review *bool
    // The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
    reviewers []ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable
    // The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
    wait_timer *int32
}
// NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody instantiates a new ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody and sets the default values.
func NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody()(*ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) {
    m := &ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody{
    }
    return m
}
// CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody(), nil
}
// GetDeploymentBranchPolicy gets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
// returns a DeploymentBranchPolicySettingsable when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetDeploymentBranchPolicy()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentBranchPolicySettingsable) {
    return m.deployment_branch_policy
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deployment_branch_policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDeploymentBranchPolicySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentBranchPolicy(val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentBranchPolicySettingsable))
        }
        return nil
    }
    res["prevent_self_review"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreventSelfReview(val)
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable)
                }
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["wait_timer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitTimer(val)
        }
        return nil
    }
    return res
}
// GetPreventSelfReview gets the prevent_self_review property value. Whether or not a user who created the job is prevented from approving their own job.
// returns a *bool when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetPreventSelfReview()(*bool) {
    return m.prevent_self_review
}
// GetReviewers gets the reviewers property value. The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
// returns a []ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetReviewers()([]ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable) {
    return m.reviewers
}
// GetWaitTimer gets the wait_timer property value. The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
// returns a *int32 when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetWaitTimer()(*int32) {
    return m.wait_timer
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deployment_branch_policy", m.GetDeploymentBranchPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("prevent_self_review", m.GetPreventSelfReview())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wait_timer", m.GetWaitTimer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentBranchPolicy sets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetDeploymentBranchPolicy(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentBranchPolicySettingsable)() {
    m.deployment_branch_policy = value
}
// SetPreventSelfReview sets the prevent_self_review property value. Whether or not a user who created the job is prevented from approving their own job.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetPreventSelfReview(value *bool)() {
    m.prevent_self_review = value
}
// SetReviewers sets the reviewers property value. The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetReviewers(value []ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable)() {
    m.reviewers = value
}
// SetWaitTimer sets the wait_timer property value. The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetWaitTimer(value *int32)() {
    m.wait_timer = value
}
type ItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentBranchPolicy()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentBranchPolicySettingsable)
    GetPreventSelfReview()(*bool)
    GetReviewers()([]ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable)
    GetWaitTimer()(*int32)
    SetDeploymentBranchPolicy(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentBranchPolicySettingsable)()
    SetPreventSelfReview(value *bool)()
    SetReviewers(value []ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable)()
    SetWaitTimer(value *int32)()
}
