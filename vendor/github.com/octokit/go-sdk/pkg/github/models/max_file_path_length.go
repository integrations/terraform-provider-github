package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Max_file_path_length note: max_file_path_length is in beta and subject to change.Prevent commits that include file paths that exceed a specified character limit from being pushed to the commit graph.
type Max_file_path_length struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters Max_file_path_length_parametersable
    // The type property
    typeEscaped *Max_file_path_length_type
}
// NewMax_file_path_length instantiates a new Max_file_path_length and sets the default values.
func NewMax_file_path_length()(*Max_file_path_length) {
    m := &Max_file_path_length{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMax_file_path_lengthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMax_file_path_lengthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMax_file_path_length(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Max_file_path_length) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Max_file_path_length) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMax_file_path_length_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(Max_file_path_length_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMax_file_path_length_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*Max_file_path_length_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a Max_file_path_length_parametersable when successful
func (m *Max_file_path_length) GetParameters()(Max_file_path_length_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *Max_file_path_length_type when successful
func (m *Max_file_path_length) GetTypeEscaped()(*Max_file_path_length_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Max_file_path_length) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parameters", m.GetParameters())
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
func (m *Max_file_path_length) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *Max_file_path_length) SetParameters(value Max_file_path_length_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Max_file_path_length) SetTypeEscaped(value *Max_file_path_length_type)() {
    m.typeEscaped = value
}
type Max_file_path_lengthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(Max_file_path_length_parametersable)
    GetTypeEscaped()(*Max_file_path_length_type)
    SetParameters(value Max_file_path_length_parametersable)()
    SetTypeEscaped(value *Max_file_path_length_type)()
}
