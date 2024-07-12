package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMembersItemCodespacesResponse 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type ItemMembersItemCodespacesResponse struct {
    ItemMembersItemCodespacesGetResponse
}
// NewItemMembersItemCodespacesResponse instantiates a new ItemMembersItemCodespacesResponse and sets the default values.
func NewItemMembersItemCodespacesResponse()(*ItemMembersItemCodespacesResponse) {
    m := &ItemMembersItemCodespacesResponse{
        ItemMembersItemCodespacesGetResponse: *NewItemMembersItemCodespacesGetResponse(),
    }
    return m
}
// CreateItemMembersItemCodespacesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMembersItemCodespacesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMembersItemCodespacesResponse(), nil
}
// ItemMembersItemCodespacesResponseable 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type ItemMembersItemCodespacesResponseable interface {
    ItemMembersItemCodespacesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
