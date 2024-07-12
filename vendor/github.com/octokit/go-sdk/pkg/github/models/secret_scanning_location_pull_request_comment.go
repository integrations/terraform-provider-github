package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationPullRequestComment represents a 'pull_request_comment' secret scanning location type. This location type shows that a secret was detected in a comment on a pull request.
type SecretScanningLocationPullRequestComment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the pull request comment where the secret was detected.
    pull_request_comment_url *string
}
// NewSecretScanningLocationPullRequestComment instantiates a new SecretScanningLocationPullRequestComment and sets the default values.
func NewSecretScanningLocationPullRequestComment()(*SecretScanningLocationPullRequestComment) {
    m := &SecretScanningLocationPullRequestComment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationPullRequestCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationPullRequestCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationPullRequestComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationPullRequestComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationPullRequestComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pull_request_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestCommentUrl(val)
        }
        return nil
    }
    return res
}
// GetPullRequestCommentUrl gets the pull_request_comment_url property value. The API URL to get the pull request comment where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationPullRequestComment) GetPullRequestCommentUrl()(*string) {
    return m.pull_request_comment_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationPullRequestComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pull_request_comment_url", m.GetPullRequestCommentUrl())
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
func (m *SecretScanningLocationPullRequestComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullRequestCommentUrl sets the pull_request_comment_url property value. The API URL to get the pull request comment where the secret was detected.
func (m *SecretScanningLocationPullRequestComment) SetPullRequestCommentUrl(value *string)() {
    m.pull_request_comment_url = value
}
type SecretScanningLocationPullRequestCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullRequestCommentUrl()(*string)
    SetPullRequestCommentUrl(value *string)()
}
