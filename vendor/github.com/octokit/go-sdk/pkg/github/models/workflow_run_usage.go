package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowRunUsage workflow Run Usage
type WorkflowRunUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The billable property
    billable WorkflowRunUsage_billableable
    // The run_duration_ms property
    run_duration_ms *int32
}
// NewWorkflowRunUsage instantiates a new WorkflowRunUsage and sets the default values.
func NewWorkflowRunUsage()(*WorkflowRunUsage) {
    m := &WorkflowRunUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowRunUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkflowRunUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowRunUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkflowRunUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillable gets the billable property value. The billable property
// returns a WorkflowRunUsage_billableable when successful
func (m *WorkflowRunUsage) GetBillable()(WorkflowRunUsage_billableable) {
    return m.billable
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkflowRunUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowRunUsage_billableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillable(val.(WorkflowRunUsage_billableable))
        }
        return nil
    }
    res["run_duration_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunDurationMs(val)
        }
        return nil
    }
    return res
}
// GetRunDurationMs gets the run_duration_ms property value. The run_duration_ms property
// returns a *int32 when successful
func (m *WorkflowRunUsage) GetRunDurationMs()(*int32) {
    return m.run_duration_ms
}
// Serialize serializes information the current object
func (m *WorkflowRunUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("billable", m.GetBillable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("run_duration_ms", m.GetRunDurationMs())
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
func (m *WorkflowRunUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillable sets the billable property value. The billable property
func (m *WorkflowRunUsage) SetBillable(value WorkflowRunUsage_billableable)() {
    m.billable = value
}
// SetRunDurationMs sets the run_duration_ms property value. The run_duration_ms property
func (m *WorkflowRunUsage) SetRunDurationMs(value *int32)() {
    m.run_duration_ms = value
}
type WorkflowRunUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillable()(WorkflowRunUsage_billableable)
    GetRunDurationMs()(*int32)
    SetBillable(value WorkflowRunUsage_billableable)()
    SetRunDurationMs(value *int32)()
}
