package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitCommit low-level Git commit operations within a repository
type GitCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifying information for the git-user
    author GitCommit_authorable
    // Identifying information for the git-user
    committer GitCommit_committerable
    // The html_url property
    html_url *string
    // Message describing the purpose of the commit
    message *string
    // The node_id property
    node_id *string
    // The parents property
    parents []GitCommit_parentsable
    // SHA for the commit
    sha *string
    // The tree property
    tree GitCommit_treeable
    // The url property
    url *string
    // The verification property
    verification GitCommit_verificationable
}
// NewGitCommit instantiates a new GitCommit and sets the default values.
func NewGitCommit()(*GitCommit) {
    m := &GitCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GitCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. Identifying information for the git-user
// returns a GitCommit_authorable when successful
func (m *GitCommit) GetAuthor()(GitCommit_authorable) {
    return m.author
}
// GetCommitter gets the committer property value. Identifying information for the git-user
// returns a GitCommit_committerable when successful
func (m *GitCommit) GetCommitter()(GitCommit_committerable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitCommit_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(GitCommit_authorable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitCommit_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(GitCommit_committerable))
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
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateGitCommit_parentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GitCommit_parentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GitCommit_parentsable)
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
    res["tree"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitCommit_treeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTree(val.(GitCommit_treeable))
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
    res["verification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitCommit_verificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerification(val.(GitCommit_verificationable))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *GitCommit) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetMessage gets the message property value. Message describing the purpose of the commit
// returns a *string when successful
func (m *GitCommit) GetMessage()(*string) {
    return m.message
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *GitCommit) GetNodeId()(*string) {
    return m.node_id
}
// GetParents gets the parents property value. The parents property
// returns a []GitCommit_parentsable when successful
func (m *GitCommit) GetParents()([]GitCommit_parentsable) {
    return m.parents
}
// GetSha gets the sha property value. SHA for the commit
// returns a *string when successful
func (m *GitCommit) GetSha()(*string) {
    return m.sha
}
// GetTree gets the tree property value. The tree property
// returns a GitCommit_treeable when successful
func (m *GitCommit) GetTree()(GitCommit_treeable) {
    return m.tree
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GitCommit) GetUrl()(*string) {
    return m.url
}
// GetVerification gets the verification property value. The verification property
// returns a GitCommit_verificationable when successful
func (m *GitCommit) GetVerification()(GitCommit_verificationable) {
    return m.verification
}
// Serialize serializes information the current object
func (m *GitCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
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
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteObjectValue("tree", m.GetTree())
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
        err := writer.WriteObjectValue("verification", m.GetVerification())
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
func (m *GitCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. Identifying information for the git-user
func (m *GitCommit) SetAuthor(value GitCommit_authorable)() {
    m.author = value
}
// SetCommitter sets the committer property value. Identifying information for the git-user
func (m *GitCommit) SetCommitter(value GitCommit_committerable)() {
    m.committer = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *GitCommit) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetMessage sets the message property value. Message describing the purpose of the commit
func (m *GitCommit) SetMessage(value *string)() {
    m.message = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *GitCommit) SetNodeId(value *string)() {
    m.node_id = value
}
// SetParents sets the parents property value. The parents property
func (m *GitCommit) SetParents(value []GitCommit_parentsable)() {
    m.parents = value
}
// SetSha sets the sha property value. SHA for the commit
func (m *GitCommit) SetSha(value *string)() {
    m.sha = value
}
// SetTree sets the tree property value. The tree property
func (m *GitCommit) SetTree(value GitCommit_treeable)() {
    m.tree = value
}
// SetUrl sets the url property value. The url property
func (m *GitCommit) SetUrl(value *string)() {
    m.url = value
}
// SetVerification sets the verification property value. The verification property
func (m *GitCommit) SetVerification(value GitCommit_verificationable)() {
    m.verification = value
}
type GitCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(GitCommit_authorable)
    GetCommitter()(GitCommit_committerable)
    GetHtmlUrl()(*string)
    GetMessage()(*string)
    GetNodeId()(*string)
    GetParents()([]GitCommit_parentsable)
    GetSha()(*string)
    GetTree()(GitCommit_treeable)
    GetUrl()(*string)
    GetVerification()(GitCommit_verificationable)
    SetAuthor(value GitCommit_authorable)()
    SetCommitter(value GitCommit_committerable)()
    SetHtmlUrl(value *string)()
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetParents(value []GitCommit_parentsable)()
    SetSha(value *string)()
    SetTree(value GitCommit_treeable)()
    SetUrl(value *string)()
    SetVerification(value GitCommit_verificationable)()
}
