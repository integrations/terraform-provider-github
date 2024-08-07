package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeployKey an SSH key granting access to a single repository.
type DeployKey struct {
    // The added_by property
    added_by *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The id property
    id *int32
    // The key property
    key *string
    // The last_used property
    last_used *string
    // The read_only property
    read_only *bool
    // The title property
    title *string
    // The url property
    url *string
    // The verified property
    verified *bool
}
// NewDeployKey instantiates a new DeployKey and sets the default values.
func NewDeployKey()(*DeployKey) {
    m := &DeployKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeployKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeployKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeployKey(), nil
}
// GetAddedBy gets the added_by property value. The added_by property
// returns a *string when successful
func (m *DeployKey) GetAddedBy()(*string) {
    return m.added_by
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeployKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *DeployKey) GetCreatedAt()(*string) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeployKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedBy(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["last_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val)
        }
        return nil
    }
    res["read_only"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnly(val)
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
    res["verified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerified(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *DeployKey) GetId()(*int32) {
    return m.id
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *DeployKey) GetKey()(*string) {
    return m.key
}
// GetLastUsed gets the last_used property value. The last_used property
// returns a *string when successful
func (m *DeployKey) GetLastUsed()(*string) {
    return m.last_used
}
// GetReadOnly gets the read_only property value. The read_only property
// returns a *bool when successful
func (m *DeployKey) GetReadOnly()(*bool) {
    return m.read_only
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *DeployKey) GetTitle()(*string) {
    return m.title
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *DeployKey) GetUrl()(*string) {
    return m.url
}
// GetVerified gets the verified property value. The verified property
// returns a *bool when successful
func (m *DeployKey) GetVerified()(*bool) {
    return m.verified
}
// Serialize serializes information the current object
func (m *DeployKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("added_by", m.GetAddedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_used", m.GetLastUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("read_only", m.GetReadOnly())
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
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("verified", m.GetVerified())
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
// SetAddedBy sets the added_by property value. The added_by property
func (m *DeployKey) SetAddedBy(value *string)() {
    m.added_by = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeployKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *DeployKey) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetId sets the id property value. The id property
func (m *DeployKey) SetId(value *int32)() {
    m.id = value
}
// SetKey sets the key property value. The key property
func (m *DeployKey) SetKey(value *string)() {
    m.key = value
}
// SetLastUsed sets the last_used property value. The last_used property
func (m *DeployKey) SetLastUsed(value *string)() {
    m.last_used = value
}
// SetReadOnly sets the read_only property value. The read_only property
func (m *DeployKey) SetReadOnly(value *bool)() {
    m.read_only = value
}
// SetTitle sets the title property value. The title property
func (m *DeployKey) SetTitle(value *string)() {
    m.title = value
}
// SetUrl sets the url property value. The url property
func (m *DeployKey) SetUrl(value *string)() {
    m.url = value
}
// SetVerified sets the verified property value. The verified property
func (m *DeployKey) SetVerified(value *bool)() {
    m.verified = value
}
type DeployKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedBy()(*string)
    GetCreatedAt()(*string)
    GetId()(*int32)
    GetKey()(*string)
    GetLastUsed()(*string)
    GetReadOnly()(*bool)
    GetTitle()(*string)
    GetUrl()(*string)
    GetVerified()(*bool)
    SetAddedBy(value *string)()
    SetCreatedAt(value *string)()
    SetId(value *int32)()
    SetKey(value *string)()
    SetLastUsed(value *string)()
    SetReadOnly(value *bool)()
    SetTitle(value *string)()
    SetUrl(value *string)()
    SetVerified(value *bool)()
}
