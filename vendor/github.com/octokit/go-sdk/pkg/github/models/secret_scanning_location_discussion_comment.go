package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationDiscussionComment represents a 'discussion_comment' secret scanning location type. This location type shows that a secret was detected in a comment on a discussion.
type SecretScanningLocationDiscussionComment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the discussion comment where the secret was detected.
    discussion_comment_url *string
}
// NewSecretScanningLocationDiscussionComment instantiates a new SecretScanningLocationDiscussionComment and sets the default values.
func NewSecretScanningLocationDiscussionComment()(*SecretScanningLocationDiscussionComment) {
    m := &SecretScanningLocationDiscussionComment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationDiscussionCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationDiscussionCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationDiscussionComment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationDiscussionComment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDiscussionCommentUrl gets the discussion_comment_url property value. The API URL to get the discussion comment where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationDiscussionComment) GetDiscussionCommentUrl()(*string) {
    return m.discussion_comment_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationDiscussionComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["discussion_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscussionCommentUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SecretScanningLocationDiscussionComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("discussion_comment_url", m.GetDiscussionCommentUrl())
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
func (m *SecretScanningLocationDiscussionComment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDiscussionCommentUrl sets the discussion_comment_url property value. The API URL to get the discussion comment where the secret was detected.
func (m *SecretScanningLocationDiscussionComment) SetDiscussionCommentUrl(value *string)() {
    m.discussion_comment_url = value
}
type SecretScanningLocationDiscussionCommentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDiscussionCommentUrl()(*string)
    SetDiscussionCommentUrl(value *string)()
}
