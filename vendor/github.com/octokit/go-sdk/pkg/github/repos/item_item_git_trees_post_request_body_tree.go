package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemGitTreesPostRequestBody_tree struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content you want this file to have. GitHub will write this blob out and use that SHA for this entry. Use either this, or `tree.sha`.    **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
    content *string
    // The file referenced in the tree.
    path *string
    // The SHA1 checksum ID of the object in the tree. Also called `tree.sha`. If the value is `null` then the file will be deleted.    **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
    sha *string
}
// NewItemItemGitTreesPostRequestBody_tree instantiates a new ItemItemGitTreesPostRequestBody_tree and sets the default values.
func NewItemItemGitTreesPostRequestBody_tree()(*ItemItemGitTreesPostRequestBody_tree) {
    m := &ItemItemGitTreesPostRequestBody_tree{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGitTreesPostRequestBody_treeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGitTreesPostRequestBody_treeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitTreesPostRequestBody_tree(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGitTreesPostRequestBody_tree) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content you want this file to have. GitHub will write this blob out and use that SHA for this entry. Use either this, or `tree.sha`.    **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
// returns a *string when successful
func (m *ItemItemGitTreesPostRequestBody_tree) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGitTreesPostRequestBody_tree) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
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
// GetPath gets the path property value. The file referenced in the tree.
// returns a *string when successful
func (m *ItemItemGitTreesPostRequestBody_tree) GetPath()(*string) {
    return m.path
}
// GetSha gets the sha property value. The SHA1 checksum ID of the object in the tree. Also called `tree.sha`. If the value is `null` then the file will be deleted.    **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
// returns a *string when successful
func (m *ItemItemGitTreesPostRequestBody_tree) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *ItemItemGitTreesPostRequestBody_tree) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
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
func (m *ItemItemGitTreesPostRequestBody_tree) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content you want this file to have. GitHub will write this blob out and use that SHA for this entry. Use either this, or `tree.sha`.    **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
func (m *ItemItemGitTreesPostRequestBody_tree) SetContent(value *string)() {
    m.content = value
}
// SetPath sets the path property value. The file referenced in the tree.
func (m *ItemItemGitTreesPostRequestBody_tree) SetPath(value *string)() {
    m.path = value
}
// SetSha sets the sha property value. The SHA1 checksum ID of the object in the tree. Also called `tree.sha`. If the value is `null` then the file will be deleted.    **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
func (m *ItemItemGitTreesPostRequestBody_tree) SetSha(value *string)() {
    m.sha = value
}
type ItemItemGitTreesPostRequestBody_treeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetPath()(*string)
    GetSha()(*string)
    SetContent(value *string)()
    SetPath(value *string)()
    SetSha(value *string)()
}
