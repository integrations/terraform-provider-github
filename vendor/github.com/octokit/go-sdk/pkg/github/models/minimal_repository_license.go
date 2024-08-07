package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MinimalRepository_license struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The key property
    key *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The spdx_id property
    spdx_id *string
    // The url property
    url *string
}
// NewMinimalRepository_license instantiates a new MinimalRepository_license and sets the default values.
func NewMinimalRepository_license()(*MinimalRepository_license) {
    m := &MinimalRepository_license{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMinimalRepository_licenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMinimalRepository_licenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMinimalRepository_license(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MinimalRepository_license) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MinimalRepository_license) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["spdx_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpdxId(val)
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
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *MinimalRepository_license) GetKey()(*string) {
    return m.key
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *MinimalRepository_license) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *MinimalRepository_license) GetNodeId()(*string) {
    return m.node_id
}
// GetSpdxId gets the spdx_id property value. The spdx_id property
// returns a *string when successful
func (m *MinimalRepository_license) GetSpdxId()(*string) {
    return m.spdx_id
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *MinimalRepository_license) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *MinimalRepository_license) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("spdx_id", m.GetSpdxId())
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
func (m *MinimalRepository_license) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *MinimalRepository_license) SetKey(value *string)() {
    m.key = value
}
// SetName sets the name property value. The name property
func (m *MinimalRepository_license) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *MinimalRepository_license) SetNodeId(value *string)() {
    m.node_id = value
}
// SetSpdxId sets the spdx_id property value. The spdx_id property
func (m *MinimalRepository_license) SetSpdxId(value *string)() {
    m.spdx_id = value
}
// SetUrl sets the url property value. The url property
func (m *MinimalRepository_license) SetUrl(value *string)() {
    m.url = value
}
type MinimalRepository_licenseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetSpdxId()(*string)
    GetUrl()(*string)
    SetKey(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetSpdxId(value *string)()
    SetUrl(value *string)()
}
