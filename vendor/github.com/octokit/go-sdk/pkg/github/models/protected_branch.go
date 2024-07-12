package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranch branch protections protect branches
type ProtectedBranch struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_deletions property
    allow_deletions ProtectedBranch_allow_deletionsable
    // The allow_force_pushes property
    allow_force_pushes ProtectedBranch_allow_force_pushesable
    // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
    allow_fork_syncing ProtectedBranch_allow_fork_syncingable
    // The block_creations property
    block_creations ProtectedBranch_block_creationsable
    // The enforce_admins property
    enforce_admins ProtectedBranch_enforce_adminsable
    // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
    lock_branch ProtectedBranch_lock_branchable
    // The required_conversation_resolution property
    required_conversation_resolution ProtectedBranch_required_conversation_resolutionable
    // The required_linear_history property
    required_linear_history ProtectedBranch_required_linear_historyable
    // The required_pull_request_reviews property
    required_pull_request_reviews ProtectedBranch_required_pull_request_reviewsable
    // The required_signatures property
    required_signatures ProtectedBranch_required_signaturesable
    // Status Check Policy
    required_status_checks StatusCheckPolicyable
    // Branch Restriction Policy
    restrictions BranchRestrictionPolicyable
    // The url property
    url *string
}
// NewProtectedBranch instantiates a new ProtectedBranch and sets the default values.
func NewProtectedBranch()(*ProtectedBranch) {
    m := &ProtectedBranch{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedBranchFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedBranchFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranch(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtectedBranch) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDeletions gets the allow_deletions property value. The allow_deletions property
// returns a ProtectedBranch_allow_deletionsable when successful
func (m *ProtectedBranch) GetAllowDeletions()(ProtectedBranch_allow_deletionsable) {
    return m.allow_deletions
}
// GetAllowForcePushes gets the allow_force_pushes property value. The allow_force_pushes property
// returns a ProtectedBranch_allow_force_pushesable when successful
func (m *ProtectedBranch) GetAllowForcePushes()(ProtectedBranch_allow_force_pushesable) {
    return m.allow_force_pushes
}
// GetAllowForkSyncing gets the allow_fork_syncing property value. Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
// returns a ProtectedBranch_allow_fork_syncingable when successful
func (m *ProtectedBranch) GetAllowForkSyncing()(ProtectedBranch_allow_fork_syncingable) {
    return m.allow_fork_syncing
}
// GetBlockCreations gets the block_creations property value. The block_creations property
// returns a ProtectedBranch_block_creationsable when successful
func (m *ProtectedBranch) GetBlockCreations()(ProtectedBranch_block_creationsable) {
    return m.block_creations
}
// GetEnforceAdmins gets the enforce_admins property value. The enforce_admins property
// returns a ProtectedBranch_enforce_adminsable when successful
func (m *ProtectedBranch) GetEnforceAdmins()(ProtectedBranch_enforce_adminsable) {
    return m.enforce_admins
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedBranch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_deletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_allow_deletionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeletions(val.(ProtectedBranch_allow_deletionsable))
        }
        return nil
    }
    res["allow_force_pushes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_allow_force_pushesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForcePushes(val.(ProtectedBranch_allow_force_pushesable))
        }
        return nil
    }
    res["allow_fork_syncing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_allow_fork_syncingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForkSyncing(val.(ProtectedBranch_allow_fork_syncingable))
        }
        return nil
    }
    res["block_creations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_block_creationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCreations(val.(ProtectedBranch_block_creationsable))
        }
        return nil
    }
    res["enforce_admins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_enforce_adminsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceAdmins(val.(ProtectedBranch_enforce_adminsable))
        }
        return nil
    }
    res["lock_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_lock_branchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockBranch(val.(ProtectedBranch_lock_branchable))
        }
        return nil
    }
    res["required_conversation_resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_required_conversation_resolutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredConversationResolution(val.(ProtectedBranch_required_conversation_resolutionable))
        }
        return nil
    }
    res["required_linear_history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_required_linear_historyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredLinearHistory(val.(ProtectedBranch_required_linear_historyable))
        }
        return nil
    }
    res["required_pull_request_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_required_pull_request_reviewsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPullRequestReviews(val.(ProtectedBranch_required_pull_request_reviewsable))
        }
        return nil
    }
    res["required_signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProtectedBranch_required_signaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredSignatures(val.(ProtectedBranch_required_signaturesable))
        }
        return nil
    }
    res["required_status_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatusCheckPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredStatusChecks(val.(StatusCheckPolicyable))
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
// returns a ProtectedBranch_lock_branchable when successful
func (m *ProtectedBranch) GetLockBranch()(ProtectedBranch_lock_branchable) {
    return m.lock_branch
}
// GetRequiredConversationResolution gets the required_conversation_resolution property value. The required_conversation_resolution property
// returns a ProtectedBranch_required_conversation_resolutionable when successful
func (m *ProtectedBranch) GetRequiredConversationResolution()(ProtectedBranch_required_conversation_resolutionable) {
    return m.required_conversation_resolution
}
// GetRequiredLinearHistory gets the required_linear_history property value. The required_linear_history property
// returns a ProtectedBranch_required_linear_historyable when successful
func (m *ProtectedBranch) GetRequiredLinearHistory()(ProtectedBranch_required_linear_historyable) {
    return m.required_linear_history
}
// GetRequiredPullRequestReviews gets the required_pull_request_reviews property value. The required_pull_request_reviews property
// returns a ProtectedBranch_required_pull_request_reviewsable when successful
func (m *ProtectedBranch) GetRequiredPullRequestReviews()(ProtectedBranch_required_pull_request_reviewsable) {
    return m.required_pull_request_reviews
}
// GetRequiredSignatures gets the required_signatures property value. The required_signatures property
// returns a ProtectedBranch_required_signaturesable when successful
func (m *ProtectedBranch) GetRequiredSignatures()(ProtectedBranch_required_signaturesable) {
    return m.required_signatures
}
// GetRequiredStatusChecks gets the required_status_checks property value. Status Check Policy
// returns a StatusCheckPolicyable when successful
func (m *ProtectedBranch) GetRequiredStatusChecks()(StatusCheckPolicyable) {
    return m.required_status_checks
}
// GetRestrictions gets the restrictions property value. Branch Restriction Policy
// returns a BranchRestrictionPolicyable when successful
func (m *ProtectedBranch) GetRestrictions()(BranchRestrictionPolicyable) {
    return m.restrictions
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ProtectedBranch) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ProtectedBranch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ProtectedBranch) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDeletions sets the allow_deletions property value. The allow_deletions property
func (m *ProtectedBranch) SetAllowDeletions(value ProtectedBranch_allow_deletionsable)() {
    m.allow_deletions = value
}
// SetAllowForcePushes sets the allow_force_pushes property value. The allow_force_pushes property
func (m *ProtectedBranch) SetAllowForcePushes(value ProtectedBranch_allow_force_pushesable)() {
    m.allow_force_pushes = value
}
// SetAllowForkSyncing sets the allow_fork_syncing property value. Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
func (m *ProtectedBranch) SetAllowForkSyncing(value ProtectedBranch_allow_fork_syncingable)() {
    m.allow_fork_syncing = value
}
// SetBlockCreations sets the block_creations property value. The block_creations property
func (m *ProtectedBranch) SetBlockCreations(value ProtectedBranch_block_creationsable)() {
    m.block_creations = value
}
// SetEnforceAdmins sets the enforce_admins property value. The enforce_admins property
func (m *ProtectedBranch) SetEnforceAdmins(value ProtectedBranch_enforce_adminsable)() {
    m.enforce_admins = value
}
// SetLockBranch sets the lock_branch property value. Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
func (m *ProtectedBranch) SetLockBranch(value ProtectedBranch_lock_branchable)() {
    m.lock_branch = value
}
// SetRequiredConversationResolution sets the required_conversation_resolution property value. The required_conversation_resolution property
func (m *ProtectedBranch) SetRequiredConversationResolution(value ProtectedBranch_required_conversation_resolutionable)() {
    m.required_conversation_resolution = value
}
// SetRequiredLinearHistory sets the required_linear_history property value. The required_linear_history property
func (m *ProtectedBranch) SetRequiredLinearHistory(value ProtectedBranch_required_linear_historyable)() {
    m.required_linear_history = value
}
// SetRequiredPullRequestReviews sets the required_pull_request_reviews property value. The required_pull_request_reviews property
func (m *ProtectedBranch) SetRequiredPullRequestReviews(value ProtectedBranch_required_pull_request_reviewsable)() {
    m.required_pull_request_reviews = value
}
// SetRequiredSignatures sets the required_signatures property value. The required_signatures property
func (m *ProtectedBranch) SetRequiredSignatures(value ProtectedBranch_required_signaturesable)() {
    m.required_signatures = value
}
// SetRequiredStatusChecks sets the required_status_checks property value. Status Check Policy
func (m *ProtectedBranch) SetRequiredStatusChecks(value StatusCheckPolicyable)() {
    m.required_status_checks = value
}
// SetRestrictions sets the restrictions property value. Branch Restriction Policy
func (m *ProtectedBranch) SetRestrictions(value BranchRestrictionPolicyable)() {
    m.restrictions = value
}
// SetUrl sets the url property value. The url property
func (m *ProtectedBranch) SetUrl(value *string)() {
    m.url = value
}
type ProtectedBranchable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeletions()(ProtectedBranch_allow_deletionsable)
    GetAllowForcePushes()(ProtectedBranch_allow_force_pushesable)
    GetAllowForkSyncing()(ProtectedBranch_allow_fork_syncingable)
    GetBlockCreations()(ProtectedBranch_block_creationsable)
    GetEnforceAdmins()(ProtectedBranch_enforce_adminsable)
    GetLockBranch()(ProtectedBranch_lock_branchable)
    GetRequiredConversationResolution()(ProtectedBranch_required_conversation_resolutionable)
    GetRequiredLinearHistory()(ProtectedBranch_required_linear_historyable)
    GetRequiredPullRequestReviews()(ProtectedBranch_required_pull_request_reviewsable)
    GetRequiredSignatures()(ProtectedBranch_required_signaturesable)
    GetRequiredStatusChecks()(StatusCheckPolicyable)
    GetRestrictions()(BranchRestrictionPolicyable)
    GetUrl()(*string)
    SetAllowDeletions(value ProtectedBranch_allow_deletionsable)()
    SetAllowForcePushes(value ProtectedBranch_allow_force_pushesable)()
    SetAllowForkSyncing(value ProtectedBranch_allow_fork_syncingable)()
    SetBlockCreations(value ProtectedBranch_block_creationsable)()
    SetEnforceAdmins(value ProtectedBranch_enforce_adminsable)()
    SetLockBranch(value ProtectedBranch_lock_branchable)()
    SetRequiredConversationResolution(value ProtectedBranch_required_conversation_resolutionable)()
    SetRequiredLinearHistory(value ProtectedBranch_required_linear_historyable)()
    SetRequiredPullRequestReviews(value ProtectedBranch_required_pull_request_reviewsable)()
    SetRequiredSignatures(value ProtectedBranch_required_signaturesable)()
    SetRequiredStatusChecks(value StatusCheckPolicyable)()
    SetRestrictions(value BranchRestrictionPolicyable)()
    SetUrl(value *string)()
}
