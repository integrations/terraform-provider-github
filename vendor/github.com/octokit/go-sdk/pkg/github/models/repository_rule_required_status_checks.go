package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleRequiredStatusChecks choose which status checks must pass before the ref is updated. When enabled, commits must first be pushed to another ref where the checks pass.
type RepositoryRuleRequiredStatusChecks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleRequiredStatusChecks_parametersable
    // The type property
    typeEscaped *RepositoryRuleRequiredStatusChecks_type
}
// NewRepositoryRuleRequiredStatusChecks instantiates a new RepositoryRuleRequiredStatusChecks and sets the default values.
func NewRepositoryRuleRequiredStatusChecks()(*RepositoryRuleRequiredStatusChecks) {
    m := &RepositoryRuleRequiredStatusChecks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleRequiredStatusChecksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleRequiredStatusChecksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleRequiredStatusChecks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleRequiredStatusChecks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleRequiredStatusChecks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleRequiredStatusChecks_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleRequiredStatusChecks_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleRequiredStatusChecks_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleRequiredStatusChecks_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleRequiredStatusChecks_parametersable when successful
func (m *RepositoryRuleRequiredStatusChecks) GetParameters()(RepositoryRuleRequiredStatusChecks_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleRequiredStatusChecks_type when successful
func (m *RepositoryRuleRequiredStatusChecks) GetTypeEscaped()(*RepositoryRuleRequiredStatusChecks_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleRequiredStatusChecks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parameters", m.GetParameters())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryRuleRequiredStatusChecks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleRequiredStatusChecks) SetParameters(value RepositoryRuleRequiredStatusChecks_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleRequiredStatusChecks) SetTypeEscaped(value *RepositoryRuleRequiredStatusChecks_type)() {
    m.typeEscaped = value
}
type RepositoryRuleRequiredStatusChecksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleRequiredStatusChecks_parametersable)
    GetTypeEscaped()(*RepositoryRuleRequiredStatusChecks_type)
    SetParameters(value RepositoryRuleRequiredStatusChecks_parametersable)()
    SetTypeEscaped(value *RepositoryRuleRequiredStatusChecks_type)()
}
