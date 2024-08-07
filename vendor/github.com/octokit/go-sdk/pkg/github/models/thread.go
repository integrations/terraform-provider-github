package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Thread thread
type Thread struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *string
    // The last_read_at property
    last_read_at *string
    // The reason property
    reason *string
    // Minimal Repository
    repository MinimalRepositoryable
    // The subject property
    subject Thread_subjectable
    // The subscription_url property
    subscription_url *string
    // The unread property
    unread *bool
    // The updated_at property
    updated_at *string
    // The url property
    url *string
}
// NewThread instantiates a new Thread and sets the default values.
func NewThread()(*Thread) {
    m := &Thread{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateThreadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThread(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Thread) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Thread) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["last_read_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReadAt(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThread_subjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(Thread_subjectable))
        }
        return nil
    }
    res["subscription_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionUrl(val)
        }
        return nil
    }
    res["unread"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnread(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
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
    return res
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *Thread) GetId()(*string) {
    return m.id
}
// GetLastReadAt gets the last_read_at property value. The last_read_at property
// returns a *string when successful
func (m *Thread) GetLastReadAt()(*string) {
    return m.last_read_at
}
// GetReason gets the reason property value. The reason property
// returns a *string when successful
func (m *Thread) GetReason()(*string) {
    return m.reason
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *Thread) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetSubject gets the subject property value. The subject property
// returns a Thread_subjectable when successful
func (m *Thread) GetSubject()(Thread_subjectable) {
    return m.subject
}
// GetSubscriptionUrl gets the subscription_url property value. The subscription_url property
// returns a *string when successful
func (m *Thread) GetSubscriptionUrl()(*string) {
    return m.subscription_url
}
// GetUnread gets the unread property value. The unread property
// returns a *bool when successful
func (m *Thread) GetUnread()(*bool) {
    return m.unread
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *Thread) GetUpdatedAt()(*string) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Thread) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Thread) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_read_at", m.GetLastReadAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscription_url", m.GetSubscriptionUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unread", m.GetUnread())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Thread) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *Thread) SetId(value *string)() {
    m.id = value
}
// SetLastReadAt sets the last_read_at property value. The last_read_at property
func (m *Thread) SetLastReadAt(value *string)() {
    m.last_read_at = value
}
// SetReason sets the reason property value. The reason property
func (m *Thread) SetReason(value *string)() {
    m.reason = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *Thread) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetSubject sets the subject property value. The subject property
func (m *Thread) SetSubject(value Thread_subjectable)() {
    m.subject = value
}
// SetSubscriptionUrl sets the subscription_url property value. The subscription_url property
func (m *Thread) SetSubscriptionUrl(value *string)() {
    m.subscription_url = value
}
// SetUnread sets the unread property value. The unread property
func (m *Thread) SetUnread(value *bool)() {
    m.unread = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Thread) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Thread) SetUrl(value *string)() {
    m.url = value
}
type Threadable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetLastReadAt()(*string)
    GetReason()(*string)
    GetRepository()(MinimalRepositoryable)
    GetSubject()(Thread_subjectable)
    GetSubscriptionUrl()(*string)
    GetUnread()(*bool)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetId(value *string)()
    SetLastReadAt(value *string)()
    SetReason(value *string)()
    SetRepository(value MinimalRepositoryable)()
    SetSubject(value Thread_subjectable)()
    SetSubscriptionUrl(value *string)()
    SetUnread(value *bool)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
