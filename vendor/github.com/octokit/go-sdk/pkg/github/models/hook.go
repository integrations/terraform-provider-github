package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Hook webhooks for repositories.
type Hook struct {
    // Determines whether the hook is actually triggered on pushes.
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Configuration object of the webhook
    config WebhookConfigable
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deliveries_url property
    deliveries_url *string
    // Determines what events the hook is triggered for. Default: ['push'].
    events []string
    // Unique identifier of the webhook.
    id *int32
    // The last_response property
    last_response HookResponseable
    // The name of a valid service, use 'web' for a webhook.
    name *string
    // The ping_url property
    ping_url *string
    // The test_url property
    test_url *string
    // The type property
    typeEscaped *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewHook instantiates a new Hook and sets the default values.
func NewHook()(*Hook) {
    m := &Hook{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHook(), nil
}
// GetActive gets the active property value. Determines whether the hook is actually triggered on pushes.
// returns a *bool when successful
func (m *Hook) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Hook) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. Configuration object of the webhook
// returns a WebhookConfigable when successful
func (m *Hook) GetConfig()(WebhookConfigable) {
    return m.config
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Hook) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDeliveriesUrl gets the deliveries_url property value. The deliveries_url property
// returns a *string when successful
func (m *Hook) GetDeliveriesUrl()(*string) {
    return m.deliveries_url
}
// GetEvents gets the events property value. Determines what events the hook is triggered for. Default: ['push'].
// returns a []string when successful
func (m *Hook) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Hook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateWebhookConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(WebhookConfigable))
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["deliveries_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveriesUrl(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["last_response"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHookResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastResponse(val.(HookResponseable))
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
    res["ping_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPingUrl(val)
        }
        return nil
    }
    res["test_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestUrl(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier of the webhook.
// returns a *int32 when successful
func (m *Hook) GetId()(*int32) {
    return m.id
}
// GetLastResponse gets the last_response property value. The last_response property
// returns a HookResponseable when successful
func (m *Hook) GetLastResponse()(HookResponseable) {
    return m.last_response
}
// GetName gets the name property value. The name of a valid service, use 'web' for a webhook.
// returns a *string when successful
func (m *Hook) GetName()(*string) {
    return m.name
}
// GetPingUrl gets the ping_url property value. The ping_url property
// returns a *string when successful
func (m *Hook) GetPingUrl()(*string) {
    return m.ping_url
}
// GetTestUrl gets the test_url property value. The test_url property
// returns a *string when successful
func (m *Hook) GetTestUrl()(*string) {
    return m.test_url
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Hook) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Hook) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Hook) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Hook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deliveries_url", m.GetDeliveriesUrl())
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
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("last_response", m.GetLastResponse())
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
        err := writer.WriteStringValue("ping_url", m.GetPingUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("test_url", m.GetTestUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
// SetActive sets the active property value. Determines whether the hook is actually triggered on pushes.
func (m *Hook) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Hook) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. Configuration object of the webhook
func (m *Hook) SetConfig(value WebhookConfigable)() {
    m.config = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Hook) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDeliveriesUrl sets the deliveries_url property value. The deliveries_url property
func (m *Hook) SetDeliveriesUrl(value *string)() {
    m.deliveries_url = value
}
// SetEvents sets the events property value. Determines what events the hook is triggered for. Default: ['push'].
func (m *Hook) SetEvents(value []string)() {
    m.events = value
}
// SetId sets the id property value. Unique identifier of the webhook.
func (m *Hook) SetId(value *int32)() {
    m.id = value
}
// SetLastResponse sets the last_response property value. The last_response property
func (m *Hook) SetLastResponse(value HookResponseable)() {
    m.last_response = value
}
// SetName sets the name property value. The name of a valid service, use 'web' for a webhook.
func (m *Hook) SetName(value *string)() {
    m.name = value
}
// SetPingUrl sets the ping_url property value. The ping_url property
func (m *Hook) SetPingUrl(value *string)() {
    m.ping_url = value
}
// SetTestUrl sets the test_url property value. The test_url property
func (m *Hook) SetTestUrl(value *string)() {
    m.test_url = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Hook) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Hook) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Hook) SetUrl(value *string)() {
    m.url = value
}
type Hookable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(WebhookConfigable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeliveriesUrl()(*string)
    GetEvents()([]string)
    GetId()(*int32)
    GetLastResponse()(HookResponseable)
    GetName()(*string)
    GetPingUrl()(*string)
    GetTestUrl()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetActive(value *bool)()
    SetConfig(value WebhookConfigable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeliveriesUrl(value *string)()
    SetEvents(value []string)()
    SetId(value *int32)()
    SetLastResponse(value HookResponseable)()
    SetName(value *string)()
    SetPingUrl(value *string)()
    SetTestUrl(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
