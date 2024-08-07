package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestSimple pull Request Simple
type PullRequestSimple struct {
    // The _links property
    _links PullRequestSimple__linksable
    // The active_lock_reason property
    active_lock_reason *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    assignee NullableSimpleUserable
    // The assignees property
    assignees []SimpleUserable
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The status of auto merging a pull request.
    auto_merge AutoMergeable
    // The base property
    base PullRequestSimple_baseable
    // The body property
    body *string
    // The closed_at property
    closed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The comments_url property
    comments_url *string
    // The commits_url property
    commits_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The diff_url property
    diff_url *string
    // Indicates whether or not the pull request is a draft.
    draft *bool
    // The head property
    head PullRequestSimple_headable
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The issue_url property
    issue_url *string
    // The labels property
    labels []PullRequestSimple_labelsable
    // The locked property
    locked *bool
    // The merge_commit_sha property
    merge_commit_sha *string
    // The merged_at property
    merged_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of related issues and pull requests.
    milestone NullableMilestoneable
    // The node_id property
    node_id *string
    // The number property
    number *int32
    // The patch_url property
    patch_url *string
    // The requested_reviewers property
    requested_reviewers []SimpleUserable
    // The requested_teams property
    requested_teams []Teamable
    // The review_comment_url property
    review_comment_url *string
    // The review_comments_url property
    review_comments_url *string
    // The state property
    state *string
    // The statuses_url property
    statuses_url *string
    // The title property
    title *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewPullRequestSimple instantiates a new PullRequestSimple and sets the default values.
func NewPullRequestSimple()(*PullRequestSimple) {
    m := &PullRequestSimple{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestSimpleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestSimpleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestSimple(), nil
}
// GetActiveLockReason gets the active_lock_reason property value. The active_lock_reason property
// returns a *string when successful
func (m *PullRequestSimple) GetActiveLockReason()(*string) {
    return m.active_lock_reason
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestSimple) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *PullRequestSimple) GetAssignee()(NullableSimpleUserable) {
    return m.assignee
}
// GetAssignees gets the assignees property value. The assignees property
// returns a []SimpleUserable when successful
func (m *PullRequestSimple) GetAssignees()([]SimpleUserable) {
    return m.assignees
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *PullRequestSimple) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetAutoMerge gets the auto_merge property value. The status of auto merging a pull request.
// returns a AutoMergeable when successful
func (m *PullRequestSimple) GetAutoMerge()(AutoMergeable) {
    return m.auto_merge
}
// GetBase gets the base property value. The base property
// returns a PullRequestSimple_baseable when successful
func (m *PullRequestSimple) GetBase()(PullRequestSimple_baseable) {
    return m.base
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *PullRequestSimple) GetBody()(*string) {
    return m.body
}
// GetClosedAt gets the closed_at property value. The closed_at property
// returns a *Time when successful
func (m *PullRequestSimple) GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closed_at
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *PullRequestSimple) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommitsUrl gets the commits_url property value. The commits_url property
// returns a *string when successful
func (m *PullRequestSimple) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *PullRequestSimple) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDiffUrl gets the diff_url property value. The diff_url property
// returns a *string when successful
func (m *PullRequestSimple) GetDiffUrl()(*string) {
    return m.diff_url
}
// GetDraft gets the draft property value. Indicates whether or not the pull request is a draft.
// returns a *bool when successful
func (m *PullRequestSimple) GetDraft()(*bool) {
    return m.draft
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestSimple) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestSimple__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(PullRequestSimple__linksable))
        }
        return nil
    }
    res["active_lock_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveLockReason(val)
        }
        return nil
    }
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["assignees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleUserable)
                }
            }
            m.SetAssignees(res)
        }
        return nil
    }
    res["author_association"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthorAssociation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorAssociation(val.(*AuthorAssociation))
        }
        return nil
    }
    res["auto_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutoMergeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoMerge(val.(AutoMergeable))
        }
        return nil
    }
    res["base"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestSimple_baseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBase(val.(PullRequestSimple_baseable))
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["closed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedAt(val)
        }
        return nil
    }
    res["comments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsUrl(val)
        }
        return nil
    }
    res["commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitsUrl(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["diff_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiffUrl(val)
        }
        return nil
    }
    res["draft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraft(val)
        }
        return nil
    }
    res["head"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestSimple_headFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHead(val.(PullRequestSimple_headable))
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["issue_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueUrl(val)
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePullRequestSimple_labelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullRequestSimple_labelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PullRequestSimple_labelsable)
                }
            }
            m.SetLabels(res)
        }
        return nil
    }
    res["locked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocked(val)
        }
        return nil
    }
    res["merge_commit_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeCommitSha(val)
        }
        return nil
    }
    res["merged_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergedAt(val)
        }
        return nil
    }
    res["milestone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableMilestoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestone(val.(NullableMilestoneable))
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["patch_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatchUrl(val)
        }
        return nil
    }
    res["requested_reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleUserable)
                }
            }
            m.SetRequestedReviewers(res)
        }
        return nil
    }
    res["requested_teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Teamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Teamable)
                }
            }
            m.SetRequestedTeams(res)
        }
        return nil
    }
    res["review_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewCommentUrl(val)
        }
        return nil
    }
    res["review_comments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewCommentsUrl(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["statuses_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusesUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    return res
}
// GetHead gets the head property value. The head property
// returns a PullRequestSimple_headable when successful
func (m *PullRequestSimple) GetHead()(PullRequestSimple_headable) {
    return m.head
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *PullRequestSimple) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *PullRequestSimple) GetId()(*int64) {
    return m.id
}
// GetIssueUrl gets the issue_url property value. The issue_url property
// returns a *string when successful
func (m *PullRequestSimple) GetIssueUrl()(*string) {
    return m.issue_url
}
// GetLabels gets the labels property value. The labels property
// returns a []PullRequestSimple_labelsable when successful
func (m *PullRequestSimple) GetLabels()([]PullRequestSimple_labelsable) {
    return m.labels
}
// GetLinks gets the _links property value. The _links property
// returns a PullRequestSimple__linksable when successful
func (m *PullRequestSimple) GetLinks()(PullRequestSimple__linksable) {
    return m._links
}
// GetLocked gets the locked property value. The locked property
// returns a *bool when successful
func (m *PullRequestSimple) GetLocked()(*bool) {
    return m.locked
}
// GetMergeCommitSha gets the merge_commit_sha property value. The merge_commit_sha property
// returns a *string when successful
func (m *PullRequestSimple) GetMergeCommitSha()(*string) {
    return m.merge_commit_sha
}
// GetMergedAt gets the merged_at property value. The merged_at property
// returns a *Time when successful
func (m *PullRequestSimple) GetMergedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.merged_at
}
// GetMilestone gets the milestone property value. A collection of related issues and pull requests.
// returns a NullableMilestoneable when successful
func (m *PullRequestSimple) GetMilestone()(NullableMilestoneable) {
    return m.milestone
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *PullRequestSimple) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *PullRequestSimple) GetNumber()(*int32) {
    return m.number
}
// GetPatchUrl gets the patch_url property value. The patch_url property
// returns a *string when successful
func (m *PullRequestSimple) GetPatchUrl()(*string) {
    return m.patch_url
}
// GetRequestedReviewers gets the requested_reviewers property value. The requested_reviewers property
// returns a []SimpleUserable when successful
func (m *PullRequestSimple) GetRequestedReviewers()([]SimpleUserable) {
    return m.requested_reviewers
}
// GetRequestedTeams gets the requested_teams property value. The requested_teams property
// returns a []Teamable when successful
func (m *PullRequestSimple) GetRequestedTeams()([]Teamable) {
    return m.requested_teams
}
// GetReviewCommentsUrl gets the review_comments_url property value. The review_comments_url property
// returns a *string when successful
func (m *PullRequestSimple) GetReviewCommentsUrl()(*string) {
    return m.review_comments_url
}
// GetReviewCommentUrl gets the review_comment_url property value. The review_comment_url property
// returns a *string when successful
func (m *PullRequestSimple) GetReviewCommentUrl()(*string) {
    return m.review_comment_url
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *PullRequestSimple) GetState()(*string) {
    return m.state
}
// GetStatusesUrl gets the statuses_url property value. The statuses_url property
// returns a *string when successful
func (m *PullRequestSimple) GetStatusesUrl()(*string) {
    return m.statuses_url
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *PullRequestSimple) GetTitle()(*string) {
    return m.title
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *PullRequestSimple) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PullRequestSimple) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *PullRequestSimple) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *PullRequestSimple) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("active_lock_reason", m.GetActiveLockReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    if m.GetAssignees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignees()))
        for i, v := range m.GetAssignees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("assignees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthorAssociation() != nil {
        cast := (*m.GetAuthorAssociation()).String()
        err := writer.WriteStringValue("author_association", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("auto_merge", m.GetAutoMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("closed_at", m.GetClosedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comments_url", m.GetCommentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commits_url", m.GetCommitsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("diff_url", m.GetDiffUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("draft", m.GetDraft())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("head", m.GetHead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue_url", m.GetIssueUrl())
        if err != nil {
            return err
        }
    }
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("locked", m.GetLocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("merged_at", m.GetMergedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("merge_commit_sha", m.GetMergeCommitSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("milestone", m.GetMilestone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("patch_url", m.GetPatchUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRequestedReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequestedReviewers()))
        for i, v := range m.GetRequestedReviewers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("requested_reviewers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequestedTeams()))
        for i, v := range m.GetRequestedTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("requested_teams", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("review_comments_url", m.GetReviewCommentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("review_comment_url", m.GetReviewCommentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("statuses_url", m.GetStatusesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_links", m.GetLinks())
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
// SetActiveLockReason sets the active_lock_reason property value. The active_lock_reason property
func (m *PullRequestSimple) SetActiveLockReason(value *string)() {
    m.active_lock_reason = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullRequestSimple) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. A GitHub user.
func (m *PullRequestSimple) SetAssignee(value NullableSimpleUserable)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. The assignees property
func (m *PullRequestSimple) SetAssignees(value []SimpleUserable)() {
    m.assignees = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *PullRequestSimple) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetAutoMerge sets the auto_merge property value. The status of auto merging a pull request.
func (m *PullRequestSimple) SetAutoMerge(value AutoMergeable)() {
    m.auto_merge = value
}
// SetBase sets the base property value. The base property
func (m *PullRequestSimple) SetBase(value PullRequestSimple_baseable)() {
    m.base = value
}
// SetBody sets the body property value. The body property
func (m *PullRequestSimple) SetBody(value *string)() {
    m.body = value
}
// SetClosedAt sets the closed_at property value. The closed_at property
func (m *PullRequestSimple) SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closed_at = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *PullRequestSimple) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommitsUrl sets the commits_url property value. The commits_url property
func (m *PullRequestSimple) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *PullRequestSimple) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDiffUrl sets the diff_url property value. The diff_url property
func (m *PullRequestSimple) SetDiffUrl(value *string)() {
    m.diff_url = value
}
// SetDraft sets the draft property value. Indicates whether or not the pull request is a draft.
func (m *PullRequestSimple) SetDraft(value *bool)() {
    m.draft = value
}
// SetHead sets the head property value. The head property
func (m *PullRequestSimple) SetHead(value PullRequestSimple_headable)() {
    m.head = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *PullRequestSimple) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *PullRequestSimple) SetId(value *int64)() {
    m.id = value
}
// SetIssueUrl sets the issue_url property value. The issue_url property
func (m *PullRequestSimple) SetIssueUrl(value *string)() {
    m.issue_url = value
}
// SetLabels sets the labels property value. The labels property
func (m *PullRequestSimple) SetLabels(value []PullRequestSimple_labelsable)() {
    m.labels = value
}
// SetLinks sets the _links property value. The _links property
func (m *PullRequestSimple) SetLinks(value PullRequestSimple__linksable)() {
    m._links = value
}
// SetLocked sets the locked property value. The locked property
func (m *PullRequestSimple) SetLocked(value *bool)() {
    m.locked = value
}
// SetMergeCommitSha sets the merge_commit_sha property value. The merge_commit_sha property
func (m *PullRequestSimple) SetMergeCommitSha(value *string)() {
    m.merge_commit_sha = value
}
// SetMergedAt sets the merged_at property value. The merged_at property
func (m *PullRequestSimple) SetMergedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.merged_at = value
}
// SetMilestone sets the milestone property value. A collection of related issues and pull requests.
func (m *PullRequestSimple) SetMilestone(value NullableMilestoneable)() {
    m.milestone = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *PullRequestSimple) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The number property
func (m *PullRequestSimple) SetNumber(value *int32)() {
    m.number = value
}
// SetPatchUrl sets the patch_url property value. The patch_url property
func (m *PullRequestSimple) SetPatchUrl(value *string)() {
    m.patch_url = value
}
// SetRequestedReviewers sets the requested_reviewers property value. The requested_reviewers property
func (m *PullRequestSimple) SetRequestedReviewers(value []SimpleUserable)() {
    m.requested_reviewers = value
}
// SetRequestedTeams sets the requested_teams property value. The requested_teams property
func (m *PullRequestSimple) SetRequestedTeams(value []Teamable)() {
    m.requested_teams = value
}
// SetReviewCommentsUrl sets the review_comments_url property value. The review_comments_url property
func (m *PullRequestSimple) SetReviewCommentsUrl(value *string)() {
    m.review_comments_url = value
}
// SetReviewCommentUrl sets the review_comment_url property value. The review_comment_url property
func (m *PullRequestSimple) SetReviewCommentUrl(value *string)() {
    m.review_comment_url = value
}
// SetState sets the state property value. The state property
func (m *PullRequestSimple) SetState(value *string)() {
    m.state = value
}
// SetStatusesUrl sets the statuses_url property value. The statuses_url property
func (m *PullRequestSimple) SetStatusesUrl(value *string)() {
    m.statuses_url = value
}
// SetTitle sets the title property value. The title property
func (m *PullRequestSimple) SetTitle(value *string)() {
    m.title = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *PullRequestSimple) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *PullRequestSimple) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *PullRequestSimple) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type PullRequestSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveLockReason()(*string)
    GetAssignee()(NullableSimpleUserable)
    GetAssignees()([]SimpleUserable)
    GetAuthorAssociation()(*AuthorAssociation)
    GetAutoMerge()(AutoMergeable)
    GetBase()(PullRequestSimple_baseable)
    GetBody()(*string)
    GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiffUrl()(*string)
    GetDraft()(*bool)
    GetHead()(PullRequestSimple_headable)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetIssueUrl()(*string)
    GetLabels()([]PullRequestSimple_labelsable)
    GetLinks()(PullRequestSimple__linksable)
    GetLocked()(*bool)
    GetMergeCommitSha()(*string)
    GetMergedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMilestone()(NullableMilestoneable)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetPatchUrl()(*string)
    GetRequestedReviewers()([]SimpleUserable)
    GetRequestedTeams()([]Teamable)
    GetReviewCommentsUrl()(*string)
    GetReviewCommentUrl()(*string)
    GetState()(*string)
    GetStatusesUrl()(*string)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetActiveLockReason(value *string)()
    SetAssignee(value NullableSimpleUserable)()
    SetAssignees(value []SimpleUserable)()
    SetAuthorAssociation(value *AuthorAssociation)()
    SetAutoMerge(value AutoMergeable)()
    SetBase(value PullRequestSimple_baseable)()
    SetBody(value *string)()
    SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiffUrl(value *string)()
    SetDraft(value *bool)()
    SetHead(value PullRequestSimple_headable)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetIssueUrl(value *string)()
    SetLabels(value []PullRequestSimple_labelsable)()
    SetLinks(value PullRequestSimple__linksable)()
    SetLocked(value *bool)()
    SetMergeCommitSha(value *string)()
    SetMergedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMilestone(value NullableMilestoneable)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetPatchUrl(value *string)()
    SetRequestedReviewers(value []SimpleUserable)()
    SetRequestedTeams(value []Teamable)()
    SetReviewCommentsUrl(value *string)()
    SetReviewCommentUrl(value *string)()
    SetState(value *string)()
    SetStatusesUrl(value *string)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
