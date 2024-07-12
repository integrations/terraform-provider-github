package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleRequiredDeployments_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The environments that must be successfully deployed to before branches can be merged.
    required_deployment_environments []string
}
// NewRepositoryRuleRequiredDeployments_parameters instantiates a new RepositoryRuleRequiredDeployments_parameters and sets the default values.
func NewRepositoryRuleRequiredDeployments_parameters()(*RepositoryRuleRequiredDeployments_parameters) {
    m := &RepositoryRuleRequiredDeployments_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleRequiredDeployments_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleRequiredDeployments_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleRequiredDeployments_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleRequiredDeployments_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleRequiredDeployments_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["required_deployment_environments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRequiredDeploymentEnvironments(res)
        }
        return nil
    }
    return res
}
// GetRequiredDeploymentEnvironments gets the required_deployment_environments property value. The environments that must be successfully deployed to before branches can be merged.
// returns a []string when successful
func (m *RepositoryRuleRequiredDeployments_parameters) GetRequiredDeploymentEnvironments()([]string) {
    return m.required_deployment_environments
}
// Serialize serializes information the current object
func (m *RepositoryRuleRequiredDeployments_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRequiredDeploymentEnvironments() != nil {
        err := writer.WriteCollectionOfStringValues("required_deployment_environments", m.GetRequiredDeploymentEnvironments())
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
func (m *RepositoryRuleRequiredDeployments_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRequiredDeploymentEnvironments sets the required_deployment_environments property value. The environments that must be successfully deployed to before branches can be merged.
func (m *RepositoryRuleRequiredDeployments_parameters) SetRequiredDeploymentEnvironments(value []string)() {
    m.required_deployment_environments = value
}
type RepositoryRuleRequiredDeployments_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequiredDeploymentEnvironments()([]string)
    SetRequiredDeploymentEnvironments(value []string)()
}
