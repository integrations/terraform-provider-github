package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsWorkflowAccessToRepository struct {
    // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within therepository.`none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
    access_level *ActionsWorkflowAccessToRepository_access_level
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewActionsWorkflowAccessToRepository instantiates a new ActionsWorkflowAccessToRepository and sets the default values.
func NewActionsWorkflowAccessToRepository()(*ActionsWorkflowAccessToRepository) {
    m := &ActionsWorkflowAccessToRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsWorkflowAccessToRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsWorkflowAccessToRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsWorkflowAccessToRepository(), nil
}
// GetAccessLevel gets the access_level property value. Defines the level of access that workflows outside of the repository have to actions and reusable workflows within therepository.`none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
// returns a *ActionsWorkflowAccessToRepository_access_level when successful
func (m *ActionsWorkflowAccessToRepository) GetAccessLevel()(*ActionsWorkflowAccessToRepository_access_level) {
    return m.access_level
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsWorkflowAccessToRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsWorkflowAccessToRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionsWorkflowAccessToRepository_access_level)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*ActionsWorkflowAccessToRepository_access_level))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ActionsWorkflowAccessToRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
        err := writer.WriteStringValue("access_level", &cast)
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
// SetAccessLevel sets the access_level property value. Defines the level of access that workflows outside of the repository have to actions and reusable workflows within therepository.`none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
func (m *ActionsWorkflowAccessToRepository) SetAccessLevel(value *ActionsWorkflowAccessToRepository_access_level)() {
    m.access_level = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionsWorkflowAccessToRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ActionsWorkflowAccessToRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessLevel()(*ActionsWorkflowAccessToRepository_access_level)
    SetAccessLevel(value *ActionsWorkflowAccessToRepository_access_level)()
}
