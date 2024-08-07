package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelSearchResultItem label Search Result Item
type LabelSearchResultItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *string
    // The default property
    defaultEscaped *bool
    // The description property
    description *string
    // The id property
    id *int32
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The score property
    score *float64
    // The text_matches property
    text_matches []Labelsable
    // The url property
    url *string
}
// NewLabelSearchResultItem instantiates a new LabelSearchResultItem and sets the default values.
func NewLabelSearchResultItem()(*LabelSearchResultItem) {
    m := &LabelSearchResultItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelSearchResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelSearchResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelSearchResultItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LabelSearchResultItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
// returns a *string when successful
func (m *LabelSearchResultItem) GetColor()(*string) {
    return m.color
}
// GetDefaultEscaped gets the default property value. The default property
// returns a *bool when successful
func (m *LabelSearchResultItem) GetDefaultEscaped()(*bool) {
    return m.defaultEscaped
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *LabelSearchResultItem) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelSearchResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetInt32Value()
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
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["text_matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Labelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Labelsable)
                }
            }
            m.SetTextMatches(res)
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
func (m *LabelSearchResultItem) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *LabelSearchResultItem) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *LabelSearchResultItem) GetNodeId()(*string) {
    return m.node_id
}
// GetScore gets the score property value. The score property
// returns a *float64 when successful
func (m *LabelSearchResultItem) GetScore()(*float64) {
    return m.score
}
// GetTextMatches gets the text_matches property value. The text_matches property
// returns a []Labelsable when successful
func (m *LabelSearchResultItem) GetTextMatches()([]Labelsable) {
    return m.text_matches
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *LabelSearchResultItem) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *LabelSearchResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("id", m.GetId())
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
        err := writer.WriteFloat64Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    if m.GetTextMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTextMatches()))
        for i, v := range m.GetTextMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("text_matches", cast)
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
func (m *LabelSearchResultItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *LabelSearchResultItem) SetColor(value *string)() {
    m.color = value
}
// SetDefaultEscaped sets the default property value. The default property
func (m *LabelSearchResultItem) SetDefaultEscaped(value *bool)() {
    m.defaultEscaped = value
}
// SetDescription sets the description property value. The description property
func (m *LabelSearchResultItem) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The id property
func (m *LabelSearchResultItem) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *LabelSearchResultItem) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *LabelSearchResultItem) SetNodeId(value *string)() {
    m.node_id = value
}
// SetScore sets the score property value. The score property
func (m *LabelSearchResultItem) SetScore(value *float64)() {
    m.score = value
}
// SetTextMatches sets the text_matches property value. The text_matches property
func (m *LabelSearchResultItem) SetTextMatches(value []Labelsable)() {
    m.text_matches = value
}
// SetUrl sets the url property value. The url property
func (m *LabelSearchResultItem) SetUrl(value *string)() {
    m.url = value
}
type LabelSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetDefaultEscaped()(*bool)
    GetDescription()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetScore()(*float64)
    GetTextMatches()([]Labelsable)
    GetUrl()(*string)
    SetColor(value *string)()
    SetDefaultEscaped(value *bool)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetScore(value *float64)()
    SetTextMatches(value []Labelsable)()
    SetUrl(value *string)()
}
