package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemOutside_collaboratorsItemWithUsernamePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // When set to `true`, the request will be performed asynchronously. Returns a 202 status code when the job is successfully queued.
    async *bool
}
// NewItemOutside_collaboratorsItemWithUsernamePutRequestBody instantiates a new ItemOutside_collaboratorsItemWithUsernamePutRequestBody and sets the default values.
func NewItemOutside_collaboratorsItemWithUsernamePutRequestBody()(*ItemOutside_collaboratorsItemWithUsernamePutRequestBody) {
    m := &ItemOutside_collaboratorsItemWithUsernamePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOutside_collaboratorsItemWithUsernamePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOutside_collaboratorsItemWithUsernamePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutside_collaboratorsItemWithUsernamePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemOutside_collaboratorsItemWithUsernamePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAsync gets the async property value. When set to `true`, the request will be performed asynchronously. Returns a 202 status code when the job is successfully queued.
// returns a *bool when successful
func (m *ItemOutside_collaboratorsItemWithUsernamePutRequestBody) GetAsync()(*bool) {
    return m.async
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemOutside_collaboratorsItemWithUsernamePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["async"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAsync(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemOutside_collaboratorsItemWithUsernamePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("async", m.GetAsync())
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
func (m *ItemOutside_collaboratorsItemWithUsernamePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAsync sets the async property value. When set to `true`, the request will be performed asynchronously. Returns a 202 status code when the job is successfully queued.
func (m *ItemOutside_collaboratorsItemWithUsernamePutRequestBody) SetAsync(value *bool)() {
    m.async = value
}
type ItemOutside_collaboratorsItemWithUsernamePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAsync()(*bool)
    SetAsync(value *bool)()
}
