package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ValidationError_errors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The field property
    field *string
    // The index property
    index *int32
    // The message property
    message *string
    // The resource property
    resource *string
    // The value property
    value ValidationError_errors_ValidationError_errors_valueable
}
// ValidationError_errors_ValidationError_errors_value composed type wrapper for classes int32, string
type ValidationError_errors_ValidationError_errors_value struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewValidationError_errors_ValidationError_errors_value instantiates a new ValidationError_errors_ValidationError_errors_value and sets the default values.
func NewValidationError_errors_ValidationError_errors_value()(*ValidationError_errors_ValidationError_errors_value) {
    m := &ValidationError_errors_ValidationError_errors_value{
    }
    return m
}
// CreateValidationError_errors_ValidationError_errors_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidationError_errors_ValidationError_errors_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewValidationError_errors_ValidationError_errors_value()
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
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidationError_errors_ValidationError_errors_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *ValidationError_errors_ValidationError_errors_value) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ValidationError_errors_ValidationError_errors_value) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ValidationError_errors_ValidationError_errors_value) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ValidationError_errors_ValidationError_errors_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type int32
func (m *ValidationError_errors_ValidationError_errors_value) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ValidationError_errors_ValidationError_errors_value) SetString(value *string)() {
    m.string = value
}
type ValidationError_errors_ValidationError_errors_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewValidationError_errors instantiates a new ValidationError_errors and sets the default values.
func NewValidationError_errors()(*ValidationError_errors) {
    m := &ValidationError_errors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidationError_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateValidationError_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidationError_errors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ValidationError_errors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *ValidationError_errors) GetCode()(*string) {
    return m.code
}
// GetField gets the field property value. The field property
// returns a *string when successful
func (m *ValidationError_errors) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ValidationError_errors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateValidationError_errors_ValidationError_errors_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ValidationError_errors_ValidationError_errors_valueable))
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The index property
// returns a *int32 when successful
func (m *ValidationError_errors) GetIndex()(*int32) {
    return m.index
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *ValidationError_errors) GetMessage()(*string) {
    return m.message
}
// GetResource gets the resource property value. The resource property
// returns a *string when successful
func (m *ValidationError_errors) GetResource()(*string) {
    return m.resource
}
// GetValue gets the value property value. The value property
// returns a ValidationError_errors_ValidationError_errors_valueable when successful
func (m *ValidationError_errors) GetValue()(ValidationError_errors_ValidationError_errors_valueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ValidationError_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("field", m.GetField())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resource", m.GetResource())
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
func (m *ValidationError_errors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *ValidationError_errors) SetCode(value *string)() {
    m.code = value
}
// SetField sets the field property value. The field property
func (m *ValidationError_errors) SetField(value *string)() {
    m.field = value
}
// SetIndex sets the index property value. The index property
func (m *ValidationError_errors) SetIndex(value *int32)() {
    m.index = value
}
// SetMessage sets the message property value. The message property
func (m *ValidationError_errors) SetMessage(value *string)() {
    m.message = value
}
// SetResource sets the resource property value. The resource property
func (m *ValidationError_errors) SetResource(value *string)() {
    m.resource = value
}
// SetValue sets the value property value. The value property
func (m *ValidationError_errors) SetValue(value ValidationError_errors_ValidationError_errors_valueable)() {
    m.value = value
}
type ValidationError_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetField()(*string)
    GetIndex()(*int32)
    GetMessage()(*string)
    GetResource()(*string)
    GetValue()(ValidationError_errors_ValidationError_errors_valueable)
    SetCode(value *string)()
    SetField(value *string)()
    SetIndex(value *int32)()
    SetMessage(value *string)()
    SetResource(value *string)()
    SetValue(value ValidationError_errors_ValidationError_errors_valueable)()
}
