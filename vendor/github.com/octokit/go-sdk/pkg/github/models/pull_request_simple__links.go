package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullRequestSimple__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hypermedia Link
    comments Linkable
    // Hypermedia Link
    commits Linkable
    // Hypermedia Link
    html Linkable
    // Hypermedia Link
    issue Linkable
    // Hypermedia Link
    review_comment Linkable
    // Hypermedia Link
    review_comments Linkable
    // Hypermedia Link
    self Linkable
    // Hypermedia Link
    statuses Linkable
}
// NewPullRequestSimple__links instantiates a new PullRequestSimple__links and sets the default values.
func NewPullRequestSimple__links()(*PullRequestSimple__links) {
    m := &PullRequestSimple__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestSimple__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestSimple__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestSimple__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestSimple__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetComments()(Linkable) {
    return m.comments
}
// GetCommits gets the commits property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetCommits()(Linkable) {
    return m.commits
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestSimple__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val.(Linkable))
        }
        return nil
    }
    res["commits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommits(val.(Linkable))
        }
        return nil
    }
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
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val.(Linkable))
        }
        return nil
    }
    res["review_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewComment(val.(Linkable))
        }
        return nil
    }
    res["review_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewComments(val.(Linkable))
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
    res["statuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatuses(val.(Linkable))
        }
        return nil
    }
    return res
}
// GetHtml gets the html property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetHtml()(Linkable) {
    return m.html
}
// GetIssue gets the issue property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetIssue()(Linkable) {
    return m.issue
}
// GetReviewComment gets the review_comment property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetReviewComment()(Linkable) {
    return m.review_comment
}
// GetReviewComments gets the review_comments property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetReviewComments()(Linkable) {
    return m.review_comments
}
// GetSelf gets the self property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetSelf()(Linkable) {
    return m.self
}
// GetStatuses gets the statuses property value. Hypermedia Link
// returns a Linkable when successful
func (m *PullRequestSimple__links) GetStatuses()(Linkable) {
    return m.statuses
}
// Serialize serializes information the current object
func (m *PullRequestSimple__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("commits", m.GetCommits())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("html", m.GetHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("review_comment", m.GetReviewComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("review_comments", m.GetReviewComments())
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
        err := writer.WriteObjectValue("statuses", m.GetStatuses())
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
func (m *PullRequestSimple__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. Hypermedia Link
func (m *PullRequestSimple__links) SetComments(value Linkable)() {
    m.comments = value
}
// SetCommits sets the commits property value. Hypermedia Link
func (m *PullRequestSimple__links) SetCommits(value Linkable)() {
    m.commits = value
}
// SetHtml sets the html property value. Hypermedia Link
func (m *PullRequestSimple__links) SetHtml(value Linkable)() {
    m.html = value
}
// SetIssue sets the issue property value. Hypermedia Link
func (m *PullRequestSimple__links) SetIssue(value Linkable)() {
    m.issue = value
}
// SetReviewComment sets the review_comment property value. Hypermedia Link
func (m *PullRequestSimple__links) SetReviewComment(value Linkable)() {
    m.review_comment = value
}
// SetReviewComments sets the review_comments property value. Hypermedia Link
func (m *PullRequestSimple__links) SetReviewComments(value Linkable)() {
    m.review_comments = value
}
// SetSelf sets the self property value. Hypermedia Link
func (m *PullRequestSimple__links) SetSelf(value Linkable)() {
    m.self = value
}
// SetStatuses sets the statuses property value. Hypermedia Link
func (m *PullRequestSimple__links) SetStatuses(value Linkable)() {
    m.statuses = value
}
type PullRequestSimple__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(Linkable)
    GetCommits()(Linkable)
    GetHtml()(Linkable)
    GetIssue()(Linkable)
    GetReviewComment()(Linkable)
    GetReviewComments()(Linkable)
    GetSelf()(Linkable)
    GetStatuses()(Linkable)
    SetComments(value Linkable)()
    SetCommits(value Linkable)()
    SetHtml(value Linkable)()
    SetIssue(value Linkable)()
    SetReviewComment(value Linkable)()
    SetReviewComments(value Linkable)()
    SetSelf(value Linkable)()
    SetStatuses(value Linkable)()
}
