package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryAdvisoryUpdate_credits struct {
    // The username of the user credited.
    login *string
    // The type of credit the user is receiving.
    typeEscaped *SecurityAdvisoryCreditTypes
}
// NewRepositoryAdvisoryUpdate_credits instantiates a new RepositoryAdvisoryUpdate_credits and sets the default values.
func NewRepositoryAdvisoryUpdate_credits()(*RepositoryAdvisoryUpdate_credits) {
    m := &RepositoryAdvisoryUpdate_credits{
    }
    return m
}
// CreateRepositoryAdvisoryUpdate_creditsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisoryUpdate_creditsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryUpdate_credits(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisoryUpdate_credits) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAdvisoryCreditTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SecurityAdvisoryCreditTypes))
        }
        return nil
    }
    return res
}
// GetLogin gets the login property value. The username of the user credited.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate_credits) GetLogin()(*string) {
    return m.login
}
// GetTypeEscaped gets the type property value. The type of credit the user is receiving.
// returns a *SecurityAdvisoryCreditTypes when successful
func (m *RepositoryAdvisoryUpdate_credits) GetTypeEscaped()(*SecurityAdvisoryCreditTypes) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryUpdate_credits) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLogin sets the login property value. The username of the user credited.
func (m *RepositoryAdvisoryUpdate_credits) SetLogin(value *string)() {
    m.login = value
}
// SetTypeEscaped sets the type property value. The type of credit the user is receiving.
func (m *RepositoryAdvisoryUpdate_credits) SetTypeEscaped(value *SecurityAdvisoryCreditTypes)() {
    m.typeEscaped = value
}
type RepositoryAdvisoryUpdate_creditsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLogin()(*string)
    GetTypeEscaped()(*SecurityAdvisoryCreditTypes)
    SetLogin(value *string)()
    SetTypeEscaped(value *SecurityAdvisoryCreditTypes)()
}
