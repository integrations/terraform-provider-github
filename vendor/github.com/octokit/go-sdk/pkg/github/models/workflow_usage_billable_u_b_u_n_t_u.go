package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkflowUsage_billable_UBUNTU struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total_ms property
    total_ms *int32
}
// NewWorkflowUsage_billable_UBUNTU instantiates a new WorkflowUsage_billable_UBUNTU and sets the default values.
func NewWorkflowUsage_billable_UBUNTU()(*WorkflowUsage_billable_UBUNTU) {
    m := &WorkflowUsage_billable_UBUNTU{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowUsage_billable_UBUNTUFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkflowUsage_billable_UBUNTUFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowUsage_billable_UBUNTU(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkflowUsage_billable_UBUNTU) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkflowUsage_billable_UBUNTU) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["total_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMs(val)
        }
        return nil
    }
    return res
}
// GetTotalMs gets the total_ms property value. The total_ms property
// returns a *int32 when successful
func (m *WorkflowUsage_billable_UBUNTU) GetTotalMs()(*int32) {
    return m.total_ms
}
// Serialize serializes information the current object
func (m *WorkflowUsage_billable_UBUNTU) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_ms", m.GetTotalMs())
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
func (m *WorkflowUsage_billable_UBUNTU) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalMs sets the total_ms property value. The total_ms property
func (m *WorkflowUsage_billable_UBUNTU) SetTotalMs(value *int32)() {
    m.total_ms = value
}
type WorkflowUsage_billable_UBUNTUable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalMs()(*int32)
    SetTotalMs(value *int32)()
}
