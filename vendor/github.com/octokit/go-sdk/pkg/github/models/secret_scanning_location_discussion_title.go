package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationDiscussionTitle represents a 'discussion_title' secret scanning location type. This location type shows that a secret was detected in the title of a discussion.
type SecretScanningLocationDiscussionTitle struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL to the discussion where the secret was detected.
    discussion_title_url *string
}
// NewSecretScanningLocationDiscussionTitle instantiates a new SecretScanningLocationDiscussionTitle and sets the default values.
func NewSecretScanningLocationDiscussionTitle()(*SecretScanningLocationDiscussionTitle) {
    m := &SecretScanningLocationDiscussionTitle{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationDiscussionTitleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationDiscussionTitleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationDiscussionTitle(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationDiscussionTitle) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDiscussionTitleUrl gets the discussion_title_url property value. The URL to the discussion where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationDiscussionTitle) GetDiscussionTitleUrl()(*string) {
    return m.discussion_title_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationDiscussionTitle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["discussion_title_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscussionTitleUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SecretScanningLocationDiscussionTitle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("discussion_title_url", m.GetDiscussionTitleUrl())
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
func (m *SecretScanningLocationDiscussionTitle) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDiscussionTitleUrl sets the discussion_title_url property value. The URL to the discussion where the secret was detected.
func (m *SecretScanningLocationDiscussionTitle) SetDiscussionTitleUrl(value *string)() {
    m.discussion_title_url = value
}
type SecretScanningLocationDiscussionTitleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDiscussionTitleUrl()(*string)
    SetDiscussionTitleUrl(value *string)()
}
