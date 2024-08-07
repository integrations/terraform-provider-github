package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Blob blob
type Blob struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content *string
    // The encoding property
    encoding *string
    // The highlighted_content property
    highlighted_content *string
    // The node_id property
    node_id *string
    // The sha property
    sha *string
    // The size property
    size *int32
    // The url property
    url *string
}
// NewBlob instantiates a new Blob and sets the default values.
func NewBlob()(*Blob) {
    m := &Blob{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBlobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlob(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Blob) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
// returns a *string when successful
func (m *Blob) GetContent()(*string) {
    return m.content
}
// GetEncoding gets the encoding property value. The encoding property
// returns a *string when successful
func (m *Blob) GetEncoding()(*string) {
    return m.encoding
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Blob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["highlighted_content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighlightedContent(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHighlightedContent gets the highlighted_content property value. The highlighted_content property
// returns a *string when successful
func (m *Blob) GetHighlightedContent()(*string) {
    return m.highlighted_content
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Blob) GetNodeId()(*string) {
    return m.node_id
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *Blob) GetSha()(*string) {
    return m.sha
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *Blob) GetSize()(*int32) {
    return m.size
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Blob) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Blob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("highlighted_content", m.GetHighlightedContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *Blob) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *Blob) SetContent(value *string)() {
    m.content = value
}
// SetEncoding sets the encoding property value. The encoding property
func (m *Blob) SetEncoding(value *string)() {
    m.encoding = value
}
// SetHighlightedContent sets the highlighted_content property value. The highlighted_content property
func (m *Blob) SetHighlightedContent(value *string)() {
    m.highlighted_content = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Blob) SetNodeId(value *string)() {
    m.node_id = value
}
// SetSha sets the sha property value. The sha property
func (m *Blob) SetSha(value *string)() {
    m.sha = value
}
// SetSize sets the size property value. The size property
func (m *Blob) SetSize(value *int32)() {
    m.size = value
}
// SetUrl sets the url property value. The url property
func (m *Blob) SetUrl(value *string)() {
    m.url = value
}
type Blobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetEncoding()(*string)
    GetHighlightedContent()(*string)
    GetNodeId()(*string)
    GetSha()(*string)
    GetSize()(*int32)
    GetUrl()(*string)
    SetContent(value *string)()
    SetEncoding(value *string)()
    SetHighlightedContent(value *string)()
    SetNodeId(value *string)()
    SetSha(value *string)()
    SetSize(value *int32)()
    SetUrl(value *string)()
}
