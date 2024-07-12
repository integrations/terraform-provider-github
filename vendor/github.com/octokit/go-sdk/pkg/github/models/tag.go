package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Tag tag
type Tag struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit property
    commit Tag_commitable
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The tarball_url property
    tarball_url *string
    // The zipball_url property
    zipball_url *string
}
// NewTag instantiates a new Tag and sets the default values.
func NewTag()(*Tag) {
    m := &Tag{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTag(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Tag) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommit gets the commit property value. The commit property
// returns a Tag_commitable when successful
func (m *Tag) GetCommit()(Tag_commitable) {
    return m.commit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Tag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTag_commitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(Tag_commitable))
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
    res["tarball_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarballUrl(val)
        }
        return nil
    }
    res["zipball_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZipballUrl(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Tag) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Tag) GetNodeId()(*string) {
    return m.node_id
}
// GetTarballUrl gets the tarball_url property value. The tarball_url property
// returns a *string when successful
func (m *Tag) GetTarballUrl()(*string) {
    return m.tarball_url
}
// GetZipballUrl gets the zipball_url property value. The zipball_url property
// returns a *string when successful
func (m *Tag) GetZipballUrl()(*string) {
    return m.zipball_url
}
// Serialize serializes information the current object
func (m *Tag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("commit", m.GetCommit())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tarball_url", m.GetTarballUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("zipball_url", m.GetZipballUrl())
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
func (m *Tag) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommit sets the commit property value. The commit property
func (m *Tag) SetCommit(value Tag_commitable)() {
    m.commit = value
}
// SetName sets the name property value. The name property
func (m *Tag) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Tag) SetNodeId(value *string)() {
    m.node_id = value
}
// SetTarballUrl sets the tarball_url property value. The tarball_url property
func (m *Tag) SetTarballUrl(value *string)() {
    m.tarball_url = value
}
// SetZipballUrl sets the zipball_url property value. The zipball_url property
func (m *Tag) SetZipballUrl(value *string)() {
    m.zipball_url = value
}
type Tagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(Tag_commitable)
    GetName()(*string)
    GetNodeId()(*string)
    GetTarballUrl()(*string)
    GetZipballUrl()(*string)
    SetCommit(value Tag_commitable)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetTarballUrl(value *string)()
    SetZipballUrl(value *string)()
}
