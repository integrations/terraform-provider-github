package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimpleClassroom a GitHub Classroom classroom
type SimpleClassroom struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Returns whether classroom is archived or not.
    archived *bool
    // Unique identifier of the classroom.
    id *int32
    // The name of the classroom.
    name *string
    // The url of the classroom on GitHub Classroom.
    url *string
}
// NewSimpleClassroom instantiates a new SimpleClassroom and sets the default values.
func NewSimpleClassroom()(*SimpleClassroom) {
    m := &SimpleClassroom{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimpleClassroomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimpleClassroomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimpleClassroom(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SimpleClassroom) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchived gets the archived property value. Returns whether classroom is archived or not.
// returns a *bool when successful
func (m *SimpleClassroom) GetArchived()(*bool) {
    return m.archived
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimpleClassroom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchived(val)
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
// GetId gets the id property value. Unique identifier of the classroom.
// returns a *int32 when successful
func (m *SimpleClassroom) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the classroom.
// returns a *string when successful
func (m *SimpleClassroom) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The url of the classroom on GitHub Classroom.
// returns a *string when successful
func (m *SimpleClassroom) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *SimpleClassroom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("archived", m.GetArchived())
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
        err := writer.WriteStringValue("name", m.GetName())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimpleClassroom) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchived sets the archived property value. Returns whether classroom is archived or not.
func (m *SimpleClassroom) SetArchived(value *bool)() {
    m.archived = value
}
// SetId sets the id property value. Unique identifier of the classroom.
func (m *SimpleClassroom) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the classroom.
func (m *SimpleClassroom) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The url of the classroom on GitHub Classroom.
func (m *SimpleClassroom) SetUrl(value *string)() {
    m.url = value
}
type SimpleClassroomable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchived()(*bool)
    GetId()(*int32)
    GetName()(*string)
    GetUrl()(*string)
    SetArchived(value *bool)()
    SetId(value *int32)()
    SetName(value *string)()
    SetUrl(value *string)()
}
