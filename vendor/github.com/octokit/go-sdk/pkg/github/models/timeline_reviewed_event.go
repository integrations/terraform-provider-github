package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineReviewedEvent timeline Reviewed Event
type TimelineReviewedEvent struct {
    // The _links property
    _links TimelineReviewedEvent__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The text of the review.
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // A commit SHA for the review.
    commit_id *string
    // The event property
    event *string
    // The html_url property
    html_url *string
    // Unique identifier of the review
    id *int32
    // The node_id property
    node_id *string
    // The pull_request_url property
    pull_request_url *string
    // The state property
    state *string
    // The submitted_at property
    submitted_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    user SimpleUserable
}
// NewTimelineReviewedEvent instantiates a new TimelineReviewedEvent and sets the default values.
func NewTimelineReviewedEvent()(*TimelineReviewedEvent) {
    m := &TimelineReviewedEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineReviewedEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineReviewedEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineReviewedEvent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineReviewedEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *TimelineReviewedEvent) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. The text of the review.
// returns a *string when successful
func (m *TimelineReviewedEvent) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetBodyText()(*string) {
    return m.body_text
}
// GetCommitId gets the commit_id property value. A commit SHA for the review.
// returns a *string when successful
func (m *TimelineReviewedEvent) GetCommitId()(*string) {
    return m.commit_id
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineReviewedEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineReviewedEvent__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(TimelineReviewedEvent__linksable))
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
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
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
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["submitted_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedAt(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the review
// returns a *int32 when successful
func (m *TimelineReviewedEvent) GetId()(*int32) {
    return m.id
}
// GetLinks gets the _links property value. The _links property
// returns a TimelineReviewedEvent__linksable when successful
func (m *TimelineReviewedEvent) GetLinks()(TimelineReviewedEvent__linksable) {
    return m._links
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetNodeId()(*string) {
    return m.node_id
}
// GetPullRequestUrl gets the pull_request_url property value. The pull_request_url property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetPullRequestUrl()(*string) {
    return m.pull_request_url
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *TimelineReviewedEvent) GetState()(*string) {
    return m.state
}
// GetSubmittedAt gets the submitted_at property value. The submitted_at property
// returns a *Time when successful
func (m *TimelineReviewedEvent) GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.submitted_at
}
// GetUser gets the user property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *TimelineReviewedEvent) GetUser()(SimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *TimelineReviewedEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("event", m.GetEvent())
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
        err := writer.WriteInt32Value("id", m.GetId())
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
        err := writer.WriteStringValue("pull_request_url", m.GetPullRequestUrl())
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
        err := writer.WriteTimeValue("submitted_at", m.GetSubmittedAt())
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
func (m *TimelineReviewedEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *TimelineReviewedEvent) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. The text of the review.
func (m *TimelineReviewedEvent) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *TimelineReviewedEvent) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *TimelineReviewedEvent) SetBodyText(value *string)() {
    m.body_text = value
}
// SetCommitId sets the commit_id property value. A commit SHA for the review.
func (m *TimelineReviewedEvent) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetEvent sets the event property value. The event property
func (m *TimelineReviewedEvent) SetEvent(value *string)() {
    m.event = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TimelineReviewedEvent) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the review
func (m *TimelineReviewedEvent) SetId(value *int32)() {
    m.id = value
}
// SetLinks sets the _links property value. The _links property
func (m *TimelineReviewedEvent) SetLinks(value TimelineReviewedEvent__linksable)() {
    m._links = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TimelineReviewedEvent) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPullRequestUrl sets the pull_request_url property value. The pull_request_url property
func (m *TimelineReviewedEvent) SetPullRequestUrl(value *string)() {
    m.pull_request_url = value
}
// SetState sets the state property value. The state property
func (m *TimelineReviewedEvent) SetState(value *string)() {
    m.state = value
}
// SetSubmittedAt sets the submitted_at property value. The submitted_at property
func (m *TimelineReviewedEvent) SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submitted_at = value
}
// SetUser sets the user property value. A GitHub user.
func (m *TimelineReviewedEvent) SetUser(value SimpleUserable)() {
    m.user = value
}
type TimelineReviewedEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorAssociation()(*AuthorAssociation)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCommitId()(*string)
    GetEvent()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLinks()(TimelineReviewedEvent__linksable)
    GetNodeId()(*string)
    GetPullRequestUrl()(*string)
    GetState()(*string)
    GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUser()(SimpleUserable)
    SetAuthorAssociation(value *AuthorAssociation)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCommitId(value *string)()
    SetEvent(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLinks(value TimelineReviewedEvent__linksable)()
    SetNodeId(value *string)()
    SetPullRequestUrl(value *string)()
    SetState(value *string)()
    SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUser(value SimpleUserable)()
}
