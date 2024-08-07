package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CheckRun_output struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The annotations_count property
    annotations_count *int32
    // The annotations_url property
    annotations_url *string
    // The summary property
    summary *string
    // The text property
    text *string
    // The title property
    title *string
}
// NewCheckRun_output instantiates a new CheckRun_output and sets the default values.
func NewCheckRun_output()(*CheckRun_output) {
    m := &CheckRun_output{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckRun_outputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckRun_outputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckRun_output(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckRun_output) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnnotationsCount gets the annotations_count property value. The annotations_count property
// returns a *int32 when successful
func (m *CheckRun_output) GetAnnotationsCount()(*int32) {
    return m.annotations_count
}
// GetAnnotationsUrl gets the annotations_url property value. The annotations_url property
// returns a *string when successful
func (m *CheckRun_output) GetAnnotationsUrl()(*string) {
    return m.annotations_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckRun_output) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["annotations_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnnotationsCount(val)
        }
        return nil
    }
    res["annotations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnnotationsUrl(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
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
// GetSummary gets the summary property value. The summary property
// returns a *string when successful
func (m *CheckRun_output) GetSummary()(*string) {
    return m.summary
}
// GetText gets the text property value. The text property
// returns a *string when successful
func (m *CheckRun_output) GetText()(*string) {
    return m.text
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *CheckRun_output) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *CheckRun_output) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("annotations_count", m.GetAnnotationsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("annotations_url", m.GetAnnotationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
func (m *CheckRun_output) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnnotationsCount sets the annotations_count property value. The annotations_count property
func (m *CheckRun_output) SetAnnotationsCount(value *int32)() {
    m.annotations_count = value
}
// SetAnnotationsUrl sets the annotations_url property value. The annotations_url property
func (m *CheckRun_output) SetAnnotationsUrl(value *string)() {
    m.annotations_url = value
}
// SetSummary sets the summary property value. The summary property
func (m *CheckRun_output) SetSummary(value *string)() {
    m.summary = value
}
// SetText sets the text property value. The text property
func (m *CheckRun_output) SetText(value *string)() {
    m.text = value
}
// SetTitle sets the title property value. The title property
func (m *CheckRun_output) SetTitle(value *string)() {
    m.title = value
}
type CheckRun_outputable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnnotationsCount()(*int32)
    GetAnnotationsUrl()(*string)
    GetSummary()(*string)
    GetText()(*string)
    GetTitle()(*string)
    SetAnnotationsCount(value *int32)()
    SetAnnotationsUrl(value *string)()
    SetSummary(value *string)()
    SetText(value *string)()
    SetTitle(value *string)()
}
