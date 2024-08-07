package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectedBranch_allow_force_pushes struct {
    // The enabled property
    enabled *bool
}
// NewProtectedBranch_allow_force_pushes instantiates a new ProtectedBranch_allow_force_pushes and sets the default values.
func NewProtectedBranch_allow_force_pushes()(*ProtectedBranch_allow_force_pushes) {
    m := &ProtectedBranch_allow_force_pushes{
    }
    return m
}
// CreateProtectedBranch_allow_force_pushesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedBranch_allow_force_pushesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranch_allow_force_pushes(), nil
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *ProtectedBranch_allow_force_pushes) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedBranch_allow_force_pushes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProtectedBranch_allow_force_pushes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnabled sets the enabled property value. The enabled property
func (m *ProtectedBranch_allow_force_pushes) SetEnabled(value *bool)() {
    m.enabled = value
}
type ProtectedBranch_allow_force_pushesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    SetEnabled(value *bool)()
}
