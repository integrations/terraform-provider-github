package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TopicSearchResultItem_aliases struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The topic_relation property
    topic_relation TopicSearchResultItem_aliases_topic_relationable
}
// NewTopicSearchResultItem_aliases instantiates a new TopicSearchResultItem_aliases and sets the default values.
func NewTopicSearchResultItem_aliases()(*TopicSearchResultItem_aliases) {
    m := &TopicSearchResultItem_aliases{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTopicSearchResultItem_aliasesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTopicSearchResultItem_aliasesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicSearchResultItem_aliases(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TopicSearchResultItem_aliases) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TopicSearchResultItem_aliases) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["topic_relation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTopicSearchResultItem_aliases_topic_relationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicRelation(val.(TopicSearchResultItem_aliases_topic_relationable))
        }
        return nil
    }
    return res
}
// GetTopicRelation gets the topic_relation property value. The topic_relation property
// returns a TopicSearchResultItem_aliases_topic_relationable when successful
func (m *TopicSearchResultItem_aliases) GetTopicRelation()(TopicSearchResultItem_aliases_topic_relationable) {
    return m.topic_relation
}
// Serialize serializes information the current object
func (m *TopicSearchResultItem_aliases) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TopicSearchResultItem_aliases) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTopicRelation sets the topic_relation property value. The topic_relation property
func (m *TopicSearchResultItem_aliases) SetTopicRelation(value TopicSearchResultItem_aliases_topic_relationable)() {
    m.topic_relation = value
}
type TopicSearchResultItem_aliasesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTopicRelation()(TopicSearchResultItem_aliases_topic_relationable)
    SetTopicRelation(value TopicSearchResultItem_aliases_topic_relationable)()
}
