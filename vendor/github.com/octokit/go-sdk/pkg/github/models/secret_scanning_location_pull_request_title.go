package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationPullRequestTitle represents a 'pull_request_title' secret scanning location type. This location type shows that a secret was detected in the title of a pull request.
type SecretScanningLocationPullRequestTitle struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the pull request where the secret was detected.
    pull_request_title_url *string
}
// NewSecretScanningLocationPullRequestTitle instantiates a new SecretScanningLocationPullRequestTitle and sets the default values.
func NewSecretScanningLocationPullRequestTitle()(*SecretScanningLocationPullRequestTitle) {
    m := &SecretScanningLocationPullRequestTitle{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationPullRequestTitleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationPullRequestTitleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationPullRequestTitle(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationPullRequestTitle) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationPullRequestTitle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pull_request_title_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestTitleUrl(val)
        }
        return nil
    }
    return res
}
// GetPullRequestTitleUrl gets the pull_request_title_url property value. The API URL to get the pull request where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationPullRequestTitle) GetPullRequestTitleUrl()(*string) {
    return m.pull_request_title_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationPullRequestTitle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pull_request_title_url", m.GetPullRequestTitleUrl())
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
func (m *SecretScanningLocationPullRequestTitle) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullRequestTitleUrl sets the pull_request_title_url property value. The API URL to get the pull request where the secret was detected.
func (m *SecretScanningLocationPullRequestTitle) SetPullRequestTitleUrl(value *string)() {
    m.pull_request_title_url = value
}
type SecretScanningLocationPullRequestTitleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullRequestTitleUrl()(*string)
    SetPullRequestTitleUrl(value *string)()
}
