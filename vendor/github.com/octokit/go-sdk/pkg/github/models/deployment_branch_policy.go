package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentBranchPolicy details of a deployment branch or tag policy.
type DeploymentBranchPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the branch or tag policy.
    id *int32
    // The name pattern that branches or tags must match in order to deploy to the environment.
    name *string
    // The node_id property
    node_id *string
    // Whether this rule targets a branch or tag.
    typeEscaped *DeploymentBranchPolicy_type
}
// NewDeploymentBranchPolicy instantiates a new DeploymentBranchPolicy and sets the default values.
func NewDeploymentBranchPolicy()(*DeploymentBranchPolicy) {
    m := &DeploymentBranchPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentBranchPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentBranchPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentBranchPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeploymentBranchPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeploymentBranchPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentBranchPolicy_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*DeploymentBranchPolicy_type))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of the branch or tag policy.
// returns a *int32 when successful
func (m *DeploymentBranchPolicy) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name pattern that branches or tags must match in order to deploy to the environment.
// returns a *string when successful
func (m *DeploymentBranchPolicy) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *DeploymentBranchPolicy) GetNodeId()(*string) {
    return m.node_id
}
// GetTypeEscaped gets the type property value. Whether this rule targets a branch or tag.
// returns a *DeploymentBranchPolicy_type when successful
func (m *DeploymentBranchPolicy) GetTypeEscaped()(*DeploymentBranchPolicy_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *DeploymentBranchPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *DeploymentBranchPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of the branch or tag policy.
func (m *DeploymentBranchPolicy) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name pattern that branches or tags must match in order to deploy to the environment.
func (m *DeploymentBranchPolicy) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *DeploymentBranchPolicy) SetNodeId(value *string)() {
    m.node_id = value
}
// SetTypeEscaped sets the type property value. Whether this rule targets a branch or tag.
func (m *DeploymentBranchPolicy) SetTypeEscaped(value *DeploymentBranchPolicy_type)() {
    m.typeEscaped = value
}
type DeploymentBranchPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetTypeEscaped()(*DeploymentBranchPolicy_type)
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetTypeEscaped(value *DeploymentBranchPolicy_type)()
}
