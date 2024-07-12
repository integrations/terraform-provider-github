package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineCommitCommentedEvent timeline Commit Commented Event
type TimelineCommitCommentedEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments []CommitCommentable
    // The commit_id property
    commit_id *string
    // The event property
    event *string
    // The node_id property
    node_id *string
}
// NewTimelineCommitCommentedEvent instantiates a new TimelineCommitCommentedEvent and sets the default values.
func NewTimelineCommitCommentedEvent()(*TimelineCommitCommentedEvent) {
    m := &TimelineCommitCommentedEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineCommitCommentedEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineCommitCommentedEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineCommitCommentedEvent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineCommitCommentedEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
// returns a []CommitCommentable when successful
func (m *TimelineCommitCommentedEvent) GetComments()([]CommitCommentable) {
    return m.comments
}
// GetCommitId gets the commit_id property value. The commit_id property
// returns a *string when successful
func (m *TimelineCommitCommentedEvent) GetCommitId()(*string) {
    return m.commit_id
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *TimelineCommitCommentedEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineCommitCommentedEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommitCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommitCommentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CommitCommentable)
                }
            }
            m.SetComments(res)
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
    return res
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TimelineCommitCommentedEvent) GetNodeId()(*string) {
    return m.node_id
}
// Serialize serializes information the current object
func (m *TimelineCommitCommentedEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("comments", cast)
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
        err := writer.WriteStringValue("event", m.GetEvent())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimelineCommitCommentedEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *TimelineCommitCommentedEvent) SetComments(value []CommitCommentable)() {
    m.comments = value
}
// SetCommitId sets the commit_id property value. The commit_id property
func (m *TimelineCommitCommentedEvent) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetEvent sets the event property value. The event property
func (m *TimelineCommitCommentedEvent) SetEvent(value *string)() {
    m.event = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TimelineCommitCommentedEvent) SetNodeId(value *string)() {
    m.node_id = value
}
type TimelineCommitCommentedEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()([]CommitCommentable)
    GetCommitId()(*string)
    GetEvent()(*string)
    GetNodeId()(*string)
    SetComments(value []CommitCommentable)()
    SetCommitId(value *string)()
    SetEvent(value *string)()
    SetNodeId(value *string)()
}
