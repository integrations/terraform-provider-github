package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationPullRequestReviewComment represents a 'pull_request_review_comment' secret scanning location type. This location type shows that a secret was detected in a review comment on a pull request.
type SecretScanningLocationPullRequestReviewComment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the pull request review comment where the secret was detected.
    pull_request_review_comment_url *string
}
// NewSecretScanningLocationPullRequestReviewComment instantiates a new SecretScanningLocationPullRequestReviewComment and sets the default values.
func NewSecretScanningLocationPullRequestReviewComment()(*SecretScanningLocationPullRequestReviewComment) {
    m := &SecretScanningLocationPullRequestReviewComment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationPullRequestReviewCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationPullRequestReviewCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationPullRequestReviewComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationPullRequestReviewComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationPullRequestReviewComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pull_request_review_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestReviewCommentUrl(val)
        }
        return nil
    }
    return res
}
// GetPullRequestReviewCommentUrl gets the pull_request_review_comment_url property value. The API URL to get the pull request review comment where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationPullRequestReviewComment) GetPullRequestReviewCommentUrl()(*string) {
    return m.pull_request_review_comment_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationPullRequestReviewComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pull_request_review_comment_url", m.GetPullRequestReviewCommentUrl())
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
func (m *SecretScanningLocationPullRequestReviewComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullRequestReviewCommentUrl sets the pull_request_review_comment_url property value. The API URL to get the pull request review comment where the secret was detected.
func (m *SecretScanningLocationPullRequestReviewComment) SetPullRequestReviewCommentUrl(value *string)() {
    m.pull_request_review_comment_url = value
}
type SecretScanningLocationPullRequestReviewCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullRequestReviewCommentUrl()(*string)
    SetPullRequestReviewCommentUrl(value *string)()
}
