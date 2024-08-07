package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitSearchResultItem commit Search Result Item
type CommitSearchResultItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    author NullableSimpleUserable
    // The comments_url property
    comments_url *string
    // The commit property
    commit CommitSearchResultItem_commitable
    // Metaproperties for Git author/committer information.
    committer NullableGitUserable
    // The html_url property
    html_url *string
    // The node_id property
    node_id *string
    // The parents property
    parents []CommitSearchResultItem_parentsable
    // Minimal Repository
    repository MinimalRepositoryable
    // The score property
    score *float64
    // The sha property
    sha *string
    // The text_matches property
    text_matches []Commitsable
    // The url property
    url *string
}
// NewCommitSearchResultItem instantiates a new CommitSearchResultItem and sets the default values.
func NewCommitSearchResultItem()(*CommitSearchResultItem) {
    m := &CommitSearchResultItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommitSearchResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommitSearchResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommitSearchResultItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommitSearchResultItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *CommitSearchResultItem) GetAuthor()(NullableSimpleUserable) {
    return m.author
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *CommitSearchResultItem) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommit gets the commit property value. The commit property
// returns a CommitSearchResultItem_commitable when successful
func (m *CommitSearchResultItem) GetCommit()(CommitSearchResultItem_commitable) {
    return m.commit
}
// GetCommitter gets the committer property value. Metaproperties for Git author/committer information.
// returns a NullableGitUserable when successful
func (m *CommitSearchResultItem) GetCommitter()(NullableGitUserable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommitSearchResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(NullableSimpleUserable))
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
    res["commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommitSearchResultItem_commitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(CommitSearchResultItem_commitable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableGitUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(NullableGitUserable))
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
    res["parents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommitSearchResultItem_parentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommitSearchResultItem_parentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CommitSearchResultItem_parentsable)
                }
            }
            m.SetParents(res)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
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
    res["text_matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommitsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Commitsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Commitsable)
                }
            }
            m.SetTextMatches(res)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *CommitSearchResultItem) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *CommitSearchResultItem) GetNodeId()(*string) {
    return m.node_id
}
// GetParents gets the parents property value. The parents property
// returns a []CommitSearchResultItem_parentsable when successful
func (m *CommitSearchResultItem) GetParents()([]CommitSearchResultItem_parentsable) {
    return m.parents
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *CommitSearchResultItem) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetScore gets the score property value. The score property
// returns a *float64 when successful
func (m *CommitSearchResultItem) GetScore()(*float64) {
    return m.score
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *CommitSearchResultItem) GetSha()(*string) {
    return m.sha
}
// GetTextMatches gets the text_matches property value. The text_matches property
// returns a []Commitsable when successful
func (m *CommitSearchResultItem) GetTextMatches()([]Commitsable) {
    return m.text_matches
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CommitSearchResultItem) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CommitSearchResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
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
        err := writer.WriteObjectValue("commit", m.GetCommit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("committer", m.GetCommitter())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    if m.GetParents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParents()))
        for i, v := range m.GetParents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("parents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("score", m.GetScore())
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
    if m.GetTextMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTextMatches()))
        for i, v := range m.GetTextMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("text_matches", cast)
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
func (m *CommitSearchResultItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. A GitHub user.
func (m *CommitSearchResultItem) SetAuthor(value NullableSimpleUserable)() {
    m.author = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *CommitSearchResultItem) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommit sets the commit property value. The commit property
func (m *CommitSearchResultItem) SetCommit(value CommitSearchResultItem_commitable)() {
    m.commit = value
}
// SetCommitter sets the committer property value. Metaproperties for Git author/committer information.
func (m *CommitSearchResultItem) SetCommitter(value NullableGitUserable)() {
    m.committer = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *CommitSearchResultItem) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *CommitSearchResultItem) SetNodeId(value *string)() {
    m.node_id = value
}
// SetParents sets the parents property value. The parents property
func (m *CommitSearchResultItem) SetParents(value []CommitSearchResultItem_parentsable)() {
    m.parents = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *CommitSearchResultItem) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetScore sets the score property value. The score property
func (m *CommitSearchResultItem) SetScore(value *float64)() {
    m.score = value
}
// SetSha sets the sha property value. The sha property
func (m *CommitSearchResultItem) SetSha(value *string)() {
    m.sha = value
}
// SetTextMatches sets the text_matches property value. The text_matches property
func (m *CommitSearchResultItem) SetTextMatches(value []Commitsable)() {
    m.text_matches = value
}
// SetUrl sets the url property value. The url property
func (m *CommitSearchResultItem) SetUrl(value *string)() {
    m.url = value
}
type CommitSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetCommentsUrl()(*string)
    GetCommit()(CommitSearchResultItem_commitable)
    GetCommitter()(NullableGitUserable)
    GetHtmlUrl()(*string)
    GetNodeId()(*string)
    GetParents()([]CommitSearchResultItem_parentsable)
    GetRepository()(MinimalRepositoryable)
    GetScore()(*float64)
    GetSha()(*string)
    GetTextMatches()([]Commitsable)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetCommentsUrl(value *string)()
    SetCommit(value CommitSearchResultItem_commitable)()
    SetCommitter(value NullableGitUserable)()
    SetHtmlUrl(value *string)()
    SetNodeId(value *string)()
    SetParents(value []CommitSearchResultItem_parentsable)()
    SetRepository(value MinimalRepositoryable)()
    SetScore(value *float64)()
    SetSha(value *string)()
    SetTextMatches(value []Commitsable)()
    SetUrl(value *string)()
}
