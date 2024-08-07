package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsSetDefaultWorkflowPermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
    can_approve_pull_request_reviews *bool
    // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
    default_workflow_permissions *ActionsDefaultWorkflowPermissions
}
// NewActionsSetDefaultWorkflowPermissions instantiates a new ActionsSetDefaultWorkflowPermissions and sets the default values.
func NewActionsSetDefaultWorkflowPermissions()(*ActionsSetDefaultWorkflowPermissions) {
    m := &ActionsSetDefaultWorkflowPermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsSetDefaultWorkflowPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsSetDefaultWorkflowPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsSetDefaultWorkflowPermissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsSetDefaultWorkflowPermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCanApprovePullRequestReviews gets the can_approve_pull_request_reviews property value. Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
// returns a *bool when successful
func (m *ActionsSetDefaultWorkflowPermissions) GetCanApprovePullRequestReviews()(*bool) {
    return m.can_approve_pull_request_reviews
}
// GetDefaultWorkflowPermissions gets the default_workflow_permissions property value. The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
// returns a *ActionsDefaultWorkflowPermissions when successful
func (m *ActionsSetDefaultWorkflowPermissions) GetDefaultWorkflowPermissions()(*ActionsDefaultWorkflowPermissions) {
    return m.default_workflow_permissions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsSetDefaultWorkflowPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["can_approve_pull_request_reviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanApprovePullRequestReviews(val)
        }
        return nil
    }
    res["default_workflow_permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionsDefaultWorkflowPermissions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultWorkflowPermissions(val.(*ActionsDefaultWorkflowPermissions))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ActionsSetDefaultWorkflowPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("can_approve_pull_request_reviews", m.GetCanApprovePullRequestReviews())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultWorkflowPermissions() != nil {
        cast := (*m.GetDefaultWorkflowPermissions()).String()
        err := writer.WriteStringValue("default_workflow_permissions", &cast)
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
func (m *ActionsSetDefaultWorkflowPermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCanApprovePullRequestReviews sets the can_approve_pull_request_reviews property value. Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
func (m *ActionsSetDefaultWorkflowPermissions) SetCanApprovePullRequestReviews(value *bool)() {
    m.can_approve_pull_request_reviews = value
}
// SetDefaultWorkflowPermissions sets the default_workflow_permissions property value. The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
func (m *ActionsSetDefaultWorkflowPermissions) SetDefaultWorkflowPermissions(value *ActionsDefaultWorkflowPermissions)() {
    m.default_workflow_permissions = value
}
type ActionsSetDefaultWorkflowPermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanApprovePullRequestReviews()(*bool)
    GetDefaultWorkflowPermissions()(*ActionsDefaultWorkflowPermissions)
    SetCanApprovePullRequestReviews(value *bool)()
    SetDefaultWorkflowPermissions(value *ActionsDefaultWorkflowPermissions)()
}
