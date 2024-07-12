package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOutside_collaboratorsItemWithUsernameResponse 
// Deprecated: This class is obsolete. Use WithUsernamePutResponse instead.
type ItemOutside_collaboratorsItemWithUsernameResponse struct {
    ItemOutside_collaboratorsItemWithUsernamePutResponse
}
// NewItemOutside_collaboratorsItemWithUsernameResponse instantiates a new ItemOutside_collaboratorsItemWithUsernameResponse and sets the default values.
func NewItemOutside_collaboratorsItemWithUsernameResponse()(*ItemOutside_collaboratorsItemWithUsernameResponse) {
    m := &ItemOutside_collaboratorsItemWithUsernameResponse{
        ItemOutside_collaboratorsItemWithUsernamePutResponse: *NewItemOutside_collaboratorsItemWithUsernamePutResponse(),
    }
    return m
}
// CreateItemOutside_collaboratorsItemWithUsernameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOutside_collaboratorsItemWithUsernameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutside_collaboratorsItemWithUsernameResponse(), nil
}
// ItemOutside_collaboratorsItemWithUsernameResponseable 
// Deprecated: This class is obsolete. Use WithUsernamePutResponse instead.
type ItemOutside_collaboratorsItemWithUsernameResponseable interface {
    ItemOutside_collaboratorsItemWithUsernamePutResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
