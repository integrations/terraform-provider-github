package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodespacesPostRequestBodyMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Path to devcontainer.json config to use for this codespace
    devcontainer_path *string
    // Time in minutes before codespace stops from inactivity
    idle_timeout_minutes *int32
    // The requested location for a new codespace. Best efforts are made to respect this upon creation. Assigned by IP if not provided.
    location *string
    // Machine type to use for this codespace
    machine *string
    // Pull request number for this codespace
    pull_request CodespacesPostRequestBodyMember2_pull_requestable
    // Working directory for this codespace
    working_directory *string
}
// NewCodespacesPostRequestBodyMember2 instantiates a new CodespacesPostRequestBodyMember2 and sets the default values.
func NewCodespacesPostRequestBodyMember2()(*CodespacesPostRequestBodyMember2) {
    m := &CodespacesPostRequestBodyMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesPostRequestBodyMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespacesPostRequestBodyMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesPostRequestBodyMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespacesPostRequestBodyMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDevcontainerPath gets the devcontainer_path property value. Path to devcontainer.json config to use for this codespace
// returns a *string when successful
func (m *CodespacesPostRequestBodyMember2) GetDevcontainerPath()(*string) {
    return m.devcontainer_path
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespacesPostRequestBodyMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["pull_request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodespacesPostRequestBodyMember2_pull_requestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequest(val.(CodespacesPostRequestBodyMember2_pull_requestable))
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
func (m *CodespacesPostRequestBodyMember2) GetIdleTimeoutMinutes()(*int32) {
    return m.idle_timeout_minutes
}
// GetLocation gets the location property value. The requested location for a new codespace. Best efforts are made to respect this upon creation. Assigned by IP if not provided.
// returns a *string when successful
func (m *CodespacesPostRequestBodyMember2) GetLocation()(*string) {
    return m.location
}
// GetMachine gets the machine property value. Machine type to use for this codespace
// returns a *string when successful
func (m *CodespacesPostRequestBodyMember2) GetMachine()(*string) {
    return m.machine
}
// GetPullRequest gets the pull_request property value. Pull request number for this codespace
// returns a CodespacesPostRequestBodyMember2_pull_requestable when successful
func (m *CodespacesPostRequestBodyMember2) GetPullRequest()(CodespacesPostRequestBodyMember2_pull_requestable) {
    return m.pull_request
}
// GetWorkingDirectory gets the working_directory property value. Working directory for this codespace
// returns a *string when successful
func (m *CodespacesPostRequestBodyMember2) GetWorkingDirectory()(*string) {
    return m.working_directory
}
// Serialize serializes information the current object
func (m *CodespacesPostRequestBodyMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("devcontainer_path", m.GetDevcontainerPath())
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
        err := writer.WriteObjectValue("pull_request", m.GetPullRequest())
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
func (m *CodespacesPostRequestBodyMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDevcontainerPath sets the devcontainer_path property value. Path to devcontainer.json config to use for this codespace
func (m *CodespacesPostRequestBodyMember2) SetDevcontainerPath(value *string)() {
    m.devcontainer_path = value
}
// SetIdleTimeoutMinutes sets the idle_timeout_minutes property value. Time in minutes before codespace stops from inactivity
func (m *CodespacesPostRequestBodyMember2) SetIdleTimeoutMinutes(value *int32)() {
    m.idle_timeout_minutes = value
}
// SetLocation sets the location property value. The requested location for a new codespace. Best efforts are made to respect this upon creation. Assigned by IP if not provided.
func (m *CodespacesPostRequestBodyMember2) SetLocation(value *string)() {
    m.location = value
}
// SetMachine sets the machine property value. Machine type to use for this codespace
func (m *CodespacesPostRequestBodyMember2) SetMachine(value *string)() {
    m.machine = value
}
// SetPullRequest sets the pull_request property value. Pull request number for this codespace
func (m *CodespacesPostRequestBodyMember2) SetPullRequest(value CodespacesPostRequestBodyMember2_pull_requestable)() {
    m.pull_request = value
}
// SetWorkingDirectory sets the working_directory property value. Working directory for this codespace
func (m *CodespacesPostRequestBodyMember2) SetWorkingDirectory(value *string)() {
    m.working_directory = value
}
type CodespacesPostRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevcontainerPath()(*string)
    GetIdleTimeoutMinutes()(*int32)
    GetLocation()(*string)
    GetMachine()(*string)
    GetPullRequest()(CodespacesPostRequestBodyMember2_pull_requestable)
    GetWorkingDirectory()(*string)
    SetDevcontainerPath(value *string)()
    SetIdleTimeoutMinutes(value *int32)()
    SetLocation(value *string)()
    SetMachine(value *string)()
    SetPullRequest(value CodespacesPostRequestBodyMember2_pull_requestable)()
    SetWorkingDirectory(value *string)()
}
