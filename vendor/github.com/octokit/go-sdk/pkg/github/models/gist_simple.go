package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistSimple gist Simple
type GistSimple struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments *int32
    // The comments_url property
    comments_url *string
    // The commits_url property
    commits_url *string
    // The created_at property
    created_at *string
    // The description property
    description *string
    // The files property
    files GistSimple_filesable
    // Gist
    fork_of GistSimple_fork_ofable
    // The forks property
    // Deprecated: 
    forks []GistSimple_forksable
    // The forks_url property
    forks_url *string
    // The git_pull_url property
    git_pull_url *string
    // The git_push_url property
    git_push_url *string
    // The history property
    // Deprecated: 
    history []GistHistoryable
    // The html_url property
    html_url *string
    // The id property
    id *string
    // The node_id property
    node_id *string
    // A GitHub user.
    owner SimpleUserable
    // The public property
    public *bool
    // The truncated property
    truncated *bool
    // The updated_at property
    updated_at *string
    // The url property
    url *string
    // The user property
    user *string
}
// NewGistSimple instantiates a new GistSimple and sets the default values.
func NewGistSimple()(*GistSimple) {
    m := &GistSimple{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistSimpleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistSimpleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistSimple(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GistSimple) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
// returns a *int32 when successful
func (m *GistSimple) GetComments()(*int32) {
    return m.comments
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *GistSimple) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommitsUrl gets the commits_url property value. The commits_url property
// returns a *string when successful
func (m *GistSimple) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *GistSimple) GetCreatedAt()(*string) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *GistSimple) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistSimple) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
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
        val, err := n.GetObjectValue(CreateGistSimple_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(GistSimple_filesable))
        }
        return nil
    }
    res["fork_of"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistSimple_fork_ofFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForkOf(val.(GistSimple_fork_ofable))
        }
        return nil
    }
    res["forks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGistSimple_forksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GistSimple_forksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GistSimple_forksable)
                }
            }
            m.SetForks(res)
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
        val, err := n.GetCollectionOfObjectValues(CreateGistHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GistHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GistHistoryable)
                }
            }
            m.SetHistory(res)
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
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(SimpleUserable))
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
        val, err := n.GetStringValue()
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The files property
// returns a GistSimple_filesable when successful
func (m *GistSimple) GetFiles()(GistSimple_filesable) {
    return m.files
}
// GetForkOf gets the fork_of property value. Gist
// returns a GistSimple_fork_ofable when successful
func (m *GistSimple) GetForkOf()(GistSimple_fork_ofable) {
    return m.fork_of
}
// GetForks gets the forks property value. The forks property
// Deprecated: 
// returns a []GistSimple_forksable when successful
func (m *GistSimple) GetForks()([]GistSimple_forksable) {
    return m.forks
}
// GetForksUrl gets the forks_url property value. The forks_url property
// returns a *string when successful
func (m *GistSimple) GetForksUrl()(*string) {
    return m.forks_url
}
// GetGitPullUrl gets the git_pull_url property value. The git_pull_url property
// returns a *string when successful
func (m *GistSimple) GetGitPullUrl()(*string) {
    return m.git_pull_url
}
// GetGitPushUrl gets the git_push_url property value. The git_push_url property
// returns a *string when successful
func (m *GistSimple) GetGitPushUrl()(*string) {
    return m.git_push_url
}
// GetHistory gets the history property value. The history property
// Deprecated: 
// returns a []GistHistoryable when successful
func (m *GistSimple) GetHistory()([]GistHistoryable) {
    return m.history
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *GistSimple) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *GistSimple) GetId()(*string) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *GistSimple) GetNodeId()(*string) {
    return m.node_id
}
// GetOwner gets the owner property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *GistSimple) GetOwner()(SimpleUserable) {
    return m.owner
}
// GetPublic gets the public property value. The public property
// returns a *bool when successful
func (m *GistSimple) GetPublic()(*bool) {
    return m.public
}
// GetTruncated gets the truncated property value. The truncated property
// returns a *bool when successful
func (m *GistSimple) GetTruncated()(*bool) {
    return m.truncated
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *GistSimple) GetUpdatedAt()(*string) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GistSimple) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. The user property
// returns a *string when successful
func (m *GistSimple) GetUser()(*string) {
    return m.user
}
// Serialize serializes information the current object
func (m *GistSimple) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
    if m.GetForks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForks()))
        for i, v := range m.GetForks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("forks", cast)
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
        err := writer.WriteObjectValue("fork_of", m.GetForkOf())
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
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("history", cast)
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
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *GistSimple) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *GistSimple) SetComments(value *int32)() {
    m.comments = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *GistSimple) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommitsUrl sets the commits_url property value. The commits_url property
func (m *GistSimple) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *GistSimple) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *GistSimple) SetDescription(value *string)() {
    m.description = value
}
// SetFiles sets the files property value. The files property
func (m *GistSimple) SetFiles(value GistSimple_filesable)() {
    m.files = value
}
// SetForkOf sets the fork_of property value. Gist
func (m *GistSimple) SetForkOf(value GistSimple_fork_ofable)() {
    m.fork_of = value
}
// SetForks sets the forks property value. The forks property
// Deprecated: 
func (m *GistSimple) SetForks(value []GistSimple_forksable)() {
    m.forks = value
}
// SetForksUrl sets the forks_url property value. The forks_url property
func (m *GistSimple) SetForksUrl(value *string)() {
    m.forks_url = value
}
// SetGitPullUrl sets the git_pull_url property value. The git_pull_url property
func (m *GistSimple) SetGitPullUrl(value *string)() {
    m.git_pull_url = value
}
// SetGitPushUrl sets the git_push_url property value. The git_push_url property
func (m *GistSimple) SetGitPushUrl(value *string)() {
    m.git_push_url = value
}
// SetHistory sets the history property value. The history property
// Deprecated: 
func (m *GistSimple) SetHistory(value []GistHistoryable)() {
    m.history = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *GistSimple) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *GistSimple) SetId(value *string)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *GistSimple) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *GistSimple) SetOwner(value SimpleUserable)() {
    m.owner = value
}
// SetPublic sets the public property value. The public property
func (m *GistSimple) SetPublic(value *bool)() {
    m.public = value
}
// SetTruncated sets the truncated property value. The truncated property
func (m *GistSimple) SetTruncated(value *bool)() {
    m.truncated = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *GistSimple) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *GistSimple) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. The user property
func (m *GistSimple) SetUser(value *string)() {
    m.user = value
}
type GistSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCreatedAt()(*string)
    GetDescription()(*string)
    GetFiles()(GistSimple_filesable)
    GetForkOf()(GistSimple_fork_ofable)
    GetForks()([]GistSimple_forksable)
    GetForksUrl()(*string)
    GetGitPullUrl()(*string)
    GetGitPushUrl()(*string)
    GetHistory()([]GistHistoryable)
    GetHtmlUrl()(*string)
    GetId()(*string)
    GetNodeId()(*string)
    GetOwner()(SimpleUserable)
    GetPublic()(*bool)
    GetTruncated()(*bool)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    GetUser()(*string)
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCreatedAt(value *string)()
    SetDescription(value *string)()
    SetFiles(value GistSimple_filesable)()
    SetForkOf(value GistSimple_fork_ofable)()
    SetForks(value []GistSimple_forksable)()
    SetForksUrl(value *string)()
    SetGitPullUrl(value *string)()
    SetGitPushUrl(value *string)()
    SetHistory(value []GistHistoryable)()
    SetHtmlUrl(value *string)()
    SetId(value *string)()
    SetNodeId(value *string)()
    SetOwner(value SimpleUserable)()
    SetPublic(value *bool)()
    SetTruncated(value *bool)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
    SetUser(value *string)()
}
