package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleTagNamePattern_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How this rule will appear to users.
    name *string
    // If true, the rule will fail if the pattern matches.
    negate *bool
    // The operator to use for matching.
    operator *RepositoryRuleTagNamePattern_parameters_operator
    // The pattern to match with.
    pattern *string
}
// NewRepositoryRuleTagNamePattern_parameters instantiates a new RepositoryRuleTagNamePattern_parameters and sets the default values.
func NewRepositoryRuleTagNamePattern_parameters()(*RepositoryRuleTagNamePattern_parameters) {
    m := &RepositoryRuleTagNamePattern_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleTagNamePattern_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleTagNamePattern_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleTagNamePattern_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleTagNamePattern_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleTagNamePattern_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["negate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNegate(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleTagNamePattern_parameters_operator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*RepositoryRuleTagNamePattern_parameters_operator))
        }
        return nil
    }
    res["pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. How this rule will appear to users.
// returns a *string when successful
func (m *RepositoryRuleTagNamePattern_parameters) GetName()(*string) {
    return m.name
}
// GetNegate gets the negate property value. If true, the rule will fail if the pattern matches.
// returns a *bool when successful
func (m *RepositoryRuleTagNamePattern_parameters) GetNegate()(*bool) {
    return m.negate
}
// GetOperator gets the operator property value. The operator to use for matching.
// returns a *RepositoryRuleTagNamePattern_parameters_operator when successful
func (m *RepositoryRuleTagNamePattern_parameters) GetOperator()(*RepositoryRuleTagNamePattern_parameters_operator) {
    return m.operator
}
// GetPattern gets the pattern property value. The pattern to match with.
// returns a *string when successful
func (m *RepositoryRuleTagNamePattern_parameters) GetPattern()(*string) {
    return m.pattern
}
// Serialize serializes information the current object
func (m *RepositoryRuleTagNamePattern_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("negate", m.GetNegate())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pattern", m.GetPattern())
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
func (m *RepositoryRuleTagNamePattern_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. How this rule will appear to users.
func (m *RepositoryRuleTagNamePattern_parameters) SetName(value *string)() {
    m.name = value
}
// SetNegate sets the negate property value. If true, the rule will fail if the pattern matches.
func (m *RepositoryRuleTagNamePattern_parameters) SetNegate(value *bool)() {
    m.negate = value
}
// SetOperator sets the operator property value. The operator to use for matching.
func (m *RepositoryRuleTagNamePattern_parameters) SetOperator(value *RepositoryRuleTagNamePattern_parameters_operator)() {
    m.operator = value
}
// SetPattern sets the pattern property value. The pattern to match with.
func (m *RepositoryRuleTagNamePattern_parameters) SetPattern(value *string)() {
    m.pattern = value
}
type RepositoryRuleTagNamePattern_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNegate()(*bool)
    GetOperator()(*RepositoryRuleTagNamePattern_parameters_operator)
    GetPattern()(*string)
    SetName(value *string)()
    SetNegate(value *bool)()
    SetOperator(value *RepositoryRuleTagNamePattern_parameters_operator)()
    SetPattern(value *string)()
}
