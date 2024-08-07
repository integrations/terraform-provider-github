package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Milestone a collection of related issues and pull requests.
type Milestone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The closed_at property
    closed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The closed_issues property
    closed_issues *int32
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    creator NullableSimpleUserable
    // The description property
    description *string
    // The due_on property
    due_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The labels_url property
    labels_url *string
    // The node_id property
    node_id *string
    // The number of the milestone.
    number *int32
    // The open_issues property
    open_issues *int32
    // The state of the milestone.
    state *Milestone_state
    // The title of the milestone.
    title *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewMilestone instantiates a new Milestone and sets the default values.
func NewMilestone()(*Milestone) {
    m := &Milestone{
    }
    m.SetAdditionalData(make(map[string]any))
    stateValue := OPEN_MILESTONE_STATE
    m.SetState(&stateValue)
    return m
}
// CreateMilestoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMilestoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMilestone(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Milestone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClosedAt gets the closed_at property value. The closed_at property
// returns a *Time when successful
func (m *Milestone) GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closed_at
}
// GetClosedIssues gets the closed_issues property value. The closed_issues property
// returns a *int32 when successful
func (m *Milestone) GetClosedIssues()(*int32) {
    return m.closed_issues
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Milestone) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Milestone) GetCreator()(NullableSimpleUserable) {
    return m.creator
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Milestone) GetDescription()(*string) {
    return m.description
}
// GetDueOn gets the due_on property value. The due_on property
// returns a *Time when successful
func (m *Milestone) GetDueOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.due_on
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Milestone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["closed_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedIssues(val)
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
    res["creator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["due_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueOn(val)
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
    res["labels_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelsUrl(val)
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
    res["open_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenIssues(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMilestone_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*Milestone_state))
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
func (m *Milestone) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Milestone) GetId()(*int32) {
    return m.id
}
// GetLabelsUrl gets the labels_url property value. The labels_url property
// returns a *string when successful
func (m *Milestone) GetLabelsUrl()(*string) {
    return m.labels_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Milestone) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The number of the milestone.
// returns a *int32 when successful
func (m *Milestone) GetNumber()(*int32) {
    return m.number
}
// GetOpenIssues gets the open_issues property value. The open_issues property
// returns a *int32 when successful
func (m *Milestone) GetOpenIssues()(*int32) {
    return m.open_issues
}
// GetState gets the state property value. The state of the milestone.
// returns a *Milestone_state when successful
func (m *Milestone) GetState()(*Milestone_state) {
    return m.state
}
// GetTitle gets the title property value. The title of the milestone.
// returns a *string when successful
func (m *Milestone) GetTitle()(*string) {
    return m.title
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Milestone) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Milestone) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Milestone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("closed_at", m.GetClosedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("closed_issues", m.GetClosedIssues())
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
        err := writer.WriteObjectValue("creator", m.GetCreator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("due_on", m.GetDueOn())
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
        err := writer.WriteStringValue("labels_url", m.GetLabelsUrl())
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
        err := writer.WriteInt32Value("open_issues", m.GetOpenIssues())
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
func (m *Milestone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClosedAt sets the closed_at property value. The closed_at property
func (m *Milestone) SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closed_at = value
}
// SetClosedIssues sets the closed_issues property value. The closed_issues property
func (m *Milestone) SetClosedIssues(value *int32)() {
    m.closed_issues = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Milestone) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *Milestone) SetCreator(value NullableSimpleUserable)() {
    m.creator = value
}
// SetDescription sets the description property value. The description property
func (m *Milestone) SetDescription(value *string)() {
    m.description = value
}
// SetDueOn sets the due_on property value. The due_on property
func (m *Milestone) SetDueOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.due_on = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Milestone) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Milestone) SetId(value *int32)() {
    m.id = value
}
// SetLabelsUrl sets the labels_url property value. The labels_url property
func (m *Milestone) SetLabelsUrl(value *string)() {
    m.labels_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Milestone) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The number of the milestone.
func (m *Milestone) SetNumber(value *int32)() {
    m.number = value
}
// SetOpenIssues sets the open_issues property value. The open_issues property
func (m *Milestone) SetOpenIssues(value *int32)() {
    m.open_issues = value
}
// SetState sets the state property value. The state of the milestone.
func (m *Milestone) SetState(value *Milestone_state)() {
    m.state = value
}
// SetTitle sets the title property value. The title of the milestone.
func (m *Milestone) SetTitle(value *string)() {
    m.title = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Milestone) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Milestone) SetUrl(value *string)() {
    m.url = value
}
type Milestoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetClosedIssues()(*int32)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetDescription()(*string)
    GetDueOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLabelsUrl()(*string)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetOpenIssues()(*int32)
    GetState()(*Milestone_state)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetClosedIssues(value *int32)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetDescription(value *string)()
    SetDueOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLabelsUrl(value *string)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetOpenIssues(value *int32)()
    SetState(value *Milestone_state)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
