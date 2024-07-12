package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Event_payload struct {
    // The action property
    action *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Comments provide a way for people to collaborate on an issue.
    comment IssueCommentable
    // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
    issue Issueable
    // The pages property
    pages []Event_payload_pagesable
}
// NewEvent_payload instantiates a new Event_payload and sets the default values.
func NewEvent_payload()(*Event_payload) {
    m := &Event_payload{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEvent_payloadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEvent_payloadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvent_payload(), nil
}
// GetAction gets the action property value. The action property
// returns a *string when successful
func (m *Event_payload) GetAction()(*string) {
    return m.action
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Event_payload) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. Comments provide a way for people to collaborate on an issue.
// returns a IssueCommentable when successful
func (m *Event_payload) GetComment()(IssueCommentable) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Event_payload) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val.(IssueCommentable))
        }
        return nil
    }
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val.(Issueable))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEvent_payload_pagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event_payload_pagesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Event_payload_pagesable)
                }
            }
            m.SetPages(res)
        }
        return nil
    }
    return res
}
// GetIssue gets the issue property value. Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
// returns a Issueable when successful
func (m *Event_payload) GetIssue()(Issueable) {
    return m.issue
}
// GetPages gets the pages property value. The pages property
// returns a []Event_payload_pagesable when successful
func (m *Event_payload) GetPages()([]Event_payload_pagesable) {
    return m.pages
}
// Serialize serializes information the current object
func (m *Event_payload) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("comment", m.GetComment())
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
    if m.GetPages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPages()))
        for i, v := range m.GetPages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("pages", cast)
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
// SetAction sets the action property value. The action property
func (m *Event_payload) SetAction(value *string)() {
    m.action = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Event_payload) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. Comments provide a way for people to collaborate on an issue.
func (m *Event_payload) SetComment(value IssueCommentable)() {
    m.comment = value
}
// SetIssue sets the issue property value. Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
func (m *Event_payload) SetIssue(value Issueable)() {
    m.issue = value
}
// SetPages sets the pages property value. The pages property
func (m *Event_payload) SetPages(value []Event_payload_pagesable)() {
    m.pages = value
}
type Event_payloadable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetComment()(IssueCommentable)
    GetIssue()(Issueable)
    GetPages()([]Event_payload_pagesable)
    SetAction(value *string)()
    SetComment(value IssueCommentable)()
    SetIssue(value Issueable)()
    SetPages(value []Event_payload_pagesable)()
}
