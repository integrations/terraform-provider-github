package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunnerLabel a label for a self hosted runner
type RunnerLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique identifier of the label.
    id *int32
    // Name of the label.
    name *string
    // The type of label. Read-only labels are applied automatically when the runner is configured.
    typeEscaped *RunnerLabel_type
}
// NewRunnerLabel instantiates a new RunnerLabel and sets the default values.
func NewRunnerLabel()(*RunnerLabel) {
    m := &RunnerLabel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRunnerLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunnerLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunnerLabel(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RunnerLabel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RunnerLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunnerLabel_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RunnerLabel_type))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier of the label.
// returns a *int32 when successful
func (m *RunnerLabel) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. Name of the label.
// returns a *string when successful
func (m *RunnerLabel) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. The type of label. Read-only labels are applied automatically when the runner is configured.
// returns a *RunnerLabel_type when successful
func (m *RunnerLabel) GetTypeEscaped()(*RunnerLabel_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RunnerLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *RunnerLabel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Unique identifier of the label.
func (m *RunnerLabel) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. Name of the label.
func (m *RunnerLabel) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. The type of label. Read-only labels are applied automatically when the runner is configured.
func (m *RunnerLabel) SetTypeEscaped(value *RunnerLabel_type)() {
    m.typeEscaped = value
}
type RunnerLabelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetTypeEscaped()(*RunnerLabel_type)
    SetId(value *int32)()
    SetName(value *string)()
    SetTypeEscaped(value *RunnerLabel_type)()
}
