package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Reason for approving or denying the request. Max 1024 characters.
    reason *string
}
// NewItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody instantiates a new ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody and sets the default values.
func NewItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody()(*ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) {
    m := &ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetReason gets the reason property value. Reason for approving or denying the request. Max 1024 characters.
// returns a *string when successful
func (m *ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) GetReason()(*string) {
    return m.reason
}
// Serialize serializes information the current object
func (m *ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReason sets the reason property value. Reason for approving or denying the request. Max 1024 characters.
func (m *ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBody) SetReason(value *string)() {
    m.reason = value
}
type ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReason()(*string)
    SetReason(value *string)()
}
