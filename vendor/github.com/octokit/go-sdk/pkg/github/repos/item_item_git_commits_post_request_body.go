package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemGitCommitsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Information about the author of the commit. By default, the `author` will be the authenticated user and the current date. See the `author` and `committer` object below for details.
    author ItemItemGitCommitsPostRequestBody_authorable
    // Information about the person who is making the commit. By default, `committer` will use the information set in `author`. See the `author` and `committer` object below for details.
    committer ItemItemGitCommitsPostRequestBody_committerable
    // The commit message
    message *string
    // The SHAs of the commits that were the parents of this commit. If omitted or empty, the commit will be written as a root commit. For a single parent, an array of one SHA should be provided; for a merge commit, an array of more than one should be provided.
    parents []string
    // The [PGP signature](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) of the commit. GitHub adds the signature to the `gpgsig` header of the created commit. For a commit signature to be verifiable by Git or GitHub, it must be an ASCII-armored detached PGP signature over the string commit as it would be written to the object database. To pass a `signature` parameter, you need to first manually create a valid PGP signature, which can be complicated. You may find it easier to [use the command line](https://git-scm.com/book/id/v2/Git-Tools-Signing-Your-Work) to create signed commits.
    signature *string
    // The SHA of the tree object this commit points to
    tree *string
}
// NewItemItemGitCommitsPostRequestBody instantiates a new ItemItemGitCommitsPostRequestBody and sets the default values.
func NewItemItemGitCommitsPostRequestBody()(*ItemItemGitCommitsPostRequestBody) {
    m := &ItemItemGitCommitsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGitCommitsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGitCommitsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitCommitsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGitCommitsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. Information about the author of the commit. By default, the `author` will be the authenticated user and the current date. See the `author` and `committer` object below for details.
// returns a ItemItemGitCommitsPostRequestBody_authorable when successful
func (m *ItemItemGitCommitsPostRequestBody) GetAuthor()(ItemItemGitCommitsPostRequestBody_authorable) {
    return m.author
}
// GetCommitter gets the committer property value. Information about the person who is making the commit. By default, `committer` will use the information set in `author`. See the `author` and `committer` object below for details.
// returns a ItemItemGitCommitsPostRequestBody_committerable when successful
func (m *ItemItemGitCommitsPostRequestBody) GetCommitter()(ItemItemGitCommitsPostRequestBody_committerable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGitCommitsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemGitCommitsPostRequestBody_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(ItemItemGitCommitsPostRequestBody_authorable))
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemGitCommitsPostRequestBody_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(ItemItemGitCommitsPostRequestBody_committerable))
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
    res["parents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetParents(res)
        }
        return nil
    }
    res["signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val)
        }
        return nil
    }
    res["tree"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTree(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The commit message
// returns a *string when successful
func (m *ItemItemGitCommitsPostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetParents gets the parents property value. The SHAs of the commits that were the parents of this commit. If omitted or empty, the commit will be written as a root commit. For a single parent, an array of one SHA should be provided; for a merge commit, an array of more than one should be provided.
// returns a []string when successful
func (m *ItemItemGitCommitsPostRequestBody) GetParents()([]string) {
    return m.parents
}
// GetSignature gets the signature property value. The [PGP signature](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) of the commit. GitHub adds the signature to the `gpgsig` header of the created commit. For a commit signature to be verifiable by Git or GitHub, it must be an ASCII-armored detached PGP signature over the string commit as it would be written to the object database. To pass a `signature` parameter, you need to first manually create a valid PGP signature, which can be complicated. You may find it easier to [use the command line](https://git-scm.com/book/id/v2/Git-Tools-Signing-Your-Work) to create signed commits.
// returns a *string when successful
func (m *ItemItemGitCommitsPostRequestBody) GetSignature()(*string) {
    return m.signature
}
// GetTree gets the tree property value. The SHA of the tree object this commit points to
// returns a *string when successful
func (m *ItemItemGitCommitsPostRequestBody) GetTree()(*string) {
    return m.tree
}
// Serialize serializes information the current object
func (m *ItemItemGitCommitsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetParents() != nil {
        err := writer.WriteCollectionOfStringValues("parents", m.GetParents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signature", m.GetSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tree", m.GetTree())
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
func (m *ItemItemGitCommitsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. Information about the author of the commit. By default, the `author` will be the authenticated user and the current date. See the `author` and `committer` object below for details.
func (m *ItemItemGitCommitsPostRequestBody) SetAuthor(value ItemItemGitCommitsPostRequestBody_authorable)() {
    m.author = value
}
// SetCommitter sets the committer property value. Information about the person who is making the commit. By default, `committer` will use the information set in `author`. See the `author` and `committer` object below for details.
func (m *ItemItemGitCommitsPostRequestBody) SetCommitter(value ItemItemGitCommitsPostRequestBody_committerable)() {
    m.committer = value
}
// SetMessage sets the message property value. The commit message
func (m *ItemItemGitCommitsPostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetParents sets the parents property value. The SHAs of the commits that were the parents of this commit. If omitted or empty, the commit will be written as a root commit. For a single parent, an array of one SHA should be provided; for a merge commit, an array of more than one should be provided.
func (m *ItemItemGitCommitsPostRequestBody) SetParents(value []string)() {
    m.parents = value
}
// SetSignature sets the signature property value. The [PGP signature](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) of the commit. GitHub adds the signature to the `gpgsig` header of the created commit. For a commit signature to be verifiable by Git or GitHub, it must be an ASCII-armored detached PGP signature over the string commit as it would be written to the object database. To pass a `signature` parameter, you need to first manually create a valid PGP signature, which can be complicated. You may find it easier to [use the command line](https://git-scm.com/book/id/v2/Git-Tools-Signing-Your-Work) to create signed commits.
func (m *ItemItemGitCommitsPostRequestBody) SetSignature(value *string)() {
    m.signature = value
}
// SetTree sets the tree property value. The SHA of the tree object this commit points to
func (m *ItemItemGitCommitsPostRequestBody) SetTree(value *string)() {
    m.tree = value
}
type ItemItemGitCommitsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(ItemItemGitCommitsPostRequestBody_authorable)
    GetCommitter()(ItemItemGitCommitsPostRequestBody_committerable)
    GetMessage()(*string)
    GetParents()([]string)
    GetSignature()(*string)
    GetTree()(*string)
    SetAuthor(value ItemItemGitCommitsPostRequestBody_authorable)()
    SetCommitter(value ItemItemGitCommitsPostRequestBody_committerable)()
    SetMessage(value *string)()
    SetParents(value []string)()
    SetSignature(value *string)()
    SetTree(value *string)()
}
