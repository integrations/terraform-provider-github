package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesDevcontainersResponse 
// Deprecated: This class is obsolete. Use devcontainersGetResponse instead.
type ItemItemCodespacesDevcontainersResponse struct {
    ItemItemCodespacesDevcontainersGetResponse
}
// NewItemItemCodespacesDevcontainersResponse instantiates a new ItemItemCodespacesDevcontainersResponse and sets the default values.
func NewItemItemCodespacesDevcontainersResponse()(*ItemItemCodespacesDevcontainersResponse) {
    m := &ItemItemCodespacesDevcontainersResponse{
        ItemItemCodespacesDevcontainersGetResponse: *NewItemItemCodespacesDevcontainersGetResponse(),
    }
    return m
}
// CreateItemItemCodespacesDevcontainersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesDevcontainersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesDevcontainersResponse(), nil
}
// ItemItemCodespacesDevcontainersResponseable 
// Deprecated: This class is obsolete. Use devcontainersGetResponse instead.
type ItemItemCodespacesDevcontainersResponseable interface {
    ItemItemCodespacesDevcontainersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
