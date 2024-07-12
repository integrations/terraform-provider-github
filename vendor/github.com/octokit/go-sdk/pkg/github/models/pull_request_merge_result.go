package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestMergeResult pull Request Merge Result
type PullRequestMergeResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The merged property
    merged *bool
    // The message property
    message *string
    // The sha property
    sha *string
}
// NewPullRequestMergeResult instantiates a new PullRequestMergeResult and sets the default values.
func NewPullRequestMergeResult()(*PullRequestMergeResult) {
    m := &PullRequestMergeResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestMergeResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestMergeResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestMergeResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestMergeResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestMergeResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["merged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMerged(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    return res
}
// GetMerged gets the merged property value. The merged property
// returns a *bool when successful
func (m *PullRequestMergeResult) GetMerged()(*bool) {
    return m.merged
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *PullRequestMergeResult) GetMessage()(*string) {
    return m.message
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *PullRequestMergeResult) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *PullRequestMergeResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("merged", m.GetMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
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
func (m *PullRequestMergeResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMerged sets the merged property value. The merged property
func (m *PullRequestMergeResult) SetMerged(value *bool)() {
    m.merged = value
}
// SetMessage sets the message property value. The message property
func (m *PullRequestMergeResult) SetMessage(value *string)() {
    m.message = value
}
// SetSha sets the sha property value. The sha property
func (m *PullRequestMergeResult) SetSha(value *string)() {
    m.sha = value
}
type PullRequestMergeResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMerged()(*bool)
    GetMessage()(*string)
    GetSha()(*string)
    SetMerged(value *bool)()
    SetMessage(value *string)()
    SetSha(value *string)()
}
