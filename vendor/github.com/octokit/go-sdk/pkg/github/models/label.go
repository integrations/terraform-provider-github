package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Label color-coded labels help you categorize and filter your issues (just like labels in Gmail).
type Label struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // 6-character hex code, without the leading #, identifying the color
    color *string
    // The default property
    defaultEscaped *bool
    // The description property
    description *string
    // The id property
    id *int64
    // The name of the label.
    name *string
    // The node_id property
    node_id *string
    // URL for the label
    url *string
}
// NewLabel instantiates a new Label and sets the default values.
func NewLabel()(*Label) {
    m := &Label{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabel(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Label) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. 6-character hex code, without the leading #, identifying the color
// returns a *string when successful
func (m *Label) GetColor()(*string) {
    return m.color
}
// GetDefaultEscaped gets the default property value. The default property
// returns a *bool when successful
func (m *Label) GetDefaultEscaped()(*bool) {
    return m.defaultEscaped
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Label) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Label) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEscaped(val)
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
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// returns a *int64 when successful
func (m *Label) GetId()(*int64) {
    return m.id
}
// GetName gets the name property value. The name of the label.
// returns a *string when successful
func (m *Label) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Label) GetNodeId()(*string) {
    return m.node_id
}
// GetUrl gets the url property value. URL for the label
// returns a *string when successful
func (m *Label) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Label) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("default", m.GetDefaultEscaped())
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
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
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
func (m *Label) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. 6-character hex code, without the leading #, identifying the color
func (m *Label) SetColor(value *string)() {
    m.color = value
}
// SetDefaultEscaped sets the default property value. The default property
func (m *Label) SetDefaultEscaped(value *bool)() {
    m.defaultEscaped = value
}
// SetDescription sets the description property value. The description property
func (m *Label) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The id property
func (m *Label) SetId(value *int64)() {
    m.id = value
}
// SetName sets the name property value. The name of the label.
func (m *Label) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Label) SetNodeId(value *string)() {
    m.node_id = value
}
// SetUrl sets the url property value. URL for the label
func (m *Label) SetUrl(value *string)() {
    m.url = value
}
type Labelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetDefaultEscaped()(*bool)
    GetDescription()(*string)
    GetId()(*int64)
    GetName()(*string)
    GetNodeId()(*string)
    GetUrl()(*string)
    SetColor(value *string)()
    SetDefaultEscaped(value *bool)()
    SetDescription(value *string)()
    SetId(value *int64)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetUrl(value *string)()
}
