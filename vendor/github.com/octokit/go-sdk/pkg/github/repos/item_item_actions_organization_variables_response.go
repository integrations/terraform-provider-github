package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsOrganizationVariablesResponse 
// Deprecated: This class is obsolete. Use organizationVariablesGetResponse instead.
type ItemItemActionsOrganizationVariablesResponse struct {
    ItemItemActionsOrganizationVariablesGetResponse
}
// NewItemItemActionsOrganizationVariablesResponse instantiates a new ItemItemActionsOrganizationVariablesResponse and sets the default values.
func NewItemItemActionsOrganizationVariablesResponse()(*ItemItemActionsOrganizationVariablesResponse) {
    m := &ItemItemActionsOrganizationVariablesResponse{
        ItemItemActionsOrganizationVariablesGetResponse: *NewItemItemActionsOrganizationVariablesGetResponse(),
    }
    return m
}
// CreateItemItemActionsOrganizationVariablesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsOrganizationVariablesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsOrganizationVariablesResponse(), nil
}
// ItemItemActionsOrganizationVariablesResponseable 
// Deprecated: This class is obsolete. Use organizationVariablesGetResponse instead.
type ItemItemActionsOrganizationVariablesResponseable interface {
    ItemItemActionsOrganizationVariablesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
