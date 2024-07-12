package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryInvitation repository invitations let you manage who you collaborate with.
type RepositoryInvitation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Whether or not the invitation has expired
    expired *bool
    // The html_url property
    html_url *string
    // Unique identifier of the repository invitation.
    id *int64
    // A GitHub user.
    invitee NullableSimpleUserable
    // A GitHub user.
    inviter NullableSimpleUserable
    // The node_id property
    node_id *string
    // The permission associated with the invitation.
    permissions *RepositoryInvitation_permissions
    // Minimal Repository
    repository MinimalRepositoryable
    // URL for the repository invitation
    url *string
}
// NewRepositoryInvitation instantiates a new RepositoryInvitation and sets the default values.
func NewRepositoryInvitation()(*RepositoryInvitation) {
    m := &RepositoryInvitation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryInvitationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryInvitationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryInvitation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryInvitation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *RepositoryInvitation) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetExpired gets the expired property value. Whether or not the invitation has expired
// returns a *bool when successful
func (m *RepositoryInvitation) GetExpired()(*bool) {
    return m.expired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryInvitation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["expired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpired(val)
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
    res["invitee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitee(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["inviter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviter(val.(NullableSimpleUserable))
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
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryInvitation_permissions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(*RepositoryInvitation_permissions))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
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
func (m *RepositoryInvitation) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the repository invitation.
// returns a *int64 when successful
func (m *RepositoryInvitation) GetId()(*int64) {
    return m.id
}
// GetInvitee gets the invitee property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *RepositoryInvitation) GetInvitee()(NullableSimpleUserable) {
    return m.invitee
}
// GetInviter gets the inviter property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *RepositoryInvitation) GetInviter()(NullableSimpleUserable) {
    return m.inviter
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *RepositoryInvitation) GetNodeId()(*string) {
    return m.node_id
}
// GetPermissions gets the permissions property value. The permission associated with the invitation.
// returns a *RepositoryInvitation_permissions when successful
func (m *RepositoryInvitation) GetPermissions()(*RepositoryInvitation_permissions) {
    return m.permissions
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *RepositoryInvitation) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetUrl gets the url property value. URL for the repository invitation
// returns a *string when successful
func (m *RepositoryInvitation) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *RepositoryInvitation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("expired", m.GetExpired())
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
        err := writer.WriteObjectValue("invitee", m.GetInvitee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("inviter", m.GetInviter())
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
    if m.GetPermissions() != nil {
        cast := (*m.GetPermissions()).String()
        err := writer.WriteStringValue("permissions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
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
func (m *RepositoryInvitation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *RepositoryInvitation) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetExpired sets the expired property value. Whether or not the invitation has expired
func (m *RepositoryInvitation) SetExpired(value *bool)() {
    m.expired = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *RepositoryInvitation) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the repository invitation.
func (m *RepositoryInvitation) SetId(value *int64)() {
    m.id = value
}
// SetInvitee sets the invitee property value. A GitHub user.
func (m *RepositoryInvitation) SetInvitee(value NullableSimpleUserable)() {
    m.invitee = value
}
// SetInviter sets the inviter property value. A GitHub user.
func (m *RepositoryInvitation) SetInviter(value NullableSimpleUserable)() {
    m.inviter = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *RepositoryInvitation) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPermissions sets the permissions property value. The permission associated with the invitation.
func (m *RepositoryInvitation) SetPermissions(value *RepositoryInvitation_permissions)() {
    m.permissions = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *RepositoryInvitation) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetUrl sets the url property value. URL for the repository invitation
func (m *RepositoryInvitation) SetUrl(value *string)() {
    m.url = value
}
type RepositoryInvitationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpired()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetInvitee()(NullableSimpleUserable)
    GetInviter()(NullableSimpleUserable)
    GetNodeId()(*string)
    GetPermissions()(*RepositoryInvitation_permissions)
    GetRepository()(MinimalRepositoryable)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpired(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetInvitee(value NullableSimpleUserable)()
    SetInviter(value NullableSimpleUserable)()
    SetNodeId(value *string)()
    SetPermissions(value *RepositoryInvitation_permissions)()
    SetRepository(value MinimalRepositoryable)()
    SetUrl(value *string)()
}
