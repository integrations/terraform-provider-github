package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleRequiredDeployments choose which environments must be successfully deployed to before refs can be pushed into a ref that matches this rule.
type RepositoryRuleRequiredDeployments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleRequiredDeployments_parametersable
    // The type property
    typeEscaped *RepositoryRuleRequiredDeployments_type
}
// NewRepositoryRuleRequiredDeployments instantiates a new RepositoryRuleRequiredDeployments and sets the default values.
func NewRepositoryRuleRequiredDeployments()(*RepositoryRuleRequiredDeployments) {
    m := &RepositoryRuleRequiredDeployments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleRequiredDeploymentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleRequiredDeploymentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleRequiredDeployments(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleRequiredDeployments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleRequiredDeployments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleRequiredDeployments_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleRequiredDeployments_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleRequiredDeployments_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleRequiredDeployments_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleRequiredDeployments_parametersable when successful
func (m *RepositoryRuleRequiredDeployments) GetParameters()(RepositoryRuleRequiredDeployments_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleRequiredDeployments_type when successful
func (m *RepositoryRuleRequiredDeployments) GetTypeEscaped()(*RepositoryRuleRequiredDeployments_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleRequiredDeployments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RepositoryRuleRequiredDeployments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleRequiredDeployments) SetParameters(value RepositoryRuleRequiredDeployments_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleRequiredDeployments) SetTypeEscaped(value *RepositoryRuleRequiredDeployments_type)() {
    m.typeEscaped = value
}
type RepositoryRuleRequiredDeploymentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleRequiredDeployments_parametersable)
    GetTypeEscaped()(*RepositoryRuleRequiredDeployments_type)
    SetParameters(value RepositoryRuleRequiredDeployments_parametersable)()
    SetTypeEscaped(value *RepositoryRuleRequiredDeployments_type)()
}
