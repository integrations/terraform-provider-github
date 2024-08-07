package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitRef git references within a repository
type GitRef struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The node_id property
    node_id *string
    // The object property
    object GitRef_objectable
    // The ref property
    ref *string
    // The url property
    url *string
}
// NewGitRef instantiates a new GitRef and sets the default values.
func NewGitRef()(*GitRef) {
    m := &GitRef{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitRefFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitRefFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitRef(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GitRef) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitRef) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["object"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitRef_objectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObject(val.(GitRef_objectable))
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
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
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *GitRef) GetNodeId()(*string) {
    return m.node_id
}
// GetObject gets the object property value. The object property
// returns a GitRef_objectable when successful
func (m *GitRef) GetObject()(GitRef_objectable) {
    return m.object
}
// GetRef gets the ref property value. The ref property
// returns a *string when successful
func (m *GitRef) GetRef()(*string) {
    return m.ref
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GitRef) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *GitRef) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("object", m.GetObject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
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
func (m *GitRef) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *GitRef) SetNodeId(value *string)() {
    m.node_id = value
}
// SetObject sets the object property value. The object property
func (m *GitRef) SetObject(value GitRef_objectable)() {
    m.object = value
}
// SetRef sets the ref property value. The ref property
func (m *GitRef) SetRef(value *string)() {
    m.ref = value
}
// SetUrl sets the url property value. The url property
func (m *GitRef) SetUrl(value *string)() {
    m.url = value
}
type GitRefable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodeId()(*string)
    GetObject()(GitRef_objectable)
    GetRef()(*string)
    GetUrl()(*string)
    SetNodeId(value *string)()
    SetObject(value GitRef_objectable)()
    SetRef(value *string)()
    SetUrl(value *string)()
}
