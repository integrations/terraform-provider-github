package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunnersItemLabelsResponse 
// Deprecated: This class is obsolete. Use labelsGetResponse instead.
type ItemItemActionsRunnersItemLabelsResponse struct {
    ItemItemActionsRunnersItemLabelsGetResponse
}
// NewItemItemActionsRunnersItemLabelsResponse instantiates a new ItemItemActionsRunnersItemLabelsResponse and sets the default values.
func NewItemItemActionsRunnersItemLabelsResponse()(*ItemItemActionsRunnersItemLabelsResponse) {
    m := &ItemItemActionsRunnersItemLabelsResponse{
        ItemItemActionsRunnersItemLabelsGetResponse: *NewItemItemActionsRunnersItemLabelsGetResponse(),
    }
    return m
}
// CreateItemItemActionsRunnersItemLabelsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunnersItemLabelsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunnersItemLabelsResponse(), nil
}
// ItemItemActionsRunnersItemLabelsResponseable 
// Deprecated: This class is obsolete. Use labelsGetResponse instead.
type ItemItemActionsRunnersItemLabelsResponseable interface {
    ItemItemActionsRunnersItemLabelsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
