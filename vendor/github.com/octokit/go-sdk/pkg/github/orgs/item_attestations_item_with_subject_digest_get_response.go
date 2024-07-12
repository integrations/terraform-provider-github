package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemAttestationsItemWithSubject_digestGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attestations property
    attestations []ItemAttestationsItemWithSubject_digestGetResponse_attestationsable
}
// NewItemAttestationsItemWithSubject_digestGetResponse instantiates a new ItemAttestationsItemWithSubject_digestGetResponse and sets the default values.
func NewItemAttestationsItemWithSubject_digestGetResponse()(*ItemAttestationsItemWithSubject_digestGetResponse) {
    m := &ItemAttestationsItemWithSubject_digestGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAttestationsItemWithSubject_digestGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAttestationsItemWithSubject_digestGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAttestationsItemWithSubject_digestGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttestations gets the attestations property value. The attestations property
// returns a []ItemAttestationsItemWithSubject_digestGetResponse_attestationsable when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse) GetAttestations()([]ItemAttestationsItemWithSubject_digestGetResponse_attestationsable) {
    return m.attestations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attestations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemAttestationsItemWithSubject_digestGetResponse_attestationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemAttestationsItemWithSubject_digestGetResponse_attestationsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemAttestationsItemWithSubject_digestGetResponse_attestationsable)
                }
            }
            m.SetAttestations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemAttestationsItemWithSubject_digestGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttestations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttestations()))
        for i, v := range m.GetAttestations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attestations", cast)
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
func (m *ItemAttestationsItemWithSubject_digestGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttestations sets the attestations property value. The attestations property
func (m *ItemAttestationsItemWithSubject_digestGetResponse) SetAttestations(value []ItemAttestationsItemWithSubject_digestGetResponse_attestationsable)() {
    m.attestations = value
}
type ItemAttestationsItemWithSubject_digestGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttestations()([]ItemAttestationsItemWithSubject_digestGetResponse_attestationsable)
    SetAttestations(value []ItemAttestationsItemWithSubject_digestGetResponse_attestationsable)()
}
