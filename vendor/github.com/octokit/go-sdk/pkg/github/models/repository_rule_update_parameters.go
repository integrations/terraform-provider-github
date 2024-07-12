package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleUpdate_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Branch can pull changes from its upstream repository
    update_allows_fetch_and_merge *bool
}
// NewRepositoryRuleUpdate_parameters instantiates a new RepositoryRuleUpdate_parameters and sets the default values.
func NewRepositoryRuleUpdate_parameters()(*RepositoryRuleUpdate_parameters) {
    m := &RepositoryRuleUpdate_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleUpdate_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleUpdate_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleUpdate_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleUpdate_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleUpdate_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["update_allows_fetch_and_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateAllowsFetchAndMerge(val)
        }
        return nil
    }
    return res
}
// GetUpdateAllowsFetchAndMerge gets the update_allows_fetch_and_merge property value. Branch can pull changes from its upstream repository
// returns a *bool when successful
func (m *RepositoryRuleUpdate_parameters) GetUpdateAllowsFetchAndMerge()(*bool) {
    return m.update_allows_fetch_and_merge
}
// Serialize serializes information the current object
func (m *RepositoryRuleUpdate_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("update_allows_fetch_and_merge", m.GetUpdateAllowsFetchAndMerge())
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
func (m *RepositoryRuleUpdate_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUpdateAllowsFetchAndMerge sets the update_allows_fetch_and_merge property value. Branch can pull changes from its upstream repository
func (m *RepositoryRuleUpdate_parameters) SetUpdateAllowsFetchAndMerge(value *bool)() {
    m.update_allows_fetch_and_merge = value
}
type RepositoryRuleUpdate_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUpdateAllowsFetchAndMerge()(*bool)
    SetUpdateAllowsFetchAndMerge(value *bool)()
}
