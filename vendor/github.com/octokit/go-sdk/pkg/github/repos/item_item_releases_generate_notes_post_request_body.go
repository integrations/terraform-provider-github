package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemReleasesGenerateNotesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies a path to a file in the repository containing configuration settings used for generating the release notes. If unspecified, the configuration file located in the repository at '.github/release.yml' or '.github/release.yaml' will be used. If that is not present, the default configuration will be used.
    configuration_file_path *string
    // The name of the previous tag to use as the starting point for the release notes. Use to manually specify the range for the set of changes considered as part this release.
    previous_tag_name *string
    // The tag name for the release. This can be an existing tag or a new one.
    tag_name *string
    // Specifies the commitish value that will be the target for the release's tag. Required if the supplied tag_name does not reference an existing tag. Ignored if the tag_name already exists.
    target_commitish *string
}
// NewItemItemReleasesGenerateNotesPostRequestBody instantiates a new ItemItemReleasesGenerateNotesPostRequestBody and sets the default values.
func NewItemItemReleasesGenerateNotesPostRequestBody()(*ItemItemReleasesGenerateNotesPostRequestBody) {
    m := &ItemItemReleasesGenerateNotesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemReleasesGenerateNotesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemReleasesGenerateNotesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemReleasesGenerateNotesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemReleasesGenerateNotesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigurationFilePath gets the configuration_file_path property value. Specifies a path to a file in the repository containing configuration settings used for generating the release notes. If unspecified, the configuration file located in the repository at '.github/release.yml' or '.github/release.yaml' will be used. If that is not present, the default configuration will be used.
// returns a *string when successful
func (m *ItemItemReleasesGenerateNotesPostRequestBody) GetConfigurationFilePath()(*string) {
    return m.configuration_file_path
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemReleasesGenerateNotesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuration_file_path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationFilePath(val)
        }
        return nil
    }
    res["previous_tag_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousTagName(val)
        }
        return nil
    }
    res["tag_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagName(val)
        }
        return nil
    }
    res["target_commitish"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetCommitish(val)
        }
        return nil
    }
    return res
}
// GetPreviousTagName gets the previous_tag_name property value. The name of the previous tag to use as the starting point for the release notes. Use to manually specify the range for the set of changes considered as part this release.
// returns a *string when successful
func (m *ItemItemReleasesGenerateNotesPostRequestBody) GetPreviousTagName()(*string) {
    return m.previous_tag_name
}
// GetTagName gets the tag_name property value. The tag name for the release. This can be an existing tag or a new one.
// returns a *string when successful
func (m *ItemItemReleasesGenerateNotesPostRequestBody) GetTagName()(*string) {
    return m.tag_name
}
// GetTargetCommitish gets the target_commitish property value. Specifies the commitish value that will be the target for the release's tag. Required if the supplied tag_name does not reference an existing tag. Ignored if the tag_name already exists.
// returns a *string when successful
func (m *ItemItemReleasesGenerateNotesPostRequestBody) GetTargetCommitish()(*string) {
    return m.target_commitish
}
// Serialize serializes information the current object
func (m *ItemItemReleasesGenerateNotesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("configuration_file_path", m.GetConfigurationFilePath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previous_tag_name", m.GetPreviousTagName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tag_name", m.GetTagName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_commitish", m.GetTargetCommitish())
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
func (m *ItemItemReleasesGenerateNotesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigurationFilePath sets the configuration_file_path property value. Specifies a path to a file in the repository containing configuration settings used for generating the release notes. If unspecified, the configuration file located in the repository at '.github/release.yml' or '.github/release.yaml' will be used. If that is not present, the default configuration will be used.
func (m *ItemItemReleasesGenerateNotesPostRequestBody) SetConfigurationFilePath(value *string)() {
    m.configuration_file_path = value
}
// SetPreviousTagName sets the previous_tag_name property value. The name of the previous tag to use as the starting point for the release notes. Use to manually specify the range for the set of changes considered as part this release.
func (m *ItemItemReleasesGenerateNotesPostRequestBody) SetPreviousTagName(value *string)() {
    m.previous_tag_name = value
}
// SetTagName sets the tag_name property value. The tag name for the release. This can be an existing tag or a new one.
func (m *ItemItemReleasesGenerateNotesPostRequestBody) SetTagName(value *string)() {
    m.tag_name = value
}
// SetTargetCommitish sets the target_commitish property value. Specifies the commitish value that will be the target for the release's tag. Required if the supplied tag_name does not reference an existing tag. Ignored if the tag_name already exists.
func (m *ItemItemReleasesGenerateNotesPostRequestBody) SetTargetCommitish(value *string)() {
    m.target_commitish = value
}
type ItemItemReleasesGenerateNotesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationFilePath()(*string)
    GetPreviousTagName()(*string)
    GetTagName()(*string)
    GetTargetCommitish()(*string)
    SetConfigurationFilePath(value *string)()
    SetPreviousTagName(value *string)()
    SetTagName(value *string)()
    SetTargetCommitish(value *string)()
}
