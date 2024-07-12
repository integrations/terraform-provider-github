package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsVariablesResponse 
// Deprecated: This class is obsolete. Use variablesGetResponse instead.
type ItemActionsVariablesResponse struct {
    ItemActionsVariablesGetResponse
}
// NewItemActionsVariablesResponse instantiates a new ItemActionsVariablesResponse and sets the default values.
func NewItemActionsVariablesResponse()(*ItemActionsVariablesResponse) {
    m := &ItemActionsVariablesResponse{
        ItemActionsVariablesGetResponse: *NewItemActionsVariablesGetResponse(),
    }
    return m
}
// CreateItemActionsVariablesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsVariablesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsVariablesResponse(), nil
}
// ItemActionsVariablesResponseable 
// Deprecated: This class is obsolete. Use variablesGetResponse instead.
type ItemActionsVariablesResponseable interface {
    ItemActionsVariablesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
