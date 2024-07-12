package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThreadSubscription thread Subscription
type ThreadSubscription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ignored property
    ignored *bool
    // The reason property
    reason *string
    // The repository_url property
    repository_url *string
    // The subscribed property
    subscribed *bool
    // The thread_url property
    thread_url *string
    // The url property
    url *string
}
// NewThreadSubscription instantiates a new ThreadSubscription and sets the default values.
func NewThreadSubscription()(*ThreadSubscription) {
    m := &ThreadSubscription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateThreadSubscriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreadSubscriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreadSubscription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ThreadSubscription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *ThreadSubscription) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ThreadSubscription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["ignored"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnored(val)
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
    res["repository_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryUrl(val)
        }
        return nil
    }
    res["subscribed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscribed(val)
        }
        return nil
    }
    res["thread_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreadUrl(val)
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
// GetIgnored gets the ignored property value. The ignored property
// returns a *bool when successful
func (m *ThreadSubscription) GetIgnored()(*bool) {
    return m.ignored
}
// GetReason gets the reason property value. The reason property
// returns a *string when successful
func (m *ThreadSubscription) GetReason()(*string) {
    return m.reason
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *ThreadSubscription) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetSubscribed gets the subscribed property value. The subscribed property
// returns a *bool when successful
func (m *ThreadSubscription) GetSubscribed()(*bool) {
    return m.subscribed
}
// GetThreadUrl gets the thread_url property value. The thread_url property
// returns a *string when successful
func (m *ThreadSubscription) GetThreadUrl()(*string) {
    return m.thread_url
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ThreadSubscription) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ThreadSubscription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignored", m.GetIgnored())
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
        err := writer.WriteStringValue("repository_url", m.GetRepositoryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("subscribed", m.GetSubscribed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thread_url", m.GetThreadUrl())
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
func (m *ThreadSubscription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ThreadSubscription) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetIgnored sets the ignored property value. The ignored property
func (m *ThreadSubscription) SetIgnored(value *bool)() {
    m.ignored = value
}
// SetReason sets the reason property value. The reason property
func (m *ThreadSubscription) SetReason(value *string)() {
    m.reason = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *ThreadSubscription) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetSubscribed sets the subscribed property value. The subscribed property
func (m *ThreadSubscription) SetSubscribed(value *bool)() {
    m.subscribed = value
}
// SetThreadUrl sets the thread_url property value. The thread_url property
func (m *ThreadSubscription) SetThreadUrl(value *string)() {
    m.thread_url = value
}
// SetUrl sets the url property value. The url property
func (m *ThreadSubscription) SetUrl(value *string)() {
    m.url = value
}
type ThreadSubscriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIgnored()(*bool)
    GetReason()(*string)
    GetRepositoryUrl()(*string)
    GetSubscribed()(*bool)
    GetThreadUrl()(*string)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIgnored(value *bool)()
    SetReason(value *string)()
    SetRepositoryUrl(value *string)()
    SetSubscribed(value *bool)()
    SetThreadUrl(value *string)()
    SetUrl(value *string)()
}
