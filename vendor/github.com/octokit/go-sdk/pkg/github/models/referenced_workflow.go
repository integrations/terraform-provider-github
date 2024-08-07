package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferencedWorkflow a workflow referenced/reused by the initial caller workflow
type ReferencedWorkflow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The path property
    path *string
    // The ref property
    ref *string
    // The sha property
    sha *string
}
// NewReferencedWorkflow instantiates a new ReferencedWorkflow and sets the default values.
func NewReferencedWorkflow()(*ReferencedWorkflow) {
    m := &ReferencedWorkflow{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReferencedWorkflowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReferencedWorkflowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferencedWorkflow(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReferencedWorkflow) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReferencedWorkflow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
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
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ReferencedWorkflow) GetPath()(*string) {
    return m.path
}
// GetRef gets the ref property value. The ref property
// returns a *string when successful
func (m *ReferencedWorkflow) GetRef()(*string) {
    return m.ref
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *ReferencedWorkflow) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *ReferencedWorkflow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
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
func (m *ReferencedWorkflow) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPath sets the path property value. The path property
func (m *ReferencedWorkflow) SetPath(value *string)() {
    m.path = value
}
// SetRef sets the ref property value. The ref property
func (m *ReferencedWorkflow) SetRef(value *string)() {
    m.ref = value
}
// SetSha sets the sha property value. The sha property
func (m *ReferencedWorkflow) SetSha(value *string)() {
    m.sha = value
}
type ReferencedWorkflowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPath()(*string)
    GetRef()(*string)
    GetSha()(*string)
    SetPath(value *string)()
    SetRef(value *string)()
    SetSha(value *string)()
}
