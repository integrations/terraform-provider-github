package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReleaseAsset data related to a release.
type ReleaseAsset struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The browser_download_url property
    browser_download_url *string
    // The content_type property
    content_type *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The download_count property
    download_count *int32
    // The id property
    id *int32
    // The label property
    label *string
    // The file name of the asset.
    name *string
    // The node_id property
    node_id *string
    // The size property
    size *int32
    // State of the release asset.
    state *ReleaseAsset_state
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    uploader NullableSimpleUserable
    // The url property
    url *string
}
// NewReleaseAsset instantiates a new ReleaseAsset and sets the default values.
func NewReleaseAsset()(*ReleaseAsset) {
    m := &ReleaseAsset{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReleaseAssetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReleaseAssetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReleaseAsset(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReleaseAsset) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBrowserDownloadUrl gets the browser_download_url property value. The browser_download_url property
// returns a *string when successful
func (m *ReleaseAsset) GetBrowserDownloadUrl()(*string) {
    return m.browser_download_url
}
// GetContentType gets the content_type property value. The content_type property
// returns a *string when successful
func (m *ReleaseAsset) GetContentType()(*string) {
    return m.content_type
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *ReleaseAsset) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDownloadCount gets the download_count property value. The download_count property
// returns a *int32 when successful
func (m *ReleaseAsset) GetDownloadCount()(*int32) {
    return m.download_count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReleaseAsset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["browser_download_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserDownloadUrl(val)
        }
        return nil
    }
    res["content_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["download_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadCount(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReleaseAsset_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ReleaseAsset_state))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["uploader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploader(val.(NullableSimpleUserable))
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *ReleaseAsset) GetId()(*int32) {
    return m.id
}
// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *ReleaseAsset) GetLabel()(*string) {
    return m.label
}
// GetName gets the name property value. The file name of the asset.
// returns a *string when successful
func (m *ReleaseAsset) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *ReleaseAsset) GetNodeId()(*string) {
    return m.node_id
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *ReleaseAsset) GetSize()(*int32) {
    return m.size
}
// GetState gets the state property value. State of the release asset.
// returns a *ReleaseAsset_state when successful
func (m *ReleaseAsset) GetState()(*ReleaseAsset_state) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *ReleaseAsset) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUploader gets the uploader property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *ReleaseAsset) GetUploader()(NullableSimpleUserable) {
    return m.uploader
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ReleaseAsset) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ReleaseAsset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("browser_download_url", m.GetBrowserDownloadUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("download_count", m.GetDownloadCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("uploader", m.GetUploader())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReleaseAsset) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBrowserDownloadUrl sets the browser_download_url property value. The browser_download_url property
func (m *ReleaseAsset) SetBrowserDownloadUrl(value *string)() {
    m.browser_download_url = value
}
// SetContentType sets the content_type property value. The content_type property
func (m *ReleaseAsset) SetContentType(value *string)() {
    m.content_type = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ReleaseAsset) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDownloadCount sets the download_count property value. The download_count property
func (m *ReleaseAsset) SetDownloadCount(value *int32)() {
    m.download_count = value
}
// SetId sets the id property value. The id property
func (m *ReleaseAsset) SetId(value *int32)() {
    m.id = value
}
// SetLabel sets the label property value. The label property
func (m *ReleaseAsset) SetLabel(value *string)() {
    m.label = value
}
// SetName sets the name property value. The file name of the asset.
func (m *ReleaseAsset) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *ReleaseAsset) SetNodeId(value *string)() {
    m.node_id = value
}
// SetSize sets the size property value. The size property
func (m *ReleaseAsset) SetSize(value *int32)() {
    m.size = value
}
// SetState sets the state property value. State of the release asset.
func (m *ReleaseAsset) SetState(value *ReleaseAsset_state)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *ReleaseAsset) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUploader sets the uploader property value. A GitHub user.
func (m *ReleaseAsset) SetUploader(value NullableSimpleUserable)() {
    m.uploader = value
}
// SetUrl sets the url property value. The url property
func (m *ReleaseAsset) SetUrl(value *string)() {
    m.url = value
}
type ReleaseAssetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrowserDownloadUrl()(*string)
    GetContentType()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDownloadCount()(*int32)
    GetId()(*int32)
    GetLabel()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetSize()(*int32)
    GetState()(*ReleaseAsset_state)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUploader()(NullableSimpleUserable)
    GetUrl()(*string)
    SetBrowserDownloadUrl(value *string)()
    SetContentType(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDownloadCount(value *int32)()
    SetId(value *int32)()
    SetLabel(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetSize(value *int32)()
    SetState(value *ReleaseAsset_state)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUploader(value NullableSimpleUserable)()
    SetUrl(value *string)()
}
