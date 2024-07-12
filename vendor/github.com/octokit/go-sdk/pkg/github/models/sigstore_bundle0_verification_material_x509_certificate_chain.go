package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial_x509CertificateChain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The certificates property
    certificates []SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable
}
// NewSigstoreBundle0_verificationMaterial_x509CertificateChain instantiates a new SigstoreBundle0_verificationMaterial_x509CertificateChain and sets the default values.
func NewSigstoreBundle0_verificationMaterial_x509CertificateChain()(*SigstoreBundle0_verificationMaterial_x509CertificateChain) {
    m := &SigstoreBundle0_verificationMaterial_x509CertificateChain{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterial_x509CertificateChainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterial_x509CertificateChainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial_x509CertificateChain(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificates gets the certificates property value. The certificates property
// returns a []SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable when successful
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain) GetCertificates()([]SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable) {
    return m.certificates
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable)
                }
            }
            m.SetCertificates(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificates()))
        for i, v := range m.GetCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("certificates", cast)
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
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificates sets the certificates property value. The certificates property
func (m *SigstoreBundle0_verificationMaterial_x509CertificateChain) SetCertificates(value []SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable)() {
    m.certificates = value
}
type SigstoreBundle0_verificationMaterial_x509CertificateChainable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificates()([]SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable)
    SetCertificates(value []SigstoreBundle0_verificationMaterial_x509CertificateChain_certificatesable)()
}
