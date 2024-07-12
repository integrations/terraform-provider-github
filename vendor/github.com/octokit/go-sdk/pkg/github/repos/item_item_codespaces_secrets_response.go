package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemItemCodespacesSecretsResponse struct {
    ItemItemCodespacesSecretsGetResponse
}
// NewItemItemCodespacesSecretsResponse instantiates a new ItemItemCodespacesSecretsResponse and sets the default values.
func NewItemItemCodespacesSecretsResponse()(*ItemItemCodespacesSecretsResponse) {
    m := &ItemItemCodespacesSecretsResponse{
        ItemItemCodespacesSecretsGetResponse: *NewItemItemCodespacesSecretsGetResponse(),
    }
    return m
}
// CreateItemItemCodespacesSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesSecretsResponse(), nil
}
// ItemItemCodespacesSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemItemCodespacesSecretsResponseable interface {
    ItemItemCodespacesSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
