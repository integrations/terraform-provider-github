package repositories

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEnvironmentsItemSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemEnvironmentsItemSecretsResponse struct {
    ItemEnvironmentsItemSecretsGetResponse
}
// NewItemEnvironmentsItemSecretsResponse instantiates a new ItemEnvironmentsItemSecretsResponse and sets the default values.
func NewItemEnvironmentsItemSecretsResponse()(*ItemEnvironmentsItemSecretsResponse) {
    m := &ItemEnvironmentsItemSecretsResponse{
        ItemEnvironmentsItemSecretsGetResponse: *NewItemEnvironmentsItemSecretsGetResponse(),
    }
    return m
}
// CreateItemEnvironmentsItemSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEnvironmentsItemSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEnvironmentsItemSecretsResponse(), nil
}
// ItemEnvironmentsItemSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type ItemEnvironmentsItemSecretsResponseable interface {
    ItemEnvironmentsItemSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
