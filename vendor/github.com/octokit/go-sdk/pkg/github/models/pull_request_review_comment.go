package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestReviewComment pull Request Review Comments are comments on a portion of the Pull Request's diff.
type PullRequestReviewComment struct {
    // The _links property
    _links PullRequestReviewComment__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The text of the comment.
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // The SHA of the commit to which the comment applies.
    commit_id *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The diff of the line that the comment refers to.
    diff_hunk *string
    // HTML URL for the pull request review comment.
    html_url *string
    // The ID of the pull request review comment.
    id *int64
    // The comment ID to reply to.
    in_reply_to_id *int32
    // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
    line *int32
    // The node ID of the pull request review comment.
    node_id *string
    // The SHA of the original commit to which the comment applies.
    original_commit_id *string
    // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
    original_line *int32
    // The index of the original line in the diff to which the comment applies. This field is deprecated; use `original_line` instead.
    original_position *int32
    // The first line of the range for a multi-line comment.
    original_start_line *int32
    // The relative path of the file to which the comment applies.
    path *string
    // The line index in the diff to which the comment applies. This field is deprecated; use `line` instead.
    position *int32
    // The ID of the pull request review to which the comment belongs.
    pull_request_review_id *int64
    // URL for the pull request that the review comment belongs to.
    pull_request_url *string
    // The reactions property
    reactions ReactionRollupable
    // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
    side *PullRequestReviewComment_side
    // The first line of the range for a multi-line comment.
    start_line *int32
    // The side of the first line of the range for a multi-line comment.
    start_side *PullRequestReviewComment_start_side
    // The level at which the comment is targeted, can be a diff line or a file.
    subject_type *PullRequestReviewComment_subject_type
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // URL for the pull request review comment
    url *string
    // A GitHub user.
    user SimpleUserable
}
// NewPullRequestReviewComment instantiates a new PullRequestReviewComment and sets the default values.
func NewPullRequestReviewComment()(*PullRequestReviewComment) {
    m := &PullRequestReviewComment{
    }
    m.SetAdditionalData(make(map[string]any))
    sideValue := RIGHT_PULLREQUESTREVIEWCOMMENT_SIDE
    m.SetSide(&sideValue)
    start_sideValue := RIGHT_PULLREQUESTREVIEWCOMMENT_START_SIDE
    m.SetStartSide(&start_sideValue)
    return m
}
// CreatePullRequestReviewCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestReviewCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestReviewComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestReviewComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *PullRequestReviewComment) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. The text of the comment.
// returns a *string when successful
func (m *PullRequestReviewComment) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *PullRequestReviewComment) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *PullRequestReviewComment) GetBodyText()(*string) {
    return m.body_text
}
// GetCommitId gets the commit_id property value. The SHA of the commit to which the comment applies.
// returns a *string when successful
func (m *PullRequestReviewComment) GetCommitId()(*string) {
    return m.commit_id
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *PullRequestReviewComment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDiffHunk gets the diff_hunk property value. The diff of the line that the comment refers to.
// returns a *string when successful
func (m *PullRequestReviewComment) GetDiffHunk()(*string) {
    return m.diff_hunk
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestReviewComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestReviewComment__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(PullRequestReviewComment__linksable))
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
    res["body_html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyHtml(val)
        }
        return nil
    }
    res["body_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyText(val)
        }
        return nil
    }
    res["commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitId(val)
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
    res["diff_hunk"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiffHunk(val)
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
    res["in_reply_to_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInReplyToId(val)
        }
        return nil
    }
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val)
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
    res["original_commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalCommitId(val)
        }
        return nil
    }
    res["original_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalLine(val)
        }
        return nil
    }
    res["original_position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalPosition(val)
        }
        return nil
    }
    res["original_start_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalStartLine(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["pull_request_review_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestReviewId(val)
        }
        return nil
    }
    res["pull_request_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestUrl(val)
        }
        return nil
    }
    res["reactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReactionRollupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactions(val.(ReactionRollupable))
        }
        return nil
    }
    res["side"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestReviewComment_side)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSide(val.(*PullRequestReviewComment_side))
        }
        return nil
    }
    res["start_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartLine(val)
        }
        return nil
    }
    res["start_side"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestReviewComment_start_side)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartSide(val.(*PullRequestReviewComment_start_side))
        }
        return nil
    }
    res["subject_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePullRequestReviewComment_subject_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectType(val.(*PullRequestReviewComment_subject_type))
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
// GetHtmlUrl gets the html_url property value. HTML URL for the pull request review comment.
// returns a *string when successful
func (m *PullRequestReviewComment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The ID of the pull request review comment.
// returns a *int64 when successful
func (m *PullRequestReviewComment) GetId()(*int64) {
    return m.id
}
// GetInReplyToId gets the in_reply_to_id property value. The comment ID to reply to.
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetInReplyToId()(*int32) {
    return m.in_reply_to_id
}
// GetLine gets the line property value. The line of the blob to which the comment applies. The last line of the range for a multi-line comment
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetLine()(*int32) {
    return m.line
}
// GetLinks gets the _links property value. The _links property
// returns a PullRequestReviewComment__linksable when successful
func (m *PullRequestReviewComment) GetLinks()(PullRequestReviewComment__linksable) {
    return m._links
}
// GetNodeId gets the node_id property value. The node ID of the pull request review comment.
// returns a *string when successful
func (m *PullRequestReviewComment) GetNodeId()(*string) {
    return m.node_id
}
// GetOriginalCommitId gets the original_commit_id property value. The SHA of the original commit to which the comment applies.
// returns a *string when successful
func (m *PullRequestReviewComment) GetOriginalCommitId()(*string) {
    return m.original_commit_id
}
// GetOriginalLine gets the original_line property value. The line of the blob to which the comment applies. The last line of the range for a multi-line comment
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetOriginalLine()(*int32) {
    return m.original_line
}
// GetOriginalPosition gets the original_position property value. The index of the original line in the diff to which the comment applies. This field is deprecated; use `original_line` instead.
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetOriginalPosition()(*int32) {
    return m.original_position
}
// GetOriginalStartLine gets the original_start_line property value. The first line of the range for a multi-line comment.
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetOriginalStartLine()(*int32) {
    return m.original_start_line
}
// GetPath gets the path property value. The relative path of the file to which the comment applies.
// returns a *string when successful
func (m *PullRequestReviewComment) GetPath()(*string) {
    return m.path
}
// GetPosition gets the position property value. The line index in the diff to which the comment applies. This field is deprecated; use `line` instead.
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetPosition()(*int32) {
    return m.position
}
// GetPullRequestReviewId gets the pull_request_review_id property value. The ID of the pull request review to which the comment belongs.
// returns a *int64 when successful
func (m *PullRequestReviewComment) GetPullRequestReviewId()(*int64) {
    return m.pull_request_review_id
}
// GetPullRequestUrl gets the pull_request_url property value. URL for the pull request that the review comment belongs to.
// returns a *string when successful
func (m *PullRequestReviewComment) GetPullRequestUrl()(*string) {
    return m.pull_request_url
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *PullRequestReviewComment) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetSide gets the side property value. The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
// returns a *PullRequestReviewComment_side when successful
func (m *PullRequestReviewComment) GetSide()(*PullRequestReviewComment_side) {
    return m.side
}
// GetStartLine gets the start_line property value. The first line of the range for a multi-line comment.
// returns a *int32 when successful
func (m *PullRequestReviewComment) GetStartLine()(*int32) {
    return m.start_line
}
// GetStartSide gets the start_side property value. The side of the first line of the range for a multi-line comment.
// returns a *PullRequestReviewComment_start_side when successful
func (m *PullRequestReviewComment) GetStartSide()(*PullRequestReviewComment_start_side) {
    return m.start_side
}
// GetSubjectType gets the subject_type property value. The level at which the comment is targeted, can be a diff line or a file.
// returns a *PullRequestReviewComment_subject_type when successful
func (m *PullRequestReviewComment) GetSubjectType()(*PullRequestReviewComment_subject_type) {
    return m.subject_type
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *PullRequestReviewComment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. URL for the pull request review comment
// returns a *string when successful
func (m *PullRequestReviewComment) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *PullRequestReviewComment) GetUser()(SimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *PullRequestReviewComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAuthorAssociation() != nil {
        cast := (*m.GetAuthorAssociation()).String()
        err := writer.WriteStringValue("author_association", &cast)
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
        err := writer.WriteStringValue("body_html", m.GetBodyHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body_text", m.GetBodyText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_id", m.GetCommitId())
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
        err := writer.WriteStringValue("diff_hunk", m.GetDiffHunk())
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
        err := writer.WriteInt32Value("in_reply_to_id", m.GetInReplyToId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("line", m.GetLine())
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
        err := writer.WriteStringValue("original_commit_id", m.GetOriginalCommitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("original_line", m.GetOriginalLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("original_position", m.GetOriginalPosition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("original_start_line", m.GetOriginalStartLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("pull_request_review_id", m.GetPullRequestReviewId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pull_request_url", m.GetPullRequestUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reactions", m.GetReactions())
        if err != nil {
            return err
        }
    }
    if m.GetSide() != nil {
        cast := (*m.GetSide()).String()
        err := writer.WriteStringValue("side", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("start_line", m.GetStartLine())
        if err != nil {
            return err
        }
    }
    if m.GetStartSide() != nil {
        cast := (*m.GetStartSide()).String()
        err := writer.WriteStringValue("start_side", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectType() != nil {
        cast := (*m.GetSubjectType()).String()
        err := writer.WriteStringValue("subject_type", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullRequestReviewComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *PullRequestReviewComment) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. The text of the comment.
func (m *PullRequestReviewComment) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *PullRequestReviewComment) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *PullRequestReviewComment) SetBodyText(value *string)() {
    m.body_text = value
}
// SetCommitId sets the commit_id property value. The SHA of the commit to which the comment applies.
func (m *PullRequestReviewComment) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *PullRequestReviewComment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDiffHunk sets the diff_hunk property value. The diff of the line that the comment refers to.
func (m *PullRequestReviewComment) SetDiffHunk(value *string)() {
    m.diff_hunk = value
}
// SetHtmlUrl sets the html_url property value. HTML URL for the pull request review comment.
func (m *PullRequestReviewComment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The ID of the pull request review comment.
func (m *PullRequestReviewComment) SetId(value *int64)() {
    m.id = value
}
// SetInReplyToId sets the in_reply_to_id property value. The comment ID to reply to.
func (m *PullRequestReviewComment) SetInReplyToId(value *int32)() {
    m.in_reply_to_id = value
}
// SetLine sets the line property value. The line of the blob to which the comment applies. The last line of the range for a multi-line comment
func (m *PullRequestReviewComment) SetLine(value *int32)() {
    m.line = value
}
// SetLinks sets the _links property value. The _links property
func (m *PullRequestReviewComment) SetLinks(value PullRequestReviewComment__linksable)() {
    m._links = value
}
// SetNodeId sets the node_id property value. The node ID of the pull request review comment.
func (m *PullRequestReviewComment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOriginalCommitId sets the original_commit_id property value. The SHA of the original commit to which the comment applies.
func (m *PullRequestReviewComment) SetOriginalCommitId(value *string)() {
    m.original_commit_id = value
}
// SetOriginalLine sets the original_line property value. The line of the blob to which the comment applies. The last line of the range for a multi-line comment
func (m *PullRequestReviewComment) SetOriginalLine(value *int32)() {
    m.original_line = value
}
// SetOriginalPosition sets the original_position property value. The index of the original line in the diff to which the comment applies. This field is deprecated; use `original_line` instead.
func (m *PullRequestReviewComment) SetOriginalPosition(value *int32)() {
    m.original_position = value
}
// SetOriginalStartLine sets the original_start_line property value. The first line of the range for a multi-line comment.
func (m *PullRequestReviewComment) SetOriginalStartLine(value *int32)() {
    m.original_start_line = value
}
// SetPath sets the path property value. The relative path of the file to which the comment applies.
func (m *PullRequestReviewComment) SetPath(value *string)() {
    m.path = value
}
// SetPosition sets the position property value. The line index in the diff to which the comment applies. This field is deprecated; use `line` instead.
func (m *PullRequestReviewComment) SetPosition(value *int32)() {
    m.position = value
}
// SetPullRequestReviewId sets the pull_request_review_id property value. The ID of the pull request review to which the comment belongs.
func (m *PullRequestReviewComment) SetPullRequestReviewId(value *int64)() {
    m.pull_request_review_id = value
}
// SetPullRequestUrl sets the pull_request_url property value. URL for the pull request that the review comment belongs to.
func (m *PullRequestReviewComment) SetPullRequestUrl(value *string)() {
    m.pull_request_url = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *PullRequestReviewComment) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetSide sets the side property value. The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
func (m *PullRequestReviewComment) SetSide(value *PullRequestReviewComment_side)() {
    m.side = value
}
// SetStartLine sets the start_line property value. The first line of the range for a multi-line comment.
func (m *PullRequestReviewComment) SetStartLine(value *int32)() {
    m.start_line = value
}
// SetStartSide sets the start_side property value. The side of the first line of the range for a multi-line comment.
func (m *PullRequestReviewComment) SetStartSide(value *PullRequestReviewComment_start_side)() {
    m.start_side = value
}
// SetSubjectType sets the subject_type property value. The level at which the comment is targeted, can be a diff line or a file.
func (m *PullRequestReviewComment) SetSubjectType(value *PullRequestReviewComment_subject_type)() {
    m.subject_type = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *PullRequestReviewComment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. URL for the pull request review comment
func (m *PullRequestReviewComment) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *PullRequestReviewComment) SetUser(value SimpleUserable)() {
    m.user = value
}
type PullRequestReviewCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorAssociation()(*AuthorAssociation)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCommitId()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiffHunk()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetInReplyToId()(*int32)
    GetLine()(*int32)
    GetLinks()(PullRequestReviewComment__linksable)
    GetNodeId()(*string)
    GetOriginalCommitId()(*string)
    GetOriginalLine()(*int32)
    GetOriginalPosition()(*int32)
    GetOriginalStartLine()(*int32)
    GetPath()(*string)
    GetPosition()(*int32)
    GetPullRequestReviewId()(*int64)
    GetPullRequestUrl()(*string)
    GetReactions()(ReactionRollupable)
    GetSide()(*PullRequestReviewComment_side)
    GetStartLine()(*int32)
    GetStartSide()(*PullRequestReviewComment_start_side)
    GetSubjectType()(*PullRequestReviewComment_subject_type)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(SimpleUserable)
    SetAuthorAssociation(value *AuthorAssociation)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCommitId(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiffHunk(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetInReplyToId(value *int32)()
    SetLine(value *int32)()
    SetLinks(value PullRequestReviewComment__linksable)()
    SetNodeId(value *string)()
    SetOriginalCommitId(value *string)()
    SetOriginalLine(value *int32)()
    SetOriginalPosition(value *int32)()
    SetOriginalStartLine(value *int32)()
    SetPath(value *string)()
    SetPosition(value *int32)()
    SetPullRequestReviewId(value *int64)()
    SetPullRequestUrl(value *string)()
    SetReactions(value ReactionRollupable)()
    SetSide(value *PullRequestReviewComment_side)()
    SetStartLine(value *int32)()
    SetStartSide(value *PullRequestReviewComment_start_side)()
    SetSubjectType(value *PullRequestReviewComment_subject_type)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value SimpleUserable)()
}
