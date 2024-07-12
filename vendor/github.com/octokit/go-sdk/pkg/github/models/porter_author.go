package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PorterAuthor porter Author
type PorterAuthor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The id property
    id *int32
    // The import_url property
    import_url *string
    // The name property
    name *string
    // The remote_id property
    remote_id *string
    // The remote_name property
    remote_name *string
    // The url property
    url *string
}
// NewPorterAuthor instantiates a new PorterAuthor and sets the default values.
func NewPorterAuthor()(*PorterAuthor) {
    m := &PorterAuthor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePorterAuthorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePorterAuthorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPorterAuthor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PorterAuthor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *PorterAuthor) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PorterAuthor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    res["import_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportUrl(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["remote_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteId(val)
        }
        return nil
    }
    res["remote_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteName(val)
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
// returns a *int32 when successful
func (m *PorterAuthor) GetId()(*int32) {
    return m.id
}
// GetImportUrl gets the import_url property value. The import_url property
// returns a *string when successful
func (m *PorterAuthor) GetImportUrl()(*string) {
    return m.import_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *PorterAuthor) GetName()(*string) {
    return m.name
}
// GetRemoteId gets the remote_id property value. The remote_id property
// returns a *string when successful
func (m *PorterAuthor) GetRemoteId()(*string) {
    return m.remote_id
}
// GetRemoteName gets the remote_name property value. The remote_name property
// returns a *string when successful
func (m *PorterAuthor) GetRemoteName()(*string) {
    return m.remote_name
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PorterAuthor) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PorterAuthor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteStringValue("import_url", m.GetImportUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remote_id", m.GetRemoteId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remote_name", m.GetRemoteName())
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
func (m *PorterAuthor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *PorterAuthor) SetEmail(value *string)() {
    m.email = value
}
// SetId sets the id property value. The id property
func (m *PorterAuthor) SetId(value *int32)() {
    m.id = value
}
// SetImportUrl sets the import_url property value. The import_url property
func (m *PorterAuthor) SetImportUrl(value *string)() {
    m.import_url = value
}
// SetName sets the name property value. The name property
func (m *PorterAuthor) SetName(value *string)() {
    m.name = value
}
// SetRemoteId sets the remote_id property value. The remote_id property
func (m *PorterAuthor) SetRemoteId(value *string)() {
    m.remote_id = value
}
// SetRemoteName sets the remote_name property value. The remote_name property
func (m *PorterAuthor) SetRemoteName(value *string)() {
    m.remote_name = value
}
// SetUrl sets the url property value. The url property
func (m *PorterAuthor) SetUrl(value *string)() {
    m.url = value
}
type PorterAuthorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetId()(*int32)
    GetImportUrl()(*string)
    GetName()(*string)
    GetRemoteId()(*string)
    GetRemoteName()(*string)
    GetUrl()(*string)
    SetEmail(value *string)()
    SetId(value *int32)()
    SetImportUrl(value *string)()
    SetName(value *string)()
    SetRemoteId(value *string)()
    SetRemoteName(value *string)()
    SetUrl(value *string)()
}
