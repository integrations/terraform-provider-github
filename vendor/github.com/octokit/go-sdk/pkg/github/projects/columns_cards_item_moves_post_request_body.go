package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ColumnsCardsItemMovesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the column the card should be moved to
    column_id *int32
    // The position of the card in a column. Can be one of: `top`, `bottom`, or `after:<card_id>` to place after the specified card.
    position *string
}
// NewColumnsCardsItemMovesPostRequestBody instantiates a new ColumnsCardsItemMovesPostRequestBody and sets the default values.
func NewColumnsCardsItemMovesPostRequestBody()(*ColumnsCardsItemMovesPostRequestBody) {
    m := &ColumnsCardsItemMovesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnsCardsItemMovesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateColumnsCardsItemMovesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsCardsItemMovesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ColumnsCardsItemMovesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColumnId gets the column_id property value. The unique identifier of the column the card should be moved to
// returns a *int32 when successful
func (m *ColumnsCardsItemMovesPostRequestBody) GetColumnId()(*int32) {
    return m.column_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ColumnsCardsItemMovesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["column_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnId(val)
        }
        return nil
    }
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
// GetPosition gets the position property value. The position of the card in a column. Can be one of: `top`, `bottom`, or `after:<card_id>` to place after the specified card.
// returns a *string when successful
func (m *ColumnsCardsItemMovesPostRequestBody) GetPosition()(*string) {
    return m.position
}
// Serialize serializes information the current object
func (m *ColumnsCardsItemMovesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("column_id", m.GetColumnId())
        if err != nil {
            return err
        }
    }
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
func (m *ColumnsCardsItemMovesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColumnId sets the column_id property value. The unique identifier of the column the card should be moved to
func (m *ColumnsCardsItemMovesPostRequestBody) SetColumnId(value *int32)() {
    m.column_id = value
}
// SetPosition sets the position property value. The position of the card in a column. Can be one of: `top`, `bottom`, or `after:<card_id>` to place after the specified card.
func (m *ColumnsCardsItemMovesPostRequestBody) SetPosition(value *string)() {
    m.position = value
}
type ColumnsCardsItemMovesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumnId()(*int32)
    GetPosition()(*string)
    SetColumnId(value *int32)()
    SetPosition(value *string)()
}
