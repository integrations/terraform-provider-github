package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableIntegration_permissions the set of permissions for the GitHub app
type NullableIntegration_permissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The checks property
    checks *string
    // The contents property
    contents *string
    // The deployments property
    deployments *string
    // The issues property
    issues *string
    // The metadata property
    metadata *string
}
// NewNullableIntegration_permissions instantiates a new NullableIntegration_permissions and sets the default values.
func NewNullableIntegration_permissions()(*NullableIntegration_permissions) {
    m := &NullableIntegration_permissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableIntegration_permissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableIntegration_permissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableIntegration_permissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableIntegration_permissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChecks gets the checks property value. The checks property
// returns a *string when successful
func (m *NullableIntegration_permissions) GetChecks()(*string) {
    return m.checks
}
// GetContents gets the contents property value. The contents property
// returns a *string when successful
func (m *NullableIntegration_permissions) GetContents()(*string) {
    return m.contents
}
// GetDeployments gets the deployments property value. The deployments property
// returns a *string when successful
func (m *NullableIntegration_permissions) GetDeployments()(*string) {
    return m.deployments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableIntegration_permissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecks(val)
        }
        return nil
    }
    res["contents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContents(val)
        }
        return nil
    }
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployments(val)
        }
        return nil
    }
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssues(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    return res
}
// GetIssues gets the issues property value. The issues property
// returns a *string when successful
func (m *NullableIntegration_permissions) GetIssues()(*string) {
    return m.issues
}
// GetMetadata gets the metadata property value. The metadata property
// returns a *string when successful
func (m *NullableIntegration_permissions) GetMetadata()(*string) {
    return m.metadata
}
// Serialize serializes information the current object
func (m *NullableIntegration_permissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("checks", m.GetChecks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contents", m.GetContents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployments", m.GetDeployments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issues", m.GetIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("metadata", m.GetMetadata())
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
func (m *NullableIntegration_permissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChecks sets the checks property value. The checks property
func (m *NullableIntegration_permissions) SetChecks(value *string)() {
    m.checks = value
}
// SetContents sets the contents property value. The contents property
func (m *NullableIntegration_permissions) SetContents(value *string)() {
    m.contents = value
}
// SetDeployments sets the deployments property value. The deployments property
func (m *NullableIntegration_permissions) SetDeployments(value *string)() {
    m.deployments = value
}
// SetIssues sets the issues property value. The issues property
func (m *NullableIntegration_permissions) SetIssues(value *string)() {
    m.issues = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *NullableIntegration_permissions) SetMetadata(value *string)() {
    m.metadata = value
}
type NullableIntegration_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()(*string)
    GetContents()(*string)
    GetDeployments()(*string)
    GetIssues()(*string)
    GetMetadata()(*string)
    SetChecks(value *string)()
    SetContents(value *string)()
    SetDeployments(value *string)()
    SetIssues(value *string)()
    SetMetadata(value *string)()
}
