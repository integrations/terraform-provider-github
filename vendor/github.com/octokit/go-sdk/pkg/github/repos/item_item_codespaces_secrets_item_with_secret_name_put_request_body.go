package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get a repository public key](https://docs.github.com/rest/codespaces/repository-secrets#get-a-repository-public-key) endpoint.
    encrypted_value *string
    // ID of the key you used to encrypt the secret.
    key_id *string
}
// NewItemItemCodespacesSecretsItemWithSecret_namePutRequestBody instantiates a new ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody and sets the default values.
func NewItemItemCodespacesSecretsItemWithSecret_namePutRequestBody()(*ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) {
    m := &ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesSecretsItemWithSecret_namePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodespacesSecretsItemWithSecret_namePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesSecretsItemWithSecret_namePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncryptedValue gets the encrypted_value property value. Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get a repository public key](https://docs.github.com/rest/codespaces/repository-secrets#get-a-repository-public-key) endpoint.
// returns a *string when successful
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) GetEncryptedValue()(*string) {
    return m.encrypted_value
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encrypted_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedValue(val)
        }
        return nil
    }
    res["key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the key_id property value. ID of the key you used to encrypt the secret.
// returns a *string when successful
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) GetKeyId()(*string) {
    return m.key_id
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encrypted_value", m.GetEncryptedValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_id", m.GetKeyId())
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
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncryptedValue sets the encrypted_value property value. Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get a repository public key](https://docs.github.com/rest/codespaces/repository-secrets#get-a-repository-public-key) endpoint.
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) SetEncryptedValue(value *string)() {
    m.encrypted_value = value
}
// SetKeyId sets the key_id property value. ID of the key you used to encrypt the secret.
func (m *ItemItemCodespacesSecretsItemWithSecret_namePutRequestBody) SetKeyId(value *string)() {
    m.key_id = value
}
type ItemItemCodespacesSecretsItemWithSecret_namePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedValue()(*string)
    GetKeyId()(*string)
    SetEncryptedValue(value *string)()
    SetKeyId(value *string)()
}
