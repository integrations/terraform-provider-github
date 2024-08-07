package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Thread_subject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The latest_comment_url property
    latest_comment_url *string
    // The title property
    title *string
    // The type property
    typeEscaped *string
    // The url property
    url *string
}
// NewThread_subject instantiates a new Thread_subject and sets the default values.
func NewThread_subject()(*Thread_subject) {
    m := &Thread_subject{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateThread_subjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThread_subjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThread_subject(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Thread_subject) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Thread_subject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["latest_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestCommentUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
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
// GetLatestCommentUrl gets the latest_comment_url property value. The latest_comment_url property
// returns a *string when successful
func (m *Thread_subject) GetLatestCommentUrl()(*string) {
    return m.latest_comment_url
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *Thread_subject) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Thread_subject) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Thread_subject) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Thread_subject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("latest_comment_url", m.GetLatestCommentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *Thread_subject) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLatestCommentUrl sets the latest_comment_url property value. The latest_comment_url property
func (m *Thread_subject) SetLatestCommentUrl(value *string)() {
    m.latest_comment_url = value
}
// SetTitle sets the title property value. The title property
func (m *Thread_subject) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Thread_subject) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *Thread_subject) SetUrl(value *string)() {
    m.url = value
}
type Thread_subjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLatestCommentUrl()(*string)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetLatestCommentUrl(value *string)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
