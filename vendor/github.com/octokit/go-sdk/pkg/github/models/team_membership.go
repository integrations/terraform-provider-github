package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamMembership team Membership
type TeamMembership struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The role of the user in the team.
    role *TeamMembership_role
    // The state of the user's membership in the team.
    state *TeamMembership_state
    // The url property
    url *string
}
// NewTeamMembership instantiates a new TeamMembership and sets the default values.
func NewTeamMembership()(*TeamMembership) {
    m := &TeamMembership{
    }
    m.SetAdditionalData(make(map[string]any))
    roleValue := MEMBER_TEAMMEMBERSHIP_ROLE
    m.SetRole(&roleValue)
    return m
}
// CreateTeamMembershipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamMembershipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamMembership(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamMembership) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamMembership) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamMembership_role)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*TeamMembership_role))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamMembership_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*TeamMembership_state))
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
// GetRole gets the role property value. The role of the user in the team.
// returns a *TeamMembership_role when successful
func (m *TeamMembership) GetRole()(*TeamMembership_role) {
    return m.role
}
// GetState gets the state property value. The state of the user's membership in the team.
// returns a *TeamMembership_state when successful
func (m *TeamMembership) GetState()(*TeamMembership_state) {
    return m.state
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TeamMembership) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TeamMembership) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMembership) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRole sets the role property value. The role of the user in the team.
func (m *TeamMembership) SetRole(value *TeamMembership_role)() {
    m.role = value
}
// SetState sets the state property value. The state of the user's membership in the team.
func (m *TeamMembership) SetState(value *TeamMembership_state)() {
    m.state = value
}
// SetUrl sets the url property value. The url property
func (m *TeamMembership) SetUrl(value *string)() {
    m.url = value
}
type TeamMembershipable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRole()(*TeamMembership_role)
    GetState()(*TeamMembership_state)
    GetUrl()(*string)
    SetRole(value *TeamMembership_role)()
    SetState(value *TeamMembership_state)()
    SetUrl(value *string)()
}
