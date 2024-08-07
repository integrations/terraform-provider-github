package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GitCommit_verification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The payload property
    payload *string
    // The reason property
    reason *string
    // The signature property
    signature *string
    // The verified property
    verified *bool
}
// NewGitCommit_verification instantiates a new GitCommit_verification and sets the default values.
func NewGitCommit_verification()(*GitCommit_verification) {
    m := &GitCommit_verification{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitCommit_verificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitCommit_verificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitCommit_verification(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GitCommit_verification) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitCommit_verification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val)
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
// GetPayload gets the payload property value. The payload property
// returns a *string when successful
func (m *GitCommit_verification) GetPayload()(*string) {
    return m.payload
}
// GetReason gets the reason property value. The reason property
// returns a *string when successful
func (m *GitCommit_verification) GetReason()(*string) {
    return m.reason
}
// GetSignature gets the signature property value. The signature property
// returns a *string when successful
func (m *GitCommit_verification) GetSignature()(*string) {
    return m.signature
}
// GetVerified gets the verified property value. The verified property
// returns a *bool when successful
func (m *GitCommit_verification) GetVerified()(*bool) {
    return m.verified
}
// Serialize serializes information the current object
func (m *GitCommit_verification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signature", m.GetSignature())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GitCommit_verification) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPayload sets the payload property value. The payload property
func (m *GitCommit_verification) SetPayload(value *string)() {
    m.payload = value
}
// SetReason sets the reason property value. The reason property
func (m *GitCommit_verification) SetReason(value *string)() {
    m.reason = value
}
// SetSignature sets the signature property value. The signature property
func (m *GitCommit_verification) SetSignature(value *string)() {
    m.signature = value
}
// SetVerified sets the verified property value. The verified property
func (m *GitCommit_verification) SetVerified(value *bool)() {
    m.verified = value
}
type GitCommit_verificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPayload()(*string)
    GetReason()(*string)
    GetSignature()(*string)
    GetVerified()(*bool)
    SetPayload(value *string)()
    SetReason(value *string)()
    SetSignature(value *string)()
    SetVerified(value *bool)()
}
