package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReviewCustomGatesCommentRequired struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Comment associated with the pending deployment protection rule. **Required when state is not provided.**
    comment *string
    // The name of the environment to approve or reject.
    environment_name *string
}
// NewReviewCustomGatesCommentRequired instantiates a new ReviewCustomGatesCommentRequired and sets the default values.
func NewReviewCustomGatesCommentRequired()(*ReviewCustomGatesCommentRequired) {
    m := &ReviewCustomGatesCommentRequired{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewCustomGatesCommentRequired(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReviewCustomGatesCommentRequired) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. Comment associated with the pending deployment protection rule. **Required when state is not provided.**
// returns a *string when successful
func (m *ReviewCustomGatesCommentRequired) GetComment()(*string) {
    return m.comment
}
// GetEnvironmentName gets the environment_name property value. The name of the environment to approve or reject.
// returns a *string when successful
func (m *ReviewCustomGatesCommentRequired) GetEnvironmentName()(*string) {
    return m.environment_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReviewCustomGatesCommentRequired) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *ReviewCustomGatesCommentRequired) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReviewCustomGatesCommentRequired) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. Comment associated with the pending deployment protection rule. **Required when state is not provided.**
func (m *ReviewCustomGatesCommentRequired) SetComment(value *string)() {
    m.comment = value
}
// SetEnvironmentName sets the environment_name property value. The name of the environment to approve or reject.
func (m *ReviewCustomGatesCommentRequired) SetEnvironmentName(value *string)() {
    m.environment_name = value
}
type ReviewCustomGatesCommentRequiredable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetEnvironmentName()(*string)
    SetComment(value *string)()
    SetEnvironmentName(value *string)()
}
