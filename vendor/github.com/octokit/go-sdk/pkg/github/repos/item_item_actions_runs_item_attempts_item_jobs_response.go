package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunsItemAttemptsItemJobsResponse 
// Deprecated: This class is obsolete. Use jobsGetResponse instead.
type ItemItemActionsRunsItemAttemptsItemJobsResponse struct {
    ItemItemActionsRunsItemAttemptsItemJobsGetResponse
}
// NewItemItemActionsRunsItemAttemptsItemJobsResponse instantiates a new ItemItemActionsRunsItemAttemptsItemJobsResponse and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemJobsResponse()(*ItemItemActionsRunsItemAttemptsItemJobsResponse) {
    m := &ItemItemActionsRunsItemAttemptsItemJobsResponse{
        ItemItemActionsRunsItemAttemptsItemJobsGetResponse: *NewItemItemActionsRunsItemAttemptsItemJobsGetResponse(),
    }
    return m
}
// CreateItemItemActionsRunsItemAttemptsItemJobsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunsItemAttemptsItemJobsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemAttemptsItemJobsResponse(), nil
}
// ItemItemActionsRunsItemAttemptsItemJobsResponseable 
// Deprecated: This class is obsolete. Use jobsGetResponse instead.
type ItemItemActionsRunsItemAttemptsItemJobsResponseable interface {
    ItemItemActionsRunsItemAttemptsItemJobsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
