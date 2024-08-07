package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HookDelivery delivery made by a webhook.
type HookDelivery struct {
    // The type of activity for the event that triggered the delivery.
    action *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time when the delivery was delivered.
    delivered_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Time spent delivering.
    duration *float64
    // The event that triggered the delivery.
    event *string
    // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
    guid *string
    // Unique identifier of the delivery.
    id *int32
    // The id of the GitHub App installation associated with this event.
    installation_id *int32
    // Whether the delivery is a redelivery.
    redelivery *bool
    // The id of the repository associated with this event.
    repository_id *int32
    // The request property
    request HookDelivery_requestable
    // The response property
    response HookDelivery_responseable
    // Description of the status of the attempted delivery
    status *string
    // Status code received when delivery was made.
    status_code *int32
    // Time when the webhook delivery was throttled.
    throttled_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL target of the delivery.
    url *string
}
// NewHookDelivery instantiates a new HookDelivery and sets the default values.
func NewHookDelivery()(*HookDelivery) {
    m := &HookDelivery{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHookDeliveryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHookDeliveryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHookDelivery(), nil
}
// GetAction gets the action property value. The type of activity for the event that triggered the delivery.
// returns a *string when successful
func (m *HookDelivery) GetAction()(*string) {
    return m.action
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HookDelivery) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeliveredAt gets the delivered_at property value. Time when the delivery was delivered.
// returns a *Time when successful
func (m *HookDelivery) GetDeliveredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.delivered_at
}
// GetDuration gets the duration property value. Time spent delivering.
// returns a *float64 when successful
func (m *HookDelivery) GetDuration()(*float64) {
    return m.duration
}
// GetEvent gets the event property value. The event that triggered the delivery.
// returns a *string when successful
func (m *HookDelivery) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HookDelivery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val)
        }
        return nil
    }
    res["delivered_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveredAt(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
        }
        return nil
    }
    res["guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
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
    res["installation_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallationId(val)
        }
        return nil
    }
    res["redelivery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedelivery(val)
        }
        return nil
    }
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    res["request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHookDelivery_requestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(HookDelivery_requestable))
        }
        return nil
    }
    res["response"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHookDelivery_responseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponse(val.(HookDelivery_responseable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["status_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusCode(val)
        }
        return nil
    }
    res["throttled_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThrottledAt(val)
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
// GetGuid gets the guid property value. Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
// returns a *string when successful
func (m *HookDelivery) GetGuid()(*string) {
    return m.guid
}
// GetId gets the id property value. Unique identifier of the delivery.
// returns a *int32 when successful
func (m *HookDelivery) GetId()(*int32) {
    return m.id
}
// GetInstallationId gets the installation_id property value. The id of the GitHub App installation associated with this event.
// returns a *int32 when successful
func (m *HookDelivery) GetInstallationId()(*int32) {
    return m.installation_id
}
// GetRedelivery gets the redelivery property value. Whether the delivery is a redelivery.
// returns a *bool when successful
func (m *HookDelivery) GetRedelivery()(*bool) {
    return m.redelivery
}
// GetRepositoryId gets the repository_id property value. The id of the repository associated with this event.
// returns a *int32 when successful
func (m *HookDelivery) GetRepositoryId()(*int32) {
    return m.repository_id
}
// GetRequest gets the request property value. The request property
// returns a HookDelivery_requestable when successful
func (m *HookDelivery) GetRequest()(HookDelivery_requestable) {
    return m.request
}
// GetResponse gets the response property value. The response property
// returns a HookDelivery_responseable when successful
func (m *HookDelivery) GetResponse()(HookDelivery_responseable) {
    return m.response
}
// GetStatus gets the status property value. Description of the status of the attempted delivery
// returns a *string when successful
func (m *HookDelivery) GetStatus()(*string) {
    return m.status
}
// GetStatusCode gets the status_code property value. Status code received when delivery was made.
// returns a *int32 when successful
func (m *HookDelivery) GetStatusCode()(*int32) {
    return m.status_code
}
// GetThrottledAt gets the throttled_at property value. Time when the webhook delivery was throttled.
// returns a *Time when successful
func (m *HookDelivery) GetThrottledAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.throttled_at
}
// GetUrl gets the url property value. The URL target of the delivery.
// returns a *string when successful
func (m *HookDelivery) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *HookDelivery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("delivered_at", m.GetDeliveredAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("guid", m.GetGuid())
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
        err := writer.WriteInt32Value("installation_id", m.GetInstallationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("redelivery", m.GetRedelivery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repository_id", m.GetRepositoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("request", m.GetRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("response", m.GetResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("status_code", m.GetStatusCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("throttled_at", m.GetThrottledAt())
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
// SetAction sets the action property value. The type of activity for the event that triggered the delivery.
func (m *HookDelivery) SetAction(value *string)() {
    m.action = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HookDelivery) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeliveredAt sets the delivered_at property value. Time when the delivery was delivered.
func (m *HookDelivery) SetDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.delivered_at = value
}
// SetDuration sets the duration property value. Time spent delivering.
func (m *HookDelivery) SetDuration(value *float64)() {
    m.duration = value
}
// SetEvent sets the event property value. The event that triggered the delivery.
func (m *HookDelivery) SetEvent(value *string)() {
    m.event = value
}
// SetGuid sets the guid property value. Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
func (m *HookDelivery) SetGuid(value *string)() {
    m.guid = value
}
// SetId sets the id property value. Unique identifier of the delivery.
func (m *HookDelivery) SetId(value *int32)() {
    m.id = value
}
// SetInstallationId sets the installation_id property value. The id of the GitHub App installation associated with this event.
func (m *HookDelivery) SetInstallationId(value *int32)() {
    m.installation_id = value
}
// SetRedelivery sets the redelivery property value. Whether the delivery is a redelivery.
func (m *HookDelivery) SetRedelivery(value *bool)() {
    m.redelivery = value
}
// SetRepositoryId sets the repository_id property value. The id of the repository associated with this event.
func (m *HookDelivery) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
// SetRequest sets the request property value. The request property
func (m *HookDelivery) SetRequest(value HookDelivery_requestable)() {
    m.request = value
}
// SetResponse sets the response property value. The response property
func (m *HookDelivery) SetResponse(value HookDelivery_responseable)() {
    m.response = value
}
// SetStatus sets the status property value. Description of the status of the attempted delivery
func (m *HookDelivery) SetStatus(value *string)() {
    m.status = value
}
// SetStatusCode sets the status_code property value. Status code received when delivery was made.
func (m *HookDelivery) SetStatusCode(value *int32)() {
    m.status_code = value
}
// SetThrottledAt sets the throttled_at property value. Time when the webhook delivery was throttled.
func (m *HookDelivery) SetThrottledAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.throttled_at = value
}
// SetUrl sets the url property value. The URL target of the delivery.
func (m *HookDelivery) SetUrl(value *string)() {
    m.url = value
}
type HookDeliveryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetDeliveredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDuration()(*float64)
    GetEvent()(*string)
    GetGuid()(*string)
    GetId()(*int32)
    GetInstallationId()(*int32)
    GetRedelivery()(*bool)
    GetRepositoryId()(*int32)
    GetRequest()(HookDelivery_requestable)
    GetResponse()(HookDelivery_responseable)
    GetStatus()(*string)
    GetStatusCode()(*int32)
    GetThrottledAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAction(value *string)()
    SetDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDuration(value *float64)()
    SetEvent(value *string)()
    SetGuid(value *string)()
    SetId(value *int32)()
    SetInstallationId(value *int32)()
    SetRedelivery(value *bool)()
    SetRepositoryId(value *int32)()
    SetRequest(value HookDelivery_requestable)()
    SetResponse(value HookDelivery_responseable)()
    SetStatus(value *string)()
    SetStatusCode(value *int32)()
    SetThrottledAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
