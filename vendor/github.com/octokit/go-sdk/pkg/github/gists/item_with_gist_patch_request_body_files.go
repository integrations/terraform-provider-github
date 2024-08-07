package gists

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemWithGist_PatchRequestBody_files the gist files to be updated, renamed, or deleted. Each `key` must match the current filename(including extension) of the targeted gist file. For example: `hello.py`.To delete a file, set the whole file to null. For example: `hello.py : null`. The file will also bedeleted if the specified object does not contain at least one of `content` or `filename`.
type ItemWithGist_PatchRequestBody_files struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemWithGist_PatchRequestBody_files instantiates a new ItemWithGist_PatchRequestBody_files and sets the default values.
func NewItemWithGist_PatchRequestBody_files()(*ItemWithGist_PatchRequestBody_files) {
    m := &ItemWithGist_PatchRequestBody_files{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithGist_PatchRequestBody_filesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithGist_PatchRequestBody_filesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithGist_PatchRequestBody_files(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithGist_PatchRequestBody_files) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithGist_PatchRequestBody_files) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemWithGist_PatchRequestBody_files) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemWithGist_PatchRequestBody_files) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemWithGist_PatchRequestBody_filesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
