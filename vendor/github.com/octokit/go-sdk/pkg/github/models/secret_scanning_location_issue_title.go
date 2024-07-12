package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationIssueTitle represents an 'issue_title' secret scanning location type. This location type shows that a secret was detected in the title of an issue.
type SecretScanningLocationIssueTitle struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The API URL to get the issue where the secret was detected.
    issue_title_url *string
}
// NewSecretScanningLocationIssueTitle instantiates a new SecretScanningLocationIssueTitle and sets the default values.
func NewSecretScanningLocationIssueTitle()(*SecretScanningLocationIssueTitle) {
    m := &SecretScanningLocationIssueTitle{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationIssueTitleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationIssueTitleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationIssueTitle(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationIssueTitle) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationIssueTitle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["issue_title_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueTitleUrl(val)
        }
        return nil
    }
    return res
}
// GetIssueTitleUrl gets the issue_title_url property value. The API URL to get the issue where the secret was detected.
// returns a *string when successful
func (m *SecretScanningLocationIssueTitle) GetIssueTitleUrl()(*string) {
    return m.issue_title_url
}
// Serialize serializes information the current object
func (m *SecretScanningLocationIssueTitle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("issue_title_url", m.GetIssueTitleUrl())
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
func (m *SecretScanningLocationIssueTitle) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIssueTitleUrl sets the issue_title_url property value. The API URL to get the issue where the secret was detected.
func (m *SecretScanningLocationIssueTitle) SetIssueTitleUrl(value *string)() {
    m.issue_title_url = value
}
type SecretScanningLocationIssueTitleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssueTitleUrl()(*string)
    SetIssueTitleUrl(value *string)()
}
