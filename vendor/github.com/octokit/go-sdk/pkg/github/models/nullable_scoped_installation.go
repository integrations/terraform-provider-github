package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NullableScopedInstallation struct {
    // A GitHub user.
    account SimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The has_multiple_single_files property
    has_multiple_single_files *bool
    // The permissions granted to the user access token.
    permissions AppPermissionsable
    // The repositories_url property
    repositories_url *string
    // Describe whether all repositories have been selected or there's a selection involved
    repository_selection *NullableScopedInstallation_repository_selection
    // The single_file_name property
    single_file_name *string
    // The single_file_paths property
    single_file_paths []string
}
// NewNullableScopedInstallation instantiates a new NullableScopedInstallation and sets the default values.
func NewNullableScopedInstallation()(*NullableScopedInstallation) {
    m := &NullableScopedInstallation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableScopedInstallationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableScopedInstallationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableScopedInstallation(), nil
}
// GetAccount gets the account property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *NullableScopedInstallation) GetAccount()(SimpleUserable) {
    return m.account
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableScopedInstallation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableScopedInstallation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(SimpleUserable))
        }
        return nil
    }
    res["has_multiple_single_files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMultipleSingleFiles(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppPermissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(AppPermissionsable))
        }
        return nil
    }
    res["repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoriesUrl(val)
        }
        return nil
    }
    res["repository_selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNullableScopedInstallation_repository_selection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val.(*NullableScopedInstallation_repository_selection))
        }
        return nil
    }
    res["single_file_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleFileName(val)
        }
        return nil
    }
    res["single_file_paths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSingleFilePaths(res)
        }
        return nil
    }
    return res
}
// GetHasMultipleSingleFiles gets the has_multiple_single_files property value. The has_multiple_single_files property
// returns a *bool when successful
func (m *NullableScopedInstallation) GetHasMultipleSingleFiles()(*bool) {
    return m.has_multiple_single_files
}
// GetPermissions gets the permissions property value. The permissions granted to the user access token.
// returns a AppPermissionsable when successful
func (m *NullableScopedInstallation) GetPermissions()(AppPermissionsable) {
    return m.permissions
}
// GetRepositoriesUrl gets the repositories_url property value. The repositories_url property
// returns a *string when successful
func (m *NullableScopedInstallation) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetRepositorySelection gets the repository_selection property value. Describe whether all repositories have been selected or there's a selection involved
// returns a *NullableScopedInstallation_repository_selection when successful
func (m *NullableScopedInstallation) GetRepositorySelection()(*NullableScopedInstallation_repository_selection) {
    return m.repository_selection
}
// GetSingleFileName gets the single_file_name property value. The single_file_name property
// returns a *string when successful
func (m *NullableScopedInstallation) GetSingleFileName()(*string) {
    return m.single_file_name
}
// GetSingleFilePaths gets the single_file_paths property value. The single_file_paths property
// returns a []string when successful
func (m *NullableScopedInstallation) GetSingleFilePaths()([]string) {
    return m.single_file_paths
}
// Serialize serializes information the current object
func (m *NullableScopedInstallation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_multiple_single_files", m.GetHasMultipleSingleFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositories_url", m.GetRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRepositorySelection() != nil {
        cast := (*m.GetRepositorySelection()).String()
        err := writer.WriteStringValue("repository_selection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("single_file_name", m.GetSingleFileName())
        if err != nil {
            return err
        }
    }
    if m.GetSingleFilePaths() != nil {
        err := writer.WriteCollectionOfStringValues("single_file_paths", m.GetSingleFilePaths())
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
// SetAccount sets the account property value. A GitHub user.
func (m *NullableScopedInstallation) SetAccount(value SimpleUserable)() {
    m.account = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NullableScopedInstallation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHasMultipleSingleFiles sets the has_multiple_single_files property value. The has_multiple_single_files property
func (m *NullableScopedInstallation) SetHasMultipleSingleFiles(value *bool)() {
    m.has_multiple_single_files = value
}
// SetPermissions sets the permissions property value. The permissions granted to the user access token.
func (m *NullableScopedInstallation) SetPermissions(value AppPermissionsable)() {
    m.permissions = value
}
// SetRepositoriesUrl sets the repositories_url property value. The repositories_url property
func (m *NullableScopedInstallation) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetRepositorySelection sets the repository_selection property value. Describe whether all repositories have been selected or there's a selection involved
func (m *NullableScopedInstallation) SetRepositorySelection(value *NullableScopedInstallation_repository_selection)() {
    m.repository_selection = value
}
// SetSingleFileName sets the single_file_name property value. The single_file_name property
func (m *NullableScopedInstallation) SetSingleFileName(value *string)() {
    m.single_file_name = value
}
// SetSingleFilePaths sets the single_file_paths property value. The single_file_paths property
func (m *NullableScopedInstallation) SetSingleFilePaths(value []string)() {
    m.single_file_paths = value
}
type NullableScopedInstallationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(SimpleUserable)
    GetHasMultipleSingleFiles()(*bool)
    GetPermissions()(AppPermissionsable)
    GetRepositoriesUrl()(*string)
    GetRepositorySelection()(*NullableScopedInstallation_repository_selection)
    GetSingleFileName()(*string)
    GetSingleFilePaths()([]string)
    SetAccount(value SimpleUserable)()
    SetHasMultipleSingleFiles(value *bool)()
    SetPermissions(value AppPermissionsable)()
    SetRepositoriesUrl(value *string)()
    SetRepositorySelection(value *NullableScopedInstallation_repository_selection)()
    SetSingleFileName(value *string)()
    SetSingleFilePaths(value []string)()
}
