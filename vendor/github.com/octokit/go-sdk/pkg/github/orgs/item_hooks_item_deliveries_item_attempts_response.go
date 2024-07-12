package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemHooksItemDeliveriesItemAttemptsResponse 
// Deprecated: This class is obsolete. Use attemptsPostResponse instead.
type ItemHooksItemDeliveriesItemAttemptsResponse struct {
    ItemHooksItemDeliveriesItemAttemptsPostResponse
}
// NewItemHooksItemDeliveriesItemAttemptsResponse instantiates a new ItemHooksItemDeliveriesItemAttemptsResponse and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsResponse()(*ItemHooksItemDeliveriesItemAttemptsResponse) {
    m := &ItemHooksItemDeliveriesItemAttemptsResponse{
        ItemHooksItemDeliveriesItemAttemptsPostResponse: *NewItemHooksItemDeliveriesItemAttemptsPostResponse(),
    }
    return m
}
// CreateItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemHooksItemDeliveriesItemAttemptsResponse(), nil
}
// ItemHooksItemDeliveriesItemAttemptsResponseable 
// Deprecated: This class is obsolete. Use attemptsPostResponse instead.
type ItemHooksItemDeliveriesItemAttemptsResponseable interface {
    ItemHooksItemDeliveriesItemAttemptsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
