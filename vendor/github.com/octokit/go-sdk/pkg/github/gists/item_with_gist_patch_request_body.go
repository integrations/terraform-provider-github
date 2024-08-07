package gists

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWithGist_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the gist.
    description *string
    // The gist files to be updated, renamed, or deleted. Each `key` must match the current filename(including extension) of the targeted gist file. For example: `hello.py`.To delete a file, set the whole file to null. For example: `hello.py : null`. The file will also bedeleted if the specified object does not contain at least one of `content` or `filename`.
    files ItemWithGist_PatchRequestBody_filesable
}
// NewItemWithGist_PatchRequestBody instantiates a new ItemWithGist_PatchRequestBody and sets the default values.
func NewItemWithGist_PatchRequestBody()(*ItemWithGist_PatchRequestBody) {
    m := &ItemWithGist_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithGist_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithGist_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithGist_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithGist_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of the gist.
// returns a *string when successful
func (m *ItemWithGist_PatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithGist_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateItemWithGist_PatchRequestBody_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(ItemWithGist_PatchRequestBody_filesable))
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The gist files to be updated, renamed, or deleted. Each `key` must match the current filename(including extension) of the targeted gist file. For example: `hello.py`.To delete a file, set the whole file to null. For example: `hello.py : null`. The file will also bedeleted if the specified object does not contain at least one of `content` or `filename`.
// returns a ItemWithGist_PatchRequestBody_filesable when successful
func (m *ItemWithGist_PatchRequestBody) GetFiles()(ItemWithGist_PatchRequestBody_filesable) {
    return m.files
}
// Serialize serializes information the current object
func (m *ItemWithGist_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemWithGist_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of the gist.
func (m *ItemWithGist_PatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetFiles sets the files property value. The gist files to be updated, renamed, or deleted. Each `key` must match the current filename(including extension) of the targeted gist file. For example: `hello.py`.To delete a file, set the whole file to null. For example: `hello.py : null`. The file will also bedeleted if the specified object does not contain at least one of `content` or `filename`.
func (m *ItemWithGist_PatchRequestBody) SetFiles(value ItemWithGist_PatchRequestBody_filesable)() {
    m.files = value
}
type ItemWithGist_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetFiles()(ItemWithGist_PatchRequestBody_filesable)
    SetDescription(value *string)()
    SetFiles(value ItemWithGist_PatchRequestBody_filesable)()
}
