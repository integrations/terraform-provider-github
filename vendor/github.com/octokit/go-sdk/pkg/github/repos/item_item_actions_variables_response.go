package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsVariablesResponse 
// Deprecated: This class is obsolete. Use variablesGetResponse instead.
type ItemItemActionsVariablesResponse struct {
    ItemItemActionsVariablesGetResponse
}
// NewItemItemActionsVariablesResponse instantiates a new ItemItemActionsVariablesResponse and sets the default values.
func NewItemItemActionsVariablesResponse()(*ItemItemActionsVariablesResponse) {
    m := &ItemItemActionsVariablesResponse{
        ItemItemActionsVariablesGetResponse: *NewItemItemActionsVariablesGetResponse(),
    }
    return m
}
// CreateItemItemActionsVariablesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsVariablesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsVariablesResponse(), nil
}
// ItemItemActionsVariablesResponseable 
// Deprecated: This class is obsolete. Use variablesGetResponse instead.
type ItemItemActionsVariablesResponseable interface {
    ItemItemActionsVariablesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
