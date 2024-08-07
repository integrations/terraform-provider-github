package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullRequestReviewComment__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html property
    html PullRequestReviewComment__links_htmlable
    // The pull_request property
    pull_request PullRequestReviewComment__links_pull_requestable
    // The self property
    self PullRequestReviewComment__links_selfable
}
// NewPullRequestReviewComment__links instantiates a new PullRequestReviewComment__links and sets the default values.
func NewPullRequestReviewComment__links()(*PullRequestReviewComment__links) {
    m := &PullRequestReviewComment__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestReviewComment__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestReviewComment__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestReviewComment__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestReviewComment__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestReviewComment__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestReviewComment__links_htmlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val.(PullRequestReviewComment__links_htmlable))
        }
        return nil
    }
    res["pull_request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestReviewComment__links_pull_requestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequest(val.(PullRequestReviewComment__links_pull_requestable))
        }
        return nil
    }
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestReviewComment__links_selfFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val.(PullRequestReviewComment__links_selfable))
        }
        return nil
    }
    return res
}
// GetHtml gets the html property value. The html property
// returns a PullRequestReviewComment__links_htmlable when successful
func (m *PullRequestReviewComment__links) GetHtml()(PullRequestReviewComment__links_htmlable) {
    return m.html
}
// GetPullRequest gets the pull_request property value. The pull_request property
// returns a PullRequestReviewComment__links_pull_requestable when successful
func (m *PullRequestReviewComment__links) GetPullRequest()(PullRequestReviewComment__links_pull_requestable) {
    return m.pull_request
}
// GetSelf gets the self property value. The self property
// returns a PullRequestReviewComment__links_selfable when successful
func (m *PullRequestReviewComment__links) GetSelf()(PullRequestReviewComment__links_selfable) {
    return m.self
}
// Serialize serializes information the current object
func (m *PullRequestReviewComment__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PullRequestReviewComment__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtml sets the html property value. The html property
func (m *PullRequestReviewComment__links) SetHtml(value PullRequestReviewComment__links_htmlable)() {
    m.html = value
}
// SetPullRequest sets the pull_request property value. The pull_request property
func (m *PullRequestReviewComment__links) SetPullRequest(value PullRequestReviewComment__links_pull_requestable)() {
    m.pull_request = value
}
// SetSelf sets the self property value. The self property
func (m *PullRequestReviewComment__links) SetSelf(value PullRequestReviewComment__links_selfable)() {
    m.self = value
}
type PullRequestReviewComment__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(PullRequestReviewComment__links_htmlable)
    GetPullRequest()(PullRequestReviewComment__links_pull_requestable)
    GetSelf()(PullRequestReviewComment__links_selfable)
    SetHtml(value PullRequestReviewComment__links_htmlable)()
    SetPullRequest(value PullRequestReviewComment__links_pull_requestable)()
    SetSelf(value PullRequestReviewComment__links_selfable)()
}
