package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingDeployment details of a deployment that is waiting for protection rules to pass
type PendingDeployment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the currently authenticated user can approve the deployment
    current_user_can_approve *bool
    // The environment property
    environment PendingDeployment_environmentable
    // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
    reviewers []PendingDeployment_reviewersable
    // The set duration of the wait timer
    wait_timer *int32
    // The time that the wait timer began.
    wait_timer_started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewPendingDeployment instantiates a new PendingDeployment and sets the default values.
func NewPendingDeployment()(*PendingDeployment) {
    m := &PendingDeployment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePendingDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePendingDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPendingDeployment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PendingDeployment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentUserCanApprove gets the current_user_can_approve property value. Whether the currently authenticated user can approve the deployment
// returns a *bool when successful
func (m *PendingDeployment) GetCurrentUserCanApprove()(*bool) {
    return m.current_user_can_approve
}
// GetEnvironment gets the environment property value. The environment property
// returns a PendingDeployment_environmentable when successful
func (m *PendingDeployment) GetEnvironment()(PendingDeployment_environmentable) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PendingDeployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["current_user_can_approve"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserCanApprove(val)
        }
        return nil
    }
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingDeployment_environmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val.(PendingDeployment_environmentable))
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePendingDeployment_reviewersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PendingDeployment_reviewersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PendingDeployment_reviewersable)
                }
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["wait_timer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitTimer(val)
        }
        return nil
    }
    res["wait_timer_started_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitTimerStartedAt(val)
        }
        return nil
    }
    return res
}
// GetReviewers gets the reviewers property value. The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
// returns a []PendingDeployment_reviewersable when successful
func (m *PendingDeployment) GetReviewers()([]PendingDeployment_reviewersable) {
    return m.reviewers
}
// GetWaitTimer gets the wait_timer property value. The set duration of the wait timer
// returns a *int32 when successful
func (m *PendingDeployment) GetWaitTimer()(*int32) {
    return m.wait_timer
}
// GetWaitTimerStartedAt gets the wait_timer_started_at property value. The time that the wait timer began.
// returns a *Time when successful
func (m *PendingDeployment) GetWaitTimerStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.wait_timer_started_at
}
// Serialize serializes information the current object
func (m *PendingDeployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("current_user_can_approve", m.GetCurrentUserCanApprove())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wait_timer", m.GetWaitTimer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("wait_timer_started_at", m.GetWaitTimerStartedAt())
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
func (m *PendingDeployment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentUserCanApprove sets the current_user_can_approve property value. Whether the currently authenticated user can approve the deployment
func (m *PendingDeployment) SetCurrentUserCanApprove(value *bool)() {
    m.current_user_can_approve = value
}
// SetEnvironment sets the environment property value. The environment property
func (m *PendingDeployment) SetEnvironment(value PendingDeployment_environmentable)() {
    m.environment = value
}
// SetReviewers sets the reviewers property value. The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *PendingDeployment) SetReviewers(value []PendingDeployment_reviewersable)() {
    m.reviewers = value
}
// SetWaitTimer sets the wait_timer property value. The set duration of the wait timer
func (m *PendingDeployment) SetWaitTimer(value *int32)() {
    m.wait_timer = value
}
// SetWaitTimerStartedAt sets the wait_timer_started_at property value. The time that the wait timer began.
func (m *PendingDeployment) SetWaitTimerStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.wait_timer_started_at = value
}
type PendingDeploymentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUserCanApprove()(*bool)
    GetEnvironment()(PendingDeployment_environmentable)
    GetReviewers()([]PendingDeployment_reviewersable)
    GetWaitTimer()(*int32)
    GetWaitTimerStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCurrentUserCanApprove(value *bool)()
    SetEnvironment(value PendingDeployment_environmentable)()
    SetReviewers(value []PendingDeployment_reviewersable)()
    SetWaitTimer(value *int32)()
    SetWaitTimerStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
