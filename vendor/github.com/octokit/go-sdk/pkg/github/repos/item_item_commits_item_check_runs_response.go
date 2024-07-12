package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCommitsItemCheckRunsResponse 
// Deprecated: This class is obsolete. Use checkRunsGetResponse instead.
type ItemItemCommitsItemCheckRunsResponse struct {
    ItemItemCommitsItemCheckRunsGetResponse
}
// NewItemItemCommitsItemCheckRunsResponse instantiates a new ItemItemCommitsItemCheckRunsResponse and sets the default values.
func NewItemItemCommitsItemCheckRunsResponse()(*ItemItemCommitsItemCheckRunsResponse) {
    m := &ItemItemCommitsItemCheckRunsResponse{
        ItemItemCommitsItemCheckRunsGetResponse: *NewItemItemCommitsItemCheckRunsGetResponse(),
    }
    return m
}
// CreateItemItemCommitsItemCheckRunsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCommitsItemCheckRunsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCommitsItemCheckRunsResponse(), nil
}
// ItemItemCommitsItemCheckRunsResponseable 
// Deprecated: This class is obsolete. Use checkRunsGetResponse instead.
type ItemItemCommitsItemCheckRunsResponseable interface {
    ItemItemCommitsItemCheckRunsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
