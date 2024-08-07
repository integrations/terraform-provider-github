package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamDiscussionComment a reply to a discussion within a team.
type TeamDiscussionComment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    author NullableSimpleUserable
    // The main text of the comment.
    body *string
    // The body_html property
    body_html *string
    // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
    body_version *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The discussion_url property
    discussion_url *string
    // The html_url property
    html_url *string
    // The last_edited_at property
    last_edited_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The node_id property
    node_id *string
    // The unique sequence number of a team discussion comment.
    number *int32
    // The reactions property
    reactions ReactionRollupable
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewTeamDiscussionComment instantiates a new TeamDiscussionComment and sets the default values.
func NewTeamDiscussionComment()(*TeamDiscussionComment) {
    m := &TeamDiscussionComment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamDiscussionCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamDiscussionCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamDiscussionComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamDiscussionComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *TeamDiscussionComment) GetAuthor()(NullableSimpleUserable) {
    return m.author
}
// GetBody gets the body property value. The main text of the comment.
// returns a *string when successful
func (m *TeamDiscussionComment) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *TeamDiscussionComment) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyVersion gets the body_version property value. The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
// returns a *string when successful
func (m *TeamDiscussionComment) GetBodyVersion()(*string) {
    return m.body_version
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *TeamDiscussionComment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDiscussionUrl gets the discussion_url property value. The discussion_url property
// returns a *string when successful
func (m *TeamDiscussionComment) GetDiscussionUrl()(*string) {
    return m.discussion_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamDiscussionComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(NullableSimpleUserable))
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
    res["body_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyVersion(val)
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
    res["discussion_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscussionUrl(val)
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
    res["last_edited_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEditedAt(val)
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
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *TeamDiscussionComment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLastEditedAt gets the last_edited_at property value. The last_edited_at property
// returns a *Time when successful
func (m *TeamDiscussionComment) GetLastEditedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_edited_at
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TeamDiscussionComment) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The unique sequence number of a team discussion comment.
// returns a *int32 when successful
func (m *TeamDiscussionComment) GetNumber()(*int32) {
    return m.number
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *TeamDiscussionComment) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *TeamDiscussionComment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TeamDiscussionComment) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TeamDiscussionComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
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
        err := writer.WriteStringValue("body_version", m.GetBodyVersion())
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
        err := writer.WriteStringValue("discussion_url", m.GetDiscussionUrl())
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
        err := writer.WriteTimeValue("last_edited_at", m.GetLastEditedAt())
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
        err := writer.WriteObjectValue("reactions", m.GetReactions())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamDiscussionComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. A GitHub user.
func (m *TeamDiscussionComment) SetAuthor(value NullableSimpleUserable)() {
    m.author = value
}
// SetBody sets the body property value. The main text of the comment.
func (m *TeamDiscussionComment) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *TeamDiscussionComment) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyVersion sets the body_version property value. The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
func (m *TeamDiscussionComment) SetBodyVersion(value *string)() {
    m.body_version = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *TeamDiscussionComment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDiscussionUrl sets the discussion_url property value. The discussion_url property
func (m *TeamDiscussionComment) SetDiscussionUrl(value *string)() {
    m.discussion_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TeamDiscussionComment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLastEditedAt sets the last_edited_at property value. The last_edited_at property
func (m *TeamDiscussionComment) SetLastEditedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_edited_at = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TeamDiscussionComment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The unique sequence number of a team discussion comment.
func (m *TeamDiscussionComment) SetNumber(value *int32)() {
    m.number = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *TeamDiscussionComment) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *TeamDiscussionComment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *TeamDiscussionComment) SetUrl(value *string)() {
    m.url = value
}
type TeamDiscussionCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyVersion()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiscussionUrl()(*string)
    GetHtmlUrl()(*string)
    GetLastEditedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetReactions()(ReactionRollupable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyVersion(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiscussionUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLastEditedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetReactions(value ReactionRollupable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
