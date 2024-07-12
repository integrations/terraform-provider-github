package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ColumnsCardsItemWithCard_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether or not the card is archived
    archived *bool
    // The project card's note
    note *string
}
// NewColumnsCardsItemWithCard_PatchRequestBody instantiates a new ColumnsCardsItemWithCard_PatchRequestBody and sets the default values.
func NewColumnsCardsItemWithCard_PatchRequestBody()(*ColumnsCardsItemWithCard_PatchRequestBody) {
    m := &ColumnsCardsItemWithCard_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnsCardsItemWithCard_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateColumnsCardsItemWithCard_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsCardsItemWithCard_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ColumnsCardsItemWithCard_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchived gets the archived property value. Whether or not the card is archived
// returns a *bool when successful
func (m *ColumnsCardsItemWithCard_PatchRequestBody) GetArchived()(*bool) {
    return m.archived
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ColumnsCardsItemWithCard_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchived(val)
        }
        return nil
    }
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    return res
}
// GetNote gets the note property value. The project card's note
// returns a *string when successful
func (m *ColumnsCardsItemWithCard_PatchRequestBody) GetNote()(*string) {
    return m.note
}
// Serialize serializes information the current object
func (m *ColumnsCardsItemWithCard_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("archived", m.GetArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note", m.GetNote())
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
func (m *ColumnsCardsItemWithCard_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchived sets the archived property value. Whether or not the card is archived
func (m *ColumnsCardsItemWithCard_PatchRequestBody) SetArchived(value *bool)() {
    m.archived = value
}
// SetNote sets the note property value. The project card's note
func (m *ColumnsCardsItemWithCard_PatchRequestBody) SetNote(value *string)() {
    m.note = value
}
type ColumnsCardsItemWithCard_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchived()(*bool)
    GetNote()(*string)
    SetArchived(value *bool)()
    SetNote(value *string)()
}
