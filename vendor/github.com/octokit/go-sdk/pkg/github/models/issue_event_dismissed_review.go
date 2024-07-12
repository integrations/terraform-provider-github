package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IssueEventDismissedReview struct {
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
// NewIssueEventDismissedReview instantiates a new IssueEventDismissedReview and sets the default values.
func NewIssueEventDismissedReview()(*IssueEventDismissedReview) {
    m := &IssueEventDismissedReview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssueEventDismissedReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueEventDismissedReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssueEventDismissedReview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IssueEventDismissedReview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDismissalCommitId gets the dismissal_commit_id property value. The dismissal_commit_id property
// returns a *string when successful
func (m *IssueEventDismissedReview) GetDismissalCommitId()(*string) {
    return m.dismissal_commit_id
}
// GetDismissalMessage gets the dismissal_message property value. The dismissal_message property
// returns a *string when successful
func (m *IssueEventDismissedReview) GetDismissalMessage()(*string) {
    return m.dismissal_message
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueEventDismissedReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *IssueEventDismissedReview) GetReviewId()(*int32) {
    return m.review_id
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *IssueEventDismissedReview) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *IssueEventDismissedReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IssueEventDismissedReview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDismissalCommitId sets the dismissal_commit_id property value. The dismissal_commit_id property
func (m *IssueEventDismissedReview) SetDismissalCommitId(value *string)() {
    m.dismissal_commit_id = value
}
// SetDismissalMessage sets the dismissal_message property value. The dismissal_message property
func (m *IssueEventDismissedReview) SetDismissalMessage(value *string)() {
    m.dismissal_message = value
}
// SetReviewId sets the review_id property value. The review_id property
func (m *IssueEventDismissedReview) SetReviewId(value *int32)() {
    m.review_id = value
}
// SetState sets the state property value. The state property
func (m *IssueEventDismissedReview) SetState(value *string)() {
    m.state = value
}
type IssueEventDismissedReviewable interface {
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
