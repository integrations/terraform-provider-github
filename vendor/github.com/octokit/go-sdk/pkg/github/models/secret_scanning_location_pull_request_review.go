package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationPullRequestReview represents a 'pull_request_review' secret scanning location type. This location type shows that a secret was detected in a review on a pull request.
type SecretScanningLocationPullRequestReview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the pull request review where the secret was detected.
    pull_request_review_url *string
}
// NewSecretScanningLocationPullRequestReview instantiates a new SecretScanningLocationPullRequestReview and sets the default values.
func NewSecretScanningLocationPullRequestReview()(*SecretScanningLocationPullRequestReview) {
    m := &SecretScanningLocationPullRequestReview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationPullRequestReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationPullRequestReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationPullRequestReview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationPullRequestReview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationPullRequestReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pull_request_review_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestReviewUrl(val)
        }
        return nil
    }
    return res
}
// GetPullRequestReviewUrl gets the pull_request_review_url property value. The API URL to get the pull request review where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationPullRequestReview) GetPullRequestReviewUrl()(*string) {
    return m.pull_request_review_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationPullRequestReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pull_request_review_url", m.GetPullRequestReviewUrl())
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
func (m *SecretScanningLocationPullRequestReview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullRequestReviewUrl sets the pull_request_review_url property value. The API URL to get the pull request review where the secret was detected.
func (m *SecretScanningLocationPullRequestReview) SetPullRequestReviewUrl(value *string)() {
    m.pull_request_review_url = value
}
type SecretScanningLocationPullRequestReviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullRequestReviewUrl()(*string)
    SetPullRequestReviewUrl(value *string)()
}
