package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TimelineReviewedEvent__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html property
    html TimelineReviewedEvent__links_htmlable
    // The pull_request property
    pull_request TimelineReviewedEvent__links_pull_requestable
}
// NewTimelineReviewedEvent__links instantiates a new TimelineReviewedEvent__links and sets the default values.
func NewTimelineReviewedEvent__links()(*TimelineReviewedEvent__links) {
    m := &TimelineReviewedEvent__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineReviewedEvent__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineReviewedEvent__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineReviewedEvent__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineReviewedEvent__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineReviewedEvent__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineReviewedEvent__links_htmlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val.(TimelineReviewedEvent__links_htmlable))
        }
        return nil
    }
    res["pull_request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineReviewedEvent__links_pull_requestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequest(val.(TimelineReviewedEvent__links_pull_requestable))
        }
        return nil
    }
    return res
}
// GetHtml gets the html property value. The html property
// returns a TimelineReviewedEvent__links_htmlable when successful
func (m *TimelineReviewedEvent__links) GetHtml()(TimelineReviewedEvent__links_htmlable) {
    return m.html
}
// GetPullRequest gets the pull_request property value. The pull_request property
// returns a TimelineReviewedEvent__links_pull_requestable when successful
func (m *TimelineReviewedEvent__links) GetPullRequest()(TimelineReviewedEvent__links_pull_requestable) {
    return m.pull_request
}
// Serialize serializes information the current object
func (m *TimelineReviewedEvent__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("html", m.GetHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pull_request", m.GetPullRequest())
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
func (m *TimelineReviewedEvent__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtml sets the html property value. The html property
func (m *TimelineReviewedEvent__links) SetHtml(value TimelineReviewedEvent__links_htmlable)() {
    m.html = value
}
// SetPullRequest sets the pull_request property value. The pull_request property
func (m *TimelineReviewedEvent__links) SetPullRequest(value TimelineReviewedEvent__links_pull_requestable)() {
    m.pull_request = value
}
type TimelineReviewedEvent__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(TimelineReviewedEvent__links_htmlable)
    GetPullRequest()(TimelineReviewedEvent__links_pull_requestable)
    SetHtml(value TimelineReviewedEvent__links_htmlable)()
    SetPullRequest(value TimelineReviewedEvent__links_pull_requestable)()
}
