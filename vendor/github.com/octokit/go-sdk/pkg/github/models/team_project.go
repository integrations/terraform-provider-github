package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamProject a team's access to a project.
type TeamProject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The body property
    body *string
    // The columns_url property
    columns_url *string
    // The created_at property
    created_at *string
    // A GitHub user.
    creator SimpleUserable
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The number property
    number *int32
    // The organization permission for this project. Only present when owner is an organization.
    organization_permission *string
    // The owner_url property
    owner_url *string
    // The permissions property
    permissions TeamProject_permissionsable
    // Whether the project is private or not. Only present when owner is an organization.
    private *bool
    // The state property
    state *string
    // The updated_at property
    updated_at *string
    // The url property
    url *string
}
// NewTeamProject instantiates a new TeamProject and sets the default values.
func NewTeamProject()(*TeamProject) {
    m := &TeamProject{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamProjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamProjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamProject(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamProject) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *TeamProject) GetBody()(*string) {
    return m.body
}
// GetColumnsUrl gets the columns_url property value. The columns_url property
// returns a *string when successful
func (m *TeamProject) GetColumnsUrl()(*string) {
    return m.columns_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *TeamProject) GetCreatedAt()(*string) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *TeamProject) GetCreator()(SimpleUserable) {
    return m.creator
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamProject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["creator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val.(SimpleUserable))
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPermission(val)
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
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamProject_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(TeamProject_permissionsable))
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
        val, err := n.GetStringValue()
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
func (m *TeamProject) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *TeamProject) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TeamProject) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TeamProject) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *TeamProject) GetNumber()(*int32) {
    return m.number
}
// GetOrganizationPermission gets the organization_permission property value. The organization permission for this project. Only present when owner is an organization.
// returns a *string when successful
func (m *TeamProject) GetOrganizationPermission()(*string) {
    return m.organization_permission
}
// GetOwnerUrl gets the owner_url property value. The owner_url property
// returns a *string when successful
func (m *TeamProject) GetOwnerUrl()(*string) {
    return m.owner_url
}
// GetPermissions gets the permissions property value. The permissions property
// returns a TeamProject_permissionsable when successful
func (m *TeamProject) GetPermissions()(TeamProject_permissionsable) {
    return m.permissions
}
// GetPrivate gets the private property value. Whether the project is private or not. Only present when owner is an organization.
// returns a *bool when successful
func (m *TeamProject) GetPrivate()(*bool) {
    return m.private
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *TeamProject) GetState()(*string) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *TeamProject) GetUpdatedAt()(*string) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TeamProject) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TeamProject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
    {
        err := writer.WriteStringValue("organization_permission", m.GetOrganizationPermission())
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
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
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
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
func (m *TeamProject) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The body property
func (m *TeamProject) SetBody(value *string)() {
    m.body = value
}
// SetColumnsUrl sets the columns_url property value. The columns_url property
func (m *TeamProject) SetColumnsUrl(value *string)() {
    m.columns_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *TeamProject) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *TeamProject) SetCreator(value SimpleUserable)() {
    m.creator = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TeamProject) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *TeamProject) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *TeamProject) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TeamProject) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The number property
func (m *TeamProject) SetNumber(value *int32)() {
    m.number = value
}
// SetOrganizationPermission sets the organization_permission property value. The organization permission for this project. Only present when owner is an organization.
func (m *TeamProject) SetOrganizationPermission(value *string)() {
    m.organization_permission = value
}
// SetOwnerUrl sets the owner_url property value. The owner_url property
func (m *TeamProject) SetOwnerUrl(value *string)() {
    m.owner_url = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *TeamProject) SetPermissions(value TeamProject_permissionsable)() {
    m.permissions = value
}
// SetPrivate sets the private property value. Whether the project is private or not. Only present when owner is an organization.
func (m *TeamProject) SetPrivate(value *bool)() {
    m.private = value
}
// SetState sets the state property value. The state property
func (m *TeamProject) SetState(value *string)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *TeamProject) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *TeamProject) SetUrl(value *string)() {
    m.url = value
}
type TeamProjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetColumnsUrl()(*string)
    GetCreatedAt()(*string)
    GetCreator()(SimpleUserable)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetOrganizationPermission()(*string)
    GetOwnerUrl()(*string)
    GetPermissions()(TeamProject_permissionsable)
    GetPrivate()(*bool)
    GetState()(*string)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetBody(value *string)()
    SetColumnsUrl(value *string)()
    SetCreatedAt(value *string)()
    SetCreator(value SimpleUserable)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetOrganizationPermission(value *string)()
    SetOwnerUrl(value *string)()
    SetPermissions(value TeamProject_permissionsable)()
    SetPrivate(value *bool)()
    SetState(value *string)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
