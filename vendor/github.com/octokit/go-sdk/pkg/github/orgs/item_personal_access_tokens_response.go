package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPersonalAccessTokensResponse 
// Deprecated: This class is obsolete. Use personalAccessTokensPostResponse instead.
type ItemPersonalAccessTokensResponse struct {
    ItemPersonalAccessTokensPostResponse
}
// NewItemPersonalAccessTokensResponse instantiates a new ItemPersonalAccessTokensResponse and sets the default values.
func NewItemPersonalAccessTokensResponse()(*ItemPersonalAccessTokensResponse) {
    m := &ItemPersonalAccessTokensResponse{
        ItemPersonalAccessTokensPostResponse: *NewItemPersonalAccessTokensPostResponse(),
    }
    return m
}
// CreateItemPersonalAccessTokensResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPersonalAccessTokensResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPersonalAccessTokensResponse(), nil
}
// ItemPersonalAccessTokensResponseable 
// Deprecated: This class is obsolete. Use personalAccessTokensPostResponse instead.
type ItemPersonalAccessTokensResponseable interface {
    ItemPersonalAccessTokensPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
