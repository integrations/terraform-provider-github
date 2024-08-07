package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_dsseEnvelope_signatures struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The keyid property
    keyid *string
    // The sig property
    sig *string
}
// NewSigstoreBundle0_dsseEnvelope_signatures instantiates a new SigstoreBundle0_dsseEnvelope_signatures and sets the default values.
func NewSigstoreBundle0_dsseEnvelope_signatures()(*SigstoreBundle0_dsseEnvelope_signatures) {
    m := &SigstoreBundle0_dsseEnvelope_signatures{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_dsseEnvelope_signaturesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_dsseEnvelope_signaturesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_dsseEnvelope_signatures(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_dsseEnvelope_signatures) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_dsseEnvelope_signatures) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyid(val)
        }
        return nil
    }
    res["sig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSig(val)
        }
        return nil
    }
    return res
}
// GetKeyid gets the keyid property value. The keyid property
// returns a *string when successful
func (m *SigstoreBundle0_dsseEnvelope_signatures) GetKeyid()(*string) {
    return m.keyid
}
// GetSig gets the sig property value. The sig property
// returns a *string when successful
func (m *SigstoreBundle0_dsseEnvelope_signatures) GetSig()(*string) {
    return m.sig
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_dsseEnvelope_signatures) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("keyid", m.GetKeyid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sig", m.GetSig())
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
func (m *SigstoreBundle0_dsseEnvelope_signatures) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyid sets the keyid property value. The keyid property
func (m *SigstoreBundle0_dsseEnvelope_signatures) SetKeyid(value *string)() {
    m.keyid = value
}
// SetSig sets the sig property value. The sig property
func (m *SigstoreBundle0_dsseEnvelope_signatures) SetSig(value *string)() {
    m.sig = value
}
type SigstoreBundle0_dsseEnvelope_signaturesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyid()(*string)
    GetSig()(*string)
    SetKeyid(value *string)()
    SetSig(value *string)()
}
