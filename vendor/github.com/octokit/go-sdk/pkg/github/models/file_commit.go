package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileCommit file Commit
type FileCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit property
    commit FileCommit_commitable
    // The content property
    content FileCommit_contentable
}
// NewFileCommit instantiates a new FileCommit and sets the default values.
func NewFileCommit()(*FileCommit) {
    m := &FileCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommit gets the commit property value. The commit property
// returns a FileCommit_commitable when successful
func (m *FileCommit) GetCommit()(FileCommit_commitable) {
    return m.commit
}
// GetContent gets the content property value. The content property
// returns a FileCommit_contentable when successful
func (m *FileCommit) GetContent()(FileCommit_contentable) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileCommit_commitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(FileCommit_commitable))
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileCommit_contentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(FileCommit_contentable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FileCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("commit", m.GetCommit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("content", m.GetContent())
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
func (m *FileCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommit sets the commit property value. The commit property
func (m *FileCommit) SetCommit(value FileCommit_commitable)() {
    m.commit = value
}
// SetContent sets the content property value. The content property
func (m *FileCommit) SetContent(value FileCommit_contentable)() {
    m.content = value
}
type FileCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(FileCommit_commitable)
    GetContent()(FileCommit_contentable)
    SetCommit(value FileCommit_commitable)()
    SetContent(value FileCommit_contentable)()
}
