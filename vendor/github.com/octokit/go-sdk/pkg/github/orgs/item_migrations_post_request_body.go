package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemMigrationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether attachments should be excluded from the migration (to reduce migration archive file size).
    exclude_attachments *bool
    // Indicates whether the repository git data should be excluded from the migration.
    exclude_git_data *bool
    // Indicates whether metadata should be excluded and only git source should be included for the migration.
    exclude_metadata *bool
    // Indicates whether projects owned by the organization or users should be excluded. from the migration.
    exclude_owner_projects *bool
    // Indicates whether releases should be excluded from the migration (to reduce migration archive file size).
    exclude_releases *bool
    // Indicates whether repositories should be locked (to prevent manipulation) while migrating data.
    lock_repositories *bool
    // Indicates whether this should only include organization metadata (repositories array should be empty and will ignore other flags).
    org_metadata_only *bool
    // A list of arrays indicating which repositories should be migrated.
    repositories []string
}
// NewItemMigrationsPostRequestBody instantiates a new ItemMigrationsPostRequestBody and sets the default values.
func NewItemMigrationsPostRequestBody()(*ItemMigrationsPostRequestBody) {
    m := &ItemMigrationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMigrationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMigrationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMigrationsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemMigrationsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExcludeAttachments gets the exclude_attachments property value. Indicates whether attachments should be excluded from the migration (to reduce migration archive file size).
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetExcludeAttachments()(*bool) {
    return m.exclude_attachments
}
// GetExcludeGitData gets the exclude_git_data property value. Indicates whether the repository git data should be excluded from the migration.
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetExcludeGitData()(*bool) {
    return m.exclude_git_data
}
// GetExcludeMetadata gets the exclude_metadata property value. Indicates whether metadata should be excluded and only git source should be included for the migration.
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetExcludeMetadata()(*bool) {
    return m.exclude_metadata
}
// GetExcludeOwnerProjects gets the exclude_owner_projects property value. Indicates whether projects owned by the organization or users should be excluded. from the migration.
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetExcludeOwnerProjects()(*bool) {
    return m.exclude_owner_projects
}
// GetExcludeReleases gets the exclude_releases property value. Indicates whether releases should be excluded from the migration (to reduce migration archive file size).
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetExcludeReleases()(*bool) {
    return m.exclude_releases
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemMigrationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exclude_attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeAttachments(val)
        }
        return nil
    }
    res["exclude_git_data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeGitData(val)
        }
        return nil
    }
    res["exclude_metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeMetadata(val)
        }
        return nil
    }
    res["exclude_owner_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeOwnerProjects(val)
        }
        return nil
    }
    res["exclude_releases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeReleases(val)
        }
        return nil
    }
    res["lock_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockRepositories(val)
        }
        return nil
    }
    res["org_metadata_only"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrgMetadataOnly(val)
        }
        return nil
    }
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRepositories(res)
        }
        return nil
    }
    return res
}
// GetLockRepositories gets the lock_repositories property value. Indicates whether repositories should be locked (to prevent manipulation) while migrating data.
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetLockRepositories()(*bool) {
    return m.lock_repositories
}
// GetOrgMetadataOnly gets the org_metadata_only property value. Indicates whether this should only include organization metadata (repositories array should be empty and will ignore other flags).
// returns a *bool when successful
func (m *ItemMigrationsPostRequestBody) GetOrgMetadataOnly()(*bool) {
    return m.org_metadata_only
}
// GetRepositories gets the repositories property value. A list of arrays indicating which repositories should be migrated.
// returns a []string when successful
func (m *ItemMigrationsPostRequestBody) GetRepositories()([]string) {
    return m.repositories
}
// Serialize serializes information the current object
func (m *ItemMigrationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("exclude_attachments", m.GetExcludeAttachments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exclude_git_data", m.GetExcludeGitData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exclude_metadata", m.GetExcludeMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exclude_owner_projects", m.GetExcludeOwnerProjects())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exclude_releases", m.GetExcludeReleases())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("lock_repositories", m.GetLockRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("org_metadata_only", m.GetOrgMetadataOnly())
        if err != nil {
            return err
        }
    }
    if m.GetRepositories() != nil {
        err := writer.WriteCollectionOfStringValues("repositories", m.GetRepositories())
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
func (m *ItemMigrationsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExcludeAttachments sets the exclude_attachments property value. Indicates whether attachments should be excluded from the migration (to reduce migration archive file size).
func (m *ItemMigrationsPostRequestBody) SetExcludeAttachments(value *bool)() {
    m.exclude_attachments = value
}
// SetExcludeGitData sets the exclude_git_data property value. Indicates whether the repository git data should be excluded from the migration.
func (m *ItemMigrationsPostRequestBody) SetExcludeGitData(value *bool)() {
    m.exclude_git_data = value
}
// SetExcludeMetadata sets the exclude_metadata property value. Indicates whether metadata should be excluded and only git source should be included for the migration.
func (m *ItemMigrationsPostRequestBody) SetExcludeMetadata(value *bool)() {
    m.exclude_metadata = value
}
// SetExcludeOwnerProjects sets the exclude_owner_projects property value. Indicates whether projects owned by the organization or users should be excluded. from the migration.
func (m *ItemMigrationsPostRequestBody) SetExcludeOwnerProjects(value *bool)() {
    m.exclude_owner_projects = value
}
// SetExcludeReleases sets the exclude_releases property value. Indicates whether releases should be excluded from the migration (to reduce migration archive file size).
func (m *ItemMigrationsPostRequestBody) SetExcludeReleases(value *bool)() {
    m.exclude_releases = value
}
// SetLockRepositories sets the lock_repositories property value. Indicates whether repositories should be locked (to prevent manipulation) while migrating data.
func (m *ItemMigrationsPostRequestBody) SetLockRepositories(value *bool)() {
    m.lock_repositories = value
}
// SetOrgMetadataOnly sets the org_metadata_only property value. Indicates whether this should only include organization metadata (repositories array should be empty and will ignore other flags).
func (m *ItemMigrationsPostRequestBody) SetOrgMetadataOnly(value *bool)() {
    m.org_metadata_only = value
}
// SetRepositories sets the repositories property value. A list of arrays indicating which repositories should be migrated.
func (m *ItemMigrationsPostRequestBody) SetRepositories(value []string)() {
    m.repositories = value
}
type ItemMigrationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeAttachments()(*bool)
    GetExcludeGitData()(*bool)
    GetExcludeMetadata()(*bool)
    GetExcludeOwnerProjects()(*bool)
    GetExcludeReleases()(*bool)
    GetLockRepositories()(*bool)
    GetOrgMetadataOnly()(*bool)
    GetRepositories()([]string)
    SetExcludeAttachments(value *bool)()
    SetExcludeGitData(value *bool)()
    SetExcludeMetadata(value *bool)()
    SetExcludeOwnerProjects(value *bool)()
    SetExcludeReleases(value *bool)()
    SetLockRepositories(value *bool)()
    SetOrgMetadataOnly(value *bool)()
    SetRepositories(value []string)()
}
