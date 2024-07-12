package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsSecretsItemRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemActionsSecretsItemRepositoriesResponse struct {
    ItemActionsSecretsItemRepositoriesGetResponse
}
// NewItemActionsSecretsItemRepositoriesResponse instantiates a new ItemActionsSecretsItemRepositoriesResponse and sets the default values.
func NewItemActionsSecretsItemRepositoriesResponse()(*ItemActionsSecretsItemRepositoriesResponse) {
    m := &ItemActionsSecretsItemRepositoriesResponse{
        ItemActionsSecretsItemRepositoriesGetResponse: *NewItemActionsSecretsItemRepositoriesGetResponse(),
    }
    return m
}
// CreateItemActionsSecretsItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsSecretsItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsSecretsItemRepositoriesResponse(), nil
}
// ItemActionsSecretsItemRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemActionsSecretsItemRepositoriesResponseable interface {
    ItemActionsSecretsItemRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
