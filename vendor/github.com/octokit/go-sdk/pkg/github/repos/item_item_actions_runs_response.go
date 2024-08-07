package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunsResponse 
// Deprecated: This class is obsolete. Use runsGetResponse instead.
type ItemItemActionsRunsResponse struct {
    ItemItemActionsRunsGetResponse
}
// NewItemItemActionsRunsResponse instantiates a new ItemItemActionsRunsResponse and sets the default values.
func NewItemItemActionsRunsResponse()(*ItemItemActionsRunsResponse) {
    m := &ItemItemActionsRunsResponse{
        ItemItemActionsRunsGetResponse: *NewItemItemActionsRunsGetResponse(),
    }
    return m
}
// CreateItemItemActionsRunsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsResponse(), nil
}
// ItemItemActionsRunsResponseable 
// Deprecated: This class is obsolete. Use runsGetResponse instead.
type ItemItemActionsRunsResponseable interface {
    ItemItemActionsRunsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
