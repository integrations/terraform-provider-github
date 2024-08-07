package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeownersErrors_errors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The column number where this errors occurs.
    column *int32
    // The type of error.
    kind *string
    // The line number where this errors occurs.
    line *int32
    // A human-readable description of the error, combining information from multiple fields, laid out for display in a monospaced typeface (for example, a command-line setting).
    message *string
    // The path of the file where the error occured.
    path *string
    // The contents of the line where the error occurs.
    source *string
    // Suggested action to fix the error. This will usually be `null`, but is provided for some common errors.
    suggestion *string
}
// NewCodeownersErrors_errors instantiates a new CodeownersErrors_errors and sets the default values.
func NewCodeownersErrors_errors()(*CodeownersErrors_errors) {
    m := &CodeownersErrors_errors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeownersErrors_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeownersErrors_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeownersErrors_errors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeownersErrors_errors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColumn gets the column property value. The column number where this errors occurs.
// returns a *int32 when successful
func (m *CodeownersErrors_errors) GetColumn()(*int32) {
    return m.column
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeownersErrors_errors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["column"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumn(val)
        }
        return nil
    }
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["suggestion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestion(val)
        }
        return nil
    }
    return res
}
// GetKind gets the kind property value. The type of error.
// returns a *string when successful
func (m *CodeownersErrors_errors) GetKind()(*string) {
    return m.kind
}
// GetLine gets the line property value. The line number where this errors occurs.
// returns a *int32 when successful
func (m *CodeownersErrors_errors) GetLine()(*int32) {
    return m.line
}
// GetMessage gets the message property value. A human-readable description of the error, combining information from multiple fields, laid out for display in a monospaced typeface (for example, a command-line setting).
// returns a *string when successful
func (m *CodeownersErrors_errors) GetMessage()(*string) {
    return m.message
}
// GetPath gets the path property value. The path of the file where the error occured.
// returns a *string when successful
func (m *CodeownersErrors_errors) GetPath()(*string) {
    return m.path
}
// GetSource gets the source property value. The contents of the line where the error occurs.
// returns a *string when successful
func (m *CodeownersErrors_errors) GetSource()(*string) {
    return m.source
}
// GetSuggestion gets the suggestion property value. Suggested action to fix the error. This will usually be `null`, but is provided for some common errors.
// returns a *string when successful
func (m *CodeownersErrors_errors) GetSuggestion()(*string) {
    return m.suggestion
}
// Serialize serializes information the current object
func (m *CodeownersErrors_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("column", m.GetColumn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("line", m.GetLine())
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
        err := writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("suggestion", m.GetSuggestion())
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
func (m *CodeownersErrors_errors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColumn sets the column property value. The column number where this errors occurs.
func (m *CodeownersErrors_errors) SetColumn(value *int32)() {
    m.column = value
}
// SetKind sets the kind property value. The type of error.
func (m *CodeownersErrors_errors) SetKind(value *string)() {
    m.kind = value
}
// SetLine sets the line property value. The line number where this errors occurs.
func (m *CodeownersErrors_errors) SetLine(value *int32)() {
    m.line = value
}
// SetMessage sets the message property value. A human-readable description of the error, combining information from multiple fields, laid out for display in a monospaced typeface (for example, a command-line setting).
func (m *CodeownersErrors_errors) SetMessage(value *string)() {
    m.message = value
}
// SetPath sets the path property value. The path of the file where the error occured.
func (m *CodeownersErrors_errors) SetPath(value *string)() {
    m.path = value
}
// SetSource sets the source property value. The contents of the line where the error occurs.
func (m *CodeownersErrors_errors) SetSource(value *string)() {
    m.source = value
}
// SetSuggestion sets the suggestion property value. Suggested action to fix the error. This will usually be `null`, but is provided for some common errors.
func (m *CodeownersErrors_errors) SetSuggestion(value *string)() {
    m.suggestion = value
}
type CodeownersErrors_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumn()(*int32)
    GetKind()(*string)
    GetLine()(*int32)
    GetMessage()(*string)
    GetPath()(*string)
    GetSource()(*string)
    GetSuggestion()(*string)
    SetColumn(value *int32)()
    SetKind(value *string)()
    SetLine(value *int32)()
    SetMessage(value *string)()
    SetPath(value *string)()
    SetSource(value *string)()
    SetSuggestion(value *string)()
}
