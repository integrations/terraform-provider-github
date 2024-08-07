package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunnerApplication runner Application
type RunnerApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The architecture property
    architecture *string
    // The download_url property
    download_url *string
    // The filename property
    filename *string
    // The os property
    os *string
    // The sha256_checksum property
    sha256_checksum *string
    // A short lived bearer token used to download the runner, if needed.
    temp_download_token *string
}
// NewRunnerApplication instantiates a new RunnerApplication and sets the default values.
func NewRunnerApplication()(*RunnerApplication) {
    m := &RunnerApplication{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRunnerApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunnerApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunnerApplication(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RunnerApplication) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchitecture gets the architecture property value. The architecture property
// returns a *string when successful
func (m *RunnerApplication) GetArchitecture()(*string) {
    return m.architecture
}
// GetDownloadUrl gets the download_url property value. The download_url property
// returns a *string when successful
func (m *RunnerApplication) GetDownloadUrl()(*string) {
    return m.download_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RunnerApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["architecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchitecture(val)
        }
        return nil
    }
    res["download_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUrl(val)
        }
        return nil
    }
    res["filename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilename(val)
        }
        return nil
    }
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["sha256_checksum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256Checksum(val)
        }
        return nil
    }
    res["temp_download_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTempDownloadToken(val)
        }
        return nil
    }
    return res
}
// GetFilename gets the filename property value. The filename property
// returns a *string when successful
func (m *RunnerApplication) GetFilename()(*string) {
    return m.filename
}
// GetOs gets the os property value. The os property
// returns a *string when successful
func (m *RunnerApplication) GetOs()(*string) {
    return m.os
}
// GetSha256Checksum gets the sha256_checksum property value. The sha256_checksum property
// returns a *string when successful
func (m *RunnerApplication) GetSha256Checksum()(*string) {
    return m.sha256_checksum
}
// GetTempDownloadToken gets the temp_download_token property value. A short lived bearer token used to download the runner, if needed.
// returns a *string when successful
func (m *RunnerApplication) GetTempDownloadToken()(*string) {
    return m.temp_download_token
}
// Serialize serializes information the current object
func (m *RunnerApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("architecture", m.GetArchitecture())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("download_url", m.GetDownloadUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filename", m.GetFilename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha256_checksum", m.GetSha256Checksum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("temp_download_token", m.GetTempDownloadToken())
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
func (m *RunnerApplication) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchitecture sets the architecture property value. The architecture property
func (m *RunnerApplication) SetArchitecture(value *string)() {
    m.architecture = value
}
// SetDownloadUrl sets the download_url property value. The download_url property
func (m *RunnerApplication) SetDownloadUrl(value *string)() {
    m.download_url = value
}
// SetFilename sets the filename property value. The filename property
func (m *RunnerApplication) SetFilename(value *string)() {
    m.filename = value
}
// SetOs sets the os property value. The os property
func (m *RunnerApplication) SetOs(value *string)() {
    m.os = value
}
// SetSha256Checksum sets the sha256_checksum property value. The sha256_checksum property
func (m *RunnerApplication) SetSha256Checksum(value *string)() {
    m.sha256_checksum = value
}
// SetTempDownloadToken sets the temp_download_token property value. A short lived bearer token used to download the runner, if needed.
func (m *RunnerApplication) SetTempDownloadToken(value *string)() {
    m.temp_download_token = value
}
type RunnerApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchitecture()(*string)
    GetDownloadUrl()(*string)
    GetFilename()(*string)
    GetOs()(*string)
    GetSha256Checksum()(*string)
    GetTempDownloadToken()(*string)
    SetArchitecture(value *string)()
    SetDownloadUrl(value *string)()
    SetFilename(value *string)()
    SetOs(value *string)()
    SetSha256Checksum(value *string)()
    SetTempDownloadToken(value *string)()
}
