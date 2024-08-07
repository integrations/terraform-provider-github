package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemItemActionsSecretsResponse struct {
    ItemItemActionsSecretsGetResponse
}
// NewItemItemActionsSecretsResponse instantiates a new ItemItemActionsSecretsResponse and sets the default values.
func NewItemItemActionsSecretsResponse()(*ItemItemActionsSecretsResponse) {
    m := &ItemItemActionsSecretsResponse{
        ItemItemActionsSecretsGetResponse: *NewItemItemActionsSecretsGetResponse(),
    }
    return m
}
// CreateItemItemActionsSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsSecretsResponse(), nil
}
// ItemItemActionsSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemItemActionsSecretsResponseable interface {
    ItemItemActionsSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
