package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReviewDismissedIssueEvent_dismissed_review struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dismissal_commit_id property
    dismissal_commit_id *string
    // The dismissal_message property
    dismissal_message *string
    // The review_id property
    review_id *int32
    // The state property
    state *string
}
// NewReviewDismissedIssueEvent_dismissed_review instantiates a new ReviewDismissedIssueEvent_dismissed_review and sets the default values.
func NewReviewDismissedIssueEvent_dismissed_review()(*ReviewDismissedIssueEvent_dismissed_review) {
    m := &ReviewDismissedIssueEvent_dismissed_review{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReviewDismissedIssueEvent_dismissed_reviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReviewDismissedIssueEvent_dismissed_reviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewDismissedIssueEvent_dismissed_review(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReviewDismissedIssueEvent_dismissed_review) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDismissalCommitId gets the dismissal_commit_id property value. The dismissal_commit_id property
// returns a *string when successful
func (m *ReviewDismissedIssueEvent_dismissed_review) GetDismissalCommitId()(*string) {
    return m.dismissal_commit_id
}
// GetDismissalMessage gets the dismissal_message property value. The dismissal_message property
// returns a *string when successful
func (m *ReviewDismissedIssueEvent_dismissed_review) GetDismissalMessage()(*string) {
    return m.dismissal_message
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReviewDismissedIssueEvent_dismissed_review) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dismissal_commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissalCommitId(val)
        }
        return nil
    }
    res["dismissal_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissalMessage(val)
        }
        return nil
    }
    res["review_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewId(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetReviewId gets the review_id property value. The review_id property
// returns a *int32 when successful
func (m *ReviewDismissedIssueEvent_dismissed_review) GetReviewId()(*int32) {
    return m.review_id
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *ReviewDismissedIssueEvent_dismissed_review) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *ReviewDismissedIssueEvent_dismissed_review) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dismissal_commit_id", m.GetDismissalCommitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dismissal_message", m.GetDismissalMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("review_id", m.GetReviewId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *ReviewDismissedIssueEvent_dismissed_review) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDismissalCommitId sets the dismissal_commit_id property value. The dismissal_commit_id property
func (m *ReviewDismissedIssueEvent_dismissed_review) SetDismissalCommitId(value *string)() {
    m.dismissal_commit_id = value
}
// SetDismissalMessage sets the dismissal_message property value. The dismissal_message property
func (m *ReviewDismissedIssueEvent_dismissed_review) SetDismissalMessage(value *string)() {
    m.dismissal_message = value
}
// SetReviewId sets the review_id property value. The review_id property
func (m *ReviewDismissedIssueEvent_dismissed_review) SetReviewId(value *int32)() {
    m.review_id = value
}
// SetState sets the state property value. The state property
func (m *ReviewDismissedIssueEvent_dismissed_review) SetState(value *string)() {
    m.state = value
}
type ReviewDismissedIssueEvent_dismissed_reviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDismissalCommitId()(*string)
    GetDismissalMessage()(*string)
    GetReviewId()(*int32)
    GetState()(*string)
    SetDismissalCommitId(value *string)()
    SetDismissalMessage(value *string)()
    SetReviewId(value *int32)()
    SetState(value *string)()
}
