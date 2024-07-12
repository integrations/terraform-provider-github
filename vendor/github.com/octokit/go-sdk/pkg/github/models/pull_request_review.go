package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestReview pull Request Reviews are reviews on pull requests.
type PullRequestReview struct {
    // The _links property
    _links PullRequestReview__linksable
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
    // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
    commit_id *string
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
    user NullableSimpleUserable
}
// NewPullRequestReview instantiates a new PullRequestReview and sets the default values.
func NewPullRequestReview()(*PullRequestReview) {
    m := &PullRequestReview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestReview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestReview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *PullRequestReview) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. The text of the review.
// returns a *string when successful
func (m *PullRequestReview) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *PullRequestReview) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *PullRequestReview) GetBodyText()(*string) {
    return m.body_text
}
// GetCommitId gets the commit_id property value. A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
// returns a *string when successful
func (m *PullRequestReview) GetCommitId()(*string) {
    return m.commit_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestReview__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(PullRequestReview__linksable))
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
func (m *PullRequestReview) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the review
// returns a *int32 when successful
func (m *PullRequestReview) GetId()(*int32) {
    return m.id
}
// GetLinks gets the _links property value. The _links property
// returns a PullRequestReview__linksable when successful
func (m *PullRequestReview) GetLinks()(PullRequestReview__linksable) {
    return m._links
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *PullRequestReview) GetNodeId()(*string) {
    return m.node_id
}
// GetPullRequestUrl gets the pull_request_url property value. The pull_request_url property
// returns a *string when successful
func (m *PullRequestReview) GetPullRequestUrl()(*string) {
    return m.pull_request_url
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *PullRequestReview) GetState()(*string) {
    return m.state
}
// GetSubmittedAt gets the submitted_at property value. The submitted_at property
// returns a *Time when successful
func (m *PullRequestReview) GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.submitted_at
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *PullRequestReview) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *PullRequestReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PullRequestReview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *PullRequestReview) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. The text of the review.
func (m *PullRequestReview) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *PullRequestReview) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *PullRequestReview) SetBodyText(value *string)() {
    m.body_text = value
}
// SetCommitId sets the commit_id property value. A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
func (m *PullRequestReview) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *PullRequestReview) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the review
func (m *PullRequestReview) SetId(value *int32)() {
    m.id = value
}
// SetLinks sets the _links property value. The _links property
func (m *PullRequestReview) SetLinks(value PullRequestReview__linksable)() {
    m._links = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *PullRequestReview) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPullRequestUrl sets the pull_request_url property value. The pull_request_url property
func (m *PullRequestReview) SetPullRequestUrl(value *string)() {
    m.pull_request_url = value
}
// SetState sets the state property value. The state property
func (m *PullRequestReview) SetState(value *string)() {
    m.state = value
}
// SetSubmittedAt sets the submitted_at property value. The submitted_at property
func (m *PullRequestReview) SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submitted_at = value
}
// SetUser sets the user property value. A GitHub user.
func (m *PullRequestReview) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type PullRequestReviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorAssociation()(*AuthorAssociation)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCommitId()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLinks()(PullRequestReview__linksable)
    GetNodeId()(*string)
    GetPullRequestUrl()(*string)
    GetState()(*string)
    GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUser()(NullableSimpleUserable)
    SetAuthorAssociation(value *AuthorAssociation)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCommitId(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLinks(value PullRequestReview__linksable)()
    SetNodeId(value *string)()
    SetPullRequestUrl(value *string)()
    SetState(value *string)()
    SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUser(value NullableSimpleUserable)()
}
