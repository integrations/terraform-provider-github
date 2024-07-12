package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RuleSuite_rule_evaluations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Any associated details with the rule evaluation.
    details *string
    // The enforcement level of this rule source.
    enforcement *RuleSuite_rule_evaluations_enforcement
    // The result of the evaluation of the individual rule.
    result *RuleSuite_rule_evaluations_result
    // The rule_source property
    rule_source RuleSuite_rule_evaluations_rule_sourceable
    // The type of rule.
    rule_type *string
}
// NewRuleSuite_rule_evaluations instantiates a new RuleSuite_rule_evaluations and sets the default values.
func NewRuleSuite_rule_evaluations()(*RuleSuite_rule_evaluations) {
    m := &RuleSuite_rule_evaluations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRuleSuite_rule_evaluationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRuleSuite_rule_evaluationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleSuite_rule_evaluations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RuleSuite_rule_evaluations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetails gets the details property value. Any associated details with the rule evaluation.
// returns a *string when successful
func (m *RuleSuite_rule_evaluations) GetDetails()(*string) {
    return m.details
}
// GetEnforcement gets the enforcement property value. The enforcement level of this rule source.
// returns a *RuleSuite_rule_evaluations_enforcement when successful
func (m *RuleSuite_rule_evaluations) GetEnforcement()(*RuleSuite_rule_evaluations_enforcement) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RuleSuite_rule_evaluations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val)
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleSuite_rule_evaluations_enforcement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val.(*RuleSuite_rule_evaluations_enforcement))
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleSuite_rule_evaluations_result)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*RuleSuite_rule_evaluations_result))
        }
        return nil
    }
    res["rule_source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRuleSuite_rule_evaluations_rule_sourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleSource(val.(RuleSuite_rule_evaluations_rule_sourceable))
        }
        return nil
    }
    res["rule_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val)
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. The result of the evaluation of the individual rule.
// returns a *RuleSuite_rule_evaluations_result when successful
func (m *RuleSuite_rule_evaluations) GetResult()(*RuleSuite_rule_evaluations_result) {
    return m.result
}
// GetRuleSource gets the rule_source property value. The rule_source property
// returns a RuleSuite_rule_evaluations_rule_sourceable when successful
func (m *RuleSuite_rule_evaluations) GetRuleSource()(RuleSuite_rule_evaluations_rule_sourceable) {
    return m.rule_source
}
// GetRuleType gets the rule_type property value. The type of rule.
// returns a *string when successful
func (m *RuleSuite_rule_evaluations) GetRuleType()(*string) {
    return m.rule_type
}
// Serialize serializes information the current object
func (m *RuleSuite_rule_evaluations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcement() != nil {
        cast := (*m.GetEnforcement()).String()
        err := writer.WriteStringValue("enforcement", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err := writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rule_source", m.GetRuleSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rule_type", m.GetRuleType())
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
func (m *RuleSuite_rule_evaluations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetails sets the details property value. Any associated details with the rule evaluation.
func (m *RuleSuite_rule_evaluations) SetDetails(value *string)() {
    m.details = value
}
// SetEnforcement sets the enforcement property value. The enforcement level of this rule source.
func (m *RuleSuite_rule_evaluations) SetEnforcement(value *RuleSuite_rule_evaluations_enforcement)() {
    m.enforcement = value
}
// SetResult sets the result property value. The result of the evaluation of the individual rule.
func (m *RuleSuite_rule_evaluations) SetResult(value *RuleSuite_rule_evaluations_result)() {
    m.result = value
}
// SetRuleSource sets the rule_source property value. The rule_source property
func (m *RuleSuite_rule_evaluations) SetRuleSource(value RuleSuite_rule_evaluations_rule_sourceable)() {
    m.rule_source = value
}
// SetRuleType sets the rule_type property value. The type of rule.
func (m *RuleSuite_rule_evaluations) SetRuleType(value *string)() {
    m.rule_type = value
}
type RuleSuite_rule_evaluationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetails()(*string)
    GetEnforcement()(*RuleSuite_rule_evaluations_enforcement)
    GetResult()(*RuleSuite_rule_evaluations_result)
    GetRuleSource()(RuleSuite_rule_evaluations_rule_sourceable)
    GetRuleType()(*string)
    SetDetails(value *string)()
    SetEnforcement(value *RuleSuite_rule_evaluations_enforcement)()
    SetResult(value *RuleSuite_rule_evaluations_result)()
    SetRuleSource(value RuleSuite_rule_evaluations_rule_sourceable)()
    SetRuleType(value *string)()
}
