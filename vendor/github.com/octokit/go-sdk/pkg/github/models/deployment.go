package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deployment a request for a specific ref(branch,sha,tag) to be deployed
type Deployment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    creator NullableSimpleUserable
    // The description property
    description *string
    // Name for the target deployment environment.
    environment *string
    // Unique identifier of the deployment
    id *int64
    // The node_id property
    node_id *string
    // The original_environment property
    original_environment *string
    // The payload property
    payload *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // Specifies if the given environment is one that end-users directly interact with. Default: false.
    production_environment *bool
    // The ref to deploy. This can be a branch, tag, or sha.
    ref *string
    // The repository_url property
    repository_url *string
    // The sha property
    sha *string
    // The statuses_url property
    statuses_url *string
    // Parameter to specify a task to execute
    task *string
    // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
    transient_environment *bool
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewDeployment instantiates a new Deployment and sets the default values.
func NewDeployment()(*Deployment) {
    m := &Deployment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeployment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Deployment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Deployment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Deployment) GetCreator()(NullableSimpleUserable) {
    return m.creator
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Deployment) GetDescription()(*string) {
    return m.description
}
// GetEnvironment gets the environment property value. Name for the target deployment environment.
// returns a *string when successful
func (m *Deployment) GetEnvironment()(*string) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Deployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["original_environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEnvironment(val)
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
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
    res["production_environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductionEnvironment(val)
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
    res["statuses_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusesUrl(val)
        }
        return nil
    }
    res["task"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTask(val)
        }
        return nil
    }
    res["transient_environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransientEnvironment(val)
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
// GetId gets the id property value. Unique identifier of the deployment
// returns a *int64 when successful
func (m *Deployment) GetId()(*int64) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Deployment) GetNodeId()(*string) {
    return m.node_id
}
// GetOriginalEnvironment gets the original_environment property value. The original_environment property
// returns a *string when successful
func (m *Deployment) GetOriginalEnvironment()(*string) {
    return m.original_environment
}
// GetPayload gets the payload property value. The payload property
// returns a *string when successful
func (m *Deployment) GetPayload()(*string) {
    return m.payload
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *Deployment) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetProductionEnvironment gets the production_environment property value. Specifies if the given environment is one that end-users directly interact with. Default: false.
// returns a *bool when successful
func (m *Deployment) GetProductionEnvironment()(*bool) {
    return m.production_environment
}
// GetRef gets the ref property value. The ref to deploy. This can be a branch, tag, or sha.
// returns a *string when successful
func (m *Deployment) GetRef()(*string) {
    return m.ref
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *Deployment) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *Deployment) GetSha()(*string) {
    return m.sha
}
// GetStatusesUrl gets the statuses_url property value. The statuses_url property
// returns a *string when successful
func (m *Deployment) GetStatusesUrl()(*string) {
    return m.statuses_url
}
// GetTask gets the task property value. Parameter to specify a task to execute
// returns a *string when successful
func (m *Deployment) GetTask()(*string) {
    return m.task
}
// GetTransientEnvironment gets the transient_environment property value. Specifies if the given environment is will no longer exist at some point in the future. Default: false.
// returns a *bool when successful
func (m *Deployment) GetTransientEnvironment()(*bool) {
    return m.transient_environment
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Deployment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Deployment) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Deployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteStringValue("original_environment", m.GetOriginalEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payload", m.GetPayload())
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
        err := writer.WriteBoolValue("production_environment", m.GetProductionEnvironment())
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
        err := writer.WriteStringValue("repository_url", m.GetRepositoryUrl())
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
        err := writer.WriteStringValue("statuses_url", m.GetStatusesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("task", m.GetTask())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("transient_environment", m.GetTransientEnvironment())
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
func (m *Deployment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Deployment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *Deployment) SetCreator(value NullableSimpleUserable)() {
    m.creator = value
}
// SetDescription sets the description property value. The description property
func (m *Deployment) SetDescription(value *string)() {
    m.description = value
}
// SetEnvironment sets the environment property value. Name for the target deployment environment.
func (m *Deployment) SetEnvironment(value *string)() {
    m.environment = value
}
// SetId sets the id property value. Unique identifier of the deployment
func (m *Deployment) SetId(value *int64)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Deployment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOriginalEnvironment sets the original_environment property value. The original_environment property
func (m *Deployment) SetOriginalEnvironment(value *string)() {
    m.original_environment = value
}
// SetPayload sets the payload property value. The payload property
func (m *Deployment) SetPayload(value *string)() {
    m.payload = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *Deployment) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetProductionEnvironment sets the production_environment property value. Specifies if the given environment is one that end-users directly interact with. Default: false.
func (m *Deployment) SetProductionEnvironment(value *bool)() {
    m.production_environment = value
}
// SetRef sets the ref property value. The ref to deploy. This can be a branch, tag, or sha.
func (m *Deployment) SetRef(value *string)() {
    m.ref = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *Deployment) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetSha sets the sha property value. The sha property
func (m *Deployment) SetSha(value *string)() {
    m.sha = value
}
// SetStatusesUrl sets the statuses_url property value. The statuses_url property
func (m *Deployment) SetStatusesUrl(value *string)() {
    m.statuses_url = value
}
// SetTask sets the task property value. Parameter to specify a task to execute
func (m *Deployment) SetTask(value *string)() {
    m.task = value
}
// SetTransientEnvironment sets the transient_environment property value. Specifies if the given environment is will no longer exist at some point in the future. Default: false.
func (m *Deployment) SetTransientEnvironment(value *bool)() {
    m.transient_environment = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Deployment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Deployment) SetUrl(value *string)() {
    m.url = value
}
type Deploymentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetId()(*int64)
    GetNodeId()(*string)
    GetOriginalEnvironment()(*string)
    GetPayload()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetProductionEnvironment()(*bool)
    GetRef()(*string)
    GetRepositoryUrl()(*string)
    GetSha()(*string)
    GetStatusesUrl()(*string)
    GetTask()(*string)
    GetTransientEnvironment()(*bool)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetId(value *int64)()
    SetNodeId(value *string)()
    SetOriginalEnvironment(value *string)()
    SetPayload(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetProductionEnvironment(value *bool)()
    SetRef(value *string)()
    SetRepositoryUrl(value *string)()
    SetSha(value *string)()
    SetStatusesUrl(value *string)()
    SetTask(value *string)()
    SetTransientEnvironment(value *bool)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
