package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse 
// Deprecated: This class is obsolete. Use appsGetResponse instead.
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse struct {
    ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse()(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse{
        ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse: *NewItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse(),
    }
    return m
}
// CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsResponse(), nil
}
// ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseable 
// Deprecated: This class is obsolete. Use appsGetResponse instead.
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsResponseable interface {
    ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
