package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimpleCommit a commit.
type SimpleCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Information about the Git author
    author SimpleCommit_authorable
    // Information about the Git committer
    committer SimpleCommit_committerable
    // SHA for the commit
    id *string
    // Message describing the purpose of the commit
    message *string
    // Timestamp of the commit
    timestamp *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // SHA for the commit's tree
    tree_id *string
}
// NewSimpleCommit instantiates a new SimpleCommit and sets the default values.
func NewSimpleCommit()(*SimpleCommit) {
    m := &SimpleCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimpleCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimpleCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimpleCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SimpleCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. Information about the Git author
// returns a SimpleCommit_authorable when successful
func (m *SimpleCommit) GetAuthor()(SimpleCommit_authorable) {
    return m.author
}
// GetCommitter gets the committer property value. Information about the Git committer
// returns a SimpleCommit_committerable when successful
func (m *SimpleCommit) GetCommitter()(SimpleCommit_committerable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimpleCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleCommit_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(SimpleCommit_authorable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleCommit_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(SimpleCommit_committerable))
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
func (m *SimpleCommit) GetId()(*string) {
    return m.id
}
// GetMessage gets the message property value. Message describing the purpose of the commit
// returns a *string when successful
func (m *SimpleCommit) GetMessage()(*string) {
    return m.message
}
// GetTimestamp gets the timestamp property value. Timestamp of the commit
// returns a *Time when successful
func (m *SimpleCommit) GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timestamp
}
// GetTreeId gets the tree_id property value. SHA for the commit's tree
// returns a *string when successful
func (m *SimpleCommit) GetTreeId()(*string) {
    return m.tree_id
}
// Serialize serializes information the current object
func (m *SimpleCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SimpleCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. Information about the Git author
func (m *SimpleCommit) SetAuthor(value SimpleCommit_authorable)() {
    m.author = value
}
// SetCommitter sets the committer property value. Information about the Git committer
func (m *SimpleCommit) SetCommitter(value SimpleCommit_committerable)() {
    m.committer = value
}
// SetId sets the id property value. SHA for the commit
func (m *SimpleCommit) SetId(value *string)() {
    m.id = value
}
// SetMessage sets the message property value. Message describing the purpose of the commit
func (m *SimpleCommit) SetMessage(value *string)() {
    m.message = value
}
// SetTimestamp sets the timestamp property value. Timestamp of the commit
func (m *SimpleCommit) SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timestamp = value
}
// SetTreeId sets the tree_id property value. SHA for the commit's tree
func (m *SimpleCommit) SetTreeId(value *string)() {
    m.tree_id = value
}
type SimpleCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(SimpleCommit_authorable)
    GetCommitter()(SimpleCommit_committerable)
    GetId()(*string)
    GetMessage()(*string)
    GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTreeId()(*string)
    SetAuthor(value SimpleCommit_authorable)()
    SetCommitter(value SimpleCommit_committerable)()
    SetId(value *string)()
    SetMessage(value *string)()
    SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTreeId(value *string)()
}
