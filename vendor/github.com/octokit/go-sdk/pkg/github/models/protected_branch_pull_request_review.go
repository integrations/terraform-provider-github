package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchPullRequestReview protected Branch Pull Request Review
type ProtectedBranchPullRequestReview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Allow specific users, teams, or apps to bypass pull request requirements.
    bypass_pull_request_allowances ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable
    // The dismiss_stale_reviews property
    dismiss_stale_reviews *bool
    // The dismissal_restrictions property
    dismissal_restrictions ProtectedBranchPullRequestReview_dismissal_restrictionsable
    // The require_code_owner_reviews property
    require_code_owner_reviews *bool
    // Whether the most recent push must be approved by someone other than the person who pushed it.
    require_last_push_approval *bool
    // The required_approving_review_count property
    required_approving_review_count *int32
    // The url property
    url *string
}
// NewProtectedBranchPullRequestReview instantiates a new ProtectedBranchPullRequestReview and sets the default values.
func NewProtectedBranchPullRequestReview()(*ProtectedBranchPullRequestReview) {
    m := &ProtectedBranchPullRequestReview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedBranchPullRequestReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedBranchPullRequestReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranchPullRequestReview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtectedBranchPullRequestReview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassPullRequestAllowances gets the bypass_pull_request_allowances property value. Allow specific users, teams, or apps to bypass pull request requirements.
// returns a ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable when successful
func (m *ProtectedBranchPullRequestReview) GetBypassPullRequestAllowances()(ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable) {
    return m.bypass_pull_request_allowances
}
// GetDismissalRestrictions gets the dismissal_restrictions property value. The dismissal_restrictions property
// returns a ProtectedBranchPullRequestReview_dismissal_restrictionsable when successful
func (m *ProtectedBranchPullRequestReview) GetDismissalRestrictions()(ProtectedBranchPullRequestReview_dismissal_restrictionsable) {
    return m.dismissal_restrictions
}
// GetDismissStaleReviews gets the dismiss_stale_reviews property value. The dismiss_stale_reviews property
// returns a *bool when successful
func (m *ProtectedBranchPullRequestReview) GetDismissStaleReviews()(*bool) {
    return m.dismiss_stale_reviews
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedBranchPullRequestReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_pull_request_allowances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranchPullRequestReview_bypass_pull_request_allowancesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassPullRequestAllowances(val.(ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable))
        }
        return nil
    }
    res["dismiss_stale_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissStaleReviews(val)
        }
        return nil
    }
    res["dismissal_restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranchPullRequestReview_dismissal_restrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissalRestrictions(val.(ProtectedBranchPullRequestReview_dismissal_restrictionsable))
        }
        return nil
    }
    res["require_code_owner_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCodeOwnerReviews(val)
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
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetRequireCodeOwnerReviews gets the require_code_owner_reviews property value. The require_code_owner_reviews property
// returns a *bool when successful
func (m *ProtectedBranchPullRequestReview) GetRequireCodeOwnerReviews()(*bool) {
    return m.require_code_owner_reviews
}
// GetRequiredApprovingReviewCount gets the required_approving_review_count property value. The required_approving_review_count property
// returns a *int32 when successful
func (m *ProtectedBranchPullRequestReview) GetRequiredApprovingReviewCount()(*int32) {
    return m.required_approving_review_count
}
// GetRequireLastPushApproval gets the require_last_push_approval property value. Whether the most recent push must be approved by someone other than the person who pushed it.
// returns a *bool when successful
func (m *ProtectedBranchPullRequestReview) GetRequireLastPushApproval()(*bool) {
    return m.require_last_push_approval
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ProtectedBranchPullRequestReview) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ProtectedBranchPullRequestReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bypass_pull_request_allowances", m.GetBypassPullRequestAllowances())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dismissal_restrictions", m.GetDismissalRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dismiss_stale_reviews", m.GetDismissStaleReviews())
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
        err := writer.WriteBoolValue("require_code_owner_reviews", m.GetRequireCodeOwnerReviews())
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
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *ProtectedBranchPullRequestReview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassPullRequestAllowances sets the bypass_pull_request_allowances property value. Allow specific users, teams, or apps to bypass pull request requirements.
func (m *ProtectedBranchPullRequestReview) SetBypassPullRequestAllowances(value ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable)() {
    m.bypass_pull_request_allowances = value
}
// SetDismissalRestrictions sets the dismissal_restrictions property value. The dismissal_restrictions property
func (m *ProtectedBranchPullRequestReview) SetDismissalRestrictions(value ProtectedBranchPullRequestReview_dismissal_restrictionsable)() {
    m.dismissal_restrictions = value
}
// SetDismissStaleReviews sets the dismiss_stale_reviews property value. The dismiss_stale_reviews property
func (m *ProtectedBranchPullRequestReview) SetDismissStaleReviews(value *bool)() {
    m.dismiss_stale_reviews = value
}
// SetRequireCodeOwnerReviews sets the require_code_owner_reviews property value. The require_code_owner_reviews property
func (m *ProtectedBranchPullRequestReview) SetRequireCodeOwnerReviews(value *bool)() {
    m.require_code_owner_reviews = value
}
// SetRequiredApprovingReviewCount sets the required_approving_review_count property value. The required_approving_review_count property
func (m *ProtectedBranchPullRequestReview) SetRequiredApprovingReviewCount(value *int32)() {
    m.required_approving_review_count = value
}
// SetRequireLastPushApproval sets the require_last_push_approval property value. Whether the most recent push must be approved by someone other than the person who pushed it.
func (m *ProtectedBranchPullRequestReview) SetRequireLastPushApproval(value *bool)() {
    m.require_last_push_approval = value
}
// SetUrl sets the url property value. The url property
func (m *ProtectedBranchPullRequestReview) SetUrl(value *string)() {
    m.url = value
}
type ProtectedBranchPullRequestReviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassPullRequestAllowances()(ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable)
    GetDismissalRestrictions()(ProtectedBranchPullRequestReview_dismissal_restrictionsable)
    GetDismissStaleReviews()(*bool)
    GetRequireCodeOwnerReviews()(*bool)
    GetRequiredApprovingReviewCount()(*int32)
    GetRequireLastPushApproval()(*bool)
    GetUrl()(*string)
    SetBypassPullRequestAllowances(value ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable)()
    SetDismissalRestrictions(value ProtectedBranchPullRequestReview_dismissal_restrictionsable)()
    SetDismissStaleReviews(value *bool)()
    SetRequireCodeOwnerReviews(value *bool)()
    SetRequiredApprovingReviewCount(value *int32)()
    SetRequireLastPushApproval(value *bool)()
    SetUrl(value *string)()
}
