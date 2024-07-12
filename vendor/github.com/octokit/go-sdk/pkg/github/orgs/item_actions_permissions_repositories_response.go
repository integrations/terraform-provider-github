package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsPermissionsRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemActionsPermissionsRepositoriesResponse struct {
    ItemActionsPermissionsRepositoriesGetResponse
}
// NewItemActionsPermissionsRepositoriesResponse instantiates a new ItemActionsPermissionsRepositoriesResponse and sets the default values.
func NewItemActionsPermissionsRepositoriesResponse()(*ItemActionsPermissionsRepositoriesResponse) {
    m := &ItemActionsPermissionsRepositoriesResponse{
        ItemActionsPermissionsRepositoriesGetResponse: *NewItemActionsPermissionsRepositoriesGetResponse(),
    }
    return m
}
// CreateItemActionsPermissionsRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsPermissionsRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsPermissionsRepositoriesResponse(), nil
}
// ItemActionsPermissionsRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemActionsPermissionsRepositoriesResponseable interface {
    ItemActionsPermissionsRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
