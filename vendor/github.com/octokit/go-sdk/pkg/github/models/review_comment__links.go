package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReviewComment__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hypermedia Link
    html Linkable
    // Hypermedia Link
    pull_request Linkable
    // Hypermedia Link
    self Linkable
}
// NewReviewComment__links instantiates a new ReviewComment__links and sets the default values.
func NewReviewComment__links()(*ReviewComment__links) {
    m := &ReviewComment__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReviewComment__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReviewComment__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewComment__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReviewComment__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReviewComment__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val.(Linkable))
        }
        return nil
    }
    res["pull_request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequest(val.(Linkable))
        }
        return nil
    }
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val.(Linkable))
        }
        return nil
    }
    return res
}
// GetHtml gets the html property value. Hypermedia Link
// returns a Linkable when successful
func (m *ReviewComment__links) GetHtml()(Linkable) {
    return m.html
}
// GetPullRequest gets the pull_request property value. Hypermedia Link
// returns a Linkable when successful
func (m *ReviewComment__links) GetPullRequest()(Linkable) {
    return m.pull_request
}
// GetSelf gets the self property value. Hypermedia Link
// returns a Linkable when successful
func (m *ReviewComment__links) GetSelf()(Linkable) {
    return m.self
}
// Serialize serializes information the current object
func (m *ReviewComment__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("self", m.GetSelf())
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
func (m *ReviewComment__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtml sets the html property value. Hypermedia Link
func (m *ReviewComment__links) SetHtml(value Linkable)() {
    m.html = value
}
// SetPullRequest sets the pull_request property value. Hypermedia Link
func (m *ReviewComment__links) SetPullRequest(value Linkable)() {
    m.pull_request = value
}
// SetSelf sets the self property value. Hypermedia Link
func (m *ReviewComment__links) SetSelf(value Linkable)() {
    m.self = value
}
type ReviewComment__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(Linkable)
    GetPullRequest()(Linkable)
    GetSelf()(Linkable)
    SetHtml(value Linkable)()
    SetPullRequest(value Linkable)()
    SetSelf(value Linkable)()
}
