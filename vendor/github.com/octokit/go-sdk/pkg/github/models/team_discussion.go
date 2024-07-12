package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamDiscussion a team discussion is a persistent record of a free-form conversation within a team.
type TeamDiscussion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    author NullableSimpleUserable
    // The main text of the discussion.
    body *string
    // The body_html property
    body_html *string
    // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
    body_version *string
    // The comments_count property
    comments_count *int32
    // The comments_url property
    comments_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The html_url property
    html_url *string
    // The last_edited_at property
    last_edited_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The node_id property
    node_id *string
    // The unique sequence number of a team discussion.
    number *int32
    // Whether or not this discussion should be pinned for easy retrieval.
    pinned *bool
    // Whether or not this discussion should be restricted to team members and organization owners.
    private *bool
    // The reactions property
    reactions ReactionRollupable
    // The team_url property
    team_url *string
    // The title of the discussion.
    title *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewTeamDiscussion instantiates a new TeamDiscussion and sets the default values.
func NewTeamDiscussion()(*TeamDiscussion) {
    m := &TeamDiscussion{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamDiscussionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamDiscussionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamDiscussion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamDiscussion) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *TeamDiscussion) GetAuthor()(NullableSimpleUserable) {
    return m.author
}
// GetBody gets the body property value. The main text of the discussion.
// returns a *string when successful
func (m *TeamDiscussion) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *TeamDiscussion) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyVersion gets the body_version property value. The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
// returns a *string when successful
func (m *TeamDiscussion) GetBodyVersion()(*string) {
    return m.body_version
}
// GetCommentsCount gets the comments_count property value. The comments_count property
// returns a *int32 when successful
func (m *TeamDiscussion) GetCommentsCount()(*int32) {
    return m.comments_count
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *TeamDiscussion) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *TeamDiscussion) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamDiscussion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["comments_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsCount(val)
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
    res["pinned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinned(val)
        }
        return nil
    }
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
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
    res["team_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamUrl(val)
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
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *TeamDiscussion) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLastEditedAt gets the last_edited_at property value. The last_edited_at property
// returns a *Time when successful
func (m *TeamDiscussion) GetLastEditedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_edited_at
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TeamDiscussion) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The unique sequence number of a team discussion.
// returns a *int32 when successful
func (m *TeamDiscussion) GetNumber()(*int32) {
    return m.number
}
// GetPinned gets the pinned property value. Whether or not this discussion should be pinned for easy retrieval.
// returns a *bool when successful
func (m *TeamDiscussion) GetPinned()(*bool) {
    return m.pinned
}
// GetPrivate gets the private property value. Whether or not this discussion should be restricted to team members and organization owners.
// returns a *bool when successful
func (m *TeamDiscussion) GetPrivate()(*bool) {
    return m.private
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *TeamDiscussion) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetTeamUrl gets the team_url property value. The team_url property
// returns a *string when successful
func (m *TeamDiscussion) GetTeamUrl()(*string) {
    return m.team_url
}
// GetTitle gets the title property value. The title of the discussion.
// returns a *string when successful
func (m *TeamDiscussion) GetTitle()(*string) {
    return m.title
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *TeamDiscussion) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TeamDiscussion) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TeamDiscussion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("comments_count", m.GetCommentsCount())
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
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteBoolValue("pinned", m.GetPinned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("private", m.GetPrivate())
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
        err := writer.WriteStringValue("team_url", m.GetTeamUrl())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamDiscussion) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. A GitHub user.
func (m *TeamDiscussion) SetAuthor(value NullableSimpleUserable)() {
    m.author = value
}
// SetBody sets the body property value. The main text of the discussion.
func (m *TeamDiscussion) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *TeamDiscussion) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyVersion sets the body_version property value. The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
func (m *TeamDiscussion) SetBodyVersion(value *string)() {
    m.body_version = value
}
// SetCommentsCount sets the comments_count property value. The comments_count property
func (m *TeamDiscussion) SetCommentsCount(value *int32)() {
    m.comments_count = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *TeamDiscussion) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *TeamDiscussion) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TeamDiscussion) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLastEditedAt sets the last_edited_at property value. The last_edited_at property
func (m *TeamDiscussion) SetLastEditedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_edited_at = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TeamDiscussion) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The unique sequence number of a team discussion.
func (m *TeamDiscussion) SetNumber(value *int32)() {
    m.number = value
}
// SetPinned sets the pinned property value. Whether or not this discussion should be pinned for easy retrieval.
func (m *TeamDiscussion) SetPinned(value *bool)() {
    m.pinned = value
}
// SetPrivate sets the private property value. Whether or not this discussion should be restricted to team members and organization owners.
func (m *TeamDiscussion) SetPrivate(value *bool)() {
    m.private = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *TeamDiscussion) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetTeamUrl sets the team_url property value. The team_url property
func (m *TeamDiscussion) SetTeamUrl(value *string)() {
    m.team_url = value
}
// SetTitle sets the title property value. The title of the discussion.
func (m *TeamDiscussion) SetTitle(value *string)() {
    m.title = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *TeamDiscussion) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *TeamDiscussion) SetUrl(value *string)() {
    m.url = value
}
type TeamDiscussionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyVersion()(*string)
    GetCommentsCount()(*int32)
    GetCommentsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetLastEditedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetPinned()(*bool)
    GetPrivate()(*bool)
    GetReactions()(ReactionRollupable)
    GetTeamUrl()(*string)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyVersion(value *string)()
    SetCommentsCount(value *int32)()
    SetCommentsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetLastEditedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetPinned(value *bool)()
    SetPrivate(value *bool)()
    SetReactions(value ReactionRollupable)()
    SetTeamUrl(value *string)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
