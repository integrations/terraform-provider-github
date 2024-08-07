package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryAdvisoryCredit a credit given to a user for a repository security advisory.
type RepositoryAdvisoryCredit struct {
    // The state of the user's acceptance of the credit.
    state *RepositoryAdvisoryCredit_state
    // The type of credit the user is receiving.
    typeEscaped *SecurityAdvisoryCreditTypes
    // A GitHub user.
    user SimpleUserable
}
// NewRepositoryAdvisoryCredit instantiates a new RepositoryAdvisoryCredit and sets the default values.
func NewRepositoryAdvisoryCredit()(*RepositoryAdvisoryCredit) {
    m := &RepositoryAdvisoryCredit{
    }
    return m
}
// CreateRepositoryAdvisoryCreditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisoryCreditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryCredit(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisoryCredit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryAdvisoryCredit_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*RepositoryAdvisoryCredit_state))
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(SimpleUserable))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state of the user's acceptance of the credit.
// returns a *RepositoryAdvisoryCredit_state when successful
func (m *RepositoryAdvisoryCredit) GetState()(*RepositoryAdvisoryCredit_state) {
    return m.state
}
// GetTypeEscaped gets the type property value. The type of credit the user is receiving.
// returns a *SecurityAdvisoryCreditTypes when successful
func (m *RepositoryAdvisoryCredit) GetTypeEscaped()(*SecurityAdvisoryCreditTypes) {
    return m.typeEscaped
}
// GetUser gets the user property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *RepositoryAdvisoryCredit) GetUser()(SimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryCredit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
    {
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetState sets the state property value. The state of the user's acceptance of the credit.
func (m *RepositoryAdvisoryCredit) SetState(value *RepositoryAdvisoryCredit_state)() {
    m.state = value
}
// SetTypeEscaped sets the type property value. The type of credit the user is receiving.
func (m *RepositoryAdvisoryCredit) SetTypeEscaped(value *SecurityAdvisoryCreditTypes)() {
    m.typeEscaped = value
}
// SetUser sets the user property value. A GitHub user.
func (m *RepositoryAdvisoryCredit) SetUser(value SimpleUserable)() {
    m.user = value
}
type RepositoryAdvisoryCreditable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetState()(*RepositoryAdvisoryCredit_state)
    GetTypeEscaped()(*SecurityAdvisoryCreditTypes)
    GetUser()(SimpleUserable)
    SetState(value *RepositoryAdvisoryCredit_state)()
    SetTypeEscaped(value *SecurityAdvisoryCreditTypes)()
    SetUser(value SimpleUserable)()
}
