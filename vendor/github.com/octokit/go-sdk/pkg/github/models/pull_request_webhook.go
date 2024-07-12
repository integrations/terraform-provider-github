package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullRequestWebhook struct {
    PullRequest
    // Whether to allow auto-merge for pull requests.
    allow_auto_merge *bool
    // Whether to allow updating the pull request's branch.
    allow_update_branch *bool
    // Whether to delete head branches when pull requests are merged.
    delete_branch_on_merge *bool
    // The default value for a merge commit message.- `PR_TITLE` - default to the pull request's title.- `PR_BODY` - default to the pull request's body.- `BLANK` - default to a blank commit message.
    merge_commit_message *PullRequestWebhook_merge_commit_message
    // The default value for a merge commit title.- `PR_TITLE` - default to the pull request's title.- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
    merge_commit_title *PullRequestWebhook_merge_commit_title
    // The default value for a squash merge commit message:- `PR_BODY` - default to the pull request's body.- `COMMIT_MESSAGES` - default to the branch's commit messages.- `BLANK` - default to a blank commit message.
    squash_merge_commit_message *PullRequestWebhook_squash_merge_commit_message
    // The default value for a squash merge commit title:- `PR_TITLE` - default to the pull request's title.- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
    squash_merge_commit_title *PullRequestWebhook_squash_merge_commit_title
    // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.**
    use_squash_pr_title_as_default *bool
}
// NewPullRequestWebhook instantiates a new PullRequestWebhook and sets the default values.
func NewPullRequestWebhook()(*PullRequestWebhook) {
    m := &PullRequestWebhook{
        PullRequest: *NewPullRequest(),
    }
    return m
}
// CreatePullRequestWebhookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestWebhookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestWebhook(), nil
}
// GetAllowAutoMerge gets the allow_auto_merge property value. Whether to allow auto-merge for pull requests.
// returns a *bool when successful
func (m *PullRequestWebhook) GetAllowAutoMerge()(*bool) {
    return m.allow_auto_merge
}
// GetAllowUpdateBranch gets the allow_update_branch property value. Whether to allow updating the pull request's branch.
// returns a *bool when successful
func (m *PullRequestWebhook) GetAllowUpdateBranch()(*bool) {
    return m.allow_update_branch
}
// GetDeleteBranchOnMerge gets the delete_branch_on_merge property value. Whether to delete head branches when pull requests are merged.
// returns a *bool when successful
func (m *PullRequestWebhook) GetDeleteBranchOnMerge()(*bool) {
    return m.delete_branch_on_merge
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestWebhook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PullRequest.GetFieldDeserializers()
    res["allow_auto_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAutoMerge(val)
        }
        return nil
    }
    res["allow_update_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUpdateBranch(val)
        }
        return nil
    }
    res["delete_branch_on_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteBranchOnMerge(val)
        }
        return nil
    }
    res["merge_commit_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestWebhook_merge_commit_message)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeCommitMessage(val.(*PullRequestWebhook_merge_commit_message))
        }
        return nil
    }
    res["merge_commit_title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestWebhook_merge_commit_title)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeCommitTitle(val.(*PullRequestWebhook_merge_commit_title))
        }
        return nil
    }
    res["squash_merge_commit_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestWebhook_squash_merge_commit_message)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquashMergeCommitMessage(val.(*PullRequestWebhook_squash_merge_commit_message))
        }
        return nil
    }
    res["squash_merge_commit_title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestWebhook_squash_merge_commit_title)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquashMergeCommitTitle(val.(*PullRequestWebhook_squash_merge_commit_title))
        }
        return nil
    }
    res["use_squash_pr_title_as_default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSquashPrTitleAsDefault(val)
        }
        return nil
    }
    return res
}
// GetMergeCommitMessage gets the merge_commit_message property value. The default value for a merge commit message.- `PR_TITLE` - default to the pull request's title.- `PR_BODY` - default to the pull request's body.- `BLANK` - default to a blank commit message.
// returns a *PullRequestWebhook_merge_commit_message when successful
func (m *PullRequestWebhook) GetMergeCommitMessage()(*PullRequestWebhook_merge_commit_message) {
    return m.merge_commit_message
}
// GetMergeCommitTitle gets the merge_commit_title property value. The default value for a merge commit title.- `PR_TITLE` - default to the pull request's title.- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
// returns a *PullRequestWebhook_merge_commit_title when successful
func (m *PullRequestWebhook) GetMergeCommitTitle()(*PullRequestWebhook_merge_commit_title) {
    return m.merge_commit_title
}
// GetSquashMergeCommitMessage gets the squash_merge_commit_message property value. The default value for a squash merge commit message:- `PR_BODY` - default to the pull request's body.- `COMMIT_MESSAGES` - default to the branch's commit messages.- `BLANK` - default to a blank commit message.
// returns a *PullRequestWebhook_squash_merge_commit_message when successful
func (m *PullRequestWebhook) GetSquashMergeCommitMessage()(*PullRequestWebhook_squash_merge_commit_message) {
    return m.squash_merge_commit_message
}
// GetSquashMergeCommitTitle gets the squash_merge_commit_title property value. The default value for a squash merge commit title:- `PR_TITLE` - default to the pull request's title.- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
// returns a *PullRequestWebhook_squash_merge_commit_title when successful
func (m *PullRequestWebhook) GetSquashMergeCommitTitle()(*PullRequestWebhook_squash_merge_commit_title) {
    return m.squash_merge_commit_title
}
// GetUseSquashPrTitleAsDefault gets the use_squash_pr_title_as_default property value. Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.**
// returns a *bool when successful
func (m *PullRequestWebhook) GetUseSquashPrTitleAsDefault()(*bool) {
    return m.use_squash_pr_title_as_default
}
// Serialize serializes information the current object
func (m *PullRequestWebhook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PullRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allow_auto_merge", m.GetAllowAutoMerge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allow_update_branch", m.GetAllowUpdateBranch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("delete_branch_on_merge", m.GetDeleteBranchOnMerge())
        if err != nil {
            return err
        }
    }
    if m.GetMergeCommitMessage() != nil {
        cast := (*m.GetMergeCommitMessage()).String()
        err = writer.WriteStringValue("merge_commit_message", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMergeCommitTitle() != nil {
        cast := (*m.GetMergeCommitTitle()).String()
        err = writer.WriteStringValue("merge_commit_title", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSquashMergeCommitMessage() != nil {
        cast := (*m.GetSquashMergeCommitMessage()).String()
        err = writer.WriteStringValue("squash_merge_commit_message", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSquashMergeCommitTitle() != nil {
        cast := (*m.GetSquashMergeCommitTitle()).String()
        err = writer.WriteStringValue("squash_merge_commit_title", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("use_squash_pr_title_as_default", m.GetUseSquashPrTitleAsDefault())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAutoMerge sets the allow_auto_merge property value. Whether to allow auto-merge for pull requests.
func (m *PullRequestWebhook) SetAllowAutoMerge(value *bool)() {
    m.allow_auto_merge = value
}
// SetAllowUpdateBranch sets the allow_update_branch property value. Whether to allow updating the pull request's branch.
func (m *PullRequestWebhook) SetAllowUpdateBranch(value *bool)() {
    m.allow_update_branch = value
}
// SetDeleteBranchOnMerge sets the delete_branch_on_merge property value. Whether to delete head branches when pull requests are merged.
func (m *PullRequestWebhook) SetDeleteBranchOnMerge(value *bool)() {
    m.delete_branch_on_merge = value
}
// SetMergeCommitMessage sets the merge_commit_message property value. The default value for a merge commit message.- `PR_TITLE` - default to the pull request's title.- `PR_BODY` - default to the pull request's body.- `BLANK` - default to a blank commit message.
func (m *PullRequestWebhook) SetMergeCommitMessage(value *PullRequestWebhook_merge_commit_message)() {
    m.merge_commit_message = value
}
// SetMergeCommitTitle sets the merge_commit_title property value. The default value for a merge commit title.- `PR_TITLE` - default to the pull request's title.- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
func (m *PullRequestWebhook) SetMergeCommitTitle(value *PullRequestWebhook_merge_commit_title)() {
    m.merge_commit_title = value
}
// SetSquashMergeCommitMessage sets the squash_merge_commit_message property value. The default value for a squash merge commit message:- `PR_BODY` - default to the pull request's body.- `COMMIT_MESSAGES` - default to the branch's commit messages.- `BLANK` - default to a blank commit message.
func (m *PullRequestWebhook) SetSquashMergeCommitMessage(value *PullRequestWebhook_squash_merge_commit_message)() {
    m.squash_merge_commit_message = value
}
// SetSquashMergeCommitTitle sets the squash_merge_commit_title property value. The default value for a squash merge commit title:- `PR_TITLE` - default to the pull request's title.- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
func (m *PullRequestWebhook) SetSquashMergeCommitTitle(value *PullRequestWebhook_squash_merge_commit_title)() {
    m.squash_merge_commit_title = value
}
// SetUseSquashPrTitleAsDefault sets the use_squash_pr_title_as_default property value. Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.**
func (m *PullRequestWebhook) SetUseSquashPrTitleAsDefault(value *bool)() {
    m.use_squash_pr_title_as_default = value
}
type PullRequestWebhookable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PullRequestable
    GetAllowAutoMerge()(*bool)
    GetAllowUpdateBranch()(*bool)
    GetDeleteBranchOnMerge()(*bool)
    GetMergeCommitMessage()(*PullRequestWebhook_merge_commit_message)
    GetMergeCommitTitle()(*PullRequestWebhook_merge_commit_title)
    GetSquashMergeCommitMessage()(*PullRequestWebhook_squash_merge_commit_message)
    GetSquashMergeCommitTitle()(*PullRequestWebhook_squash_merge_commit_title)
    GetUseSquashPrTitleAsDefault()(*bool)
    SetAllowAutoMerge(value *bool)()
    SetAllowUpdateBranch(value *bool)()
    SetDeleteBranchOnMerge(value *bool)()
    SetMergeCommitMessage(value *PullRequestWebhook_merge_commit_message)()
    SetMergeCommitTitle(value *PullRequestWebhook_merge_commit_title)()
    SetSquashMergeCommitMessage(value *PullRequestWebhook_squash_merge_commit_message)()
    SetSquashMergeCommitTitle(value *PullRequestWebhook_squash_merge_commit_title)()
    SetUseSquashPrTitleAsDefault(value *bool)()
}
