package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistSimple_fork_of gist
type GistSimple_fork_of struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments *int32
    // The comments_url property
    comments_url *string
    // The commits_url property
    commits_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The files property
    files GistSimple_fork_of_filesable
    // The forks property
    forks i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The forks_url property
    forks_url *string
    // The git_pull_url property
    git_pull_url *string
    // The git_push_url property
    git_push_url *string
    // The history property
    history i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The html_url property
    html_url *string
    // The id property
    id *string
    // The node_id property
    node_id *string
    // A GitHub user.
    owner NullableSimpleUserable
    // The public property
    public *bool
    // The truncated property
    truncated *bool
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewGistSimple_fork_of instantiates a new GistSimple_fork_of and sets the default values.
func NewGistSimple_fork_of()(*GistSimple_fork_of) {
    m := &GistSimple_fork_of{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistSimple_fork_ofFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistSimple_fork_ofFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistSimple_fork_of(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GistSimple_fork_of) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
// returns a *int32 when successful
func (m *GistSimple_fork_of) GetComments()(*int32) {
    return m.comments
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommitsUrl gets the commits_url property value. The commits_url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *GistSimple_fork_of) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *GistSimple_fork_of) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistSimple_fork_of) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val)
        }
        return nil
    }
    res["comments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsUrl(val)
        }
        return nil
    }
    res["commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitsUrl(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistSimple_fork_of_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(GistSimple_fork_of_filesable))
        }
        return nil
    }
    res["forks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForks(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["forks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForksUrl(val)
        }
        return nil
    }
    res["git_pull_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitPullUrl(val)
        }
        return nil
    }
    res["git_push_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitPushUrl(val)
        }
        return nil
    }
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHistory(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublic(val)
        }
        return nil
    }
    res["truncated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTruncated(val)
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The files property
// returns a GistSimple_fork_of_filesable when successful
func (m *GistSimple_fork_of) GetFiles()(GistSimple_fork_of_filesable) {
    return m.files
}
// GetForks gets the forks property value. The forks property
// returns a UntypedNodeable when successful
func (m *GistSimple_fork_of) GetForks()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.forks
}
// GetForksUrl gets the forks_url property value. The forks_url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetForksUrl()(*string) {
    return m.forks_url
}
// GetGitPullUrl gets the git_pull_url property value. The git_pull_url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetGitPullUrl()(*string) {
    return m.git_pull_url
}
// GetGitPushUrl gets the git_push_url property value. The git_push_url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetGitPushUrl()(*string) {
    return m.git_push_url
}
// GetHistory gets the history property value. The history property
// returns a UntypedNodeable when successful
func (m *GistSimple_fork_of) GetHistory()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.history
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *GistSimple_fork_of) GetId()(*string) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *GistSimple_fork_of) GetNodeId()(*string) {
    return m.node_id
}
// GetOwner gets the owner property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *GistSimple_fork_of) GetOwner()(NullableSimpleUserable) {
    return m.owner
}
// GetPublic gets the public property value. The public property
// returns a *bool when successful
func (m *GistSimple_fork_of) GetPublic()(*bool) {
    return m.public
}
// GetTruncated gets the truncated property value. The truncated property
// returns a *bool when successful
func (m *GistSimple_fork_of) GetTruncated()(*bool) {
    return m.truncated
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *GistSimple_fork_of) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GistSimple_fork_of) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *GistSimple_fork_of) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *GistSimple_fork_of) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comments_url", m.GetCommentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commits_url", m.GetCommitsUrl())
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
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("files", m.GetFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("forks", m.GetForks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("forks_url", m.GetForksUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_pull_url", m.GetGitPullUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_push_url", m.GetGitPushUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("history", m.GetHistory())
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
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("public", m.GetPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("truncated", m.GetTruncated())
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
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *GistSimple_fork_of) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *GistSimple_fork_of) SetComments(value *int32)() {
    m.comments = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *GistSimple_fork_of) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommitsUrl sets the commits_url property value. The commits_url property
func (m *GistSimple_fork_of) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *GistSimple_fork_of) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *GistSimple_fork_of) SetDescription(value *string)() {
    m.description = value
}
// SetFiles sets the files property value. The files property
func (m *GistSimple_fork_of) SetFiles(value GistSimple_fork_of_filesable)() {
    m.files = value
}
// SetForks sets the forks property value. The forks property
func (m *GistSimple_fork_of) SetForks(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.forks = value
}
// SetForksUrl sets the forks_url property value. The forks_url property
func (m *GistSimple_fork_of) SetForksUrl(value *string)() {
    m.forks_url = value
}
// SetGitPullUrl sets the git_pull_url property value. The git_pull_url property
func (m *GistSimple_fork_of) SetGitPullUrl(value *string)() {
    m.git_pull_url = value
}
// SetGitPushUrl sets the git_push_url property value. The git_push_url property
func (m *GistSimple_fork_of) SetGitPushUrl(value *string)() {
    m.git_push_url = value
}
// SetHistory sets the history property value. The history property
func (m *GistSimple_fork_of) SetHistory(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.history = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *GistSimple_fork_of) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *GistSimple_fork_of) SetId(value *string)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *GistSimple_fork_of) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *GistSimple_fork_of) SetOwner(value NullableSimpleUserable)() {
    m.owner = value
}
// SetPublic sets the public property value. The public property
func (m *GistSimple_fork_of) SetPublic(value *bool)() {
    m.public = value
}
// SetTruncated sets the truncated property value. The truncated property
func (m *GistSimple_fork_of) SetTruncated(value *bool)() {
    m.truncated = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *GistSimple_fork_of) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *GistSimple_fork_of) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *GistSimple_fork_of) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type GistSimple_fork_ofable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetFiles()(GistSimple_fork_of_filesable)
    GetForks()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetForksUrl()(*string)
    GetGitPullUrl()(*string)
    GetGitPushUrl()(*string)
    GetHistory()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetHtmlUrl()(*string)
    GetId()(*string)
    GetNodeId()(*string)
    GetOwner()(NullableSimpleUserable)
    GetPublic()(*bool)
    GetTruncated()(*bool)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetFiles(value GistSimple_fork_of_filesable)()
    SetForks(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetForksUrl(value *string)()
    SetGitPullUrl(value *string)()
    SetGitPushUrl(value *string)()
    SetHistory(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetHtmlUrl(value *string)()
    SetId(value *string)()
    SetNodeId(value *string)()
    SetOwner(value NullableSimpleUserable)()
    SetPublic(value *bool)()
    SetTruncated(value *bool)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
