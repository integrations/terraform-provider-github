package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDependabotSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemDependabotSecretsResponse struct {
    ItemDependabotSecretsGetResponse
}
// NewItemDependabotSecretsResponse instantiates a new ItemDependabotSecretsResponse and sets the default values.
func NewItemDependabotSecretsResponse()(*ItemDependabotSecretsResponse) {
    m := &ItemDependabotSecretsResponse{
        ItemDependabotSecretsGetResponse: *NewItemDependabotSecretsGetResponse(),
    }
    return m
}
// CreateItemDependabotSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDependabotSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDependabotSecretsResponse(), nil
}
// ItemDependabotSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemDependabotSecretsResponseable interface {
    ItemDependabotSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
