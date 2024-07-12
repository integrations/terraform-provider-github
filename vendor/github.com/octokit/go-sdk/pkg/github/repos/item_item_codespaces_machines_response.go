package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesMachinesResponse 
// Deprecated: This class is obsolete. Use machinesGetResponse instead.
type ItemItemCodespacesMachinesResponse struct {
    ItemItemCodespacesMachinesGetResponse
}
// NewItemItemCodespacesMachinesResponse instantiates a new ItemItemCodespacesMachinesResponse and sets the default values.
func NewItemItemCodespacesMachinesResponse()(*ItemItemCodespacesMachinesResponse) {
    m := &ItemItemCodespacesMachinesResponse{
        ItemItemCodespacesMachinesGetResponse: *NewItemItemCodespacesMachinesGetResponse(),
    }
    return m
}
// CreateItemItemCodespacesMachinesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesMachinesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesMachinesResponse(), nil
}
// ItemItemCodespacesMachinesResponseable 
// Deprecated: This class is obsolete. Use machinesGetResponse instead.
type ItemItemCodespacesMachinesResponseable interface {
    ItemItemCodespacesMachinesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
