package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodespacesSecretsItemRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemCodespacesSecretsItemRepositoriesResponse struct {
    ItemCodespacesSecretsItemRepositoriesGetResponse
}
// NewItemCodespacesSecretsItemRepositoriesResponse instantiates a new ItemCodespacesSecretsItemRepositoriesResponse and sets the default values.
func NewItemCodespacesSecretsItemRepositoriesResponse()(*ItemCodespacesSecretsItemRepositoriesResponse) {
    m := &ItemCodespacesSecretsItemRepositoriesResponse{
        ItemCodespacesSecretsItemRepositoriesGetResponse: *NewItemCodespacesSecretsItemRepositoriesGetResponse(),
    }
    return m
}
// CreateItemCodespacesSecretsItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCodespacesSecretsItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodespacesSecretsItemRepositoriesResponse(), nil
}
// ItemCodespacesSecretsItemRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemCodespacesSecretsItemRepositoriesResponseable interface {
    ItemCodespacesSecretsItemRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
