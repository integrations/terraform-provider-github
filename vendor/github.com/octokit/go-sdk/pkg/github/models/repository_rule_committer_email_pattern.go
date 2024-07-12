package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleCommitterEmailPattern parameters to be used for the committer_email_pattern rule
type RepositoryRuleCommitterEmailPattern struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleCommitterEmailPattern_parametersable
    // The type property
    typeEscaped *RepositoryRuleCommitterEmailPattern_type
}
// NewRepositoryRuleCommitterEmailPattern instantiates a new RepositoryRuleCommitterEmailPattern and sets the default values.
func NewRepositoryRuleCommitterEmailPattern()(*RepositoryRuleCommitterEmailPattern) {
    m := &RepositoryRuleCommitterEmailPattern{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleCommitterEmailPatternFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleCommitterEmailPatternFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleCommitterEmailPattern(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleCommitterEmailPattern) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleCommitterEmailPattern) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleCommitterEmailPattern_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleCommitterEmailPattern_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleCommitterEmailPattern_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleCommitterEmailPattern_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleCommitterEmailPattern_parametersable when successful
func (m *RepositoryRuleCommitterEmailPattern) GetParameters()(RepositoryRuleCommitterEmailPattern_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleCommitterEmailPattern_type when successful
func (m *RepositoryRuleCommitterEmailPattern) GetTypeEscaped()(*RepositoryRuleCommitterEmailPattern_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleCommitterEmailPattern) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RepositoryRuleCommitterEmailPattern) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleCommitterEmailPattern) SetParameters(value RepositoryRuleCommitterEmailPattern_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleCommitterEmailPattern) SetTypeEscaped(value *RepositoryRuleCommitterEmailPattern_type)() {
    m.typeEscaped = value
}
type RepositoryRuleCommitterEmailPatternable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleCommitterEmailPattern_parametersable)
    GetTypeEscaped()(*RepositoryRuleCommitterEmailPattern_type)
    SetParameters(value RepositoryRuleCommitterEmailPattern_parametersable)()
    SetTypeEscaped(value *RepositoryRuleCommitterEmailPattern_type)()
}
