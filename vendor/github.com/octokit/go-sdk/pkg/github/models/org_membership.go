package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgMembership org Membership
type OrgMembership struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub organization.
    organization OrganizationSimpleable
    // The organization_url property
    organization_url *string
    // The permissions property
    permissions OrgMembership_permissionsable
    // The user's membership type in the organization.
    role *OrgMembership_role
    // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
    state *OrgMembership_state
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewOrgMembership instantiates a new OrgMembership and sets the default values.
func NewOrgMembership()(*OrgMembership) {
    m := &OrgMembership{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrgMembershipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrgMembershipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgMembership(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrgMembership) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrgMembership) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(OrganizationSimpleable))
        }
        return nil
    }
    res["organization_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationUrl(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrgMembership_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(OrgMembership_permissionsable))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrgMembership_role)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*OrgMembership_role))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrgMembership_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*OrgMembership_state))
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
// GetOrganization gets the organization property value. A GitHub organization.
// returns a OrganizationSimpleable when successful
func (m *OrgMembership) GetOrganization()(OrganizationSimpleable) {
    return m.organization
}
// GetOrganizationUrl gets the organization_url property value. The organization_url property
// returns a *string when successful
func (m *OrgMembership) GetOrganizationUrl()(*string) {
    return m.organization_url
}
// GetPermissions gets the permissions property value. The permissions property
// returns a OrgMembership_permissionsable when successful
func (m *OrgMembership) GetPermissions()(OrgMembership_permissionsable) {
    return m.permissions
}
// GetRole gets the role property value. The user's membership type in the organization.
// returns a *OrgMembership_role when successful
func (m *OrgMembership) GetRole()(*OrgMembership_role) {
    return m.role
}
// GetState gets the state property value. The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
// returns a *OrgMembership_state when successful
func (m *OrgMembership) GetState()(*OrgMembership_state) {
    return m.state
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *OrgMembership) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *OrgMembership) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *OrgMembership) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_url", m.GetOrganizationUrl())
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
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *OrgMembership) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOrganization sets the organization property value. A GitHub organization.
func (m *OrgMembership) SetOrganization(value OrganizationSimpleable)() {
    m.organization = value
}
// SetOrganizationUrl sets the organization_url property value. The organization_url property
func (m *OrgMembership) SetOrganizationUrl(value *string)() {
    m.organization_url = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *OrgMembership) SetPermissions(value OrgMembership_permissionsable)() {
    m.permissions = value
}
// SetRole sets the role property value. The user's membership type in the organization.
func (m *OrgMembership) SetRole(value *OrgMembership_role)() {
    m.role = value
}
// SetState sets the state property value. The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
func (m *OrgMembership) SetState(value *OrgMembership_state)() {
    m.state = value
}
// SetUrl sets the url property value. The url property
func (m *OrgMembership) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *OrgMembership) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type OrgMembershipable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrganization()(OrganizationSimpleable)
    GetOrganizationUrl()(*string)
    GetPermissions()(OrgMembership_permissionsable)
    GetRole()(*OrgMembership_role)
    GetState()(*OrgMembership_state)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetOrganization(value OrganizationSimpleable)()
    SetOrganizationUrl(value *string)()
    SetPermissions(value OrgMembership_permissionsable)()
    SetRole(value *OrgMembership_role)()
    SetState(value *OrgMembership_state)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
