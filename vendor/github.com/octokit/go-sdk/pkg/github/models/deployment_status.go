package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentStatus the status of a deployment.
type DeploymentStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    creator NullableSimpleUserable
    // The deployment_url property
    deployment_url *string
    // A short description of the status.
    description *string
    // The environment of the deployment that the status is for.
    environment *string
    // The URL for accessing your environment.
    environment_url *string
    // The id property
    id *int64
    // The URL to associate with this status.
    log_url *string
    // The node_id property
    node_id *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // The repository_url property
    repository_url *string
    // The state of the status.
    state *DeploymentStatus_state
    // Deprecated: the URL to associate with this status.
    target_url *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewDeploymentStatus instantiates a new DeploymentStatus and sets the default values.
func NewDeploymentStatus()(*DeploymentStatus) {
    m := &DeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeploymentStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *DeploymentStatus) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *DeploymentStatus) GetCreator()(NullableSimpleUserable) {
    return m.creator
}
// GetDeploymentUrl gets the deployment_url property value. The deployment_url property
// returns a *string when successful
func (m *DeploymentStatus) GetDeploymentUrl()(*string) {
    return m.deployment_url
}
// GetDescription gets the description property value. A short description of the status.
// returns a *string when successful
func (m *DeploymentStatus) GetDescription()(*string) {
    return m.description
}
// GetEnvironment gets the environment property value. The environment of the deployment that the status is for.
// returns a *string when successful
func (m *DeploymentStatus) GetEnvironment()(*string) {
    return m.environment
}
// GetEnvironmentUrl gets the environment_url property value. The URL for accessing your environment.
// returns a *string when successful
func (m *DeploymentStatus) GetEnvironmentUrl()(*string) {
    return m.environment_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["creator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["deployment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentUrl(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val)
        }
        return nil
    }
    res["environment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironmentUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["log_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogUrl(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["performed_via_github_app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedViaGithubApp(val.(NullableIntegrationable))
        }
        return nil
    }
    res["repository_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryUrl(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStatus_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*DeploymentStatus_state))
        }
        return nil
    }
    res["target_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetUrl(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *DeploymentStatus) GetId()(*int64) {
    return m.id
}
// GetLogUrl gets the log_url property value. The URL to associate with this status.
// returns a *string when successful
func (m *DeploymentStatus) GetLogUrl()(*string) {
    return m.log_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *DeploymentStatus) GetNodeId()(*string) {
    return m.node_id
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *DeploymentStatus) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *DeploymentStatus) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetState gets the state property value. The state of the status.
// returns a *DeploymentStatus_state when successful
func (m *DeploymentStatus) GetState()(*DeploymentStatus_state) {
    return m.state
}
// GetTargetUrl gets the target_url property value. Deprecated: the URL to associate with this status.
// returns a *string when successful
func (m *DeploymentStatus) GetTargetUrl()(*string) {
    return m.target_url
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *DeploymentStatus) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *DeploymentStatus) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *DeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("creator", m.GetCreator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployment_url", m.GetDeploymentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment_url", m.GetEnvironmentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("log_url", m.GetLogUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_url", m.GetRepositoryUrl())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_url", m.GetTargetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *DeploymentStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *DeploymentStatus) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *DeploymentStatus) SetCreator(value NullableSimpleUserable)() {
    m.creator = value
}
// SetDeploymentUrl sets the deployment_url property value. The deployment_url property
func (m *DeploymentStatus) SetDeploymentUrl(value *string)() {
    m.deployment_url = value
}
// SetDescription sets the description property value. A short description of the status.
func (m *DeploymentStatus) SetDescription(value *string)() {
    m.description = value
}
// SetEnvironment sets the environment property value. The environment of the deployment that the status is for.
func (m *DeploymentStatus) SetEnvironment(value *string)() {
    m.environment = value
}
// SetEnvironmentUrl sets the environment_url property value. The URL for accessing your environment.
func (m *DeploymentStatus) SetEnvironmentUrl(value *string)() {
    m.environment_url = value
}
// SetId sets the id property value. The id property
func (m *DeploymentStatus) SetId(value *int64)() {
    m.id = value
}
// SetLogUrl sets the log_url property value. The URL to associate with this status.
func (m *DeploymentStatus) SetLogUrl(value *string)() {
    m.log_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *DeploymentStatus) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *DeploymentStatus) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *DeploymentStatus) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetState sets the state property value. The state of the status.
func (m *DeploymentStatus) SetState(value *DeploymentStatus_state)() {
    m.state = value
}
// SetTargetUrl sets the target_url property value. Deprecated: the URL to associate with this status.
func (m *DeploymentStatus) SetTargetUrl(value *string)() {
    m.target_url = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *DeploymentStatus) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *DeploymentStatus) SetUrl(value *string)() {
    m.url = value
}
type DeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetDeploymentUrl()(*string)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetEnvironmentUrl()(*string)
    GetId()(*int64)
    GetLogUrl()(*string)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetRepositoryUrl()(*string)
    GetState()(*DeploymentStatus_state)
    GetTargetUrl()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetDeploymentUrl(value *string)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetEnvironmentUrl(value *string)()
    SetId(value *int64)()
    SetLogUrl(value *string)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetRepositoryUrl(value *string)()
    SetState(value *DeploymentStatus_state)()
    SetTargetUrl(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
