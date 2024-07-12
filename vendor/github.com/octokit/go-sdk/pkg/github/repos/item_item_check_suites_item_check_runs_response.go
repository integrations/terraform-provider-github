package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCheckSuitesItemCheckRunsResponse 
// Deprecated: This class is obsolete. Use checkRunsGetResponse instead.
type ItemItemCheckSuitesItemCheckRunsResponse struct {
    ItemItemCheckSuitesItemCheckRunsGetResponse
}
// NewItemItemCheckSuitesItemCheckRunsResponse instantiates a new ItemItemCheckSuitesItemCheckRunsResponse and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsResponse()(*ItemItemCheckSuitesItemCheckRunsResponse) {
    m := &ItemItemCheckSuitesItemCheckRunsResponse{
        ItemItemCheckSuitesItemCheckRunsGetResponse: *NewItemItemCheckSuitesItemCheckRunsGetResponse(),
    }
    return m
}
// CreateItemItemCheckSuitesItemCheckRunsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCheckSuitesItemCheckRunsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCheckSuitesItemCheckRunsResponse(), nil
}
// ItemItemCheckSuitesItemCheckRunsResponseable 
// Deprecated: This class is obsolete. Use checkRunsGetResponse instead.
type ItemItemCheckSuitesItemCheckRunsResponseable interface {
    ItemItemCheckSuitesItemCheckRunsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
