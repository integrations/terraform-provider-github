package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunnersItemLabelsItemWithNameResponse 
// Deprecated: This class is obsolete. Use WithNameDeleteResponse instead.
type ItemItemActionsRunnersItemLabelsItemWithNameResponse struct {
    ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponse
}
// NewItemItemActionsRunnersItemLabelsItemWithNameResponse instantiates a new ItemItemActionsRunnersItemLabelsItemWithNameResponse and sets the default values.
func NewItemItemActionsRunnersItemLabelsItemWithNameResponse()(*ItemItemActionsRunnersItemLabelsItemWithNameResponse) {
    m := &ItemItemActionsRunnersItemLabelsItemWithNameResponse{
        ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponse: *NewItemItemActionsRunnersItemLabelsItemWithNameDeleteResponse(),
    }
    return m
}
// CreateItemItemActionsRunnersItemLabelsItemWithNameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunnersItemLabelsItemWithNameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunnersItemLabelsItemWithNameResponse(), nil
}
// ItemItemActionsRunnersItemLabelsItemWithNameResponseable 
// Deprecated: This class is obsolete. Use WithNameDeleteResponse instead.
type ItemItemActionsRunnersItemLabelsItemWithNameResponseable interface {
    ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
