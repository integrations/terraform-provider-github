package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationDiscussionBody represents a 'discussion_body' secret scanning location type. This location type shows that a secret was detected in the body of a discussion.
type SecretScanningLocationDiscussionBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL to the discussion where the secret was detected.
    discussion_body_url *string
}
// NewSecretScanningLocationDiscussionBody instantiates a new SecretScanningLocationDiscussionBody and sets the default values.
func NewSecretScanningLocationDiscussionBody()(*SecretScanningLocationDiscussionBody) {
    m := &SecretScanningLocationDiscussionBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationDiscussionBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationDiscussionBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationDiscussionBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationDiscussionBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDiscussionBodyUrl gets the discussion_body_url property value. The URL to the discussion where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationDiscussionBody) GetDiscussionBodyUrl()(*string) {
    return m.discussion_body_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationDiscussionBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["discussion_body_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscussionBodyUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SecretScanningLocationDiscussionBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("discussion_body_url", m.GetDiscussionBodyUrl())
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
func (m *SecretScanningLocationDiscussionBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDiscussionBodyUrl sets the discussion_body_url property value. The URL to the discussion where the secret was detected.
func (m *SecretScanningLocationDiscussionBody) SetDiscussionBodyUrl(value *string)() {
    m.discussion_body_url = value
}
type SecretScanningLocationDiscussionBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDiscussionBodyUrl()(*string)
    SetDiscussionBodyUrl(value *string)()
}
