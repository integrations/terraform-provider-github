package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InstallationToken authentication token for a GitHub App installed on a user or org.
type InstallationToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The expires_at property
    expires_at *string
    // The has_multiple_single_files property
    has_multiple_single_files *bool
    // The permissions granted to the user access token.
    permissions AppPermissionsable
    // The repositories property
    repositories []Repositoryable
    // The repository_selection property
    repository_selection *InstallationToken_repository_selection
    // The single_file property
    single_file *string
    // The single_file_paths property
    single_file_paths []string
    // The token property
    token *string
}
// NewInstallationToken instantiates a new InstallationToken and sets the default values.
func NewInstallationToken()(*InstallationToken) {
    m := &InstallationToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstallationTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstallationTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallationToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InstallationToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresAt gets the expires_at property value. The expires_at property
// returns a *string when successful
func (m *InstallationToken) GetExpiresAt()(*string) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InstallationToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
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
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Repositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Repositoryable)
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
    res["repository_selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallationToken_repository_selection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val.(*InstallationToken_repository_selection))
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
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetHasMultipleSingleFiles gets the has_multiple_single_files property value. The has_multiple_single_files property
// returns a *bool when successful
func (m *InstallationToken) GetHasMultipleSingleFiles()(*bool) {
    return m.has_multiple_single_files
}
// GetPermissions gets the permissions property value. The permissions granted to the user access token.
// returns a AppPermissionsable when successful
func (m *InstallationToken) GetPermissions()(AppPermissionsable) {
    return m.permissions
}
// GetRepositories gets the repositories property value. The repositories property
// returns a []Repositoryable when successful
func (m *InstallationToken) GetRepositories()([]Repositoryable) {
    return m.repositories
}
// GetRepositorySelection gets the repository_selection property value. The repository_selection property
// returns a *InstallationToken_repository_selection when successful
func (m *InstallationToken) GetRepositorySelection()(*InstallationToken_repository_selection) {
    return m.repository_selection
}
// GetSingleFile gets the single_file property value. The single_file property
// returns a *string when successful
func (m *InstallationToken) GetSingleFile()(*string) {
    return m.single_file
}
// GetSingleFilePaths gets the single_file_paths property value. The single_file_paths property
// returns a []string when successful
func (m *InstallationToken) GetSingleFilePaths()([]string) {
    return m.single_file_paths
}
// GetToken gets the token property value. The token property
// returns a *string when successful
func (m *InstallationToken) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *InstallationToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expires_at", m.GetExpiresAt())
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
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
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
        err := writer.WriteStringValue("single_file", m.GetSingleFile())
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
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *InstallationToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresAt sets the expires_at property value. The expires_at property
func (m *InstallationToken) SetExpiresAt(value *string)() {
    m.expires_at = value
}
// SetHasMultipleSingleFiles sets the has_multiple_single_files property value. The has_multiple_single_files property
func (m *InstallationToken) SetHasMultipleSingleFiles(value *bool)() {
    m.has_multiple_single_files = value
}
// SetPermissions sets the permissions property value. The permissions granted to the user access token.
func (m *InstallationToken) SetPermissions(value AppPermissionsable)() {
    m.permissions = value
}
// SetRepositories sets the repositories property value. The repositories property
func (m *InstallationToken) SetRepositories(value []Repositoryable)() {
    m.repositories = value
}
// SetRepositorySelection sets the repository_selection property value. The repository_selection property
func (m *InstallationToken) SetRepositorySelection(value *InstallationToken_repository_selection)() {
    m.repository_selection = value
}
// SetSingleFile sets the single_file property value. The single_file property
func (m *InstallationToken) SetSingleFile(value *string)() {
    m.single_file = value
}
// SetSingleFilePaths sets the single_file_paths property value. The single_file_paths property
func (m *InstallationToken) SetSingleFilePaths(value []string)() {
    m.single_file_paths = value
}
// SetToken sets the token property value. The token property
func (m *InstallationToken) SetToken(value *string)() {
    m.token = value
}
type InstallationTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAt()(*string)
    GetHasMultipleSingleFiles()(*bool)
    GetPermissions()(AppPermissionsable)
    GetRepositories()([]Repositoryable)
    GetRepositorySelection()(*InstallationToken_repository_selection)
    GetSingleFile()(*string)
    GetSingleFilePaths()([]string)
    GetToken()(*string)
    SetExpiresAt(value *string)()
    SetHasMultipleSingleFiles(value *bool)()
    SetPermissions(value AppPermissionsable)()
    SetRepositories(value []Repositoryable)()
    SetRepositorySelection(value *InstallationToken_repository_selection)()
    SetSingleFile(value *string)()
    SetSingleFilePaths(value []string)()
    SetToken(value *string)()
}
