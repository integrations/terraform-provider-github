package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRulesetConditionsForRepositoryNames conditions to target repositories by name and refs by name
type RepositoryRulesetConditionsForRepositoryNames struct {
    RepositoryRulesetConditions
}
// NewRepositoryRulesetConditionsForRepositoryNames instantiates a new RepositoryRulesetConditionsForRepositoryNames and sets the default values.
func NewRepositoryRulesetConditionsForRepositoryNames()(*RepositoryRulesetConditionsForRepositoryNames) {
    m := &RepositoryRulesetConditionsForRepositoryNames{
        RepositoryRulesetConditions: *NewRepositoryRulesetConditions(),
    }
    return m
}
// CreateRepositoryRulesetConditionsForRepositoryNamesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoryRulesetConditionsForRepositoryNamesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRulesetConditionsForRepositoryNames(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoryRulesetConditionsForRepositoryNames) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RepositoryRulesetConditions.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RepositoryRulesetConditionsForRepositoryNames) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RepositoryRulesetConditions.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RepositoryRulesetConditionsForRepositoryNamesable 
type RepositoryRulesetConditionsForRepositoryNamesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RepositoryRulesetConditionsable
}
