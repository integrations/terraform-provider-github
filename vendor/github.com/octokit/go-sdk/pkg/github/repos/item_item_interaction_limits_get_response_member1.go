package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemInteractionLimitsGetResponseMember1 
type ItemItemInteractionLimitsGetResponseMember1 struct {
}
// NewItemItemInteractionLimitsGetResponseMember1 instantiates a new ItemItemInteractionLimitsGetResponseMember1 and sets the default values.
func NewItemItemInteractionLimitsGetResponseMember1()(*ItemItemInteractionLimitsGetResponseMember1) {
    m := &ItemItemInteractionLimitsGetResponseMember1{
    }
    return m
}
// CreateItemItemInteractionLimitsGetResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemInteractionLimitsGetResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemInteractionLimitsGetResponseMember1(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemInteractionLimitsGetResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemItemInteractionLimitsGetResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// ItemItemInteractionLimitsGetResponseMember1able 
type ItemItemInteractionLimitsGetResponseMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
