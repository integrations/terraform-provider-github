package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The discussion comment's body text.
    body *string
}
// NewItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody instantiates a new ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody and sets the default values.
func NewItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody()(*ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) {
    m := &ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The discussion comment's body text.
// returns a *string when successful
func (m *ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
func (m *ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The discussion comment's body text.
func (m *ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBody) SetBody(value *string)() {
    m.body = value
}
type ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    SetBody(value *string)()
}
