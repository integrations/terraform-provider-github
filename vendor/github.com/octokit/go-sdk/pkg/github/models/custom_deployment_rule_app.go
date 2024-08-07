package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomDeploymentRuleApp a GitHub App that is providing a custom deployment protection rule.
type CustomDeploymentRuleApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the deployment protection rule integration.
    id *int32
    // The URL for the endpoint to get details about the app.
    integration_url *string
    // The node ID for the deployment protection rule integration.
    node_id *string
    // The slugified name of the deployment protection rule integration.
    slug *string
}
// NewCustomDeploymentRuleApp instantiates a new CustomDeploymentRuleApp and sets the default values.
func NewCustomDeploymentRuleApp()(*CustomDeploymentRuleApp) {
    m := &CustomDeploymentRuleApp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomDeploymentRuleAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomDeploymentRuleAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomDeploymentRuleApp(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomDeploymentRuleApp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomDeploymentRuleApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["integration_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationUrl(val)
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
    res["slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSlug(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of the deployment protection rule integration.
// returns a *int32 when successful
func (m *CustomDeploymentRuleApp) GetId()(*int32) {
    return m.id
}
// GetIntegrationUrl gets the integration_url property value. The URL for the endpoint to get details about the app.
// returns a *string when successful
func (m *CustomDeploymentRuleApp) GetIntegrationUrl()(*string) {
    return m.integration_url
}
// GetNodeId gets the node_id property value. The node ID for the deployment protection rule integration.
// returns a *string when successful
func (m *CustomDeploymentRuleApp) GetNodeId()(*string) {
    return m.node_id
}
// GetSlug gets the slug property value. The slugified name of the deployment protection rule integration.
// returns a *string when successful
func (m *CustomDeploymentRuleApp) GetSlug()(*string) {
    return m.slug
}
// Serialize serializes information the current object
func (m *CustomDeploymentRuleApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("integration_url", m.GetIntegrationUrl())
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
        err := writer.WriteStringValue("slug", m.GetSlug())
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
func (m *CustomDeploymentRuleApp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of the deployment protection rule integration.
func (m *CustomDeploymentRuleApp) SetId(value *int32)() {
    m.id = value
}
// SetIntegrationUrl sets the integration_url property value. The URL for the endpoint to get details about the app.
func (m *CustomDeploymentRuleApp) SetIntegrationUrl(value *string)() {
    m.integration_url = value
}
// SetNodeId sets the node_id property value. The node ID for the deployment protection rule integration.
func (m *CustomDeploymentRuleApp) SetNodeId(value *string)() {
    m.node_id = value
}
// SetSlug sets the slug property value. The slugified name of the deployment protection rule integration.
func (m *CustomDeploymentRuleApp) SetSlug(value *string)() {
    m.slug = value
}
type CustomDeploymentRuleAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetIntegrationUrl()(*string)
    GetNodeId()(*string)
    GetSlug()(*string)
    SetId(value *int32)()
    SetIntegrationUrl(value *string)()
    SetNodeId(value *string)()
    SetSlug(value *string)()
}
