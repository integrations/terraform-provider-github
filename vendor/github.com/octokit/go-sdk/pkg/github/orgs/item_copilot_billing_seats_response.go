package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSeatsResponse 
// Deprecated: This class is obsolete. Use seatsGetResponse instead.
type ItemCopilotBillingSeatsResponse struct {
    ItemCopilotBillingSeatsGetResponse
}
// NewItemCopilotBillingSeatsResponse instantiates a new ItemCopilotBillingSeatsResponse and sets the default values.
func NewItemCopilotBillingSeatsResponse()(*ItemCopilotBillingSeatsResponse) {
    m := &ItemCopilotBillingSeatsResponse{
        ItemCopilotBillingSeatsGetResponse: *NewItemCopilotBillingSeatsGetResponse(),
    }
    return m
}
// CreateItemCopilotBillingSeatsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCopilotBillingSeatsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSeatsResponse(), nil
}
// ItemCopilotBillingSeatsResponseable 
// Deprecated: This class is obsolete. Use seatsGetResponse instead.
type ItemCopilotBillingSeatsResponseable interface {
    ItemCopilotBillingSeatsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
