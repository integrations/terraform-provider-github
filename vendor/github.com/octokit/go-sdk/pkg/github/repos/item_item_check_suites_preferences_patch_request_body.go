package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCheckSuitesPreferencesPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Enables or disables automatic creation of CheckSuite events upon pushes to the repository. Enabled by default.
    auto_trigger_checks []ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable
}
// NewItemItemCheckSuitesPreferencesPatchRequestBody instantiates a new ItemItemCheckSuitesPreferencesPatchRequestBody and sets the default values.
func NewItemItemCheckSuitesPreferencesPatchRequestBody()(*ItemItemCheckSuitesPreferencesPatchRequestBody) {
    m := &ItemItemCheckSuitesPreferencesPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCheckSuitesPreferencesPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCheckSuitesPreferencesPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCheckSuitesPreferencesPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutoTriggerChecks gets the auto_trigger_checks property value. Enables or disables automatic creation of CheckSuite events upon pushes to the repository. Enabled by default.
// returns a []ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable when successful
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody) GetAutoTriggerChecks()([]ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable) {
    return m.auto_trigger_checks
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auto_trigger_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable)
                }
            }
            m.SetAutoTriggerChecks(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutoTriggerChecks sets the auto_trigger_checks property value. Enables or disables automatic creation of CheckSuite events upon pushes to the repository. Enabled by default.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody) SetAutoTriggerChecks(value []ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable)() {
    m.auto_trigger_checks = value
}
type ItemItemCheckSuitesPreferencesPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoTriggerChecks()([]ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable)
    SetAutoTriggerChecks(value []ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksable)()
}
