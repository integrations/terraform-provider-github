package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsCacheUsageByRepositoryResponse 
// Deprecated: This class is obsolete. Use usageByRepositoryGetResponse instead.
type ItemActionsCacheUsageByRepositoryResponse struct {
    ItemActionsCacheUsageByRepositoryGetResponse
}
// NewItemActionsCacheUsageByRepositoryResponse instantiates a new ItemActionsCacheUsageByRepositoryResponse and sets the default values.
func NewItemActionsCacheUsageByRepositoryResponse()(*ItemActionsCacheUsageByRepositoryResponse) {
    m := &ItemActionsCacheUsageByRepositoryResponse{
        ItemActionsCacheUsageByRepositoryGetResponse: *NewItemActionsCacheUsageByRepositoryGetResponse(),
    }
    return m
}
// CreateItemActionsCacheUsageByRepositoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsCacheUsageByRepositoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsCacheUsageByRepositoryResponse(), nil
}
// ItemActionsCacheUsageByRepositoryResponseable 
// Deprecated: This class is obsolete. Use usageByRepositoryGetResponse instead.
type ItemActionsCacheUsageByRepositoryResponseable interface {
    ItemActionsCacheUsageByRepositoryGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
