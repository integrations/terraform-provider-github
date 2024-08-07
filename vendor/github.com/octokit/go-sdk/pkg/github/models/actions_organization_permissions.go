package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsOrganizationPermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions policy that controls the actions and reusable workflows that are allowed to run.
    allowed_actions *AllowedActions
    // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
    enabled_repositories *EnabledRepositories
    // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
    selected_actions_url *string
    // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
    selected_repositories_url *string
}
// NewActionsOrganizationPermissions instantiates a new ActionsOrganizationPermissions and sets the default values.
func NewActionsOrganizationPermissions()(*ActionsOrganizationPermissions) {
    m := &ActionsOrganizationPermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsOrganizationPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsOrganizationPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsOrganizationPermissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsOrganizationPermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedActions gets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
// returns a *AllowedActions when successful
func (m *ActionsOrganizationPermissions) GetAllowedActions()(*AllowedActions) {
    return m.allowed_actions
}
// GetEnabledRepositories gets the enabled_repositories property value. The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
// returns a *EnabledRepositories when successful
func (m *ActionsOrganizationPermissions) GetEnabledRepositories()(*EnabledRepositories) {
    return m.enabled_repositories
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsOrganizationPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["enabled_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnabledRepositories)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledRepositories(val.(*EnabledRepositories))
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
    res["selected_repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectedRepositoriesUrl(val)
        }
        return nil
    }
    return res
}
// GetSelectedActionsUrl gets the selected_actions_url property value. The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
// returns a *string when successful
func (m *ActionsOrganizationPermissions) GetSelectedActionsUrl()(*string) {
    return m.selected_actions_url
}
// GetSelectedRepositoriesUrl gets the selected_repositories_url property value. The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
// returns a *string when successful
func (m *ActionsOrganizationPermissions) GetSelectedRepositoriesUrl()(*string) {
    return m.selected_repositories_url
}
// Serialize serializes information the current object
func (m *ActionsOrganizationPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedActions() != nil {
        cast := (*m.GetAllowedActions()).String()
        err := writer.WriteStringValue("allowed_actions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnabledRepositories() != nil {
        cast := (*m.GetEnabledRepositories()).String()
        err := writer.WriteStringValue("enabled_repositories", &cast)
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
        err := writer.WriteStringValue("selected_repositories_url", m.GetSelectedRepositoriesUrl())
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
func (m *ActionsOrganizationPermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedActions sets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
func (m *ActionsOrganizationPermissions) SetAllowedActions(value *AllowedActions)() {
    m.allowed_actions = value
}
// SetEnabledRepositories sets the enabled_repositories property value. The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
func (m *ActionsOrganizationPermissions) SetEnabledRepositories(value *EnabledRepositories)() {
    m.enabled_repositories = value
}
// SetSelectedActionsUrl sets the selected_actions_url property value. The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
func (m *ActionsOrganizationPermissions) SetSelectedActionsUrl(value *string)() {
    m.selected_actions_url = value
}
// SetSelectedRepositoriesUrl sets the selected_repositories_url property value. The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
func (m *ActionsOrganizationPermissions) SetSelectedRepositoriesUrl(value *string)() {
    m.selected_repositories_url = value
}
type ActionsOrganizationPermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedActions()(*AllowedActions)
    GetEnabledRepositories()(*EnabledRepositories)
    GetSelectedActionsUrl()(*string)
    GetSelectedRepositoriesUrl()(*string)
    SetAllowedActions(value *AllowedActions)()
    SetEnabledRepositories(value *EnabledRepositories)()
    SetSelectedActionsUrl(value *string)()
    SetSelectedRepositoriesUrl(value *string)()
}
