package repositories

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEnvironmentsItemVariablesResponse 
// Deprecated: This class is obsolete. Use variablesGetResponse instead.
type ItemEnvironmentsItemVariablesResponse struct {
    ItemEnvironmentsItemVariablesGetResponse
}
// NewItemEnvironmentsItemVariablesResponse instantiates a new ItemEnvironmentsItemVariablesResponse and sets the default values.
func NewItemEnvironmentsItemVariablesResponse()(*ItemEnvironmentsItemVariablesResponse) {
    m := &ItemEnvironmentsItemVariablesResponse{
        ItemEnvironmentsItemVariablesGetResponse: *NewItemEnvironmentsItemVariablesGetResponse(),
    }
    return m
}
// CreateItemEnvironmentsItemVariablesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEnvironmentsItemVariablesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEnvironmentsItemVariablesResponse(), nil
}
// ItemEnvironmentsItemVariablesResponseable 
// Deprecated: This class is obsolete. Use variablesGetResponse instead.
type ItemEnvironmentsItemVariablesResponseable interface {
    ItemEnvironmentsItemVariablesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
