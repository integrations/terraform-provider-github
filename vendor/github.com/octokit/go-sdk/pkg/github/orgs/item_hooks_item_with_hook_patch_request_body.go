package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemHooksItemWithHook_PatchRequestBody struct {
    // Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Key/value pairs to provide settings for this webhook.
    config ItemHooksItemWithHook_PatchRequestBody_configable
    // Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
    events []string
    // The name property
    name *string
}
// NewItemHooksItemWithHook_PatchRequestBody instantiates a new ItemHooksItemWithHook_PatchRequestBody and sets the default values.
func NewItemHooksItemWithHook_PatchRequestBody()(*ItemHooksItemWithHook_PatchRequestBody) {
    m := &ItemHooksItemWithHook_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemHooksItemWithHook_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemHooksItemWithHook_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemHooksItemWithHook_PatchRequestBody(), nil
}
// GetActive gets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
// returns a *bool when successful
func (m *ItemHooksItemWithHook_PatchRequestBody) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemHooksItemWithHook_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. Key/value pairs to provide settings for this webhook.
// returns a ItemHooksItemWithHook_PatchRequestBody_configable when successful
func (m *ItemHooksItemWithHook_PatchRequestBody) GetConfig()(ItemHooksItemWithHook_PatchRequestBody_configable) {
    return m.config
}
// GetEvents gets the events property value. Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
// returns a []string when successful
func (m *ItemHooksItemWithHook_PatchRequestBody) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemHooksItemWithHook_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemHooksItemWithHook_PatchRequestBody_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(ItemHooksItemWithHook_PatchRequestBody_configable))
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ItemHooksItemWithHook_PatchRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemHooksItemWithHook_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        err := writer.WriteCollectionOfStringValues("events", m.GetEvents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
// SetActive sets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
func (m *ItemHooksItemWithHook_PatchRequestBody) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemHooksItemWithHook_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. Key/value pairs to provide settings for this webhook.
func (m *ItemHooksItemWithHook_PatchRequestBody) SetConfig(value ItemHooksItemWithHook_PatchRequestBody_configable)() {
    m.config = value
}
// SetEvents sets the events property value. Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
func (m *ItemHooksItemWithHook_PatchRequestBody) SetEvents(value []string)() {
    m.events = value
}
// SetName sets the name property value. The name property
func (m *ItemHooksItemWithHook_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
type ItemHooksItemWithHook_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(ItemHooksItemWithHook_PatchRequestBody_configable)
    GetEvents()([]string)
    GetName()(*string)
    SetActive(value *bool)()
    SetConfig(value ItemHooksItemWithHook_PatchRequestBody_configable)()
    SetEvents(value []string)()
    SetName(value *string)()
}
