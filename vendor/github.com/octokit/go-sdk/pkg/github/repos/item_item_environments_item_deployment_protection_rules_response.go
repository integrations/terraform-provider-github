package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemEnvironmentsItemDeployment_protection_rulesResponse 
// Deprecated: This class is obsolete. Use deployment_protection_rulesGetResponse instead.
type ItemItemEnvironmentsItemDeployment_protection_rulesResponse struct {
    ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesResponse instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesResponse and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesResponse()(*ItemItemEnvironmentsItemDeployment_protection_rulesResponse) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesResponse{
        ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse: *NewItemItemEnvironmentsItemDeployment_protection_rulesGetResponse(),
    }
    return m
}
// CreateItemItemEnvironmentsItemDeployment_protection_rulesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsItemDeployment_protection_rulesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesResponse(), nil
}
// ItemItemEnvironmentsItemDeployment_protection_rulesResponseable 
// Deprecated: This class is obsolete. Use deployment_protection_rulesGetResponse instead.
type ItemItemEnvironmentsItemDeployment_protection_rulesResponseable interface {
    ItemItemEnvironmentsItemDeployment_protection_rulesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
