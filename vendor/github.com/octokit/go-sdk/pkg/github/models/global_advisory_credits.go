package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GlobalAdvisory_credits struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of credit the user is receiving.
    typeEscaped *SecurityAdvisoryCreditTypes
    // A GitHub user.
    user SimpleUserable
}
// NewGlobalAdvisory_credits instantiates a new GlobalAdvisory_credits and sets the default values.
func NewGlobalAdvisory_credits()(*GlobalAdvisory_credits) {
    m := &GlobalAdvisory_credits{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGlobalAdvisory_creditsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGlobalAdvisory_creditsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGlobalAdvisory_credits(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GlobalAdvisory_credits) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GlobalAdvisory_credits) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetTypeEscaped gets the type property value. The type of credit the user is receiving.
// returns a *SecurityAdvisoryCreditTypes when successful
func (m *GlobalAdvisory_credits) GetTypeEscaped()(*SecurityAdvisoryCreditTypes) {
    return m.typeEscaped
}
// GetUser gets the user property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *GlobalAdvisory_credits) GetUser()(SimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *GlobalAdvisory_credits) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GlobalAdvisory_credits) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type of credit the user is receiving.
func (m *GlobalAdvisory_credits) SetTypeEscaped(value *SecurityAdvisoryCreditTypes)() {
    m.typeEscaped = value
}
// SetUser sets the user property value. A GitHub user.
func (m *GlobalAdvisory_credits) SetUser(value SimpleUserable)() {
    m.user = value
}
type GlobalAdvisory_creditsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*SecurityAdvisoryCreditTypes)
    GetUser()(SimpleUserable)
    SetTypeEscaped(value *SecurityAdvisoryCreditTypes)()
    SetUser(value SimpleUserable)()
}
