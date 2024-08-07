package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomPropertyValue custom property name and associated value
type CustomPropertyValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the property
    property_name *string
    // The value assigned to the property
    value CustomPropertyValue_CustomPropertyValue_valueable
}
// CustomPropertyValue_CustomPropertyValue_value composed type wrapper for classes string
type CustomPropertyValue_CustomPropertyValue_value struct {
    // Composed type representation for type string
    string *string
}
// NewCustomPropertyValue_CustomPropertyValue_value instantiates a new CustomPropertyValue_CustomPropertyValue_value and sets the default values.
func NewCustomPropertyValue_CustomPropertyValue_value()(*CustomPropertyValue_CustomPropertyValue_value) {
    m := &CustomPropertyValue_CustomPropertyValue_value{
    }
    return m
}
// CreateCustomPropertyValue_CustomPropertyValue_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomPropertyValue_CustomPropertyValue_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCustomPropertyValue_CustomPropertyValue_value()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomPropertyValue_CustomPropertyValue_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CustomPropertyValue_CustomPropertyValue_value) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *CustomPropertyValue_CustomPropertyValue_value) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *CustomPropertyValue_CustomPropertyValue_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetString sets the string property value. Composed type representation for type string
func (m *CustomPropertyValue_CustomPropertyValue_value) SetString(value *string)() {
    m.string = value
}
type CustomPropertyValue_CustomPropertyValue_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    SetString(value *string)()
}
// NewCustomPropertyValue instantiates a new CustomPropertyValue and sets the default values.
func NewCustomPropertyValue()(*CustomPropertyValue) {
    m := &CustomPropertyValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomPropertyValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomPropertyValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomPropertyValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomPropertyValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomPropertyValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["property_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyName(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomPropertyValue_CustomPropertyValue_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(CustomPropertyValue_CustomPropertyValue_valueable))
        }
        return nil
    }
    return res
}
// GetPropertyName gets the property_name property value. The name of the property
// returns a *string when successful
func (m *CustomPropertyValue) GetPropertyName()(*string) {
    return m.property_name
}
// GetValue gets the value property value. The value assigned to the property
// returns a CustomPropertyValue_CustomPropertyValue_valueable when successful
func (m *CustomPropertyValue) GetValue()(CustomPropertyValue_CustomPropertyValue_valueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *CustomPropertyValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("property_name", m.GetPropertyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *CustomPropertyValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPropertyName sets the property_name property value. The name of the property
func (m *CustomPropertyValue) SetPropertyName(value *string)() {
    m.property_name = value
}
// SetValue sets the value property value. The value assigned to the property
func (m *CustomPropertyValue) SetValue(value CustomPropertyValue_CustomPropertyValue_valueable)() {
    m.value = value
}
type CustomPropertyValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPropertyName()(*string)
    GetValue()(CustomPropertyValue_CustomPropertyValue_valueable)
    SetPropertyName(value *string)()
    SetValue(value CustomPropertyValue_CustomPropertyValue_valueable)()
}
