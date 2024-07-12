package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
    allowed_values []string
    // Default value of the property
    default_value ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable
    // Short description of the property
    description *string
    // Whether the property is required.
    required *bool
}
// ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value composed type wrapper for classes string
type ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value struct {
    // Composed type representation for type string
    string *string
}
// NewItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value instantiates a new ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value and sets the default values.
func NewItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value()(*ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value) {
    m := &ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value{
    }
    return m
}
// CreateItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value()
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
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_value) SetString(value *string)() {
    m.string = value
}
type ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    SetString(value *string)()
}
// NewItemPropertiesSchemaItemWithCustom_property_namePutRequestBody instantiates a new ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody and sets the default values.
func NewItemPropertiesSchemaItemWithCustom_property_namePutRequestBody()(*ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) {
    m := &ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPropertiesSchemaItemWithCustom_property_namePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPropertiesSchemaItemWithCustom_property_namePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPropertiesSchemaItemWithCustom_property_namePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedValues gets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
// returns a []string when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) GetAllowedValues()([]string) {
    return m.allowed_values
}
// GetDefaultValue gets the default_value property value. Default value of the property
// returns a ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) GetDefaultValue()(ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable) {
    return m.default_value
}
// GetDescription gets the description property value. Short description of the property
// returns a *string when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAllowedValues(res)
        }
        return nil
    }
    res["default_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    return res
}
// GetRequired gets the required property value. Whether the property is required.
// returns a *bool when successful
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) GetRequired()(*bool) {
    return m.required
}
// Serialize serializes information the current object
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedValues() != nil {
        err := writer.WriteCollectionOfStringValues("allowed_values", m.GetAllowedValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("default_value", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
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
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedValues sets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) SetAllowedValues(value []string)() {
    m.allowed_values = value
}
// SetDefaultValue sets the default_value property value. Default value of the property
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) SetDefaultValue(value ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable)() {
    m.default_value = value
}
// SetDescription sets the description property value. Short description of the property
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetRequired sets the required property value. Whether the property is required.
func (m *ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody) SetRequired(value *bool)() {
    m.required = value
}
type ItemPropertiesSchemaItemWithCustom_property_namePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedValues()([]string)
    GetDefaultValue()(ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable)
    GetDescription()(*string)
    GetRequired()(*bool)
    SetAllowedValues(value []string)()
    SetDefaultValue(value ItemPropertiesSchemaItemWithCustom_property_namePutRequestBody_WithCustom_property_namePutRequestBody_default_valueable)()
    SetDescription(value *string)()
    SetRequired(value *bool)()
}
