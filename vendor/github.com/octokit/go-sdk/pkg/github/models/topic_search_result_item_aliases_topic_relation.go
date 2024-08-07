package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TopicSearchResultItem_aliases_topic_relation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // The name property
    name *string
    // The relation_type property
    relation_type *string
    // The topic_id property
    topic_id *int32
}
// NewTopicSearchResultItem_aliases_topic_relation instantiates a new TopicSearchResultItem_aliases_topic_relation and sets the default values.
func NewTopicSearchResultItem_aliases_topic_relation()(*TopicSearchResultItem_aliases_topic_relation) {
    m := &TopicSearchResultItem_aliases_topic_relation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTopicSearchResultItem_aliases_topic_relationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTopicSearchResultItem_aliases_topic_relationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicSearchResultItem_aliases_topic_relation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TopicSearchResultItem_aliases_topic_relation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TopicSearchResultItem_aliases_topic_relation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["relation_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationType(val)
        }
        return nil
    }
    res["topic_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *TopicSearchResultItem_aliases_topic_relation) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TopicSearchResultItem_aliases_topic_relation) GetName()(*string) {
    return m.name
}
// GetRelationType gets the relation_type property value. The relation_type property
// returns a *string when successful
func (m *TopicSearchResultItem_aliases_topic_relation) GetRelationType()(*string) {
    return m.relation_type
}
// GetTopicId gets the topic_id property value. The topic_id property
// returns a *int32 when successful
func (m *TopicSearchResultItem_aliases_topic_relation) GetTopicId()(*int32) {
    return m.topic_id
}
// Serialize serializes information the current object
func (m *TopicSearchResultItem_aliases_topic_relation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("relation_type", m.GetRelationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("topic_id", m.GetTopicId())
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
func (m *TopicSearchResultItem_aliases_topic_relation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *TopicSearchResultItem_aliases_topic_relation) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *TopicSearchResultItem_aliases_topic_relation) SetName(value *string)() {
    m.name = value
}
// SetRelationType sets the relation_type property value. The relation_type property
func (m *TopicSearchResultItem_aliases_topic_relation) SetRelationType(value *string)() {
    m.relation_type = value
}
// SetTopicId sets the topic_id property value. The topic_id property
func (m *TopicSearchResultItem_aliases_topic_relation) SetTopicId(value *int32)() {
    m.topic_id = value
}
type TopicSearchResultItem_aliases_topic_relationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetRelationType()(*string)
    GetTopicId()(*int32)
    SetId(value *int32)()
    SetName(value *string)()
    SetRelationType(value *string)()
    SetTopicId(value *int32)()
}
