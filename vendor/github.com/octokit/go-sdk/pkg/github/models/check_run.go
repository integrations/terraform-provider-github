package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckRun a check performed on the code of a given code change
type CheckRun struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    app NullableIntegrationable
    // The check_suite property
    check_suite CheckRun_check_suiteable
    // The completed_at property
    completed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The conclusion property
    conclusion *CheckRun_conclusion
    // A deployment created as the result of an Actions check run from a workflow that references an environment
    deployment DeploymentSimpleable
    // The details_url property
    details_url *string
    // The external_id property
    external_id *string
    // The SHA of the commit that is being checked.
    head_sha *string
    // The html_url property
    html_url *string
    // The id of the check.
    id *int32
    // The name of the check.
    name *string
    // The node_id property
    node_id *string
    // The output property
    output CheckRun_outputable
    // Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
    pull_requests []PullRequestMinimalable
    // The started_at property
    started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
    status *CheckRun_status
    // The url property
    url *string
}
// NewCheckRun instantiates a new CheckRun and sets the default values.
func NewCheckRun()(*CheckRun) {
    m := &CheckRun{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckRun(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckRun) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApp gets the app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *CheckRun) GetApp()(NullableIntegrationable) {
    return m.app
}
// GetCheckSuite gets the check_suite property value. The check_suite property
// returns a CheckRun_check_suiteable when successful
func (m *CheckRun) GetCheckSuite()(CheckRun_check_suiteable) {
    return m.check_suite
}
// GetCompletedAt gets the completed_at property value. The completed_at property
// returns a *Time when successful
func (m *CheckRun) GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completed_at
}
// GetConclusion gets the conclusion property value. The conclusion property
// returns a *CheckRun_conclusion when successful
func (m *CheckRun) GetConclusion()(*CheckRun_conclusion) {
    return m.conclusion
}
// GetDeployment gets the deployment property value. A deployment created as the result of an Actions check run from a workflow that references an environment
// returns a DeploymentSimpleable when successful
func (m *CheckRun) GetDeployment()(DeploymentSimpleable) {
    return m.deployment
}
// GetDetailsUrl gets the details_url property value. The details_url property
// returns a *string when successful
func (m *CheckRun) GetDetailsUrl()(*string) {
    return m.details_url
}
// GetExternalId gets the external_id property value. The external_id property
// returns a *string when successful
func (m *CheckRun) GetExternalId()(*string) {
    return m.external_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckRun) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(NullableIntegrationable))
        }
        return nil
    }
    res["check_suite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCheckRun_check_suiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckSuite(val.(CheckRun_check_suiteable))
        }
        return nil
    }
    res["completed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedAt(val)
        }
        return nil
    }
    res["conclusion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCheckRun_conclusion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConclusion(val.(*CheckRun_conclusion))
        }
        return nil
    }
    res["deployment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployment(val.(DeploymentSimpleable))
        }
        return nil
    }
    res["details_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetailsUrl(val)
        }
        return nil
    }
    res["external_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["head_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadSha(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["output"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCheckRun_outputFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutput(val.(CheckRun_outputable))
        }
        return nil
    }
    res["pull_requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePullRequestMinimalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullRequestMinimalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PullRequestMinimalable)
                }
            }
            m.SetPullRequests(res)
        }
        return nil
    }
    res["started_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedAt(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCheckRun_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CheckRun_status))
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
// GetHeadSha gets the head_sha property value. The SHA of the commit that is being checked.
// returns a *string when successful
func (m *CheckRun) GetHeadSha()(*string) {
    return m.head_sha
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *CheckRun) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id of the check.
// returns a *int32 when successful
func (m *CheckRun) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the check.
// returns a *string when successful
func (m *CheckRun) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *CheckRun) GetNodeId()(*string) {
    return m.node_id
}
// GetOutput gets the output property value. The output property
// returns a CheckRun_outputable when successful
func (m *CheckRun) GetOutput()(CheckRun_outputable) {
    return m.output
}
// GetPullRequests gets the pull_requests property value. Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
// returns a []PullRequestMinimalable when successful
func (m *CheckRun) GetPullRequests()([]PullRequestMinimalable) {
    return m.pull_requests
}
// GetStartedAt gets the started_at property value. The started_at property
// returns a *Time when successful
func (m *CheckRun) GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.started_at
}
// GetStatus gets the status property value. The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
// returns a *CheckRun_status when successful
func (m *CheckRun) GetStatus()(*CheckRun_status) {
    return m.status
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CheckRun) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CheckRun) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("check_suite", m.GetCheckSuite())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completed_at", m.GetCompletedAt())
        if err != nil {
            return err
        }
    }
    if m.GetConclusion() != nil {
        cast := (*m.GetConclusion()).String()
        err := writer.WriteStringValue("conclusion", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deployment", m.GetDeployment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("details_url", m.GetDetailsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("external_id", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("head_sha", m.GetHeadSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteObjectValue("output", m.GetOutput())
        if err != nil {
            return err
        }
    }
    if m.GetPullRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPullRequests()))
        for i, v := range m.GetPullRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("pull_requests", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("started_at", m.GetStartedAt())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CheckRun) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApp sets the app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *CheckRun) SetApp(value NullableIntegrationable)() {
    m.app = value
}
// SetCheckSuite sets the check_suite property value. The check_suite property
func (m *CheckRun) SetCheckSuite(value CheckRun_check_suiteable)() {
    m.check_suite = value
}
// SetCompletedAt sets the completed_at property value. The completed_at property
func (m *CheckRun) SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completed_at = value
}
// SetConclusion sets the conclusion property value. The conclusion property
func (m *CheckRun) SetConclusion(value *CheckRun_conclusion)() {
    m.conclusion = value
}
// SetDeployment sets the deployment property value. A deployment created as the result of an Actions check run from a workflow that references an environment
func (m *CheckRun) SetDeployment(value DeploymentSimpleable)() {
    m.deployment = value
}
// SetDetailsUrl sets the details_url property value. The details_url property
func (m *CheckRun) SetDetailsUrl(value *string)() {
    m.details_url = value
}
// SetExternalId sets the external_id property value. The external_id property
func (m *CheckRun) SetExternalId(value *string)() {
    m.external_id = value
}
// SetHeadSha sets the head_sha property value. The SHA of the commit that is being checked.
func (m *CheckRun) SetHeadSha(value *string)() {
    m.head_sha = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *CheckRun) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id of the check.
func (m *CheckRun) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the check.
func (m *CheckRun) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *CheckRun) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOutput sets the output property value. The output property
func (m *CheckRun) SetOutput(value CheckRun_outputable)() {
    m.output = value
}
// SetPullRequests sets the pull_requests property value. Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
func (m *CheckRun) SetPullRequests(value []PullRequestMinimalable)() {
    m.pull_requests = value
}
// SetStartedAt sets the started_at property value. The started_at property
func (m *CheckRun) SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.started_at = value
}
// SetStatus sets the status property value. The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
func (m *CheckRun) SetStatus(value *CheckRun_status)() {
    m.status = value
}
// SetUrl sets the url property value. The url property
func (m *CheckRun) SetUrl(value *string)() {
    m.url = value
}
type CheckRunable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(NullableIntegrationable)
    GetCheckSuite()(CheckRun_check_suiteable)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConclusion()(*CheckRun_conclusion)
    GetDeployment()(DeploymentSimpleable)
    GetDetailsUrl()(*string)
    GetExternalId()(*string)
    GetHeadSha()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOutput()(CheckRun_outputable)
    GetPullRequests()([]PullRequestMinimalable)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*CheckRun_status)
    GetUrl()(*string)
    SetApp(value NullableIntegrationable)()
    SetCheckSuite(value CheckRun_check_suiteable)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConclusion(value *CheckRun_conclusion)()
    SetDeployment(value DeploymentSimpleable)()
    SetDetailsUrl(value *string)()
    SetExternalId(value *string)()
    SetHeadSha(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOutput(value CheckRun_outputable)()
    SetPullRequests(value []PullRequestMinimalable)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *CheckRun_status)()
    SetUrl(value *string)()
}
