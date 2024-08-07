package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Status the status of a commit.
type Status struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // The context property
    context *string
    // The created_at property
    created_at *string
    // A GitHub user.
    creator NullableSimpleUserable
    // The description property
    description *string
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // The state property
    state *string
    // The target_url property
    target_url *string
    // The updated_at property
    updated_at *string
    // The url property
    url *string
}
// NewStatus instantiates a new Status and sets the default values.
func NewStatus()(*Status) {
    m := &Status{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Status) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *Status) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetContext gets the context property value. The context property
// returns a *string when successful
func (m *Status) GetContext()(*string) {
    return m.context
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *Status) GetCreatedAt()(*string) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Status) GetCreator()(NullableSimpleUserable) {
    return m.creator
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Status) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Status) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["avatar_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvatarUrl(val)
        }
        return nil
    }
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["creator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["target_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetUrl(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Status) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Status) GetNodeId()(*string) {
    return m.node_id
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *Status) GetState()(*string) {
    return m.state
}
// GetTargetUrl gets the target_url property value. The target_url property
// returns a *string when successful
func (m *Status) GetTargetUrl()(*string) {
    return m.target_url
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *Status) GetUpdatedAt()(*string) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Status) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Status) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("creator", m.GetCreator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_url", m.GetTargetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Status) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *Status) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetContext sets the context property value. The context property
func (m *Status) SetContext(value *string)() {
    m.context = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Status) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *Status) SetCreator(value NullableSimpleUserable)() {
    m.creator = value
}
// SetDescription sets the description property value. The description property
func (m *Status) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The id property
func (m *Status) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Status) SetNodeId(value *string)() {
    m.node_id = value
}
// SetState sets the state property value. The state property
func (m *Status) SetState(value *string)() {
    m.state = value
}
// SetTargetUrl sets the target_url property value. The target_url property
func (m *Status) SetTargetUrl(value *string)() {
    m.target_url = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Status) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Status) SetUrl(value *string)() {
    m.url = value
}
type Statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetContext()(*string)
    GetCreatedAt()(*string)
    GetCreator()(NullableSimpleUserable)
    GetDescription()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetState()(*string)
    GetTargetUrl()(*string)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetContext(value *string)()
    SetCreatedAt(value *string)()
    SetCreator(value NullableSimpleUserable)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetState(value *string)()
    SetTargetUrl(value *string)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
