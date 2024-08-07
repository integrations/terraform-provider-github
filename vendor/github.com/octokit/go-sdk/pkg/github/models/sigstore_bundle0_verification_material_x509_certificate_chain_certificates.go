package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The rawBytes property
    rawBytes *string
}
// NewSigstoreBundle0_verificationMaterial_x509CertificateChain_certificates instantiates a new SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates and sets the default values.
func NewSigstoreBundle0_verificationMaterial_x509CertificateChain_certificates()(*SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) {
    m := &SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial_x509CertificateChain_certificates(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["rawBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawBytes(val)
        }
        return nil
    }
    return res
}
// GetRawBytes gets the rawBytes property value. The rawBytes property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) GetRawBytes()(*string) {
    return m.rawBytes
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("rawBytes", m.GetRawBytes())
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
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRawBytes sets the rawBytes property value. The rawBytes property
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain_certificates) SetRawBytes(value *string)() {
    m.rawBytes = value
}
type SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRawBytes()(*string)
    SetRawBytes(value *string)()
}
