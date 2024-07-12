package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSelected_usersResponse the total number of seat assignments created.
// Deprecated: This class is obsolete. Use selected_usersPostResponse instead.
type ItemCopilotBillingSelected_usersResponse struct {
    ItemCopilotBillingSelected_usersPostResponse
}
// NewItemCopilotBillingSelected_usersResponse instantiates a new ItemCopilotBillingSelected_usersResponse and sets the default values.
func NewItemCopilotBillingSelected_usersResponse()(*ItemCopilotBillingSelected_usersResponse) {
    m := &ItemCopilotBillingSelected_usersResponse{
        ItemCopilotBillingSelected_usersPostResponse: *NewItemCopilotBillingSelected_usersPostResponse(),
    }
    return m
}
// CreateItemCopilotBillingSelected_usersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCopilotBillingSelected_usersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_usersResponse(), nil
}
// ItemCopilotBillingSelected_usersResponseable 
// Deprecated: This class is obsolete. Use selected_usersPostResponse instead.
type ItemCopilotBillingSelected_usersResponseable interface {
    ItemCopilotBillingSelected_usersPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
