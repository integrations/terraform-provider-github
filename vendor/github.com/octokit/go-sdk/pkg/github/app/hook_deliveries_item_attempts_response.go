package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HookDeliveriesItemAttemptsResponse 
// Deprecated: This class is obsolete. Use attemptsPostResponse instead.
type HookDeliveriesItemAttemptsResponse struct {
    HookDeliveriesItemAttemptsPostResponse
}
// NewHookDeliveriesItemAttemptsResponse instantiates a new HookDeliveriesItemAttemptsResponse and sets the default values.
func NewHookDeliveriesItemAttemptsResponse()(*HookDeliveriesItemAttemptsResponse) {
    m := &HookDeliveriesItemAttemptsResponse{
        HookDeliveriesItemAttemptsPostResponse: *NewHookDeliveriesItemAttemptsPostResponse(),
    }
    return m
}
// CreateHookDeliveriesItemAttemptsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHookDeliveriesItemAttemptsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHookDeliveriesItemAttemptsResponse(), nil
}
// HookDeliveriesItemAttemptsResponseable 
// Deprecated: This class is obsolete. Use attemptsPostResponse instead.
type HookDeliveriesItemAttemptsResponseable interface {
    HookDeliveriesItemAttemptsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
