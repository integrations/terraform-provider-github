package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInstallationsResponse 
// Deprecated: This class is obsolete. Use installationsGetResponse instead.
type ItemInstallationsResponse struct {
    ItemInstallationsGetResponse
}
// NewItemInstallationsResponse instantiates a new ItemInstallationsResponse and sets the default values.
func NewItemInstallationsResponse()(*ItemInstallationsResponse) {
    m := &ItemInstallationsResponse{
        ItemInstallationsGetResponse: *NewItemInstallationsGetResponse(),
    }
    return m
}
// CreateItemInstallationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInstallationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInstallationsResponse(), nil
}
// ItemInstallationsResponseable 
// Deprecated: This class is obsolete. Use installationsGetResponse instead.
type ItemInstallationsResponseable interface {
    ItemInstallationsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
