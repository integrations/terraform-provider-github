package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemAttestationsPostRequestBody_bundle the attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
type ItemItemAttestationsPostRequestBody_bundle struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dsseEnvelope property
    dsseEnvelope ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable
    // The mediaType property
    mediaType *string
    // The verificationMaterial property
    verificationMaterial ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable
}
// NewItemItemAttestationsPostRequestBody_bundle instantiates a new ItemItemAttestationsPostRequestBody_bundle and sets the default values.
func NewItemItemAttestationsPostRequestBody_bundle()(*ItemItemAttestationsPostRequestBody_bundle) {
    m := &ItemItemAttestationsPostRequestBody_bundle{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemAttestationsPostRequestBody_bundleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemAttestationsPostRequestBody_bundleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemAttestationsPostRequestBody_bundle(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemAttestationsPostRequestBody_bundle) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDsseEnvelope gets the dsseEnvelope property value. The dsseEnvelope property
// returns a ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable when successful
func (m *ItemItemAttestationsPostRequestBody_bundle) GetDsseEnvelope()(ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable) {
    return m.dsseEnvelope
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemAttestationsPostRequestBody_bundle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dsseEnvelope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDsseEnvelope(val.(ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable))
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
        val, err := n.GetObjectValue(CreateItemItemAttestationsPostRequestBody_bundle_verificationMaterialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerificationMaterial(val.(ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable))
        }
        return nil
    }
    return res
}
// GetMediaType gets the mediaType property value. The mediaType property
// returns a *string when successful
func (m *ItemItemAttestationsPostRequestBody_bundle) GetMediaType()(*string) {
    return m.mediaType
}
// GetVerificationMaterial gets the verificationMaterial property value. The verificationMaterial property
// returns a ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable when successful
func (m *ItemItemAttestationsPostRequestBody_bundle) GetVerificationMaterial()(ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable) {
    return m.verificationMaterial
}
// Serialize serializes information the current object
func (m *ItemItemAttestationsPostRequestBody_bundle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemAttestationsPostRequestBody_bundle) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDsseEnvelope sets the dsseEnvelope property value. The dsseEnvelope property
func (m *ItemItemAttestationsPostRequestBody_bundle) SetDsseEnvelope(value ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable)() {
    m.dsseEnvelope = value
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *ItemItemAttestationsPostRequestBody_bundle) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetVerificationMaterial sets the verificationMaterial property value. The verificationMaterial property
func (m *ItemItemAttestationsPostRequestBody_bundle) SetVerificationMaterial(value ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable)() {
    m.verificationMaterial = value
}
type ItemItemAttestationsPostRequestBody_bundleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDsseEnvelope()(ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable)
    GetMediaType()(*string)
    GetVerificationMaterial()(ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable)
    SetDsseEnvelope(value ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable)()
    SetMediaType(value *string)()
    SetVerificationMaterial(value ItemItemAttestationsPostRequestBody_bundle_verificationMaterialable)()
}
