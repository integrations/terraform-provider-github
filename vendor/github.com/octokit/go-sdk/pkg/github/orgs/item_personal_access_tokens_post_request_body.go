package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPersonalAccessTokensPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The IDs of the fine-grained personal access tokens.
    pat_ids []int32
}
// NewItemPersonalAccessTokensPostRequestBody instantiates a new ItemPersonalAccessTokensPostRequestBody and sets the default values.
func NewItemPersonalAccessTokensPostRequestBody()(*ItemPersonalAccessTokensPostRequestBody) {
    m := &ItemPersonalAccessTokensPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPersonalAccessTokensPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPersonalAccessTokensPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPersonalAccessTokensPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPersonalAccessTokensPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPersonalAccessTokensPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pat_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetPatIds(res)
        }
        return nil
    }
    return res
}
// GetPatIds gets the pat_ids property value. The IDs of the fine-grained personal access tokens.
// returns a []int32 when successful
func (m *ItemPersonalAccessTokensPostRequestBody) GetPatIds()([]int32) {
    return m.pat_ids
}
// Serialize serializes information the current object
func (m *ItemPersonalAccessTokensPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPatIds() != nil {
        err := writer.WriteCollectionOfInt32Values("pat_ids", m.GetPatIds())
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
func (m *ItemPersonalAccessTokensPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPatIds sets the pat_ids property value. The IDs of the fine-grained personal access tokens.
func (m *ItemPersonalAccessTokensPostRequestBody) SetPatIds(value []int32)() {
    m.pat_ids = value
}
type ItemPersonalAccessTokensPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPatIds()([]int32)
    SetPatIds(value []int32)()
}
