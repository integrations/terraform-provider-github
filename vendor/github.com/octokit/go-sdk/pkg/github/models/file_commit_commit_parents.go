package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileCommit_commit_parents struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html_url property
    html_url *string
    // The sha property
    sha *string
    // The url property
    url *string
}
// NewFileCommit_commit_parents instantiates a new FileCommit_commit_parents and sets the default values.
func NewFileCommit_commit_parents()(*FileCommit_commit_parents) {
    m := &FileCommit_commit_parents{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileCommit_commit_parentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileCommit_commit_parentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileCommit_commit_parents(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileCommit_commit_parents) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileCommit_commit_parents) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *FileCommit_commit_parents) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *FileCommit_commit_parents) GetSha()(*string) {
    return m.sha
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *FileCommit_commit_parents) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *FileCommit_commit_parents) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
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
func (m *FileCommit_commit_parents) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *FileCommit_commit_parents) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetSha sets the sha property value. The sha property
func (m *FileCommit_commit_parents) SetSha(value *string)() {
    m.sha = value
}
// SetUrl sets the url property value. The url property
func (m *FileCommit_commit_parents) SetUrl(value *string)() {
    m.url = value
}
type FileCommit_commit_parentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtmlUrl()(*string)
    GetSha()(*string)
    GetUrl()(*string)
    SetHtmlUrl(value *string)()
    SetSha(value *string)()
    SetUrl(value *string)()
}
