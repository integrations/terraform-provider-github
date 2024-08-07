package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemReleasesItemWithRelease_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Text describing the contents of the tag.
    body *string
    // If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. If there is already a discussion linked to the release, this parameter is ignored. For more information, see "[Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository)."
    discussion_category_name *string
    // `true` makes the release a draft, and `false` publishes the release.
    draft *bool
    // The name of the release.
    name *string
    // `true` to identify the release as a prerelease, `false` to identify the release as a full release.
    prerelease *bool
    // The name of the tag.
    tag_name *string
    // Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch.
    target_commitish *string
}
// NewItemItemReleasesItemWithRelease_PatchRequestBody instantiates a new ItemItemReleasesItemWithRelease_PatchRequestBody and sets the default values.
func NewItemItemReleasesItemWithRelease_PatchRequestBody()(*ItemItemReleasesItemWithRelease_PatchRequestBody) {
    m := &ItemItemReleasesItemWithRelease_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemReleasesItemWithRelease_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemReleasesItemWithRelease_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemReleasesItemWithRelease_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. Text describing the contents of the tag.
// returns a *string when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetBody()(*string) {
    return m.body
}
// GetDiscussionCategoryName gets the discussion_category_name property value. If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. If there is already a discussion linked to the release, this parameter is ignored. For more information, see "[Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository)."
// returns a *string when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetDiscussionCategoryName()(*string) {
    return m.discussion_category_name
}
// GetDraft gets the draft property value. `true` makes the release a draft, and `false` publishes the release.
// returns a *bool when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetDraft()(*bool) {
    return m.draft
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["discussion_category_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscussionCategoryName(val)
        }
        return nil
    }
    res["draft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraft(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["prerelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrerelease(val)
        }
        return nil
    }
    res["tag_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagName(val)
        }
        return nil
    }
    res["target_commitish"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetCommitish(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the release.
// returns a *string when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetPrerelease gets the prerelease property value. `true` to identify the release as a prerelease, `false` to identify the release as a full release.
// returns a *bool when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetPrerelease()(*bool) {
    return m.prerelease
}
// GetTagName gets the tag_name property value. The name of the tag.
// returns a *string when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetTagName()(*string) {
    return m.tag_name
}
// GetTargetCommitish gets the target_commitish property value. Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch.
// returns a *string when successful
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) GetTargetCommitish()(*string) {
    return m.target_commitish
}
// Serialize serializes information the current object
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("discussion_category_name", m.GetDiscussionCategoryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("draft", m.GetDraft())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("prerelease", m.GetPrerelease())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tag_name", m.GetTagName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_commitish", m.GetTargetCommitish())
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
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. Text describing the contents of the tag.
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetDiscussionCategoryName sets the discussion_category_name property value. If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. If there is already a discussion linked to the release, this parameter is ignored. For more information, see "[Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository)."
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetDiscussionCategoryName(value *string)() {
    m.discussion_category_name = value
}
// SetDraft sets the draft property value. `true` makes the release a draft, and `false` publishes the release.
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetDraft(value *bool)() {
    m.draft = value
}
// SetName sets the name property value. The name of the release.
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPrerelease sets the prerelease property value. `true` to identify the release as a prerelease, `false` to identify the release as a full release.
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetPrerelease(value *bool)() {
    m.prerelease = value
}
// SetTagName sets the tag_name property value. The name of the tag.
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetTagName(value *string)() {
    m.tag_name = value
}
// SetTargetCommitish sets the target_commitish property value. Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch.
func (m *ItemItemReleasesItemWithRelease_PatchRequestBody) SetTargetCommitish(value *string)() {
    m.target_commitish = value
}
type ItemItemReleasesItemWithRelease_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetDiscussionCategoryName()(*string)
    GetDraft()(*bool)
    GetName()(*string)
    GetPrerelease()(*bool)
    GetTagName()(*string)
    GetTargetCommitish()(*string)
    SetBody(value *string)()
    SetDiscussionCategoryName(value *string)()
    SetDraft(value *bool)()
    SetName(value *string)()
    SetPrerelease(value *bool)()
    SetTagName(value *string)()
    SetTargetCommitish(value *string)()
}
