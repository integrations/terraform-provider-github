package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowUsage workflow Usage
type WorkflowUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The billable property
    billable WorkflowUsage_billableable
}
// NewWorkflowUsage instantiates a new WorkflowUsage and sets the default values.
func NewWorkflowUsage()(*WorkflowUsage) {
    m := &WorkflowUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkflowUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkflowUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillable gets the billable property value. The billable property
// returns a WorkflowUsage_billableable when successful
func (m *WorkflowUsage) GetBillable()(WorkflowUsage_billableable) {
    return m.billable
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkflowUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowUsage_billableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillable(val.(WorkflowUsage_billableable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WorkflowUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("billable", m.GetBillable())
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
func (m *WorkflowUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillable sets the billable property value. The billable property
func (m *WorkflowUsage) SetBillable(value WorkflowUsage_billableable)() {
    m.billable = value
}
type WorkflowUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillable()(WorkflowUsage_billableable)
    SetBillable(value WorkflowUsage_billableable)()
}
