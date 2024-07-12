package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Event_payload_pages struct {
    // The action property
    action *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html_url property
    html_url *string
    // The page_name property
    page_name *string
    // The sha property
    sha *string
    // The summary property
    summary *string
    // The title property
    title *string
}
// NewEvent_payload_pages instantiates a new Event_payload_pages and sets the default values.
func NewEvent_payload_pages()(*Event_payload_pages) {
    m := &Event_payload_pages{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEvent_payload_pagesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEvent_payload_pagesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvent_payload_pages(), nil
}
// GetAction gets the action property value. The action property
// returns a *string when successful
func (m *Event_payload_pages) GetAction()(*string) {
    return m.action
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Event_payload_pages) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Event_payload_pages) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["page_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageName(val)
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
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Event_payload_pages) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetPageName gets the page_name property value. The page_name property
// returns a *string when successful
func (m *Event_payload_pages) GetPageName()(*string) {
    return m.page_name
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *Event_payload_pages) GetSha()(*string) {
    return m.sha
}
// GetSummary gets the summary property value. The summary property
// returns a *string when successful
func (m *Event_payload_pages) GetSummary()(*string) {
    return m.summary
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *Event_payload_pages) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *Event_payload_pages) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("action", m.GetAction())
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
        err := writer.WriteStringValue("page_name", m.GetPageName())
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
        err := writer.WriteStringValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *Event_payload_pages) SetAction(value *string)() {
    m.action = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Event_payload_pages) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Event_payload_pages) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetPageName sets the page_name property value. The page_name property
func (m *Event_payload_pages) SetPageName(value *string)() {
    m.page_name = value
}
// SetSha sets the sha property value. The sha property
func (m *Event_payload_pages) SetSha(value *string)() {
    m.sha = value
}
// SetSummary sets the summary property value. The summary property
func (m *Event_payload_pages) SetSummary(value *string)() {
    m.summary = value
}
// SetTitle sets the title property value. The title property
func (m *Event_payload_pages) SetTitle(value *string)() {
    m.title = value
}
type Event_payload_pagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetHtmlUrl()(*string)
    GetPageName()(*string)
    GetSha()(*string)
    GetSummary()(*string)
    GetTitle()(*string)
    SetAction(value *string)()
    SetHtmlUrl(value *string)()
    SetPageName(value *string)()
    SetSha(value *string)()
    SetSummary(value *string)()
    SetTitle(value *string)()
}
