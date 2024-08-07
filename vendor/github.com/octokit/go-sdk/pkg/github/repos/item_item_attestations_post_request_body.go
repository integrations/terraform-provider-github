package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemAttestationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
    bundle ItemItemAttestationsPostRequestBody_bundleable
}
// NewItemItemAttestationsPostRequestBody instantiates a new ItemItemAttestationsPostRequestBody and sets the default values.
func NewItemItemAttestationsPostRequestBody()(*ItemItemAttestationsPostRequestBody) {
    m := &ItemItemAttestationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemAttestationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemAttestationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemAttestationsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemAttestationsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBundle gets the bundle property value. The attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
// returns a ItemItemAttestationsPostRequestBody_bundleable when successful
func (m *ItemItemAttestationsPostRequestBody) GetBundle()(ItemItemAttestationsPostRequestBody_bundleable) {
    return m.bundle
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemAttestationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bundle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemAttestationsPostRequestBody_bundleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundle(val.(ItemItemAttestationsPostRequestBody_bundleable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemAttestationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bundle", m.GetBundle())
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
func (m *ItemItemAttestationsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBundle sets the bundle property value. The attestation's Sigstore Bundle.Refer to the [Sigstore Bundle Specification](https://github.com/sigstore/protobuf-specs/blob/main/protos/sigstore_bundle.proto) for more information.
func (m *ItemItemAttestationsPostRequestBody) SetBundle(value ItemItemAttestationsPostRequestBody_bundleable)() {
    m.bundle = value
}
type ItemItemAttestationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBundle()(ItemItemAttestationsPostRequestBody_bundleable)
    SetBundle(value ItemItemAttestationsPostRequestBody_bundleable)()
}
