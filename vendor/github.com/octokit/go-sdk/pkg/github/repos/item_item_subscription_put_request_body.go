package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemSubscriptionPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Determines if all notifications should be blocked from this repository.
    ignored *bool
    // Determines if notifications should be received from this repository.
    subscribed *bool
}
// NewItemItemSubscriptionPutRequestBody instantiates a new ItemItemSubscriptionPutRequestBody and sets the default values.
func NewItemItemSubscriptionPutRequestBody()(*ItemItemSubscriptionPutRequestBody) {
    m := &ItemItemSubscriptionPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemSubscriptionPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemSubscriptionPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemSubscriptionPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemSubscriptionPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemSubscriptionPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ignored"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnored(val)
        }
        return nil
    }
    res["subscribed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscribed(val)
        }
        return nil
    }
    return res
}
// GetIgnored gets the ignored property value. Determines if all notifications should be blocked from this repository.
// returns a *bool when successful
func (m *ItemItemSubscriptionPutRequestBody) GetIgnored()(*bool) {
    return m.ignored
}
// GetSubscribed gets the subscribed property value. Determines if notifications should be received from this repository.
// returns a *bool when successful
func (m *ItemItemSubscriptionPutRequestBody) GetSubscribed()(*bool) {
    return m.subscribed
}
// Serialize serializes information the current object
func (m *ItemItemSubscriptionPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("ignored", m.GetIgnored())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("subscribed", m.GetSubscribed())
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
func (m *ItemItemSubscriptionPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIgnored sets the ignored property value. Determines if all notifications should be blocked from this repository.
func (m *ItemItemSubscriptionPutRequestBody) SetIgnored(value *bool)() {
    m.ignored = value
}
// SetSubscribed sets the subscribed property value. Determines if notifications should be received from this repository.
func (m *ItemItemSubscriptionPutRequestBody) SetSubscribed(value *bool)() {
    m.subscribed = value
}
type ItemItemSubscriptionPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIgnored()(*bool)
    GetSubscribed()(*bool)
    SetIgnored(value *bool)()
    SetSubscribed(value *bool)()
}
