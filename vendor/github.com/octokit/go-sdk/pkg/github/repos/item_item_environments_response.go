package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemEnvironmentsResponse 
// Deprecated: This class is obsolete. Use environmentsGetResponse instead.
type ItemItemEnvironmentsResponse struct {
    ItemItemEnvironmentsGetResponse
}
// NewItemItemEnvironmentsResponse instantiates a new ItemItemEnvironmentsResponse and sets the default values.
func NewItemItemEnvironmentsResponse()(*ItemItemEnvironmentsResponse) {
    m := &ItemItemEnvironmentsResponse{
        ItemItemEnvironmentsGetResponse: *NewItemItemEnvironmentsGetResponse(),
    }
    return m
}
// CreateItemItemEnvironmentsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsResponse(), nil
}
// ItemItemEnvironmentsResponseable 
// Deprecated: This class is obsolete. Use environmentsGetResponse instead.
type ItemItemEnvironmentsResponseable interface {
    ItemItemEnvironmentsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
