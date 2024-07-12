package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsOrganizationSecretsResponse 
// Deprecated: This class is obsolete. Use organizationSecretsGetResponse instead.
type ItemItemActionsOrganizationSecretsResponse struct {
    ItemItemActionsOrganizationSecretsGetResponse
}
// NewItemItemActionsOrganizationSecretsResponse instantiates a new ItemItemActionsOrganizationSecretsResponse and sets the default values.
func NewItemItemActionsOrganizationSecretsResponse()(*ItemItemActionsOrganizationSecretsResponse) {
    m := &ItemItemActionsOrganizationSecretsResponse{
        ItemItemActionsOrganizationSecretsGetResponse: *NewItemItemActionsOrganizationSecretsGetResponse(),
    }
    return m
}
// CreateItemItemActionsOrganizationSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsOrganizationSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsOrganizationSecretsResponse(), nil
}
// ItemItemActionsOrganizationSecretsResponseable 
// Deprecated: This class is obsolete. Use organizationSecretsGetResponse instead.
type ItemItemActionsOrganizationSecretsResponseable interface {
    ItemItemActionsOrganizationSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
