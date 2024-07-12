package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileCommit_commit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The author property
    author FileCommit_commit_authorable
    // The committer property
    committer FileCommit_commit_committerable
    // The html_url property
    html_url *string
    // The message property
    message *string
    // The node_id property
    node_id *string
    // The parents property
    parents []FileCommit_commit_parentsable
    // The sha property
    sha *string
    // The tree property
    tree FileCommit_commit_treeable
    // The url property
    url *string
    // The verification property
    verification FileCommit_commit_verificationable
}
// NewFileCommit_commit instantiates a new FileCommit_commit and sets the default values.
func NewFileCommit_commit()(*FileCommit_commit) {
    m := &FileCommit_commit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileCommit_commitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileCommit_commitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileCommit_commit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileCommit_commit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. The author property
// returns a FileCommit_commit_authorable when successful
func (m *FileCommit_commit) GetAuthor()(FileCommit_commit_authorable) {
    return m.author
}
// GetCommitter gets the committer property value. The committer property
// returns a FileCommit_commit_committerable when successful
func (m *FileCommit_commit) GetCommitter()(FileCommit_commit_committerable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileCommit_commit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileCommit_commit_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(FileCommit_commit_authorable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileCommit_commit_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(FileCommit_commit_committerable))
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
        val, err := n.GetCollectionOfObjectValues(CreateFileCommit_commit_parentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileCommit_commit_parentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileCommit_commit_parentsable)
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
        val, err := n.GetObjectValue(CreateFileCommit_commit_treeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTree(val.(FileCommit_commit_treeable))
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
        val, err := n.GetObjectValue(CreateFileCommit_commit_verificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerification(val.(FileCommit_commit_verificationable))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *FileCommit_commit) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *FileCommit_commit) GetMessage()(*string) {
    return m.message
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *FileCommit_commit) GetNodeId()(*string) {
    return m.node_id
}
// GetParents gets the parents property value. The parents property
// returns a []FileCommit_commit_parentsable when successful
func (m *FileCommit_commit) GetParents()([]FileCommit_commit_parentsable) {
    return m.parents
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *FileCommit_commit) GetSha()(*string) {
    return m.sha
}
// GetTree gets the tree property value. The tree property
// returns a FileCommit_commit_treeable when successful
func (m *FileCommit_commit) GetTree()(FileCommit_commit_treeable) {
    return m.tree
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *FileCommit_commit) GetUrl()(*string) {
    return m.url
}
// GetVerification gets the verification property value. The verification property
// returns a FileCommit_commit_verificationable when successful
func (m *FileCommit_commit) GetVerification()(FileCommit_commit_verificationable) {
    return m.verification
}
// Serialize serializes information the current object
func (m *FileCommit_commit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *FileCommit_commit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. The author property
func (m *FileCommit_commit) SetAuthor(value FileCommit_commit_authorable)() {
    m.author = value
}
// SetCommitter sets the committer property value. The committer property
func (m *FileCommit_commit) SetCommitter(value FileCommit_commit_committerable)() {
    m.committer = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *FileCommit_commit) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetMessage sets the message property value. The message property
func (m *FileCommit_commit) SetMessage(value *string)() {
    m.message = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *FileCommit_commit) SetNodeId(value *string)() {
    m.node_id = value
}
// SetParents sets the parents property value. The parents property
func (m *FileCommit_commit) SetParents(value []FileCommit_commit_parentsable)() {
    m.parents = value
}
// SetSha sets the sha property value. The sha property
func (m *FileCommit_commit) SetSha(value *string)() {
    m.sha = value
}
// SetTree sets the tree property value. The tree property
func (m *FileCommit_commit) SetTree(value FileCommit_commit_treeable)() {
    m.tree = value
}
// SetUrl sets the url property value. The url property
func (m *FileCommit_commit) SetUrl(value *string)() {
    m.url = value
}
// SetVerification sets the verification property value. The verification property
func (m *FileCommit_commit) SetVerification(value FileCommit_commit_verificationable)() {
    m.verification = value
}
type FileCommit_commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(FileCommit_commit_authorable)
    GetCommitter()(FileCommit_commit_committerable)
    GetHtmlUrl()(*string)
    GetMessage()(*string)
    GetNodeId()(*string)
    GetParents()([]FileCommit_commit_parentsable)
    GetSha()(*string)
    GetTree()(FileCommit_commit_treeable)
    GetUrl()(*string)
    GetVerification()(FileCommit_commit_verificationable)
    SetAuthor(value FileCommit_commit_authorable)()
    SetCommitter(value FileCommit_commit_committerable)()
    SetHtmlUrl(value *string)()
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetParents(value []FileCommit_commit_parentsable)()
    SetSha(value *string)()
    SetTree(value FileCommit_commit_treeable)()
    SetUrl(value *string)()
    SetVerification(value FileCommit_commit_verificationable)()
}
