package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ColumnsCardsItemMoves403Error_errors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The field property
    field *string
    // The message property
    message *string
    // The resource property
    resource *string
}
// NewColumnsCardsItemMoves403Error_errors instantiates a new ColumnsCardsItemMoves403Error_errors and sets the default values.
func NewColumnsCardsItemMoves403Error_errors()(*ColumnsCardsItemMoves403Error_errors) {
    m := &ColumnsCardsItemMoves403Error_errors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnsCardsItemMoves403Error_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateColumnsCardsItemMoves403Error_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsCardsItemMoves403Error_errors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ColumnsCardsItemMoves403Error_errors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *ColumnsCardsItemMoves403Error_errors) GetCode()(*string) {
    return m.code
}
// GetField gets the field property value. The field property
// returns a *string when successful
func (m *ColumnsCardsItemMoves403Error_errors) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ColumnsCardsItemMoves403Error_errors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *ColumnsCardsItemMoves403Error_errors) GetMessage()(*string) {
    return m.message
}
// GetResource gets the resource property value. The resource property
// returns a *string when successful
func (m *ColumnsCardsItemMoves403Error_errors) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *ColumnsCardsItemMoves403Error_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("field", m.GetField())
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
        err := writer.WriteStringValue("resource", m.GetResource())
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
func (m *ColumnsCardsItemMoves403Error_errors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *ColumnsCardsItemMoves403Error_errors) SetCode(value *string)() {
    m.code = value
}
// SetField sets the field property value. The field property
func (m *ColumnsCardsItemMoves403Error_errors) SetField(value *string)() {
    m.field = value
}
// SetMessage sets the message property value. The message property
func (m *ColumnsCardsItemMoves403Error_errors) SetMessage(value *string)() {
    m.message = value
}
// SetResource sets the resource property value. The resource property
func (m *ColumnsCardsItemMoves403Error_errors) SetResource(value *string)() {
    m.resource = value
}
type ColumnsCardsItemMoves403Error_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetField()(*string)
    GetMessage()(*string)
    GetResource()(*string)
    SetCode(value *string)()
    SetField(value *string)()
    SetMessage(value *string)()
    SetResource(value *string)()
}
