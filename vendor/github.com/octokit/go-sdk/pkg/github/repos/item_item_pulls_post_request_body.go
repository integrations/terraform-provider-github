package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository.
    base *string
    // The contents of the pull request.
    body *string
    // Indicates whether the pull request is a draft. See "[Draft Pull Requests](https://docs.github.com/articles/about-pull-requests#draft-pull-requests)" in the GitHub Help documentation to learn more.
    draft *bool
    // The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.
    head *string
    // The name of the repository where the changes in the pull request were made. This field is required for cross-repository pull requests if both repositories are owned by the same organization.
    head_repo *string
    // An issue in the repository to convert to a pull request. The issue title, body, and comments will become the title, body, and comments on the new pull request. Required unless `title` is specified.
    issue *int64
    // Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request.
    maintainer_can_modify *bool
    // The title of the new pull request. Required unless `issue` is specified.
    title *string
}
// NewItemItemPullsPostRequestBody instantiates a new ItemItemPullsPostRequestBody and sets the default values.
func NewItemItemPullsPostRequestBody()(*ItemItemPullsPostRequestBody) {
    m := &ItemItemPullsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBase gets the base property value. The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository.
// returns a *string when successful
func (m *ItemItemPullsPostRequestBody) GetBase()(*string) {
    return m.base
}
// GetBody gets the body property value. The contents of the pull request.
// returns a *string when successful
func (m *ItemItemPullsPostRequestBody) GetBody()(*string) {
    return m.body
}
// GetDraft gets the draft property value. Indicates whether the pull request is a draft. See "[Draft Pull Requests](https://docs.github.com/articles/about-pull-requests#draft-pull-requests)" in the GitHub Help documentation to learn more.
// returns a *bool when successful
func (m *ItemItemPullsPostRequestBody) GetDraft()(*bool) {
    return m.draft
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["head_repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadRepo(val)
        }
        return nil
    }
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val)
        }
        return nil
    }
    res["maintainer_can_modify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintainerCanModify(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetHead gets the head property value. The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.
// returns a *string when successful
func (m *ItemItemPullsPostRequestBody) GetHead()(*string) {
    return m.head
}
// GetHeadRepo gets the head_repo property value. The name of the repository where the changes in the pull request were made. This field is required for cross-repository pull requests if both repositories are owned by the same organization.
// returns a *string when successful
func (m *ItemItemPullsPostRequestBody) GetHeadRepo()(*string) {
    return m.head_repo
}
// GetIssue gets the issue property value. An issue in the repository to convert to a pull request. The issue title, body, and comments will become the title, body, and comments on the new pull request. Required unless `title` is specified.
// returns a *int64 when successful
func (m *ItemItemPullsPostRequestBody) GetIssue()(*int64) {
    return m.issue
}
// GetMaintainerCanModify gets the maintainer_can_modify property value. Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request.
// returns a *bool when successful
func (m *ItemItemPullsPostRequestBody) GetMaintainerCanModify()(*bool) {
    return m.maintainer_can_modify
}
// GetTitle gets the title property value. The title of the new pull request. Required unless `issue` is specified.
// returns a *string when successful
func (m *ItemItemPullsPostRequestBody) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *ItemItemPullsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
        err := writer.WriteStringValue("head", m.GetHead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("head_repo", m.GetHeadRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maintainer_can_modify", m.GetMaintainerCanModify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *ItemItemPullsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBase sets the base property value. The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository.
func (m *ItemItemPullsPostRequestBody) SetBase(value *string)() {
    m.base = value
}
// SetBody sets the body property value. The contents of the pull request.
func (m *ItemItemPullsPostRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetDraft sets the draft property value. Indicates whether the pull request is a draft. See "[Draft Pull Requests](https://docs.github.com/articles/about-pull-requests#draft-pull-requests)" in the GitHub Help documentation to learn more.
func (m *ItemItemPullsPostRequestBody) SetDraft(value *bool)() {
    m.draft = value
}
// SetHead sets the head property value. The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.
func (m *ItemItemPullsPostRequestBody) SetHead(value *string)() {
    m.head = value
}
// SetHeadRepo sets the head_repo property value. The name of the repository where the changes in the pull request were made. This field is required for cross-repository pull requests if both repositories are owned by the same organization.
func (m *ItemItemPullsPostRequestBody) SetHeadRepo(value *string)() {
    m.head_repo = value
}
// SetIssue sets the issue property value. An issue in the repository to convert to a pull request. The issue title, body, and comments will become the title, body, and comments on the new pull request. Required unless `title` is specified.
func (m *ItemItemPullsPostRequestBody) SetIssue(value *int64)() {
    m.issue = value
}
// SetMaintainerCanModify sets the maintainer_can_modify property value. Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request.
func (m *ItemItemPullsPostRequestBody) SetMaintainerCanModify(value *bool)() {
    m.maintainer_can_modify = value
}
// SetTitle sets the title property value. The title of the new pull request. Required unless `issue` is specified.
func (m *ItemItemPullsPostRequestBody) SetTitle(value *string)() {
    m.title = value
}
type ItemItemPullsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBase()(*string)
    GetBody()(*string)
    GetDraft()(*bool)
    GetHead()(*string)
    GetHeadRepo()(*string)
    GetIssue()(*int64)
    GetMaintainerCanModify()(*bool)
    GetTitle()(*string)
    SetBase(value *string)()
    SetBody(value *string)()
    SetDraft(value *bool)()
    SetHead(value *string)()
    SetHeadRepo(value *string)()
    SetIssue(value *int64)()
    SetMaintainerCanModify(value *bool)()
    SetTitle(value *string)()
}
