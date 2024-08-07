package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCommitsItemCheckSuitesResponse 
// Deprecated: This class is obsolete. Use checkSuitesGetResponse instead.
type ItemItemCommitsItemCheckSuitesResponse struct {
    ItemItemCommitsItemCheckSuitesGetResponse
}
// NewItemItemCommitsItemCheckSuitesResponse instantiates a new ItemItemCommitsItemCheckSuitesResponse and sets the default values.
func NewItemItemCommitsItemCheckSuitesResponse()(*ItemItemCommitsItemCheckSuitesResponse) {
    m := &ItemItemCommitsItemCheckSuitesResponse{
        ItemItemCommitsItemCheckSuitesGetResponse: *NewItemItemCommitsItemCheckSuitesGetResponse(),
    }
    return m
}
// CreateItemItemCommitsItemCheckSuitesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCommitsItemCheckSuitesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCommitsItemCheckSuitesResponse(), nil
}
// ItemItemCommitsItemCheckSuitesResponseable 
// Deprecated: This class is obsolete. Use checkSuitesGetResponse instead.
type ItemItemCommitsItemCheckSuitesResponseable interface {
    ItemItemCommitsItemCheckSuitesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
