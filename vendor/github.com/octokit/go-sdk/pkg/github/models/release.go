package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Release a release.
type Release struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assets property
    assets []ReleaseAssetable
    // The assets_url property
    assets_url *string
    // A GitHub user.
    author SimpleUserable
    // The body property
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL of the release discussion.
    discussion_url *string
    // true to create a draft (unpublished) release, false to create a published one.
    draft *bool
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The mentions_count property
    mentions_count *int32
    // The name property
    name *string
    // The node_id property
    node_id *string
    // Whether to identify the release as a prerelease or a full release.
    prerelease *bool
    // The published_at property
    published_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reactions property
    reactions ReactionRollupable
    // The name of the tag.
    tag_name *string
    // The tarball_url property
    tarball_url *string
    // Specifies the commitish value that determines where the Git tag is created from.
    target_commitish *string
    // The upload_url property
    upload_url *string
    // The url property
    url *string
    // The zipball_url property
    zipball_url *string
}
// NewRelease instantiates a new Release and sets the default values.
func NewRelease()(*Release) {
    m := &Release{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReleaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReleaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelease(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Release) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssets gets the assets property value. The assets property
// returns a []ReleaseAssetable when successful
func (m *Release) GetAssets()([]ReleaseAssetable) {
    return m.assets
}
// GetAssetsUrl gets the assets_url property value. The assets_url property
// returns a *string when successful
func (m *Release) GetAssetsUrl()(*string) {
    return m.assets_url
}
// GetAuthor gets the author property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *Release) GetAuthor()(SimpleUserable) {
    return m.author
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *Release) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *Release) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *Release) GetBodyText()(*string) {
    return m.body_text
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Release) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDiscussionUrl gets the discussion_url property value. The URL of the release discussion.
// returns a *string when successful
func (m *Release) GetDiscussionUrl()(*string) {
    return m.discussion_url
}
// GetDraft gets the draft property value. true to create a draft (unpublished) release, false to create a published one.
// returns a *bool when successful
func (m *Release) GetDraft()(*bool) {
    return m.draft
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Release) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateReleaseAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReleaseAssetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ReleaseAssetable)
                }
            }
            m.SetAssets(res)
        }
        return nil
    }
    res["assets_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetsUrl(val)
        }
        return nil
    }
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(SimpleUserable))
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["body_html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyHtml(val)
        }
        return nil
    }
    res["body_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyText(val)
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
    res["discussion_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscussionUrl(val)
        }
        return nil
    }
    res["draft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraft(val)
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
    res["mentions_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionsCount(val)
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
    res["prerelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrerelease(val)
        }
        return nil
    }
    res["published_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedAt(val)
        }
        return nil
    }
    res["reactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReactionRollupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactions(val.(ReactionRollupable))
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
    res["tarball_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarballUrl(val)
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
    res["upload_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadUrl(val)
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
    res["zipball_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZipballUrl(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Release) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Release) GetId()(*int32) {
    return m.id
}
// GetMentionsCount gets the mentions_count property value. The mentions_count property
// returns a *int32 when successful
func (m *Release) GetMentionsCount()(*int32) {
    return m.mentions_count
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Release) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Release) GetNodeId()(*string) {
    return m.node_id
}
// GetPrerelease gets the prerelease property value. Whether to identify the release as a prerelease or a full release.
// returns a *bool when successful
func (m *Release) GetPrerelease()(*bool) {
    return m.prerelease
}
// GetPublishedAt gets the published_at property value. The published_at property
// returns a *Time when successful
func (m *Release) GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.published_at
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *Release) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetTagName gets the tag_name property value. The name of the tag.
// returns a *string when successful
func (m *Release) GetTagName()(*string) {
    return m.tag_name
}
// GetTarballUrl gets the tarball_url property value. The tarball_url property
// returns a *string when successful
func (m *Release) GetTarballUrl()(*string) {
    return m.tarball_url
}
// GetTargetCommitish gets the target_commitish property value. Specifies the commitish value that determines where the Git tag is created from.
// returns a *string when successful
func (m *Release) GetTargetCommitish()(*string) {
    return m.target_commitish
}
// GetUploadUrl gets the upload_url property value. The upload_url property
// returns a *string when successful
func (m *Release) GetUploadUrl()(*string) {
    return m.upload_url
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Release) GetUrl()(*string) {
    return m.url
}
// GetZipballUrl gets the zipball_url property value. The zipball_url property
// returns a *string when successful
func (m *Release) GetZipballUrl()(*string) {
    return m.zipball_url
}
// Serialize serializes information the current object
func (m *Release) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssets()))
        for i, v := range m.GetAssets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("assets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assets_url", m.GetAssetsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body_html", m.GetBodyHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body_text", m.GetBodyText())
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
        err := writer.WriteStringValue("discussion_url", m.GetDiscussionUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("draft", m.GetDraft())
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
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("mentions_count", m.GetMentionsCount())
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
        err := writer.WriteBoolValue("prerelease", m.GetPrerelease())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("published_at", m.GetPublishedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reactions", m.GetReactions())
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
        err := writer.WriteStringValue("tarball_url", m.GetTarballUrl())
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
        err := writer.WriteStringValue("upload_url", m.GetUploadUrl())
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
        err := writer.WriteStringValue("zipball_url", m.GetZipballUrl())
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
func (m *Release) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssets sets the assets property value. The assets property
func (m *Release) SetAssets(value []ReleaseAssetable)() {
    m.assets = value
}
// SetAssetsUrl sets the assets_url property value. The assets_url property
func (m *Release) SetAssetsUrl(value *string)() {
    m.assets_url = value
}
// SetAuthor sets the author property value. A GitHub user.
func (m *Release) SetAuthor(value SimpleUserable)() {
    m.author = value
}
// SetBody sets the body property value. The body property
func (m *Release) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *Release) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *Release) SetBodyText(value *string)() {
    m.body_text = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Release) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDiscussionUrl sets the discussion_url property value. The URL of the release discussion.
func (m *Release) SetDiscussionUrl(value *string)() {
    m.discussion_url = value
}
// SetDraft sets the draft property value. true to create a draft (unpublished) release, false to create a published one.
func (m *Release) SetDraft(value *bool)() {
    m.draft = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Release) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Release) SetId(value *int32)() {
    m.id = value
}
// SetMentionsCount sets the mentions_count property value. The mentions_count property
func (m *Release) SetMentionsCount(value *int32)() {
    m.mentions_count = value
}
// SetName sets the name property value. The name property
func (m *Release) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Release) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPrerelease sets the prerelease property value. Whether to identify the release as a prerelease or a full release.
func (m *Release) SetPrerelease(value *bool)() {
    m.prerelease = value
}
// SetPublishedAt sets the published_at property value. The published_at property
func (m *Release) SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.published_at = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *Release) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetTagName sets the tag_name property value. The name of the tag.
func (m *Release) SetTagName(value *string)() {
    m.tag_name = value
}
// SetTarballUrl sets the tarball_url property value. The tarball_url property
func (m *Release) SetTarballUrl(value *string)() {
    m.tarball_url = value
}
// SetTargetCommitish sets the target_commitish property value. Specifies the commitish value that determines where the Git tag is created from.
func (m *Release) SetTargetCommitish(value *string)() {
    m.target_commitish = value
}
// SetUploadUrl sets the upload_url property value. The upload_url property
func (m *Release) SetUploadUrl(value *string)() {
    m.upload_url = value
}
// SetUrl sets the url property value. The url property
func (m *Release) SetUrl(value *string)() {
    m.url = value
}
// SetZipballUrl sets the zipball_url property value. The zipball_url property
func (m *Release) SetZipballUrl(value *string)() {
    m.zipball_url = value
}
type Releaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssets()([]ReleaseAssetable)
    GetAssetsUrl()(*string)
    GetAuthor()(SimpleUserable)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiscussionUrl()(*string)
    GetDraft()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetMentionsCount()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetPrerelease()(*bool)
    GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReactions()(ReactionRollupable)
    GetTagName()(*string)
    GetTarballUrl()(*string)
    GetTargetCommitish()(*string)
    GetUploadUrl()(*string)
    GetUrl()(*string)
    GetZipballUrl()(*string)
    SetAssets(value []ReleaseAssetable)()
    SetAssetsUrl(value *string)()
    SetAuthor(value SimpleUserable)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiscussionUrl(value *string)()
    SetDraft(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetMentionsCount(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetPrerelease(value *bool)()
    SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReactions(value ReactionRollupable)()
    SetTagName(value *string)()
    SetTarballUrl(value *string)()
    SetTargetCommitish(value *string)()
    SetUploadUrl(value *string)()
    SetUrl(value *string)()
    SetZipballUrl(value *string)()
}
