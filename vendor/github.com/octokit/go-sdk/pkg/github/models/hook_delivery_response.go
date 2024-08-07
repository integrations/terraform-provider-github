package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HookDelivery_response struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The response headers received when the delivery was made.
    headers HookDelivery_response_headersable
    // The response payload received.
    payload *string
}
// NewHookDelivery_response instantiates a new HookDelivery_response and sets the default values.
func NewHookDelivery_response()(*HookDelivery_response) {
    m := &HookDelivery_response{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHookDelivery_responseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHookDelivery_responseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHookDelivery_response(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HookDelivery_response) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HookDelivery_response) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["headers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHookDelivery_response_headersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaders(val.(HookDelivery_response_headersable))
        }
        return nil
    }
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
    return res
}
// GetHeaders gets the headers property value. The response headers received when the delivery was made.
// returns a HookDelivery_response_headersable when successful
func (m *HookDelivery_response) GetHeaders()(HookDelivery_response_headersable) {
    return m.headers
}
// GetPayload gets the payload property value. The response payload received.
// returns a *string when successful
func (m *HookDelivery_response) GetPayload()(*string) {
    return m.payload
}
// Serialize serializes information the current object
func (m *HookDelivery_response) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("headers", m.GetHeaders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payload", m.GetPayload())
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
func (m *HookDelivery_response) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeaders sets the headers property value. The response headers received when the delivery was made.
func (m *HookDelivery_response) SetHeaders(value HookDelivery_response_headersable)() {
    m.headers = value
}
// SetPayload sets the payload property value. The response payload received.
func (m *HookDelivery_response) SetPayload(value *string)() {
    m.payload = value
}
type HookDelivery_responseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeaders()(HookDelivery_response_headersable)
    GetPayload()(*string)
    SetHeaders(value HookDelivery_response_headersable)()
    SetPayload(value *string)()
}
