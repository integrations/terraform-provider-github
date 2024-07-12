package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodespacesSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemCodespacesSecretsResponse struct {
    ItemCodespacesSecretsGetResponse
}
// NewItemCodespacesSecretsResponse instantiates a new ItemCodespacesSecretsResponse and sets the default values.
func NewItemCodespacesSecretsResponse()(*ItemCodespacesSecretsResponse) {
    m := &ItemCodespacesSecretsResponse{
        ItemCodespacesSecretsGetResponse: *NewItemCodespacesSecretsGetResponse(),
    }
    return m
}
// CreateItemCodespacesSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCodespacesSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodespacesSecretsResponse(), nil
}
// ItemCodespacesSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemCodespacesSecretsResponseable interface {
    ItemCodespacesSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
