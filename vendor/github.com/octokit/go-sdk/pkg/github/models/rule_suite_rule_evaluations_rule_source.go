package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RuleSuite_rule_evaluations_rule_source struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the rule source.
    id *int32
    // The name of the rule source.
    name *string
    // The type of rule source.
    typeEscaped *string
}
// NewRuleSuite_rule_evaluations_rule_source instantiates a new RuleSuite_rule_evaluations_rule_source and sets the default values.
func NewRuleSuite_rule_evaluations_rule_source()(*RuleSuite_rule_evaluations_rule_source) {
    m := &RuleSuite_rule_evaluations_rule_source{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRuleSuite_rule_evaluations_rule_sourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRuleSuite_rule_evaluations_rule_sourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleSuite_rule_evaluations_rule_source(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RuleSuite_rule_evaluations_rule_source) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RuleSuite_rule_evaluations_rule_source) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the rule source.
// returns a *int32 when successful
func (m *RuleSuite_rule_evaluations_rule_source) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the rule source.
// returns a *string when successful
func (m *RuleSuite_rule_evaluations_rule_source) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. The type of rule source.
// returns a *string when successful
func (m *RuleSuite_rule_evaluations_rule_source) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RuleSuite_rule_evaluations_rule_source) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *RuleSuite_rule_evaluations_rule_source) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The ID of the rule source.
func (m *RuleSuite_rule_evaluations_rule_source) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the rule source.
func (m *RuleSuite_rule_evaluations_rule_source) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. The type of rule source.
func (m *RuleSuite_rule_evaluations_rule_source) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type RuleSuite_rule_evaluations_rule_sourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetTypeEscaped()(*string)
    SetId(value *int32)()
    SetName(value *string)()
    SetTypeEscaped(value *string)()
}
