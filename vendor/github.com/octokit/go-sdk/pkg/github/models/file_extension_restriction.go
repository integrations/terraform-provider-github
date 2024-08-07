package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// File_extension_restriction note: file_extension_restriction is in beta and subject to change.Prevent commits that include files with specified file extensions from being pushed to the commit graph.
type File_extension_restriction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters File_extension_restriction_parametersable
    // The type property
    typeEscaped *File_extension_restriction_type
}
// NewFile_extension_restriction instantiates a new File_extension_restriction and sets the default values.
func NewFile_extension_restriction()(*File_extension_restriction) {
    m := &File_extension_restriction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFile_extension_restrictionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFile_extension_restrictionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFile_extension_restriction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *File_extension_restriction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *File_extension_restriction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFile_extension_restriction_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(File_extension_restriction_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFile_extension_restriction_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*File_extension_restriction_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a File_extension_restriction_parametersable when successful
func (m *File_extension_restriction) GetParameters()(File_extension_restriction_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *File_extension_restriction_type when successful
func (m *File_extension_restriction) GetTypeEscaped()(*File_extension_restriction_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *File_extension_restriction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parameters", m.GetParameters())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *File_extension_restriction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *File_extension_restriction) SetParameters(value File_extension_restriction_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *File_extension_restriction) SetTypeEscaped(value *File_extension_restriction_type)() {
    m.typeEscaped = value
}
type File_extension_restrictionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(File_extension_restriction_parametersable)
    GetTypeEscaped()(*File_extension_restriction_type)
    SetParameters(value File_extension_restriction_parametersable)()
    SetTypeEscaped(value *File_extension_restriction_type)()
}
