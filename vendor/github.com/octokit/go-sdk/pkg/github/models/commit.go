package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Commit commit
type Commit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    author NullableSimpleUserable
    // The comments_url property
    comments_url *string
    // The commit property
    commit Commit_commitable
    // A GitHub user.
    committer NullableSimpleUserable
    // The files property
    files []DiffEntryable
    // The html_url property
    html_url *string
    // The node_id property
    node_id *string
    // The parents property
    parents []Commit_parentsable
    // The sha property
    sha *string
    // The stats property
    stats Commit_statsable
    // The url property
    url *string
}
// NewCommit instantiates a new Commit and sets the default values.
func NewCommit()(*Commit) {
    m := &Commit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Commit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Commit) GetAuthor()(NullableSimpleUserable) {
    return m.author
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *Commit) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommit gets the commit property value. The commit property
// returns a Commit_commitable when successful
func (m *Commit) GetCommit()(Commit_commitable) {
    return m.commit
}
// GetCommitter gets the committer property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Commit) GetCommitter()(NullableSimpleUserable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Commit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateCommit_commitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(Commit_commitable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiffEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiffEntryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DiffEntryable)
                }
            }
            m.SetFiles(res)
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
        val, err := n.GetCollectionOfObjectValues(CreateCommit_parentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Commit_parentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Commit_parentsable)
                }
            }
            m.SetParents(res)
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
    res["stats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommit_statsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStats(val.(Commit_statsable))
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
// GetFiles gets the files property value. The files property
// returns a []DiffEntryable when successful
func (m *Commit) GetFiles()([]DiffEntryable) {
    return m.files
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Commit) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Commit) GetNodeId()(*string) {
    return m.node_id
}
// GetParents gets the parents property value. The parents property
// returns a []Commit_parentsable when successful
func (m *Commit) GetParents()([]Commit_parentsable) {
    return m.parents
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *Commit) GetSha()(*string) {
    return m.sha
}
// GetStats gets the stats property value. The stats property
// returns a Commit_statsable when successful
func (m *Commit) GetStats()(Commit_statsable) {
    return m.stats
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Commit) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Commit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("files", cast)
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
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("stats", m.GetStats())
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
func (m *Commit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. A GitHub user.
func (m *Commit) SetAuthor(value NullableSimpleUserable)() {
    m.author = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *Commit) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommit sets the commit property value. The commit property
func (m *Commit) SetCommit(value Commit_commitable)() {
    m.commit = value
}
// SetCommitter sets the committer property value. A GitHub user.
func (m *Commit) SetCommitter(value NullableSimpleUserable)() {
    m.committer = value
}
// SetFiles sets the files property value. The files property
func (m *Commit) SetFiles(value []DiffEntryable)() {
    m.files = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Commit) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Commit) SetNodeId(value *string)() {
    m.node_id = value
}
// SetParents sets the parents property value. The parents property
func (m *Commit) SetParents(value []Commit_parentsable)() {
    m.parents = value
}
// SetSha sets the sha property value. The sha property
func (m *Commit) SetSha(value *string)() {
    m.sha = value
}
// SetStats sets the stats property value. The stats property
func (m *Commit) SetStats(value Commit_statsable)() {
    m.stats = value
}
// SetUrl sets the url property value. The url property
func (m *Commit) SetUrl(value *string)() {
    m.url = value
}
type Commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetCommentsUrl()(*string)
    GetCommit()(Commit_commitable)
    GetCommitter()(NullableSimpleUserable)
    GetFiles()([]DiffEntryable)
    GetHtmlUrl()(*string)
    GetNodeId()(*string)
    GetParents()([]Commit_parentsable)
    GetSha()(*string)
    GetStats()(Commit_statsable)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetCommentsUrl(value *string)()
    SetCommit(value Commit_commitable)()
    SetCommitter(value NullableSimpleUserable)()
    SetFiles(value []DiffEntryable)()
    SetHtmlUrl(value *string)()
    SetNodeId(value *string)()
    SetParents(value []Commit_parentsable)()
    SetSha(value *string)()
    SetStats(value Commit_statsable)()
    SetUrl(value *string)()
}
