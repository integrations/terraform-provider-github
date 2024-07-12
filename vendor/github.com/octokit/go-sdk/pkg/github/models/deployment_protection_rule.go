package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentProtectionRule deployment protection rule
type DeploymentProtectionRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub App that is providing a custom deployment protection rule.
    app CustomDeploymentRuleAppable
    // Whether the deployment protection rule is enabled for the environment.
    enabled *bool
    // The unique identifier for the deployment protection rule.
    id *int32
    // The node ID for the deployment protection rule.
    node_id *string
}
// NewDeploymentProtectionRule instantiates a new DeploymentProtectionRule and sets the default values.
func NewDeploymentProtectionRule()(*DeploymentProtectionRule) {
    m := &DeploymentProtectionRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentProtectionRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentProtectionRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentProtectionRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeploymentProtectionRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApp gets the app property value. A GitHub App that is providing a custom deployment protection rule.
// returns a CustomDeploymentRuleAppable when successful
func (m *DeploymentProtectionRule) GetApp()(CustomDeploymentRuleAppable) {
    return m.app
}
// GetEnabled gets the enabled property value. Whether the deployment protection rule is enabled for the environment.
// returns a *bool when successful
func (m *DeploymentProtectionRule) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeploymentProtectionRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomDeploymentRuleAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(CustomDeploymentRuleAppable))
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
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
    return res
}
// GetId gets the id property value. The unique identifier for the deployment protection rule.
// returns a *int32 when successful
func (m *DeploymentProtectionRule) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node ID for the deployment protection rule.
// returns a *string when successful
func (m *DeploymentProtectionRule) GetNodeId()(*string) {
    return m.node_id
}
// Serialize serializes information the current object
func (m *DeploymentProtectionRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentProtectionRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApp sets the app property value. A GitHub App that is providing a custom deployment protection rule.
func (m *DeploymentProtectionRule) SetApp(value CustomDeploymentRuleAppable)() {
    m.app = value
}
// SetEnabled sets the enabled property value. Whether the deployment protection rule is enabled for the environment.
func (m *DeploymentProtectionRule) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetId sets the id property value. The unique identifier for the deployment protection rule.
func (m *DeploymentProtectionRule) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node ID for the deployment protection rule.
func (m *DeploymentProtectionRule) SetNodeId(value *string)() {
    m.node_id = value
}
type DeploymentProtectionRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(CustomDeploymentRuleAppable)
    GetEnabled()(*bool)
    GetId()(*int32)
    GetNodeId()(*string)
    SetApp(value CustomDeploymentRuleAppable)()
    SetEnabled(value *bool)()
    SetId(value *int32)()
    SetNodeId(value *string)()
}
