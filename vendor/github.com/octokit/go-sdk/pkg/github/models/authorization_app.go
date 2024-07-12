package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Authorization_app struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The client_id property
    client_id *string
    // The name property
    name *string
    // The url property
    url *string
}
// NewAuthorization_app instantiates a new Authorization_app and sets the default values.
func NewAuthorization_app()(*Authorization_app) {
    m := &Authorization_app{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthorization_appFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorization_appFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorization_app(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Authorization_app) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the client_id property value. The client_id property
// returns a *string when successful
func (m *Authorization_app) GetClientId()(*string) {
    return m.client_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Authorization_app) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
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
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Authorization_app) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Authorization_app) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Authorization_app) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_id", m.GetClientId())
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
func (m *Authorization_app) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the client_id property value. The client_id property
func (m *Authorization_app) SetClientId(value *string)() {
    m.client_id = value
}
// SetName sets the name property value. The name property
func (m *Authorization_app) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The url property
func (m *Authorization_app) SetUrl(value *string)() {
    m.url = value
}
type Authorization_appable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetName()(*string)
    GetUrl()(*string)
    SetClientId(value *string)()
    SetName(value *string)()
    SetUrl(value *string)()
}
