package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineLineCommentedEvent timeline Line Commented Event
type TimelineLineCommentedEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments []PullRequestReviewCommentable
    // The event property
    event *string
    // The node_id property
    node_id *string
}
// NewTimelineLineCommentedEvent instantiates a new TimelineLineCommentedEvent and sets the default values.
func NewTimelineLineCommentedEvent()(*TimelineLineCommentedEvent) {
    m := &TimelineLineCommentedEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineLineCommentedEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineLineCommentedEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineLineCommentedEvent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineLineCommentedEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
// returns a []PullRequestReviewCommentable when successful
func (m *TimelineLineCommentedEvent) GetComments()([]PullRequestReviewCommentable) {
    return m.comments
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *TimelineLineCommentedEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineLineCommentedEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePullRequestReviewCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullRequestReviewCommentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PullRequestReviewCommentable)
                }
            }
            m.SetComments(res)
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
func (m *TimelineLineCommentedEvent) GetNodeId()(*string) {
    return m.node_id
}
// Serialize serializes information the current object
func (m *TimelineLineCommentedEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TimelineLineCommentedEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *TimelineLineCommentedEvent) SetComments(value []PullRequestReviewCommentable)() {
    m.comments = value
}
// SetEvent sets the event property value. The event property
func (m *TimelineLineCommentedEvent) SetEvent(value *string)() {
    m.event = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TimelineLineCommentedEvent) SetNodeId(value *string)() {
    m.node_id = value
}
type TimelineLineCommentedEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()([]PullRequestReviewCommentable)
    GetEvent()(*string)
    GetNodeId()(*string)
    SetComments(value []PullRequestReviewCommentable)()
    SetEvent(value *string)()
    SetNodeId(value *string)()
}
