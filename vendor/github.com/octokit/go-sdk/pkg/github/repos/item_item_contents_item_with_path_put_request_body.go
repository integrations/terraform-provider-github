package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemContentsItemWithPathPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The author of the file. Default: The `committer` or the authenticated user if you omit `committer`.
    author ItemItemContentsItemWithPathPutRequestBody_authorable
    // The branch name. Default: the repository’s default branch.
    branch *string
    // The person that committed the file. Default: the authenticated user.
    committer ItemItemContentsItemWithPathPutRequestBody_committerable
    // The new file content, using Base64 encoding.
    content *string
    // The commit message.
    message *string
    // **Required if you are updating a file**. The blob SHA of the file being replaced.
    sha *string
}
// NewItemItemContentsItemWithPathPutRequestBody instantiates a new ItemItemContentsItemWithPathPutRequestBody and sets the default values.
func NewItemItemContentsItemWithPathPutRequestBody()(*ItemItemContentsItemWithPathPutRequestBody) {
    m := &ItemItemContentsItemWithPathPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemContentsItemWithPathPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemContentsItemWithPathPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemContentsItemWithPathPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. The author of the file. Default: The `committer` or the authenticated user if you omit `committer`.
// returns a ItemItemContentsItemWithPathPutRequestBody_authorable when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetAuthor()(ItemItemContentsItemWithPathPutRequestBody_authorable) {
    return m.author
}
// GetBranch gets the branch property value. The branch name. Default: the repository’s default branch.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetBranch()(*string) {
    return m.branch
}
// GetCommitter gets the committer property value. The person that committed the file. Default: the authenticated user.
// returns a ItemItemContentsItemWithPathPutRequestBody_committerable when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetCommitter()(ItemItemContentsItemWithPathPutRequestBody_committerable) {
    return m.committer
}
// GetContent gets the content property value. The new file content, using Base64 encoding.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemContentsItemWithPathPutRequestBody_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(ItemItemContentsItemWithPathPutRequestBody_authorable))
        }
        return nil
    }
    res["branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    res["committer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemContentsItemWithPathPutRequestBody_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(ItemItemContentsItemWithPathPutRequestBody_committerable))
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
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
    return res
}
// GetMessage gets the message property value. The commit message.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetMessage()(*string) {
    return m.message
}
// GetSha gets the sha property value. **Required if you are updating a file**. The blob SHA of the file being replaced.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *ItemItemContentsItemWithPathPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("branch", m.GetBranch())
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
        err := writer.WriteStringValue("content", m.GetContent())
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
        err := writer.WriteStringValue("sha", m.GetSha())
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
func (m *ItemItemContentsItemWithPathPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. The author of the file. Default: The `committer` or the authenticated user if you omit `committer`.
func (m *ItemItemContentsItemWithPathPutRequestBody) SetAuthor(value ItemItemContentsItemWithPathPutRequestBody_authorable)() {
    m.author = value
}
// SetBranch sets the branch property value. The branch name. Default: the repository’s default branch.
func (m *ItemItemContentsItemWithPathPutRequestBody) SetBranch(value *string)() {
    m.branch = value
}
// SetCommitter sets the committer property value. The person that committed the file. Default: the authenticated user.
func (m *ItemItemContentsItemWithPathPutRequestBody) SetCommitter(value ItemItemContentsItemWithPathPutRequestBody_committerable)() {
    m.committer = value
}
// SetContent sets the content property value. The new file content, using Base64 encoding.
func (m *ItemItemContentsItemWithPathPutRequestBody) SetContent(value *string)() {
    m.content = value
}
// SetMessage sets the message property value. The commit message.
func (m *ItemItemContentsItemWithPathPutRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetSha sets the sha property value. **Required if you are updating a file**. The blob SHA of the file being replaced.
func (m *ItemItemContentsItemWithPathPutRequestBody) SetSha(value *string)() {
    m.sha = value
}
type ItemItemContentsItemWithPathPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(ItemItemContentsItemWithPathPutRequestBody_authorable)
    GetBranch()(*string)
    GetCommitter()(ItemItemContentsItemWithPathPutRequestBody_committerable)
    GetContent()(*string)
    GetMessage()(*string)
    GetSha()(*string)
    SetAuthor(value ItemItemContentsItemWithPathPutRequestBody_authorable)()
    SetBranch(value *string)()
    SetCommitter(value ItemItemContentsItemWithPathPutRequestBody_committerable)()
    SetContent(value *string)()
    SetMessage(value *string)()
    SetSha(value *string)()
}
