package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationInvitation organization Invitation
type OrganizationInvitation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The email property
    email *string
    // The failed_at property
    failed_at *string
    // The failed_reason property
    failed_reason *string
    // The id property
    id *int64
    // The invitation_source property
    invitation_source *string
    // The invitation_teams_url property
    invitation_teams_url *string
    // A GitHub user.
    inviter SimpleUserable
    // The login property
    login *string
    // The node_id property
    node_id *string
    // The role property
    role *string
    // The team_count property
    team_count *int32
}
// NewOrganizationInvitation instantiates a new OrganizationInvitation and sets the default values.
func NewOrganizationInvitation()(*OrganizationInvitation) {
    m := &OrganizationInvitation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationInvitationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationInvitationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationInvitation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationInvitation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *OrganizationInvitation) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *OrganizationInvitation) GetEmail()(*string) {
    return m.email
}
// GetFailedAt gets the failed_at property value. The failed_at property
// returns a *string when successful
func (m *OrganizationInvitation) GetFailedAt()(*string) {
    return m.failed_at
}
// GetFailedReason gets the failed_reason property value. The failed_reason property
// returns a *string when successful
func (m *OrganizationInvitation) GetFailedReason()(*string) {
    return m.failed_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationInvitation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["failed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedAt(val)
        }
        return nil
    }
    res["failed_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedReason(val)
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
    res["invitation_source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationSource(val)
        }
        return nil
    }
    res["invitation_teams_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationTeamsUrl(val)
        }
        return nil
    }
    res["inviter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviter(val.(SimpleUserable))
        }
        return nil
    }
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
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
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["team_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamCount(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrganizationInvitation) GetId()(*int64) {
    return m.id
}
// GetInvitationSource gets the invitation_source property value. The invitation_source property
// returns a *string when successful
func (m *OrganizationInvitation) GetInvitationSource()(*string) {
    return m.invitation_source
}
// GetInvitationTeamsUrl gets the invitation_teams_url property value. The invitation_teams_url property
// returns a *string when successful
func (m *OrganizationInvitation) GetInvitationTeamsUrl()(*string) {
    return m.invitation_teams_url
}
// GetInviter gets the inviter property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *OrganizationInvitation) GetInviter()(SimpleUserable) {
    return m.inviter
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *OrganizationInvitation) GetLogin()(*string) {
    return m.login
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *OrganizationInvitation) GetNodeId()(*string) {
    return m.node_id
}
// GetRole gets the role property value. The role property
// returns a *string when successful
func (m *OrganizationInvitation) GetRole()(*string) {
    return m.role
}
// GetTeamCount gets the team_count property value. The team_count property
// returns a *int32 when successful
func (m *OrganizationInvitation) GetTeamCount()(*int32) {
    return m.team_count
}
// Serialize serializes information the current object
func (m *OrganizationInvitation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failed_at", m.GetFailedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failed_reason", m.GetFailedReason())
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
        err := writer.WriteStringValue("invitation_source", m.GetInvitationSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("invitation_teams_url", m.GetInvitationTeamsUrl())
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
        err := writer.WriteStringValue("login", m.GetLogin())
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
        err := writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("team_count", m.GetTeamCount())
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
func (m *OrganizationInvitation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *OrganizationInvitation) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEmail sets the email property value. The email property
func (m *OrganizationInvitation) SetEmail(value *string)() {
    m.email = value
}
// SetFailedAt sets the failed_at property value. The failed_at property
func (m *OrganizationInvitation) SetFailedAt(value *string)() {
    m.failed_at = value
}
// SetFailedReason sets the failed_reason property value. The failed_reason property
func (m *OrganizationInvitation) SetFailedReason(value *string)() {
    m.failed_reason = value
}
// SetId sets the id property value. The id property
func (m *OrganizationInvitation) SetId(value *int64)() {
    m.id = value
}
// SetInvitationSource sets the invitation_source property value. The invitation_source property
func (m *OrganizationInvitation) SetInvitationSource(value *string)() {
    m.invitation_source = value
}
// SetInvitationTeamsUrl sets the invitation_teams_url property value. The invitation_teams_url property
func (m *OrganizationInvitation) SetInvitationTeamsUrl(value *string)() {
    m.invitation_teams_url = value
}
// SetInviter sets the inviter property value. A GitHub user.
func (m *OrganizationInvitation) SetInviter(value SimpleUserable)() {
    m.inviter = value
}
// SetLogin sets the login property value. The login property
func (m *OrganizationInvitation) SetLogin(value *string)() {
    m.login = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *OrganizationInvitation) SetNodeId(value *string)() {
    m.node_id = value
}
// SetRole sets the role property value. The role property
func (m *OrganizationInvitation) SetRole(value *string)() {
    m.role = value
}
// SetTeamCount sets the team_count property value. The team_count property
func (m *OrganizationInvitation) SetTeamCount(value *int32)() {
    m.team_count = value
}
type OrganizationInvitationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetEmail()(*string)
    GetFailedAt()(*string)
    GetFailedReason()(*string)
    GetId()(*int64)
    GetInvitationSource()(*string)
    GetInvitationTeamsUrl()(*string)
    GetInviter()(SimpleUserable)
    GetLogin()(*string)
    GetNodeId()(*string)
    GetRole()(*string)
    GetTeamCount()(*int32)
    SetCreatedAt(value *string)()
    SetEmail(value *string)()
    SetFailedAt(value *string)()
    SetFailedReason(value *string)()
    SetId(value *int64)()
    SetInvitationSource(value *string)()
    SetInvitationTeamsUrl(value *string)()
    SetInviter(value SimpleUserable)()
    SetLogin(value *string)()
    SetNodeId(value *string)()
    SetRole(value *string)()
    SetTeamCount(value *int32)()
}
