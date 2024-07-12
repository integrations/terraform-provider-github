package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPullsItemUpdateBranchResponse 
// Deprecated: This class is obsolete. Use updateBranchPutResponse instead.
type ItemItemPullsItemUpdateBranchResponse struct {
    ItemItemPullsItemUpdateBranchPutResponse
}
// NewItemItemPullsItemUpdateBranchResponse instantiates a new ItemItemPullsItemUpdateBranchResponse and sets the default values.
func NewItemItemPullsItemUpdateBranchResponse()(*ItemItemPullsItemUpdateBranchResponse) {
    m := &ItemItemPullsItemUpdateBranchResponse{
        ItemItemPullsItemUpdateBranchPutResponse: *NewItemItemPullsItemUpdateBranchPutResponse(),
    }
    return m
}
// CreateItemItemPullsItemUpdateBranchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemPullsItemUpdateBranchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemUpdateBranchResponse(), nil
}
// ItemItemPullsItemUpdateBranchResponseable 
// Deprecated: This class is obsolete. Use updateBranchPutResponse instead.
type ItemItemPullsItemUpdateBranchResponseable interface {
    ItemItemPullsItemUpdateBranchPutResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
