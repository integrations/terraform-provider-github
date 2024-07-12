package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiOverview_ssh_key_fingerprints struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The SHA256_DSA property
    sHA256_DSA *string
    // The SHA256_ECDSA property
    sHA256_ECDSA *string
    // The SHA256_ED25519 property
    sHA256_ED25519 *string
    // The SHA256_RSA property
    sHA256_RSA *string
}
// NewApiOverview_ssh_key_fingerprints instantiates a new ApiOverview_ssh_key_fingerprints and sets the default values.
func NewApiOverview_ssh_key_fingerprints()(*ApiOverview_ssh_key_fingerprints) {
    m := &ApiOverview_ssh_key_fingerprints{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiOverview_ssh_key_fingerprintsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiOverview_ssh_key_fingerprintsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiOverview_ssh_key_fingerprints(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiOverview_ssh_key_fingerprints) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiOverview_ssh_key_fingerprints) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["SHA256_DSA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSHA256DSA(val)
        }
        return nil
    }
    res["SHA256_ECDSA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSHA256ECDSA(val)
        }
        return nil
    }
    res["SHA256_ED25519"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSHA256ED25519(val)
        }
        return nil
    }
    res["SHA256_RSA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSHA256RSA(val)
        }
        return nil
    }
    return res
}
// GetSHA256DSA gets the SHA256_DSA property value. The SHA256_DSA property
// returns a *string when successful
func (m *ApiOverview_ssh_key_fingerprints) GetSHA256DSA()(*string) {
    return m.sHA256_DSA
}
// GetSHA256ECDSA gets the SHA256_ECDSA property value. The SHA256_ECDSA property
// returns a *string when successful
func (m *ApiOverview_ssh_key_fingerprints) GetSHA256ECDSA()(*string) {
    return m.sHA256_ECDSA
}
// GetSHA256ED25519 gets the SHA256_ED25519 property value. The SHA256_ED25519 property
// returns a *string when successful
func (m *ApiOverview_ssh_key_fingerprints) GetSHA256ED25519()(*string) {
    return m.sHA256_ED25519
}
// GetSHA256RSA gets the SHA256_RSA property value. The SHA256_RSA property
// returns a *string when successful
func (m *ApiOverview_ssh_key_fingerprints) GetSHA256RSA()(*string) {
    return m.sHA256_RSA
}
// Serialize serializes information the current object
func (m *ApiOverview_ssh_key_fingerprints) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("SHA256_DSA", m.GetSHA256DSA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("SHA256_ECDSA", m.GetSHA256ECDSA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("SHA256_ED25519", m.GetSHA256ED25519())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("SHA256_RSA", m.GetSHA256RSA())
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
func (m *ApiOverview_ssh_key_fingerprints) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSHA256DSA sets the SHA256_DSA property value. The SHA256_DSA property
func (m *ApiOverview_ssh_key_fingerprints) SetSHA256DSA(value *string)() {
    m.sHA256_DSA = value
}
// SetSHA256ECDSA sets the SHA256_ECDSA property value. The SHA256_ECDSA property
func (m *ApiOverview_ssh_key_fingerprints) SetSHA256ECDSA(value *string)() {
    m.sHA256_ECDSA = value
}
// SetSHA256ED25519 sets the SHA256_ED25519 property value. The SHA256_ED25519 property
func (m *ApiOverview_ssh_key_fingerprints) SetSHA256ED25519(value *string)() {
    m.sHA256_ED25519 = value
}
// SetSHA256RSA sets the SHA256_RSA property value. The SHA256_RSA property
func (m *ApiOverview_ssh_key_fingerprints) SetSHA256RSA(value *string)() {
    m.sHA256_RSA = value
}
type ApiOverview_ssh_key_fingerprintsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSHA256DSA()(*string)
    GetSHA256ECDSA()(*string)
    GetSHA256ED25519()(*string)
    GetSHA256RSA()(*string)
    SetSHA256DSA(value *string)()
    SetSHA256ECDSA(value *string)()
    SetSHA256ED25519(value *string)()
    SetSHA256RSA(value *string)()
}
