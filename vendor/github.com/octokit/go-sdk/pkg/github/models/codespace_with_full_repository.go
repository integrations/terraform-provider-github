package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespaceWithFullRepository a codespace.
type CodespaceWithFullRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    billable_owner SimpleUserable
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Path to devcontainer.json from repo root used to create Codespace.
    devcontainer_path *string
    // Display name for this codespace.
    display_name *string
    // UUID identifying this codespace's environment.
    environment_id *string
    // Details about the codespace's git repository.
    git_status CodespaceWithFullRepository_git_statusable
    // The id property
    id *int64
    // The number of minutes of inactivity after which this codespace will be automatically stopped.
    idle_timeout_minutes *int32
    // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
    idle_timeout_notice *string
    // Last known time this codespace was started.
    last_used_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The initally assigned location of a new codespace.
    location *CodespaceWithFullRepository_location
    // A description of the machine powering a codespace.
    machine NullableCodespaceMachineable
    // API URL to access available alternate machine types for this codespace.
    machines_url *string
    // Automatically generated name of this codespace.
    name *string
    // A GitHub user.
    owner SimpleUserable
    // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
    pending_operation *bool
    // Text to show user when codespace is disabled by a pending operation
    pending_operation_disabled_reason *string
    // Whether the codespace was created from a prebuild.
    prebuild *bool
    // API URL to publish this codespace to a new repository.
    publish_url *string
    // API URL for the Pull Request associated with this codespace, if any.
    pulls_url *string
    // The recent_folders property
    recent_folders []string
    // Full Repository
    repository FullRepositoryable
    // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
    retention_expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
    retention_period_minutes *int32
    // The runtime_constraints property
    runtime_constraints CodespaceWithFullRepository_runtime_constraintsable
    // API URL to start this codespace.
    start_url *string
    // State of this codespace.
    state *CodespaceWithFullRepository_state
    // API URL to stop this codespace.
    stop_url *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // API URL for this codespace.
    url *string
    // URL to access this codespace on the web.
    web_url *string
}
// NewCodespaceWithFullRepository instantiates a new CodespaceWithFullRepository and sets the default values.
func NewCodespaceWithFullRepository()(*CodespaceWithFullRepository) {
    m := &CodespaceWithFullRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespaceWithFullRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespaceWithFullRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespaceWithFullRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespaceWithFullRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillableOwner gets the billable_owner property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *CodespaceWithFullRepository) GetBillableOwner()(SimpleUserable) {
    return m.billable_owner
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *CodespaceWithFullRepository) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDevcontainerPath gets the devcontainer_path property value. Path to devcontainer.json from repo root used to create Codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetDevcontainerPath()(*string) {
    return m.devcontainer_path
}
// GetDisplayName gets the display_name property value. Display name for this codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetDisplayName()(*string) {
    return m.display_name
}
// GetEnvironmentId gets the environment_id property value. UUID identifying this codespace's environment.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetEnvironmentId()(*string) {
    return m.environment_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespaceWithFullRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billable_owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillableOwner(val.(SimpleUserable))
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
    res["devcontainer_path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevcontainerPath(val)
        }
        return nil
    }
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["environment_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironmentId(val)
        }
        return nil
    }
    res["git_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodespaceWithFullRepository_git_statusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitStatus(val.(CodespaceWithFullRepository_git_statusable))
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
    res["idle_timeout_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdleTimeoutMinutes(val)
        }
        return nil
    }
    res["idle_timeout_notice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdleTimeoutNotice(val)
        }
        return nil
    }
    res["last_used_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsedAt(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodespaceWithFullRepository_location)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(*CodespaceWithFullRepository_location))
        }
        return nil
    }
    res["machine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCodespaceMachineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachine(val.(NullableCodespaceMachineable))
        }
        return nil
    }
    res["machines_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachinesUrl(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(SimpleUserable))
        }
        return nil
    }
    res["pending_operation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOperation(val)
        }
        return nil
    }
    res["pending_operation_disabled_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOperationDisabledReason(val)
        }
        return nil
    }
    res["prebuild"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebuild(val)
        }
        return nil
    }
    res["publish_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishUrl(val)
        }
        return nil
    }
    res["pulls_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullsUrl(val)
        }
        return nil
    }
    res["recent_folders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRecentFolders(res)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFullRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(FullRepositoryable))
        }
        return nil
    }
    res["retention_expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionExpiresAt(val)
        }
        return nil
    }
    res["retention_period_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionPeriodMinutes(val)
        }
        return nil
    }
    res["runtime_constraints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodespaceWithFullRepository_runtime_constraintsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuntimeConstraints(val.(CodespaceWithFullRepository_runtime_constraintsable))
        }
        return nil
    }
    res["start_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartUrl(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodespaceWithFullRepository_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*CodespaceWithFullRepository_state))
        }
        return nil
    }
    res["stop_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStopUrl(val)
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
    res["web_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetGitStatus gets the git_status property value. Details about the codespace's git repository.
// returns a CodespaceWithFullRepository_git_statusable when successful
func (m *CodespaceWithFullRepository) GetGitStatus()(CodespaceWithFullRepository_git_statusable) {
    return m.git_status
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CodespaceWithFullRepository) GetId()(*int64) {
    return m.id
}
// GetIdleTimeoutMinutes gets the idle_timeout_minutes property value. The number of minutes of inactivity after which this codespace will be automatically stopped.
// returns a *int32 when successful
func (m *CodespaceWithFullRepository) GetIdleTimeoutMinutes()(*int32) {
    return m.idle_timeout_minutes
}
// GetIdleTimeoutNotice gets the idle_timeout_notice property value. Text to show user when codespace idle timeout minutes has been overriden by an organization policy
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetIdleTimeoutNotice()(*string) {
    return m.idle_timeout_notice
}
// GetLastUsedAt gets the last_used_at property value. Last known time this codespace was started.
// returns a *Time when successful
func (m *CodespaceWithFullRepository) GetLastUsedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_used_at
}
// GetLocation gets the location property value. The initally assigned location of a new codespace.
// returns a *CodespaceWithFullRepository_location when successful
func (m *CodespaceWithFullRepository) GetLocation()(*CodespaceWithFullRepository_location) {
    return m.location
}
// GetMachine gets the machine property value. A description of the machine powering a codespace.
// returns a NullableCodespaceMachineable when successful
func (m *CodespaceWithFullRepository) GetMachine()(NullableCodespaceMachineable) {
    return m.machine
}
// GetMachinesUrl gets the machines_url property value. API URL to access available alternate machine types for this codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetMachinesUrl()(*string) {
    return m.machines_url
}
// GetName gets the name property value. Automatically generated name of this codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *CodespaceWithFullRepository) GetOwner()(SimpleUserable) {
    return m.owner
}
// GetPendingOperation gets the pending_operation property value. Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
// returns a *bool when successful
func (m *CodespaceWithFullRepository) GetPendingOperation()(*bool) {
    return m.pending_operation
}
// GetPendingOperationDisabledReason gets the pending_operation_disabled_reason property value. Text to show user when codespace is disabled by a pending operation
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetPendingOperationDisabledReason()(*string) {
    return m.pending_operation_disabled_reason
}
// GetPrebuild gets the prebuild property value. Whether the codespace was created from a prebuild.
// returns a *bool when successful
func (m *CodespaceWithFullRepository) GetPrebuild()(*bool) {
    return m.prebuild
}
// GetPublishUrl gets the publish_url property value. API URL to publish this codespace to a new repository.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetPublishUrl()(*string) {
    return m.publish_url
}
// GetPullsUrl gets the pulls_url property value. API URL for the Pull Request associated with this codespace, if any.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetPullsUrl()(*string) {
    return m.pulls_url
}
// GetRecentFolders gets the recent_folders property value. The recent_folders property
// returns a []string when successful
func (m *CodespaceWithFullRepository) GetRecentFolders()([]string) {
    return m.recent_folders
}
// GetRepository gets the repository property value. Full Repository
// returns a FullRepositoryable when successful
func (m *CodespaceWithFullRepository) GetRepository()(FullRepositoryable) {
    return m.repository
}
// GetRetentionExpiresAt gets the retention_expires_at property value. When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
// returns a *Time when successful
func (m *CodespaceWithFullRepository) GetRetentionExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.retention_expires_at
}
// GetRetentionPeriodMinutes gets the retention_period_minutes property value. Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
// returns a *int32 when successful
func (m *CodespaceWithFullRepository) GetRetentionPeriodMinutes()(*int32) {
    return m.retention_period_minutes
}
// GetRuntimeConstraints gets the runtime_constraints property value. The runtime_constraints property
// returns a CodespaceWithFullRepository_runtime_constraintsable when successful
func (m *CodespaceWithFullRepository) GetRuntimeConstraints()(CodespaceWithFullRepository_runtime_constraintsable) {
    return m.runtime_constraints
}
// GetStartUrl gets the start_url property value. API URL to start this codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetStartUrl()(*string) {
    return m.start_url
}
// GetState gets the state property value. State of this codespace.
// returns a *CodespaceWithFullRepository_state when successful
func (m *CodespaceWithFullRepository) GetState()(*CodespaceWithFullRepository_state) {
    return m.state
}
// GetStopUrl gets the stop_url property value. API URL to stop this codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetStopUrl()(*string) {
    return m.stop_url
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *CodespaceWithFullRepository) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. API URL for this codespace.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetUrl()(*string) {
    return m.url
}
// GetWebUrl gets the web_url property value. URL to access this codespace on the web.
// returns a *string when successful
func (m *CodespaceWithFullRepository) GetWebUrl()(*string) {
    return m.web_url
}
// Serialize serializes information the current object
func (m *CodespaceWithFullRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("billable_owner", m.GetBillableOwner())
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
    {
        err := writer.WriteStringValue("devcontainer_path", m.GetDevcontainerPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment_id", m.GetEnvironmentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("git_status", m.GetGitStatus())
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
        err := writer.WriteInt32Value("idle_timeout_minutes", m.GetIdleTimeoutMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("idle_timeout_notice", m.GetIdleTimeoutNotice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_used_at", m.GetLastUsedAt())
        if err != nil {
            return err
        }
    }
    if m.GetLocation() != nil {
        cast := (*m.GetLocation()).String()
        err := writer.WriteStringValue("location", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("machine", m.GetMachine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("machines_url", m.GetMachinesUrl())
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
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("pending_operation", m.GetPendingOperation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pending_operation_disabled_reason", m.GetPendingOperationDisabledReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("prebuild", m.GetPrebuild())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publish_url", m.GetPublishUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pulls_url", m.GetPullsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRecentFolders() != nil {
        err := writer.WriteCollectionOfStringValues("recent_folders", m.GetRecentFolders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("retention_expires_at", m.GetRetentionExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("retention_period_minutes", m.GetRetentionPeriodMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("runtime_constraints", m.GetRuntimeConstraints())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("start_url", m.GetStartUrl())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stop_url", m.GetStopUrl())
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
        err := writer.WriteStringValue("web_url", m.GetWebUrl())
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
func (m *CodespaceWithFullRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillableOwner sets the billable_owner property value. A GitHub user.
func (m *CodespaceWithFullRepository) SetBillableOwner(value SimpleUserable)() {
    m.billable_owner = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *CodespaceWithFullRepository) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDevcontainerPath sets the devcontainer_path property value. Path to devcontainer.json from repo root used to create Codespace.
func (m *CodespaceWithFullRepository) SetDevcontainerPath(value *string)() {
    m.devcontainer_path = value
}
// SetDisplayName sets the display_name property value. Display name for this codespace.
func (m *CodespaceWithFullRepository) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetEnvironmentId sets the environment_id property value. UUID identifying this codespace's environment.
func (m *CodespaceWithFullRepository) SetEnvironmentId(value *string)() {
    m.environment_id = value
}
// SetGitStatus sets the git_status property value. Details about the codespace's git repository.
func (m *CodespaceWithFullRepository) SetGitStatus(value CodespaceWithFullRepository_git_statusable)() {
    m.git_status = value
}
// SetId sets the id property value. The id property
func (m *CodespaceWithFullRepository) SetId(value *int64)() {
    m.id = value
}
// SetIdleTimeoutMinutes sets the idle_timeout_minutes property value. The number of minutes of inactivity after which this codespace will be automatically stopped.
func (m *CodespaceWithFullRepository) SetIdleTimeoutMinutes(value *int32)() {
    m.idle_timeout_minutes = value
}
// SetIdleTimeoutNotice sets the idle_timeout_notice property value. Text to show user when codespace idle timeout minutes has been overriden by an organization policy
func (m *CodespaceWithFullRepository) SetIdleTimeoutNotice(value *string)() {
    m.idle_timeout_notice = value
}
// SetLastUsedAt sets the last_used_at property value. Last known time this codespace was started.
func (m *CodespaceWithFullRepository) SetLastUsedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_used_at = value
}
// SetLocation sets the location property value. The initally assigned location of a new codespace.
func (m *CodespaceWithFullRepository) SetLocation(value *CodespaceWithFullRepository_location)() {
    m.location = value
}
// SetMachine sets the machine property value. A description of the machine powering a codespace.
func (m *CodespaceWithFullRepository) SetMachine(value NullableCodespaceMachineable)() {
    m.machine = value
}
// SetMachinesUrl sets the machines_url property value. API URL to access available alternate machine types for this codespace.
func (m *CodespaceWithFullRepository) SetMachinesUrl(value *string)() {
    m.machines_url = value
}
// SetName sets the name property value. Automatically generated name of this codespace.
func (m *CodespaceWithFullRepository) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *CodespaceWithFullRepository) SetOwner(value SimpleUserable)() {
    m.owner = value
}
// SetPendingOperation sets the pending_operation property value. Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
func (m *CodespaceWithFullRepository) SetPendingOperation(value *bool)() {
    m.pending_operation = value
}
// SetPendingOperationDisabledReason sets the pending_operation_disabled_reason property value. Text to show user when codespace is disabled by a pending operation
func (m *CodespaceWithFullRepository) SetPendingOperationDisabledReason(value *string)() {
    m.pending_operation_disabled_reason = value
}
// SetPrebuild sets the prebuild property value. Whether the codespace was created from a prebuild.
func (m *CodespaceWithFullRepository) SetPrebuild(value *bool)() {
    m.prebuild = value
}
// SetPublishUrl sets the publish_url property value. API URL to publish this codespace to a new repository.
func (m *CodespaceWithFullRepository) SetPublishUrl(value *string)() {
    m.publish_url = value
}
// SetPullsUrl sets the pulls_url property value. API URL for the Pull Request associated with this codespace, if any.
func (m *CodespaceWithFullRepository) SetPullsUrl(value *string)() {
    m.pulls_url = value
}
// SetRecentFolders sets the recent_folders property value. The recent_folders property
func (m *CodespaceWithFullRepository) SetRecentFolders(value []string)() {
    m.recent_folders = value
}
// SetRepository sets the repository property value. Full Repository
func (m *CodespaceWithFullRepository) SetRepository(value FullRepositoryable)() {
    m.repository = value
}
// SetRetentionExpiresAt sets the retention_expires_at property value. When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
func (m *CodespaceWithFullRepository) SetRetentionExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.retention_expires_at = value
}
// SetRetentionPeriodMinutes sets the retention_period_minutes property value. Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
func (m *CodespaceWithFullRepository) SetRetentionPeriodMinutes(value *int32)() {
    m.retention_period_minutes = value
}
// SetRuntimeConstraints sets the runtime_constraints property value. The runtime_constraints property
func (m *CodespaceWithFullRepository) SetRuntimeConstraints(value CodespaceWithFullRepository_runtime_constraintsable)() {
    m.runtime_constraints = value
}
// SetStartUrl sets the start_url property value. API URL to start this codespace.
func (m *CodespaceWithFullRepository) SetStartUrl(value *string)() {
    m.start_url = value
}
// SetState sets the state property value. State of this codespace.
func (m *CodespaceWithFullRepository) SetState(value *CodespaceWithFullRepository_state)() {
    m.state = value
}
// SetStopUrl sets the stop_url property value. API URL to stop this codespace.
func (m *CodespaceWithFullRepository) SetStopUrl(value *string)() {
    m.stop_url = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *CodespaceWithFullRepository) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. API URL for this codespace.
func (m *CodespaceWithFullRepository) SetUrl(value *string)() {
    m.url = value
}
// SetWebUrl sets the web_url property value. URL to access this codespace on the web.
func (m *CodespaceWithFullRepository) SetWebUrl(value *string)() {
    m.web_url = value
}
type CodespaceWithFullRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillableOwner()(SimpleUserable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDevcontainerPath()(*string)
    GetDisplayName()(*string)
    GetEnvironmentId()(*string)
    GetGitStatus()(CodespaceWithFullRepository_git_statusable)
    GetId()(*int64)
    GetIdleTimeoutMinutes()(*int32)
    GetIdleTimeoutNotice()(*string)
    GetLastUsedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocation()(*CodespaceWithFullRepository_location)
    GetMachine()(NullableCodespaceMachineable)
    GetMachinesUrl()(*string)
    GetName()(*string)
    GetOwner()(SimpleUserable)
    GetPendingOperation()(*bool)
    GetPendingOperationDisabledReason()(*string)
    GetPrebuild()(*bool)
    GetPublishUrl()(*string)
    GetPullsUrl()(*string)
    GetRecentFolders()([]string)
    GetRepository()(FullRepositoryable)
    GetRetentionExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRetentionPeriodMinutes()(*int32)
    GetRuntimeConstraints()(CodespaceWithFullRepository_runtime_constraintsable)
    GetStartUrl()(*string)
    GetState()(*CodespaceWithFullRepository_state)
    GetStopUrl()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetWebUrl()(*string)
    SetBillableOwner(value SimpleUserable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDevcontainerPath(value *string)()
    SetDisplayName(value *string)()
    SetEnvironmentId(value *string)()
    SetGitStatus(value CodespaceWithFullRepository_git_statusable)()
    SetId(value *int64)()
    SetIdleTimeoutMinutes(value *int32)()
    SetIdleTimeoutNotice(value *string)()
    SetLastUsedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocation(value *CodespaceWithFullRepository_location)()
    SetMachine(value NullableCodespaceMachineable)()
    SetMachinesUrl(value *string)()
    SetName(value *string)()
    SetOwner(value SimpleUserable)()
    SetPendingOperation(value *bool)()
    SetPendingOperationDisabledReason(value *string)()
    SetPrebuild(value *bool)()
    SetPublishUrl(value *string)()
    SetPullsUrl(value *string)()
    SetRecentFolders(value []string)()
    SetRepository(value FullRepositoryable)()
    SetRetentionExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRetentionPeriodMinutes(value *int32)()
    SetRuntimeConstraints(value CodespaceWithFullRepository_runtime_constraintsable)()
    SetStartUrl(value *string)()
    SetState(value *CodespaceWithFullRepository_state)()
    SetStopUrl(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetWebUrl(value *string)()
}
