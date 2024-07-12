package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsVariablesItemRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemActionsVariablesItemRepositoriesResponse struct {
    ItemActionsVariablesItemRepositoriesGetResponse
}
// NewItemActionsVariablesItemRepositoriesResponse instantiates a new ItemActionsVariablesItemRepositoriesResponse and sets the default values.
func NewItemActionsVariablesItemRepositoriesResponse()(*ItemActionsVariablesItemRepositoriesResponse) {
    m := &ItemActionsVariablesItemRepositoriesResponse{
        ItemActionsVariablesItemRepositoriesGetResponse: *NewItemActionsVariablesItemRepositoriesGetResponse(),
    }
    return m
}
// CreateItemActionsVariablesItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsVariablesItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsVariablesItemRepositoriesResponse(), nil
}
// ItemActionsVariablesItemRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemActionsVariablesItemRepositoriesResponseable interface {
    ItemActionsVariablesItemRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
