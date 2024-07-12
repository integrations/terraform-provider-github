package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemActionsPermissionsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions policy that controls the actions and reusable workflows that are allowed to run.
    allowed_actions *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AllowedActions
    // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
    enabled_repositories *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EnabledRepositories
}
// NewItemActionsPermissionsPutRequestBody instantiates a new ItemActionsPermissionsPutRequestBody and sets the default values.
func NewItemActionsPermissionsPutRequestBody()(*ItemActionsPermissionsPutRequestBody) {
    m := &ItemActionsPermissionsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsPermissionsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsPermissionsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsPermissionsPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsPermissionsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedActions gets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
// returns a *AllowedActions when successful
func (m *ItemActionsPermissionsPutRequestBody) GetAllowedActions()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AllowedActions) {
    return m.allowed_actions
}
// GetEnabledRepositories gets the enabled_repositories property value. The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
// returns a *EnabledRepositories when successful
func (m *ItemActionsPermissionsPutRequestBody) GetEnabledRepositories()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EnabledRepositories) {
    return m.enabled_repositories
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsPermissionsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseAllowedActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedActions(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AllowedActions))
        }
        return nil
    }
    res["enabled_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseEnabledRepositories)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledRepositories(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EnabledRepositories))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemActionsPermissionsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsPermissionsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedActions sets the allowed_actions property value. The permissions policy that controls the actions and reusable workflows that are allowed to run.
func (m *ItemActionsPermissionsPutRequestBody) SetAllowedActions(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AllowedActions)() {
    m.allowed_actions = value
}
// SetEnabledRepositories sets the enabled_repositories property value. The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
func (m *ItemActionsPermissionsPutRequestBody) SetEnabledRepositories(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EnabledRepositories)() {
    m.enabled_repositories = value
}
type ItemActionsPermissionsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedActions()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AllowedActions)
    GetEnabledRepositories()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EnabledRepositories)
    SetAllowedActions(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AllowedActions)()
    SetEnabledRepositories(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.EnabledRepositories)()
}
