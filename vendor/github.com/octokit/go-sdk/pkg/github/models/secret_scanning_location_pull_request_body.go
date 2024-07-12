package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationPullRequestBody represents a 'pull_request_body' secret scanning location type. This location type shows that a secret was detected in the body of a pull request.
type SecretScanningLocationPullRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the pull request where the secret was detected.
    pull_request_body_url *string
}
// NewSecretScanningLocationPullRequestBody instantiates a new SecretScanningLocationPullRequestBody and sets the default values.
func NewSecretScanningLocationPullRequestBody()(*SecretScanningLocationPullRequestBody) {
    m := &SecretScanningLocationPullRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationPullRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationPullRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationPullRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationPullRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationPullRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pull_request_body_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestBodyUrl(val)
        }
        return nil
    }
    return res
}
// GetPullRequestBodyUrl gets the pull_request_body_url property value. The API URL to get the pull request where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationPullRequestBody) GetPullRequestBodyUrl()(*string) {
    return m.pull_request_body_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationPullRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pull_request_body_url", m.GetPullRequestBodyUrl())
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
func (m *SecretScanningLocationPullRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullRequestBodyUrl sets the pull_request_body_url property value. The API URL to get the pull request where the secret was detected.
func (m *SecretScanningLocationPullRequestBody) SetPullRequestBodyUrl(value *string)() {
    m.pull_request_body_url = value
}
type SecretScanningLocationPullRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullRequestBodyUrl()(*string)
    SetPullRequestBodyUrl(value *string)()
}
