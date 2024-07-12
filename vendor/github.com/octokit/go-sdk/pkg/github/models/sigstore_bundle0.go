package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SigstoreBundle0 sigstore Bundle v0.1
type SigstoreBundle0 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dsseEnvelope property
    dsseEnvelope SigstoreBundle0_dsseEnvelopeable
    // The mediaType property
    mediaType *string
    // The verificationMaterial property
    verificationMaterial SigstoreBundle0_verificationMaterialable
}
// NewSigstoreBundle0 instantiates a new SigstoreBundle0 and sets the default values.
func NewSigstoreBundle0()(*SigstoreBundle0) {
    m := &SigstoreBundle0{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDsseEnvelope gets the dsseEnvelope property value. The dsseEnvelope property
// returns a SigstoreBundle0_dsseEnvelopeable when successful
func (m *SigstoreBundle0) GetDsseEnvelope()(SigstoreBundle0_dsseEnvelopeable) {
    return m.dsseEnvelope
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dsseEnvelope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigstoreBundle0_dsseEnvelopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDsseEnvelope(val.(SigstoreBundle0_dsseEnvelopeable))
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["verificationMaterial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigstoreBundle0_verificationMaterialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerificationMaterial(val.(SigstoreBundle0_verificationMaterialable))
        }
        return nil
    }
    return res
}
// GetMediaType gets the mediaType property value. The mediaType property
// returns a *string when successful
func (m *SigstoreBundle0) GetMediaType()(*string) {
    return m.mediaType
}
// GetVerificationMaterial gets the verificationMaterial property value. The verificationMaterial property
// returns a SigstoreBundle0_verificationMaterialable when successful
func (m *SigstoreBundle0) GetVerificationMaterial()(SigstoreBundle0_verificationMaterialable) {
    return m.verificationMaterial
}
// Serialize serializes information the current object
func (m *SigstoreBundle0) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("dsseEnvelope", m.GetDsseEnvelope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verificationMaterial", m.GetVerificationMaterial())
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
func (m *SigstoreBundle0) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDsseEnvelope sets the dsseEnvelope property value. The dsseEnvelope property
func (m *SigstoreBundle0) SetDsseEnvelope(value SigstoreBundle0_dsseEnvelopeable)() {
    m.dsseEnvelope = value
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *SigstoreBundle0) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetVerificationMaterial sets the verificationMaterial property value. The verificationMaterial property
func (m *SigstoreBundle0) SetVerificationMaterial(value SigstoreBundle0_verificationMaterialable)() {
    m.verificationMaterial = value
}
type SigstoreBundle0able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDsseEnvelope()(SigstoreBundle0_dsseEnvelopeable)
    GetMediaType()(*string)
    GetVerificationMaterial()(SigstoreBundle0_verificationMaterialable)
    SetDsseEnvelope(value SigstoreBundle0_dsseEnvelopeable)()
    SetMediaType(value *string)()
    SetVerificationMaterial(value SigstoreBundle0_verificationMaterialable)()
}
