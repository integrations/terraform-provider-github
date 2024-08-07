package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReviewComment legacy Review Comment
type ReviewComment struct {
    // The _links property
    _links ReviewComment__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The body property
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // The commit_id property
    commit_id *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The diff_hunk property
    diff_hunk *string
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The in_reply_to_id property
    in_reply_to_id *int32
    // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
    line *int32
    // The node_id property
    node_id *string
    // The original_commit_id property
    original_commit_id *string
    // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
    original_line *int32
    // The original_position property
    original_position *int32
    // The original first line of the range for a multi-line comment.
    original_start_line *int32
    // The path property
    path *string
    // The position property
    position *int32
    // The pull_request_review_id property
    pull_request_review_id *int64
    // The pull_request_url property
    pull_request_url *string
    // The reactions property
    reactions ReactionRollupable
    // The side of the first line of the range for a multi-line comment.
    side *ReviewComment_side
    // The first line of the range for a multi-line comment.
    start_line *int32
    // The side of the first line of the range for a multi-line comment.
    start_side *ReviewComment_start_side
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewReviewComment instantiates a new ReviewComment and sets the default values.
func NewReviewComment()(*ReviewComment) {
    m := &ReviewComment{
    }
    m.SetAdditionalData(make(map[string]any))
    sideValue := RIGHT_REVIEWCOMMENT_SIDE
    m.SetSide(&sideValue)
    start_sideValue := RIGHT_REVIEWCOMMENT_START_SIDE
    m.SetStartSide(&start_sideValue)
    return m
}
// CreateReviewCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReviewCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReviewComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *ReviewComment) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *ReviewComment) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *ReviewComment) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *ReviewComment) GetBodyText()(*string) {
    return m.body_text
}
// GetCommitId gets the commit_id property value. The commit_id property
// returns a *string when successful
func (m *ReviewComment) GetCommitId()(*string) {
    return m.commit_id
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *ReviewComment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDiffHunk gets the diff_hunk property value. The diff_hunk property
// returns a *string when successful
func (m *ReviewComment) GetDiffHunk()(*string) {
    return m.diff_hunk
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReviewComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReviewComment__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(ReviewComment__linksable))
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
        val, err := n.GetEnumValue(ParseReviewComment_side)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSide(val.(*ReviewComment_side))
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
        val, err := n.GetEnumValue(ParseReviewComment_start_side)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartSide(val.(*ReviewComment_start_side))
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *ReviewComment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ReviewComment) GetId()(*int64) {
    return m.id
}
// GetInReplyToId gets the in_reply_to_id property value. The in_reply_to_id property
// returns a *int32 when successful
func (m *ReviewComment) GetInReplyToId()(*int32) {
    return m.in_reply_to_id
}
// GetLine gets the line property value. The line of the blob to which the comment applies. The last line of the range for a multi-line comment
// returns a *int32 when successful
func (m *ReviewComment) GetLine()(*int32) {
    return m.line
}
// GetLinks gets the _links property value. The _links property
// returns a ReviewComment__linksable when successful
func (m *ReviewComment) GetLinks()(ReviewComment__linksable) {
    return m._links
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *ReviewComment) GetNodeId()(*string) {
    return m.node_id
}
// GetOriginalCommitId gets the original_commit_id property value. The original_commit_id property
// returns a *string when successful
func (m *ReviewComment) GetOriginalCommitId()(*string) {
    return m.original_commit_id
}
// GetOriginalLine gets the original_line property value. The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
// returns a *int32 when successful
func (m *ReviewComment) GetOriginalLine()(*int32) {
    return m.original_line
}
// GetOriginalPosition gets the original_position property value. The original_position property
// returns a *int32 when successful
func (m *ReviewComment) GetOriginalPosition()(*int32) {
    return m.original_position
}
// GetOriginalStartLine gets the original_start_line property value. The original first line of the range for a multi-line comment.
// returns a *int32 when successful
func (m *ReviewComment) GetOriginalStartLine()(*int32) {
    return m.original_start_line
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ReviewComment) GetPath()(*string) {
    return m.path
}
// GetPosition gets the position property value. The position property
// returns a *int32 when successful
func (m *ReviewComment) GetPosition()(*int32) {
    return m.position
}
// GetPullRequestReviewId gets the pull_request_review_id property value. The pull_request_review_id property
// returns a *int64 when successful
func (m *ReviewComment) GetPullRequestReviewId()(*int64) {
    return m.pull_request_review_id
}
// GetPullRequestUrl gets the pull_request_url property value. The pull_request_url property
// returns a *string when successful
func (m *ReviewComment) GetPullRequestUrl()(*string) {
    return m.pull_request_url
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *ReviewComment) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetSide gets the side property value. The side of the first line of the range for a multi-line comment.
// returns a *ReviewComment_side when successful
func (m *ReviewComment) GetSide()(*ReviewComment_side) {
    return m.side
}
// GetStartLine gets the start_line property value. The first line of the range for a multi-line comment.
// returns a *int32 when successful
func (m *ReviewComment) GetStartLine()(*int32) {
    return m.start_line
}
// GetStartSide gets the start_side property value. The side of the first line of the range for a multi-line comment.
// returns a *ReviewComment_start_side when successful
func (m *ReviewComment) GetStartSide()(*ReviewComment_start_side) {
    return m.start_side
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *ReviewComment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ReviewComment) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *ReviewComment) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *ReviewComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReviewComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *ReviewComment) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. The body property
func (m *ReviewComment) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *ReviewComment) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *ReviewComment) SetBodyText(value *string)() {
    m.body_text = value
}
// SetCommitId sets the commit_id property value. The commit_id property
func (m *ReviewComment) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ReviewComment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDiffHunk sets the diff_hunk property value. The diff_hunk property
func (m *ReviewComment) SetDiffHunk(value *string)() {
    m.diff_hunk = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *ReviewComment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *ReviewComment) SetId(value *int64)() {
    m.id = value
}
// SetInReplyToId sets the in_reply_to_id property value. The in_reply_to_id property
func (m *ReviewComment) SetInReplyToId(value *int32)() {
    m.in_reply_to_id = value
}
// SetLine sets the line property value. The line of the blob to which the comment applies. The last line of the range for a multi-line comment
func (m *ReviewComment) SetLine(value *int32)() {
    m.line = value
}
// SetLinks sets the _links property value. The _links property
func (m *ReviewComment) SetLinks(value ReviewComment__linksable)() {
    m._links = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *ReviewComment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOriginalCommitId sets the original_commit_id property value. The original_commit_id property
func (m *ReviewComment) SetOriginalCommitId(value *string)() {
    m.original_commit_id = value
}
// SetOriginalLine sets the original_line property value. The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
func (m *ReviewComment) SetOriginalLine(value *int32)() {
    m.original_line = value
}
// SetOriginalPosition sets the original_position property value. The original_position property
func (m *ReviewComment) SetOriginalPosition(value *int32)() {
    m.original_position = value
}
// SetOriginalStartLine sets the original_start_line property value. The original first line of the range for a multi-line comment.
func (m *ReviewComment) SetOriginalStartLine(value *int32)() {
    m.original_start_line = value
}
// SetPath sets the path property value. The path property
func (m *ReviewComment) SetPath(value *string)() {
    m.path = value
}
// SetPosition sets the position property value. The position property
func (m *ReviewComment) SetPosition(value *int32)() {
    m.position = value
}
// SetPullRequestReviewId sets the pull_request_review_id property value. The pull_request_review_id property
func (m *ReviewComment) SetPullRequestReviewId(value *int64)() {
    m.pull_request_review_id = value
}
// SetPullRequestUrl sets the pull_request_url property value. The pull_request_url property
func (m *ReviewComment) SetPullRequestUrl(value *string)() {
    m.pull_request_url = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *ReviewComment) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetSide sets the side property value. The side of the first line of the range for a multi-line comment.
func (m *ReviewComment) SetSide(value *ReviewComment_side)() {
    m.side = value
}
// SetStartLine sets the start_line property value. The first line of the range for a multi-line comment.
func (m *ReviewComment) SetStartLine(value *int32)() {
    m.start_line = value
}
// SetStartSide sets the start_side property value. The side of the first line of the range for a multi-line comment.
func (m *ReviewComment) SetStartSide(value *ReviewComment_start_side)() {
    m.start_side = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *ReviewComment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *ReviewComment) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *ReviewComment) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type ReviewCommentable interface {
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
    GetLinks()(ReviewComment__linksable)
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
    GetSide()(*ReviewComment_side)
    GetStartLine()(*int32)
    GetStartSide()(*ReviewComment_start_side)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
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
    SetLinks(value ReviewComment__linksable)()
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
    SetSide(value *ReviewComment_side)()
    SetStartLine(value *int32)()
    SetStartSide(value *ReviewComment_start_side)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
