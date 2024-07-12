package gists

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GistsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description of the gist
    description *string
    // Names and content for the files that make up the gist
    files GistsPostRequestBody_filesable
    // The public property
    public GistsPostRequestBody_GistsPostRequestBody_publicable
}
// GistsPostRequestBody_GistsPostRequestBody_public composed type wrapper for classes bool, string
type GistsPostRequestBody_GistsPostRequestBody_public struct {
    // Composed type representation for type bool
    boolean *bool
    // Composed type representation for type string
    string *string
}
// NewGistsPostRequestBody_GistsPostRequestBody_public instantiates a new GistsPostRequestBody_GistsPostRequestBody_public and sets the default values.
func NewGistsPostRequestBody_GistsPostRequestBody_public()(*GistsPostRequestBody_GistsPostRequestBody_public) {
    m := &GistsPostRequestBody_GistsPostRequestBody_public{
    }
    return m
}
// CreateGistsPostRequestBody_GistsPostRequestBody_publicFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistsPostRequestBody_GistsPostRequestBody_publicFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewGistsPostRequestBody_GistsPostRequestBody_public()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetBoolValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetBoolean(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBoolean gets the boolean property value. Composed type representation for type bool
// returns a *bool when successful
func (m *GistsPostRequestBody_GistsPostRequestBody_public) GetBoolean()(*bool) {
    return m.boolean
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistsPostRequestBody_GistsPostRequestBody_public) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *GistsPostRequestBody_GistsPostRequestBody_public) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *GistsPostRequestBody_GistsPostRequestBody_public) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *GistsPostRequestBody_GistsPostRequestBody_public) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBoolean() != nil {
        err := writer.WriteBoolValue("", m.GetBoolean())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBoolean sets the boolean property value. Composed type representation for type bool
func (m *GistsPostRequestBody_GistsPostRequestBody_public) SetBoolean(value *bool)() {
    m.boolean = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *GistsPostRequestBody_GistsPostRequestBody_public) SetString(value *string)() {
    m.string = value
}
type GistsPostRequestBody_GistsPostRequestBody_publicable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBoolean()(*bool)
    GetString()(*string)
    SetBoolean(value *bool)()
    SetString(value *string)()
}
// NewGistsPostRequestBody instantiates a new GistsPostRequestBody and sets the default values.
func NewGistsPostRequestBody()(*GistsPostRequestBody) {
    m := &GistsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GistsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Description of the gist
// returns a *string when successful
func (m *GistsPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistsPostRequestBody_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(GistsPostRequestBody_filesable))
        }
        return nil
    }
    res["public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistsPostRequestBody_GistsPostRequestBody_publicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublic(val.(GistsPostRequestBody_GistsPostRequestBody_publicable))
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. Names and content for the files that make up the gist
// returns a GistsPostRequestBody_filesable when successful
func (m *GistsPostRequestBody) GetFiles()(GistsPostRequestBody_filesable) {
    return m.files
}
// GetPublic gets the public property value. The public property
// returns a GistsPostRequestBody_GistsPostRequestBody_publicable when successful
func (m *GistsPostRequestBody) GetPublic()(GistsPostRequestBody_GistsPostRequestBody_publicable) {
    return m.public
}
// Serialize serializes information the current object
func (m *GistsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("files", m.GetFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("public", m.GetPublic())
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
func (m *GistsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description of the gist
func (m *GistsPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetFiles sets the files property value. Names and content for the files that make up the gist
func (m *GistsPostRequestBody) SetFiles(value GistsPostRequestBody_filesable)() {
    m.files = value
}
// SetPublic sets the public property value. The public property
func (m *GistsPostRequestBody) SetPublic(value GistsPostRequestBody_GistsPostRequestBody_publicable)() {
    m.public = value
}
type GistsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetFiles()(GistsPostRequestBody_filesable)
    GetPublic()(GistsPostRequestBody_GistsPostRequestBody_publicable)
    SetDescription(value *string)()
    SetFiles(value GistsPostRequestBody_filesable)()
    SetPublic(value GistsPostRequestBody_GistsPostRequestBody_publicable)()
}
