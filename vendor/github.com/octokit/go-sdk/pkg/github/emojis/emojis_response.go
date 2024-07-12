package emojis

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmojisResponse 
// Deprecated: This class is obsolete. Use emojisGetResponse instead.
type EmojisResponse struct {
    EmojisGetResponse
}
// NewEmojisResponse instantiates a new emojisResponse and sets the default values.
func NewEmojisResponse()(*EmojisResponse) {
    m := &EmojisResponse{
        EmojisGetResponse: *NewEmojisGetResponse(),
    }
    return m
}
// CreateEmojisResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmojisResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmojisResponse(), nil
}
// EmojisResponseable 
// Deprecated: This class is obsolete. Use emojisGetResponse instead.
type EmojisResponseable interface {
    EmojisGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
