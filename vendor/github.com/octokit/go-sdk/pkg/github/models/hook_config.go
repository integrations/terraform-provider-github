package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Hook_config 
type Hook_config struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
    content_type *string
    // The digest property
    digest *string
    // The email property
    email *string
    // The insecure_ssl property
    insecure_ssl WebhookConfigInsecureSslable
    // The password property
    password *string
    // The room property
    room *string
    // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
    secret *string
    // The subdomain property
    subdomain *string
    // The token property
    token *string
    // The URL to which the payloads will be delivered.
    url *string
}
// NewHook_config instantiates a new hook_config and sets the default values.
func NewHook_config()(*Hook_config) {
    m := &Hook_config{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHook_configFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHook_configFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHook_config(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Hook_config) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentType gets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *Hook_config) GetContentType()(*string) {
    return m.content_type
}
// GetDigest gets the digest property value. The digest property
func (m *Hook_config) GetDigest()(*string) {
    return m.digest
}
// GetEmail gets the email property value. The email property
func (m *Hook_config) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Hook_config) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["digest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
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
    res["insecure_ssl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookConfigInsecureSslFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsecureSsl(val.(WebhookConfigInsecureSslable))
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["room"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoom(val)
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    res["subdomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubdomain(val)
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
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
// GetInsecureSsl gets the insecure_ssl property value. The insecure_ssl property
func (m *Hook_config) GetInsecureSsl()(WebhookConfigInsecureSslable) {
    return m.insecure_ssl
}
// GetPassword gets the password property value. The password property
func (m *Hook_config) GetPassword()(*string) {
    return m.password
}
// GetRoom gets the room property value. The room property
func (m *Hook_config) GetRoom()(*string) {
    return m.room
}
// GetSecret gets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *Hook_config) GetSecret()(*string) {
    return m.secret
}
// GetSubdomain gets the subdomain property value. The subdomain property
func (m *Hook_config) GetSubdomain()(*string) {
    return m.subdomain
}
// GetToken gets the token property value. The token property
func (m *Hook_config) GetToken()(*string) {
    return m.token
}
// GetUrl gets the url property value. The URL to which the payloads will be delivered.
func (m *Hook_config) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Hook_config) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("insecure_ssl", m.GetInsecureSsl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("room", m.GetRoom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subdomain", m.GetSubdomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *Hook_config) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentType sets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *Hook_config) SetContentType(value *string)() {
    m.content_type = value
}
// SetDigest sets the digest property value. The digest property
func (m *Hook_config) SetDigest(value *string)() {
    m.digest = value
}
// SetEmail sets the email property value. The email property
func (m *Hook_config) SetEmail(value *string)() {
    m.email = value
}
// SetInsecureSsl sets the insecure_ssl property value. The insecure_ssl property
func (m *Hook_config) SetInsecureSsl(value WebhookConfigInsecureSslable)() {
    m.insecure_ssl = value
}
// SetPassword sets the password property value. The password property
func (m *Hook_config) SetPassword(value *string)() {
    m.password = value
}
// SetRoom sets the room property value. The room property
func (m *Hook_config) SetRoom(value *string)() {
    m.room = value
}
// SetSecret sets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *Hook_config) SetSecret(value *string)() {
    m.secret = value
}
// SetSubdomain sets the subdomain property value. The subdomain property
func (m *Hook_config) SetSubdomain(value *string)() {
    m.subdomain = value
}
// SetToken sets the token property value. The token property
func (m *Hook_config) SetToken(value *string)() {
    m.token = value
}
// SetUrl sets the url property value. The URL to which the payloads will be delivered.
func (m *Hook_config) SetUrl(value *string)() {
    m.url = value
}
// Hook_configable 
type Hook_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(*string)
    GetDigest()(*string)
    GetEmail()(*string)
    GetInsecureSsl()(WebhookConfigInsecureSslable)
    GetPassword()(*string)
    GetRoom()(*string)
    GetSecret()(*string)
    GetSubdomain()(*string)
    GetToken()(*string)
    GetUrl()(*string)
    SetContentType(value *string)()
    SetDigest(value *string)()
    SetEmail(value *string)()
    SetInsecureSsl(value WebhookConfigInsecureSslable)()
    SetPassword(value *string)()
    SetRoom(value *string)()
    SetSecret(value *string)()
    SetSubdomain(value *string)()
    SetToken(value *string)()
    SetUrl(value *string)()
}
