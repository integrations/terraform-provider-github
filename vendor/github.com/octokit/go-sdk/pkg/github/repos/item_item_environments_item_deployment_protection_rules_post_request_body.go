package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the custom app that will be enabled on the environment.
    integration_id *int32
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody()(*ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["integration_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationId(val)
        }
        return nil
    }
    return res
}
// GetIntegrationId gets the integration_id property value. The ID of the custom app that will be enabled on the environment.
// returns a *int32 when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) GetIntegrationId()(*int32) {
    return m.integration_id
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("integration_id", m.GetIntegrationId())
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
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIntegrationId sets the integration_id property value. The ID of the custom app that will be enabled on the environment.
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBody) SetIntegrationId(value *int32)() {
    m.integration_id = value
}
type ItemItemEnvironmentsItemDeployment_protection_rulesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntegrationId()(*int32)
    SetIntegrationId(value *int32)()
}
