package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Import_project_choices struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The human_name property
    human_name *string
    // The tfvc_project property
    tfvc_project *string
    // The vcs property
    vcs *string
}
// NewImport_project_choices instantiates a new Import_project_choices and sets the default values.
func NewImport_project_choices()(*Import_project_choices) {
    m := &Import_project_choices{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImport_project_choicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImport_project_choicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImport_project_choices(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Import_project_choices) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Import_project_choices) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["human_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHumanName(val)
        }
        return nil
    }
    res["tfvc_project"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTfvcProject(val)
        }
        return nil
    }
    res["vcs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVcs(val)
        }
        return nil
    }
    return res
}
// GetHumanName gets the human_name property value. The human_name property
// returns a *string when successful
func (m *Import_project_choices) GetHumanName()(*string) {
    return m.human_name
}
// GetTfvcProject gets the tfvc_project property value. The tfvc_project property
// returns a *string when successful
func (m *Import_project_choices) GetTfvcProject()(*string) {
    return m.tfvc_project
}
// GetVcs gets the vcs property value. The vcs property
// returns a *string when successful
func (m *Import_project_choices) GetVcs()(*string) {
    return m.vcs
}
// Serialize serializes information the current object
func (m *Import_project_choices) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("human_name", m.GetHumanName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tfvc_project", m.GetTfvcProject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vcs", m.GetVcs())
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
func (m *Import_project_choices) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHumanName sets the human_name property value. The human_name property
func (m *Import_project_choices) SetHumanName(value *string)() {
    m.human_name = value
}
// SetTfvcProject sets the tfvc_project property value. The tfvc_project property
func (m *Import_project_choices) SetTfvcProject(value *string)() {
    m.tfvc_project = value
}
// SetVcs sets the vcs property value. The vcs property
func (m *Import_project_choices) SetVcs(value *string)() {
    m.vcs = value
}
type Import_project_choicesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHumanName()(*string)
    GetTfvcProject()(*string)
    GetVcs()(*string)
    SetHumanName(value *string)()
    SetTfvcProject(value *string)()
    SetVcs(value *string)()
}
