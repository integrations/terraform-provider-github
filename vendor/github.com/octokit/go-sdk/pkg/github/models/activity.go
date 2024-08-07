package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Activity activity
type Activity struct {
    // The type of the activity that was performed.
    activity_type *Activity_activity_type
    // A GitHub user.
    actor NullableSimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The SHA of the commit after the activity.
    after *string
    // The SHA of the commit before the activity.
    before *string
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // The full Git reference, formatted as `refs/heads/<branch name>`.
    ref *string
    // The time when the activity occurred.
    timestamp *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewActivity instantiates a new Activity and sets the default values.
func NewActivity()(*Activity) {
    m := &Activity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivity(), nil
}
// GetActivityType gets the activity_type property value. The type of the activity that was performed.
// returns a *Activity_activity_type when successful
func (m *Activity) GetActivityType()(*Activity_activity_type) {
    return m.activity_type
}
// GetActor gets the actor property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Activity) GetActor()(NullableSimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Activity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfter gets the after property value. The SHA of the commit after the activity.
// returns a *string when successful
func (m *Activity) GetAfter()(*string) {
    return m.after
}
// GetBefore gets the before property value. The SHA of the commit before the activity.
// returns a *string when successful
func (m *Activity) GetBefore()(*string) {
    return m.before
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Activity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activity_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivity_activity_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val.(*Activity_activity_type))
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["after"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfter(val)
        }
        return nil
    }
    res["before"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBefore(val)
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Activity) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Activity) GetNodeId()(*string) {
    return m.node_id
}
// GetRef gets the ref property value. The full Git reference, formatted as `refs/heads/<branch name>`.
// returns a *string when successful
func (m *Activity) GetRef()(*string) {
    return m.ref
}
// GetTimestamp gets the timestamp property value. The time when the activity occurred.
// returns a *Time when successful
func (m *Activity) GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timestamp
}
// Serialize serializes information the current object
func (m *Activity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActivityType() != nil {
        cast := (*m.GetActivityType()).String()
        err := writer.WriteStringValue("activity_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("after", m.GetAfter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("before", m.GetBefore())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("timestamp", m.GetTimestamp())
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
// SetActivityType sets the activity_type property value. The type of the activity that was performed.
func (m *Activity) SetActivityType(value *Activity_activity_type)() {
    m.activity_type = value
}
// SetActor sets the actor property value. A GitHub user.
func (m *Activity) SetActor(value NullableSimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Activity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfter sets the after property value. The SHA of the commit after the activity.
func (m *Activity) SetAfter(value *string)() {
    m.after = value
}
// SetBefore sets the before property value. The SHA of the commit before the activity.
func (m *Activity) SetBefore(value *string)() {
    m.before = value
}
// SetId sets the id property value. The id property
func (m *Activity) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Activity) SetNodeId(value *string)() {
    m.node_id = value
}
// SetRef sets the ref property value. The full Git reference, formatted as `refs/heads/<branch name>`.
func (m *Activity) SetRef(value *string)() {
    m.ref = value
}
// SetTimestamp sets the timestamp property value. The time when the activity occurred.
func (m *Activity) SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timestamp = value
}
type Activityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityType()(*Activity_activity_type)
    GetActor()(NullableSimpleUserable)
    GetAfter()(*string)
    GetBefore()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetRef()(*string)
    GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActivityType(value *Activity_activity_type)()
    SetActor(value NullableSimpleUserable)()
    SetAfter(value *string)()
    SetBefore(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetRef(value *string)()
    SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
