package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchProtection branch Protection
type BranchProtection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_deletions property
    allow_deletions BranchProtection_allow_deletionsable
    // The allow_force_pushes property
    allow_force_pushes BranchProtection_allow_force_pushesable
    // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
    allow_fork_syncing BranchProtection_allow_fork_syncingable
    // The block_creations property
    block_creations BranchProtection_block_creationsable
    // The enabled property
    enabled *bool
    // Protected Branch Admin Enforced
    enforce_admins ProtectedBranchAdminEnforcedable
    // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
    lock_branch BranchProtection_lock_branchable
    // The name property
    name *string
    // The protection_url property
    protection_url *string
    // The required_conversation_resolution property
    required_conversation_resolution BranchProtection_required_conversation_resolutionable
    // The required_linear_history property
    required_linear_history BranchProtection_required_linear_historyable
    // Protected Branch Pull Request Review
    required_pull_request_reviews ProtectedBranchPullRequestReviewable
    // The required_signatures property
    required_signatures BranchProtection_required_signaturesable
    // Protected Branch Required Status Check
    required_status_checks ProtectedBranchRequiredStatusCheckable
    // Branch Restriction Policy
    restrictions BranchRestrictionPolicyable
    // The url property
    url *string
}
// NewBranchProtection instantiates a new BranchProtection and sets the default values.
func NewBranchProtection()(*BranchProtection) {
    m := &BranchProtection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchProtection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchProtection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDeletions gets the allow_deletions property value. The allow_deletions property
// returns a BranchProtection_allow_deletionsable when successful
func (m *BranchProtection) GetAllowDeletions()(BranchProtection_allow_deletionsable) {
    return m.allow_deletions
}
// GetAllowForcePushes gets the allow_force_pushes property value. The allow_force_pushes property
// returns a BranchProtection_allow_force_pushesable when successful
func (m *BranchProtection) GetAllowForcePushes()(BranchProtection_allow_force_pushesable) {
    return m.allow_force_pushes
}
// GetAllowForkSyncing gets the allow_fork_syncing property value. Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
// returns a BranchProtection_allow_fork_syncingable when successful
func (m *BranchProtection) GetAllowForkSyncing()(BranchProtection_allow_fork_syncingable) {
    return m.allow_fork_syncing
}
// GetBlockCreations gets the block_creations property value. The block_creations property
// returns a BranchProtection_block_creationsable when successful
func (m *BranchProtection) GetBlockCreations()(BranchProtection_block_creationsable) {
    return m.block_creations
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *BranchProtection) GetEnabled()(*bool) {
    return m.enabled
}
// GetEnforceAdmins gets the enforce_admins property value. Protected Branch Admin Enforced
// returns a ProtectedBranchAdminEnforcedable when successful
func (m *BranchProtection) GetEnforceAdmins()(ProtectedBranchAdminEnforcedable) {
    return m.enforce_admins
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_deletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_allow_deletionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeletions(val.(BranchProtection_allow_deletionsable))
        }
        return nil
    }
    res["allow_force_pushes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_allow_force_pushesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForcePushes(val.(BranchProtection_allow_force_pushesable))
        }
        return nil
    }
    res["allow_fork_syncing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_allow_fork_syncingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForkSyncing(val.(BranchProtection_allow_fork_syncingable))
        }
        return nil
    }
    res["block_creations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_block_creationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCreations(val.(BranchProtection_block_creationsable))
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["enforce_admins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranchAdminEnforcedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceAdmins(val.(ProtectedBranchAdminEnforcedable))
        }
        return nil
    }
    res["lock_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_lock_branchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockBranch(val.(BranchProtection_lock_branchable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["protection_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionUrl(val)
        }
        return nil
    }
    res["required_conversation_resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_required_conversation_resolutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredConversationResolution(val.(BranchProtection_required_conversation_resolutionable))
        }
        return nil
    }
    res["required_linear_history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_required_linear_historyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredLinearHistory(val.(BranchProtection_required_linear_historyable))
        }
        return nil
    }
    res["required_pull_request_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranchPullRequestReviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPullRequestReviews(val.(ProtectedBranchPullRequestReviewable))
        }
        return nil
    }
    res["required_signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtection_required_signaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredSignatures(val.(BranchProtection_required_signaturesable))
        }
        return nil
    }
    res["required_status_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranchRequiredStatusCheckFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredStatusChecks(val.(ProtectedBranchRequiredStatusCheckable))
        }
        return nil
    }
    res["restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchRestrictionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictions(val.(BranchRestrictionPolicyable))
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
// GetLockBranch gets the lock_branch property value. Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
// returns a BranchProtection_lock_branchable when successful
func (m *BranchProtection) GetLockBranch()(BranchProtection_lock_branchable) {
    return m.lock_branch
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BranchProtection) GetName()(*string) {
    return m.name
}
// GetProtectionUrl gets the protection_url property value. The protection_url property
// returns a *string when successful
func (m *BranchProtection) GetProtectionUrl()(*string) {
    return m.protection_url
}
// GetRequiredConversationResolution gets the required_conversation_resolution property value. The required_conversation_resolution property
// returns a BranchProtection_required_conversation_resolutionable when successful
func (m *BranchProtection) GetRequiredConversationResolution()(BranchProtection_required_conversation_resolutionable) {
    return m.required_conversation_resolution
}
// GetRequiredLinearHistory gets the required_linear_history property value. The required_linear_history property
// returns a BranchProtection_required_linear_historyable when successful
func (m *BranchProtection) GetRequiredLinearHistory()(BranchProtection_required_linear_historyable) {
    return m.required_linear_history
}
// GetRequiredPullRequestReviews gets the required_pull_request_reviews property value. Protected Branch Pull Request Review
// returns a ProtectedBranchPullRequestReviewable when successful
func (m *BranchProtection) GetRequiredPullRequestReviews()(ProtectedBranchPullRequestReviewable) {
    return m.required_pull_request_reviews
}
// GetRequiredSignatures gets the required_signatures property value. The required_signatures property
// returns a BranchProtection_required_signaturesable when successful
func (m *BranchProtection) GetRequiredSignatures()(BranchProtection_required_signaturesable) {
    return m.required_signatures
}
// GetRequiredStatusChecks gets the required_status_checks property value. Protected Branch Required Status Check
// returns a ProtectedBranchRequiredStatusCheckable when successful
func (m *BranchProtection) GetRequiredStatusChecks()(ProtectedBranchRequiredStatusCheckable) {
    return m.required_status_checks
}
// GetRestrictions gets the restrictions property value. Branch Restriction Policy
// returns a BranchRestrictionPolicyable when successful
func (m *BranchProtection) GetRestrictions()(BranchRestrictionPolicyable) {
    return m.restrictions
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *BranchProtection) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *BranchProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("allow_deletions", m.GetAllowDeletions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("allow_force_pushes", m.GetAllowForcePushes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("allow_fork_syncing", m.GetAllowForkSyncing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("block_creations", m.GetBlockCreations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("enforce_admins", m.GetEnforceAdmins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lock_branch", m.GetLockBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("protection_url", m.GetProtectionUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("required_conversation_resolution", m.GetRequiredConversationResolution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("required_linear_history", m.GetRequiredLinearHistory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("required_pull_request_reviews", m.GetRequiredPullRequestReviews())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("required_signatures", m.GetRequiredSignatures())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("required_status_checks", m.GetRequiredStatusChecks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("restrictions", m.GetRestrictions())
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
func (m *BranchProtection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDeletions sets the allow_deletions property value. The allow_deletions property
func (m *BranchProtection) SetAllowDeletions(value BranchProtection_allow_deletionsable)() {
    m.allow_deletions = value
}
// SetAllowForcePushes sets the allow_force_pushes property value. The allow_force_pushes property
func (m *BranchProtection) SetAllowForcePushes(value BranchProtection_allow_force_pushesable)() {
    m.allow_force_pushes = value
}
// SetAllowForkSyncing sets the allow_fork_syncing property value. Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
func (m *BranchProtection) SetAllowForkSyncing(value BranchProtection_allow_fork_syncingable)() {
    m.allow_fork_syncing = value
}
// SetBlockCreations sets the block_creations property value. The block_creations property
func (m *BranchProtection) SetBlockCreations(value BranchProtection_block_creationsable)() {
    m.block_creations = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *BranchProtection) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEnforceAdmins sets the enforce_admins property value. Protected Branch Admin Enforced
func (m *BranchProtection) SetEnforceAdmins(value ProtectedBranchAdminEnforcedable)() {
    m.enforce_admins = value
}
// SetLockBranch sets the lock_branch property value. Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
func (m *BranchProtection) SetLockBranch(value BranchProtection_lock_branchable)() {
    m.lock_branch = value
}
// SetName sets the name property value. The name property
func (m *BranchProtection) SetName(value *string)() {
    m.name = value
}
// SetProtectionUrl sets the protection_url property value. The protection_url property
func (m *BranchProtection) SetProtectionUrl(value *string)() {
    m.protection_url = value
}
// SetRequiredConversationResolution sets the required_conversation_resolution property value. The required_conversation_resolution property
func (m *BranchProtection) SetRequiredConversationResolution(value BranchProtection_required_conversation_resolutionable)() {
    m.required_conversation_resolution = value
}
// SetRequiredLinearHistory sets the required_linear_history property value. The required_linear_history property
func (m *BranchProtection) SetRequiredLinearHistory(value BranchProtection_required_linear_historyable)() {
    m.required_linear_history = value
}
// SetRequiredPullRequestReviews sets the required_pull_request_reviews property value. Protected Branch Pull Request Review
func (m *BranchProtection) SetRequiredPullRequestReviews(value ProtectedBranchPullRequestReviewable)() {
    m.required_pull_request_reviews = value
}
// SetRequiredSignatures sets the required_signatures property value. The required_signatures property
func (m *BranchProtection) SetRequiredSignatures(value BranchProtection_required_signaturesable)() {
    m.required_signatures = value
}
// SetRequiredStatusChecks sets the required_status_checks property value. Protected Branch Required Status Check
func (m *BranchProtection) SetRequiredStatusChecks(value ProtectedBranchRequiredStatusCheckable)() {
    m.required_status_checks = value
}
// SetRestrictions sets the restrictions property value. Branch Restriction Policy
func (m *BranchProtection) SetRestrictions(value BranchRestrictionPolicyable)() {
    m.restrictions = value
}
// SetUrl sets the url property value. The url property
func (m *BranchProtection) SetUrl(value *string)() {
    m.url = value
}
type BranchProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeletions()(BranchProtection_allow_deletionsable)
    GetAllowForcePushes()(BranchProtection_allow_force_pushesable)
    GetAllowForkSyncing()(BranchProtection_allow_fork_syncingable)
    GetBlockCreations()(BranchProtection_block_creationsable)
    GetEnabled()(*bool)
    GetEnforceAdmins()(ProtectedBranchAdminEnforcedable)
    GetLockBranch()(BranchProtection_lock_branchable)
    GetName()(*string)
    GetProtectionUrl()(*string)
    GetRequiredConversationResolution()(BranchProtection_required_conversation_resolutionable)
    GetRequiredLinearHistory()(BranchProtection_required_linear_historyable)
    GetRequiredPullRequestReviews()(ProtectedBranchPullRequestReviewable)
    GetRequiredSignatures()(BranchProtection_required_signaturesable)
    GetRequiredStatusChecks()(ProtectedBranchRequiredStatusCheckable)
    GetRestrictions()(BranchRestrictionPolicyable)
    GetUrl()(*string)
    SetAllowDeletions(value BranchProtection_allow_deletionsable)()
    SetAllowForcePushes(value BranchProtection_allow_force_pushesable)()
    SetAllowForkSyncing(value BranchProtection_allow_fork_syncingable)()
    SetBlockCreations(value BranchProtection_block_creationsable)()
    SetEnabled(value *bool)()
    SetEnforceAdmins(value ProtectedBranchAdminEnforcedable)()
    SetLockBranch(value BranchProtection_lock_branchable)()
    SetName(value *string)()
    SetProtectionUrl(value *string)()
    SetRequiredConversationResolution(value BranchProtection_required_conversation_resolutionable)()
    SetRequiredLinearHistory(value BranchProtection_required_linear_historyable)()
    SetRequiredPullRequestReviews(value ProtectedBranchPullRequestReviewable)()
    SetRequiredSignatures(value BranchProtection_required_signaturesable)()
    SetRequiredStatusChecks(value ProtectedBranchRequiredStatusCheckable)()
    SetRestrictions(value BranchRestrictionPolicyable)()
    SetUrl(value *string)()
}
