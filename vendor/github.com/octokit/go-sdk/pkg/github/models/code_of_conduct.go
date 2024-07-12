package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeOfConduct code Of Conduct
type CodeOfConduct struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The body property
    body *string
    // The html_url property
    html_url *string
    // The key property
    key *string
    // The name property
    name *string
    // The url property
    url *string
}
// NewCodeOfConduct instantiates a new CodeOfConduct and sets the default values.
func NewCodeOfConduct()(*CodeOfConduct) {
    m := &CodeOfConduct{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeOfConductFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeOfConductFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeOfConduct(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeOfConduct) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *CodeOfConduct) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeOfConduct) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *CodeOfConduct) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *CodeOfConduct) GetKey()(*string) {
    return m.key
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CodeOfConduct) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CodeOfConduct) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeOfConduct) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *CodeOfConduct) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The body property
func (m *CodeOfConduct) SetBody(value *string)() {
    m.body = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *CodeOfConduct) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetKey sets the key property value. The key property
func (m *CodeOfConduct) SetKey(value *string)() {
    m.key = value
}
// SetName sets the name property value. The name property
func (m *CodeOfConduct) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The url property
func (m *CodeOfConduct) SetUrl(value *string)() {
    m.url = value
}
type CodeOfConductable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetHtmlUrl()(*string)
    GetKey()(*string)
    GetName()(*string)
    GetUrl()(*string)
    SetBody(value *string)()
    SetHtmlUrl(value *string)()
    SetKey(value *string)()
    SetName(value *string)()
    SetUrl(value *string)()
}
