package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Classroom a GitHub Classroom classroom
type Classroom struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether classroom is archived.
    archived *bool
    // Unique identifier of the classroom.
    id *int32
    // The name of the classroom.
    name *string
    // A GitHub organization.
    organization SimpleClassroomOrganizationable
    // The URL of the classroom on GitHub Classroom.
    url *string
}
// NewClassroom instantiates a new Classroom and sets the default values.
func NewClassroom()(*Classroom) {
    m := &Classroom{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClassroomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClassroomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassroom(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Classroom) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchived gets the archived property value. Whether classroom is archived.
// returns a *bool when successful
func (m *Classroom) GetArchived()(*bool) {
    return m.archived
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Classroom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleClassroomOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(SimpleClassroomOrganizationable))
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
func (m *Classroom) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the classroom.
// returns a *string when successful
func (m *Classroom) GetName()(*string) {
    return m.name
}
// GetOrganization gets the organization property value. A GitHub organization.
// returns a SimpleClassroomOrganizationable when successful
func (m *Classroom) GetOrganization()(SimpleClassroomOrganizationable) {
    return m.organization
}
// GetUrl gets the url property value. The URL of the classroom on GitHub Classroom.
// returns a *string when successful
func (m *Classroom) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Classroom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("organization", m.GetOrganization())
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
func (m *Classroom) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchived sets the archived property value. Whether classroom is archived.
func (m *Classroom) SetArchived(value *bool)() {
    m.archived = value
}
// SetId sets the id property value. Unique identifier of the classroom.
func (m *Classroom) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the classroom.
func (m *Classroom) SetName(value *string)() {
    m.name = value
}
// SetOrganization sets the organization property value. A GitHub organization.
func (m *Classroom) SetOrganization(value SimpleClassroomOrganizationable)() {
    m.organization = value
}
// SetUrl sets the url property value. The URL of the classroom on GitHub Classroom.
func (m *Classroom) SetUrl(value *string)() {
    m.url = value
}
type Classroomable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchived()(*bool)
    GetId()(*int32)
    GetName()(*string)
    GetOrganization()(SimpleClassroomOrganizationable)
    GetUrl()(*string)
    SetArchived(value *bool)()
    SetId(value *int32)()
    SetName(value *string)()
    SetOrganization(value SimpleClassroomOrganizationable)()
    SetUrl(value *string)()
}
