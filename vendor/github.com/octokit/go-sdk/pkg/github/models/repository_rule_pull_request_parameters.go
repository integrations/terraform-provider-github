package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRulePullRequest_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // New, reviewable commits pushed will dismiss previous pull request review approvals.
    dismiss_stale_reviews_on_push *bool
    // Require an approving review in pull requests that modify files that have a designated code owner.
    require_code_owner_review *bool
    // Whether the most recent reviewable push must be approved by someone other than the person who pushed it.
    require_last_push_approval *bool
    // The number of approving reviews that are required before a pull request can be merged.
    required_approving_review_count *int32
    // All conversations on code must be resolved before a pull request can be merged.
    required_review_thread_resolution *bool
}
// NewRepositoryRulePullRequest_parameters instantiates a new RepositoryRulePullRequest_parameters and sets the default values.
func NewRepositoryRulePullRequest_parameters()(*RepositoryRulePullRequest_parameters) {
    m := &RepositoryRulePullRequest_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRulePullRequest_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRulePullRequest_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRulePullRequest_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRulePullRequest_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDismissStaleReviewsOnPush gets the dismiss_stale_reviews_on_push property value. New, reviewable commits pushed will dismiss previous pull request review approvals.
// returns a *bool when successful
func (m *RepositoryRulePullRequest_parameters) GetDismissStaleReviewsOnPush()(*bool) {
    return m.dismiss_stale_reviews_on_push
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRulePullRequest_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dismiss_stale_reviews_on_push"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissStaleReviewsOnPush(val)
        }
        return nil
    }
    res["require_code_owner_review"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCodeOwnerReview(val)
        }
        return nil
    }
    res["require_last_push_approval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireLastPushApproval(val)
        }
        return nil
    }
    res["required_approving_review_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredApprovingReviewCount(val)
        }
        return nil
    }
    res["required_review_thread_resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredReviewThreadResolution(val)
        }
        return nil
    }
    return res
}
// GetRequireCodeOwnerReview gets the require_code_owner_review property value. Require an approving review in pull requests that modify files that have a designated code owner.
// returns a *bool when successful
func (m *RepositoryRulePullRequest_parameters) GetRequireCodeOwnerReview()(*bool) {
    return m.require_code_owner_review
}
// GetRequiredApprovingReviewCount gets the required_approving_review_count property value. The number of approving reviews that are required before a pull request can be merged.
// returns a *int32 when successful
func (m *RepositoryRulePullRequest_parameters) GetRequiredApprovingReviewCount()(*int32) {
    return m.required_approving_review_count
}
// GetRequiredReviewThreadResolution gets the required_review_thread_resolution property value. All conversations on code must be resolved before a pull request can be merged.
// returns a *bool when successful
func (m *RepositoryRulePullRequest_parameters) GetRequiredReviewThreadResolution()(*bool) {
    return m.required_review_thread_resolution
}
// GetRequireLastPushApproval gets the require_last_push_approval property value. Whether the most recent reviewable push must be approved by someone other than the person who pushed it.
// returns a *bool when successful
func (m *RepositoryRulePullRequest_parameters) GetRequireLastPushApproval()(*bool) {
    return m.require_last_push_approval
}
// Serialize serializes information the current object
func (m *RepositoryRulePullRequest_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("dismiss_stale_reviews_on_push", m.GetDismissStaleReviewsOnPush())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("required_approving_review_count", m.GetRequiredApprovingReviewCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required_review_thread_resolution", m.GetRequiredReviewThreadResolution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("require_code_owner_review", m.GetRequireCodeOwnerReview())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("require_last_push_approval", m.GetRequireLastPushApproval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryRulePullRequest_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDismissStaleReviewsOnPush sets the dismiss_stale_reviews_on_push property value. New, reviewable commits pushed will dismiss previous pull request review approvals.
func (m *RepositoryRulePullRequest_parameters) SetDismissStaleReviewsOnPush(value *bool)() {
    m.dismiss_stale_reviews_on_push = value
}
// SetRequireCodeOwnerReview sets the require_code_owner_review property value. Require an approving review in pull requests that modify files that have a designated code owner.
func (m *RepositoryRulePullRequest_parameters) SetRequireCodeOwnerReview(value *bool)() {
    m.require_code_owner_review = value
}
// SetRequiredApprovingReviewCount sets the required_approving_review_count property value. The number of approving reviews that are required before a pull request can be merged.
func (m *RepositoryRulePullRequest_parameters) SetRequiredApprovingReviewCount(value *int32)() {
    m.required_approving_review_count = value
}
// SetRequiredReviewThreadResolution sets the required_review_thread_resolution property value. All conversations on code must be resolved before a pull request can be merged.
func (m *RepositoryRulePullRequest_parameters) SetRequiredReviewThreadResolution(value *bool)() {
    m.required_review_thread_resolution = value
}
// SetRequireLastPushApproval sets the require_last_push_approval property value. Whether the most recent reviewable push must be approved by someone other than the person who pushed it.
func (m *RepositoryRulePullRequest_parameters) SetRequireLastPushApproval(value *bool)() {
    m.require_last_push_approval = value
}
type RepositoryRulePullRequest_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDismissStaleReviewsOnPush()(*bool)
    GetRequireCodeOwnerReview()(*bool)
    GetRequiredApprovingReviewCount()(*int32)
    GetRequiredReviewThreadResolution()(*bool)
    GetRequireLastPushApproval()(*bool)
    SetDismissStaleReviewsOnPush(value *bool)()
    SetRequireCodeOwnerReview(value *bool)()
    SetRequiredApprovingReviewCount(value *int32)()
    SetRequiredReviewThreadResolution(value *bool)()
    SetRequireLastPushApproval(value *bool)()
}
