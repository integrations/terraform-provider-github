package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommunityProfile_files struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Code of Conduct Simple
    code_of_conduct NullableCodeOfConductSimpleable
    // The code_of_conduct_file property
    code_of_conduct_file NullableCommunityHealthFileable
    // The contributing property
    contributing NullableCommunityHealthFileable
    // The issue_template property
    issue_template NullableCommunityHealthFileable
    // License Simple
    license NullableLicenseSimpleable
    // The pull_request_template property
    pull_request_template NullableCommunityHealthFileable
    // The readme property
    readme NullableCommunityHealthFileable
}
// NewCommunityProfile_files instantiates a new CommunityProfile_files and sets the default values.
func NewCommunityProfile_files()(*CommunityProfile_files) {
    m := &CommunityProfile_files{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommunityProfile_filesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommunityProfile_filesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunityProfile_files(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommunityProfile_files) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodeOfConduct gets the code_of_conduct property value. Code of Conduct Simple
// returns a NullableCodeOfConductSimpleable when successful
func (m *CommunityProfile_files) GetCodeOfConduct()(NullableCodeOfConductSimpleable) {
    return m.code_of_conduct
}
// GetCodeOfConductFile gets the code_of_conduct_file property value. The code_of_conduct_file property
// returns a NullableCommunityHealthFileable when successful
func (m *CommunityProfile_files) GetCodeOfConductFile()(NullableCommunityHealthFileable) {
    return m.code_of_conduct_file
}
// GetContributing gets the contributing property value. The contributing property
// returns a NullableCommunityHealthFileable when successful
func (m *CommunityProfile_files) GetContributing()(NullableCommunityHealthFileable) {
    return m.contributing
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommunityProfile_files) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code_of_conduct"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCodeOfConductSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeOfConduct(val.(NullableCodeOfConductSimpleable))
        }
        return nil
    }
    res["code_of_conduct_file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCommunityHealthFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeOfConductFile(val.(NullableCommunityHealthFileable))
        }
        return nil
    }
    res["contributing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCommunityHealthFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributing(val.(NullableCommunityHealthFileable))
        }
        return nil
    }
    res["issue_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCommunityHealthFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueTemplate(val.(NullableCommunityHealthFileable))
        }
        return nil
    }
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableLicenseSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val.(NullableLicenseSimpleable))
        }
        return nil
    }
    res["pull_request_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCommunityHealthFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestTemplate(val.(NullableCommunityHealthFileable))
        }
        return nil
    }
    res["readme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCommunityHealthFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadme(val.(NullableCommunityHealthFileable))
        }
        return nil
    }
    return res
}
// GetIssueTemplate gets the issue_template property value. The issue_template property
// returns a NullableCommunityHealthFileable when successful
func (m *CommunityProfile_files) GetIssueTemplate()(NullableCommunityHealthFileable) {
    return m.issue_template
}
// GetLicense gets the license property value. License Simple
// returns a NullableLicenseSimpleable when successful
func (m *CommunityProfile_files) GetLicense()(NullableLicenseSimpleable) {
    return m.license
}
// GetPullRequestTemplate gets the pull_request_template property value. The pull_request_template property
// returns a NullableCommunityHealthFileable when successful
func (m *CommunityProfile_files) GetPullRequestTemplate()(NullableCommunityHealthFileable) {
    return m.pull_request_template
}
// GetReadme gets the readme property value. The readme property
// returns a NullableCommunityHealthFileable when successful
func (m *CommunityProfile_files) GetReadme()(NullableCommunityHealthFileable) {
    return m.readme
}
// Serialize serializes information the current object
func (m *CommunityProfile_files) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("code_of_conduct", m.GetCodeOfConduct())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("code_of_conduct_file", m.GetCodeOfConductFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contributing", m.GetContributing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("issue_template", m.GetIssueTemplate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pull_request_template", m.GetPullRequestTemplate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("readme", m.GetReadme())
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
func (m *CommunityProfile_files) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodeOfConduct sets the code_of_conduct property value. Code of Conduct Simple
func (m *CommunityProfile_files) SetCodeOfConduct(value NullableCodeOfConductSimpleable)() {
    m.code_of_conduct = value
}
// SetCodeOfConductFile sets the code_of_conduct_file property value. The code_of_conduct_file property
func (m *CommunityProfile_files) SetCodeOfConductFile(value NullableCommunityHealthFileable)() {
    m.code_of_conduct_file = value
}
// SetContributing sets the contributing property value. The contributing property
func (m *CommunityProfile_files) SetContributing(value NullableCommunityHealthFileable)() {
    m.contributing = value
}
// SetIssueTemplate sets the issue_template property value. The issue_template property
func (m *CommunityProfile_files) SetIssueTemplate(value NullableCommunityHealthFileable)() {
    m.issue_template = value
}
// SetLicense sets the license property value. License Simple
func (m *CommunityProfile_files) SetLicense(value NullableLicenseSimpleable)() {
    m.license = value
}
// SetPullRequestTemplate sets the pull_request_template property value. The pull_request_template property
func (m *CommunityProfile_files) SetPullRequestTemplate(value NullableCommunityHealthFileable)() {
    m.pull_request_template = value
}
// SetReadme sets the readme property value. The readme property
func (m *CommunityProfile_files) SetReadme(value NullableCommunityHealthFileable)() {
    m.readme = value
}
type CommunityProfile_filesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodeOfConduct()(NullableCodeOfConductSimpleable)
    GetCodeOfConductFile()(NullableCommunityHealthFileable)
    GetContributing()(NullableCommunityHealthFileable)
    GetIssueTemplate()(NullableCommunityHealthFileable)
    GetLicense()(NullableLicenseSimpleable)
    GetPullRequestTemplate()(NullableCommunityHealthFileable)
    GetReadme()(NullableCommunityHealthFileable)
    SetCodeOfConduct(value NullableCodeOfConductSimpleable)()
    SetCodeOfConductFile(value NullableCommunityHealthFileable)()
    SetContributing(value NullableCommunityHealthFileable)()
    SetIssueTemplate(value NullableCommunityHealthFileable)()
    SetLicense(value NullableLicenseSimpleable)()
    SetPullRequestTemplate(value NullableCommunityHealthFileable)()
    SetReadme(value NullableCommunityHealthFileable)()
}
