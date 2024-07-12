package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkflowRunUsage_billable_WINDOWS_job_runs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The duration_ms property
    duration_ms *int32
    // The job_id property
    job_id *int32
}
// NewWorkflowRunUsage_billable_WINDOWS_job_runs instantiates a new WorkflowRunUsage_billable_WINDOWS_job_runs and sets the default values.
func NewWorkflowRunUsage_billable_WINDOWS_job_runs()(*WorkflowRunUsage_billable_WINDOWS_job_runs) {
    m := &WorkflowRunUsage_billable_WINDOWS_job_runs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowRunUsage_billable_WINDOWS_job_runsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkflowRunUsage_billable_WINDOWS_job_runsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowRunUsage_billable_WINDOWS_job_runs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDurationMs gets the duration_ms property value. The duration_ms property
// returns a *int32 when successful
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) GetDurationMs()(*int32) {
    return m.duration_ms
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["duration_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationMs(val)
        }
        return nil
    }
    res["job_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobId(val)
        }
        return nil
    }
    return res
}
// GetJobId gets the job_id property value. The job_id property
// returns a *int32 when successful
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) GetJobId()(*int32) {
    return m.job_id
}
// Serialize serializes information the current object
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("duration_ms", m.GetDurationMs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("job_id", m.GetJobId())
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
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDurationMs sets the duration_ms property value. The duration_ms property
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) SetDurationMs(value *int32)() {
    m.duration_ms = value
}
// SetJobId sets the job_id property value. The job_id property
func (m *WorkflowRunUsage_billable_WINDOWS_job_runs) SetJobId(value *int32)() {
    m.job_id = value
}
type WorkflowRunUsage_billable_WINDOWS_job_runsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDurationMs()(*int32)
    GetJobId()(*int32)
    SetDurationMs(value *int32)()
    SetJobId(value *int32)()
}
