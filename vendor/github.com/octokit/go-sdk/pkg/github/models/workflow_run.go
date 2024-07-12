package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowRun an invocation of a workflow
type WorkflowRun struct {
    // A GitHub user.
    actor SimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL to the artifacts for the workflow run.
    artifacts_url *string
    // The URL to cancel the workflow run.
    cancel_url *string
    // The ID of the associated check suite.
    check_suite_id *int32
    // The node ID of the associated check suite.
    check_suite_node_id *string
    // The URL to the associated check suite.
    check_suite_url *string
    // The conclusion property
    conclusion *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
    display_title *string
    // The event property
    event *string
    // The head_branch property
    head_branch *string
    // A commit.
    head_commit NullableSimpleCommitable
    // Minimal Repository
    head_repository MinimalRepositoryable
    // The head_repository_id property
    head_repository_id *int32
    // The SHA of the head commit that points to the version of the workflow being run.
    head_sha *string
    // The html_url property
    html_url *string
    // The ID of the workflow run.
    id *int32
    // The URL to the jobs for the workflow run.
    jobs_url *string
    // The URL to download the logs for the workflow run.
    logs_url *string
    // The name of the workflow run.
    name *string
    // The node_id property
    node_id *string
    // The full path of the workflow
    path *string
    // The URL to the previous attempted run of this workflow, if one exists.
    previous_attempt_url *string
    // Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
    pull_requests []PullRequestMinimalable
    // The referenced_workflows property
    referenced_workflows []ReferencedWorkflowable
    // Minimal Repository
    repository MinimalRepositoryable
    // The URL to rerun the workflow run.
    rerun_url *string
    // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
    run_attempt *int32
    // The auto incrementing run number for the workflow run.
    run_number *int32
    // The start time of the latest run. Resets on re-run.
    run_started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *string
    // A GitHub user.
    triggering_actor SimpleUserable
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL to the workflow run.
    url *string
    // The ID of the parent workflow.
    workflow_id *int32
    // The URL to the workflow.
    workflow_url *string
}
// NewWorkflowRun instantiates a new WorkflowRun and sets the default values.
func NewWorkflowRun()(*WorkflowRun) {
    m := &WorkflowRun{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkflowRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowRun(), nil
}
// GetActor gets the actor property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *WorkflowRun) GetActor()(SimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkflowRun) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArtifactsUrl gets the artifacts_url property value. The URL to the artifacts for the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetArtifactsUrl()(*string) {
    return m.artifacts_url
}
// GetCancelUrl gets the cancel_url property value. The URL to cancel the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetCancelUrl()(*string) {
    return m.cancel_url
}
// GetCheckSuiteId gets the check_suite_id property value. The ID of the associated check suite.
// returns a *int32 when successful
func (m *WorkflowRun) GetCheckSuiteId()(*int32) {
    return m.check_suite_id
}
// GetCheckSuiteNodeId gets the check_suite_node_id property value. The node ID of the associated check suite.
// returns a *string when successful
func (m *WorkflowRun) GetCheckSuiteNodeId()(*string) {
    return m.check_suite_node_id
}
// GetCheckSuiteUrl gets the check_suite_url property value. The URL to the associated check suite.
// returns a *string when successful
func (m *WorkflowRun) GetCheckSuiteUrl()(*string) {
    return m.check_suite_url
}
// GetConclusion gets the conclusion property value. The conclusion property
// returns a *string when successful
func (m *WorkflowRun) GetConclusion()(*string) {
    return m.conclusion
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *WorkflowRun) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDisplayTitle gets the display_title property value. The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
// returns a *string when successful
func (m *WorkflowRun) GetDisplayTitle()(*string) {
    return m.display_title
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *WorkflowRun) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkflowRun) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(SimpleUserable))
        }
        return nil
    }
    res["artifacts_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactsUrl(val)
        }
        return nil
    }
    res["cancel_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelUrl(val)
        }
        return nil
    }
    res["check_suite_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckSuiteId(val)
        }
        return nil
    }
    res["check_suite_node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckSuiteNodeId(val)
        }
        return nil
    }
    res["check_suite_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckSuiteUrl(val)
        }
        return nil
    }
    res["conclusion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConclusion(val)
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
    res["display_title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayTitle(val)
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
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
        val, err := n.GetObjectValue(CreateNullableSimpleCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadCommit(val.(NullableSimpleCommitable))
        }
        return nil
    }
    res["head_repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["head_repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadRepositoryId(val)
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
    res["jobs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobsUrl(val)
        }
        return nil
    }
    res["logs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogsUrl(val)
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["previous_attempt_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousAttemptUrl(val)
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
    res["referenced_workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateReferencedWorkflowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReferencedWorkflowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ReferencedWorkflowable)
                }
            }
            m.SetReferencedWorkflows(res)
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
    res["rerun_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRerunUrl(val)
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
    res["run_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunNumber(val)
        }
        return nil
    }
    res["run_started_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunStartedAt(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["triggering_actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTriggeringActor(val.(SimpleUserable))
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
    res["workflow_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowId(val)
        }
        return nil
    }
    res["workflow_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowUrl(val)
        }
        return nil
    }
    return res
}
// GetHeadBranch gets the head_branch property value. The head_branch property
// returns a *string when successful
func (m *WorkflowRun) GetHeadBranch()(*string) {
    return m.head_branch
}
// GetHeadCommit gets the head_commit property value. A commit.
// returns a NullableSimpleCommitable when successful
func (m *WorkflowRun) GetHeadCommit()(NullableSimpleCommitable) {
    return m.head_commit
}
// GetHeadRepository gets the head_repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *WorkflowRun) GetHeadRepository()(MinimalRepositoryable) {
    return m.head_repository
}
// GetHeadRepositoryId gets the head_repository_id property value. The head_repository_id property
// returns a *int32 when successful
func (m *WorkflowRun) GetHeadRepositoryId()(*int32) {
    return m.head_repository_id
}
// GetHeadSha gets the head_sha property value. The SHA of the head commit that points to the version of the workflow being run.
// returns a *string when successful
func (m *WorkflowRun) GetHeadSha()(*string) {
    return m.head_sha
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *WorkflowRun) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The ID of the workflow run.
// returns a *int32 when successful
func (m *WorkflowRun) GetId()(*int32) {
    return m.id
}
// GetJobsUrl gets the jobs_url property value. The URL to the jobs for the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetJobsUrl()(*string) {
    return m.jobs_url
}
// GetLogsUrl gets the logs_url property value. The URL to download the logs for the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetLogsUrl()(*string) {
    return m.logs_url
}
// GetName gets the name property value. The name of the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *WorkflowRun) GetNodeId()(*string) {
    return m.node_id
}
// GetPath gets the path property value. The full path of the workflow
// returns a *string when successful
func (m *WorkflowRun) GetPath()(*string) {
    return m.path
}
// GetPreviousAttemptUrl gets the previous_attempt_url property value. The URL to the previous attempted run of this workflow, if one exists.
// returns a *string when successful
func (m *WorkflowRun) GetPreviousAttemptUrl()(*string) {
    return m.previous_attempt_url
}
// GetPullRequests gets the pull_requests property value. Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
// returns a []PullRequestMinimalable when successful
func (m *WorkflowRun) GetPullRequests()([]PullRequestMinimalable) {
    return m.pull_requests
}
// GetReferencedWorkflows gets the referenced_workflows property value. The referenced_workflows property
// returns a []ReferencedWorkflowable when successful
func (m *WorkflowRun) GetReferencedWorkflows()([]ReferencedWorkflowable) {
    return m.referenced_workflows
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *WorkflowRun) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetRerunUrl gets the rerun_url property value. The URL to rerun the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetRerunUrl()(*string) {
    return m.rerun_url
}
// GetRunAttempt gets the run_attempt property value. Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
// returns a *int32 when successful
func (m *WorkflowRun) GetRunAttempt()(*int32) {
    return m.run_attempt
}
// GetRunNumber gets the run_number property value. The auto incrementing run number for the workflow run.
// returns a *int32 when successful
func (m *WorkflowRun) GetRunNumber()(*int32) {
    return m.run_number
}
// GetRunStartedAt gets the run_started_at property value. The start time of the latest run. Resets on re-run.
// returns a *Time when successful
func (m *WorkflowRun) GetRunStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.run_started_at
}
// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *WorkflowRun) GetStatus()(*string) {
    return m.status
}
// GetTriggeringActor gets the triggering_actor property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *WorkflowRun) GetTriggeringActor()(SimpleUserable) {
    return m.triggering_actor
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *WorkflowRun) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The URL to the workflow run.
// returns a *string when successful
func (m *WorkflowRun) GetUrl()(*string) {
    return m.url
}
// GetWorkflowId gets the workflow_id property value. The ID of the parent workflow.
// returns a *int32 when successful
func (m *WorkflowRun) GetWorkflowId()(*int32) {
    return m.workflow_id
}
// GetWorkflowUrl gets the workflow_url property value. The URL to the workflow.
// returns a *string when successful
func (m *WorkflowRun) GetWorkflowUrl()(*string) {
    return m.workflow_url
}
// Serialize serializes information the current object
func (m *WorkflowRun) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("artifacts_url", m.GetArtifactsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cancel_url", m.GetCancelUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("check_suite_id", m.GetCheckSuiteId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("check_suite_node_id", m.GetCheckSuiteNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("check_suite_url", m.GetCheckSuiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("conclusion", m.GetConclusion())
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
        err := writer.WriteStringValue("display_title", m.GetDisplayTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
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
        err := writer.WriteObjectValue("head_repository", m.GetHeadRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("head_repository_id", m.GetHeadRepositoryId())
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
        err := writer.WriteStringValue("jobs_url", m.GetJobsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logs_url", m.GetLogsUrl())
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
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previous_attempt_url", m.GetPreviousAttemptUrl())
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
    if m.GetReferencedWorkflows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferencedWorkflows()))
        for i, v := range m.GetReferencedWorkflows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("referenced_workflows", cast)
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
        err := writer.WriteStringValue("rerun_url", m.GetRerunUrl())
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
        err := writer.WriteInt32Value("run_number", m.GetRunNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("run_started_at", m.GetRunStartedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("triggering_actor", m.GetTriggeringActor())
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
        err := writer.WriteInt32Value("workflow_id", m.GetWorkflowId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("workflow_url", m.GetWorkflowUrl())
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
// SetActor sets the actor property value. A GitHub user.
func (m *WorkflowRun) SetActor(value SimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkflowRun) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArtifactsUrl sets the artifacts_url property value. The URL to the artifacts for the workflow run.
func (m *WorkflowRun) SetArtifactsUrl(value *string)() {
    m.artifacts_url = value
}
// SetCancelUrl sets the cancel_url property value. The URL to cancel the workflow run.
func (m *WorkflowRun) SetCancelUrl(value *string)() {
    m.cancel_url = value
}
// SetCheckSuiteId sets the check_suite_id property value. The ID of the associated check suite.
func (m *WorkflowRun) SetCheckSuiteId(value *int32)() {
    m.check_suite_id = value
}
// SetCheckSuiteNodeId sets the check_suite_node_id property value. The node ID of the associated check suite.
func (m *WorkflowRun) SetCheckSuiteNodeId(value *string)() {
    m.check_suite_node_id = value
}
// SetCheckSuiteUrl sets the check_suite_url property value. The URL to the associated check suite.
func (m *WorkflowRun) SetCheckSuiteUrl(value *string)() {
    m.check_suite_url = value
}
// SetConclusion sets the conclusion property value. The conclusion property
func (m *WorkflowRun) SetConclusion(value *string)() {
    m.conclusion = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *WorkflowRun) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDisplayTitle sets the display_title property value. The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
func (m *WorkflowRun) SetDisplayTitle(value *string)() {
    m.display_title = value
}
// SetEvent sets the event property value. The event property
func (m *WorkflowRun) SetEvent(value *string)() {
    m.event = value
}
// SetHeadBranch sets the head_branch property value. The head_branch property
func (m *WorkflowRun) SetHeadBranch(value *string)() {
    m.head_branch = value
}
// SetHeadCommit sets the head_commit property value. A commit.
func (m *WorkflowRun) SetHeadCommit(value NullableSimpleCommitable)() {
    m.head_commit = value
}
// SetHeadRepository sets the head_repository property value. Minimal Repository
func (m *WorkflowRun) SetHeadRepository(value MinimalRepositoryable)() {
    m.head_repository = value
}
// SetHeadRepositoryId sets the head_repository_id property value. The head_repository_id property
func (m *WorkflowRun) SetHeadRepositoryId(value *int32)() {
    m.head_repository_id = value
}
// SetHeadSha sets the head_sha property value. The SHA of the head commit that points to the version of the workflow being run.
func (m *WorkflowRun) SetHeadSha(value *string)() {
    m.head_sha = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *WorkflowRun) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The ID of the workflow run.
func (m *WorkflowRun) SetId(value *int32)() {
    m.id = value
}
// SetJobsUrl sets the jobs_url property value. The URL to the jobs for the workflow run.
func (m *WorkflowRun) SetJobsUrl(value *string)() {
    m.jobs_url = value
}
// SetLogsUrl sets the logs_url property value. The URL to download the logs for the workflow run.
func (m *WorkflowRun) SetLogsUrl(value *string)() {
    m.logs_url = value
}
// SetName sets the name property value. The name of the workflow run.
func (m *WorkflowRun) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *WorkflowRun) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPath sets the path property value. The full path of the workflow
func (m *WorkflowRun) SetPath(value *string)() {
    m.path = value
}
// SetPreviousAttemptUrl sets the previous_attempt_url property value. The URL to the previous attempted run of this workflow, if one exists.
func (m *WorkflowRun) SetPreviousAttemptUrl(value *string)() {
    m.previous_attempt_url = value
}
// SetPullRequests sets the pull_requests property value. Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
func (m *WorkflowRun) SetPullRequests(value []PullRequestMinimalable)() {
    m.pull_requests = value
}
// SetReferencedWorkflows sets the referenced_workflows property value. The referenced_workflows property
func (m *WorkflowRun) SetReferencedWorkflows(value []ReferencedWorkflowable)() {
    m.referenced_workflows = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *WorkflowRun) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetRerunUrl sets the rerun_url property value. The URL to rerun the workflow run.
func (m *WorkflowRun) SetRerunUrl(value *string)() {
    m.rerun_url = value
}
// SetRunAttempt sets the run_attempt property value. Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
func (m *WorkflowRun) SetRunAttempt(value *int32)() {
    m.run_attempt = value
}
// SetRunNumber sets the run_number property value. The auto incrementing run number for the workflow run.
func (m *WorkflowRun) SetRunNumber(value *int32)() {
    m.run_number = value
}
// SetRunStartedAt sets the run_started_at property value. The start time of the latest run. Resets on re-run.
func (m *WorkflowRun) SetRunStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.run_started_at = value
}
// SetStatus sets the status property value. The status property
func (m *WorkflowRun) SetStatus(value *string)() {
    m.status = value
}
// SetTriggeringActor sets the triggering_actor property value. A GitHub user.
func (m *WorkflowRun) SetTriggeringActor(value SimpleUserable)() {
    m.triggering_actor = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *WorkflowRun) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The URL to the workflow run.
func (m *WorkflowRun) SetUrl(value *string)() {
    m.url = value
}
// SetWorkflowId sets the workflow_id property value. The ID of the parent workflow.
func (m *WorkflowRun) SetWorkflowId(value *int32)() {
    m.workflow_id = value
}
// SetWorkflowUrl sets the workflow_url property value. The URL to the workflow.
func (m *WorkflowRun) SetWorkflowUrl(value *string)() {
    m.workflow_url = value
}
type WorkflowRunable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(SimpleUserable)
    GetArtifactsUrl()(*string)
    GetCancelUrl()(*string)
    GetCheckSuiteId()(*int32)
    GetCheckSuiteNodeId()(*string)
    GetCheckSuiteUrl()(*string)
    GetConclusion()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayTitle()(*string)
    GetEvent()(*string)
    GetHeadBranch()(*string)
    GetHeadCommit()(NullableSimpleCommitable)
    GetHeadRepository()(MinimalRepositoryable)
    GetHeadRepositoryId()(*int32)
    GetHeadSha()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetJobsUrl()(*string)
    GetLogsUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetPath()(*string)
    GetPreviousAttemptUrl()(*string)
    GetPullRequests()([]PullRequestMinimalable)
    GetReferencedWorkflows()([]ReferencedWorkflowable)
    GetRepository()(MinimalRepositoryable)
    GetRerunUrl()(*string)
    GetRunAttempt()(*int32)
    GetRunNumber()(*int32)
    GetRunStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    GetTriggeringActor()(SimpleUserable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetWorkflowId()(*int32)
    GetWorkflowUrl()(*string)
    SetActor(value SimpleUserable)()
    SetArtifactsUrl(value *string)()
    SetCancelUrl(value *string)()
    SetCheckSuiteId(value *int32)()
    SetCheckSuiteNodeId(value *string)()
    SetCheckSuiteUrl(value *string)()
    SetConclusion(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayTitle(value *string)()
    SetEvent(value *string)()
    SetHeadBranch(value *string)()
    SetHeadCommit(value NullableSimpleCommitable)()
    SetHeadRepository(value MinimalRepositoryable)()
    SetHeadRepositoryId(value *int32)()
    SetHeadSha(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetJobsUrl(value *string)()
    SetLogsUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetPath(value *string)()
    SetPreviousAttemptUrl(value *string)()
    SetPullRequests(value []PullRequestMinimalable)()
    SetReferencedWorkflows(value []ReferencedWorkflowable)()
    SetRepository(value MinimalRepositoryable)()
    SetRerunUrl(value *string)()
    SetRunAttempt(value *int32)()
    SetRunNumber(value *int32)()
    SetRunStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
    SetTriggeringActor(value SimpleUserable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetWorkflowId(value *int32)()
    SetWorkflowUrl(value *string)()
}
