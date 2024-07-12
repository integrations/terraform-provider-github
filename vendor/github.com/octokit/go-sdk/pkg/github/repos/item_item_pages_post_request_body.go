package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPagesPostRequestBody the source branch and directory used to publish your Pages site.
type ItemItemPagesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The source branch and directory used to publish your Pages site.
    source ItemItemPagesPostRequestBody_sourceable
}
// NewItemItemPagesPostRequestBody instantiates a new ItemItemPagesPostRequestBody and sets the default values.
func NewItemItemPagesPostRequestBody()(*ItemItemPagesPostRequestBody) {
    m := &ItemItemPagesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPagesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPagesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPagesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPagesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemPagesPostRequestBody_sourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(ItemItemPagesPostRequestBody_sourceable))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source branch and directory used to publish your Pages site.
// returns a ItemItemPagesPostRequestBody_sourceable when successful
func (m *ItemItemPagesPostRequestBody) GetSource()(ItemItemPagesPostRequestBody_sourceable) {
    return m.source
}
// Serialize serializes information the current object
func (m *ItemItemPagesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("source", m.GetSource())
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
func (m *ItemItemPagesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSource sets the source property value. The source branch and directory used to publish your Pages site.
func (m *ItemItemPagesPostRequestBody) SetSource(value ItemItemPagesPostRequestBody_sourceable)() {
    m.source = value
}
type ItemItemPagesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSource()(ItemItemPagesPostRequestBody_sourceable)
    SetSource(value ItemItemPagesPostRequestBody_sourceable)()
}
