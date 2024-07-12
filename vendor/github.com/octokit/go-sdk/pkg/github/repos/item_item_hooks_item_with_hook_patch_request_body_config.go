package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemHooksItemWithHook_PatchRequestBody_config key/value pairs to provide settings for this webhook.
type ItemItemHooksItemWithHook_PatchRequestBody_config struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address property
    address *string
    // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
    content_type *string
    // The insecure_ssl property
    insecure_ssl i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigInsecureSslable
    // The room property
    room *string
    // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
    secret *string
    // The URL to which the payloads will be delivered.
    url *string
}
// NewItemItemHooksItemWithHook_PatchRequestBody_config instantiates a new ItemItemHooksItemWithHook_PatchRequestBody_config and sets the default values.
func NewItemItemHooksItemWithHook_PatchRequestBody_config()(*ItemItemHooksItemWithHook_PatchRequestBody_config) {
    m := &ItemItemHooksItemWithHook_PatchRequestBody_config{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemHooksItemWithHook_PatchRequestBody_configFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemHooksItemWithHook_PatchRequestBody_configFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemHooksItemWithHook_PatchRequestBody_config(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The address property
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetAddress()(*string) {
    return m.address
}
// GetContentType gets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetContentType()(*string) {
    return m.content_type
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
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
        val, err := n.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateWebhookConfigInsecureSslFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsecureSsl(val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigInsecureSslable))
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
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetInsecureSsl()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigInsecureSslable) {
    return m.insecure_ssl
}
// GetRoom gets the room property value. The room property
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetRoom()(*string) {
    return m.room
}
// GetSecret gets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetSecret()(*string) {
    return m.secret
}
// GetUrl gets the url property value. The URL to which the payloads will be delivered.
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
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
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The address property
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetAddress(value *string)() {
    m.address = value
}
// SetContentType sets the content_type property value. The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetContentType(value *string)() {
    m.content_type = value
}
// SetInsecureSsl sets the insecure_ssl property value. The insecure_ssl property
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetInsecureSsl(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigInsecureSslable)() {
    m.insecure_ssl = value
}
// SetRoom sets the room property value. The room property
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetRoom(value *string)() {
    m.room = value
}
// SetSecret sets the secret property value. If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetSecret(value *string)() {
    m.secret = value
}
// SetUrl sets the url property value. The URL to which the payloads will be delivered.
func (m *ItemItemHooksItemWithHook_PatchRequestBody_config) SetUrl(value *string)() {
    m.url = value
}
// ItemItemHooksItemWithHook_PatchRequestBody_configable 
type ItemItemHooksItemWithHook_PatchRequestBody_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetContentType()(*string)
    GetInsecureSsl()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigInsecureSslable)
    GetRoom()(*string)
    GetSecret()(*string)
    GetUrl()(*string)
    SetAddress(value *string)()
    SetContentType(value *string)()
    SetInsecureSsl(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WebhookConfigInsecureSslable)()
    SetRoom(value *string)()
    SetSecret(value *string)()
    SetUrl(value *string)()
}
