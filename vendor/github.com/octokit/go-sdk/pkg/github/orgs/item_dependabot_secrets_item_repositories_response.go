package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDependabotSecretsItemRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemDependabotSecretsItemRepositoriesResponse struct {
    ItemDependabotSecretsItemRepositoriesGetResponse
}
// NewItemDependabotSecretsItemRepositoriesResponse instantiates a new ItemDependabotSecretsItemRepositoriesResponse and sets the default values.
func NewItemDependabotSecretsItemRepositoriesResponse()(*ItemDependabotSecretsItemRepositoriesResponse) {
    m := &ItemDependabotSecretsItemRepositoriesResponse{
        ItemDependabotSecretsItemRepositoriesGetResponse: *NewItemDependabotSecretsItemRepositoriesGetResponse(),
    }
    return m
}
// CreateItemDependabotSecretsItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDependabotSecretsItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDependabotSecretsItemRepositoriesResponse(), nil
}
// ItemDependabotSecretsItemRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type ItemDependabotSecretsItemRepositoriesResponseable interface {
    ItemDependabotSecretsItemRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
