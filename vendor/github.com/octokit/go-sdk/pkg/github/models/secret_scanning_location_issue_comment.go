package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationIssueComment represents an 'issue_comment' secret scanning location type. This location type shows that a secret was detected in a comment on an issue.
type SecretScanningLocationIssueComment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the issue comment where the secret was detected.
    issue_comment_url *string
}
// NewSecretScanningLocationIssueComment instantiates a new SecretScanningLocationIssueComment and sets the default values.
func NewSecretScanningLocationIssueComment()(*SecretScanningLocationIssueComment) {
    m := &SecretScanningLocationIssueComment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationIssueCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationIssueCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationIssueComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationIssueComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationIssueComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["issue_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueCommentUrl(val)
        }
        return nil
    }
    return res
}
// GetIssueCommentUrl gets the issue_comment_url property value. The API URL to get the issue comment where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationIssueComment) GetIssueCommentUrl()(*string) {
    return m.issue_comment_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationIssueComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("issue_comment_url", m.GetIssueCommentUrl())
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
func (m *SecretScanningLocationIssueComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIssueCommentUrl sets the issue_comment_url property value. The API URL to get the issue comment where the secret was detected.
func (m *SecretScanningLocationIssueComment) SetIssueCommentUrl(value *string)() {
    m.issue_comment_url = value
}
type SecretScanningLocationIssueCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssueCommentUrl()(*string)
    SetIssueCommentUrl(value *string)()
}
