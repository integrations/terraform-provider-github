package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemGitTagsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tag message.
    message *string
    // The SHA of the git object this is tagging.
    object *string
    // The tag's name. This is typically a version (e.g., "v0.0.1").
    tag *string
    // An object with information about the individual creating the tag.
    tagger ItemItemGitTagsPostRequestBody_taggerable
}
// NewItemItemGitTagsPostRequestBody instantiates a new ItemItemGitTagsPostRequestBody and sets the default values.
func NewItemItemGitTagsPostRequestBody()(*ItemItemGitTagsPostRequestBody) {
    m := &ItemItemGitTagsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGitTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGitTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitTagsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGitTagsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGitTagsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["object"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObject(val)
        }
        return nil
    }
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val)
        }
        return nil
    }
    res["tagger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemGitTagsPostRequestBody_taggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagger(val.(ItemItemGitTagsPostRequestBody_taggerable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The tag message.
// returns a *string when successful
func (m *ItemItemGitTagsPostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetObject gets the object property value. The SHA of the git object this is tagging.
// returns a *string when successful
func (m *ItemItemGitTagsPostRequestBody) GetObject()(*string) {
    return m.object
}
// GetTag gets the tag property value. The tag's name. This is typically a version (e.g., "v0.0.1").
// returns a *string when successful
func (m *ItemItemGitTagsPostRequestBody) GetTag()(*string) {
    return m.tag
}
// GetTagger gets the tagger property value. An object with information about the individual creating the tag.
// returns a ItemItemGitTagsPostRequestBody_taggerable when successful
func (m *ItemItemGitTagsPostRequestBody) GetTagger()(ItemItemGitTagsPostRequestBody_taggerable) {
    return m.tagger
}
// Serialize serializes information the current object
func (m *ItemItemGitTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("object", m.GetObject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tagger", m.GetTagger())
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
func (m *ItemItemGitTagsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. The tag message.
func (m *ItemItemGitTagsPostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetObject sets the object property value. The SHA of the git object this is tagging.
func (m *ItemItemGitTagsPostRequestBody) SetObject(value *string)() {
    m.object = value
}
// SetTag sets the tag property value. The tag's name. This is typically a version (e.g., "v0.0.1").
func (m *ItemItemGitTagsPostRequestBody) SetTag(value *string)() {
    m.tag = value
}
// SetTagger sets the tagger property value. An object with information about the individual creating the tag.
func (m *ItemItemGitTagsPostRequestBody) SetTagger(value ItemItemGitTagsPostRequestBody_taggerable)() {
    m.tagger = value
}
type ItemItemGitTagsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    GetObject()(*string)
    GetTag()(*string)
    GetTagger()(ItemItemGitTagsPostRequestBody_taggerable)
    SetMessage(value *string)()
    SetObject(value *string)()
    SetTag(value *string)()
    SetTagger(value ItemItemGitTagsPostRequestBody_taggerable)()
}
