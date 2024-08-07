package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ColumnsItemMovesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The position of the column in a project. Can be one of: `first`, `last`, or `after:<column_id>` to place after the specified column.
    position *string
}
// NewColumnsItemMovesPostRequestBody instantiates a new ColumnsItemMovesPostRequestBody and sets the default values.
func NewColumnsItemMovesPostRequestBody()(*ColumnsItemMovesPostRequestBody) {
    m := &ColumnsItemMovesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnsItemMovesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateColumnsItemMovesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsItemMovesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ColumnsItemMovesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ColumnsItemMovesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetPosition gets the position property value. The position of the column in a project. Can be one of: `first`, `last`, or `after:<column_id>` to place after the specified column.
// returns a *string when successful
func (m *ColumnsItemMovesPostRequestBody) GetPosition()(*string) {
    return m.position
}
// Serialize serializes information the current object
func (m *ColumnsItemMovesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("position", m.GetPosition())
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
func (m *ColumnsItemMovesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPosition sets the position property value. The position of the column in a project. Can be one of: `first`, `last`, or `after:<column_id>` to place after the specified column.
func (m *ColumnsItemMovesPostRequestBody) SetPosition(value *string)() {
    m.position = value
}
type ColumnsItemMovesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPosition()(*string)
    SetPosition(value *string)()
}
