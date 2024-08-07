package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Allow specific users, teams, or apps to bypass pull request requirements.
    bypass_pull_request_allowances ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable
    // Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
    dismiss_stale_reviews *bool
    // Specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
    dismissal_restrictions ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable
    // Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) have reviewed.
    require_code_owner_reviews *bool
    // Whether the most recent push must be approved by someone other than the person who pushed it. Default: `false`
    require_last_push_approval *bool
    // Specifies the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
    required_approving_review_count *int32
}
// NewItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody instantiates a new ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody and sets the default values.
func NewItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody()(*ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) {
    m := &ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassPullRequestAllowances gets the bypass_pull_request_allowances property value. Allow specific users, teams, or apps to bypass pull request requirements.
// returns a ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetBypassPullRequestAllowances()(ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable) {
    return m.bypass_pull_request_allowances
}
// GetDismissalRestrictions gets the dismissal_restrictions property value. Specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
// returns a ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetDismissalRestrictions()(ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable) {
    return m.dismissal_restrictions
}
// GetDismissStaleReviews gets the dismiss_stale_reviews property value. Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetDismissStaleReviews()(*bool) {
    return m.dismiss_stale_reviews
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_pull_request_allowances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassPullRequestAllowances(val.(ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable))
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
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissalRestrictions(val.(ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable))
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
    return res
}
// GetRequireCodeOwnerReviews gets the require_code_owner_reviews property value. Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) have reviewed.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetRequireCodeOwnerReviews()(*bool) {
    return m.require_code_owner_reviews
}
// GetRequiredApprovingReviewCount gets the required_approving_review_count property value. Specifies the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
// returns a *int32 when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetRequiredApprovingReviewCount()(*int32) {
    return m.required_approving_review_count
}
// GetRequireLastPushApproval gets the require_last_push_approval property value. Whether the most recent push must be approved by someone other than the person who pushed it. Default: `false`
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) GetRequireLastPushApproval()(*bool) {
    return m.require_last_push_approval
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassPullRequestAllowances sets the bypass_pull_request_allowances property value. Allow specific users, teams, or apps to bypass pull request requirements.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetBypassPullRequestAllowances(value ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable)() {
    m.bypass_pull_request_allowances = value
}
// SetDismissalRestrictions sets the dismissal_restrictions property value. Specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetDismissalRestrictions(value ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable)() {
    m.dismissal_restrictions = value
}
// SetDismissStaleReviews sets the dismiss_stale_reviews property value. Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetDismissStaleReviews(value *bool)() {
    m.dismiss_stale_reviews = value
}
// SetRequireCodeOwnerReviews sets the require_code_owner_reviews property value. Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) have reviewed.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetRequireCodeOwnerReviews(value *bool)() {
    m.require_code_owner_reviews = value
}
// SetRequiredApprovingReviewCount sets the required_approving_review_count property value. Specifies the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetRequiredApprovingReviewCount(value *int32)() {
    m.required_approving_review_count = value
}
// SetRequireLastPushApproval sets the require_last_push_approval property value. Whether the most recent push must be approved by someone other than the person who pushed it. Default: `false`
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody) SetRequireLastPushApproval(value *bool)() {
    m.require_last_push_approval = value
}
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassPullRequestAllowances()(ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable)
    GetDismissalRestrictions()(ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable)
    GetDismissStaleReviews()(*bool)
    GetRequireCodeOwnerReviews()(*bool)
    GetRequiredApprovingReviewCount()(*int32)
    GetRequireLastPushApproval()(*bool)
    SetBypassPullRequestAllowances(value ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_bypass_pull_request_allowancesable)()
    SetDismissalRestrictions(value ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable)()
    SetDismissStaleReviews(value *bool)()
    SetRequireCodeOwnerReviews(value *bool)()
    SetRequiredApprovingReviewCount(value *int32)()
    SetRequireLastPushApproval(value *bool)()
}
