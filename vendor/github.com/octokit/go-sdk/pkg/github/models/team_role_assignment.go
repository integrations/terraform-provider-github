package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamRoleAssignment the Relationship a Team has with a role.
type TeamRoleAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The members_url property
    members_url *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The notification_setting property
    notification_setting *string
    // Groups of organization members that gives permissions on specified repositories.
    parent NullableTeamSimpleable
    // The permission property
    permission *string
    // The permissions property
    permissions TeamRoleAssignment_permissionsable
    // The privacy property
    privacy *string
    // The repositories_url property
    repositories_url *string
    // The slug property
    slug *string
    // The url property
    url *string
}
// NewTeamRoleAssignment instantiates a new TeamRoleAssignment and sets the default values.
func NewTeamRoleAssignment()(*TeamRoleAssignment) {
    m := &TeamRoleAssignment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamRoleAssignment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamRoleAssignment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TeamRoleAssignment) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["members_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersUrl(val)
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
    res["notification_setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationSetting(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableTeamSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(NullableTeamSimpleable))
        }
        return nil
    }
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamRoleAssignment_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(TeamRoleAssignment_permissionsable))
        }
        return nil
    }
    res["privacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacy(val)
        }
        return nil
    }
    res["repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoriesUrl(val)
        }
        return nil
    }
    res["slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSlug(val)
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
func (m *TeamRoleAssignment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *TeamRoleAssignment) GetId()(*int32) {
    return m.id
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *TeamRoleAssignment) GetMembersUrl()(*string) {
    return m.members_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TeamRoleAssignment) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TeamRoleAssignment) GetNodeId()(*string) {
    return m.node_id
}
// GetNotificationSetting gets the notification_setting property value. The notification_setting property
// returns a *string when successful
func (m *TeamRoleAssignment) GetNotificationSetting()(*string) {
    return m.notification_setting
}
// GetParent gets the parent property value. Groups of organization members that gives permissions on specified repositories.
// returns a NullableTeamSimpleable when successful
func (m *TeamRoleAssignment) GetParent()(NullableTeamSimpleable) {
    return m.parent
}
// GetPermission gets the permission property value. The permission property
// returns a *string when successful
func (m *TeamRoleAssignment) GetPermission()(*string) {
    return m.permission
}
// GetPermissions gets the permissions property value. The permissions property
// returns a TeamRoleAssignment_permissionsable when successful
func (m *TeamRoleAssignment) GetPermissions()(TeamRoleAssignment_permissionsable) {
    return m.permissions
}
// GetPrivacy gets the privacy property value. The privacy property
// returns a *string when successful
func (m *TeamRoleAssignment) GetPrivacy()(*string) {
    return m.privacy
}
// GetRepositoriesUrl gets the repositories_url property value. The repositories_url property
// returns a *string when successful
func (m *TeamRoleAssignment) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetSlug gets the slug property value. The slug property
// returns a *string when successful
func (m *TeamRoleAssignment) GetSlug()(*string) {
    return m.slug
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TeamRoleAssignment) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TeamRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("members_url", m.GetMembersUrl())
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
        err := writer.WriteStringValue("notification_setting", m.GetNotificationSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("permission", m.GetPermission())
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
        err := writer.WriteStringValue("privacy", m.GetPrivacy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositories_url", m.GetRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("slug", m.GetSlug())
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
func (m *TeamRoleAssignment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *TeamRoleAssignment) SetDescription(value *string)() {
    m.description = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TeamRoleAssignment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *TeamRoleAssignment) SetId(value *int32)() {
    m.id = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *TeamRoleAssignment) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetName sets the name property value. The name property
func (m *TeamRoleAssignment) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TeamRoleAssignment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNotificationSetting sets the notification_setting property value. The notification_setting property
func (m *TeamRoleAssignment) SetNotificationSetting(value *string)() {
    m.notification_setting = value
}
// SetParent sets the parent property value. Groups of organization members that gives permissions on specified repositories.
func (m *TeamRoleAssignment) SetParent(value NullableTeamSimpleable)() {
    m.parent = value
}
// SetPermission sets the permission property value. The permission property
func (m *TeamRoleAssignment) SetPermission(value *string)() {
    m.permission = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *TeamRoleAssignment) SetPermissions(value TeamRoleAssignment_permissionsable)() {
    m.permissions = value
}
// SetPrivacy sets the privacy property value. The privacy property
func (m *TeamRoleAssignment) SetPrivacy(value *string)() {
    m.privacy = value
}
// SetRepositoriesUrl sets the repositories_url property value. The repositories_url property
func (m *TeamRoleAssignment) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetSlug sets the slug property value. The slug property
func (m *TeamRoleAssignment) SetSlug(value *string)() {
    m.slug = value
}
// SetUrl sets the url property value. The url property
func (m *TeamRoleAssignment) SetUrl(value *string)() {
    m.url = value
}
type TeamRoleAssignmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetNotificationSetting()(*string)
    GetParent()(NullableTeamSimpleable)
    GetPermission()(*string)
    GetPermissions()(TeamRoleAssignment_permissionsable)
    GetPrivacy()(*string)
    GetRepositoriesUrl()(*string)
    GetSlug()(*string)
    GetUrl()(*string)
    SetDescription(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNotificationSetting(value *string)()
    SetParent(value NullableTeamSimpleable)()
    SetPermission(value *string)()
    SetPermissions(value TeamRoleAssignment_permissionsable)()
    SetPrivacy(value *string)()
    SetRepositoriesUrl(value *string)()
    SetSlug(value *string)()
    SetUrl(value *string)()
}
