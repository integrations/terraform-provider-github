package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRunnersResponse 
// Deprecated: This class is obsolete. Use runnersGetResponse instead.
type ItemActionsRunnersResponse struct {
    ItemActionsRunnersGetResponse
}
// NewItemActionsRunnersResponse instantiates a new ItemActionsRunnersResponse and sets the default values.
func NewItemActionsRunnersResponse()(*ItemActionsRunnersResponse) {
    m := &ItemActionsRunnersResponse{
        ItemActionsRunnersGetResponse: *NewItemActionsRunnersGetResponse(),
    }
    return m
}
// CreateItemActionsRunnersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersResponse(), nil
}
// ItemActionsRunnersResponseable 
// Deprecated: This class is obsolete. Use runnersGetResponse instead.
type ItemActionsRunnersResponseable interface {
    ItemActionsRunnersGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
