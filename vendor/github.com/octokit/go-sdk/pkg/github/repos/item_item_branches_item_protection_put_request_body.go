package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBranchesItemProtectionPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Allows deletion of the protected branch by anyone with write access to the repository. Set to `false` to prevent deletion of the protected branch. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://docs.github.com/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation.
    allow_deletions *bool
    // Permits force pushes to the protected branch by anyone with write access to the repository. Set to `true` to allow force pushes. Set to `false` or `null` to block force pushes. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://docs.github.com/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation."
    allow_force_pushes *bool
    // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing. Default: `false`.
    allow_fork_syncing *bool
    // If set to `true`, the `restrictions` branch protection settings which limits who can push will also block pushes which create new branches, unless the push is initiated by a user, team, or app which has the ability to push. Set to `true` to restrict new branch creation. Default: `false`.
    block_creations *bool
    // Enforce all configured restrictions for administrators. Set to `true` to enforce required status checks for repository administrators. Set to `null` to disable.
    enforce_admins *bool
    // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch. Default: `false`.
    lock_branch *bool
    // Requires all conversations on code to be resolved before a pull request can be merged into a branch that matches this rule. Set to `false` to disable. Default: `false`.
    required_conversation_resolution *bool
    // Enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch. Set to `true` to enforce a linear commit history. Set to `false` to disable a linear commit Git history. Your repository must allow squash merging or rebase merging before you can enable a linear commit history. Default: `false`. For more information, see "[Requiring a linear commit history](https://docs.github.com/github/administering-a-repository/requiring-a-linear-commit-history)" in the GitHub Help documentation.
    required_linear_history *bool
    // Require at least one approving review on a pull request, before merging. Set to `null` to disable.
    required_pull_request_reviews ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable
    // Require status checks to pass before merging. Set to `null` to disable.
    required_status_checks ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable
    // Restrict who can push to the protected branch. User, app, and team `restrictions` are only available for organization-owned repositories. Set to `null` to disable.
    restrictions ItemItemBranchesItemProtectionPutRequestBody_restrictionsable
}
// NewItemItemBranchesItemProtectionPutRequestBody instantiates a new ItemItemBranchesItemProtectionPutRequestBody and sets the default values.
func NewItemItemBranchesItemProtectionPutRequestBody()(*ItemItemBranchesItemProtectionPutRequestBody) {
    m := &ItemItemBranchesItemProtectionPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDeletions gets the allow_deletions property value. Allows deletion of the protected branch by anyone with write access to the repository. Set to `false` to prevent deletion of the protected branch. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://docs.github.com/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetAllowDeletions()(*bool) {
    return m.allow_deletions
}
// GetAllowForcePushes gets the allow_force_pushes property value. Permits force pushes to the protected branch by anyone with write access to the repository. Set to `true` to allow force pushes. Set to `false` or `null` to block force pushes. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://docs.github.com/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation."
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetAllowForcePushes()(*bool) {
    return m.allow_force_pushes
}
// GetAllowForkSyncing gets the allow_fork_syncing property value. Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing. Default: `false`.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetAllowForkSyncing()(*bool) {
    return m.allow_fork_syncing
}
// GetBlockCreations gets the block_creations property value. If set to `true`, the `restrictions` branch protection settings which limits who can push will also block pushes which create new branches, unless the push is initiated by a user, team, or app which has the ability to push. Set to `true` to restrict new branch creation. Default: `false`.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetBlockCreations()(*bool) {
    return m.block_creations
}
// GetEnforceAdmins gets the enforce_admins property value. Enforce all configured restrictions for administrators. Set to `true` to enforce required status checks for repository administrators. Set to `null` to disable.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetEnforceAdmins()(*bool) {
    return m.enforce_admins
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_deletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeletions(val)
        }
        return nil
    }
    res["allow_force_pushes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForcePushes(val)
        }
        return nil
    }
    res["allow_fork_syncing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForkSyncing(val)
        }
        return nil
    }
    res["block_creations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCreations(val)
        }
        return nil
    }
    res["enforce_admins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceAdmins(val)
        }
        return nil
    }
    res["lock_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockBranch(val)
        }
        return nil
    }
    res["required_conversation_resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredConversationResolution(val)
        }
        return nil
    }
    res["required_linear_history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredLinearHistory(val)
        }
        return nil
    }
    res["required_pull_request_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPullRequestReviews(val.(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable))
        }
        return nil
    }
    res["required_status_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionPutRequestBody_required_status_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredStatusChecks(val.(ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable))
        }
        return nil
    }
    res["restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemBranchesItemProtectionPutRequestBody_restrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictions(val.(ItemItemBranchesItemProtectionPutRequestBody_restrictionsable))
        }
        return nil
    }
    return res
}
// GetLockBranch gets the lock_branch property value. Whether to set the branch as read-only. If this is true, users will not be able to push to the branch. Default: `false`.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetLockBranch()(*bool) {
    return m.lock_branch
}
// GetRequiredConversationResolution gets the required_conversation_resolution property value. Requires all conversations on code to be resolved before a pull request can be merged into a branch that matches this rule. Set to `false` to disable. Default: `false`.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetRequiredConversationResolution()(*bool) {
    return m.required_conversation_resolution
}
// GetRequiredLinearHistory gets the required_linear_history property value. Enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch. Set to `true` to enforce a linear commit history. Set to `false` to disable a linear commit Git history. Your repository must allow squash merging or rebase merging before you can enable a linear commit history. Default: `false`. For more information, see "[Requiring a linear commit history](https://docs.github.com/github/administering-a-repository/requiring-a-linear-commit-history)" in the GitHub Help documentation.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetRequiredLinearHistory()(*bool) {
    return m.required_linear_history
}
// GetRequiredPullRequestReviews gets the required_pull_request_reviews property value. Require at least one approving review on a pull request, before merging. Set to `null` to disable.
// returns a ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetRequiredPullRequestReviews()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable) {
    return m.required_pull_request_reviews
}
// GetRequiredStatusChecks gets the required_status_checks property value. Require status checks to pass before merging. Set to `null` to disable.
// returns a ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetRequiredStatusChecks()(ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable) {
    return m.required_status_checks
}
// GetRestrictions gets the restrictions property value. Restrict who can push to the protected branch. User, app, and team `restrictions` are only available for organization-owned repositories. Set to `null` to disable.
// returns a ItemItemBranchesItemProtectionPutRequestBody_restrictionsable when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody) GetRestrictions()(ItemItemBranchesItemProtectionPutRequestBody_restrictionsable) {
    return m.restrictions
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_deletions", m.GetAllowDeletions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_force_pushes", m.GetAllowForcePushes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_fork_syncing", m.GetAllowForkSyncing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("block_creations", m.GetBlockCreations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enforce_admins", m.GetEnforceAdmins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("lock_branch", m.GetLockBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required_conversation_resolution", m.GetRequiredConversationResolution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required_linear_history", m.GetRequiredLinearHistory())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDeletions sets the allow_deletions property value. Allows deletion of the protected branch by anyone with write access to the repository. Set to `false` to prevent deletion of the protected branch. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://docs.github.com/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetAllowDeletions(value *bool)() {
    m.allow_deletions = value
}
// SetAllowForcePushes sets the allow_force_pushes property value. Permits force pushes to the protected branch by anyone with write access to the repository. Set to `true` to allow force pushes. Set to `false` or `null` to block force pushes. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://docs.github.com/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation."
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetAllowForcePushes(value *bool)() {
    m.allow_force_pushes = value
}
// SetAllowForkSyncing sets the allow_fork_syncing property value. Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing. Default: `false`.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetAllowForkSyncing(value *bool)() {
    m.allow_fork_syncing = value
}
// SetBlockCreations sets the block_creations property value. If set to `true`, the `restrictions` branch protection settings which limits who can push will also block pushes which create new branches, unless the push is initiated by a user, team, or app which has the ability to push. Set to `true` to restrict new branch creation. Default: `false`.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetBlockCreations(value *bool)() {
    m.block_creations = value
}
// SetEnforceAdmins sets the enforce_admins property value. Enforce all configured restrictions for administrators. Set to `true` to enforce required status checks for repository administrators. Set to `null` to disable.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetEnforceAdmins(value *bool)() {
    m.enforce_admins = value
}
// SetLockBranch sets the lock_branch property value. Whether to set the branch as read-only. If this is true, users will not be able to push to the branch. Default: `false`.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetLockBranch(value *bool)() {
    m.lock_branch = value
}
// SetRequiredConversationResolution sets the required_conversation_resolution property value. Requires all conversations on code to be resolved before a pull request can be merged into a branch that matches this rule. Set to `false` to disable. Default: `false`.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetRequiredConversationResolution(value *bool)() {
    m.required_conversation_resolution = value
}
// SetRequiredLinearHistory sets the required_linear_history property value. Enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch. Set to `true` to enforce a linear commit history. Set to `false` to disable a linear commit Git history. Your repository must allow squash merging or rebase merging before you can enable a linear commit history. Default: `false`. For more information, see "[Requiring a linear commit history](https://docs.github.com/github/administering-a-repository/requiring-a-linear-commit-history)" in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetRequiredLinearHistory(value *bool)() {
    m.required_linear_history = value
}
// SetRequiredPullRequestReviews sets the required_pull_request_reviews property value. Require at least one approving review on a pull request, before merging. Set to `null` to disable.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetRequiredPullRequestReviews(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable)() {
    m.required_pull_request_reviews = value
}
// SetRequiredStatusChecks sets the required_status_checks property value. Require status checks to pass before merging. Set to `null` to disable.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetRequiredStatusChecks(value ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable)() {
    m.required_status_checks = value
}
// SetRestrictions sets the restrictions property value. Restrict who can push to the protected branch. User, app, and team `restrictions` are only available for organization-owned repositories. Set to `null` to disable.
func (m *ItemItemBranchesItemProtectionPutRequestBody) SetRestrictions(value ItemItemBranchesItemProtectionPutRequestBody_restrictionsable)() {
    m.restrictions = value
}
type ItemItemBranchesItemProtectionPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeletions()(*bool)
    GetAllowForcePushes()(*bool)
    GetAllowForkSyncing()(*bool)
    GetBlockCreations()(*bool)
    GetEnforceAdmins()(*bool)
    GetLockBranch()(*bool)
    GetRequiredConversationResolution()(*bool)
    GetRequiredLinearHistory()(*bool)
    GetRequiredPullRequestReviews()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable)
    GetRequiredStatusChecks()(ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable)
    GetRestrictions()(ItemItemBranchesItemProtectionPutRequestBody_restrictionsable)
    SetAllowDeletions(value *bool)()
    SetAllowForcePushes(value *bool)()
    SetAllowForkSyncing(value *bool)()
    SetBlockCreations(value *bool)()
    SetEnforceAdmins(value *bool)()
    SetLockBranch(value *bool)()
    SetRequiredConversationResolution(value *bool)()
    SetRequiredLinearHistory(value *bool)()
    SetRequiredPullRequestReviews(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable)()
    SetRequiredStatusChecks(value ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable)()
    SetRestrictions(value ItemItemBranchesItemProtectionPutRequestBody_restrictionsable)()
}
