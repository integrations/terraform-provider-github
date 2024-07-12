package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSelected_teamsResponse the total number of seat assignments created.
// Deprecated: This class is obsolete. Use selected_teamsPostResponse instead.
type ItemCopilotBillingSelected_teamsResponse struct {
    ItemCopilotBillingSelected_teamsPostResponse
}
// NewItemCopilotBillingSelected_teamsResponse instantiates a new ItemCopilotBillingSelected_teamsResponse and sets the default values.
func NewItemCopilotBillingSelected_teamsResponse()(*ItemCopilotBillingSelected_teamsResponse) {
    m := &ItemCopilotBillingSelected_teamsResponse{
        ItemCopilotBillingSelected_teamsPostResponse: *NewItemCopilotBillingSelected_teamsPostResponse(),
    }
    return m
}
// CreateItemCopilotBillingSelected_teamsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCopilotBillingSelected_teamsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_teamsResponse(), nil
}
// ItemCopilotBillingSelected_teamsResponseable 
// Deprecated: This class is obsolete. Use selected_teamsPostResponse instead.
type ItemCopilotBillingSelected_teamsResponseable interface {
    ItemCopilotBillingSelected_teamsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
