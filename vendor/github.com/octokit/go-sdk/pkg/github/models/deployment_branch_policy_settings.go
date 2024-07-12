package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentBranchPolicySettings the type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
type DeploymentBranchPolicySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether only branches that match the specified name patterns can deploy to this environment.  If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
    custom_branch_policies *bool
    // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
    protected_branches *bool
}
// NewDeploymentBranchPolicySettings instantiates a new DeploymentBranchPolicySettings and sets the default values.
func NewDeploymentBranchPolicySettings()(*DeploymentBranchPolicySettings) {
    m := &DeploymentBranchPolicySettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentBranchPolicySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentBranchPolicySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentBranchPolicySettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeploymentBranchPolicySettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomBranchPolicies gets the custom_branch_policies property value. Whether only branches that match the specified name patterns can deploy to this environment.  If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
// returns a *bool when successful
func (m *DeploymentBranchPolicySettings) GetCustomBranchPolicies()(*bool) {
    return m.custom_branch_policies
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeploymentBranchPolicySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custom_branch_policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBranchPolicies(val)
        }
        return nil
    }
    res["protected_branches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectedBranches(val)
        }
        return nil
    }
    return res
}
// GetProtectedBranches gets the protected_branches property value. Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
// returns a *bool when successful
func (m *DeploymentBranchPolicySettings) GetProtectedBranches()(*bool) {
    return m.protected_branches
}
// Serialize serializes information the current object
func (m *DeploymentBranchPolicySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("custom_branch_policies", m.GetCustomBranchPolicies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("protected_branches", m.GetProtectedBranches())
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
func (m *DeploymentBranchPolicySettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomBranchPolicies sets the custom_branch_policies property value. Whether only branches that match the specified name patterns can deploy to this environment.  If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
func (m *DeploymentBranchPolicySettings) SetCustomBranchPolicies(value *bool)() {
    m.custom_branch_policies = value
}
// SetProtectedBranches sets the protected_branches property value. Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
func (m *DeploymentBranchPolicySettings) SetProtectedBranches(value *bool)() {
    m.protected_branches = value
}
type DeploymentBranchPolicySettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomBranchPolicies()(*bool)
    GetProtectedBranches()(*bool)
    SetCustomBranchPolicies(value *bool)()
    SetProtectedBranches(value *bool)()
}
