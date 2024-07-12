package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HookDeliveryItem delivery made by a webhook, without request and response information.
type HookDeliveryItem struct {
    // The type of activity for the event that triggered the delivery.
    action *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time when the webhook delivery occurred.
    delivered_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Time spent delivering.
    duration *float64
    // The event that triggered the delivery.
    event *string
    // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
    guid *string
    // Unique identifier of the webhook delivery.
    id *int32
    // The id of the GitHub App installation associated with this event.
    installation_id *int32
    // Whether the webhook delivery is a redelivery.
    redelivery *bool
    // The id of the repository associated with this event.
    repository_id *int32
    // Describes the response returned after attempting the delivery.
    status *string
    // Status code received when delivery was made.
    status_code *int32
    // Time when the webhook delivery was throttled.
    throttled_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewHookDeliveryItem instantiates a new HookDeliveryItem and sets the default values.
func NewHookDeliveryItem()(*HookDeliveryItem) {
    m := &HookDeliveryItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHookDeliveryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHookDeliveryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHookDeliveryItem(), nil
}
// GetAction gets the action property value. The type of activity for the event that triggered the delivery.
// returns a *string when successful
func (m *HookDeliveryItem) GetAction()(*string) {
    return m.action
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HookDeliveryItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeliveredAt gets the delivered_at property value. Time when the webhook delivery occurred.
// returns a *Time when successful
func (m *HookDeliveryItem) GetDeliveredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.delivered_at
}
// GetDuration gets the duration property value. Time spent delivering.
// returns a *float64 when successful
func (m *HookDeliveryItem) GetDuration()(*float64) {
    return m.duration
}
// GetEvent gets the event property value. The event that triggered the delivery.
// returns a *string when successful
func (m *HookDeliveryItem) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HookDeliveryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetGuid gets the guid property value. Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
// returns a *string when successful
func (m *HookDeliveryItem) GetGuid()(*string) {
    return m.guid
}
// GetId gets the id property value. Unique identifier of the webhook delivery.
// returns a *int32 when successful
func (m *HookDeliveryItem) GetId()(*int32) {
    return m.id
}
// GetInstallationId gets the installation_id property value. The id of the GitHub App installation associated with this event.
// returns a *int32 when successful
func (m *HookDeliveryItem) GetInstallationId()(*int32) {
    return m.installation_id
}
// GetRedelivery gets the redelivery property value. Whether the webhook delivery is a redelivery.
// returns a *bool when successful
func (m *HookDeliveryItem) GetRedelivery()(*bool) {
    return m.redelivery
}
// GetRepositoryId gets the repository_id property value. The id of the repository associated with this event.
// returns a *int32 when successful
func (m *HookDeliveryItem) GetRepositoryId()(*int32) {
    return m.repository_id
}
// GetStatus gets the status property value. Describes the response returned after attempting the delivery.
// returns a *string when successful
func (m *HookDeliveryItem) GetStatus()(*string) {
    return m.status
}
// GetStatusCode gets the status_code property value. Status code received when delivery was made.
// returns a *int32 when successful
func (m *HookDeliveryItem) GetStatusCode()(*int32) {
    return m.status_code
}
// GetThrottledAt gets the throttled_at property value. Time when the webhook delivery was throttled.
// returns a *Time when successful
func (m *HookDeliveryItem) GetThrottledAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.throttled_at
}
// Serialize serializes information the current object
func (m *HookDeliveryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The type of activity for the event that triggered the delivery.
func (m *HookDeliveryItem) SetAction(value *string)() {
    m.action = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HookDeliveryItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeliveredAt sets the delivered_at property value. Time when the webhook delivery occurred.
func (m *HookDeliveryItem) SetDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.delivered_at = value
}
// SetDuration sets the duration property value. Time spent delivering.
func (m *HookDeliveryItem) SetDuration(value *float64)() {
    m.duration = value
}
// SetEvent sets the event property value. The event that triggered the delivery.
func (m *HookDeliveryItem) SetEvent(value *string)() {
    m.event = value
}
// SetGuid sets the guid property value. Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
func (m *HookDeliveryItem) SetGuid(value *string)() {
    m.guid = value
}
// SetId sets the id property value. Unique identifier of the webhook delivery.
func (m *HookDeliveryItem) SetId(value *int32)() {
    m.id = value
}
// SetInstallationId sets the installation_id property value. The id of the GitHub App installation associated with this event.
func (m *HookDeliveryItem) SetInstallationId(value *int32)() {
    m.installation_id = value
}
// SetRedelivery sets the redelivery property value. Whether the webhook delivery is a redelivery.
func (m *HookDeliveryItem) SetRedelivery(value *bool)() {
    m.redelivery = value
}
// SetRepositoryId sets the repository_id property value. The id of the repository associated with this event.
func (m *HookDeliveryItem) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
// SetStatus sets the status property value. Describes the response returned after attempting the delivery.
func (m *HookDeliveryItem) SetStatus(value *string)() {
    m.status = value
}
// SetStatusCode sets the status_code property value. Status code received when delivery was made.
func (m *HookDeliveryItem) SetStatusCode(value *int32)() {
    m.status_code = value
}
// SetThrottledAt sets the throttled_at property value. Time when the webhook delivery was throttled.
func (m *HookDeliveryItem) SetThrottledAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.throttled_at = value
}
type HookDeliveryItemable interface {
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
    GetStatus()(*string)
    GetStatusCode()(*int32)
    GetThrottledAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAction(value *string)()
    SetDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDuration(value *float64)()
    SetEvent(value *string)()
    SetGuid(value *string)()
    SetId(value *int32)()
    SetInstallationId(value *int32)()
    SetRedelivery(value *bool)()
    SetRepositoryId(value *int32)()
    SetStatus(value *string)()
    SetStatusCode(value *int32)()
    SetThrottledAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
