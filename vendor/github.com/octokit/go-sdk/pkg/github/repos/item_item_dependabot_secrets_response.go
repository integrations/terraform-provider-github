package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDependabotSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemItemDependabotSecretsResponse struct {
    ItemItemDependabotSecretsGetResponse
}
// NewItemItemDependabotSecretsResponse instantiates a new ItemItemDependabotSecretsResponse and sets the default values.
func NewItemItemDependabotSecretsResponse()(*ItemItemDependabotSecretsResponse) {
    m := &ItemItemDependabotSecretsResponse{
        ItemItemDependabotSecretsGetResponse: *NewItemItemDependabotSecretsGetResponse(),
    }
    return m
}
// CreateItemItemDependabotSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDependabotSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDependabotSecretsResponse(), nil
}
// ItemItemDependabotSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemItemDependabotSecretsResponseable interface {
    ItemItemDependabotSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
