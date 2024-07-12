package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemHooksItemWithHook_PatchRequestBody struct {
    // Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
    active *bool
    // Determines a list of events to be added to the list of events that the Hook triggers for.
    add_events []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Configuration object of the webhook
    config i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigable
    // Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. This replaces the entire array of events.
    events []string
    // Determines a list of events to be removed from the list of events that the Hook triggers for.
    remove_events []string
}
// NewItemItemHooksItemWithHook_PatchRequestBody instantiates a new ItemItemHooksItemWithHook_PatchRequestBody and sets the default values.
func NewItemItemHooksItemWithHook_PatchRequestBody()(*ItemItemHooksItemWithHook_PatchRequestBody) {
    m := &ItemItemHooksItemWithHook_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemHooksItemWithHook_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemHooksItemWithHook_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemHooksItemWithHook_PatchRequestBody(), nil
}
// GetActive gets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
// returns a *bool when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetActive()(*bool) {
    return m.active
}
// GetAddEvents gets the add_events property value. Determines a list of events to be added to the list of events that the Hook triggers for.
// returns a []string when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetAddEvents()([]string) {
    return m.add_events
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. Configuration object of the webhook
// returns a WebhookConfigable when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigable) {
    return m.config
}
// GetEvents gets the events property value. Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. This replaces the entire array of events.
// returns a []string when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["add_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAddEvents(res)
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateWebhookConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigable))
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
    res["remove_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRemoveEvents(res)
        }
        return nil
    }
    return res
}
// GetRemoveEvents gets the remove_events property value. Determines a list of events to be removed from the list of events that the Hook triggers for.
// returns a []string when successful
func (m *ItemItemHooksItemWithHook_PatchRequestBody) GetRemoveEvents()([]string) {
    return m.remove_events
}
// Serialize serializes information the current object
func (m *ItemItemHooksItemWithHook_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    if m.GetAddEvents() != nil {
        err := writer.WriteCollectionOfStringValues("add_events", m.GetAddEvents())
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
    if m.GetRemoveEvents() != nil {
        err := writer.WriteCollectionOfStringValues("remove_events", m.GetRemoveEvents())
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
func (m *ItemItemHooksItemWithHook_PatchRequestBody) SetActive(value *bool)() {
    m.active = value
}
// SetAddEvents sets the add_events property value. Determines a list of events to be added to the list of events that the Hook triggers for.
func (m *ItemItemHooksItemWithHook_PatchRequestBody) SetAddEvents(value []string)() {
    m.add_events = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemHooksItemWithHook_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. Configuration object of the webhook
func (m *ItemItemHooksItemWithHook_PatchRequestBody) SetConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigable)() {
    m.config = value
}
// SetEvents sets the events property value. Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for. This replaces the entire array of events.
func (m *ItemItemHooksItemWithHook_PatchRequestBody) SetEvents(value []string)() {
    m.events = value
}
// SetRemoveEvents sets the remove_events property value. Determines a list of events to be removed from the list of events that the Hook triggers for.
func (m *ItemItemHooksItemWithHook_PatchRequestBody) SetRemoveEvents(value []string)() {
    m.remove_events = value
}
type ItemItemHooksItemWithHook_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetAddEvents()([]string)
    GetConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigable)
    GetEvents()([]string)
    GetRemoveEvents()([]string)
    SetActive(value *bool)()
    SetAddEvents(value []string)()
    SetConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigable)()
    SetEvents(value []string)()
    SetRemoveEvents(value []string)()
}
