package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunnersGenerateJitconfigResponse 
// Deprecated: This class is obsolete. Use generateJitconfigPostResponse instead.
type ItemItemActionsRunnersGenerateJitconfigResponse struct {
    ItemItemActionsRunnersGenerateJitconfigPostResponse
}
// NewItemItemActionsRunnersGenerateJitconfigResponse instantiates a new ItemItemActionsRunnersGenerateJitconfigResponse and sets the default values.
func NewItemItemActionsRunnersGenerateJitconfigResponse()(*ItemItemActionsRunnersGenerateJitconfigResponse) {
    m := &ItemItemActionsRunnersGenerateJitconfigResponse{
        ItemItemActionsRunnersGenerateJitconfigPostResponse: *NewItemItemActionsRunnersGenerateJitconfigPostResponse(),
    }
    return m
}
// CreateItemItemActionsRunnersGenerateJitconfigResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunnersGenerateJitconfigResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunnersGenerateJitconfigResponse(), nil
}
// ItemItemActionsRunnersGenerateJitconfigResponseable 
// Deprecated: This class is obsolete. Use generateJitconfigPostResponse instead.
type ItemItemActionsRunnersGenerateJitconfigResponseable interface {
    ItemItemActionsRunnersGenerateJitconfigPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
