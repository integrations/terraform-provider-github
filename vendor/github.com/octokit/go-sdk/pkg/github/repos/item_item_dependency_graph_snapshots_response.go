package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDependencyGraphSnapshotsResponse 
// Deprecated: This class is obsolete. Use snapshotsPostResponse instead.
type ItemItemDependencyGraphSnapshotsResponse struct {
    ItemItemDependencyGraphSnapshotsPostResponse
}
// NewItemItemDependencyGraphSnapshotsResponse instantiates a new ItemItemDependencyGraphSnapshotsResponse and sets the default values.
func NewItemItemDependencyGraphSnapshotsResponse()(*ItemItemDependencyGraphSnapshotsResponse) {
    m := &ItemItemDependencyGraphSnapshotsResponse{
        ItemItemDependencyGraphSnapshotsPostResponse: *NewItemItemDependencyGraphSnapshotsPostResponse(),
    }
    return m
}
// CreateItemItemDependencyGraphSnapshotsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDependencyGraphSnapshotsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDependencyGraphSnapshotsResponse(), nil
}
// ItemItemDependencyGraphSnapshotsResponseable 
// Deprecated: This class is obsolete. Use snapshotsPostResponse instead.
type ItemItemDependencyGraphSnapshotsResponseable interface {
    ItemItemDependencyGraphSnapshotsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
