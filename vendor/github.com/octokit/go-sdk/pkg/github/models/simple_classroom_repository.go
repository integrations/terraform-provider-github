package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimpleClassroomRepository a GitHub repository view for Classroom
type SimpleClassroomRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The default branch for the repository.
    default_branch *string
    // The full, globally unique name of the repository.
    full_name *string
    // The URL to view the repository on GitHub.com.
    html_url *string
    // A unique identifier of the repository.
    id *int32
    // The GraphQL identifier of the repository.
    node_id *string
    // Whether the repository is private.
    private *bool
}
// NewSimpleClassroomRepository instantiates a new SimpleClassroomRepository and sets the default values.
func NewSimpleClassroomRepository()(*SimpleClassroomRepository) {
    m := &SimpleClassroomRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimpleClassroomRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimpleClassroomRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimpleClassroomRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SimpleClassroomRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultBranch gets the default_branch property value. The default branch for the repository.
// returns a *string when successful
func (m *SimpleClassroomRepository) GetDefaultBranch()(*string) {
    return m.default_branch
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimpleClassroomRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["default_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBranch(val)
        }
        return nil
    }
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    return res
}
// GetFullName gets the full_name property value. The full, globally unique name of the repository.
// returns a *string when successful
func (m *SimpleClassroomRepository) GetFullName()(*string) {
    return m.full_name
}
// GetHtmlUrl gets the html_url property value. The URL to view the repository on GitHub.com.
// returns a *string when successful
func (m *SimpleClassroomRepository) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. A unique identifier of the repository.
// returns a *int32 when successful
func (m *SimpleClassroomRepository) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The GraphQL identifier of the repository.
// returns a *string when successful
func (m *SimpleClassroomRepository) GetNodeId()(*string) {
    return m.node_id
}
// GetPrivate gets the private property value. Whether the repository is private.
// returns a *bool when successful
func (m *SimpleClassroomRepository) GetPrivate()(*bool) {
    return m.private
}
// Serialize serializes information the current object
func (m *SimpleClassroomRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("default_branch", m.GetDefaultBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("private", m.GetPrivate())
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
func (m *SimpleClassroomRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultBranch sets the default_branch property value. The default branch for the repository.
func (m *SimpleClassroomRepository) SetDefaultBranch(value *string)() {
    m.default_branch = value
}
// SetFullName sets the full_name property value. The full, globally unique name of the repository.
func (m *SimpleClassroomRepository) SetFullName(value *string)() {
    m.full_name = value
}
// SetHtmlUrl sets the html_url property value. The URL to view the repository on GitHub.com.
func (m *SimpleClassroomRepository) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. A unique identifier of the repository.
func (m *SimpleClassroomRepository) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The GraphQL identifier of the repository.
func (m *SimpleClassroomRepository) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPrivate sets the private property value. Whether the repository is private.
func (m *SimpleClassroomRepository) SetPrivate(value *bool)() {
    m.private = value
}
type SimpleClassroomRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultBranch()(*string)
    GetFullName()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetPrivate()(*bool)
    SetDefaultBranch(value *string)()
    SetFullName(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetPrivate(value *bool)()
}
