package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemWithOrgResponse 
// Deprecated: This class is obsolete. Use WithOrgDeleteResponse instead.
type ItemWithOrgResponse struct {
    ItemWithOrgDeleteResponse
}
// NewItemWithOrgResponse instantiates a new ItemWithOrgResponse and sets the default values.
func NewItemWithOrgResponse()(*ItemWithOrgResponse) {
    m := &ItemWithOrgResponse{
        ItemWithOrgDeleteResponse: *NewItemWithOrgDeleteResponse(),
    }
    return m
}
// CreateItemWithOrgResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemWithOrgResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithOrgResponse(), nil
}
// ItemWithOrgResponseable 
// Deprecated: This class is obsolete. Use WithOrgDeleteResponse instead.
type ItemWithOrgResponseable interface {
    ItemWithOrgDeleteResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
