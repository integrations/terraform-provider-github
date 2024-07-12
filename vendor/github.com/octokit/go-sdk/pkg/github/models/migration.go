package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Migration a migration.
type Migration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The archive_url property
    archive_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
    exclude []string
    // The exclude_attachments property
    exclude_attachments *bool
    // The exclude_git_data property
    exclude_git_data *bool
    // The exclude_metadata property
    exclude_metadata *bool
    // The exclude_owner_projects property
    exclude_owner_projects *bool
    // The exclude_releases property
    exclude_releases *bool
    // The guid property
    guid *string
    // The id property
    id *int64
    // The lock_repositories property
    lock_repositories *bool
    // The node_id property
    node_id *string
    // The org_metadata_only property
    org_metadata_only *bool
    // A GitHub user.
    owner NullableSimpleUserable
    // The repositories included in the migration. Only returned for export migrations.
    repositories []Repositoryable
    // The state property
    state *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewMigration instantiates a new Migration and sets the default values.
func NewMigration()(*Migration) {
    m := &Migration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMigrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMigrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMigration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Migration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchiveUrl gets the archive_url property value. The archive_url property
// returns a *string when successful
func (m *Migration) GetArchiveUrl()(*string) {
    return m.archive_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Migration) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetExclude gets the exclude property value. Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
// returns a []string when successful
func (m *Migration) GetExclude()([]string) {
    return m.exclude
}
// GetExcludeAttachments gets the exclude_attachments property value. The exclude_attachments property
// returns a *bool when successful
func (m *Migration) GetExcludeAttachments()(*bool) {
    return m.exclude_attachments
}
// GetExcludeGitData gets the exclude_git_data property value. The exclude_git_data property
// returns a *bool when successful
func (m *Migration) GetExcludeGitData()(*bool) {
    return m.exclude_git_data
}
// GetExcludeMetadata gets the exclude_metadata property value. The exclude_metadata property
// returns a *bool when successful
func (m *Migration) GetExcludeMetadata()(*bool) {
    return m.exclude_metadata
}
// GetExcludeOwnerProjects gets the exclude_owner_projects property value. The exclude_owner_projects property
// returns a *bool when successful
func (m *Migration) GetExcludeOwnerProjects()(*bool) {
    return m.exclude_owner_projects
}
// GetExcludeReleases gets the exclude_releases property value. The exclude_releases property
// returns a *bool when successful
func (m *Migration) GetExcludeReleases()(*bool) {
    return m.exclude_releases
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Migration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archive_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveUrl(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["exclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExclude(res)
        }
        return nil
    }
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
    res["guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(NullableSimpleUserable))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetGuid gets the guid property value. The guid property
// returns a *string when successful
func (m *Migration) GetGuid()(*string) {
    return m.guid
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *Migration) GetId()(*int64) {
    return m.id
}
// GetLockRepositories gets the lock_repositories property value. The lock_repositories property
// returns a *bool when successful
func (m *Migration) GetLockRepositories()(*bool) {
    return m.lock_repositories
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Migration) GetNodeId()(*string) {
    return m.node_id
}
// GetOrgMetadataOnly gets the org_metadata_only property value. The org_metadata_only property
// returns a *bool when successful
func (m *Migration) GetOrgMetadataOnly()(*bool) {
    return m.org_metadata_only
}
// GetOwner gets the owner property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Migration) GetOwner()(NullableSimpleUserable) {
    return m.owner
}
// GetRepositories gets the repositories property value. The repositories included in the migration. Only returned for export migrations.
// returns a []Repositoryable when successful
func (m *Migration) GetRepositories()([]Repositoryable) {
    return m.repositories
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *Migration) GetState()(*string) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Migration) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Migration) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Migration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archive_url", m.GetArchiveUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetExclude() != nil {
        err := writer.WriteCollectionOfStringValues("exclude", m.GetExclude())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("guid", m.GetGuid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
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
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
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
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *Migration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchiveUrl sets the archive_url property value. The archive_url property
func (m *Migration) SetArchiveUrl(value *string)() {
    m.archive_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Migration) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetExclude sets the exclude property value. Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
func (m *Migration) SetExclude(value []string)() {
    m.exclude = value
}
// SetExcludeAttachments sets the exclude_attachments property value. The exclude_attachments property
func (m *Migration) SetExcludeAttachments(value *bool)() {
    m.exclude_attachments = value
}
// SetExcludeGitData sets the exclude_git_data property value. The exclude_git_data property
func (m *Migration) SetExcludeGitData(value *bool)() {
    m.exclude_git_data = value
}
// SetExcludeMetadata sets the exclude_metadata property value. The exclude_metadata property
func (m *Migration) SetExcludeMetadata(value *bool)() {
    m.exclude_metadata = value
}
// SetExcludeOwnerProjects sets the exclude_owner_projects property value. The exclude_owner_projects property
func (m *Migration) SetExcludeOwnerProjects(value *bool)() {
    m.exclude_owner_projects = value
}
// SetExcludeReleases sets the exclude_releases property value. The exclude_releases property
func (m *Migration) SetExcludeReleases(value *bool)() {
    m.exclude_releases = value
}
// SetGuid sets the guid property value. The guid property
func (m *Migration) SetGuid(value *string)() {
    m.guid = value
}
// SetId sets the id property value. The id property
func (m *Migration) SetId(value *int64)() {
    m.id = value
}
// SetLockRepositories sets the lock_repositories property value. The lock_repositories property
func (m *Migration) SetLockRepositories(value *bool)() {
    m.lock_repositories = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Migration) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOrgMetadataOnly sets the org_metadata_only property value. The org_metadata_only property
func (m *Migration) SetOrgMetadataOnly(value *bool)() {
    m.org_metadata_only = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *Migration) SetOwner(value NullableSimpleUserable)() {
    m.owner = value
}
// SetRepositories sets the repositories property value. The repositories included in the migration. Only returned for export migrations.
func (m *Migration) SetRepositories(value []Repositoryable)() {
    m.repositories = value
}
// SetState sets the state property value. The state property
func (m *Migration) SetState(value *string)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Migration) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Migration) SetUrl(value *string)() {
    m.url = value
}
type Migrationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExclude()([]string)
    GetExcludeAttachments()(*bool)
    GetExcludeGitData()(*bool)
    GetExcludeMetadata()(*bool)
    GetExcludeOwnerProjects()(*bool)
    GetExcludeReleases()(*bool)
    GetGuid()(*string)
    GetId()(*int64)
    GetLockRepositories()(*bool)
    GetNodeId()(*string)
    GetOrgMetadataOnly()(*bool)
    GetOwner()(NullableSimpleUserable)
    GetRepositories()([]Repositoryable)
    GetState()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetArchiveUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExclude(value []string)()
    SetExcludeAttachments(value *bool)()
    SetExcludeGitData(value *bool)()
    SetExcludeMetadata(value *bool)()
    SetExcludeOwnerProjects(value *bool)()
    SetExcludeReleases(value *bool)()
    SetGuid(value *string)()
    SetId(value *int64)()
    SetLockRepositories(value *bool)()
    SetNodeId(value *string)()
    SetOrgMetadataOnly(value *bool)()
    SetOwner(value NullableSimpleUserable)()
    SetRepositories(value []Repositoryable)()
    SetState(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
