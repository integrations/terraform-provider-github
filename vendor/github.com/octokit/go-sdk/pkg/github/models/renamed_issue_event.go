package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RenamedIssueEvent renamed Issue Event
type RenamedIssueEvent struct {
    // A GitHub user.
    actor SimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit_id property
    commit_id *string
    // The commit_url property
    commit_url *string
    // The created_at property
    created_at *string
    // The event property
    event *string
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // The rename property
    rename RenamedIssueEvent_renameable
    // The url property
    url *string
}
// NewRenamedIssueEvent instantiates a new RenamedIssueEvent and sets the default values.
func NewRenamedIssueEvent()(*RenamedIssueEvent) {
    m := &RenamedIssueEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRenamedIssueEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRenamedIssueEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRenamedIssueEvent(), nil
}
// GetActor gets the actor property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *RenamedIssueEvent) GetActor()(SimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RenamedIssueEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitId gets the commit_id property value. The commit_id property
// returns a *string when successful
func (m *RenamedIssueEvent) GetCommitId()(*string) {
    return m.commit_id
}
// GetCommitUrl gets the commit_url property value. The commit_url property
// returns a *string when successful
func (m *RenamedIssueEvent) GetCommitUrl()(*string) {
    return m.commit_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *RenamedIssueEvent) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *RenamedIssueEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RenamedIssueEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(SimpleUserable))
        }
        return nil
    }
    res["commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitId(val)
        }
        return nil
    }
    res["commit_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitUrl(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["performed_via_github_app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedViaGithubApp(val.(NullableIntegrationable))
        }
        return nil
    }
    res["rename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRenamedIssueEvent_renameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRename(val.(RenamedIssueEvent_renameable))
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *RenamedIssueEvent) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *RenamedIssueEvent) GetNodeId()(*string) {
    return m.node_id
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *RenamedIssueEvent) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetRename gets the rename property value. The rename property
// returns a RenamedIssueEvent_renameable when successful
func (m *RenamedIssueEvent) GetRename()(RenamedIssueEvent_renameable) {
    return m.rename
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *RenamedIssueEvent) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *RenamedIssueEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_id", m.GetCommitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_url", m.GetCommitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rename", m.GetRename())
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
// SetActor sets the actor property value. A GitHub user.
func (m *RenamedIssueEvent) SetActor(value SimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RenamedIssueEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitId sets the commit_id property value. The commit_id property
func (m *RenamedIssueEvent) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetCommitUrl sets the commit_url property value. The commit_url property
func (m *RenamedIssueEvent) SetCommitUrl(value *string)() {
    m.commit_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *RenamedIssueEvent) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEvent sets the event property value. The event property
func (m *RenamedIssueEvent) SetEvent(value *string)() {
    m.event = value
}
// SetId sets the id property value. The id property
func (m *RenamedIssueEvent) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *RenamedIssueEvent) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *RenamedIssueEvent) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetRename sets the rename property value. The rename property
func (m *RenamedIssueEvent) SetRename(value RenamedIssueEvent_renameable)() {
    m.rename = value
}
// SetUrl sets the url property value. The url property
func (m *RenamedIssueEvent) SetUrl(value *string)() {
    m.url = value
}
type RenamedIssueEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(SimpleUserable)
    GetCommitId()(*string)
    GetCommitUrl()(*string)
    GetCreatedAt()(*string)
    GetEvent()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetRename()(RenamedIssueEvent_renameable)
    GetUrl()(*string)
    SetActor(value SimpleUserable)()
    SetCommitId(value *string)()
    SetCommitUrl(value *string)()
    SetCreatedAt(value *string)()
    SetEvent(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetRename(value RenamedIssueEvent_renameable)()
    SetUrl(value *string)()
}
