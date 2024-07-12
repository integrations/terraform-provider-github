package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentFile content File
type ContentFile struct {
    // The _links property
    _links ContentFile__linksable
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
    // The name property
    name *string
    // The path property
    path *string
    // The sha property
    sha *string
    // The size property
    size *int32
    // The submodule_git_url property
    submodule_git_url *string
    // The target property
    target *string
    // The type property
    typeEscaped *ContentFile_type
    // The url property
    url *string
}
// NewContentFile instantiates a new ContentFile and sets the default values.
func NewContentFile()(*ContentFile) {
    m := &ContentFile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContentFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentFile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContentFile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
// returns a *string when successful
func (m *ContentFile) GetContent()(*string) {
    return m.content
}
// GetDownloadUrl gets the download_url property value. The download_url property
// returns a *string when successful
func (m *ContentFile) GetDownloadUrl()(*string) {
    return m.download_url
}
// GetEncoding gets the encoding property value. The encoding property
// returns a *string when successful
func (m *ContentFile) GetEncoding()(*string) {
    return m.encoding
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContentFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentFile__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(ContentFile__linksable))
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
    res["submodule_git_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmoduleGitUrl(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentFile_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ContentFile_type))
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
func (m *ContentFile) GetGitUrl()(*string) {
    return m.git_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *ContentFile) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLinks gets the _links property value. The _links property
// returns a ContentFile__linksable when successful
func (m *ContentFile) GetLinks()(ContentFile__linksable) {
    return m._links
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ContentFile) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ContentFile) GetPath()(*string) {
    return m.path
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *ContentFile) GetSha()(*string) {
    return m.sha
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *ContentFile) GetSize()(*int32) {
    return m.size
}
// GetSubmoduleGitUrl gets the submodule_git_url property value. The submodule_git_url property
// returns a *string when successful
func (m *ContentFile) GetSubmoduleGitUrl()(*string) {
    return m.submodule_git_url
}
// GetTarget gets the target property value. The target property
// returns a *string when successful
func (m *ContentFile) GetTarget()(*string) {
    return m.target
}
// GetTypeEscaped gets the type property value. The type property
// returns a *ContentFile_type when successful
func (m *ContentFile) GetTypeEscaped()(*ContentFile_type) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ContentFile) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ContentFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("submodule_git_url", m.GetSubmoduleGitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *ContentFile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *ContentFile) SetContent(value *string)() {
    m.content = value
}
// SetDownloadUrl sets the download_url property value. The download_url property
func (m *ContentFile) SetDownloadUrl(value *string)() {
    m.download_url = value
}
// SetEncoding sets the encoding property value. The encoding property
func (m *ContentFile) SetEncoding(value *string)() {
    m.encoding = value
}
// SetGitUrl sets the git_url property value. The git_url property
func (m *ContentFile) SetGitUrl(value *string)() {
    m.git_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *ContentFile) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLinks sets the _links property value. The _links property
func (m *ContentFile) SetLinks(value ContentFile__linksable)() {
    m._links = value
}
// SetName sets the name property value. The name property
func (m *ContentFile) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. The path property
func (m *ContentFile) SetPath(value *string)() {
    m.path = value
}
// SetSha sets the sha property value. The sha property
func (m *ContentFile) SetSha(value *string)() {
    m.sha = value
}
// SetSize sets the size property value. The size property
func (m *ContentFile) SetSize(value *int32)() {
    m.size = value
}
// SetSubmoduleGitUrl sets the submodule_git_url property value. The submodule_git_url property
func (m *ContentFile) SetSubmoduleGitUrl(value *string)() {
    m.submodule_git_url = value
}
// SetTarget sets the target property value. The target property
func (m *ContentFile) SetTarget(value *string)() {
    m.target = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ContentFile) SetTypeEscaped(value *ContentFile_type)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *ContentFile) SetUrl(value *string)() {
    m.url = value
}
type ContentFileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetDownloadUrl()(*string)
    GetEncoding()(*string)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLinks()(ContentFile__linksable)
    GetName()(*string)
    GetPath()(*string)
    GetSha()(*string)
    GetSize()(*int32)
    GetSubmoduleGitUrl()(*string)
    GetTarget()(*string)
    GetTypeEscaped()(*ContentFile_type)
    GetUrl()(*string)
    SetContent(value *string)()
    SetDownloadUrl(value *string)()
    SetEncoding(value *string)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLinks(value ContentFile__linksable)()
    SetName(value *string)()
    SetPath(value *string)()
    SetSha(value *string)()
    SetSize(value *int32)()
    SetSubmoduleGitUrl(value *string)()
    SetTarget(value *string)()
    SetTypeEscaped(value *ContentFile_type)()
    SetUrl(value *string)()
}
