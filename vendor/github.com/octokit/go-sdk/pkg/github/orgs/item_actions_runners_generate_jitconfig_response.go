package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRunnersGenerateJitconfigResponse 
// Deprecated: This class is obsolete. Use generateJitconfigPostResponse instead.
type ItemActionsRunnersGenerateJitconfigResponse struct {
    ItemActionsRunnersGenerateJitconfigPostResponse
}
// NewItemActionsRunnersGenerateJitconfigResponse instantiates a new ItemActionsRunnersGenerateJitconfigResponse and sets the default values.
func NewItemActionsRunnersGenerateJitconfigResponse()(*ItemActionsRunnersGenerateJitconfigResponse) {
    m := &ItemActionsRunnersGenerateJitconfigResponse{
        ItemActionsRunnersGenerateJitconfigPostResponse: *NewItemActionsRunnersGenerateJitconfigPostResponse(),
    }
    return m
}
// CreateItemActionsRunnersGenerateJitconfigResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRunnersGenerateJitconfigResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersGenerateJitconfigResponse(), nil
}
// ItemActionsRunnersGenerateJitconfigResponseable 
// Deprecated: This class is obsolete. Use generateJitconfigPostResponse instead.
type ItemActionsRunnersGenerateJitconfigResponseable interface {
    ItemActionsRunnersGenerateJitconfigPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
