package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesResponse 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type ItemItemCodespacesResponse struct {
    ItemItemCodespacesGetResponse
}
// NewItemItemCodespacesResponse instantiates a new ItemItemCodespacesResponse and sets the default values.
func NewItemItemCodespacesResponse()(*ItemItemCodespacesResponse) {
    m := &ItemItemCodespacesResponse{
        ItemItemCodespacesGetResponse: *NewItemItemCodespacesGetResponse(),
    }
    return m
}
// CreateItemItemCodespacesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesResponse(), nil
}
// ItemItemCodespacesResponseable 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type ItemItemCodespacesResponseable interface {
    ItemItemCodespacesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
