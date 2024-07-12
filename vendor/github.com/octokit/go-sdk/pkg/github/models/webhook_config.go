package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookConfig configuration object of the webhook
type WebhookConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
    content_type *string
    // The insecure_ssl property
    insecure_ssl WebhookConfigInsecureSslable
    // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
    secret *string
    // The URL to which the payloads will be delivered.
    url *string
}
// NewWebhookConfig instantiates a new WebhookConfig and sets the default values.
func NewWebhookConfig()(*WebhookConfig) {
    m := &WebhookConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebhookConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebhookConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WebhookConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentType gets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
// returns a *string when successful
func (m *WebhookConfig) GetContentType()(*string) {
    return m.content_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebhookConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// returns a WebhookConfigInsecureSslable when successful
func (m *WebhookConfig) GetInsecureSsl()(WebhookConfigInsecureSslable) {
    return m.insecure_ssl
}
// GetSecret gets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
// returns a *string when successful
func (m *WebhookConfig) GetSecret()(*string) {
    return m.secret
}
// GetUrl gets the url property value. The URL to which the payloads will be delivered.
// returns a *string when successful
func (m *WebhookConfig) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *WebhookConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
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
        err := writer.WriteStringValue("secret", m.GetSecret())
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
func (m *WebhookConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentType sets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *WebhookConfig) SetContentType(value *string)() {
    m.content_type = value
}
// SetInsecureSsl sets the insecure_ssl property value. The insecure_ssl property
func (m *WebhookConfig) SetInsecureSsl(value WebhookConfigInsecureSslable)() {
    m.insecure_ssl = value
}
// SetSecret sets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *WebhookConfig) SetSecret(value *string)() {
    m.secret = value
}
// SetUrl sets the url property value. The URL to which the payloads will be delivered.
func (m *WebhookConfig) SetUrl(value *string)() {
    m.url = value
}
type WebhookConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(*string)
    GetInsecureSsl()(WebhookConfigInsecureSslable)
    GetSecret()(*string)
    GetUrl()(*string)
    SetContentType(value *string)()
    SetInsecureSsl(value WebhookConfigInsecureSslable)()
    SetSecret(value *string)()
    SetUrl(value *string)()
}
