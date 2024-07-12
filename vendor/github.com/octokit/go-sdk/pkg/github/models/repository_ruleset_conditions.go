package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRulesetConditions parameters for a repository ruleset ref name condition
type RepositoryRulesetConditions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ref_name property
    ref_name RepositoryRulesetConditions_ref_nameable
}
// NewRepositoryRulesetConditions instantiates a new RepositoryRulesetConditions and sets the default values.
func NewRepositoryRulesetConditions()(*RepositoryRulesetConditions) {
    m := &RepositoryRulesetConditions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRulesetConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRulesetConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRulesetConditions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRulesetConditions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRulesetConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ref_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRulesetConditions_ref_nameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRefName(val.(RepositoryRulesetConditions_ref_nameable))
        }
        return nil
    }
    return res
}
// GetRefName gets the ref_name property value. The ref_name property
// returns a RepositoryRulesetConditions_ref_nameable when successful
func (m *RepositoryRulesetConditions) GetRefName()(RepositoryRulesetConditions_ref_nameable) {
    return m.ref_name
}
// Serialize serializes information the current object
func (m *RepositoryRulesetConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ref_name", m.GetRefName())
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
func (m *RepositoryRulesetConditions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRefName sets the ref_name property value. The ref_name property
func (m *RepositoryRulesetConditions) SetRefName(value RepositoryRulesetConditions_ref_nameable)() {
    m.ref_name = value
}
type RepositoryRulesetConditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRefName()(RepositoryRulesetConditions_ref_nameable)
    SetRefName(value RepositoryRulesetConditions_ref_nameable)()
}
