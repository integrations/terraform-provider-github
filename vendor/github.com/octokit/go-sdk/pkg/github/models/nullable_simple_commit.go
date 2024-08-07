package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableSimpleCommit a commit.
type NullableSimpleCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Information about the Git author
    author NullableSimpleCommit_authorable
    // Information about the Git committer
    committer NullableSimpleCommit_committerable
    // SHA for the commit
    id *string
    // Message describing the purpose of the commit
    message *string
    // Timestamp of the commit
    timestamp *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // SHA for the commit's tree
    tree_id *string
}
// NewNullableSimpleCommit instantiates a new NullableSimpleCommit and sets the default values.
func NewNullableSimpleCommit()(*NullableSimpleCommit) {
    m := &NullableSimpleCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableSimpleCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNullableSimpleCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableSimpleCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NullableSimpleCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. Information about the Git author
// returns a NullableSimpleCommit_authorable when successful
func (m *NullableSimpleCommit) GetAuthor()(NullableSimpleCommit_authorable) {
    return m.author
}
// GetCommitter gets the committer property value. Information about the Git committer
// returns a NullableSimpleCommit_committerable when successful
func (m *NullableSimpleCommit) GetCommitter()(NullableSimpleCommit_committerable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NullableSimpleCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleCommit_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(NullableSimpleCommit_authorable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleCommit_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(NullableSimpleCommit_committerable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    res["tree_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTreeId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. SHA for the commit
// returns a *string when successful
func (m *NullableSimpleCommit) GetId()(*string) {
    return m.id
}
// GetMessage gets the message property value. Message describing the purpose of the commit
// returns a *string when successful
func (m *NullableSimpleCommit) GetMessage()(*string) {
    return m.message
}
// GetTimestamp gets the timestamp property value. Timestamp of the commit
// returns a *Time when successful
func (m *NullableSimpleCommit) GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timestamp
}
// GetTreeId gets the tree_id property value. SHA for the commit's tree
// returns a *string when successful
func (m *NullableSimpleCommit) GetTreeId()(*string) {
    return m.tree_id
}
// Serialize serializes information the current object
func (m *NullableSimpleCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("committer", m.GetCommitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tree_id", m.GetTreeId())
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
func (m *NullableSimpleCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. Information about the Git author
func (m *NullableSimpleCommit) SetAuthor(value NullableSimpleCommit_authorable)() {
    m.author = value
}
// SetCommitter sets the committer property value. Information about the Git committer
func (m *NullableSimpleCommit) SetCommitter(value NullableSimpleCommit_committerable)() {
    m.committer = value
}
// SetId sets the id property value. SHA for the commit
func (m *NullableSimpleCommit) SetId(value *string)() {
    m.id = value
}
// SetMessage sets the message property value. Message describing the purpose of the commit
func (m *NullableSimpleCommit) SetMessage(value *string)() {
    m.message = value
}
// SetTimestamp sets the timestamp property value. Timestamp of the commit
func (m *NullableSimpleCommit) SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timestamp = value
}
// SetTreeId sets the tree_id property value. SHA for the commit's tree
func (m *NullableSimpleCommit) SetTreeId(value *string)() {
    m.tree_id = value
}
type NullableSimpleCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleCommit_authorable)
    GetCommitter()(NullableSimpleCommit_committerable)
    GetId()(*string)
    GetMessage()(*string)
    GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTreeId()(*string)
    SetAuthor(value NullableSimpleCommit_authorable)()
    SetCommitter(value NullableSimpleCommit_committerable)()
    SetId(value *string)()
    SetMessage(value *string)()
    SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTreeId(value *string)()
}
