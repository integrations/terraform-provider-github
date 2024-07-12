package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GistComment403Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The block property
    block GistComment403Error_blockable
    // The documentation_url property
    documentation_url *string
    // The message property
    message *string
}
// NewGistComment403Error instantiates a new GistComment403Error and sets the default values.
func NewGistComment403Error()(*GistComment403Error) {
    m := &GistComment403Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistComment403ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistComment403ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistComment403Error(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *GistComment403Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GistComment403Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlock gets the block property value. The block property
// returns a GistComment403Error_blockable when successful
func (m *GistComment403Error) GetBlock()(GistComment403Error_blockable) {
    return m.block
}
// GetDocumentationUrl gets the documentation_url property value. The documentation_url property
// returns a *string when successful
func (m *GistComment403Error) GetDocumentationUrl()(*string) {
    return m.documentation_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistComment403Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["block"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistComment403Error_blockFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlock(val.(GistComment403Error_blockable))
        }
        return nil
    }
    res["documentation_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
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
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *GistComment403Error) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *GistComment403Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("block", m.GetBlock())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentation_url", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *GistComment403Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlock sets the block property value. The block property
func (m *GistComment403Error) SetBlock(value GistComment403Error_blockable)() {
    m.block = value
}
// SetDocumentationUrl sets the documentation_url property value. The documentation_url property
func (m *GistComment403Error) SetDocumentationUrl(value *string)() {
    m.documentation_url = value
}
// SetMessage sets the message property value. The message property
func (m *GistComment403Error) SetMessage(value *string)() {
    m.message = value
}
type GistComment403Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlock()(GistComment403Error_blockable)
    GetDocumentationUrl()(*string)
    GetMessage()(*string)
    SetBlock(value GistComment403Error_blockable)()
    SetDocumentationUrl(value *string)()
    SetMessage(value *string)()
}
