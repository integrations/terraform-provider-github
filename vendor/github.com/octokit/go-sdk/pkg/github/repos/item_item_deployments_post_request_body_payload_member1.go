package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDeploymentsPostRequestBody_payloadMember1 
type ItemItemDeploymentsPostRequestBody_payloadMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemItemDeploymentsPostRequestBody_payloadMember1 instantiates a new ItemItemDeploymentsPostRequestBody_payloadMember1 and sets the default values.
func NewItemItemDeploymentsPostRequestBody_payloadMember1()(*ItemItemDeploymentsPostRequestBody_payloadMember1) {
    m := &ItemItemDeploymentsPostRequestBody_payloadMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemDeploymentsPostRequestBody_payloadMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDeploymentsPostRequestBody_payloadMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDeploymentsPostRequestBody_payloadMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemDeploymentsPostRequestBody_payloadMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemDeploymentsPostRequestBody_payloadMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemItemDeploymentsPostRequestBody_payloadMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemDeploymentsPostRequestBody_payloadMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ItemItemDeploymentsPostRequestBody_payloadMember1able 
type ItemItemDeploymentsPostRequestBody_payloadMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
