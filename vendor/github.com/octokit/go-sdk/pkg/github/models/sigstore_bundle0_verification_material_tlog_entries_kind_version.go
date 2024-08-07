package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The kind property
    kind *string
    // The version property
    version *string
}
// NewSigstoreBundle0_verificationMaterial_tlogEntries_kindVersion instantiates a new SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion and sets the default values.
func NewSigstoreBundle0_verificationMaterial_tlogEntries_kindVersion()(*SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) {
    m := &SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterial_tlogEntries_kindVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterial_tlogEntries_kindVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial_tlogEntries_kindVersion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetKind gets the kind property value. The kind property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) GetKind()(*string) {
    return m.kind
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKind sets the kind property value. The kind property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) SetKind(value *string)() {
    m.kind = value
}
// SetVersion sets the version property value. The version property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries_kindVersion) SetVersion(value *string)() {
    m.version = value
}
type SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKind()(*string)
    GetVersion()(*string)
    SetKind(value *string)()
    SetVersion(value *string)()
}
