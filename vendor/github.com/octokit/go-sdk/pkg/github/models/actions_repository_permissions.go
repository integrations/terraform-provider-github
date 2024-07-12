package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsRepositoryPermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions policy that controls the actions and reusable workflows that are allowed to run.
    allowed_actions *AllowedActions
    // Whether GitHub Actions is enabled on the repository.
    enabled *bool
    // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
    selected_actions_url *string
}
// NewActionsRepositoryPermissions instantiates a new ActionsRepositoryPermissions and sets the default values.
func NewActionsRepositoryPermissions()(*ActionsRepositoryPermissions) {
    m := &ActionsRepositoryPermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsRepositoryPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsRepositoryPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsRepositoryPermissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsRepositoryPermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedActions gets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
// returns a *AllowedActions when successful
func (m *ActionsRepositoryPermissions) GetAllowedActions()(*AllowedActions) {
    return m.allowed_actions
}
// GetEnabled gets the enabled property value. Whether GitHub Actions is enabled on the repository.
// returns a *bool when successful
func (m *ActionsRepositoryPermissions) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsRepositoryPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedActions(val.(*AllowedActions))
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["selected_actions_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectedActionsUrl(val)
        }
        return nil
    }
    return res
}
// GetSelectedActionsUrl gets the selected_actions_url property value. The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
// returns a *string when successful
func (m *ActionsRepositoryPermissions) GetSelectedActionsUrl()(*string) {
    return m.selected_actions_url
}
// Serialize serializes information the current object
func (m *ActionsRepositoryPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedActions() != nil {
        cast := (*m.GetAllowedActions()).String()
        err := writer.WriteStringValue("allowed_actions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("selected_actions_url", m.GetSelectedActionsUrl())
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
func (m *ActionsRepositoryPermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedActions sets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
func (m *ActionsRepositoryPermissions) SetAllowedActions(value *AllowedActions)() {
    m.allowed_actions = value
}
// SetEnabled sets the enabled property value. Whether GitHub Actions is enabled on the repository.
func (m *ActionsRepositoryPermissions) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetSelectedActionsUrl sets the selected_actions_url property value. The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
func (m *ActionsRepositoryPermissions) SetSelectedActionsUrl(value *string)() {
    m.selected_actions_url = value
}
type ActionsRepositoryPermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedActions()(*AllowedActions)
    GetEnabled()(*bool)
    GetSelectedActionsUrl()(*string)
    SetAllowedActions(value *AllowedActions)()
    SetEnabled(value *bool)()
    SetSelectedActionsUrl(value *string)()
}
