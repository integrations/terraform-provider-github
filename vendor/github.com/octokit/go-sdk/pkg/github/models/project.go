package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Project projects are a way to organize columns and cards of work.
type Project struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Body of the project
    body *string
    // The columns_url property
    columns_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    creator NullableSimpleUserable
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // Name of the project
    name *string
    // The node_id property
    node_id *string
    // The number property
    number *int32
    // The baseline permission that all organization members have on this project. Only present if owner is an organization.
    organization_permission *Project_organization_permission
    // The owner_url property
    owner_url *string
    // Whether or not this project can be seen by everyone. Only present if owner is an organization.
    private *bool
    // State of the project; either 'open' or 'closed'
    state *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewProject instantiates a new Project and sets the default values.
func NewProject()(*Project) {
    m := &Project{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProject(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Project) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. Body of the project
// returns a *string when successful
func (m *Project) GetBody()(*string) {
    return m.body
}
// GetColumnsUrl gets the columns_url property value. The columns_url property
// returns a *string when successful
func (m *Project) GetColumnsUrl()(*string) {
    return m.columns_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Project) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Project) GetCreator()(NullableSimpleUserable) {
    return m.creator
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Project) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["columns_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnsUrl(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["organization_permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProject_organization_permission)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPermission(val.(*Project_organization_permission))
        }
        return nil
    }
    res["owner_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUrl(val)
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
func (m *Project) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Project) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Name of the project
// returns a *string when successful
func (m *Project) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Project) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *Project) GetNumber()(*int32) {
    return m.number
}
// GetOrganizationPermission gets the organization_permission property value. The baseline permission that all organization members have on this project. Only present if owner is an organization.
// returns a *Project_organization_permission when successful
func (m *Project) GetOrganizationPermission()(*Project_organization_permission) {
    return m.organization_permission
}
// GetOwnerUrl gets the owner_url property value. The owner_url property
// returns a *string when successful
func (m *Project) GetOwnerUrl()(*string) {
    return m.owner_url
}
// GetPrivate gets the private property value. Whether or not this project can be seen by everyone. Only present if owner is an organization.
// returns a *bool when successful
func (m *Project) GetPrivate()(*bool) {
    return m.private
}
// GetState gets the state property value. State of the project; either 'open' or 'closed'
// returns a *string when successful
func (m *Project) GetState()(*string) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Project) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Project) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Project) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("columns_url", m.GetColumnsUrl())
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
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetOrganizationPermission() != nil {
        cast := (*m.GetOrganizationPermission()).String()
        err := writer.WriteStringValue("organization_permission", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("owner_url", m.GetOwnerUrl())
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
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *Project) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. Body of the project
func (m *Project) SetBody(value *string)() {
    m.body = value
}
// SetColumnsUrl sets the columns_url property value. The columns_url property
func (m *Project) SetColumnsUrl(value *string)() {
    m.columns_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Project) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *Project) SetCreator(value NullableSimpleUserable)() {
    m.creator = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Project) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Project) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Name of the project
func (m *Project) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Project) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The number property
func (m *Project) SetNumber(value *int32)() {
    m.number = value
}
// SetOrganizationPermission sets the organization_permission property value. The baseline permission that all organization members have on this project. Only present if owner is an organization.
func (m *Project) SetOrganizationPermission(value *Project_organization_permission)() {
    m.organization_permission = value
}
// SetOwnerUrl sets the owner_url property value. The owner_url property
func (m *Project) SetOwnerUrl(value *string)() {
    m.owner_url = value
}
// SetPrivate sets the private property value. Whether or not this project can be seen by everyone. Only present if owner is an organization.
func (m *Project) SetPrivate(value *bool)() {
    m.private = value
}
// SetState sets the state property value. State of the project; either 'open' or 'closed'
func (m *Project) SetState(value *string)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Project) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Project) SetUrl(value *string)() {
    m.url = value
}
type Projectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetColumnsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetOrganizationPermission()(*Project_organization_permission)
    GetOwnerUrl()(*string)
    GetPrivate()(*bool)
    GetState()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetBody(value *string)()
    SetColumnsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetOrganizationPermission(value *Project_organization_permission)()
    SetOwnerUrl(value *string)()
    SetPrivate(value *bool)()
    SetState(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
