package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckAnnotation check Annotation
type CheckAnnotation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The annotation_level property
    annotation_level *string
    // The blob_href property
    blob_href *string
    // The end_column property
    end_column *int32
    // The end_line property
    end_line *int32
    // The message property
    message *string
    // The path property
    path *string
    // The raw_details property
    raw_details *string
    // The start_column property
    start_column *int32
    // The start_line property
    start_line *int32
    // The title property
    title *string
}
// NewCheckAnnotation instantiates a new CheckAnnotation and sets the default values.
func NewCheckAnnotation()(*CheckAnnotation) {
    m := &CheckAnnotation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckAnnotationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckAnnotationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckAnnotation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckAnnotation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnnotationLevel gets the annotation_level property value. The annotation_level property
// returns a *string when successful
func (m *CheckAnnotation) GetAnnotationLevel()(*string) {
    return m.annotation_level
}
// GetBlobHref gets the blob_href property value. The blob_href property
// returns a *string when successful
func (m *CheckAnnotation) GetBlobHref()(*string) {
    return m.blob_href
}
// GetEndColumn gets the end_column property value. The end_column property
// returns a *int32 when successful
func (m *CheckAnnotation) GetEndColumn()(*int32) {
    return m.end_column
}
// GetEndLine gets the end_line property value. The end_line property
// returns a *int32 when successful
func (m *CheckAnnotation) GetEndLine()(*int32) {
    return m.end_line
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckAnnotation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["annotation_level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnnotationLevel(val)
        }
        return nil
    }
    res["blob_href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobHref(val)
        }
        return nil
    }
    res["end_column"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndColumn(val)
        }
        return nil
    }
    res["end_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndLine(val)
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["raw_details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawDetails(val)
        }
        return nil
    }
    res["start_column"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartColumn(val)
        }
        return nil
    }
    res["start_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartLine(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *CheckAnnotation) GetMessage()(*string) {
    return m.message
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *CheckAnnotation) GetPath()(*string) {
    return m.path
}
// GetRawDetails gets the raw_details property value. The raw_details property
// returns a *string when successful
func (m *CheckAnnotation) GetRawDetails()(*string) {
    return m.raw_details
}
// GetStartColumn gets the start_column property value. The start_column property
// returns a *int32 when successful
func (m *CheckAnnotation) GetStartColumn()(*int32) {
    return m.start_column
}
// GetStartLine gets the start_line property value. The start_line property
// returns a *int32 when successful
func (m *CheckAnnotation) GetStartLine()(*int32) {
    return m.start_line
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *CheckAnnotation) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *CheckAnnotation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("annotation_level", m.GetAnnotationLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blob_href", m.GetBlobHref())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("end_column", m.GetEndColumn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("end_line", m.GetEndLine())
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
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("raw_details", m.GetRawDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("start_column", m.GetStartColumn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("start_line", m.GetStartLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *CheckAnnotation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnnotationLevel sets the annotation_level property value. The annotation_level property
func (m *CheckAnnotation) SetAnnotationLevel(value *string)() {
    m.annotation_level = value
}
// SetBlobHref sets the blob_href property value. The blob_href property
func (m *CheckAnnotation) SetBlobHref(value *string)() {
    m.blob_href = value
}
// SetEndColumn sets the end_column property value. The end_column property
func (m *CheckAnnotation) SetEndColumn(value *int32)() {
    m.end_column = value
}
// SetEndLine sets the end_line property value. The end_line property
func (m *CheckAnnotation) SetEndLine(value *int32)() {
    m.end_line = value
}
// SetMessage sets the message property value. The message property
func (m *CheckAnnotation) SetMessage(value *string)() {
    m.message = value
}
// SetPath sets the path property value. The path property
func (m *CheckAnnotation) SetPath(value *string)() {
    m.path = value
}
// SetRawDetails sets the raw_details property value. The raw_details property
func (m *CheckAnnotation) SetRawDetails(value *string)() {
    m.raw_details = value
}
// SetStartColumn sets the start_column property value. The start_column property
func (m *CheckAnnotation) SetStartColumn(value *int32)() {
    m.start_column = value
}
// SetStartLine sets the start_line property value. The start_line property
func (m *CheckAnnotation) SetStartLine(value *int32)() {
    m.start_line = value
}
// SetTitle sets the title property value. The title property
func (m *CheckAnnotation) SetTitle(value *string)() {
    m.title = value
}
type CheckAnnotationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnnotationLevel()(*string)
    GetBlobHref()(*string)
    GetEndColumn()(*int32)
    GetEndLine()(*int32)
    GetMessage()(*string)
    GetPath()(*string)
    GetRawDetails()(*string)
    GetStartColumn()(*int32)
    GetStartLine()(*int32)
    GetTitle()(*string)
    SetAnnotationLevel(value *string)()
    SetBlobHref(value *string)()
    SetEndColumn(value *int32)()
    SetEndLine(value *int32)()
    SetMessage(value *string)()
    SetPath(value *string)()
    SetRawDetails(value *string)()
    SetStartColumn(value *int32)()
    SetStartLine(value *int32)()
    SetTitle(value *string)()
}
