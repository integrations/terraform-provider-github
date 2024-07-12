package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRunnersItemLabelsResponse 
// Deprecated: This class is obsolete. Use labelsGetResponse instead.
type ItemActionsRunnersItemLabelsResponse struct {
    ItemActionsRunnersItemLabelsGetResponse
}
// NewItemActionsRunnersItemLabelsResponse instantiates a new ItemActionsRunnersItemLabelsResponse and sets the default values.
func NewItemActionsRunnersItemLabelsResponse()(*ItemActionsRunnersItemLabelsResponse) {
    m := &ItemActionsRunnersItemLabelsResponse{
        ItemActionsRunnersItemLabelsGetResponse: *NewItemActionsRunnersItemLabelsGetResponse(),
    }
    return m
}
// CreateItemActionsRunnersItemLabelsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersItemLabelsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersItemLabelsResponse(), nil
}
// ItemActionsRunnersItemLabelsResponseable 
// Deprecated: This class is obsolete. Use labelsGetResponse instead.
type ItemActionsRunnersItemLabelsResponseable interface {
    ItemActionsRunnersItemLabelsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
