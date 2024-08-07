package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationIssueBody represents an 'issue_body' secret scanning location type. This location type shows that a secret was detected in the body of an issue.
type SecretScanningLocationIssueBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the issue where the secret was detected.
    issue_body_url *string
}
// NewSecretScanningLocationIssueBody instantiates a new SecretScanningLocationIssueBody and sets the default values.
func NewSecretScanningLocationIssueBody()(*SecretScanningLocationIssueBody) {
    m := &SecretScanningLocationIssueBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationIssueBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationIssueBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationIssueBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationIssueBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationIssueBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["issue_body_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueBodyUrl(val)
        }
        return nil
    }
    return res
}
// GetIssueBodyUrl gets the issue_body_url property value. The API URL to get the issue where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationIssueBody) GetIssueBodyUrl()(*string) {
    return m.issue_body_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationIssueBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("issue_body_url", m.GetIssueBodyUrl())
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
func (m *SecretScanningLocationIssueBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIssueBodyUrl sets the issue_body_url property value. The API URL to get the issue where the secret was detected.
func (m *SecretScanningLocationIssueBody) SetIssueBodyUrl(value *string)() {
    m.issue_body_url = value
}
type SecretScanningLocationIssueBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssueBodyUrl()(*string)
    SetIssueBodyUrl(value *string)()
}
