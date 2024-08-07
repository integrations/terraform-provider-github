package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkflowRunUsage_billable struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The MACOS property
    mACOS WorkflowRunUsage_billable_MACOSable
    // The UBUNTU property
    uBUNTU WorkflowRunUsage_billable_UBUNTUable
    // The WINDOWS property
    wINDOWS WorkflowRunUsage_billable_WINDOWSable
}
// NewWorkflowRunUsage_billable instantiates a new WorkflowRunUsage_billable and sets the default values.
func NewWorkflowRunUsage_billable()(*WorkflowRunUsage_billable) {
    m := &WorkflowRunUsage_billable{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowRunUsage_billableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkflowRunUsage_billableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowRunUsage_billable(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkflowRunUsage_billable) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkflowRunUsage_billable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["MACOS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowRunUsage_billable_MACOSFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMACOS(val.(WorkflowRunUsage_billable_MACOSable))
        }
        return nil
    }
    res["UBUNTU"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowRunUsage_billable_UBUNTUFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUBUNTU(val.(WorkflowRunUsage_billable_UBUNTUable))
        }
        return nil
    }
    res["WINDOWS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowRunUsage_billable_WINDOWSFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWINDOWS(val.(WorkflowRunUsage_billable_WINDOWSable))
        }
        return nil
    }
    return res
}
// GetMACOS gets the MACOS property value. The MACOS property
// returns a WorkflowRunUsage_billable_MACOSable when successful
func (m *WorkflowRunUsage_billable) GetMACOS()(WorkflowRunUsage_billable_MACOSable) {
    return m.mACOS
}
// GetUBUNTU gets the UBUNTU property value. The UBUNTU property
// returns a WorkflowRunUsage_billable_UBUNTUable when successful
func (m *WorkflowRunUsage_billable) GetUBUNTU()(WorkflowRunUsage_billable_UBUNTUable) {
    return m.uBUNTU
}
// GetWINDOWS gets the WINDOWS property value. The WINDOWS property
// returns a WorkflowRunUsage_billable_WINDOWSable when successful
func (m *WorkflowRunUsage_billable) GetWINDOWS()(WorkflowRunUsage_billable_WINDOWSable) {
    return m.wINDOWS
}
// Serialize serializes information the current object
func (m *WorkflowRunUsage_billable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("MACOS", m.GetMACOS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("UBUNTU", m.GetUBUNTU())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("WINDOWS", m.GetWINDOWS())
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
func (m *WorkflowRunUsage_billable) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMACOS sets the MACOS property value. The MACOS property
func (m *WorkflowRunUsage_billable) SetMACOS(value WorkflowRunUsage_billable_MACOSable)() {
    m.mACOS = value
}
// SetUBUNTU sets the UBUNTU property value. The UBUNTU property
func (m *WorkflowRunUsage_billable) SetUBUNTU(value WorkflowRunUsage_billable_UBUNTUable)() {
    m.uBUNTU = value
}
// SetWINDOWS sets the WINDOWS property value. The WINDOWS property
func (m *WorkflowRunUsage_billable) SetWINDOWS(value WorkflowRunUsage_billable_WINDOWSable)() {
    m.wINDOWS = value
}
type WorkflowRunUsage_billableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMACOS()(WorkflowRunUsage_billable_MACOSable)
    GetUBUNTU()(WorkflowRunUsage_billable_UBUNTUable)
    GetWINDOWS()(WorkflowRunUsage_billable_WINDOWSable)
    SetMACOS(value WorkflowRunUsage_billable_MACOSable)()
    SetUBUNTU(value WorkflowRunUsage_billable_UBUNTUable)()
    SetWINDOWS(value WorkflowRunUsage_billable_WINDOWSable)()
}
