package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HookDelivery_request struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The request headers sent with the webhook delivery.
    headers HookDelivery_request_headersable
    // The webhook payload.
    payload HookDelivery_request_payloadable
}
// NewHookDelivery_request instantiates a new HookDelivery_request and sets the default values.
func NewHookDelivery_request()(*HookDelivery_request) {
    m := &HookDelivery_request{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHookDelivery_requestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHookDelivery_requestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHookDelivery_request(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HookDelivery_request) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HookDelivery_request) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["headers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHookDelivery_request_headersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaders(val.(HookDelivery_request_headersable))
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHookDelivery_request_payloadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(HookDelivery_request_payloadable))
        }
        return nil
    }
    return res
}
// GetHeaders gets the headers property value. The request headers sent with the webhook delivery.
// returns a HookDelivery_request_headersable when successful
func (m *HookDelivery_request) GetHeaders()(HookDelivery_request_headersable) {
    return m.headers
}
// GetPayload gets the payload property value. The webhook payload.
// returns a HookDelivery_request_payloadable when successful
func (m *HookDelivery_request) GetPayload()(HookDelivery_request_payloadable) {
    return m.payload
}
// Serialize serializes information the current object
func (m *HookDelivery_request) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("headers", m.GetHeaders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("payload", m.GetPayload())
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
func (m *HookDelivery_request) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeaders sets the headers property value. The request headers sent with the webhook delivery.
func (m *HookDelivery_request) SetHeaders(value HookDelivery_request_headersable)() {
    m.headers = value
}
// SetPayload sets the payload property value. The webhook payload.
func (m *HookDelivery_request) SetPayload(value HookDelivery_request_payloadable)() {
    m.payload = value
}
type HookDelivery_requestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeaders()(HookDelivery_request_headersable)
    GetPayload()(HookDelivery_request_payloadable)
    SetHeaders(value HookDelivery_request_headersable)()
    SetPayload(value HookDelivery_request_payloadable)()
}
