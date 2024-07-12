package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
type ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Can be `enabled` or `disabled`.
    status *string
}
// NewItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning instantiates a new ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning and sets the default values.
func NewItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning()(*ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) {
    m := &ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Can be `enabled` or `disabled`.
// returns a *string when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. Can be `enabled` or `disabled`.
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning) SetStatus(value *string)() {
    m.status = value
}
type ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*string)
    SetStatus(value *string)()
}
