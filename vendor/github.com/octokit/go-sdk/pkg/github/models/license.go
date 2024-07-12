package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// License license
type License struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The body property
    body *string
    // The conditions property
    conditions []string
    // The description property
    description *string
    // The featured property
    featured *bool
    // The html_url property
    html_url *string
    // The implementation property
    implementation *string
    // The key property
    key *string
    // The limitations property
    limitations []string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The permissions property
    permissions []string
    // The spdx_id property
    spdx_id *string
    // The url property
    url *string
}
// NewLicense instantiates a new License and sets the default values.
func NewLicense()(*License) {
    m := &License{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLicenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLicenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicense(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *License) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *License) GetBody()(*string) {
    return m.body
}
// GetConditions gets the conditions property value. The conditions property
// returns a []string when successful
func (m *License) GetConditions()([]string) {
    return m.conditions
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *License) GetDescription()(*string) {
    return m.description
}
// GetFeatured gets the featured property value. The featured property
// returns a *bool when successful
func (m *License) GetFeatured()(*bool) {
    return m.featured
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *License) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetConditions(res)
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
    res["featured"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatured(val)
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
    res["implementation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplementation(val)
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
    res["limitations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetLimitations(res)
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
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPermissions(res)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *License) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetImplementation gets the implementation property value. The implementation property
// returns a *string when successful
func (m *License) GetImplementation()(*string) {
    return m.implementation
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *License) GetKey()(*string) {
    return m.key
}
// GetLimitations gets the limitations property value. The limitations property
// returns a []string when successful
func (m *License) GetLimitations()([]string) {
    return m.limitations
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *License) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *License) GetNodeId()(*string) {
    return m.node_id
}
// GetPermissions gets the permissions property value. The permissions property
// returns a []string when successful
func (m *License) GetPermissions()([]string) {
    return m.permissions
}
// GetSpdxId gets the spdx_id property value. The spdx_id property
// returns a *string when successful
func (m *License) GetSpdxId()(*string) {
    return m.spdx_id
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *License) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *License) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetConditions() != nil {
        err := writer.WriteCollectionOfStringValues("conditions", m.GetConditions())
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
        err := writer.WriteBoolValue("featured", m.GetFeatured())
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
        err := writer.WriteStringValue("implementation", m.GetImplementation())
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
    if m.GetLimitations() != nil {
        err := writer.WriteCollectionOfStringValues("limitations", m.GetLimitations())
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
    if m.GetPermissions() != nil {
        err := writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
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
func (m *License) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The body property
func (m *License) SetBody(value *string)() {
    m.body = value
}
// SetConditions sets the conditions property value. The conditions property
func (m *License) SetConditions(value []string)() {
    m.conditions = value
}
// SetDescription sets the description property value. The description property
func (m *License) SetDescription(value *string)() {
    m.description = value
}
// SetFeatured sets the featured property value. The featured property
func (m *License) SetFeatured(value *bool)() {
    m.featured = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *License) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetImplementation sets the implementation property value. The implementation property
func (m *License) SetImplementation(value *string)() {
    m.implementation = value
}
// SetKey sets the key property value. The key property
func (m *License) SetKey(value *string)() {
    m.key = value
}
// SetLimitations sets the limitations property value. The limitations property
func (m *License) SetLimitations(value []string)() {
    m.limitations = value
}
// SetName sets the name property value. The name property
func (m *License) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *License) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *License) SetPermissions(value []string)() {
    m.permissions = value
}
// SetSpdxId sets the spdx_id property value. The spdx_id property
func (m *License) SetSpdxId(value *string)() {
    m.spdx_id = value
}
// SetUrl sets the url property value. The url property
func (m *License) SetUrl(value *string)() {
    m.url = value
}
type Licenseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetConditions()([]string)
    GetDescription()(*string)
    GetFeatured()(*bool)
    GetHtmlUrl()(*string)
    GetImplementation()(*string)
    GetKey()(*string)
    GetLimitations()([]string)
    GetName()(*string)
    GetNodeId()(*string)
    GetPermissions()([]string)
    GetSpdxId()(*string)
    GetUrl()(*string)
    SetBody(value *string)()
    SetConditions(value []string)()
    SetDescription(value *string)()
    SetFeatured(value *bool)()
    SetHtmlUrl(value *string)()
    SetImplementation(value *string)()
    SetKey(value *string)()
    SetLimitations(value []string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetPermissions(value []string)()
    SetSpdxId(value *string)()
    SetUrl(value *string)()
}
