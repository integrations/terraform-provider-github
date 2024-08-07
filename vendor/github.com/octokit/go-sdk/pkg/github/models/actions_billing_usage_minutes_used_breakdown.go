package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsBillingUsage_minutes_used_breakdown struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total minutes used on macOS runner machines.
    mACOS *int32
    // Total minutes used on macOS 12 core runner machines.
    macos_12_core *int32
    // Total minutes used on all runner machines.
    total *int32
    // Total minutes used on Ubuntu runner machines.
    uBUNTU *int32
    // Total minutes used on Ubuntu 16 core runner machines.
    ubuntu_16_core *int32
    // Total minutes used on Ubuntu 32 core runner machines.
    ubuntu_32_core *int32
    // Total minutes used on Ubuntu 4 core runner machines.
    ubuntu_4_core *int32
    // Total minutes used on Ubuntu 64 core runner machines.
    ubuntu_64_core *int32
    // Total minutes used on Ubuntu 8 core runner machines.
    ubuntu_8_core *int32
    // Total minutes used on Windows runner machines.
    wINDOWS *int32
    // Total minutes used on Windows 16 core runner machines.
    windows_16_core *int32
    // Total minutes used on Windows 32 core runner machines.
    windows_32_core *int32
    // Total minutes used on Windows 4 core runner machines.
    windows_4_core *int32
    // Total minutes used on Windows 64 core runner machines.
    windows_64_core *int32
    // Total minutes used on Windows 8 core runner machines.
    windows_8_core *int32
}
// NewActionsBillingUsage_minutes_used_breakdown instantiates a new ActionsBillingUsage_minutes_used_breakdown and sets the default values.
func NewActionsBillingUsage_minutes_used_breakdown()(*ActionsBillingUsage_minutes_used_breakdown) {
    m := &ActionsBillingUsage_minutes_used_breakdown{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsBillingUsage_minutes_used_breakdownFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsBillingUsage_minutes_used_breakdownFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsBillingUsage_minutes_used_breakdown(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["MACOS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMACOS(val)
        }
        return nil
    }
    res["macos_12_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacos12Core(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["UBUNTU"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUBUNTU(val)
        }
        return nil
    }
    res["ubuntu_16_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUbuntu16Core(val)
        }
        return nil
    }
    res["ubuntu_32_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUbuntu32Core(val)
        }
        return nil
    }
    res["ubuntu_4_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUbuntu4Core(val)
        }
        return nil
    }
    res["ubuntu_64_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUbuntu64Core(val)
        }
        return nil
    }
    res["ubuntu_8_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUbuntu8Core(val)
        }
        return nil
    }
    res["WINDOWS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWINDOWS(val)
        }
        return nil
    }
    res["windows_16_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows16Core(val)
        }
        return nil
    }
    res["windows_32_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows32Core(val)
        }
        return nil
    }
    res["windows_4_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows4Core(val)
        }
        return nil
    }
    res["windows_64_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows64Core(val)
        }
        return nil
    }
    res["windows_8_core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows8Core(val)
        }
        return nil
    }
    return res
}
// GetMACOS gets the MACOS property value. Total minutes used on macOS runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetMACOS()(*int32) {
    return m.mACOS
}
// GetMacos12Core gets the macos_12_core property value. Total minutes used on macOS 12 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetMacos12Core()(*int32) {
    return m.macos_12_core
}
// GetTotal gets the total property value. Total minutes used on all runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetTotal()(*int32) {
    return m.total
}
// GetUBUNTU gets the UBUNTU property value. Total minutes used on Ubuntu runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetUBUNTU()(*int32) {
    return m.uBUNTU
}
// GetUbuntu16Core gets the ubuntu_16_core property value. Total minutes used on Ubuntu 16 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetUbuntu16Core()(*int32) {
    return m.ubuntu_16_core
}
// GetUbuntu32Core gets the ubuntu_32_core property value. Total minutes used on Ubuntu 32 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetUbuntu32Core()(*int32) {
    return m.ubuntu_32_core
}
// GetUbuntu4Core gets the ubuntu_4_core property value. Total minutes used on Ubuntu 4 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetUbuntu4Core()(*int32) {
    return m.ubuntu_4_core
}
// GetUbuntu64Core gets the ubuntu_64_core property value. Total minutes used on Ubuntu 64 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetUbuntu64Core()(*int32) {
    return m.ubuntu_64_core
}
// GetUbuntu8Core gets the ubuntu_8_core property value. Total minutes used on Ubuntu 8 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetUbuntu8Core()(*int32) {
    return m.ubuntu_8_core
}
// GetWINDOWS gets the WINDOWS property value. Total minutes used on Windows runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetWINDOWS()(*int32) {
    return m.wINDOWS
}
// GetWindows16Core gets the windows_16_core property value. Total minutes used on Windows 16 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetWindows16Core()(*int32) {
    return m.windows_16_core
}
// GetWindows32Core gets the windows_32_core property value. Total minutes used on Windows 32 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetWindows32Core()(*int32) {
    return m.windows_32_core
}
// GetWindows4Core gets the windows_4_core property value. Total minutes used on Windows 4 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetWindows4Core()(*int32) {
    return m.windows_4_core
}
// GetWindows64Core gets the windows_64_core property value. Total minutes used on Windows 64 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetWindows64Core()(*int32) {
    return m.windows_64_core
}
// GetWindows8Core gets the windows_8_core property value. Total minutes used on Windows 8 core runner machines.
// returns a *int32 when successful
func (m *ActionsBillingUsage_minutes_used_breakdown) GetWindows8Core()(*int32) {
    return m.windows_8_core
}
// Serialize serializes information the current object
func (m *ActionsBillingUsage_minutes_used_breakdown) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("MACOS", m.GetMACOS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("macos_12_core", m.GetMacos12Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("UBUNTU", m.GetUBUNTU())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ubuntu_16_core", m.GetUbuntu16Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ubuntu_32_core", m.GetUbuntu32Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ubuntu_4_core", m.GetUbuntu4Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ubuntu_64_core", m.GetUbuntu64Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ubuntu_8_core", m.GetUbuntu8Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("WINDOWS", m.GetWINDOWS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows_16_core", m.GetWindows16Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows_32_core", m.GetWindows32Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows_4_core", m.GetWindows4Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows_64_core", m.GetWindows64Core())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows_8_core", m.GetWindows8Core())
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
func (m *ActionsBillingUsage_minutes_used_breakdown) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMACOS sets the MACOS property value. Total minutes used on macOS runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetMACOS(value *int32)() {
    m.mACOS = value
}
// SetMacos12Core sets the macos_12_core property value. Total minutes used on macOS 12 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetMacos12Core(value *int32)() {
    m.macos_12_core = value
}
// SetTotal sets the total property value. Total minutes used on all runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetTotal(value *int32)() {
    m.total = value
}
// SetUBUNTU sets the UBUNTU property value. Total minutes used on Ubuntu runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetUBUNTU(value *int32)() {
    m.uBUNTU = value
}
// SetUbuntu16Core sets the ubuntu_16_core property value. Total minutes used on Ubuntu 16 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetUbuntu16Core(value *int32)() {
    m.ubuntu_16_core = value
}
// SetUbuntu32Core sets the ubuntu_32_core property value. Total minutes used on Ubuntu 32 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetUbuntu32Core(value *int32)() {
    m.ubuntu_32_core = value
}
// SetUbuntu4Core sets the ubuntu_4_core property value. Total minutes used on Ubuntu 4 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetUbuntu4Core(value *int32)() {
    m.ubuntu_4_core = value
}
// SetUbuntu64Core sets the ubuntu_64_core property value. Total minutes used on Ubuntu 64 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetUbuntu64Core(value *int32)() {
    m.ubuntu_64_core = value
}
// SetUbuntu8Core sets the ubuntu_8_core property value. Total minutes used on Ubuntu 8 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetUbuntu8Core(value *int32)() {
    m.ubuntu_8_core = value
}
// SetWINDOWS sets the WINDOWS property value. Total minutes used on Windows runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetWINDOWS(value *int32)() {
    m.wINDOWS = value
}
// SetWindows16Core sets the windows_16_core property value. Total minutes used on Windows 16 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetWindows16Core(value *int32)() {
    m.windows_16_core = value
}
// SetWindows32Core sets the windows_32_core property value. Total minutes used on Windows 32 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetWindows32Core(value *int32)() {
    m.windows_32_core = value
}
// SetWindows4Core sets the windows_4_core property value. Total minutes used on Windows 4 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetWindows4Core(value *int32)() {
    m.windows_4_core = value
}
// SetWindows64Core sets the windows_64_core property value. Total minutes used on Windows 64 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetWindows64Core(value *int32)() {
    m.windows_64_core = value
}
// SetWindows8Core sets the windows_8_core property value. Total minutes used on Windows 8 core runner machines.
func (m *ActionsBillingUsage_minutes_used_breakdown) SetWindows8Core(value *int32)() {
    m.windows_8_core = value
}
type ActionsBillingUsage_minutes_used_breakdownable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMACOS()(*int32)
    GetMacos12Core()(*int32)
    GetTotal()(*int32)
    GetUBUNTU()(*int32)
    GetUbuntu16Core()(*int32)
    GetUbuntu32Core()(*int32)
    GetUbuntu4Core()(*int32)
    GetUbuntu64Core()(*int32)
    GetUbuntu8Core()(*int32)
    GetWINDOWS()(*int32)
    GetWindows16Core()(*int32)
    GetWindows32Core()(*int32)
    GetWindows4Core()(*int32)
    GetWindows64Core()(*int32)
    GetWindows8Core()(*int32)
    SetMACOS(value *int32)()
    SetMacos12Core(value *int32)()
    SetTotal(value *int32)()
    SetUBUNTU(value *int32)()
    SetUbuntu16Core(value *int32)()
    SetUbuntu32Core(value *int32)()
    SetUbuntu4Core(value *int32)()
    SetUbuntu64Core(value *int32)()
    SetUbuntu8Core(value *int32)()
    SetWINDOWS(value *int32)()
    SetWindows16Core(value *int32)()
    SetWindows32Core(value *int32)()
    SetWindows4Core(value *int32)()
    SetWindows64Core(value *int32)()
    SetWindows8Core(value *int32)()
}
