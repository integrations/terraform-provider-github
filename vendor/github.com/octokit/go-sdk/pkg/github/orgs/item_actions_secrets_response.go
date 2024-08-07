package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemActionsSecretsResponse struct {
    ItemActionsSecretsGetResponse
}
// NewItemActionsSecretsResponse instantiates a new ItemActionsSecretsResponse and sets the default values.
func NewItemActionsSecretsResponse()(*ItemActionsSecretsResponse) {
    m := &ItemActionsSecretsResponse{
        ItemActionsSecretsGetResponse: *NewItemActionsSecretsGetResponse(),
    }
    return m
}
// CreateItemActionsSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsSecretsResponse(), nil
}
// ItemActionsSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemActionsSecretsResponseable interface {
    ItemActionsSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
