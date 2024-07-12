package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRulesetConditionsForRepositoryIDs conditions to target repositories by id and refs by name
type RepositoryRulesetConditionsForRepositoryIDs struct {
    RepositoryRulesetConditions
}
// NewRepositoryRulesetConditionsForRepositoryIDs instantiates a new RepositoryRulesetConditionsForRepositoryIDs and sets the default values.
func NewRepositoryRulesetConditionsForRepositoryIDs()(*RepositoryRulesetConditionsForRepositoryIDs) {
    m := &RepositoryRulesetConditionsForRepositoryIDs{
        RepositoryRulesetConditions: *NewRepositoryRulesetConditions(),
    }
    return m
}
// CreateRepositoryRulesetConditionsForRepositoryIDsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoryRulesetConditionsForRepositoryIDsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRulesetConditionsForRepositoryIDs(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoryRulesetConditionsForRepositoryIDs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RepositoryRulesetConditions.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RepositoryRulesetConditionsForRepositoryIDs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RepositoryRulesetConditions.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RepositoryRulesetConditionsForRepositoryIDsable 
type RepositoryRulesetConditionsForRepositoryIDsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RepositoryRulesetConditionsable
}
