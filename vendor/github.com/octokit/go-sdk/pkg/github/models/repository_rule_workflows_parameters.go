package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleWorkflows_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Workflows that must pass for this rule to pass.
    workflows []RepositoryRuleParamsWorkflowFileReferenceable
}
// NewRepositoryRuleWorkflows_parameters instantiates a new RepositoryRuleWorkflows_parameters and sets the default values.
func NewRepositoryRuleWorkflows_parameters()(*RepositoryRuleWorkflows_parameters) {
    m := &RepositoryRuleWorkflows_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleWorkflows_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleWorkflows_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleWorkflows_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleWorkflows_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleWorkflows_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryRuleParamsWorkflowFileReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryRuleParamsWorkflowFileReferenceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryRuleParamsWorkflowFileReferenceable)
                }
            }
            m.SetWorkflows(res)
        }
        return nil
    }
    return res
}
// GetWorkflows gets the workflows property value. Workflows that must pass for this rule to pass.
// returns a []RepositoryRuleParamsWorkflowFileReferenceable when successful
func (m *RepositoryRuleWorkflows_parameters) GetWorkflows()([]RepositoryRuleParamsWorkflowFileReferenceable) {
    return m.workflows
}
// Serialize serializes information the current object
func (m *RepositoryRuleWorkflows_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWorkflows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkflows()))
        for i, v := range m.GetWorkflows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("workflows", cast)
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
func (m *RepositoryRuleWorkflows_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWorkflows sets the workflows property value. Workflows that must pass for this rule to pass.
func (m *RepositoryRuleWorkflows_parameters) SetWorkflows(value []RepositoryRuleParamsWorkflowFileReferenceable)() {
    m.workflows = value
}
type RepositoryRuleWorkflows_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWorkflows()([]RepositoryRuleParamsWorkflowFileReferenceable)
    SetWorkflows(value []RepositoryRuleParamsWorkflowFileReferenceable)()
}
