package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The timestampVerificationData property
    timestampVerificationData *string
    // The tlogEntries property
    tlogEntries []SigstoreBundle0_verificationMaterial_tlogEntriesable
    // The x509CertificateChain property
    x509CertificateChain SigstoreBundle0_verificationMaterial_x509CertificateChainable
}
// NewSigstoreBundle0_verificationMaterial instantiates a new SigstoreBundle0_verificationMaterial and sets the default values.
func NewSigstoreBundle0_verificationMaterial()(*SigstoreBundle0_verificationMaterial) {
    m := &SigstoreBundle0_verificationMaterial{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["timestampVerificationData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestampVerificationData(val)
        }
        return nil
    }
    res["tlogEntries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSigstoreBundle0_verificationMaterial_tlogEntriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SigstoreBundle0_verificationMaterial_tlogEntriesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SigstoreBundle0_verificationMaterial_tlogEntriesable)
                }
            }
            m.SetTlogEntries(res)
        }
        return nil
    }
    res["x509CertificateChain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigstoreBundle0_verificationMaterial_x509CertificateChainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateChain(val.(SigstoreBundle0_verificationMaterial_x509CertificateChainable))
        }
        return nil
    }
    return res
}
// GetTimestampVerificationData gets the timestampVerificationData property value. The timestampVerificationData property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial) GetTimestampVerificationData()(*string) {
    return m.timestampVerificationData
}
// GetTlogEntries gets the tlogEntries property value. The tlogEntries property
// returns a []SigstoreBundle0_verificationMaterial_tlogEntriesable when successful
func (m *SigstoreBundle0_verificationMaterial) GetTlogEntries()([]SigstoreBundle0_verificationMaterial_tlogEntriesable) {
    return m.tlogEntries
}
// GetX509CertificateChain gets the x509CertificateChain property value. The x509CertificateChain property
// returns a SigstoreBundle0_verificationMaterial_x509CertificateChainable when successful
func (m *SigstoreBundle0_verificationMaterial) GetX509CertificateChain()(SigstoreBundle0_verificationMaterial_x509CertificateChainable) {
    return m.x509CertificateChain
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("timestampVerificationData", m.GetTimestampVerificationData())
        if err != nil {
            return err
        }
    }
    if m.GetTlogEntries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTlogEntries()))
        for i, v := range m.GetTlogEntries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("tlogEntries", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("x509CertificateChain", m.GetX509CertificateChain())
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
func (m *SigstoreBundle0_verificationMaterial) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTimestampVerificationData sets the timestampVerificationData property value. The timestampVerificationData property
func (m *SigstoreBundle0_verificationMaterial) SetTimestampVerificationData(value *string)() {
    m.timestampVerificationData = value
}
// SetTlogEntries sets the tlogEntries property value. The tlogEntries property
func (m *SigstoreBundle0_verificationMaterial) SetTlogEntries(value []SigstoreBundle0_verificationMaterial_tlogEntriesable)() {
    m.tlogEntries = value
}
// SetX509CertificateChain sets the x509CertificateChain property value. The x509CertificateChain property
func (m *SigstoreBundle0_verificationMaterial) SetX509CertificateChain(value SigstoreBundle0_verificationMaterial_x509CertificateChainable)() {
    m.x509CertificateChain = value
}
type SigstoreBundle0_verificationMaterialable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTimestampVerificationData()(*string)
    GetTlogEntries()([]SigstoreBundle0_verificationMaterial_tlogEntriesable)
    GetX509CertificateChain()(SigstoreBundle0_verificationMaterial_x509CertificateChainable)
    SetTimestampVerificationData(value *string)()
    SetTlogEntries(value []SigstoreBundle0_verificationMaterial_tlogEntriesable)()
    SetX509CertificateChain(value SigstoreBundle0_verificationMaterial_x509CertificateChainable)()
}
