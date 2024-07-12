package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileCommit_content__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The git property
    git *string
    // The html property
    html *string
    // The self property
    self *string
}
// NewFileCommit_content__links instantiates a new FileCommit_content__links and sets the default values.
func NewFileCommit_content__links()(*FileCommit_content__links) {
    m := &FileCommit_content__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileCommit_content__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileCommit_content__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileCommit_content__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileCommit_content__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileCommit_content__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["git"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGit(val)
        }
        return nil
    }
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val)
        }
        return nil
    }
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val)
        }
        return nil
    }
    return res
}
// GetGit gets the git property value. The git property
// returns a *string when successful
func (m *FileCommit_content__links) GetGit()(*string) {
    return m.git
}
// GetHtml gets the html property value. The html property
// returns a *string when successful
func (m *FileCommit_content__links) GetHtml()(*string) {
    return m.html
}
// GetSelf gets the self property value. The self property
// returns a *string when successful
func (m *FileCommit_content__links) GetSelf()(*string) {
    return m.self
}
// Serialize serializes information the current object
func (m *FileCommit_content__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("git", m.GetGit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html", m.GetHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("self", m.GetSelf())
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
func (m *FileCommit_content__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGit sets the git property value. The git property
func (m *FileCommit_content__links) SetGit(value *string)() {
    m.git = value
}
// SetHtml sets the html property value. The html property
func (m *FileCommit_content__links) SetHtml(value *string)() {
    m.html = value
}
// SetSelf sets the self property value. The self property
func (m *FileCommit_content__links) SetSelf(value *string)() {
    m.self = value
}
type FileCommit_content__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGit()(*string)
    GetHtml()(*string)
    GetSelf()(*string)
    SetGit(value *string)()
    SetHtml(value *string)()
    SetSelf(value *string)()
}
