package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GlobalAdvisory_identifiers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of identifier.
    typeEscaped *GlobalAdvisory_identifiers_type
    // The identifier value.
    value *string
}
// NewGlobalAdvisory_identifiers instantiates a new GlobalAdvisory_identifiers and sets the default values.
func NewGlobalAdvisory_identifiers()(*GlobalAdvisory_identifiers) {
    m := &GlobalAdvisory_identifiers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGlobalAdvisory_identifiersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGlobalAdvisory_identifiersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGlobalAdvisory_identifiers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GlobalAdvisory_identifiers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GlobalAdvisory_identifiers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGlobalAdvisory_identifiers_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*GlobalAdvisory_identifiers_type))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type of identifier.
// returns a *GlobalAdvisory_identifiers_type when successful
func (m *GlobalAdvisory_identifiers) GetTypeEscaped()(*GlobalAdvisory_identifiers_type) {
    return m.typeEscaped
}
// GetValue gets the value property value. The identifier value.
// returns a *string when successful
func (m *GlobalAdvisory_identifiers) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *GlobalAdvisory_identifiers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *GlobalAdvisory_identifiers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type of identifier.
func (m *GlobalAdvisory_identifiers) SetTypeEscaped(value *GlobalAdvisory_identifiers_type)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The identifier value.
func (m *GlobalAdvisory_identifiers) SetValue(value *string)() {
    m.value = value
}
type GlobalAdvisory_identifiersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*GlobalAdvisory_identifiers_type)
    GetValue()(*string)
    SetTypeEscaped(value *GlobalAdvisory_identifiers_type)()
    SetValue(value *string)()
}
