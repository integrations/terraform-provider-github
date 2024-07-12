package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InteractionLimitsGetResponseMember1 
type InteractionLimitsGetResponseMember1 struct {
}
// NewInteractionLimitsGetResponseMember1 instantiates a new InteractionLimitsGetResponseMember1 and sets the default values.
func NewInteractionLimitsGetResponseMember1()(*InteractionLimitsGetResponseMember1) {
    m := &InteractionLimitsGetResponseMember1{
    }
    return m
}
// CreateInteractionLimitsGetResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInteractionLimitsGetResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInteractionLimitsGetResponseMember1(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InteractionLimitsGetResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *InteractionLimitsGetResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// InteractionLimitsGetResponseMember1able 
type InteractionLimitsGetResponseMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
