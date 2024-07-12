package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMembersItemCodespacesItemWithCodespace_nameResponse 
// Deprecated: This class is obsolete. Use WithCodespace_nameDeleteResponse instead.
type ItemMembersItemCodespacesItemWithCodespace_nameResponse struct {
    ItemMembersItemCodespacesItemWithCodespace_nameDeleteResponse
}
// NewItemMembersItemCodespacesItemWithCodespace_nameResponse instantiates a new ItemMembersItemCodespacesItemWithCodespace_nameResponse and sets the default values.
func NewItemMembersItemCodespacesItemWithCodespace_nameResponse()(*ItemMembersItemCodespacesItemWithCodespace_nameResponse) {
    m := &ItemMembersItemCodespacesItemWithCodespace_nameResponse{
        ItemMembersItemCodespacesItemWithCodespace_nameDeleteResponse: *NewItemMembersItemCodespacesItemWithCodespace_nameDeleteResponse(),
    }
    return m
}
// CreateItemMembersItemCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMembersItemCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMembersItemCodespacesItemWithCodespace_nameResponse(), nil
}
// ItemMembersItemCodespacesItemWithCodespace_nameResponseable 
// Deprecated: This class is obsolete. Use WithCodespace_nameDeleteResponse instead.
type ItemMembersItemCodespacesItemWithCodespace_nameResponseable interface {
    ItemMembersItemCodespacesItemWithCodespace_nameDeleteResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
