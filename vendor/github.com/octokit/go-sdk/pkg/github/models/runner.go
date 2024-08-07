package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Runner a self hosted runner
type Runner struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The busy property
    busy *bool
    // The id of the runner.
    id *int32
    // The labels property
    labels []RunnerLabelable
    // The name of the runner.
    name *string
    // The Operating System of the runner.
    os *string
    // The id of the runner group.
    runner_group_id *int32
    // The status of the runner.
    status *string
}
// NewRunner instantiates a new Runner and sets the default values.
func NewRunner()(*Runner) {
    m := &Runner{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRunnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunner(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Runner) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBusy gets the busy property value. The busy property
// returns a *bool when successful
func (m *Runner) GetBusy()(*bool) {
    return m.busy
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Runner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["busy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusy(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateRunnerLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RunnerLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RunnerLabelable)
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
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
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
    return res
}
// GetId gets the id property value. The id of the runner.
// returns a *int32 when successful
func (m *Runner) GetId()(*int32) {
    return m.id
}
// GetLabels gets the labels property value. The labels property
// returns a []RunnerLabelable when successful
func (m *Runner) GetLabels()([]RunnerLabelable) {
    return m.labels
}
// GetName gets the name property value. The name of the runner.
// returns a *string when successful
func (m *Runner) GetName()(*string) {
    return m.name
}
// GetOs gets the os property value. The Operating System of the runner.
// returns a *string when successful
func (m *Runner) GetOs()(*string) {
    return m.os
}
// GetRunnerGroupId gets the runner_group_id property value. The id of the runner group.
// returns a *int32 when successful
func (m *Runner) GetRunnerGroupId()(*int32) {
    return m.runner_group_id
}
// GetStatus gets the status property value. The status of the runner.
// returns a *string when successful
func (m *Runner) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *Runner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("busy", m.GetBusy())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
        err := writer.WriteStringValue("os", m.GetOs())
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
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *Runner) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBusy sets the busy property value. The busy property
func (m *Runner) SetBusy(value *bool)() {
    m.busy = value
}
// SetId sets the id property value. The id of the runner.
func (m *Runner) SetId(value *int32)() {
    m.id = value
}
// SetLabels sets the labels property value. The labels property
func (m *Runner) SetLabels(value []RunnerLabelable)() {
    m.labels = value
}
// SetName sets the name property value. The name of the runner.
func (m *Runner) SetName(value *string)() {
    m.name = value
}
// SetOs sets the os property value. The Operating System of the runner.
func (m *Runner) SetOs(value *string)() {
    m.os = value
}
// SetRunnerGroupId sets the runner_group_id property value. The id of the runner group.
func (m *Runner) SetRunnerGroupId(value *int32)() {
    m.runner_group_id = value
}
// SetStatus sets the status property value. The status of the runner.
func (m *Runner) SetStatus(value *string)() {
    m.status = value
}
type Runnerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusy()(*bool)
    GetId()(*int32)
    GetLabels()([]RunnerLabelable)
    GetName()(*string)
    GetOs()(*string)
    GetRunnerGroupId()(*int32)
    GetStatus()(*string)
    SetBusy(value *bool)()
    SetId(value *int32)()
    SetLabels(value []RunnerLabelable)()
    SetName(value *string)()
    SetOs(value *string)()
    SetRunnerGroupId(value *int32)()
    SetStatus(value *string)()
}
