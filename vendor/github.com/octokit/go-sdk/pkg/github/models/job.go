package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Job information of a job execution in a workflow run
type Job struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The check_run_url property
    check_run_url *string
    // The time that the job finished, in ISO 8601 format.
    completed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The outcome of the job.
    conclusion *Job_conclusion
    // The time that the job created, in ISO 8601 format.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the current branch.
    head_branch *string
    // The SHA of the commit that is being run.
    head_sha *string
    // The html_url property
    html_url *string
    // The id of the job.
    id *int32
    // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
    labels []string
    // The name of the job.
    name *string
    // The node_id property
    node_id *string
    // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
    run_attempt *int32
    // The id of the associated workflow run.
    run_id *int32
    // The run_url property
    run_url *string
    // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
    runner_group_id *int32
    // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
    runner_group_name *string
    // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
    runner_id *int32
    // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
    runner_name *string
    // The time that the job started, in ISO 8601 format.
    started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The phase of the lifecycle that the job is currently in.
    status *Job_status
    // Steps in this job.
    steps []Job_stepsable
    // The url property
    url *string
    // The name of the workflow.
    workflow_name *string
}
// NewJob instantiates a new Job and sets the default values.
func NewJob()(*Job) {
    m := &Job{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJob(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Job) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckRunUrl gets the check_run_url property value. The check_run_url property
// returns a *string when successful
func (m *Job) GetCheckRunUrl()(*string) {
    return m.check_run_url
}
// GetCompletedAt gets the completed_at property value. The time that the job finished, in ISO 8601 format.
// returns a *Time when successful
func (m *Job) GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completed_at
}
// GetConclusion gets the conclusion property value. The outcome of the job.
// returns a *Job_conclusion when successful
func (m *Job) GetConclusion()(*Job_conclusion) {
    return m.conclusion
}
// GetCreatedAt gets the created_at property value. The time that the job created, in ISO 8601 format.
// returns a *Time when successful
func (m *Job) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Job) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_run_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckRunUrl(val)
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
        val, err := n.GetEnumValue(ParseJob_conclusion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConclusion(val.(*Job_conclusion))
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
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetLabels(res)
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
    res["run_attempt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAttempt(val)
        }
        return nil
    }
    res["run_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
        }
        return nil
    }
    res["run_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunUrl(val)
        }
        return nil
    }
    res["runner_group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerGroupId(val)
        }
        return nil
    }
    res["runner_group_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerGroupName(val)
        }
        return nil
    }
    res["runner_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerId(val)
        }
        return nil
    }
    res["runner_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerName(val)
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
        val, err := n.GetEnumValue(ParseJob_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Job_status))
        }
        return nil
    }
    res["steps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJob_stepsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Job_stepsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Job_stepsable)
                }
            }
            m.SetSteps(res)
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
    res["workflow_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowName(val)
        }
        return nil
    }
    return res
}
// GetHeadBranch gets the head_branch property value. The name of the current branch.
// returns a *string when successful
func (m *Job) GetHeadBranch()(*string) {
    return m.head_branch
}
// GetHeadSha gets the head_sha property value. The SHA of the commit that is being run.
// returns a *string when successful
func (m *Job) GetHeadSha()(*string) {
    return m.head_sha
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Job) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id of the job.
// returns a *int32 when successful
func (m *Job) GetId()(*int32) {
    return m.id
}
// GetLabels gets the labels property value. Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
// returns a []string when successful
func (m *Job) GetLabels()([]string) {
    return m.labels
}
// GetName gets the name property value. The name of the job.
// returns a *string when successful
func (m *Job) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Job) GetNodeId()(*string) {
    return m.node_id
}
// GetRunAttempt gets the run_attempt property value. Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
// returns a *int32 when successful
func (m *Job) GetRunAttempt()(*int32) {
    return m.run_attempt
}
// GetRunId gets the run_id property value. The id of the associated workflow run.
// returns a *int32 when successful
func (m *Job) GetRunId()(*int32) {
    return m.run_id
}
// GetRunnerGroupId gets the runner_group_id property value. The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
// returns a *int32 when successful
func (m *Job) GetRunnerGroupId()(*int32) {
    return m.runner_group_id
}
// GetRunnerGroupName gets the runner_group_name property value. The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
// returns a *string when successful
func (m *Job) GetRunnerGroupName()(*string) {
    return m.runner_group_name
}
// GetRunnerId gets the runner_id property value. The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
// returns a *int32 when successful
func (m *Job) GetRunnerId()(*int32) {
    return m.runner_id
}
// GetRunnerName gets the runner_name property value. The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
// returns a *string when successful
func (m *Job) GetRunnerName()(*string) {
    return m.runner_name
}
// GetRunUrl gets the run_url property value. The run_url property
// returns a *string when successful
func (m *Job) GetRunUrl()(*string) {
    return m.run_url
}
// GetStartedAt gets the started_at property value. The time that the job started, in ISO 8601 format.
// returns a *Time when successful
func (m *Job) GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.started_at
}
// GetStatus gets the status property value. The phase of the lifecycle that the job is currently in.
// returns a *Job_status when successful
func (m *Job) GetStatus()(*Job_status) {
    return m.status
}
// GetSteps gets the steps property value. Steps in this job.
// returns a []Job_stepsable when successful
func (m *Job) GetSteps()([]Job_stepsable) {
    return m.steps
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Job) GetUrl()(*string) {
    return m.url
}
// GetWorkflowName gets the workflow_name property value. The name of the workflow.
// returns a *string when successful
func (m *Job) GetWorkflowName()(*string) {
    return m.workflow_name
}
// Serialize serializes information the current object
func (m *Job) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("check_run_url", m.GetCheckRunUrl())
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
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
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
        err := writer.WriteInt32Value("runner_group_id", m.GetRunnerGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("runner_group_name", m.GetRunnerGroupName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("runner_id", m.GetRunnerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("runner_name", m.GetRunnerName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("run_attempt", m.GetRunAttempt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("run_id", m.GetRunId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("run_url", m.GetRunUrl())
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
    if m.GetSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSteps()))
        for i, v := range m.GetSteps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("steps", cast)
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
        err := writer.WriteStringValue("workflow_name", m.GetWorkflowName())
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
func (m *Job) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckRunUrl sets the check_run_url property value. The check_run_url property
func (m *Job) SetCheckRunUrl(value *string)() {
    m.check_run_url = value
}
// SetCompletedAt sets the completed_at property value. The time that the job finished, in ISO 8601 format.
func (m *Job) SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completed_at = value
}
// SetConclusion sets the conclusion property value. The outcome of the job.
func (m *Job) SetConclusion(value *Job_conclusion)() {
    m.conclusion = value
}
// SetCreatedAt sets the created_at property value. The time that the job created, in ISO 8601 format.
func (m *Job) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHeadBranch sets the head_branch property value. The name of the current branch.
func (m *Job) SetHeadBranch(value *string)() {
    m.head_branch = value
}
// SetHeadSha sets the head_sha property value. The SHA of the commit that is being run.
func (m *Job) SetHeadSha(value *string)() {
    m.head_sha = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Job) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id of the job.
func (m *Job) SetId(value *int32)() {
    m.id = value
}
// SetLabels sets the labels property value. Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
func (m *Job) SetLabels(value []string)() {
    m.labels = value
}
// SetName sets the name property value. The name of the job.
func (m *Job) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Job) SetNodeId(value *string)() {
    m.node_id = value
}
// SetRunAttempt sets the run_attempt property value. Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
func (m *Job) SetRunAttempt(value *int32)() {
    m.run_attempt = value
}
// SetRunId sets the run_id property value. The id of the associated workflow run.
func (m *Job) SetRunId(value *int32)() {
    m.run_id = value
}
// SetRunnerGroupId sets the runner_group_id property value. The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
func (m *Job) SetRunnerGroupId(value *int32)() {
    m.runner_group_id = value
}
// SetRunnerGroupName sets the runner_group_name property value. The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
func (m *Job) SetRunnerGroupName(value *string)() {
    m.runner_group_name = value
}
// SetRunnerId sets the runner_id property value. The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
func (m *Job) SetRunnerId(value *int32)() {
    m.runner_id = value
}
// SetRunnerName sets the runner_name property value. The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
func (m *Job) SetRunnerName(value *string)() {
    m.runner_name = value
}
// SetRunUrl sets the run_url property value. The run_url property
func (m *Job) SetRunUrl(value *string)() {
    m.run_url = value
}
// SetStartedAt sets the started_at property value. The time that the job started, in ISO 8601 format.
func (m *Job) SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.started_at = value
}
// SetStatus sets the status property value. The phase of the lifecycle that the job is currently in.
func (m *Job) SetStatus(value *Job_status)() {
    m.status = value
}
// SetSteps sets the steps property value. Steps in this job.
func (m *Job) SetSteps(value []Job_stepsable)() {
    m.steps = value
}
// SetUrl sets the url property value. The url property
func (m *Job) SetUrl(value *string)() {
    m.url = value
}
// SetWorkflowName sets the workflow_name property value. The name of the workflow.
func (m *Job) SetWorkflowName(value *string)() {
    m.workflow_name = value
}
type Jobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckRunUrl()(*string)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConclusion()(*Job_conclusion)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHeadBranch()(*string)
    GetHeadSha()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLabels()([]string)
    GetName()(*string)
    GetNodeId()(*string)
    GetRunAttempt()(*int32)
    GetRunId()(*int32)
    GetRunnerGroupId()(*int32)
    GetRunnerGroupName()(*string)
    GetRunnerId()(*int32)
    GetRunnerName()(*string)
    GetRunUrl()(*string)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*Job_status)
    GetSteps()([]Job_stepsable)
    GetUrl()(*string)
    GetWorkflowName()(*string)
    SetCheckRunUrl(value *string)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConclusion(value *Job_conclusion)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHeadBranch(value *string)()
    SetHeadSha(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLabels(value []string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetRunAttempt(value *int32)()
    SetRunId(value *int32)()
    SetRunnerGroupId(value *int32)()
    SetRunnerGroupName(value *string)()
    SetRunnerId(value *int32)()
    SetRunnerName(value *string)()
    SetRunUrl(value *string)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *Job_status)()
    SetSteps(value []Job_stepsable)()
    SetUrl(value *string)()
    SetWorkflowName(value *string)()
}
