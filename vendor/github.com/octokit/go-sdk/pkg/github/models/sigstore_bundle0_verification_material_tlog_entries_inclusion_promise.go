package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The signedEntryTimestamp property
    signedEntryTimestamp *string
}
// NewSigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise instantiates a new SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise and sets the default values.
func NewSigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise()(*SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) {
    m := &SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["signedEntryTimestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedEntryTimestamp(val)
        }
        return nil
    }
    return res
}
// GetSignedEntryTimestamp gets the signedEntryTimestamp property value. The signedEntryTimestamp property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) GetSignedEntryTimestamp()(*string) {
    return m.signedEntryTimestamp
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("signedEntryTimestamp", m.GetSignedEntryTimestamp())
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
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSignedEntryTimestamp sets the signedEntryTimestamp property value. The signedEntryTimestamp property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromise) SetSignedEntryTimestamp(value *string)() {
    m.signedEntryTimestamp = value
}
type SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSignedEntryTimestamp()(*string)
    SetSignedEntryTimestamp(value *string)()
}
