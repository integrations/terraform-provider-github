package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TimelineCrossReferencedEvent_source struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
    issue Issueable
    // The type property
    typeEscaped *string
}
// NewTimelineCrossReferencedEvent_source instantiates a new TimelineCrossReferencedEvent_source and sets the default values.
func NewTimelineCrossReferencedEvent_source()(*TimelineCrossReferencedEvent_source) {
    m := &TimelineCrossReferencedEvent_source{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimelineCrossReferencedEvent_sourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineCrossReferencedEvent_sourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimelineCrossReferencedEvent_source(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TimelineCrossReferencedEvent_source) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineCrossReferencedEvent_source) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetIssue gets the issue property value. Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
// returns a Issueable when successful
func (m *TimelineCrossReferencedEvent_source) GetIssue()(Issueable) {
    return m.issue
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *TimelineCrossReferencedEvent_source) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *TimelineCrossReferencedEvent_source) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *TimelineCrossReferencedEvent_source) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIssue sets the issue property value. Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
func (m *TimelineCrossReferencedEvent_source) SetIssue(value Issueable)() {
    m.issue = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *TimelineCrossReferencedEvent_source) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type TimelineCrossReferencedEvent_sourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssue()(Issueable)
    GetTypeEscaped()(*string)
    SetIssue(value Issueable)()
    SetTypeEscaped(value *string)()
}
