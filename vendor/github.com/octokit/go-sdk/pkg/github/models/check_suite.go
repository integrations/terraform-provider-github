package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckSuite a suite of checks performed on the code of a given code change
type CheckSuite struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The after property
    after *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    app NullableIntegrationable
    // The before property
    before *string
    // The check_runs_url property
    check_runs_url *string
    // The conclusion property
    conclusion *CheckSuite_conclusion
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The head_branch property
    head_branch *string
    // A commit.
    head_commit SimpleCommitable
    // The SHA of the head commit that is being checked.
    head_sha *string
    // The id property
    id *int32
    // The latest_check_runs_count property
    latest_check_runs_count *int32
    // The node_id property
    node_id *string
    // The pull_requests property
    pull_requests []PullRequestMinimalable
    // Minimal Repository
    repository MinimalRepositoryable
    // The rerequestable property
    rerequestable *bool
    // The runs_rerequestable property
    runs_rerequestable *bool
    // The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
    status *CheckSuite_status
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewCheckSuite instantiates a new CheckSuite and sets the default values.
func NewCheckSuite()(*CheckSuite) {
    m := &CheckSuite{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckSuiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckSuiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckSuite(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckSuite) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfter gets the after property value. The after property
// returns a *string when successful
func (m *CheckSuite) GetAfter()(*string) {
    return m.after
}
// GetApp gets the app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *CheckSuite) GetApp()(NullableIntegrationable) {
    return m.app
}
// GetBefore gets the before property value. The before property
// returns a *string when successful
func (m *CheckSuite) GetBefore()(*string) {
    return m.before
}
// GetCheckRunsUrl gets the check_runs_url property value. The check_runs_url property
// returns a *string when successful
func (m *CheckSuite) GetCheckRunsUrl()(*string) {
    return m.check_runs_url
}
// GetConclusion gets the conclusion property value. The conclusion property
// returns a *CheckSuite_conclusion when successful
func (m *CheckSuite) GetConclusion()(*CheckSuite_conclusion) {
    return m.conclusion
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *CheckSuite) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckSuite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["after"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfter(val)
        }
        return nil
    }
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
    res["before"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBefore(val)
        }
        return nil
    }
    res["check_runs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckRunsUrl(val)
        }
        return nil
    }
    res["conclusion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCheckSuite_conclusion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConclusion(val.(*CheckSuite_conclusion))
        }
        return nil
    }
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
    res["head_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadBranch(val)
        }
        return nil
    }
    res["head_commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadCommit(val.(SimpleCommitable))
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
    res["latest_check_runs_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestCheckRunsCount(val)
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
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["rerequestable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRerequestable(val)
        }
        return nil
    }
    res["runs_rerequestable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunsRerequestable(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCheckSuite_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CheckSuite_status))
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
// GetHeadBranch gets the head_branch property value. The head_branch property
// returns a *string when successful
func (m *CheckSuite) GetHeadBranch()(*string) {
    return m.head_branch
}
// GetHeadCommit gets the head_commit property value. A commit.
// returns a SimpleCommitable when successful
func (m *CheckSuite) GetHeadCommit()(SimpleCommitable) {
    return m.head_commit
}
// GetHeadSha gets the head_sha property value. The SHA of the head commit that is being checked.
// returns a *string when successful
func (m *CheckSuite) GetHeadSha()(*string) {
    return m.head_sha
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *CheckSuite) GetId()(*int32) {
    return m.id
}
// GetLatestCheckRunsCount gets the latest_check_runs_count property value. The latest_check_runs_count property
// returns a *int32 when successful
func (m *CheckSuite) GetLatestCheckRunsCount()(*int32) {
    return m.latest_check_runs_count
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *CheckSuite) GetNodeId()(*string) {
    return m.node_id
}
// GetPullRequests gets the pull_requests property value. The pull_requests property
// returns a []PullRequestMinimalable when successful
func (m *CheckSuite) GetPullRequests()([]PullRequestMinimalable) {
    return m.pull_requests
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *CheckSuite) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetRerequestable gets the rerequestable property value. The rerequestable property
// returns a *bool when successful
func (m *CheckSuite) GetRerequestable()(*bool) {
    return m.rerequestable
}
// GetRunsRerequestable gets the runs_rerequestable property value. The runs_rerequestable property
// returns a *bool when successful
func (m *CheckSuite) GetRunsRerequestable()(*bool) {
    return m.runs_rerequestable
}
// GetStatus gets the status property value. The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
// returns a *CheckSuite_status when successful
func (m *CheckSuite) GetStatus()(*CheckSuite_status) {
    return m.status
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *CheckSuite) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CheckSuite) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CheckSuite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("after", m.GetAfter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("before", m.GetBefore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("check_runs_url", m.GetCheckRunsUrl())
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
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("head_branch", m.GetHeadBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("head_commit", m.GetHeadCommit())
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
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("latest_check_runs_count", m.GetLatestCheckRunsCount())
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
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("rerequestable", m.GetRerequestable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("runs_rerequestable", m.GetRunsRerequestable())
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
func (m *CheckSuite) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfter sets the after property value. The after property
func (m *CheckSuite) SetAfter(value *string)() {
    m.after = value
}
// SetApp sets the app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *CheckSuite) SetApp(value NullableIntegrationable)() {
    m.app = value
}
// SetBefore sets the before property value. The before property
func (m *CheckSuite) SetBefore(value *string)() {
    m.before = value
}
// SetCheckRunsUrl sets the check_runs_url property value. The check_runs_url property
func (m *CheckSuite) SetCheckRunsUrl(value *string)() {
    m.check_runs_url = value
}
// SetConclusion sets the conclusion property value. The conclusion property
func (m *CheckSuite) SetConclusion(value *CheckSuite_conclusion)() {
    m.conclusion = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *CheckSuite) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHeadBranch sets the head_branch property value. The head_branch property
func (m *CheckSuite) SetHeadBranch(value *string)() {
    m.head_branch = value
}
// SetHeadCommit sets the head_commit property value. A commit.
func (m *CheckSuite) SetHeadCommit(value SimpleCommitable)() {
    m.head_commit = value
}
// SetHeadSha sets the head_sha property value. The SHA of the head commit that is being checked.
func (m *CheckSuite) SetHeadSha(value *string)() {
    m.head_sha = value
}
// SetId sets the id property value. The id property
func (m *CheckSuite) SetId(value *int32)() {
    m.id = value
}
// SetLatestCheckRunsCount sets the latest_check_runs_count property value. The latest_check_runs_count property
func (m *CheckSuite) SetLatestCheckRunsCount(value *int32)() {
    m.latest_check_runs_count = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *CheckSuite) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPullRequests sets the pull_requests property value. The pull_requests property
func (m *CheckSuite) SetPullRequests(value []PullRequestMinimalable)() {
    m.pull_requests = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *CheckSuite) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetRerequestable sets the rerequestable property value. The rerequestable property
func (m *CheckSuite) SetRerequestable(value *bool)() {
    m.rerequestable = value
}
// SetRunsRerequestable sets the runs_rerequestable property value. The runs_rerequestable property
func (m *CheckSuite) SetRunsRerequestable(value *bool)() {
    m.runs_rerequestable = value
}
// SetStatus sets the status property value. The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
func (m *CheckSuite) SetStatus(value *CheckSuite_status)() {
    m.status = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *CheckSuite) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *CheckSuite) SetUrl(value *string)() {
    m.url = value
}
type CheckSuiteable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAfter()(*string)
    GetApp()(NullableIntegrationable)
    GetBefore()(*string)
    GetCheckRunsUrl()(*string)
    GetConclusion()(*CheckSuite_conclusion)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHeadBranch()(*string)
    GetHeadCommit()(SimpleCommitable)
    GetHeadSha()(*string)
    GetId()(*int32)
    GetLatestCheckRunsCount()(*int32)
    GetNodeId()(*string)
    GetPullRequests()([]PullRequestMinimalable)
    GetRepository()(MinimalRepositoryable)
    GetRerequestable()(*bool)
    GetRunsRerequestable()(*bool)
    GetStatus()(*CheckSuite_status)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAfter(value *string)()
    SetApp(value NullableIntegrationable)()
    SetBefore(value *string)()
    SetCheckRunsUrl(value *string)()
    SetConclusion(value *CheckSuite_conclusion)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHeadBranch(value *string)()
    SetHeadCommit(value SimpleCommitable)()
    SetHeadSha(value *string)()
    SetId(value *int32)()
    SetLatestCheckRunsCount(value *int32)()
    SetNodeId(value *string)()
    SetPullRequests(value []PullRequestMinimalable)()
    SetRepository(value MinimalRepositoryable)()
    SetRerequestable(value *bool)()
    SetRunsRerequestable(value *bool)()
    SetStatus(value *CheckSuite_status)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
