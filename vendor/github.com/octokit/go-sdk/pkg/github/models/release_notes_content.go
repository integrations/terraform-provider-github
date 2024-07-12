package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReleaseNotesContent generated name and body describing a release
type ReleaseNotesContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The generated body describing the contents of the release supporting markdown formatting
    body *string
    // The generated name of the release
    name *string
}
// NewReleaseNotesContent instantiates a new ReleaseNotesContent and sets the default values.
func NewReleaseNotesContent()(*ReleaseNotesContent) {
    m := &ReleaseNotesContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReleaseNotesContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReleaseNotesContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReleaseNotesContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReleaseNotesContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The generated body describing the contents of the release supporting markdown formatting
// returns a *string when successful
func (m *ReleaseNotesContent) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReleaseNotesContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
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
    return res
}
// GetName gets the name property value. The generated name of the release
// returns a *string when successful
func (m *ReleaseNotesContent) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ReleaseNotesContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReleaseNotesContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The generated body describing the contents of the release supporting markdown formatting
func (m *ReleaseNotesContent) SetBody(value *string)() {
    m.body = value
}
// SetName sets the name property value. The generated name of the release
func (m *ReleaseNotesContent) SetName(value *string)() {
    m.name = value
}
type ReleaseNotesContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetName()(*string)
    SetBody(value *string)()
    SetName(value *string)()
}
