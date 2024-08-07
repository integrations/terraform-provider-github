package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSelected_usersPostResponse the total number of seat assignments created.
type ItemCopilotBillingSelected_usersPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The seats_created property
    seats_created *int32
}
// NewItemCopilotBillingSelected_usersPostResponse instantiates a new ItemCopilotBillingSelected_usersPostResponse and sets the default values.
func NewItemCopilotBillingSelected_usersPostResponse()(*ItemCopilotBillingSelected_usersPostResponse) {
    m := &ItemCopilotBillingSelected_usersPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSelected_usersPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCopilotBillingSelected_usersPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_usersPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCopilotBillingSelected_usersPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCopilotBillingSelected_usersPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["seats_created"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeatsCreated(val)
        }
        return nil
    }
    return res
}
// GetSeatsCreated gets the seats_created property value. The seats_created property
// returns a *int32 when successful
func (m *ItemCopilotBillingSelected_usersPostResponse) GetSeatsCreated()(*int32) {
    return m.seats_created
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSelected_usersPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("seats_created", m.GetSeatsCreated())
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
func (m *ItemCopilotBillingSelected_usersPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSeatsCreated sets the seats_created property value. The seats_created property
func (m *ItemCopilotBillingSelected_usersPostResponse) SetSeatsCreated(value *int32)() {
    m.seats_created = value
}
type ItemCopilotBillingSelected_usersPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSeatsCreated()(*int32)
    SetSeatsCreated(value *int32)()
}
