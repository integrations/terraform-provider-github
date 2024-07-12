package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReviewCustomGatesStateRequired struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional comment to include with the review.
    comment *string
    // The name of the environment to approve or reject.
    environment_name *string
    // Whether to approve or reject deployment to the specified environments.
    state *ReviewCustomGatesStateRequired_state
}
// NewReviewCustomGatesStateRequired instantiates a new ReviewCustomGatesStateRequired and sets the default values.
func NewReviewCustomGatesStateRequired()(*ReviewCustomGatesStateRequired) {
    m := &ReviewCustomGatesStateRequired{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReviewCustomGatesStateRequiredFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReviewCustomGatesStateRequiredFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewCustomGatesStateRequired(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReviewCustomGatesStateRequired) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. Optional comment to include with the review.
// returns a *string when successful
func (m *ReviewCustomGatesStateRequired) GetComment()(*string) {
    return m.comment
}
// GetEnvironmentName gets the environment_name property value. The name of the environment to approve or reject.
// returns a *string when successful
func (m *ReviewCustomGatesStateRequired) GetEnvironmentName()(*string) {
    return m.environment_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReviewCustomGatesStateRequired) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["environment_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironmentName(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReviewCustomGatesStateRequired_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ReviewCustomGatesStateRequired_state))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. Whether to approve or reject deployment to the specified environments.
// returns a *ReviewCustomGatesStateRequired_state when successful
func (m *ReviewCustomGatesStateRequired) GetState()(*ReviewCustomGatesStateRequired_state) {
    return m.state
}
// Serialize serializes information the current object
func (m *ReviewCustomGatesStateRequired) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment_name", m.GetEnvironmentName())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReviewCustomGatesStateRequired) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. Optional comment to include with the review.
func (m *ReviewCustomGatesStateRequired) SetComment(value *string)() {
    m.comment = value
}
// SetEnvironmentName sets the environment_name property value. The name of the environment to approve or reject.
func (m *ReviewCustomGatesStateRequired) SetEnvironmentName(value *string)() {
    m.environment_name = value
}
// SetState sets the state property value. Whether to approve or reject deployment to the specified environments.
func (m *ReviewCustomGatesStateRequired) SetState(value *ReviewCustomGatesStateRequired_state)() {
    m.state = value
}
type ReviewCustomGatesStateRequiredable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetEnvironmentName()(*string)
    GetState()(*ReviewCustomGatesStateRequired_state)
    SetComment(value *string)()
    SetEnvironmentName(value *string)()
    SetState(value *ReviewCustomGatesStateRequired_state)()
}
