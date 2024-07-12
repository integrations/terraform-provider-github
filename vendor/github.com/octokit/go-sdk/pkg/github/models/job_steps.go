package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Job_steps struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the job finished, in ISO 8601 format.
    completed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The outcome of the job.
    conclusion *string
    // The name of the job.
    name *string
    // The number property
    number *int32
    // The time that the step started, in ISO 8601 format.
    started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The phase of the lifecycle that the job is currently in.
    status *Job_steps_status
}
// NewJob_steps instantiates a new Job_steps and sets the default values.
func NewJob_steps()(*Job_steps) {
    m := &Job_steps{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJob_stepsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJob_stepsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJob_steps(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Job_steps) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompletedAt gets the completed_at property value. The time that the job finished, in ISO 8601 format.
// returns a *Time when successful
func (m *Job_steps) GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completed_at
}
// GetConclusion gets the conclusion property value. The outcome of the job.
// returns a *string when successful
func (m *Job_steps) GetConclusion()(*string) {
    return m.conclusion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Job_steps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConclusion(val)
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
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
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
        val, err := n.GetEnumValue(ParseJob_steps_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Job_steps_status))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the job.
// returns a *string when successful
func (m *Job_steps) GetName()(*string) {
    return m.name
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *Job_steps) GetNumber()(*int32) {
    return m.number
}
// GetStartedAt gets the started_at property value. The time that the step started, in ISO 8601 format.
// returns a *Time when successful
func (m *Job_steps) GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.started_at
}
// GetStatus gets the status property value. The phase of the lifecycle that the job is currently in.
// returns a *Job_steps_status when successful
func (m *Job_steps) GetStatus()(*Job_steps_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *Job_steps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("completed_at", m.GetCompletedAt())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("number", m.GetNumber())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Job_steps) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompletedAt sets the completed_at property value. The time that the job finished, in ISO 8601 format.
func (m *Job_steps) SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completed_at = value
}
// SetConclusion sets the conclusion property value. The outcome of the job.
func (m *Job_steps) SetConclusion(value *string)() {
    m.conclusion = value
}
// SetName sets the name property value. The name of the job.
func (m *Job_steps) SetName(value *string)() {
    m.name = value
}
// SetNumber sets the number property value. The number property
func (m *Job_steps) SetNumber(value *int32)() {
    m.number = value
}
// SetStartedAt sets the started_at property value. The time that the step started, in ISO 8601 format.
func (m *Job_steps) SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.started_at = value
}
// SetStatus sets the status property value. The phase of the lifecycle that the job is currently in.
func (m *Job_steps) SetStatus(value *Job_steps_status)() {
    m.status = value
}
type Job_stepsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConclusion()(*string)
    GetName()(*string)
    GetNumber()(*int32)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*Job_steps_status)
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConclusion(value *string)()
    SetName(value *string)()
    SetNumber(value *int32)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *Job_steps_status)()
}
