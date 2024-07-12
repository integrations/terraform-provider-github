package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseContent license Content
type LicenseContent struct {
    // The _links property
    _links LicenseContent__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content *string
    // The download_url property
    download_url *string
    // The encoding property
    encoding *string
    // The git_url property
    git_url *string
    // The html_url property
    html_url *string
    // License Simple
    license NullableLicenseSimpleable
    // The name property
    name *string
    // The path property
    path *string
    // The sha property
    sha *string
    // The size property
    size *int32
    // The type property
    typeEscaped *string
    // The url property
    url *string
}
// NewLicenseContent instantiates a new LicenseContent and sets the default values.
func NewLicenseContent()(*LicenseContent) {
    m := &LicenseContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLicenseContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLicenseContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LicenseContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
// returns a *string when successful
func (m *LicenseContent) GetContent()(*string) {
    return m.content
}
// GetDownloadUrl gets the download_url property value. The download_url property
// returns a *string when successful
func (m *LicenseContent) GetDownloadUrl()(*string) {
    return m.download_url
}
// GetEncoding gets the encoding property value. The encoding property
// returns a *string when successful
func (m *LicenseContent) GetEncoding()(*string) {
    return m.encoding
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LicenseContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseContent__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(LicenseContent__linksable))
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
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
    res["encoding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncoding(val)
        }
        return nil
    }
    res["git_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitUrl(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetGitUrl gets the git_url property value. The git_url property
// returns a *string when successful
func (m *LicenseContent) GetGitUrl()(*string) {
    return m.git_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *LicenseContent) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLicense gets the license property value. License Simple
// returns a NullableLicenseSimpleable when successful
func (m *LicenseContent) GetLicense()(NullableLicenseSimpleable) {
    return m.license
}
// GetLinks gets the _links property value. The _links property
// returns a LicenseContent__linksable when successful
func (m *LicenseContent) GetLinks()(LicenseContent__linksable) {
    return m._links
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *LicenseContent) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *LicenseContent) GetPath()(*string) {
    return m.path
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *LicenseContent) GetSha()(*string) {
    return m.sha
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *LicenseContent) GetSize()(*int32) {
    return m.size
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *LicenseContent) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *LicenseContent) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *LicenseContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
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
        err := writer.WriteStringValue("encoding", m.GetEncoding())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_url", m.GetGitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_links", m.GetLinks())
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
func (m *LicenseContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *LicenseContent) SetContent(value *string)() {
    m.content = value
}
// SetDownloadUrl sets the download_url property value. The download_url property
func (m *LicenseContent) SetDownloadUrl(value *string)() {
    m.download_url = value
}
// SetEncoding sets the encoding property value. The encoding property
func (m *LicenseContent) SetEncoding(value *string)() {
    m.encoding = value
}
// SetGitUrl sets the git_url property value. The git_url property
func (m *LicenseContent) SetGitUrl(value *string)() {
    m.git_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *LicenseContent) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLicense sets the license property value. License Simple
func (m *LicenseContent) SetLicense(value NullableLicenseSimpleable)() {
    m.license = value
}
// SetLinks sets the _links property value. The _links property
func (m *LicenseContent) SetLinks(value LicenseContent__linksable)() {
    m._links = value
}
// SetName sets the name property value. The name property
func (m *LicenseContent) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. The path property
func (m *LicenseContent) SetPath(value *string)() {
    m.path = value
}
// SetSha sets the sha property value. The sha property
func (m *LicenseContent) SetSha(value *string)() {
    m.sha = value
}
// SetSize sets the size property value. The size property
func (m *LicenseContent) SetSize(value *int32)() {
    m.size = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *LicenseContent) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *LicenseContent) SetUrl(value *string)() {
    m.url = value
}
type LicenseContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetDownloadUrl()(*string)
    GetEncoding()(*string)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLicense()(NullableLicenseSimpleable)
    GetLinks()(LicenseContent__linksable)
    GetName()(*string)
    GetPath()(*string)
    GetSha()(*string)
    GetSize()(*int32)
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetContent(value *string)()
    SetDownloadUrl(value *string)()
    SetEncoding(value *string)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLicense(value NullableLicenseSimpleable)()
    SetLinks(value LicenseContent__linksable)()
    SetName(value *string)()
    SetPath(value *string)()
    SetSha(value *string)()
    SetSize(value *int32)()
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
