package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ColumnsItemCardsPostRequestBodyMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the content associated with the card
    content_id *int32
    // The piece of content associated with the card
    content_type *string
}
// NewColumnsItemCardsPostRequestBodyMember2 instantiates a new ColumnsItemCardsPostRequestBodyMember2 and sets the default values.
func NewColumnsItemCardsPostRequestBodyMember2()(*ColumnsItemCardsPostRequestBodyMember2) {
    m := &ColumnsItemCardsPostRequestBodyMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnsItemCardsPostRequestBodyMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateColumnsItemCardsPostRequestBodyMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsItemCardsPostRequestBodyMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ColumnsItemCardsPostRequestBodyMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentId gets the content_id property value. The unique identifier of the content associated with the card
// returns a *int32 when successful
func (m *ColumnsItemCardsPostRequestBodyMember2) GetContentId()(*int32) {
    return m.content_id
}
// GetContentType gets the content_type property value. The piece of content associated with the card
// returns a *string when successful
func (m *ColumnsItemCardsPostRequestBodyMember2) GetContentType()(*string) {
    return m.content_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ColumnsItemCardsPostRequestBodyMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentId(val)
        }
        return nil
    }
    res["content_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ColumnsItemCardsPostRequestBodyMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("content_id", m.GetContentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
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
func (m *ColumnsItemCardsPostRequestBodyMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentId sets the content_id property value. The unique identifier of the content associated with the card
func (m *ColumnsItemCardsPostRequestBodyMember2) SetContentId(value *int32)() {
    m.content_id = value
}
// SetContentType sets the content_type property value. The piece of content associated with the card
func (m *ColumnsItemCardsPostRequestBodyMember2) SetContentType(value *string)() {
    m.content_type = value
}
type ColumnsItemCardsPostRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentId()(*int32)
    GetContentType()(*string)
    SetContentId(value *int32)()
    SetContentType(value *string)()
}
