package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BranchRestrictionPolicy_apps_permissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contents property
    contents *string
    // The issues property
    issues *string
    // The metadata property
    metadata *string
    // The single_file property
    single_file *string
}
// NewBranchRestrictionPolicy_apps_permissions instantiates a new BranchRestrictionPolicy_apps_permissions and sets the default values.
func NewBranchRestrictionPolicy_apps_permissions()(*BranchRestrictionPolicy_apps_permissions) {
    m := &BranchRestrictionPolicy_apps_permissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchRestrictionPolicy_apps_permissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchRestrictionPolicy_apps_permissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchRestrictionPolicy_apps_permissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchRestrictionPolicy_apps_permissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContents gets the contents property value. The contents property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_permissions) GetContents()(*string) {
    return m.contents
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchRestrictionPolicy_apps_permissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["single_file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleFile(val)
        }
        return nil
    }
    return res
}
// GetIssues gets the issues property value. The issues property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_permissions) GetIssues()(*string) {
    return m.issues
}
// GetMetadata gets the metadata property value. The metadata property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_permissions) GetMetadata()(*string) {
    return m.metadata
}
// GetSingleFile gets the single_file property value. The single_file property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_permissions) GetSingleFile()(*string) {
    return m.single_file
}
// Serialize serializes information the current object
func (m *BranchRestrictionPolicy_apps_permissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contents", m.GetContents())
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
        err := writer.WriteStringValue("single_file", m.GetSingleFile())
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
func (m *BranchRestrictionPolicy_apps_permissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContents sets the contents property value. The contents property
func (m *BranchRestrictionPolicy_apps_permissions) SetContents(value *string)() {
    m.contents = value
}
// SetIssues sets the issues property value. The issues property
func (m *BranchRestrictionPolicy_apps_permissions) SetIssues(value *string)() {
    m.issues = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *BranchRestrictionPolicy_apps_permissions) SetMetadata(value *string)() {
    m.metadata = value
}
// SetSingleFile sets the single_file property value. The single_file property
func (m *BranchRestrictionPolicy_apps_permissions) SetSingleFile(value *string)() {
    m.single_file = value
}
type BranchRestrictionPolicy_apps_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContents()(*string)
    GetIssues()(*string)
    GetMetadata()(*string)
    GetSingleFile()(*string)
    SetContents(value *string)()
    SetIssues(value *string)()
    SetMetadata(value *string)()
    SetSingleFile(value *string)()
}
