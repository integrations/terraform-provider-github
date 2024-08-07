package notifications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ThreadsItemSubscriptionPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to block all notifications from a thread.
    ignored *bool
}
// NewThreadsItemSubscriptionPutRequestBody instantiates a new ThreadsItemSubscriptionPutRequestBody and sets the default values.
func NewThreadsItemSubscriptionPutRequestBody()(*ThreadsItemSubscriptionPutRequestBody) {
    m := &ThreadsItemSubscriptionPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateThreadsItemSubscriptionPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreadsItemSubscriptionPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreadsItemSubscriptionPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ThreadsItemSubscriptionPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ThreadsItemSubscriptionPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetIgnored gets the ignored property value. Whether to block all notifications from a thread.
// returns a *bool when successful
func (m *ThreadsItemSubscriptionPutRequestBody) GetIgnored()(*bool) {
    return m.ignored
}
// Serialize serializes information the current object
func (m *ThreadsItemSubscriptionPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("ignored", m.GetIgnored())
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
func (m *ThreadsItemSubscriptionPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIgnored sets the ignored property value. Whether to block all notifications from a thread.
func (m *ThreadsItemSubscriptionPutRequestBody) SetIgnored(value *bool)() {
    m.ignored = value
}
type ThreadsItemSubscriptionPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIgnored()(*bool)
    SetIgnored(value *bool)()
}
