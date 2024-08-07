package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CommitSearchResultItem_commit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The author property
    author CommitSearchResultItem_commit_authorable
    // The comment_count property
    comment_count *int32
    // Metaproperties for Git author/committer information.
    committer NullableGitUserable
    // The message property
    message *string
    // The tree property
    tree CommitSearchResultItem_commit_treeable
    // The url property
    url *string
    // The verification property
    verification Verificationable
}
// NewCommitSearchResultItem_commit instantiates a new CommitSearchResultItem_commit and sets the default values.
func NewCommitSearchResultItem_commit()(*CommitSearchResultItem_commit) {
    m := &CommitSearchResultItem_commit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommitSearchResultItem_commitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommitSearchResultItem_commitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommitSearchResultItem_commit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommitSearchResultItem_commit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. The author property
// returns a CommitSearchResultItem_commit_authorable when successful
func (m *CommitSearchResultItem_commit) GetAuthor()(CommitSearchResultItem_commit_authorable) {
    return m.author
}
// GetCommentCount gets the comment_count property value. The comment_count property
// returns a *int32 when successful
func (m *CommitSearchResultItem_commit) GetCommentCount()(*int32) {
    return m.comment_count
}
// GetCommitter gets the committer property value. Metaproperties for Git author/committer information.
// returns a NullableGitUserable when successful
func (m *CommitSearchResultItem_commit) GetCommitter()(NullableGitUserable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommitSearchResultItem_commit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommitSearchResultItem_commit_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(CommitSearchResultItem_commit_authorable))
        }
        return nil
    }
    res["comment_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentCount(val)
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
    res["tree"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommitSearchResultItem_commit_treeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTree(val.(CommitSearchResultItem_commit_treeable))
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
        val, err := n.GetObjectValue(CreateVerificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerification(val.(Verificationable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *CommitSearchResultItem_commit) GetMessage()(*string) {
    return m.message
}
// GetTree gets the tree property value. The tree property
// returns a CommitSearchResultItem_commit_treeable when successful
func (m *CommitSearchResultItem_commit) GetTree()(CommitSearchResultItem_commit_treeable) {
    return m.tree
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CommitSearchResultItem_commit) GetUrl()(*string) {
    return m.url
}
// GetVerification gets the verification property value. The verification property
// returns a Verificationable when successful
func (m *CommitSearchResultItem_commit) GetVerification()(Verificationable) {
    return m.verification
}
// Serialize serializes information the current object
func (m *CommitSearchResultItem_commit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("comment_count", m.GetCommentCount())
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
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *CommitSearchResultItem_commit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. The author property
func (m *CommitSearchResultItem_commit) SetAuthor(value CommitSearchResultItem_commit_authorable)() {
    m.author = value
}
// SetCommentCount sets the comment_count property value. The comment_count property
func (m *CommitSearchResultItem_commit) SetCommentCount(value *int32)() {
    m.comment_count = value
}
// SetCommitter sets the committer property value. Metaproperties for Git author/committer information.
func (m *CommitSearchResultItem_commit) SetCommitter(value NullableGitUserable)() {
    m.committer = value
}
// SetMessage sets the message property value. The message property
func (m *CommitSearchResultItem_commit) SetMessage(value *string)() {
    m.message = value
}
// SetTree sets the tree property value. The tree property
func (m *CommitSearchResultItem_commit) SetTree(value CommitSearchResultItem_commit_treeable)() {
    m.tree = value
}
// SetUrl sets the url property value. The url property
func (m *CommitSearchResultItem_commit) SetUrl(value *string)() {
    m.url = value
}
// SetVerification sets the verification property value. The verification property
func (m *CommitSearchResultItem_commit) SetVerification(value Verificationable)() {
    m.verification = value
}
type CommitSearchResultItem_commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(CommitSearchResultItem_commit_authorable)
    GetCommentCount()(*int32)
    GetCommitter()(NullableGitUserable)
    GetMessage()(*string)
    GetTree()(CommitSearchResultItem_commit_treeable)
    GetUrl()(*string)
    GetVerification()(Verificationable)
    SetAuthor(value CommitSearchResultItem_commit_authorable)()
    SetCommentCount(value *int32)()
    SetCommitter(value NullableGitUserable)()
    SetMessage(value *string)()
    SetTree(value CommitSearchResultItem_commit_treeable)()
    SetUrl(value *string)()
    SetVerification(value Verificationable)()
}
