package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodespacesResponse 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type ItemCodespacesResponse struct {
    ItemCodespacesGetResponse
}
// NewItemCodespacesResponse instantiates a new ItemCodespacesResponse and sets the default values.
func NewItemCodespacesResponse()(*ItemCodespacesResponse) {
    m := &ItemCodespacesResponse{
        ItemCodespacesGetResponse: *NewItemCodespacesGetResponse(),
    }
    return m
}
// CreateItemCodespacesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCodespacesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodespacesResponse(), nil
}
// ItemCodespacesResponseable 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type ItemCodespacesResponseable interface {
    ItemCodespacesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
