package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableCodespaceMachine a description of the machine powering a codespace.
type NullableCodespaceMachine struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How many cores are available to the codespace.
    cpus *int32
    // The display name of the machine includes cores, memory, and storage.
    display_name *string
    // How much memory is available to the codespace.
    memory_in_bytes *int32
    // The name of the machine.
    name *string
    // The operating system of the machine.
    operating_system *string
    // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
    prebuild_availability *NullableCodespaceMachine_prebuild_availability
    // How much storage is available to the codespace.
    storage_in_bytes *int32
}
// NewNullableCodespaceMachine instantiates a new NullableCodespaceMachine and sets the default values.
func NewNullableCodespaceMachine()(*NullableCodespaceMachine) {
    m := &NullableCodespaceMachine{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableCodespaceMachineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableCodespaceMachineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableCodespaceMachine(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableCodespaceMachine) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpus gets the cpus property value. How many cores are available to the codespace.
// returns a *int32 when successful
func (m *NullableCodespaceMachine) GetCpus()(*int32) {
    return m.cpus
}
// GetDisplayName gets the display_name property value. The display name of the machine includes cores, memory, and storage.
// returns a *string when successful
func (m *NullableCodespaceMachine) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableCodespaceMachine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpus(val)
        }
        return nil
    }
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["memory_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryInBytes(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["operating_system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["prebuild_availability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNullableCodespaceMachine_prebuild_availability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebuildAvailability(val.(*NullableCodespaceMachine_prebuild_availability))
        }
        return nil
    }
    res["storage_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageInBytes(val)
        }
        return nil
    }
    return res
}
// GetMemoryInBytes gets the memory_in_bytes property value. How much memory is available to the codespace.
// returns a *int32 when successful
func (m *NullableCodespaceMachine) GetMemoryInBytes()(*int32) {
    return m.memory_in_bytes
}
// GetName gets the name property value. The name of the machine.
// returns a *string when successful
func (m *NullableCodespaceMachine) GetName()(*string) {
    return m.name
}
// GetOperatingSystem gets the operating_system property value. The operating system of the machine.
// returns a *string when successful
func (m *NullableCodespaceMachine) GetOperatingSystem()(*string) {
    return m.operating_system
}
// GetPrebuildAvailability gets the prebuild_availability property value. Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
// returns a *NullableCodespaceMachine_prebuild_availability when successful
func (m *NullableCodespaceMachine) GetPrebuildAvailability()(*NullableCodespaceMachine_prebuild_availability) {
    return m.prebuild_availability
}
// GetStorageInBytes gets the storage_in_bytes property value. How much storage is available to the codespace.
// returns a *int32 when successful
func (m *NullableCodespaceMachine) GetStorageInBytes()(*int32) {
    return m.storage_in_bytes
}
// Serialize serializes information the current object
func (m *NullableCodespaceMachine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("cpus", m.GetCpus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("memory_in_bytes", m.GetMemoryInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operating_system", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    if m.GetPrebuildAvailability() != nil {
        cast := (*m.GetPrebuildAvailability()).String()
        err := writer.WriteStringValue("prebuild_availability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("storage_in_bytes", m.GetStorageInBytes())
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
func (m *NullableCodespaceMachine) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpus sets the cpus property value. How many cores are available to the codespace.
func (m *NullableCodespaceMachine) SetCpus(value *int32)() {
    m.cpus = value
}
// SetDisplayName sets the display_name property value. The display name of the machine includes cores, memory, and storage.
func (m *NullableCodespaceMachine) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetMemoryInBytes sets the memory_in_bytes property value. How much memory is available to the codespace.
func (m *NullableCodespaceMachine) SetMemoryInBytes(value *int32)() {
    m.memory_in_bytes = value
}
// SetName sets the name property value. The name of the machine.
func (m *NullableCodespaceMachine) SetName(value *string)() {
    m.name = value
}
// SetOperatingSystem sets the operating_system property value. The operating system of the machine.
func (m *NullableCodespaceMachine) SetOperatingSystem(value *string)() {
    m.operating_system = value
}
// SetPrebuildAvailability sets the prebuild_availability property value. Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
func (m *NullableCodespaceMachine) SetPrebuildAvailability(value *NullableCodespaceMachine_prebuild_availability)() {
    m.prebuild_availability = value
}
// SetStorageInBytes sets the storage_in_bytes property value. How much storage is available to the codespace.
func (m *NullableCodespaceMachine) SetStorageInBytes(value *int32)() {
    m.storage_in_bytes = value
}
type NullableCodespaceMachineable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpus()(*int32)
    GetDisplayName()(*string)
    GetMemoryInBytes()(*int32)
    GetName()(*string)
    GetOperatingSystem()(*string)
    GetPrebuildAvailability()(*NullableCodespaceMachine_prebuild_availability)
    GetStorageInBytes()(*int32)
    SetCpus(value *int32)()
    SetDisplayName(value *string)()
    SetMemoryInBytes(value *int32)()
    SetName(value *string)()
    SetOperatingSystem(value *string)()
    SetPrebuildAvailability(value *NullableCodespaceMachine_prebuild_availability)()
    SetStorageInBytes(value *int32)()
}
