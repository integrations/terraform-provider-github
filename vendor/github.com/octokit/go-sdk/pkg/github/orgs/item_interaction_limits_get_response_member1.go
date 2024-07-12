package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInteractionLimitsGetResponseMember1 
type ItemInteractionLimitsGetResponseMember1 struct {
}
// NewItemInteractionLimitsGetResponseMember1 instantiates a new ItemInteractionLimitsGetResponseMember1 and sets the default values.
func NewItemInteractionLimitsGetResponseMember1()(*ItemInteractionLimitsGetResponseMember1) {
    m := &ItemInteractionLimitsGetResponseMember1{
    }
    return m
}
// CreateItemInteractionLimitsGetResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInteractionLimitsGetResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInteractionLimitsGetResponseMember1(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInteractionLimitsGetResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemInteractionLimitsGetResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// ItemInteractionLimitsGetResponseMember1able 
type ItemInteractionLimitsGetResponseMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
