package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Actor actor
type Actor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // The display_login property
    display_login *string
    // The gravatar_id property
    gravatar_id *string
    // The id property
    id *int32
    // The login property
    login *string
    // The url property
    url *string
}
// NewActor instantiates a new Actor and sets the default values.
func NewActor()(*Actor) {
    m := &Actor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Actor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *Actor) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetDisplayLogin gets the display_login property value. The display_login property
// returns a *string when successful
func (m *Actor) GetDisplayLogin()(*string) {
    return m.display_login
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Actor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["avatar_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvatarUrl(val)
        }
        return nil
    }
    res["display_login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayLogin(val)
        }
        return nil
    }
    res["gravatar_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGravatarId(val)
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
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
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
// GetGravatarId gets the gravatar_id property value. The gravatar_id property
// returns a *string when successful
func (m *Actor) GetGravatarId()(*string) {
    return m.gravatar_id
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Actor) GetId()(*int32) {
    return m.id
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *Actor) GetLogin()(*string) {
    return m.login
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Actor) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Actor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("display_login", m.GetDisplayLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("gravatar_id", m.GetGravatarId())
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
        err := writer.WriteStringValue("login", m.GetLogin())
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
func (m *Actor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *Actor) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetDisplayLogin sets the display_login property value. The display_login property
func (m *Actor) SetDisplayLogin(value *string)() {
    m.display_login = value
}
// SetGravatarId sets the gravatar_id property value. The gravatar_id property
func (m *Actor) SetGravatarId(value *string)() {
    m.gravatar_id = value
}
// SetId sets the id property value. The id property
func (m *Actor) SetId(value *int32)() {
    m.id = value
}
// SetLogin sets the login property value. The login property
func (m *Actor) SetLogin(value *string)() {
    m.login = value
}
// SetUrl sets the url property value. The url property
func (m *Actor) SetUrl(value *string)() {
    m.url = value
}
type Actorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetDisplayLogin()(*string)
    GetGravatarId()(*string)
    GetId()(*int32)
    GetLogin()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetDisplayLogin(value *string)()
    SetGravatarId(value *string)()
    SetId(value *int32)()
    SetLogin(value *string)()
    SetUrl(value *string)()
}
