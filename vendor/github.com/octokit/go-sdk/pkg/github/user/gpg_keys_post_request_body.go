package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Gpg_keysPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GPG key in ASCII-armored format.
    armored_public_key *string
    // A descriptive name for the new key.
    name *string
}
// NewGpg_keysPostRequestBody instantiates a new Gpg_keysPostRequestBody and sets the default values.
func NewGpg_keysPostRequestBody()(*Gpg_keysPostRequestBody) {
    m := &Gpg_keysPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGpg_keysPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGpg_keysPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGpg_keysPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Gpg_keysPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArmoredPublicKey gets the armored_public_key property value. A GPG key in ASCII-armored format.
// returns a *string when successful
func (m *Gpg_keysPostRequestBody) GetArmoredPublicKey()(*string) {
    return m.armored_public_key
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Gpg_keysPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["armored_public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArmoredPublicKey(val)
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
    return res
}
// GetName gets the name property value. A descriptive name for the new key.
// returns a *string when successful
func (m *Gpg_keysPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *Gpg_keysPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("armored_public_key", m.GetArmoredPublicKey())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Gpg_keysPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArmoredPublicKey sets the armored_public_key property value. A GPG key in ASCII-armored format.
func (m *Gpg_keysPostRequestBody) SetArmoredPublicKey(value *string)() {
    m.armored_public_key = value
}
// SetName sets the name property value. A descriptive name for the new key.
func (m *Gpg_keysPostRequestBody) SetName(value *string)() {
    m.name = value
}
type Gpg_keysPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArmoredPublicKey()(*string)
    GetName()(*string)
    SetArmoredPublicKey(value *string)()
    SetName(value *string)()
}
