package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CheckSuitePreference_preferences struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The auto_trigger_checks property
    auto_trigger_checks []CheckSuitePreference_preferences_auto_trigger_checksable
}
// NewCheckSuitePreference_preferences instantiates a new CheckSuitePreference_preferences and sets the default values.
func NewCheckSuitePreference_preferences()(*CheckSuitePreference_preferences) {
    m := &CheckSuitePreference_preferences{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckSuitePreference_preferencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckSuitePreference_preferencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckSuitePreference_preferences(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckSuitePreference_preferences) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutoTriggerChecks gets the auto_trigger_checks property value. The auto_trigger_checks property
// returns a []CheckSuitePreference_preferences_auto_trigger_checksable when successful
func (m *CheckSuitePreference_preferences) GetAutoTriggerChecks()([]CheckSuitePreference_preferences_auto_trigger_checksable) {
    return m.auto_trigger_checks
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckSuitePreference_preferences) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auto_trigger_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCheckSuitePreference_preferences_auto_trigger_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CheckSuitePreference_preferences_auto_trigger_checksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CheckSuitePreference_preferences_auto_trigger_checksable)
                }
            }
            m.SetAutoTriggerChecks(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CheckSuitePreference_preferences) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAutoTriggerChecks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAutoTriggerChecks()))
        for i, v := range m.GetAutoTriggerChecks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("auto_trigger_checks", cast)
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
func (m *CheckSuitePreference_preferences) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutoTriggerChecks sets the auto_trigger_checks property value. The auto_trigger_checks property
func (m *CheckSuitePreference_preferences) SetAutoTriggerChecks(value []CheckSuitePreference_preferences_auto_trigger_checksable)() {
    m.auto_trigger_checks = value
}
type CheckSuitePreference_preferencesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoTriggerChecks()([]CheckSuitePreference_preferences_auto_trigger_checksable)
    SetAutoTriggerChecks(value []CheckSuitePreference_preferences_auto_trigger_checksable)()
}
