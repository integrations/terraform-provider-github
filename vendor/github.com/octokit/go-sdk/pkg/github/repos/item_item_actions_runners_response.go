package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunnersResponse 
// Deprecated: This class is obsolete. Use runnersGetResponse instead.
type ItemItemActionsRunnersResponse struct {
    ItemItemActionsRunnersGetResponse
}
// NewItemItemActionsRunnersResponse instantiates a new ItemItemActionsRunnersResponse and sets the default values.
func NewItemItemActionsRunnersResponse()(*ItemItemActionsRunnersResponse) {
    m := &ItemItemActionsRunnersResponse{
        ItemItemActionsRunnersGetResponse: *NewItemItemActionsRunnersGetResponse(),
    }
    return m
}
// CreateItemItemActionsRunnersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunnersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunnersResponse(), nil
}
// ItemItemActionsRunnersResponseable 
// Deprecated: This class is obsolete. Use runnersGetResponse instead.
type ItemItemActionsRunnersResponseable interface {
    ItemItemActionsRunnersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
