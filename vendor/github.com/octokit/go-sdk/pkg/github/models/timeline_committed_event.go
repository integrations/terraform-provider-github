package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineCommittedEvent timeline Committed Event
type TimelineCommittedEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifying information for the git-user
    author TimelineCommittedEvent_authorable
    // Identifying information for the git-user
    committer TimelineCommittedEvent_committerable
    // The event property
    event *string
    // The html_url property
    html_url *string
    // Message describing the purpose of the commit
    message *string
    // The node_id property
    node_id *string
    // The parents property
    parents []TimelineCommittedEvent_parentsable
    // SHA for the commit
    sha *string
    // The tree property
    tree TimelineCommittedEvent_treeable
    // The url property
    url *string
    // The verification property
    verification TimelineCommittedEvent_verificationable
}
// NewTimelineCommittedEvent instantiates a new TimelineCommittedEvent and sets the default values.
func NewTimelineCommittedEvent()(*TimelineCommittedEvent) {
    m := &TimelineCommittedEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineCommittedEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineCommittedEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineCommittedEvent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineCommittedEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. Identifying information for the git-user
// returns a TimelineCommittedEvent_authorable when successful
func (m *TimelineCommittedEvent) GetAuthor()(TimelineCommittedEvent_authorable) {
    return m.author
}
// GetCommitter gets the committer property value. Identifying information for the git-user
// returns a TimelineCommittedEvent_committerable when successful
func (m *TimelineCommittedEvent) GetCommitter()(TimelineCommittedEvent_committerable) {
    return m.committer
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *TimelineCommittedEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineCommittedEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineCommittedEvent_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(TimelineCommittedEvent_authorable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineCommittedEvent_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(TimelineCommittedEvent_committerable))
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["parents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimelineCommittedEvent_parentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimelineCommittedEvent_parentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TimelineCommittedEvent_parentsable)
                }
            }
            m.SetParents(res)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["tree"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineCommittedEvent_treeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTree(val.(TimelineCommittedEvent_treeable))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["verification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimelineCommittedEvent_verificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerification(val.(TimelineCommittedEvent_verificationable))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *TimelineCommittedEvent) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetMessage gets the message property value. Message describing the purpose of the commit
// returns a *string when successful
func (m *TimelineCommittedEvent) GetMessage()(*string) {
    return m.message
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TimelineCommittedEvent) GetNodeId()(*string) {
    return m.node_id
}
// GetParents gets the parents property value. The parents property
// returns a []TimelineCommittedEvent_parentsable when successful
func (m *TimelineCommittedEvent) GetParents()([]TimelineCommittedEvent_parentsable) {
    return m.parents
}
// GetSha gets the sha property value. SHA for the commit
// returns a *string when successful
func (m *TimelineCommittedEvent) GetSha()(*string) {
    return m.sha
}
// GetTree gets the tree property value. The tree property
// returns a TimelineCommittedEvent_treeable when successful
func (m *TimelineCommittedEvent) GetTree()(TimelineCommittedEvent_treeable) {
    return m.tree
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TimelineCommittedEvent) GetUrl()(*string) {
    return m.url
}
// GetVerification gets the verification property value. The verification property
// returns a TimelineCommittedEvent_verificationable when successful
func (m *TimelineCommittedEvent) GetVerification()(TimelineCommittedEvent_verificationable) {
    return m.verification
}
// Serialize serializes information the current object
func (m *TimelineCommittedEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("committer", m.GetCommitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    if m.GetParents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParents()))
        for i, v := range m.GetParents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("parents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tree", m.GetTree())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verification", m.GetVerification())
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
func (m *TimelineCommittedEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. Identifying information for the git-user
func (m *TimelineCommittedEvent) SetAuthor(value TimelineCommittedEvent_authorable)() {
    m.author = value
}
// SetCommitter sets the committer property value. Identifying information for the git-user
func (m *TimelineCommittedEvent) SetCommitter(value TimelineCommittedEvent_committerable)() {
    m.committer = value
}
// SetEvent sets the event property value. The event property
func (m *TimelineCommittedEvent) SetEvent(value *string)() {
    m.event = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TimelineCommittedEvent) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetMessage sets the message property value. Message describing the purpose of the commit
func (m *TimelineCommittedEvent) SetMessage(value *string)() {
    m.message = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TimelineCommittedEvent) SetNodeId(value *string)() {
    m.node_id = value
}
// SetParents sets the parents property value. The parents property
func (m *TimelineCommittedEvent) SetParents(value []TimelineCommittedEvent_parentsable)() {
    m.parents = value
}
// SetSha sets the sha property value. SHA for the commit
func (m *TimelineCommittedEvent) SetSha(value *string)() {
    m.sha = value
}
// SetTree sets the tree property value. The tree property
func (m *TimelineCommittedEvent) SetTree(value TimelineCommittedEvent_treeable)() {
    m.tree = value
}
// SetUrl sets the url property value. The url property
func (m *TimelineCommittedEvent) SetUrl(value *string)() {
    m.url = value
}
// SetVerification sets the verification property value. The verification property
func (m *TimelineCommittedEvent) SetVerification(value TimelineCommittedEvent_verificationable)() {
    m.verification = value
}
type TimelineCommittedEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(TimelineCommittedEvent_authorable)
    GetCommitter()(TimelineCommittedEvent_committerable)
    GetEvent()(*string)
    GetHtmlUrl()(*string)
    GetMessage()(*string)
    GetNodeId()(*string)
    GetParents()([]TimelineCommittedEvent_parentsable)
    GetSha()(*string)
    GetTree()(TimelineCommittedEvent_treeable)
    GetUrl()(*string)
    GetVerification()(TimelineCommittedEvent_verificationable)
    SetAuthor(value TimelineCommittedEvent_authorable)()
    SetCommitter(value TimelineCommittedEvent_committerable)()
    SetEvent(value *string)()
    SetHtmlUrl(value *string)()
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetParents(value []TimelineCommittedEvent_parentsable)()
    SetSha(value *string)()
    SetTree(value TimelineCommittedEvent_treeable)()
    SetUrl(value *string)()
    SetVerification(value TimelineCommittedEvent_verificationable)()
}
