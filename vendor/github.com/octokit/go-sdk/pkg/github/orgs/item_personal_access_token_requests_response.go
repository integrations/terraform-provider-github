package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPersonalAccessTokenRequestsResponse 
// Deprecated: This class is obsolete. Use personalAccessTokenRequestsPostResponse instead.
type ItemPersonalAccessTokenRequestsResponse struct {
    ItemPersonalAccessTokenRequestsPostResponse
}
// NewItemPersonalAccessTokenRequestsResponse instantiates a new ItemPersonalAccessTokenRequestsResponse and sets the default values.
func NewItemPersonalAccessTokenRequestsResponse()(*ItemPersonalAccessTokenRequestsResponse) {
    m := &ItemPersonalAccessTokenRequestsResponse{
        ItemPersonalAccessTokenRequestsPostResponse: *NewItemPersonalAccessTokenRequestsPostResponse(),
    }
    return m
}
// CreateItemPersonalAccessTokenRequestsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPersonalAccessTokenRequestsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPersonalAccessTokenRequestsResponse(), nil
}
// ItemPersonalAccessTokenRequestsResponseable 
// Deprecated: This class is obsolete. Use personalAccessTokenRequestsPostResponse instead.
type ItemPersonalAccessTokenRequestsResponseable interface {
    ItemPersonalAccessTokenRequestsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
