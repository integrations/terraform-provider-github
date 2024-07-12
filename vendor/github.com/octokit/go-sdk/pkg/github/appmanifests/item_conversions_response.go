package appmanifests

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemConversionsResponse 
// Deprecated: This class is obsolete. Use conversionsPostResponse instead.
type ItemConversionsResponse struct {
    ItemConversionsPostResponse
}
// NewItemConversionsResponse instantiates a new ItemConversionsResponse and sets the default values.
func NewItemConversionsResponse()(*ItemConversionsResponse) {
    m := &ItemConversionsResponse{
        ItemConversionsPostResponse: *NewItemConversionsPostResponse(),
    }
    return m
}
// CreateItemConversionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemConversionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversionsResponse(), nil
}
// ItemConversionsResponseable 
// Deprecated: This class is obsolete. Use conversionsPostResponse instead.
type ItemConversionsResponseable interface {
    ItemConversionsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
