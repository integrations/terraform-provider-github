package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemGitBlobsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The new blob's content.
    content *string
    // The encoding used for `content`. Currently, `"utf-8"` and `"base64"` are supported.
    encoding *string
}
// NewItemItemGitBlobsPostRequestBody instantiates a new ItemItemGitBlobsPostRequestBody and sets the default values.
func NewItemItemGitBlobsPostRequestBody()(*ItemItemGitBlobsPostRequestBody) {
    m := &ItemItemGitBlobsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    encodingValue := "utf-8"
    m.SetEncoding(&encodingValue)
    return m
}
// CreateItemItemGitBlobsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGitBlobsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitBlobsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGitBlobsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The new blob's content.
// returns a *string when successful
func (m *ItemItemGitBlobsPostRequestBody) GetContent()(*string) {
    return m.content
}
// GetEncoding gets the encoding property value. The encoding used for `content`. Currently, `"utf-8"` and `"base64"` are supported.
// returns a *string when successful
func (m *ItemItemGitBlobsPostRequestBody) GetEncoding()(*string) {
    return m.encoding
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGitBlobsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["encoding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncoding(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemGitBlobsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encoding", m.GetEncoding())
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
func (m *ItemItemGitBlobsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The new blob's content.
func (m *ItemItemGitBlobsPostRequestBody) SetContent(value *string)() {
    m.content = value
}
// SetEncoding sets the encoding property value. The encoding used for `content`. Currently, `"utf-8"` and `"base64"` are supported.
func (m *ItemItemGitBlobsPostRequestBody) SetEncoding(value *string)() {
    m.encoding = value
}
type ItemItemGitBlobsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetEncoding()(*string)
    SetContent(value *string)()
    SetEncoding(value *string)()
}
