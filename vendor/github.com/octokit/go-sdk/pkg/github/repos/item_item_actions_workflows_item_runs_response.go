package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsWorkflowsItemRunsResponse 
// Deprecated: This class is obsolete. Use runsGetResponse instead.
type ItemItemActionsWorkflowsItemRunsResponse struct {
    ItemItemActionsWorkflowsItemRunsGetResponse
}
// NewItemItemActionsWorkflowsItemRunsResponse instantiates a new ItemItemActionsWorkflowsItemRunsResponse and sets the default values.
func NewItemItemActionsWorkflowsItemRunsResponse()(*ItemItemActionsWorkflowsItemRunsResponse) {
    m := &ItemItemActionsWorkflowsItemRunsResponse{
        ItemItemActionsWorkflowsItemRunsGetResponse: *NewItemItemActionsWorkflowsItemRunsGetResponse(),
    }
    return m
}
// CreateItemItemActionsWorkflowsItemRunsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsWorkflowsItemRunsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsWorkflowsItemRunsResponse(), nil
}
// ItemItemActionsWorkflowsItemRunsResponseable 
// Deprecated: This class is obsolete. Use runsGetResponse instead.
type ItemItemActionsWorkflowsItemRunsResponseable interface {
    ItemItemActionsWorkflowsItemRunsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
