package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AutoMerge the status of auto merging a pull request.
type AutoMerge struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Commit message for the merge commit.
    commit_message *string
    // Title for the merge commit message.
    commit_title *string
    // A GitHub user.
    enabled_by SimpleUserable
    // The merge method to use.
    merge_method *AutoMerge_merge_method
}
// NewAutoMerge instantiates a new AutoMerge and sets the default values.
func NewAutoMerge()(*AutoMerge) {
    m := &AutoMerge{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAutoMergeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutoMergeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutoMerge(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutoMerge) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitMessage gets the commit_message property value. Commit message for the merge commit.
// returns a *string when successful
func (m *AutoMerge) GetCommitMessage()(*string) {
    return m.commit_message
}
// GetCommitTitle gets the commit_title property value. Title for the merge commit message.
// returns a *string when successful
func (m *AutoMerge) GetCommitTitle()(*string) {
    return m.commit_title
}
// GetEnabledBy gets the enabled_by property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *AutoMerge) GetEnabledBy()(SimpleUserable) {
    return m.enabled_by
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutoMerge) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["enabled_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledBy(val.(SimpleUserable))
        }
        return nil
    }
    res["merge_method"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutoMerge_merge_method)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeMethod(val.(*AutoMerge_merge_method))
        }
        return nil
    }
    return res
}
// GetMergeMethod gets the merge_method property value. The merge method to use.
// returns a *AutoMerge_merge_method when successful
func (m *AutoMerge) GetMergeMethod()(*AutoMerge_merge_method) {
    return m.merge_method
}
// Serialize serializes information the current object
func (m *AutoMerge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("enabled_by", m.GetEnabledBy())
        if err != nil {
            return err
        }
    }
    if m.GetMergeMethod() != nil {
        cast := (*m.GetMergeMethod()).String()
        err := writer.WriteStringValue("merge_method", &cast)
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
func (m *AutoMerge) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitMessage sets the commit_message property value. Commit message for the merge commit.
func (m *AutoMerge) SetCommitMessage(value *string)() {
    m.commit_message = value
}
// SetCommitTitle sets the commit_title property value. Title for the merge commit message.
func (m *AutoMerge) SetCommitTitle(value *string)() {
    m.commit_title = value
}
// SetEnabledBy sets the enabled_by property value. A GitHub user.
func (m *AutoMerge) SetEnabledBy(value SimpleUserable)() {
    m.enabled_by = value
}
// SetMergeMethod sets the merge_method property value. The merge method to use.
func (m *AutoMerge) SetMergeMethod(value *AutoMerge_merge_method)() {
    m.merge_method = value
}
type AutoMergeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommitMessage()(*string)
    GetCommitTitle()(*string)
    GetEnabledBy()(SimpleUserable)
    GetMergeMethod()(*AutoMerge_merge_method)
    SetCommitMessage(value *string)()
    SetCommitTitle(value *string)()
    SetEnabledBy(value SimpleUserable)()
    SetMergeMethod(value *AutoMerge_merge_method)()
}
