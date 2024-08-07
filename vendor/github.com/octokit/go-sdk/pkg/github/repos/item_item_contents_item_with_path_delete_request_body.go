package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemContentsItemWithPathDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // object containing information about the author.
    author ItemItemContentsItemWithPathDeleteRequestBody_authorable
    // The branch name. Default: the repository’s default branch
    branch *string
    // object containing information about the committer.
    committer ItemItemContentsItemWithPathDeleteRequestBody_committerable
    // The commit message.
    message *string
    // The blob SHA of the file being deleted.
    sha *string
}
// NewItemItemContentsItemWithPathDeleteRequestBody instantiates a new ItemItemContentsItemWithPathDeleteRequestBody and sets the default values.
func NewItemItemContentsItemWithPathDeleteRequestBody()(*ItemItemContentsItemWithPathDeleteRequestBody) {
    m := &ItemItemContentsItemWithPathDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemContentsItemWithPathDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemContentsItemWithPathDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemContentsItemWithPathDeleteRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. object containing information about the author.
// returns a ItemItemContentsItemWithPathDeleteRequestBody_authorable when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetAuthor()(ItemItemContentsItemWithPathDeleteRequestBody_authorable) {
    return m.author
}
// GetBranch gets the branch property value. The branch name. Default: the repository’s default branch
// returns a *string when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetBranch()(*string) {
    return m.branch
}
// GetCommitter gets the committer property value. object containing information about the committer.
// returns a ItemItemContentsItemWithPathDeleteRequestBody_committerable when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetCommitter()(ItemItemContentsItemWithPathDeleteRequestBody_committerable) {
    return m.committer
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemContentsItemWithPathDeleteRequestBody_authorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(ItemItemContentsItemWithPathDeleteRequestBody_authorable))
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
        val, err := n.GetObjectValue(CreateItemItemContentsItemWithPathDeleteRequestBody_committerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitter(val.(ItemItemContentsItemWithPathDeleteRequestBody_committerable))
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
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetMessage()(*string) {
    return m.message
}
// GetSha gets the sha property value. The blob SHA of the file being deleted.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *ItemItemContentsItemWithPathDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemContentsItemWithPathDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. object containing information about the author.
func (m *ItemItemContentsItemWithPathDeleteRequestBody) SetAuthor(value ItemItemContentsItemWithPathDeleteRequestBody_authorable)() {
    m.author = value
}
// SetBranch sets the branch property value. The branch name. Default: the repository’s default branch
func (m *ItemItemContentsItemWithPathDeleteRequestBody) SetBranch(value *string)() {
    m.branch = value
}
// SetCommitter sets the committer property value. object containing information about the committer.
func (m *ItemItemContentsItemWithPathDeleteRequestBody) SetCommitter(value ItemItemContentsItemWithPathDeleteRequestBody_committerable)() {
    m.committer = value
}
// SetMessage sets the message property value. The commit message.
func (m *ItemItemContentsItemWithPathDeleteRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetSha sets the sha property value. The blob SHA of the file being deleted.
func (m *ItemItemContentsItemWithPathDeleteRequestBody) SetSha(value *string)() {
    m.sha = value
}
type ItemItemContentsItemWithPathDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(ItemItemContentsItemWithPathDeleteRequestBody_authorable)
    GetBranch()(*string)
    GetCommitter()(ItemItemContentsItemWithPathDeleteRequestBody_committerable)
    GetMessage()(*string)
    GetSha()(*string)
    SetAuthor(value ItemItemContentsItemWithPathDeleteRequestBody_authorable)()
    SetBranch(value *string)()
    SetCommitter(value ItemItemContentsItemWithPathDeleteRequestBody_committerable)()
    SetMessage(value *string)()
    SetSha(value *string)()
}
