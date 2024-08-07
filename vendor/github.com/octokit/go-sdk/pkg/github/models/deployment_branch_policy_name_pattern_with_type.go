package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeploymentBranchPolicyNamePatternWithType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name pattern that branches or tags must match in order to deploy to the environment.Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`.For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
    name *string
    // Whether this rule targets a branch or tag
    typeEscaped *DeploymentBranchPolicyNamePatternWithType_type
}
// NewDeploymentBranchPolicyNamePatternWithType instantiates a new DeploymentBranchPolicyNamePatternWithType and sets the default values.
func NewDeploymentBranchPolicyNamePatternWithType()(*DeploymentBranchPolicyNamePatternWithType) {
    m := &DeploymentBranchPolicyNamePatternWithType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentBranchPolicyNamePatternWithTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentBranchPolicyNamePatternWithTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentBranchPolicyNamePatternWithType(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeploymentBranchPolicyNamePatternWithType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeploymentBranchPolicyNamePatternWithType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentBranchPolicyNamePatternWithType_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*DeploymentBranchPolicyNamePatternWithType_type))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name pattern that branches or tags must match in order to deploy to the environment.Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`.For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
// returns a *string when successful
func (m *DeploymentBranchPolicyNamePatternWithType) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. Whether this rule targets a branch or tag
// returns a *DeploymentBranchPolicyNamePatternWithType_type when successful
func (m *DeploymentBranchPolicyNamePatternWithType) GetTypeEscaped()(*DeploymentBranchPolicyNamePatternWithType_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *DeploymentBranchPolicyNamePatternWithType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *DeploymentBranchPolicyNamePatternWithType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name pattern that branches or tags must match in order to deploy to the environment.Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`.For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
func (m *DeploymentBranchPolicyNamePatternWithType) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. Whether this rule targets a branch or tag
func (m *DeploymentBranchPolicyNamePatternWithType) SetTypeEscaped(value *DeploymentBranchPolicyNamePatternWithType_type)() {
    m.typeEscaped = value
}
type DeploymentBranchPolicyNamePatternWithTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetTypeEscaped()(*DeploymentBranchPolicyNamePatternWithType_type)
    SetName(value *string)()
    SetTypeEscaped(value *DeploymentBranchPolicyNamePatternWithType_type)()
}
