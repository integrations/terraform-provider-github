package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAlertLocation describe a region within a file for the alert.
type CodeScanningAlertLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The end_column property
    end_column *int32
    // The end_line property
    end_line *int32
    // The path property
    path *string
    // The start_column property
    start_column *int32
    // The start_line property
    start_line *int32
}
// NewCodeScanningAlertLocation instantiates a new CodeScanningAlertLocation and sets the default values.
func NewCodeScanningAlertLocation()(*CodeScanningAlertLocation) {
    m := &CodeScanningAlertLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAlertLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAlertLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAlertLocation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAlertLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndColumn gets the end_column property value. The end_column property
// returns a *int32 when successful
func (m *CodeScanningAlertLocation) GetEndColumn()(*int32) {
    return m.end_column
}
// GetEndLine gets the end_line property value. The end_line property
// returns a *int32 when successful
func (m *CodeScanningAlertLocation) GetEndLine()(*int32) {
    return m.end_line
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAlertLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *CodeScanningAlertLocation) GetPath()(*string) {
    return m.path
}
// GetStartColumn gets the start_column property value. The start_column property
// returns a *int32 when successful
func (m *CodeScanningAlertLocation) GetStartColumn()(*int32) {
    return m.start_column
}
// GetStartLine gets the start_line property value. The start_line property
// returns a *int32 when successful
func (m *CodeScanningAlertLocation) GetStartLine()(*int32) {
    return m.start_line
}
// Serialize serializes information the current object
func (m *CodeScanningAlertLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("path", m.GetPath())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningAlertLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndColumn sets the end_column property value. The end_column property
func (m *CodeScanningAlertLocation) SetEndColumn(value *int32)() {
    m.end_column = value
}
// SetEndLine sets the end_line property value. The end_line property
func (m *CodeScanningAlertLocation) SetEndLine(value *int32)() {
    m.end_line = value
}
// SetPath sets the path property value. The path property
func (m *CodeScanningAlertLocation) SetPath(value *string)() {
    m.path = value
}
// SetStartColumn sets the start_column property value. The start_column property
func (m *CodeScanningAlertLocation) SetStartColumn(value *int32)() {
    m.start_column = value
}
// SetStartLine sets the start_line property value. The start_line property
func (m *CodeScanningAlertLocation) SetStartLine(value *int32)() {
    m.start_line = value
}
type CodeScanningAlertLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndColumn()(*int32)
    GetEndLine()(*int32)
    GetPath()(*string)
    GetStartColumn()(*int32)
    GetStartLine()(*int32)
    SetEndColumn(value *int32)()
    SetEndLine(value *int32)()
    SetPath(value *string)()
    SetStartColumn(value *int32)()
    SetStartLine(value *int32)()
}
