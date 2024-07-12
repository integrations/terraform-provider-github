package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NullableCollaborator_permissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The admin property
    admin *bool
    // The maintain property
    maintain *bool
    // The pull property
    pull *bool
    // The push property
    push *bool
    // The triage property
    triage *bool
}
// NewNullableCollaborator_permissions instantiates a new NullableCollaborator_permissions and sets the default values.
func NewNullableCollaborator_permissions()(*NullableCollaborator_permissions) {
    m := &NullableCollaborator_permissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableCollaborator_permissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableCollaborator_permissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableCollaborator_permissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableCollaborator_permissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdmin gets the admin property value. The admin property
// returns a *bool when successful
func (m *NullableCollaborator_permissions) GetAdmin()(*bool) {
    return m.admin
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableCollaborator_permissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["admin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmin(val)
        }
        return nil
    }
    res["maintain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintain(val)
        }
        return nil
    }
    res["pull"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPull(val)
        }
        return nil
    }
    res["push"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPush(val)
        }
        return nil
    }
    res["triage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTriage(val)
        }
        return nil
    }
    return res
}
// GetMaintain gets the maintain property value. The maintain property
// returns a *bool when successful
func (m *NullableCollaborator_permissions) GetMaintain()(*bool) {
    return m.maintain
}
// GetPull gets the pull property value. The pull property
// returns a *bool when successful
func (m *NullableCollaborator_permissions) GetPull()(*bool) {
    return m.pull
}
// GetPush gets the push property value. The push property
// returns a *bool when successful
func (m *NullableCollaborator_permissions) GetPush()(*bool) {
    return m.push
}
// GetTriage gets the triage property value. The triage property
// returns a *bool when successful
func (m *NullableCollaborator_permissions) GetTriage()(*bool) {
    return m.triage
}
// Serialize serializes information the current object
func (m *NullableCollaborator_permissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("admin", m.GetAdmin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maintain", m.GetMaintain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("pull", m.GetPull())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("push", m.GetPush())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("triage", m.GetTriage())
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
func (m *NullableCollaborator_permissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdmin sets the admin property value. The admin property
func (m *NullableCollaborator_permissions) SetAdmin(value *bool)() {
    m.admin = value
}
// SetMaintain sets the maintain property value. The maintain property
func (m *NullableCollaborator_permissions) SetMaintain(value *bool)() {
    m.maintain = value
}
// SetPull sets the pull property value. The pull property
func (m *NullableCollaborator_permissions) SetPull(value *bool)() {
    m.pull = value
}
// SetPush sets the push property value. The push property
func (m *NullableCollaborator_permissions) SetPush(value *bool)() {
    m.push = value
}
// SetTriage sets the triage property value. The triage property
func (m *NullableCollaborator_permissions) SetTriage(value *bool)() {
    m.triage = value
}
type NullableCollaborator_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdmin()(*bool)
    GetMaintain()(*bool)
    GetPull()(*bool)
    GetPush()(*bool)
    GetTriage()(*bool)
    SetAdmin(value *bool)()
    SetMaintain(value *bool)()
    SetPull(value *bool)()
    SetPush(value *bool)()
    SetTriage(value *bool)()
}
