package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequest pull requests let you tell others about changes you've pushed to a repository on GitHub. Once a pull request is sent, interested parties can review the set of changes, discuss potential modifications, and even push follow-up commits if necessary.
type PullRequest struct {
    // The _links property
    _links PullRequest__linksable
    // The active_lock_reason property
    active_lock_reason *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The additions property
    additions *int32
    // A GitHub user.
    assignee NullableSimpleUserable
    // The assignees property
    assignees []SimpleUserable
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The status of auto merging a pull request.
    auto_merge AutoMergeable
    // The base property
    base PullRequest_baseable
    // The body property
    body *string
    // The changed_files property
    changed_files *int32
    // The closed_at property
    closed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The comments property
    comments *int32
    // The comments_url property
    comments_url *string
    // The commits property
    commits *int32
    // The commits_url property
    commits_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deletions property
    deletions *int32
    // The diff_url property
    diff_url *string
    // Indicates whether or not the pull request is a draft.
    draft *bool
    // The head property
    head PullRequest_headable
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The issue_url property
    issue_url *string
    // The labels property
    labels []PullRequest_labelsable
    // The locked property
    locked *bool
    // Indicates whether maintainers can modify the pull request.
    maintainer_can_modify *bool
    // The merge_commit_sha property
    merge_commit_sha *string
    // The mergeable property
    mergeable *bool
    // The mergeable_state property
    mergeable_state *string
    // The merged property
    merged *bool
    // The merged_at property
    merged_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    merged_by NullableSimpleUserable
    // A collection of related issues and pull requests.
    milestone NullableMilestoneable
    // The node_id property
    node_id *string
    // Number uniquely identifying the pull request within its repository.
    number *int32
    // The patch_url property
    patch_url *string
    // The rebaseable property
    rebaseable *bool
    // The requested_reviewers property
    requested_reviewers []SimpleUserable
    // The requested_teams property
    requested_teams []TeamSimpleable
    // The review_comment_url property
    review_comment_url *string
    // The review_comments property
    review_comments *int32
    // The review_comments_url property
    review_comments_url *string
    // State of this Pull Request. Either `open` or `closed`.
    state *PullRequest_state
    // The statuses_url property
    statuses_url *string
    // The title of the pull request.
    title *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user SimpleUserable
}
// NewPullRequest instantiates a new PullRequest and sets the default values.
func NewPullRequest()(*PullRequest) {
    m := &PullRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequest(), nil
}
// GetActiveLockReason gets the active_lock_reason property value. The active_lock_reason property
// returns a *string when successful
func (m *PullRequest) GetActiveLockReason()(*string) {
    return m.active_lock_reason
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdditions gets the additions property value. The additions property
// returns a *int32 when successful
func (m *PullRequest) GetAdditions()(*int32) {
    return m.additions
}
// GetAssignee gets the assignee property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *PullRequest) GetAssignee()(NullableSimpleUserable) {
    return m.assignee
}
// GetAssignees gets the assignees property value. The assignees property
// returns a []SimpleUserable when successful
func (m *PullRequest) GetAssignees()([]SimpleUserable) {
    return m.assignees
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *PullRequest) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetAutoMerge gets the auto_merge property value. The status of auto merging a pull request.
// returns a AutoMergeable when successful
func (m *PullRequest) GetAutoMerge()(AutoMergeable) {
    return m.auto_merge
}
// GetBase gets the base property value. The base property
// returns a PullRequest_baseable when successful
func (m *PullRequest) GetBase()(PullRequest_baseable) {
    return m.base
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *PullRequest) GetBody()(*string) {
    return m.body
}
// GetChangedFiles gets the changed_files property value. The changed_files property
// returns a *int32 when successful
func (m *PullRequest) GetChangedFiles()(*int32) {
    return m.changed_files
}
// GetClosedAt gets the closed_at property value. The closed_at property
// returns a *Time when successful
func (m *PullRequest) GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closed_at
}
// GetComments gets the comments property value. The comments property
// returns a *int32 when successful
func (m *PullRequest) GetComments()(*int32) {
    return m.comments
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *PullRequest) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommits gets the commits property value. The commits property
// returns a *int32 when successful
func (m *PullRequest) GetCommits()(*int32) {
    return m.commits
}
// GetCommitsUrl gets the commits_url property value. The commits_url property
// returns a *string when successful
func (m *PullRequest) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *PullRequest) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDeletions gets the deletions property value. The deletions property
// returns a *int32 when successful
func (m *PullRequest) GetDeletions()(*int32) {
    return m.deletions
}
// GetDiffUrl gets the diff_url property value. The diff_url property
// returns a *string when successful
func (m *PullRequest) GetDiffUrl()(*string) {
    return m.diff_url
}
// GetDraft gets the draft property value. Indicates whether or not the pull request is a draft.
// returns a *bool when successful
func (m *PullRequest) GetDraft()(*bool) {
    return m.draft
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequest__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(PullRequest__linksable))
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
    res["additions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditions(val)
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
        val, err := n.GetObjectValue(CreatePullRequest_baseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBase(val.(PullRequest_baseable))
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
    res["changed_files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangedFiles(val)
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
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val)
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
    res["commits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommits(val)
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
    res["deletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletions(val)
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
        val, err := n.GetObjectValue(CreatePullRequest_headFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHead(val.(PullRequest_headable))
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
        val, err := n.GetCollectionOfObjectValues(CreatePullRequest_labelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullRequest_labelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PullRequest_labelsable)
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
    res["maintainer_can_modify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintainerCanModify(val)
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
    res["mergeable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeable(val)
        }
        return nil
    }
    res["mergeable_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeableState(val)
        }
        return nil
    }
    res["merged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMerged(val)
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
    res["merged_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergedBy(val.(NullableSimpleUserable))
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
    res["rebaseable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRebaseable(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateTeamSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamSimpleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamSimpleable)
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
    res["review_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewComments(val)
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
        val, err := n.GetEnumValue(ParsePullRequest_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PullRequest_state))
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
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(SimpleUserable))
        }
        return nil
    }
    return res
}
// GetHead gets the head property value. The head property
// returns a PullRequest_headable when successful
func (m *PullRequest) GetHead()(PullRequest_headable) {
    return m.head
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *PullRequest) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *PullRequest) GetId()(*int64) {
    return m.id
}
// GetIssueUrl gets the issue_url property value. The issue_url property
// returns a *string when successful
func (m *PullRequest) GetIssueUrl()(*string) {
    return m.issue_url
}
// GetLabels gets the labels property value. The labels property
// returns a []PullRequest_labelsable when successful
func (m *PullRequest) GetLabels()([]PullRequest_labelsable) {
    return m.labels
}
// GetLinks gets the _links property value. The _links property
// returns a PullRequest__linksable when successful
func (m *PullRequest) GetLinks()(PullRequest__linksable) {
    return m._links
}
// GetLocked gets the locked property value. The locked property
// returns a *bool when successful
func (m *PullRequest) GetLocked()(*bool) {
    return m.locked
}
// GetMaintainerCanModify gets the maintainer_can_modify property value. Indicates whether maintainers can modify the pull request.
// returns a *bool when successful
func (m *PullRequest) GetMaintainerCanModify()(*bool) {
    return m.maintainer_can_modify
}
// GetMergeable gets the mergeable property value. The mergeable property
// returns a *bool when successful
func (m *PullRequest) GetMergeable()(*bool) {
    return m.mergeable
}
// GetMergeableState gets the mergeable_state property value. The mergeable_state property
// returns a *string when successful
func (m *PullRequest) GetMergeableState()(*string) {
    return m.mergeable_state
}
// GetMergeCommitSha gets the merge_commit_sha property value. The merge_commit_sha property
// returns a *string when successful
func (m *PullRequest) GetMergeCommitSha()(*string) {
    return m.merge_commit_sha
}
// GetMerged gets the merged property value. The merged property
// returns a *bool when successful
func (m *PullRequest) GetMerged()(*bool) {
    return m.merged
}
// GetMergedAt gets the merged_at property value. The merged_at property
// returns a *Time when successful
func (m *PullRequest) GetMergedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.merged_at
}
// GetMergedBy gets the merged_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *PullRequest) GetMergedBy()(NullableSimpleUserable) {
    return m.merged_by
}
// GetMilestone gets the milestone property value. A collection of related issues and pull requests.
// returns a NullableMilestoneable when successful
func (m *PullRequest) GetMilestone()(NullableMilestoneable) {
    return m.milestone
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *PullRequest) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. Number uniquely identifying the pull request within its repository.
// returns a *int32 when successful
func (m *PullRequest) GetNumber()(*int32) {
    return m.number
}
// GetPatchUrl gets the patch_url property value. The patch_url property
// returns a *string when successful
func (m *PullRequest) GetPatchUrl()(*string) {
    return m.patch_url
}
// GetRebaseable gets the rebaseable property value. The rebaseable property
// returns a *bool when successful
func (m *PullRequest) GetRebaseable()(*bool) {
    return m.rebaseable
}
// GetRequestedReviewers gets the requested_reviewers property value. The requested_reviewers property
// returns a []SimpleUserable when successful
func (m *PullRequest) GetRequestedReviewers()([]SimpleUserable) {
    return m.requested_reviewers
}
// GetRequestedTeams gets the requested_teams property value. The requested_teams property
// returns a []TeamSimpleable when successful
func (m *PullRequest) GetRequestedTeams()([]TeamSimpleable) {
    return m.requested_teams
}
// GetReviewComments gets the review_comments property value. The review_comments property
// returns a *int32 when successful
func (m *PullRequest) GetReviewComments()(*int32) {
    return m.review_comments
}
// GetReviewCommentsUrl gets the review_comments_url property value. The review_comments_url property
// returns a *string when successful
func (m *PullRequest) GetReviewCommentsUrl()(*string) {
    return m.review_comments_url
}
// GetReviewCommentUrl gets the review_comment_url property value. The review_comment_url property
// returns a *string when successful
func (m *PullRequest) GetReviewCommentUrl()(*string) {
    return m.review_comment_url
}
// GetState gets the state property value. State of this Pull Request. Either `open` or `closed`.
// returns a *PullRequest_state when successful
func (m *PullRequest) GetState()(*PullRequest_state) {
    return m.state
}
// GetStatusesUrl gets the statuses_url property value. The statuses_url property
// returns a *string when successful
func (m *PullRequest) GetStatusesUrl()(*string) {
    return m.statuses_url
}
// GetTitle gets the title property value. The title of the pull request.
// returns a *string when successful
func (m *PullRequest) GetTitle()(*string) {
    return m.title
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *PullRequest) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PullRequest) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *PullRequest) GetUser()(SimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *PullRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("active_lock_reason", m.GetActiveLockReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("additions", m.GetAdditions())
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
        err := writer.WriteInt32Value("changed_files", m.GetChangedFiles())
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
        err := writer.WriteInt32Value("comments", m.GetComments())
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
        err := writer.WriteInt32Value("commits", m.GetCommits())
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
        err := writer.WriteInt32Value("deletions", m.GetDeletions())
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
        err := writer.WriteBoolValue("maintainer_can_modify", m.GetMaintainerCanModify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mergeable", m.GetMergeable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mergeable_state", m.GetMergeableState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("merged", m.GetMerged())
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
        err := writer.WriteObjectValue("merged_by", m.GetMergedBy())
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
    {
        err := writer.WriteBoolValue("rebaseable", m.GetRebaseable())
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
        err := writer.WriteInt32Value("review_comments", m.GetReviewComments())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *PullRequest) SetActiveLockReason(value *string)() {
    m.active_lock_reason = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdditions sets the additions property value. The additions property
func (m *PullRequest) SetAdditions(value *int32)() {
    m.additions = value
}
// SetAssignee sets the assignee property value. A GitHub user.
func (m *PullRequest) SetAssignee(value NullableSimpleUserable)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. The assignees property
func (m *PullRequest) SetAssignees(value []SimpleUserable)() {
    m.assignees = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *PullRequest) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetAutoMerge sets the auto_merge property value. The status of auto merging a pull request.
func (m *PullRequest) SetAutoMerge(value AutoMergeable)() {
    m.auto_merge = value
}
// SetBase sets the base property value. The base property
func (m *PullRequest) SetBase(value PullRequest_baseable)() {
    m.base = value
}
// SetBody sets the body property value. The body property
func (m *PullRequest) SetBody(value *string)() {
    m.body = value
}
// SetChangedFiles sets the changed_files property value. The changed_files property
func (m *PullRequest) SetChangedFiles(value *int32)() {
    m.changed_files = value
}
// SetClosedAt sets the closed_at property value. The closed_at property
func (m *PullRequest) SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closed_at = value
}
// SetComments sets the comments property value. The comments property
func (m *PullRequest) SetComments(value *int32)() {
    m.comments = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *PullRequest) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommits sets the commits property value. The commits property
func (m *PullRequest) SetCommits(value *int32)() {
    m.commits = value
}
// SetCommitsUrl sets the commits_url property value. The commits_url property
func (m *PullRequest) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *PullRequest) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDeletions sets the deletions property value. The deletions property
func (m *PullRequest) SetDeletions(value *int32)() {
    m.deletions = value
}
// SetDiffUrl sets the diff_url property value. The diff_url property
func (m *PullRequest) SetDiffUrl(value *string)() {
    m.diff_url = value
}
// SetDraft sets the draft property value. Indicates whether or not the pull request is a draft.
func (m *PullRequest) SetDraft(value *bool)() {
    m.draft = value
}
// SetHead sets the head property value. The head property
func (m *PullRequest) SetHead(value PullRequest_headable)() {
    m.head = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *PullRequest) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *PullRequest) SetId(value *int64)() {
    m.id = value
}
// SetIssueUrl sets the issue_url property value. The issue_url property
func (m *PullRequest) SetIssueUrl(value *string)() {
    m.issue_url = value
}
// SetLabels sets the labels property value. The labels property
func (m *PullRequest) SetLabels(value []PullRequest_labelsable)() {
    m.labels = value
}
// SetLinks sets the _links property value. The _links property
func (m *PullRequest) SetLinks(value PullRequest__linksable)() {
    m._links = value
}
// SetLocked sets the locked property value. The locked property
func (m *PullRequest) SetLocked(value *bool)() {
    m.locked = value
}
// SetMaintainerCanModify sets the maintainer_can_modify property value. Indicates whether maintainers can modify the pull request.
func (m *PullRequest) SetMaintainerCanModify(value *bool)() {
    m.maintainer_can_modify = value
}
// SetMergeable sets the mergeable property value. The mergeable property
func (m *PullRequest) SetMergeable(value *bool)() {
    m.mergeable = value
}
// SetMergeableState sets the mergeable_state property value. The mergeable_state property
func (m *PullRequest) SetMergeableState(value *string)() {
    m.mergeable_state = value
}
// SetMergeCommitSha sets the merge_commit_sha property value. The merge_commit_sha property
func (m *PullRequest) SetMergeCommitSha(value *string)() {
    m.merge_commit_sha = value
}
// SetMerged sets the merged property value. The merged property
func (m *PullRequest) SetMerged(value *bool)() {
    m.merged = value
}
// SetMergedAt sets the merged_at property value. The merged_at property
func (m *PullRequest) SetMergedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.merged_at = value
}
// SetMergedBy sets the merged_by property value. A GitHub user.
func (m *PullRequest) SetMergedBy(value NullableSimpleUserable)() {
    m.merged_by = value
}
// SetMilestone sets the milestone property value. A collection of related issues and pull requests.
func (m *PullRequest) SetMilestone(value NullableMilestoneable)() {
    m.milestone = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *PullRequest) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. Number uniquely identifying the pull request within its repository.
func (m *PullRequest) SetNumber(value *int32)() {
    m.number = value
}
// SetPatchUrl sets the patch_url property value. The patch_url property
func (m *PullRequest) SetPatchUrl(value *string)() {
    m.patch_url = value
}
// SetRebaseable sets the rebaseable property value. The rebaseable property
func (m *PullRequest) SetRebaseable(value *bool)() {
    m.rebaseable = value
}
// SetRequestedReviewers sets the requested_reviewers property value. The requested_reviewers property
func (m *PullRequest) SetRequestedReviewers(value []SimpleUserable)() {
    m.requested_reviewers = value
}
// SetRequestedTeams sets the requested_teams property value. The requested_teams property
func (m *PullRequest) SetRequestedTeams(value []TeamSimpleable)() {
    m.requested_teams = value
}
// SetReviewComments sets the review_comments property value. The review_comments property
func (m *PullRequest) SetReviewComments(value *int32)() {
    m.review_comments = value
}
// SetReviewCommentsUrl sets the review_comments_url property value. The review_comments_url property
func (m *PullRequest) SetReviewCommentsUrl(value *string)() {
    m.review_comments_url = value
}
// SetReviewCommentUrl sets the review_comment_url property value. The review_comment_url property
func (m *PullRequest) SetReviewCommentUrl(value *string)() {
    m.review_comment_url = value
}
// SetState sets the state property value. State of this Pull Request. Either `open` or `closed`.
func (m *PullRequest) SetState(value *PullRequest_state)() {
    m.state = value
}
// SetStatusesUrl sets the statuses_url property value. The statuses_url property
func (m *PullRequest) SetStatusesUrl(value *string)() {
    m.statuses_url = value
}
// SetTitle sets the title property value. The title of the pull request.
func (m *PullRequest) SetTitle(value *string)() {
    m.title = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *PullRequest) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *PullRequest) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *PullRequest) SetUser(value SimpleUserable)() {
    m.user = value
}
type PullRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveLockReason()(*string)
    GetAdditions()(*int32)
    GetAssignee()(NullableSimpleUserable)
    GetAssignees()([]SimpleUserable)
    GetAuthorAssociation()(*AuthorAssociation)
    GetAutoMerge()(AutoMergeable)
    GetBase()(PullRequest_baseable)
    GetBody()(*string)
    GetChangedFiles()(*int32)
    GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCommits()(*int32)
    GetCommitsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeletions()(*int32)
    GetDiffUrl()(*string)
    GetDraft()(*bool)
    GetHead()(PullRequest_headable)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetIssueUrl()(*string)
    GetLabels()([]PullRequest_labelsable)
    GetLinks()(PullRequest__linksable)
    GetLocked()(*bool)
    GetMaintainerCanModify()(*bool)
    GetMergeable()(*bool)
    GetMergeableState()(*string)
    GetMergeCommitSha()(*string)
    GetMerged()(*bool)
    GetMergedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMergedBy()(NullableSimpleUserable)
    GetMilestone()(NullableMilestoneable)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetPatchUrl()(*string)
    GetRebaseable()(*bool)
    GetRequestedReviewers()([]SimpleUserable)
    GetRequestedTeams()([]TeamSimpleable)
    GetReviewComments()(*int32)
    GetReviewCommentsUrl()(*string)
    GetReviewCommentUrl()(*string)
    GetState()(*PullRequest_state)
    GetStatusesUrl()(*string)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(SimpleUserable)
    SetActiveLockReason(value *string)()
    SetAdditions(value *int32)()
    SetAssignee(value NullableSimpleUserable)()
    SetAssignees(value []SimpleUserable)()
    SetAuthorAssociation(value *AuthorAssociation)()
    SetAutoMerge(value AutoMergeable)()
    SetBase(value PullRequest_baseable)()
    SetBody(value *string)()
    SetChangedFiles(value *int32)()
    SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCommits(value *int32)()
    SetCommitsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeletions(value *int32)()
    SetDiffUrl(value *string)()
    SetDraft(value *bool)()
    SetHead(value PullRequest_headable)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetIssueUrl(value *string)()
    SetLabels(value []PullRequest_labelsable)()
    SetLinks(value PullRequest__linksable)()
    SetLocked(value *bool)()
    SetMaintainerCanModify(value *bool)()
    SetMergeable(value *bool)()
    SetMergeableState(value *string)()
    SetMergeCommitSha(value *string)()
    SetMerged(value *bool)()
    SetMergedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMergedBy(value NullableSimpleUserable)()
    SetMilestone(value NullableMilestoneable)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetPatchUrl(value *string)()
    SetRebaseable(value *bool)()
    SetRequestedReviewers(value []SimpleUserable)()
    SetRequestedTeams(value []TeamSimpleable)()
    SetReviewComments(value *int32)()
    SetReviewCommentsUrl(value *string)()
    SetReviewCommentUrl(value *string)()
    SetState(value *PullRequest_state)()
    SetStatusesUrl(value *string)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value SimpleUserable)()
}
