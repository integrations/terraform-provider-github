package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemMergePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Extra detail to append to automatic commit message.
    commit_message *string
    // Title for the automatic commit message.
    commit_title *string
    // SHA that pull request head must match to allow merge.
    sha *string
}
// NewItemItemPullsItemMergePutRequestBody instantiates a new ItemItemPullsItemMergePutRequestBody and sets the default values.
func NewItemItemPullsItemMergePutRequestBody()(*ItemItemPullsItemMergePutRequestBody) {
    m := &ItemItemPullsItemMergePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemMergePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemMergePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemMergePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemMergePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitMessage gets the commit_message property value. Extra detail to append to automatic commit message.
// returns a *string when successful
func (m *ItemItemPullsItemMergePutRequestBody) GetCommitMessage()(*string) {
    return m.commit_message
}
// GetCommitTitle gets the commit_title property value. Title for the automatic commit message.
// returns a *string when successful
func (m *ItemItemPullsItemMergePutRequestBody) GetCommitTitle()(*string) {
    return m.commit_title
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemMergePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["commit_title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitTitle(val)
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
// GetSha gets the sha property value. SHA that pull request head must match to allow merge.
// returns a *string when successful
func (m *ItemItemPullsItemMergePutRequestBody) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemMergePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("commit_message", m.GetCommitMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_title", m.GetCommitTitle())
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
func (m *ItemItemPullsItemMergePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitMessage sets the commit_message property value. Extra detail to append to automatic commit message.
func (m *ItemItemPullsItemMergePutRequestBody) SetCommitMessage(value *string)() {
    m.commit_message = value
}
// SetCommitTitle sets the commit_title property value. Title for the automatic commit message.
func (m *ItemItemPullsItemMergePutRequestBody) SetCommitTitle(value *string)() {
    m.commit_title = value
}
// SetSha sets the sha property value. SHA that pull request head must match to allow merge.
func (m *ItemItemPullsItemMergePutRequestBody) SetSha(value *string)() {
    m.sha = value
}
type ItemItemPullsItemMergePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommitMessage()(*string)
    GetCommitTitle()(*string)
    GetSha()(*string)
    SetCommitMessage(value *string)()
    SetCommitTitle(value *string)()
    SetSha(value *string)()
}
