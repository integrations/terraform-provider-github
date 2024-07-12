package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRunnersItemLabelsItemWithNameResponse 
// Deprecated: This class is obsolete. Use WithNameDeleteResponse instead.
type ItemActionsRunnersItemLabelsItemWithNameResponse struct {
    ItemActionsRunnersItemLabelsItemWithNameDeleteResponse
}
// NewItemActionsRunnersItemLabelsItemWithNameResponse instantiates a new ItemActionsRunnersItemLabelsItemWithNameResponse and sets the default values.
func NewItemActionsRunnersItemLabelsItemWithNameResponse()(*ItemActionsRunnersItemLabelsItemWithNameResponse) {
    m := &ItemActionsRunnersItemLabelsItemWithNameResponse{
        ItemActionsRunnersItemLabelsItemWithNameDeleteResponse: *NewItemActionsRunnersItemLabelsItemWithNameDeleteResponse(),
    }
    return m
}
// CreateItemActionsRunnersItemLabelsItemWithNameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersItemLabelsItemWithNameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersItemLabelsItemWithNameResponse(), nil
}
// ItemActionsRunnersItemLabelsItemWithNameResponseable 
// Deprecated: This class is obsolete. Use WithNameDeleteResponse instead.
type ItemActionsRunnersItemLabelsItemWithNameResponseable interface {
    ItemActionsRunnersItemLabelsItemWithNameDeleteResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
