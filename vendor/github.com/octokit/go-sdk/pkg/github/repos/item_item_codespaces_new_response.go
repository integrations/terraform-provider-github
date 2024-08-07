package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesNewResponse 
// Deprecated: This class is obsolete. Use newGetResponse instead.
type ItemItemCodespacesNewResponse struct {
    ItemItemCodespacesNewGetResponse
}
// NewItemItemCodespacesNewResponse instantiates a new ItemItemCodespacesNewResponse and sets the default values.
func NewItemItemCodespacesNewResponse()(*ItemItemCodespacesNewResponse) {
    m := &ItemItemCodespacesNewResponse{
        ItemItemCodespacesNewGetResponse: *NewItemItemCodespacesNewGetResponse(),
    }
    return m
}
// CreateItemItemCodespacesNewResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesNewResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesNewResponse(), nil
}
// ItemItemCodespacesNewResponseable 
// Deprecated: This class is obsolete. Use newGetResponse instead.
type ItemItemCodespacesNewResponseable interface {
    ItemItemCodespacesNewGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
