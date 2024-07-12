package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemMergesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the base branch that the head will be merged into.
    base *string
    // Commit message to use for the merge commit. If omitted, a default message will be used.
    commit_message *string
    // The head to merge. This can be a branch name or a commit SHA1.
    head *string
}
// NewItemItemMergesPostRequestBody instantiates a new ItemItemMergesPostRequestBody and sets the default values.
func NewItemItemMergesPostRequestBody()(*ItemItemMergesPostRequestBody) {
    m := &ItemItemMergesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemMergesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemMergesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemMergesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemMergesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBase gets the base property value. The name of the base branch that the head will be merged into.
// returns a *string when successful
func (m *ItemItemMergesPostRequestBody) GetBase()(*string) {
    return m.base
}
// GetCommitMessage gets the commit_message property value. Commit message to use for the merge commit. If omitted, a default message will be used.
// returns a *string when successful
func (m *ItemItemMergesPostRequestBody) GetCommitMessage()(*string) {
    return m.commit_message
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemMergesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["base"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBase(val)
        }
        return nil
    }
    res["commit_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitMessage(val)
        }
        return nil
    }
    res["head"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHead(val)
        }
        return nil
    }
    return res
}
// GetHead gets the head property value. The head to merge. This can be a branch name or a commit SHA1.
// returns a *string when successful
func (m *ItemItemMergesPostRequestBody) GetHead()(*string) {
    return m.head
}
// Serialize serializes information the current object
func (m *ItemItemMergesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_message", m.GetCommitMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("head", m.GetHead())
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
func (m *ItemItemMergesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBase sets the base property value. The name of the base branch that the head will be merged into.
func (m *ItemItemMergesPostRequestBody) SetBase(value *string)() {
    m.base = value
}
// SetCommitMessage sets the commit_message property value. Commit message to use for the merge commit. If omitted, a default message will be used.
func (m *ItemItemMergesPostRequestBody) SetCommitMessage(value *string)() {
    m.commit_message = value
}
// SetHead sets the head property value. The head to merge. This can be a branch name or a commit SHA1.
func (m *ItemItemMergesPostRequestBody) SetHead(value *string)() {
    m.head = value
}
type ItemItemMergesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBase()(*string)
    GetCommitMessage()(*string)
    GetHead()(*string)
    SetBase(value *string)()
    SetCommitMessage(value *string)()
    SetHead(value *string)()
}
