package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueComment comments provide a way for people to collaborate on an issue.
type IssueComment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // Contents of the issue comment
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The html_url property
    html_url *string
    // Unique identifier of the issue comment
    id *int64
    // The issue_url property
    issue_url *string
    // The node_id property
    node_id *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // The reactions property
    reactions ReactionRollupable
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // URL for the issue comment
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewIssueComment instantiates a new IssueComment and sets the default values.
func NewIssueComment()(*IssueComment) {
    m := &IssueComment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssueCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssueComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IssueComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *IssueComment) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. Contents of the issue comment
// returns a *string when successful
func (m *IssueComment) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *IssueComment) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *IssueComment) GetBodyText()(*string) {
    return m.body_text
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *IssueComment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["performed_via_github_app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedViaGithubApp(val.(NullableIntegrationable))
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
func (m *IssueComment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the issue comment
// returns a *int64 when successful
func (m *IssueComment) GetId()(*int64) {
    return m.id
}
// GetIssueUrl gets the issue_url property value. The issue_url property
// returns a *string when successful
func (m *IssueComment) GetIssueUrl()(*string) {
    return m.issue_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *IssueComment) GetNodeId()(*string) {
    return m.node_id
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *IssueComment) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *IssueComment) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *IssueComment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. URL for the issue comment
// returns a *string when successful
func (m *IssueComment) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueComment) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *IssueComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
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
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *IssueComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *IssueComment) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. Contents of the issue comment
func (m *IssueComment) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *IssueComment) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *IssueComment) SetBodyText(value *string)() {
    m.body_text = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *IssueComment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *IssueComment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the issue comment
func (m *IssueComment) SetId(value *int64)() {
    m.id = value
}
// SetIssueUrl sets the issue_url property value. The issue_url property
func (m *IssueComment) SetIssueUrl(value *string)() {
    m.issue_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *IssueComment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *IssueComment) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *IssueComment) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *IssueComment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. URL for the issue comment
func (m *IssueComment) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *IssueComment) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type IssueCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorAssociation()(*AuthorAssociation)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetIssueUrl()(*string)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetReactions()(ReactionRollupable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetAuthorAssociation(value *AuthorAssociation)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetIssueUrl(value *string)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetReactions(value ReactionRollupable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
