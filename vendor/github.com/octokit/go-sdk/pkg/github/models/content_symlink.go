package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentSymlink an object describing a symlink
type ContentSymlink struct {
    // The _links property
    _links ContentSymlink__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The download_url property
    download_url *string
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
    // The target property
    target *string
    // The type property
    typeEscaped *ContentSymlink_type
    // The url property
    url *string
}
// NewContentSymlink instantiates a new ContentSymlink and sets the default values.
func NewContentSymlink()(*ContentSymlink) {
    m := &ContentSymlink{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentSymlinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContentSymlinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentSymlink(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContentSymlink) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDownloadUrl gets the download_url property value. The download_url property
// returns a *string when successful
func (m *ContentSymlink) GetDownloadUrl()(*string) {
    return m.download_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContentSymlink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentSymlink__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(ContentSymlink__linksable))
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
        val, err := n.GetEnumValue(ParseContentSymlink_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ContentSymlink_type))
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
func (m *ContentSymlink) GetGitUrl()(*string) {
    return m.git_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *ContentSymlink) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLinks gets the _links property value. The _links property
// returns a ContentSymlink__linksable when successful
func (m *ContentSymlink) GetLinks()(ContentSymlink__linksable) {
    return m._links
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ContentSymlink) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ContentSymlink) GetPath()(*string) {
    return m.path
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *ContentSymlink) GetSha()(*string) {
    return m.sha
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *ContentSymlink) GetSize()(*int32) {
    return m.size
}
// GetTarget gets the target property value. The target property
// returns a *string when successful
func (m *ContentSymlink) GetTarget()(*string) {
    return m.target
}
// GetTypeEscaped gets the type property value. The type property
// returns a *ContentSymlink_type when successful
func (m *ContentSymlink) GetTypeEscaped()(*ContentSymlink_type) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ContentSymlink) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ContentSymlink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("download_url", m.GetDownloadUrl())
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
func (m *ContentSymlink) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDownloadUrl sets the download_url property value. The download_url property
func (m *ContentSymlink) SetDownloadUrl(value *string)() {
    m.download_url = value
}
// SetGitUrl sets the git_url property value. The git_url property
func (m *ContentSymlink) SetGitUrl(value *string)() {
    m.git_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *ContentSymlink) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLinks sets the _links property value. The _links property
func (m *ContentSymlink) SetLinks(value ContentSymlink__linksable)() {
    m._links = value
}
// SetName sets the name property value. The name property
func (m *ContentSymlink) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. The path property
func (m *ContentSymlink) SetPath(value *string)() {
    m.path = value
}
// SetSha sets the sha property value. The sha property
func (m *ContentSymlink) SetSha(value *string)() {
    m.sha = value
}
// SetSize sets the size property value. The size property
func (m *ContentSymlink) SetSize(value *int32)() {
    m.size = value
}
// SetTarget sets the target property value. The target property
func (m *ContentSymlink) SetTarget(value *string)() {
    m.target = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ContentSymlink) SetTypeEscaped(value *ContentSymlink_type)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *ContentSymlink) SetUrl(value *string)() {
    m.url = value
}
type ContentSymlinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDownloadUrl()(*string)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLinks()(ContentSymlink__linksable)
    GetName()(*string)
    GetPath()(*string)
    GetSha()(*string)
    GetSize()(*int32)
    GetTarget()(*string)
    GetTypeEscaped()(*ContentSymlink_type)
    GetUrl()(*string)
    SetDownloadUrl(value *string)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLinks(value ContentSymlink__linksable)()
    SetName(value *string)()
    SetPath(value *string)()
    SetSha(value *string)()
    SetSize(value *int32)()
    SetTarget(value *string)()
    SetTypeEscaped(value *ContentSymlink_type)()
    SetUrl(value *string)()
}
