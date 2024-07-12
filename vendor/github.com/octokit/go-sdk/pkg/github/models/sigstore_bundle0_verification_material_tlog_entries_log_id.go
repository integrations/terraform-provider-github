package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial_tlogEntries_logId struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The keyId property
    keyId *string
}
// NewSigstoreBundle0_verificationMaterial_tlogEntries_logId instantiates a new SigstoreBundle0_verificationMaterial_tlogEntries_logId and sets the default values.
func NewSigstoreBundle0_verificationMaterial_tlogEntries_logId()(*SigstoreBundle0_verificationMaterial_tlogEntries_logId) {
    m := &SigstoreBundle0_verificationMaterial_tlogEntries_logId{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterial_tlogEntries_logIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterial_tlogEntries_logIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial_tlogEntries_logId(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_logId) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_logId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The keyId property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_logId) GetKeyId()(*string) {
    return m.keyId
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_logId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("keyId", m.GetKeyId())
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
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_logId) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyId sets the keyId property value. The keyId property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_logId) SetKeyId(value *string)() {
    m.keyId = value
}
type SigstoreBundle0_verificationMaterial_tlogEntries_logIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyId()(*string)
    SetKeyId(value *string)()
}
