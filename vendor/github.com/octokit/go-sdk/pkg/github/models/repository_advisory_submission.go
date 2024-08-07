package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryAdvisory_submission struct {
    // Whether a private vulnerability report was accepted by the repository's administrators.
    accepted *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewRepositoryAdvisory_submission instantiates a new RepositoryAdvisory_submission and sets the default values.
func NewRepositoryAdvisory_submission()(*RepositoryAdvisory_submission) {
    m := &RepositoryAdvisory_submission{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryAdvisory_submissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisory_submissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisory_submission(), nil
}
// GetAccepted gets the accepted property value. Whether a private vulnerability report was accepted by the repository's administrators.
// returns a *bool when successful
func (m *RepositoryAdvisory_submission) GetAccepted()(*bool) {
    return m.accepted
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryAdvisory_submission) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisory_submission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccepted(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RepositoryAdvisory_submission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccepted sets the accepted property value. Whether a private vulnerability report was accepted by the repository's administrators.
func (m *RepositoryAdvisory_submission) SetAccepted(value *bool)() {
    m.accepted = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryAdvisory_submission) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type RepositoryAdvisory_submissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccepted()(*bool)
    SetAccepted(value *bool)()
}
