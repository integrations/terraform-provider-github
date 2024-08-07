package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespaceWithFullRepository_git_status details about the codespace's git repository.
type CodespaceWithFullRepository_git_status struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of commits the local repository is ahead of the remote.
    ahead *int32
    // The number of commits the local repository is behind the remote.
    behind *int32
    // Whether the local repository has uncommitted changes.
    has_uncommitted_changes *bool
    // Whether the local repository has unpushed changes.
    has_unpushed_changes *bool
    // The current branch (or SHA if in detached HEAD state) of the local repository.
    ref *string
}
// NewCodespaceWithFullRepository_git_status instantiates a new CodespaceWithFullRepository_git_status and sets the default values.
func NewCodespaceWithFullRepository_git_status()(*CodespaceWithFullRepository_git_status) {
    m := &CodespaceWithFullRepository_git_status{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespaceWithFullRepository_git_statusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespaceWithFullRepository_git_statusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespaceWithFullRepository_git_status(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespaceWithFullRepository_git_status) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAhead gets the ahead property value. The number of commits the local repository is ahead of the remote.
// returns a *int32 when successful
func (m *CodespaceWithFullRepository_git_status) GetAhead()(*int32) {
    return m.ahead
}
// GetBehind gets the behind property value. The number of commits the local repository is behind the remote.
// returns a *int32 when successful
func (m *CodespaceWithFullRepository_git_status) GetBehind()(*int32) {
    return m.behind
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespaceWithFullRepository_git_status) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ahead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAhead(val)
        }
        return nil
    }
    res["behind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBehind(val)
        }
        return nil
    }
    res["has_uncommitted_changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasUncommittedChanges(val)
        }
        return nil
    }
    res["has_unpushed_changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasUnpushedChanges(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    return res
}
// GetHasUncommittedChanges gets the has_uncommitted_changes property value. Whether the local repository has uncommitted changes.
// returns a *bool when successful
func (m *CodespaceWithFullRepository_git_status) GetHasUncommittedChanges()(*bool) {
    return m.has_uncommitted_changes
}
// GetHasUnpushedChanges gets the has_unpushed_changes property value. Whether the local repository has unpushed changes.
// returns a *bool when successful
func (m *CodespaceWithFullRepository_git_status) GetHasUnpushedChanges()(*bool) {
    return m.has_unpushed_changes
}
// GetRef gets the ref property value. The current branch (or SHA if in detached HEAD state) of the local repository.
// returns a *string when successful
func (m *CodespaceWithFullRepository_git_status) GetRef()(*string) {
    return m.ref
}
// Serialize serializes information the current object
func (m *CodespaceWithFullRepository_git_status) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("ahead", m.GetAhead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("behind", m.GetBehind())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_uncommitted_changes", m.GetHasUncommittedChanges())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_unpushed_changes", m.GetHasUnpushedChanges())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
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
func (m *CodespaceWithFullRepository_git_status) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAhead sets the ahead property value. The number of commits the local repository is ahead of the remote.
func (m *CodespaceWithFullRepository_git_status) SetAhead(value *int32)() {
    m.ahead = value
}
// SetBehind sets the behind property value. The number of commits the local repository is behind the remote.
func (m *CodespaceWithFullRepository_git_status) SetBehind(value *int32)() {
    m.behind = value
}
// SetHasUncommittedChanges sets the has_uncommitted_changes property value. Whether the local repository has uncommitted changes.
func (m *CodespaceWithFullRepository_git_status) SetHasUncommittedChanges(value *bool)() {
    m.has_uncommitted_changes = value
}
// SetHasUnpushedChanges sets the has_unpushed_changes property value. Whether the local repository has unpushed changes.
func (m *CodespaceWithFullRepository_git_status) SetHasUnpushedChanges(value *bool)() {
    m.has_unpushed_changes = value
}
// SetRef sets the ref property value. The current branch (or SHA if in detached HEAD state) of the local repository.
func (m *CodespaceWithFullRepository_git_status) SetRef(value *string)() {
    m.ref = value
}
type CodespaceWithFullRepository_git_statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAhead()(*int32)
    GetBehind()(*int32)
    GetHasUncommittedChanges()(*bool)
    GetHasUnpushedChanges()(*bool)
    GetRef()(*string)
    SetAhead(value *int32)()
    SetBehind(value *int32)()
    SetHasUncommittedChanges(value *bool)()
    SetHasUnpushedChanges(value *bool)()
    SetRef(value *string)()
}
