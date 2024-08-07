package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunsItemJobsResponse 
// Deprecated: This class is obsolete. Use jobsGetResponse instead.
type ItemItemActionsRunsItemJobsResponse struct {
    ItemItemActionsRunsItemJobsGetResponse
}
// NewItemItemActionsRunsItemJobsResponse instantiates a new ItemItemActionsRunsItemJobsResponse and sets the default values.
func NewItemItemActionsRunsItemJobsResponse()(*ItemItemActionsRunsItemJobsResponse) {
    m := &ItemItemActionsRunsItemJobsResponse{
        ItemItemActionsRunsItemJobsGetResponse: *NewItemItemActionsRunsItemJobsGetResponse(),
    }
    return m
}
// CreateItemItemActionsRunsItemJobsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunsItemJobsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemJobsResponse(), nil
}
// ItemItemActionsRunsItemJobsResponseable 
// Deprecated: This class is obsolete. Use jobsGetResponse instead.
type ItemItemActionsRunsItemJobsResponseable interface {
    ItemItemActionsRunsItemJobsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
