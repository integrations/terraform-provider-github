package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemDispatchesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // JSON payload with extra information about the webhook event that your action or workflow may use. The maximum number of top-level properties is 10.
    client_payload ItemItemDispatchesPostRequestBody_client_payloadable
    // A custom webhook event name. Must be 100 characters or fewer.
    event_type *string
}
// NewItemItemDispatchesPostRequestBody instantiates a new ItemItemDispatchesPostRequestBody and sets the default values.
func NewItemItemDispatchesPostRequestBody()(*ItemItemDispatchesPostRequestBody) {
    m := &ItemItemDispatchesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemDispatchesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemDispatchesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDispatchesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemDispatchesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientPayload gets the client_payload property value. JSON payload with extra information about the webhook event that your action or workflow may use. The maximum number of top-level properties is 10.
// returns a ItemItemDispatchesPostRequestBody_client_payloadable when successful
func (m *ItemItemDispatchesPostRequestBody) GetClientPayload()(ItemItemDispatchesPostRequestBody_client_payloadable) {
    return m.client_payload
}
// GetEventType gets the event_type property value. A custom webhook event name. Must be 100 characters or fewer.
// returns a *string when successful
func (m *ItemItemDispatchesPostRequestBody) GetEventType()(*string) {
    return m.event_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemDispatchesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemDispatchesPostRequestBody_client_payloadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientPayload(val.(ItemItemDispatchesPostRequestBody_client_payloadable))
        }
        return nil
    }
    res["event_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemDispatchesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("client_payload", m.GetClientPayload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event_type", m.GetEventType())
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
func (m *ItemItemDispatchesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientPayload sets the client_payload property value. JSON payload with extra information about the webhook event that your action or workflow may use. The maximum number of top-level properties is 10.
func (m *ItemItemDispatchesPostRequestBody) SetClientPayload(value ItemItemDispatchesPostRequestBody_client_payloadable)() {
    m.client_payload = value
}
// SetEventType sets the event_type property value. A custom webhook event name. Must be 100 characters or fewer.
func (m *ItemItemDispatchesPostRequestBody) SetEventType(value *string)() {
    m.event_type = value
}
type ItemItemDispatchesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientPayload()(ItemItemDispatchesPostRequestBody_client_payloadable)
    GetEventType()(*string)
    SetClientPayload(value ItemItemDispatchesPostRequestBody_client_payloadable)()
    SetEventType(value *string)()
}
