package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemSecurityAdvisoriesItemCveResponse 
// Deprecated: This class is obsolete. Use cvePostResponse instead.
type ItemItemSecurityAdvisoriesItemCveResponse struct {
    ItemItemSecurityAdvisoriesItemCvePostResponse
}
// NewItemItemSecurityAdvisoriesItemCveResponse instantiates a new ItemItemSecurityAdvisoriesItemCveResponse and sets the default values.
func NewItemItemSecurityAdvisoriesItemCveResponse()(*ItemItemSecurityAdvisoriesItemCveResponse) {
    m := &ItemItemSecurityAdvisoriesItemCveResponse{
        ItemItemSecurityAdvisoriesItemCvePostResponse: *NewItemItemSecurityAdvisoriesItemCvePostResponse(),
    }
    return m
}
// CreateItemItemSecurityAdvisoriesItemCveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemSecurityAdvisoriesItemCveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemSecurityAdvisoriesItemCveResponse(), nil
}
// ItemItemSecurityAdvisoriesItemCveResponseable 
// Deprecated: This class is obsolete. Use cvePostResponse instead.
type ItemItemSecurityAdvisoriesItemCveResponseable interface {
    ItemItemSecurityAdvisoriesItemCvePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
