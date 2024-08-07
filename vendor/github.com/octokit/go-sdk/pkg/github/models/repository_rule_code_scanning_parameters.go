package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleCodeScanning_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Tools that must provide code scanning results for this rule to pass.
    code_scanning_tools []RepositoryRuleParamsCodeScanningToolable
}
// NewRepositoryRuleCodeScanning_parameters instantiates a new RepositoryRuleCodeScanning_parameters and sets the default values.
func NewRepositoryRuleCodeScanning_parameters()(*RepositoryRuleCodeScanning_parameters) {
    m := &RepositoryRuleCodeScanning_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleCodeScanning_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleCodeScanning_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleCodeScanning_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleCodeScanning_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodeScanningTools gets the code_scanning_tools property value. Tools that must provide code scanning results for this rule to pass.
// returns a []RepositoryRuleParamsCodeScanningToolable when successful
func (m *RepositoryRuleCodeScanning_parameters) GetCodeScanningTools()([]RepositoryRuleParamsCodeScanningToolable) {
    return m.code_scanning_tools
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleCodeScanning_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code_scanning_tools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryRuleParamsCodeScanningToolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryRuleParamsCodeScanningToolable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryRuleParamsCodeScanningToolable)
                }
            }
            m.SetCodeScanningTools(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RepositoryRuleCodeScanning_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCodeScanningTools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCodeScanningTools()))
        for i, v := range m.GetCodeScanningTools() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("code_scanning_tools", cast)
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
func (m *RepositoryRuleCodeScanning_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodeScanningTools sets the code_scanning_tools property value. Tools that must provide code scanning results for this rule to pass.
func (m *RepositoryRuleCodeScanning_parameters) SetCodeScanningTools(value []RepositoryRuleParamsCodeScanningToolable)() {
    m.code_scanning_tools = value
}
type RepositoryRuleCodeScanning_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodeScanningTools()([]RepositoryRuleParamsCodeScanningToolable)
    SetCodeScanningTools(value []RepositoryRuleParamsCodeScanningToolable)()
}
