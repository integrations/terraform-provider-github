package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TopicSearchResultItem_related struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The topic_relation property
    topic_relation TopicSearchResultItem_related_topic_relationable
}
// NewTopicSearchResultItem_related instantiates a new TopicSearchResultItem_related and sets the default values.
func NewTopicSearchResultItem_related()(*TopicSearchResultItem_related) {
    m := &TopicSearchResultItem_related{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTopicSearchResultItem_relatedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTopicSearchResultItem_relatedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicSearchResultItem_related(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TopicSearchResultItem_related) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TopicSearchResultItem_related) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["topic_relation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTopicSearchResultItem_related_topic_relationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicRelation(val.(TopicSearchResultItem_related_topic_relationable))
        }
        return nil
    }
    return res
}
// GetTopicRelation gets the topic_relation property value. The topic_relation property
// returns a TopicSearchResultItem_related_topic_relationable when successful
func (m *TopicSearchResultItem_related) GetTopicRelation()(TopicSearchResultItem_related_topic_relationable) {
    return m.topic_relation
}
// Serialize serializes information the current object
func (m *TopicSearchResultItem_related) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("topic_relation", m.GetTopicRelation())
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
func (m *TopicSearchResultItem_related) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTopicRelation sets the topic_relation property value. The topic_relation property
func (m *TopicSearchResultItem_related) SetTopicRelation(value TopicSearchResultItem_related_topic_relationable)() {
    m.topic_relation = value
}
type TopicSearchResultItem_relatedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTopicRelation()(TopicSearchResultItem_related_topic_relationable)
    SetTopicRelation(value TopicSearchResultItem_related_topic_relationable)()
}
