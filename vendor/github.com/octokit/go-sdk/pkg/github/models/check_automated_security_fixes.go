package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckAutomatedSecurityFixes check Automated Security Fixes
type CheckAutomatedSecurityFixes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether automated security fixes are enabled for the repository.
    enabled *bool
    // Whether automated security fixes are paused for the repository.
    paused *bool
}
// NewCheckAutomatedSecurityFixes instantiates a new CheckAutomatedSecurityFixes and sets the default values.
func NewCheckAutomatedSecurityFixes()(*CheckAutomatedSecurityFixes) {
    m := &CheckAutomatedSecurityFixes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckAutomatedSecurityFixesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckAutomatedSecurityFixesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckAutomatedSecurityFixes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckAutomatedSecurityFixes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. Whether automated security fixes are enabled for the repository.
// returns a *bool when successful
func (m *CheckAutomatedSecurityFixes) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckAutomatedSecurityFixes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["paused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaused(val)
        }
        return nil
    }
    return res
}
// GetPaused gets the paused property value. Whether automated security fixes are paused for the repository.
// returns a *bool when successful
func (m *CheckAutomatedSecurityFixes) GetPaused()(*bool) {
    return m.paused
}
// Serialize serializes information the current object
func (m *CheckAutomatedSecurityFixes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("paused", m.GetPaused())
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
func (m *CheckAutomatedSecurityFixes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. Whether automated security fixes are enabled for the repository.
func (m *CheckAutomatedSecurityFixes) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetPaused sets the paused property value. Whether automated security fixes are paused for the repository.
func (m *CheckAutomatedSecurityFixes) SetPaused(value *bool)() {
    m.paused = value
}
type CheckAutomatedSecurityFixesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetPaused()(*bool)
    SetEnabled(value *bool)()
    SetPaused(value *bool)()
}
