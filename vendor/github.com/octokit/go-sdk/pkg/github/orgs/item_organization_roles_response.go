package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOrganizationRolesResponse 
// Deprecated: This class is obsolete. Use organizationRolesGetResponse instead.
type ItemOrganizationRolesResponse struct {
    ItemOrganizationRolesGetResponse
}
// NewItemOrganizationRolesResponse instantiates a new ItemOrganizationRolesResponse and sets the default values.
func NewItemOrganizationRolesResponse()(*ItemOrganizationRolesResponse) {
    m := &ItemOrganizationRolesResponse{
        ItemOrganizationRolesGetResponse: *NewItemOrganizationRolesGetResponse(),
    }
    return m
}
// CreateItemOrganizationRolesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOrganizationRolesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOrganizationRolesResponse(), nil
}
// ItemOrganizationRolesResponseable 
// Deprecated: This class is obsolete. Use organizationRolesGetResponse instead.
type ItemOrganizationRolesResponseable interface {
    ItemOrganizationRolesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
