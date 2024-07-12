package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Autolink an autolink reference.
type Autolink struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
    is_alphanumeric *bool
    // The prefix of a key that is linkified.
    key_prefix *string
    // A template for the target URL that is generated if a key was found.
    url_template *string
}
// NewAutolink instantiates a new Autolink and sets the default values.
func NewAutolink()(*Autolink) {
    m := &Autolink{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAutolinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutolinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutolink(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Autolink) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Autolink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["is_alphanumeric"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAlphanumeric(val)
        }
        return nil
    }
    res["key_prefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPrefix(val)
        }
        return nil
    }
    res["url_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrlTemplate(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Autolink) GetId()(*int32) {
    return m.id
}
// GetIsAlphanumeric gets the is_alphanumeric property value. Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
// returns a *bool when successful
func (m *Autolink) GetIsAlphanumeric()(*bool) {
    return m.is_alphanumeric
}
// GetKeyPrefix gets the key_prefix property value. The prefix of a key that is linkified.
// returns a *string when successful
func (m *Autolink) GetKeyPrefix()(*string) {
    return m.key_prefix
}
// GetUrlTemplate gets the url_template property value. A template for the target URL that is generated if a key was found.
// returns a *string when successful
func (m *Autolink) GetUrlTemplate()(*string) {
    return m.url_template
}
// Serialize serializes information the current object
func (m *Autolink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_alphanumeric", m.GetIsAlphanumeric())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_prefix", m.GetKeyPrefix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url_template", m.GetUrlTemplate())
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
func (m *Autolink) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *Autolink) SetId(value *int32)() {
    m.id = value
}
// SetIsAlphanumeric sets the is_alphanumeric property value. Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
func (m *Autolink) SetIsAlphanumeric(value *bool)() {
    m.is_alphanumeric = value
}
// SetKeyPrefix sets the key_prefix property value. The prefix of a key that is linkified.
func (m *Autolink) SetKeyPrefix(value *string)() {
    m.key_prefix = value
}
// SetUrlTemplate sets the url_template property value. A template for the target URL that is generated if a key was found.
func (m *Autolink) SetUrlTemplate(value *string)() {
    m.url_template = value
}
type Autolinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetIsAlphanumeric()(*bool)
    GetKeyPrefix()(*string)
    GetUrlTemplate()(*string)
    SetId(value *int32)()
    SetIsAlphanumeric(value *bool)()
    SetKeyPrefix(value *string)()
    SetUrlTemplate(value *string)()
}
