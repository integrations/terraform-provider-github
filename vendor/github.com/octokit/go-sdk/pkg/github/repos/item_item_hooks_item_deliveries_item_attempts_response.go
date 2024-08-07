package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemHooksItemDeliveriesItemAttemptsResponse 
// Deprecated: This class is obsolete. Use attemptsPostResponse instead.
type ItemItemHooksItemDeliveriesItemAttemptsResponse struct {
    ItemItemHooksItemDeliveriesItemAttemptsPostResponse
}
// NewItemItemHooksItemDeliveriesItemAttemptsResponse instantiates a new ItemItemHooksItemDeliveriesItemAttemptsResponse and sets the default values.
func NewItemItemHooksItemDeliveriesItemAttemptsResponse()(*ItemItemHooksItemDeliveriesItemAttemptsResponse) {
    m := &ItemItemHooksItemDeliveriesItemAttemptsResponse{
        ItemItemHooksItemDeliveriesItemAttemptsPostResponse: *NewItemItemHooksItemDeliveriesItemAttemptsPostResponse(),
    }
    return m
}
// CreateItemItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemHooksItemDeliveriesItemAttemptsResponse(), nil
}
// ItemItemHooksItemDeliveriesItemAttemptsResponseable 
// Deprecated: This class is obsolete. Use attemptsPostResponse instead.
type ItemItemHooksItemDeliveriesItemAttemptsResponseable interface {
    ItemItemHooksItemDeliveriesItemAttemptsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
