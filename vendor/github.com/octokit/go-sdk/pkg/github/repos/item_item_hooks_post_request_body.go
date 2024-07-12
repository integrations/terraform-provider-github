package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemHooksPostRequestBody struct {
    // Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
    active *bool
    // Key/value pairs to provide settings for this webhook.
    config ItemItemHooksPostRequestBody_configable
    // Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
    events []string
    // Use `web` to create a webhook. Default: `web`. This parameter only accepts the value `web`.
    name *string
}
// NewItemItemHooksPostRequestBody instantiates a new ItemItemHooksPostRequestBody and sets the default values.
func NewItemItemHooksPostRequestBody()(*ItemItemHooksPostRequestBody) {
    m := &ItemItemHooksPostRequestBody{
    }
    return m
}
// CreateItemItemHooksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemHooksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemHooksPostRequestBody(), nil
}
// GetActive gets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
// returns a *bool when successful
func (m *ItemItemHooksPostRequestBody) GetActive()(*bool) {
    return m.active
}
// GetConfig gets the config property value. Key/value pairs to provide settings for this webhook.
// returns a ItemItemHooksPostRequestBody_configable when successful
func (m *ItemItemHooksPostRequestBody) GetConfig()(ItemItemHooksPostRequestBody_configable) {
    return m.config
}
// GetEvents gets the events property value. Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
// returns a []string when successful
func (m *ItemItemHooksPostRequestBody) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemHooksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateItemItemHooksPostRequestBody_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(ItemItemHooksPostRequestBody_configable))
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
// GetName gets the name property value. Use `web` to create a webhook. Default: `web`. This parameter only accepts the value `web`.
// returns a *string when successful
func (m *ItemItemHooksPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemHooksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetActive sets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
func (m *ItemItemHooksPostRequestBody) SetActive(value *bool)() {
    m.active = value
}
// SetConfig sets the config property value. Key/value pairs to provide settings for this webhook.
func (m *ItemItemHooksPostRequestBody) SetConfig(value ItemItemHooksPostRequestBody_configable)() {
    m.config = value
}
// SetEvents sets the events property value. Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
func (m *ItemItemHooksPostRequestBody) SetEvents(value []string)() {
    m.events = value
}
// SetName sets the name property value. Use `web` to create a webhook. Default: `web`. This parameter only accepts the value `web`.
func (m *ItemItemHooksPostRequestBody) SetName(value *string)() {
    m.name = value
}
type ItemItemHooksPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(ItemItemHooksPostRequestBody_configable)
    GetEvents()([]string)
    GetName()(*string)
    SetActive(value *bool)()
    SetConfig(value ItemItemHooksPostRequestBody_configable)()
    SetEvents(value []string)()
    SetName(value *string)()
}
