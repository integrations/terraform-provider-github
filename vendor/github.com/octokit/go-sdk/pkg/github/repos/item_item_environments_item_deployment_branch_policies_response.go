package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse 
// Deprecated: This class is obsolete. Use deploymentBranchPoliciesGetResponse instead.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse struct {
    ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesResponse instantiates a new ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesResponse()(*ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) {
    m := &ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse{
        ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse: *NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse(),
    }
    return m
}
// CreateItemItemEnvironmentsItemDeploymentBranchPoliciesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsItemDeploymentBranchPoliciesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesResponse(), nil
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesResponseable 
// Deprecated: This class is obsolete. Use deploymentBranchPoliciesGetResponse instead.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesResponseable interface {
    ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
