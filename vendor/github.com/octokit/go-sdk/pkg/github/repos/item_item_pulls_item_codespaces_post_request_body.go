package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemCodespacesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // IP for location auto-detection when proxying a request
    client_ip *string
    // Path to devcontainer.json config to use for this codespace
    devcontainer_path *string
    // Display name for this codespace
    display_name *string
    // Time in minutes before codespace stops from inactivity
    idle_timeout_minutes *int32
    // The requested location for a new codespace. Best efforts are made to respect this upon creation. Assigned by IP if not provided.
    location *string
    // Machine type to use for this codespace
    machine *string
    // Whether to authorize requested permissions from devcontainer.json
    multi_repo_permissions_opt_out *bool
    // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
    retention_period_minutes *int32
    // Working directory for this codespace
    working_directory *string
}
// NewItemItemPullsItemCodespacesPostRequestBody instantiates a new ItemItemPullsItemCodespacesPostRequestBody and sets the default values.
func NewItemItemPullsItemCodespacesPostRequestBody()(*ItemItemPullsItemCodespacesPostRequestBody) {
    m := &ItemItemPullsItemCodespacesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemCodespacesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemCodespacesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemCodespacesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientIp gets the client_ip property value. IP for location auto-detection when proxying a request
// returns a *string when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetClientIp()(*string) {
    return m.client_ip
}
// GetDevcontainerPath gets the devcontainer_path property value. Path to devcontainer.json config to use for this codespace
// returns a *string when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetDevcontainerPath()(*string) {
    return m.devcontainer_path
}
// GetDisplayName gets the display_name property value. Display name for this codespace
// returns a *string when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientIp(val)
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
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["machine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachine(val)
        }
        return nil
    }
    res["multi_repo_permissions_opt_out"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultiRepoPermissionsOptOut(val)
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
    res["working_directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkingDirectory(val)
        }
        return nil
    }
    return res
}
// GetIdleTimeoutMinutes gets the idle_timeout_minutes property value. Time in minutes before codespace stops from inactivity
// returns a *int32 when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetIdleTimeoutMinutes()(*int32) {
    return m.idle_timeout_minutes
}
// GetLocation gets the location property value. The requested location for a new codespace. Best efforts are made to respect this upon creation. Assigned by IP if not provided.
// returns a *string when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetLocation()(*string) {
    return m.location
}
// GetMachine gets the machine property value. Machine type to use for this codespace
// returns a *string when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetMachine()(*string) {
    return m.machine
}
// GetMultiRepoPermissionsOptOut gets the multi_repo_permissions_opt_out property value. Whether to authorize requested permissions from devcontainer.json
// returns a *bool when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetMultiRepoPermissionsOptOut()(*bool) {
    return m.multi_repo_permissions_opt_out
}
// GetRetentionPeriodMinutes gets the retention_period_minutes property value. Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
// returns a *int32 when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetRetentionPeriodMinutes()(*int32) {
    return m.retention_period_minutes
}
// GetWorkingDirectory gets the working_directory property value. Working directory for this codespace
// returns a *string when successful
func (m *ItemItemPullsItemCodespacesPostRequestBody) GetWorkingDirectory()(*string) {
    return m.working_directory
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemCodespacesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_ip", m.GetClientIp())
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
        err := writer.WriteInt32Value("idle_timeout_minutes", m.GetIdleTimeoutMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("machine", m.GetMachine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("multi_repo_permissions_opt_out", m.GetMultiRepoPermissionsOptOut())
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
        err := writer.WriteStringValue("working_directory", m.GetWorkingDirectory())
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
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientIp sets the client_ip property value. IP for location auto-detection when proxying a request
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetClientIp(value *string)() {
    m.client_ip = value
}
// SetDevcontainerPath sets the devcontainer_path property value. Path to devcontainer.json config to use for this codespace
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetDevcontainerPath(value *string)() {
    m.devcontainer_path = value
}
// SetDisplayName sets the display_name property value. Display name for this codespace
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetIdleTimeoutMinutes sets the idle_timeout_minutes property value. Time in minutes before codespace stops from inactivity
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetIdleTimeoutMinutes(value *int32)() {
    m.idle_timeout_minutes = value
}
// SetLocation sets the location property value. The requested location for a new codespace. Best efforts are made to respect this upon creation. Assigned by IP if not provided.
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetLocation(value *string)() {
    m.location = value
}
// SetMachine sets the machine property value. Machine type to use for this codespace
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetMachine(value *string)() {
    m.machine = value
}
// SetMultiRepoPermissionsOptOut sets the multi_repo_permissions_opt_out property value. Whether to authorize requested permissions from devcontainer.json
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetMultiRepoPermissionsOptOut(value *bool)() {
    m.multi_repo_permissions_opt_out = value
}
// SetRetentionPeriodMinutes sets the retention_period_minutes property value. Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetRetentionPeriodMinutes(value *int32)() {
    m.retention_period_minutes = value
}
// SetWorkingDirectory sets the working_directory property value. Working directory for this codespace
func (m *ItemItemPullsItemCodespacesPostRequestBody) SetWorkingDirectory(value *string)() {
    m.working_directory = value
}
type ItemItemPullsItemCodespacesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientIp()(*string)
    GetDevcontainerPath()(*string)
    GetDisplayName()(*string)
    GetIdleTimeoutMinutes()(*int32)
    GetLocation()(*string)
    GetMachine()(*string)
    GetMultiRepoPermissionsOptOut()(*bool)
    GetRetentionPeriodMinutes()(*int32)
    GetWorkingDirectory()(*string)
    SetClientIp(value *string)()
    SetDevcontainerPath(value *string)()
    SetDisplayName(value *string)()
    SetIdleTimeoutMinutes(value *int32)()
    SetLocation(value *string)()
    SetMachine(value *string)()
    SetMultiRepoPermissionsOptOut(value *bool)()
    SetRetentionPeriodMinutes(value *int32)()
    SetWorkingDirectory(value *string)()
}
