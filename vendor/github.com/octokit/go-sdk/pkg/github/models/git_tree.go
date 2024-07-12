package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitTree the hierarchy between files in a Git repository.
type GitTree struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sha property
    sha *string
    // Objects specifying a tree structure
    tree []GitTree_treeable
    // The truncated property
    truncated *bool
    // The url property
    url *string
}
// NewGitTree instantiates a new GitTree and sets the default values.
func NewGitTree()(*GitTree) {
    m := &GitTree{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitTreeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitTreeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitTree(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GitTree) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitTree) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetCollectionOfObjectValues(CreateGitTree_treeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GitTree_treeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GitTree_treeable)
                }
            }
            m.SetTree(res)
        }
        return nil
    }
    res["truncated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTruncated(val)
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
    return res
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *GitTree) GetSha()(*string) {
    return m.sha
}
// GetTree gets the tree property value. Objects specifying a tree structure
// returns a []GitTree_treeable when successful
func (m *GitTree) GetTree()([]GitTree_treeable) {
    return m.tree
}
// GetTruncated gets the truncated property value. The truncated property
// returns a *bool when successful
func (m *GitTree) GetTruncated()(*bool) {
    return m.truncated
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GitTree) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *GitTree) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    if m.GetTree() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTree()))
        for i, v := range m.GetTree() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("tree", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("truncated", m.GetTruncated())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GitTree) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSha sets the sha property value. The sha property
func (m *GitTree) SetSha(value *string)() {
    m.sha = value
}
// SetTree sets the tree property value. Objects specifying a tree structure
func (m *GitTree) SetTree(value []GitTree_treeable)() {
    m.tree = value
}
// SetTruncated sets the truncated property value. The truncated property
func (m *GitTree) SetTruncated(value *bool)() {
    m.truncated = value
}
// SetUrl sets the url property value. The url property
func (m *GitTree) SetUrl(value *string)() {
    m.url = value
}
type GitTreeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSha()(*string)
    GetTree()([]GitTree_treeable)
    GetTruncated()(*bool)
    GetUrl()(*string)
    SetSha(value *string)()
    SetTree(value []GitTree_treeable)()
    SetTruncated(value *bool)()
    SetUrl(value *string)()
}
