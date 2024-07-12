package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemOutside_collaboratorsItemWithUsernamePutResponse struct {
}
// NewItemOutside_collaboratorsItemWithUsernamePutResponse instantiates a new ItemOutside_collaboratorsItemWithUsernamePutResponse and sets the default values.
func NewItemOutside_collaboratorsItemWithUsernamePutResponse()(*ItemOutside_collaboratorsItemWithUsernamePutResponse) {
    m := &ItemOutside_collaboratorsItemWithUsernamePutResponse{
    }
    return m
}
// CreateItemOutside_collaboratorsItemWithUsernamePutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOutside_collaboratorsItemWithUsernamePutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutside_collaboratorsItemWithUsernamePutResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemOutside_collaboratorsItemWithUsernamePutResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemOutside_collaboratorsItemWithUsernamePutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
type ItemOutside_collaboratorsItemWithUsernamePutResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
