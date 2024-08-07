package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_dsseEnvelope struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The payload property
    payload *string
    // The payloadType property
    payloadType *string
    // The signatures property
    signatures []SigstoreBundle0_dsseEnvelope_signaturesable
}
// NewSigstoreBundle0_dsseEnvelope instantiates a new SigstoreBundle0_dsseEnvelope and sets the default values.
func NewSigstoreBundle0_dsseEnvelope()(*SigstoreBundle0_dsseEnvelope) {
    m := &SigstoreBundle0_dsseEnvelope{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_dsseEnvelopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_dsseEnvelopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_dsseEnvelope(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_dsseEnvelope) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_dsseEnvelope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
        }
        return nil
    }
    res["payloadType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadType(val)
        }
        return nil
    }
    res["signatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSigstoreBundle0_dsseEnvelope_signaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SigstoreBundle0_dsseEnvelope_signaturesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SigstoreBundle0_dsseEnvelope_signaturesable)
                }
            }
            m.SetSignatures(res)
        }
        return nil
    }
    return res
}
// GetPayload gets the payload property value. The payload property
// returns a *string when successful
func (m *SigstoreBundle0_dsseEnvelope) GetPayload()(*string) {
    return m.payload
}
// GetPayloadType gets the payloadType property value. The payloadType property
// returns a *string when successful
func (m *SigstoreBundle0_dsseEnvelope) GetPayloadType()(*string) {
    return m.payloadType
}
// GetSignatures gets the signatures property value. The signatures property
// returns a []SigstoreBundle0_dsseEnvelope_signaturesable when successful
func (m *SigstoreBundle0_dsseEnvelope) GetSignatures()([]SigstoreBundle0_dsseEnvelope_signaturesable) {
    return m.signatures
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_dsseEnvelope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payloadType", m.GetPayloadType())
        if err != nil {
            return err
        }
    }
    if m.GetSignatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignatures()))
        for i, v := range m.GetSignatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("signatures", cast)
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
func (m *SigstoreBundle0_dsseEnvelope) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPayload sets the payload property value. The payload property
func (m *SigstoreBundle0_dsseEnvelope) SetPayload(value *string)() {
    m.payload = value
}
// SetPayloadType sets the payloadType property value. The payloadType property
func (m *SigstoreBundle0_dsseEnvelope) SetPayloadType(value *string)() {
    m.payloadType = value
}
// SetSignatures sets the signatures property value. The signatures property
func (m *SigstoreBundle0_dsseEnvelope) SetSignatures(value []SigstoreBundle0_dsseEnvelope_signaturesable)() {
    m.signatures = value
}
type SigstoreBundle0_dsseEnvelopeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPayload()(*string)
    GetPayloadType()(*string)
    GetSignatures()([]SigstoreBundle0_dsseEnvelope_signaturesable)
    SetPayload(value *string)()
    SetPayloadType(value *string)()
    SetSignatures(value []SigstoreBundle0_dsseEnvelope_signaturesable)()
}
