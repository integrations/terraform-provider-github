package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Max_file_path_length_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The maximum amount of characters allowed in file paths
    max_file_path_length *int32
}
// NewMax_file_path_length_parameters instantiates a new Max_file_path_length_parameters and sets the default values.
func NewMax_file_path_length_parameters()(*Max_file_path_length_parameters) {
    m := &Max_file_path_length_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMax_file_path_length_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMax_file_path_length_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMax_file_path_length_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Max_file_path_length_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Max_file_path_length_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["max_file_path_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxFilePathLength(val)
        }
        return nil
    }
    return res
}
// GetMaxFilePathLength gets the max_file_path_length property value. The maximum amount of characters allowed in file paths
// returns a *int32 when successful
func (m *Max_file_path_length_parameters) GetMaxFilePathLength()(*int32) {
    return m.max_file_path_length
}
// Serialize serializes information the current object
func (m *Max_file_path_length_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("max_file_path_length", m.GetMaxFilePathLength())
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
func (m *Max_file_path_length_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxFilePathLength sets the max_file_path_length property value. The maximum amount of characters allowed in file paths
func (m *Max_file_path_length_parameters) SetMaxFilePathLength(value *int32)() {
    m.max_file_path_length = value
}
type Max_file_path_length_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxFilePathLength()(*int32)
    SetMaxFilePathLength(value *int32)()
}
