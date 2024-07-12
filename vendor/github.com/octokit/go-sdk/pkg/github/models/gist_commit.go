package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistCommit gist Commit
type GistCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The change_status property
    change_status GistCommit_change_statusable
    // The committed_at property
    committed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
    // The version property
    version *string
}
// NewGistCommit instantiates a new GistCommit and sets the default values.
func NewGistCommit()(*GistCommit) {
    m := &GistCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GistCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChangeStatus gets the change_status property value. The change_status property
// returns a GistCommit_change_statusable when successful
func (m *GistCommit) GetChangeStatus()(GistCommit_change_statusable) {
    return m.change_status
}
// GetCommittedAt gets the committed_at property value. The committed_at property
// returns a *Time when successful
func (m *GistCommit) GetCommittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.committed_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["change_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistCommit_change_statusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeStatus(val.(GistCommit_change_statusable))
        }
        return nil
    }
    res["committed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommittedAt(val)
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GistCommit) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *GistCommit) GetUser()(NullableSimpleUserable) {
    return m.user
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *GistCommit) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *GistCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("change_status", m.GetChangeStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("committed_at", m.GetCommittedAt())
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
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *GistCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChangeStatus sets the change_status property value. The change_status property
func (m *GistCommit) SetChangeStatus(value GistCommit_change_statusable)() {
    m.change_status = value
}
// SetCommittedAt sets the committed_at property value. The committed_at property
func (m *GistCommit) SetCommittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.committed_at = value
}
// SetUrl sets the url property value. The url property
func (m *GistCommit) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *GistCommit) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
// SetVersion sets the version property value. The version property
func (m *GistCommit) SetVersion(value *string)() {
    m.version = value
}
type GistCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeStatus()(GistCommit_change_statusable)
    GetCommittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    GetVersion()(*string)
    SetChangeStatus(value GistCommit_change_statusable)()
    SetCommittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
    SetVersion(value *string)()
}
