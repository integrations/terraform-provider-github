package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPersonalAccessTokenRequestsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique identifiers of the requests for access via fine-grained personal access token. Must be formed of between 1 and 100 `pat_request_id` values.
    pat_request_ids []int32
    // Reason for approving or denying the requests. Max 1024 characters.
    reason *string
}
// NewItemPersonalAccessTokenRequestsPostRequestBody instantiates a new ItemPersonalAccessTokenRequestsPostRequestBody and sets the default values.
func NewItemPersonalAccessTokenRequestsPostRequestBody()(*ItemPersonalAccessTokenRequestsPostRequestBody) {
    m := &ItemPersonalAccessTokenRequestsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPersonalAccessTokenRequestsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPersonalAccessTokenRequestsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPersonalAccessTokenRequestsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pat_request_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPatRequestIds(res)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    return res
}
// GetPatRequestIds gets the pat_request_ids property value. Unique identifiers of the requests for access via fine-grained personal access token. Must be formed of between 1 and 100 `pat_request_id` values.
// returns a []int32 when successful
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) GetPatRequestIds()([]int32) {
    return m.pat_request_ids
}
// GetReason gets the reason property value. Reason for approving or denying the requests. Max 1024 characters.
// returns a *string when successful
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) GetReason()(*string) {
    return m.reason
}
// Serialize serializes information the current object
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPatRequestIds() != nil {
        err := writer.WriteCollectionOfInt32Values("pat_request_ids", m.GetPatRequestIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
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
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPatRequestIds sets the pat_request_ids property value. Unique identifiers of the requests for access via fine-grained personal access token. Must be formed of between 1 and 100 `pat_request_id` values.
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) SetPatRequestIds(value []int32)() {
    m.pat_request_ids = value
}
// SetReason sets the reason property value. Reason for approving or denying the requests. Max 1024 characters.
func (m *ItemPersonalAccessTokenRequestsPostRequestBody) SetReason(value *string)() {
    m.reason = value
}
type ItemPersonalAccessTokenRequestsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPatRequestIds()([]int32)
    GetReason()(*string)
    SetPatRequestIds(value []int32)()
    SetReason(value *string)()
}
