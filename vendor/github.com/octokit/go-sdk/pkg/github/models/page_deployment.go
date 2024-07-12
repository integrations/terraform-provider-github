package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageDeployment the GitHub Pages deployment status.
type PageDeployment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
    id PageDeployment_PageDeployment_idable
    // The URI to the deployed GitHub Pages.
    page_url *string
    // The URI to the deployed GitHub Pages preview.
    preview_url *string
    // The URI to monitor GitHub Pages deployment status.
    status_url *string
}
// PageDeployment_PageDeployment_id composed type wrapper for classes int32, string
type PageDeployment_PageDeployment_id struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewPageDeployment_PageDeployment_id instantiates a new PageDeployment_PageDeployment_id and sets the default values.
func NewPageDeployment_PageDeployment_id()(*PageDeployment_PageDeployment_id) {
    m := &PageDeployment_PageDeployment_id{
    }
    return m
}
// CreatePageDeployment_PageDeployment_idFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePageDeployment_PageDeployment_idFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewPageDeployment_PageDeployment_id()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PageDeployment_PageDeployment_id) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *PageDeployment_PageDeployment_id) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *PageDeployment_PageDeployment_id) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *PageDeployment_PageDeployment_id) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *PageDeployment_PageDeployment_id) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type int32
func (m *PageDeployment_PageDeployment_id) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *PageDeployment_PageDeployment_id) SetString(value *string)() {
    m.string = value
}
type PageDeployment_PageDeployment_idable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewPageDeployment instantiates a new PageDeployment and sets the default values.
func NewPageDeployment()(*PageDeployment) {
    m := &PageDeployment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePageDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePageDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPageDeployment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PageDeployment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PageDeployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePageDeployment_PageDeployment_idFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val.(PageDeployment_PageDeployment_idable))
        }
        return nil
    }
    res["page_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageUrl(val)
        }
        return nil
    }
    res["preview_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewUrl(val)
        }
        return nil
    }
    res["status_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
// returns a PageDeployment_PageDeployment_idable when successful
func (m *PageDeployment) GetId()(PageDeployment_PageDeployment_idable) {
    return m.id
}
// GetPageUrl gets the page_url property value. The URI to the deployed GitHub Pages.
// returns a *string when successful
func (m *PageDeployment) GetPageUrl()(*string) {
    return m.page_url
}
// GetPreviewUrl gets the preview_url property value. The URI to the deployed GitHub Pages preview.
// returns a *string when successful
func (m *PageDeployment) GetPreviewUrl()(*string) {
    return m.preview_url
}
// GetStatusUrl gets the status_url property value. The URI to monitor GitHub Pages deployment status.
// returns a *string when successful
func (m *PageDeployment) GetStatusUrl()(*string) {
    return m.status_url
}
// Serialize serializes information the current object
func (m *PageDeployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("page_url", m.GetPageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("preview_url", m.GetPreviewUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status_url", m.GetStatusUrl())
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
func (m *PageDeployment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
func (m *PageDeployment) SetId(value PageDeployment_PageDeployment_idable)() {
    m.id = value
}
// SetPageUrl sets the page_url property value. The URI to the deployed GitHub Pages.
func (m *PageDeployment) SetPageUrl(value *string)() {
    m.page_url = value
}
// SetPreviewUrl sets the preview_url property value. The URI to the deployed GitHub Pages preview.
func (m *PageDeployment) SetPreviewUrl(value *string)() {
    m.preview_url = value
}
// SetStatusUrl sets the status_url property value. The URI to monitor GitHub Pages deployment status.
func (m *PageDeployment) SetStatusUrl(value *string)() {
    m.status_url = value
}
type PageDeploymentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(PageDeployment_PageDeployment_idable)
    GetPageUrl()(*string)
    GetPreviewUrl()(*string)
    GetStatusUrl()(*string)
    SetId(value PageDeployment_PageDeployment_idable)()
    SetPageUrl(value *string)()
    SetPreviewUrl(value *string)()
    SetStatusUrl(value *string)()
}
