package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemWithPull_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository.
    base *string
    // The contents of the pull request.
    body *string
    // Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request.
    maintainer_can_modify *bool
    // The title of the pull request.
    title *string
}
// NewItemItemPullsItemWithPull_numberPatchRequestBody instantiates a new ItemItemPullsItemWithPull_numberPatchRequestBody and sets the default values.
func NewItemItemPullsItemWithPull_numberPatchRequestBody()(*ItemItemPullsItemWithPull_numberPatchRequestBody) {
    m := &ItemItemPullsItemWithPull_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemWithPull_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemWithPull_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemWithPull_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBase gets the base property value. The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository.
// returns a *string when successful
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) GetBase()(*string) {
    return m.base
}
// GetBody gets the body property value. The contents of the pull request.
// returns a *string when successful
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetMaintainerCanModify gets the maintainer_can_modify property value. Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request.
// returns a *bool when successful
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) GetMaintainerCanModify()(*bool) {
    return m.maintainer_can_modify
}
// GetTitle gets the title property value. The title of the pull request.
// returns a *string when successful
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBase sets the base property value. The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository.
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) SetBase(value *string)() {
    m.base = value
}
// SetBody sets the body property value. The contents of the pull request.
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetMaintainerCanModify sets the maintainer_can_modify property value. Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request.
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) SetMaintainerCanModify(value *bool)() {
    m.maintainer_can_modify = value
}
// SetTitle sets the title property value. The title of the pull request.
func (m *ItemItemPullsItemWithPull_numberPatchRequestBody) SetTitle(value *string)() {
    m.title = value
}
type ItemItemPullsItemWithPull_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBase()(*string)
    GetBody()(*string)
    GetMaintainerCanModify()(*bool)
    GetTitle()(*string)
    SetBase(value *string)()
    SetBody(value *string)()
    SetMaintainerCanModify(value *bool)()
    SetTitle(value *string)()
}
