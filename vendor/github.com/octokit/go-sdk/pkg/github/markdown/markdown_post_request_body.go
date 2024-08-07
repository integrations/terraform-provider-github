package markdown

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MarkdownPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repository context to use when creating references in `gfm` mode.  For example, setting `context` to `octo-org/octo-repo` will change the text `#42` into an HTML link to issue 42 in the `octo-org/octo-repo` repository.
    context *string
    // The rendering mode.
    mode *MarkdownPostRequestBody_mode
    // The Markdown text to render in HTML.
    text *string
}
// NewMarkdownPostRequestBody instantiates a new MarkdownPostRequestBody and sets the default values.
func NewMarkdownPostRequestBody()(*MarkdownPostRequestBody) {
    m := &MarkdownPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    modeValue := MARKDOWN_MARKDOWNPOSTREQUESTBODY_MODE
    m.SetMode(&modeValue)
    return m
}
// CreateMarkdownPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarkdownPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarkdownPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MarkdownPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContext gets the context property value. The repository context to use when creating references in `gfm` mode.  For example, setting `context` to `octo-org/octo-repo` will change the text `#42` into an HTML link to issue 42 in the `octo-org/octo-repo` repository.
// returns a *string when successful
func (m *MarkdownPostRequestBody) GetContext()(*string) {
    return m.context
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarkdownPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMarkdownPostRequestBody_mode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*MarkdownPostRequestBody_mode))
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetMode gets the mode property value. The rendering mode.
// returns a *MarkdownPostRequestBody_mode when successful
func (m *MarkdownPostRequestBody) GetMode()(*MarkdownPostRequestBody_mode) {
    return m.mode
}
// GetText gets the text property value. The Markdown text to render in HTML.
// returns a *string when successful
func (m *MarkdownPostRequestBody) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *MarkdownPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
        err := writer.WriteStringValue("mode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *MarkdownPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContext sets the context property value. The repository context to use when creating references in `gfm` mode.  For example, setting `context` to `octo-org/octo-repo` will change the text `#42` into an HTML link to issue 42 in the `octo-org/octo-repo` repository.
func (m *MarkdownPostRequestBody) SetContext(value *string)() {
    m.context = value
}
// SetMode sets the mode property value. The rendering mode.
func (m *MarkdownPostRequestBody) SetMode(value *MarkdownPostRequestBody_mode)() {
    m.mode = value
}
// SetText sets the text property value. The Markdown text to render in HTML.
func (m *MarkdownPostRequestBody) SetText(value *string)() {
    m.text = value
}
type MarkdownPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(*string)
    GetMode()(*MarkdownPostRequestBody_mode)
    GetText()(*string)
    SetContext(value *string)()
    SetMode(value *MarkdownPostRequestBody_mode)()
    SetText(value *string)()
}
