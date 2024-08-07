package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsWorkflowsResponse 
// Deprecated: This class is obsolete. Use workflowsGetResponse instead.
type ItemItemActionsWorkflowsResponse struct {
    ItemItemActionsWorkflowsGetResponse
}
// NewItemItemActionsWorkflowsResponse instantiates a new ItemItemActionsWorkflowsResponse and sets the default values.
func NewItemItemActionsWorkflowsResponse()(*ItemItemActionsWorkflowsResponse) {
    m := &ItemItemActionsWorkflowsResponse{
        ItemItemActionsWorkflowsGetResponse: *NewItemItemActionsWorkflowsGetResponse(),
    }
    return m
}
// CreateItemItemActionsWorkflowsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsWorkflowsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsWorkflowsResponse(), nil
}
// ItemItemActionsWorkflowsResponseable 
// Deprecated: This class is obsolete. Use workflowsGetResponse instead.
type ItemItemActionsWorkflowsResponseable interface {
    ItemItemActionsWorkflowsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
